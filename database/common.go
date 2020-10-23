package database

import (
	"database/sql"
	"fmt"
	"github.com/CardinalDevLab/Schwi-Backend/utils"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Database *sql.DB
	err      error
)

func InitDatabase() {
	Database, err = sql.Open("sqlite3", "./database.db")
	utils.CheckErr(err)
	fmt.Println("Init Database")
}
