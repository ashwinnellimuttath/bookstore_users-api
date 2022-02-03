package users_db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	//_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Client *sql.DB
)

func init() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "users_db",
		AllowNativePasswords: true,
	}
	//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
	//	"root",
	//	"password",
	//	"localhost:3306",
	//	"users_db",
	//)
	var err error
	Client, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	if pingErr := Client.Ping(); pingErr != nil {
		panic(pingErr)
	}
	log.Println("successfully configured")
}
