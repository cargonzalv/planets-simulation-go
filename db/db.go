package db

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // dialecto postgres
)

// Database db instance
var Database *gorm.DB

// Init creates a new connection to the database
func Init() {
	password := os.Getenv("password")
	dbConfig := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		os.Getenv("DATABASE_URL"),
		"postgres",
		"galaxia",
		"disable",
		password,
	)
	db, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		fmt.Println("bd connection err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)

	Database = db

	// auto-connect，ping per 60s, re-connect on fail or error with intervels 3s, 3s, 15s, 30s, 60s, 60s ...
	go func(dbConfig string) {
		var intervals = []time.Duration{3 * time.Second, 3 * time.Second, 15 * time.Second, 30 * time.Second, 60 * time.Second}
		for {
			time.Sleep(60 * time.Second)
			if e := Database.DB().Ping(); e != nil {
			L:
				for i := 0; i < len(intervals); i++ {
					e2 := RetryHandler(3, func() (bool, error) {
						var e error
						Database, e = gorm.Open("postgres", dbConfig)
						if e != nil {
							return false, errors.Unwrap(e)
						}
						return true, nil
					})
					if e2 != nil {
						fmt.Println(e.Error())
						time.Sleep(intervals[i])
						if i == len(intervals)-1 {
							i--
						}
						continue
					}
					break L
				}

			}
		}
	}(dbConfig)
}

// RetryHandler Try f() n times on fail and one time on success
func RetryHandler(n int, f func() (bool, error)) error {
	ok, er := f()
	if ok && er == nil {
		return nil
	}
	if n-1 > 0 {
		return RetryHandler(n-1, f)
	}
	return er
}
