package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // dialecto postgres
)

// Database db instance
var Database *gorm.DB = New()

// New creates a new connection to the database
func New() *gorm.DB {
	password := os.Getenv("password")
	db, err := gorm.Open("postgres", "host="+os.Getenv("DATABASE_URL")+
		" user=postgres dbname=galaxia password="+password)
	if err != nil {
		fmt.Println("bd connection err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	// me aseguro de cerrar la conexión si el servidor se cierra
	defer db.Close()
	return db
}
