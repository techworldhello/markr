package db

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/techworldhello/markr/pkg/data"
	"log"
	"testing"
)

func TestSave(t *testing.T) {
	sqlDb, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("error opening stub database connection: %v", err)
	}
	defer sqlDb.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO student_result")
	mock.ExpectExec("INSERT INTO student_result").
		WithArgs(1234, 1, "Daniel", "Craig", 20, 18, data.ScannedTime, Now).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO student_result").
		WithArgs(4321, 2, "Jane", "Doe", 20, 10, data.ScannedTime, Now).
		WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectExec("INSERT INTO student_result").
		WithArgs(1212, 3, "Spongebob", "Squarepants", 20, 14, data.ScannedTime, Now).
		WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()

	c := Store{Db: sqlDb}
	if err := c.Save(data.GetTestResults()); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}