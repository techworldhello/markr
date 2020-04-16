package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/techworldhello/markr/internal/data"
	"sync"
	"time"
)

type Store struct {
	Db *sql.DB
	mu sync.RWMutex
}

func New(db *sql.DB) *Store {
	return &Store{Db: db}
}

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/markr")
	if err != nil {
		log.Errorf("error connecting to db: %+v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Errorf("cannot ping db: %v", err)
		return nil, err
	}

	log.Print("connection established")
	return db, nil
}

func (s Store) SaveResults(data data.McqTestResults) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	txn, err := s.Db.Begin()
	if err != nil {
		log.Errorf("error starting db transaction: %v", err)
		return err
	}

	stmt, err := txn.Prepare(`INSERT INTO student_result 
(student_number, test_id, first_name, last_name, total_available, total_obtained, scanned_on, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		return err
	}

	for _, result := range data.Results {
		sqlSum, err := saveToTable(stmt, result)
		if err != nil {
			log.Errorf("error saving record: %v", err)
			return err
		}
		rs, err := sqlSum.RowsAffected()
		if rs != 1 || err != nil {
			log.Errorf("rows affected: %d\nerror: %v", rs, err)
			txn.Rollback()
			return err
		}
	}

	if err := txn.Commit(); err != nil {
		log.Errorf("error commiting db transaction: %v", err)
		return err
	}

	return nil
}

func (s Store) RetrieveMarks(testId string) ([]DBMarksRecord, error) {
	rows, err := s.Db.Query(`SELECT student_number, total_available, total_obtained FROM student_result 
WHERE test_id = ? ORDER BY student_number DESC, total_obtained DESC`, testId)

	if err != nil {
		log.Errorf("error querying DB for testId %s: %v", testId, err)
		return []DBMarksRecord{}, err
	}

	return getMarks(rows)
}

var Now, _ = time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")

func saveToTable(stmt *sql.Stmt, r data.TestResult) (sql.Result, error) {
	result, err := stmt.Exec(r.StudentNumber, r.TestID, r.FirstName, r.LastName, r.SummaryMarks.Available, r.SummaryMarks.Obtained, r.ScannedOn, Now)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type DBMarksRecord struct {
	StudentId, Available, Obtained int
}

func getMarks(rows *sql.Rows) (record []DBMarksRecord, err error) {
	defer rows.Close()

	for rows.Next() {
		var (
			studentNumber, marksAvailable, marksObtained int
		)

		if err := rows.Scan(&studentNumber, &marksAvailable, &marksObtained); err != nil {
			log.Errorf("error copying from columns: %v", err)
		}

		log.Infof("found row containing %d, %d, %d", studentNumber, marksAvailable, marksObtained)

		record = append(record, DBMarksRecord{studentNumber, marksAvailable, marksObtained})
	}

	if err = rows.Err(); err != nil {
		log.Errorf("error preparing row: %v", err)
		return record, err
	}

	return record, nil
}
