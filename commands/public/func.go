package public

import (
	"fmt"
	"regexp"
)

func (conn *connect) FindScript(table string) error {
	rows, cols, callback, err := conn.QueryRows("SELECT * FROM `" + table + "`")
	if err != nil {
		return err
	}

	regularizer := regexp.MustCompile(`(?i)<script.*(</script[^>]*>)?`)

	for rows.Next() {
		for i := range cols {
			columnValue := callback(i)
			findAllCase := regularizer.FindAllString(columnValue, -1)
			for _, findCase := range findAllCase {
				conn.replace(table, cols[i], columnValue, findCase, "")
			}
		}
	}
	return nil
}

func (conn *connect) replace(table string, column string, value string, oldCase string, newCase string) error {
	query := fmt.Sprintf("UPDATE `%s` SET `%s` = REPLACE(`%s`, ?, ?) WHERE `%s` = ?", table, column, column, column)

	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(oldCase, newCase, value)
	if err != nil {
		return err
	}
	return nil
}
