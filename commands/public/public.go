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

func (conn *connect) QueryRows(querySql string) (rows *sql.Rows, columns []string, callback func(columnIndex int) string, err error) {
	rows, err = conn.db.Query(querySql)
	if err != nil {
		return nil, nil, nil, err
	}

	columns, err = rows.Columns()
	if err != nil {
		return nil, nil, nil, err
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	valuesReferences := make([]interface{}, len(values))
	for i := range values {
		valuesReferences[i] = &values[i]
	}

	// Return the Closure function
	return rows, columns, func(columnIndex int) string {
		if columnIndex > len(values)-1 {
			return fmt.Sprintf("error: Bad column index %d", columnIndex)
		}
		if err := rows.Scan(valuesReferences...); err != nil {
			return err.Error()
		}
		return string(values[columnIndex])
	}, nil
}

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
