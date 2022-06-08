package public

import "database/sql"

func GetTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return []string{}, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return []string{}, err
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	tables := []string{}

	// Fetch rows
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return []string{}, err
		}

		for _, col := range values {
			if col != nil {
				tables = append(tables, string(col))
			}
		}
	}
	if err = rows.Err(); err != nil {
		return tables, err
	}
	return tables, nil
}
