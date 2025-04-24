package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var once sync.Once

func InitDB() {
	once.Do(func() {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
			Db_user, DB_Pass, Db_host, Db_port, Db_name)
		DB, err := sql.Open("mysql", dsn)

		if err != nil {
			log.Fatal("DB Connection Error: ", err)
		}

		pingErr := DB.Ping()
		if pingErr != nil {
			log.Fatal("DB ping error: ", pingErr)
		}

		fmt.Println("DB Connected!")
	})
}
