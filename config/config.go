package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MySQL *sql.DB
	Error error
	Hola  string
)

func ConnectDataBase() {

	//adminuno
	//

	if os.Getenv("DB_USER") == "" {
		os.Setenv("DB_USER", "XdxDzKZ6Pz")
		os.Setenv("DB_PASSWORD", "vBZnwIpcj0")
		os.Setenv("DB_NAME", "XdxDzKZ6Pz")
		os.Setenv("SERVER_NAME", "remotemysql.com:3306")
	}

	var CONN_STRING = fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("SERVER_NAME"), os.Getenv("DB_NAME"))
	MySQL, Error = sql.Open("mysql", CONN_STRING)

	Error = MySQL.Ping()
	if Error != nil {
		fmt.Print(Error.Error())
	}

}
