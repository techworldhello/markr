package db

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("MYSQL_HOST_URL"),
		os.Getenv("MYSQL_HOST_PORT"),
		os.Getenv("MYSQL_DATABASE")))

	if err != nil {
		log.Errorf("error connecting to db: %+v", err)
		return nil, err
	}

	if err = pingDB(db); err != nil {
		return nil, err
	}
	log.Infof("connection established")
	return db, nil
}

func pingDB(db *sql.DB) (err error) {
	var (
		count         int
		maxRetryTimes = 5
	)
	for count < maxRetryTimes {
		if err = db.Ping(); err != nil {
			log.Warn("cannot ping db, trying again in 3 seconds..")
			time.Sleep(3 * time.Second)
			count ++
		} else {
			break
		}
	}
	if count == maxRetryTimes {
		log.Errorf("cannot ping db: %v", err)
		return err
	}
	return nil
}
