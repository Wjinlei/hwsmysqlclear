package public

import (
	"database/sql"
)

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

	return rows, columns,
		func(columnIndex int) string {
			if columnIndex > len(values)-1 {
				return ""
			}
			if err := rows.Scan(valuesReferences...); err != nil {
				return ""
			}
			return string(values[columnIndex])
		}, nil
}
