package public

import (
	"fmt"

	"github.com/Wjinlei/golib"
)

func (conn *connect) FindScript(table string) error {
	rows, columns, callback, err := conn.QueryRows("SELECT * FROM `" + table + "`")
	if err != nil {
		return err
	}

	for rows.Next() {
		// Take the 0th field as the unique identifier
		id := fmt.Sprintf("%s=%s", columns[0], callback(0))

		for i := range columns {
			value := callback(i)
			for _, findCase := range regularizer.FindAllString(value, -1) {
				conn.replace(table, id, columns[i], value, findCase, "")
			}
		}
	}
	return nil
}

func (conn *connect) replace(table string, id string, column string, value string, oldCase string, newCase string) error {
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

	// Log
	content := fmt.Sprintf("[%s] 表: %s\t字段: %s\t标识: %s\t原内容: %s\n", golib.GetNowTime(), table, column, id, value)
	golib.FileWrite(Logfile, content, golib.FileAppend)
	return nil
}
