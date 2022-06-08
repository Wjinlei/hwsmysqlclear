package public

import (
	"database/sql"
)

func (conn *connect) QueryRows(querySql string) (*sql.Rows, func(columnIndex int) string, error) {
	rows, err := conn.db.Query(querySql)
	if err != nil {
		return nil, nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	return rows, func(columnIndex int) string {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return ""
		}
		if columnIndex > len(values)-1 {
			return ""
		}
		return string(values[columnIndex])
	}, nil
}
