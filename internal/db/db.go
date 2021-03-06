package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/techworldhello/markr/internal/data"
	"time"
)

type Store struct {
	Db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s Store) SaveResults(data data.McqTestResults) error {
	txn, err := s.Db.Begin()
	if err != nil {
		log.Errorf("error starting db transaction: %v", err)
		return err
	}

	stmt, err := txn.Prepare(`INSERT INTO student_results 
(student_number, test_id, first_name, last_name, total_available, total_obtained, scanned_on, created_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	for i, result := range data.Results {
		sqlSum, err := saveToTable(stmt, result)
		if err != nil {
			log.Errorf("error saving record no.%d: %v", i, err)
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
	rows, err := s.Db.Query(`SELECT student_number, total_available, total_obtained FROM student_results 
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
	StudentId           string
	Available, Obtained int
}

func getMarks(rows *sql.Rows) (record []DBMarksRecord, err error) {
	defer rows.Close()

	for rows.Next() {
		var (
			studentNumber                 string
			marksAvailable, marksObtained int
		)

		if err := rows.Scan(&studentNumber, &marksAvailable, &marksObtained); err != nil {
			log.Errorf("error copying from columns: %v", err)
		}

		log.Debugf("found row containing %s, %d, %d", studentNumber, marksAvailable, marksObtained)

		record = append(record, DBMarksRecord{studentNumber, marksAvailable, marksObtained})
	}

	if err = rows.Err(); err != nil {
		log.Errorf("error preparing row: %v", err)
		return record, err
	}

	return record, nil
}
