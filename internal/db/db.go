package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/techworldhello/markr/internal/data"
	"log"
	"time"
)

type Database interface {
	Save(m data.McqTestResults) error
	Get(testID string) (string, error)
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
		log.Printf("error connecting to db: %+v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Printf("cannot ping db: %v", err)
		return nil, err
	}

	log.Print("connection established")
	return db, nil
}

func (s Store) Save(m data.McqTestResults) error {
	txn, err := s.Db.Begin()
	if err != nil {
		log.Printf("error starting db transaction: %v", err)
		return err
	}

	stmt, err := txn.Prepare(`INSERT INTO student_result 
(student_number, test_id, first_name, last_name, total_available, total_obtained, scanned_on, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		log.Print(err)
		return err
	}

	for _, result := range m.Results {
		sqlSum, err := saveToTable(stmt, result)
		if err != nil {
			return err
		}
		rs, err := sqlSum.RowsAffected()
		if rs != 1 || err != nil {
			log.Printf("rows affected: %d\nerror: %v", rs, err)
			txn.Rollback()
			return err
		}
	}

	if err := txn.Commit(); err != nil {
		log.Printf("error commiting db transaction: %v", err)
		return err
	}

	return nil
}

func (s Store) Get(testID string) (string, error) {
	return `{"mean":65.0,"stddev":0.0,"min":65.0,"max":65.0,"p25":65.0,"p50":65.0,"p75":65.0,"count":1}`, nil
}

var Now = time.Now()

func saveToTable(stmt *sql.Stmt, r *data.TestResult) (sql.Result, error) {
	result, err := stmt.Exec(r.StudentNumber, r.TestID, r.FirstName, r.LastName, r.SummaryMarks.Available, r.SummaryMarks.Obtained, r.ScannedOn, Now)
	if err != nil {
		return nil, err
	}
	return result, nil
}
