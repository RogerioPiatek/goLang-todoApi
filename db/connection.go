package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rogeriopiatek/goLang-todoAPI/configs"
)

// Opens the connection with the database
func OpenConnection() (*sql.DB, error) {
	//receives the Database configs
	conf := configs.GetDB()

	//mount the connection string based on the configs
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	//oppens a connection with a postgres database
	conn, err := sql.Open("postgres", sc)
	if err != nil {
		//if on production, panic is not good
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
