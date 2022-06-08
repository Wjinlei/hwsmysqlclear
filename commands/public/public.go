package public

import (
	"database/sql"
	"errors"
	"fmt"
)

/* Save connect info */
type connect struct {
	db     *sql.DB
	dbUser string
	dbPass string
	dbName string
}

var conn connect

func Connect(dbUser, dbPass, dbName string) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName))
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	SetConnect(db, dbUser, dbPass, dbName)
	return nil
}

func SetConnect(db *sql.DB, dbUser, dbPass, dbName string) {
	conn.db = db
	conn.dbUser = dbUser
	conn.dbPass = dbPass
	conn.dbName = dbName
}

func GetConnect() (*connect, error) {
	if conn.db != nil {
		return &conn, nil
	}
	return nil, errors.New("GetConnect fail. db is nil")
}

func (conn *connect) Close() {
	if conn.db != nil {
		conn.db.Close()
	}
}
