package repo

import (
	"fortress/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

var lock = &sync.Mutex{}
var connection *model.Connection

func GetConnection() (*model.Connection, error) {
	if connection == nil {
		lock.Lock()
		defer lock.Unlock()

		db, err := sqlx.Open("sqlite3", "fortress.sqlite")
		if err != nil {
			log.Fatalln("Failed to connect to internal Fortress db: " + err.Error())
			return nil, err
		} else {
			log.Println("Internal Fortress database connection established")
		}

		connection = &model.Connection{
			Db: db,
		}

		InitialiseDatabase(connection)
	}

	return connection, nil
}

func InitialiseDatabase(db *model.Connection) {
	stmt := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			name_first TEXT NOT NULL,
			name_last TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := db.Db.Exec(stmt)
	if err != nil {
		log.Fatalln("Failed to initialise database: " + err.Error())
	}
}
