package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/techworldhello/markr/internal/data"
	"os"
	"time"
)

type Database interface {
	Save(m data.McqTestResults) error
	RetrieveScores(testId string) ([]float64, error)
}

type Store struct {
	Db *sql.DB
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

func (s Store) Save(m data.McqTestResults) error {
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

	for _, result := range m.Results {
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

func (s Store) RetrieveScores(testId string) ([]float64, error) {
	rows, err := s.Db.Query(fmt.Sprintf("SELECT * FROM %s WHERE test_id = ?", os.Getenv("DB_NAME")), testId)
	if err != nil {
		log.Errorf("error querying DB for testId %s: %v", testId, err)
		return []float64{}, err
	}
	return getScores(rows)
}

func getScores(rows *sql.Rows) (scores []float64, err error) {
	defer rows.Close()
	var tr data.TestResult
	for rows.Next() {
		if err := rows.Scan(&tr.Id, &tr.ScannedOn, &tr.StudentNumber, &tr.FirstName, &tr.LastName,
			&tr.TestID, &tr.SummaryMarks.Available, &tr.SummaryMarks.Obtained, &tr.CreatedAt); err != nil {
			log.Printf("error copying from columns: %v", err)
		}
		scores = append(scores, float64(tr.SummaryMarks.Obtained))
	}
	return scores, nil
}

var Now = time.Now()

func saveToTable(stmt *sql.Stmt, r *data.TestResult) (sql.Result, error) {
	result, err := stmt.Exec(r.StudentNumber, r.TestID, r.FirstName, r.LastName, r.SummaryMarks.Available, r.SummaryMarks.Obtained, r.ScannedOn, Now)
	if err != nil {
		return nil, err
	}
	return result, nil
}
