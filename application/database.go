package application

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDB(config *Config) *sql.DB {
	datasource := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.Mysql.User,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database,
	)
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	return db
}
