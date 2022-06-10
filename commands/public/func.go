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
		whereId := columns[0]
		whereIdVal := callback(0)
		for i := range columns {
			columnValue := callback(i)
			for _, findCase := range regularizer.FindAllString(columnValue, -1) {
				conn.update(table, whereId, whereIdVal, columns[i], columnValue, findCase, "")
			}
		}
	}
	return nil
}

// Update oldCase to newCase
func (conn *connect) update(table string, whereId, whereIdVal string, field, value string, oldCase string, newCase string) error {
	stmt, err := conn.db.Prepare("UPDATE `" + table + "` SET `" + field + "` = REPLACE(`" + field + "`, ?, ?) WHERE `" + whereId + "` = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(oldCase, newCase, whereIdVal)
	if err != nil {
		return err
	}

	// Write Log
	golib.FileWrite(
		Logfile,
		fmt.Sprintf("[%s] 表: %s\t字段: %s\t标识: %s\t原内容: %s\n",
			golib.GetNowTime(), table, field, whereId+"="+whereIdVal, value),
		golib.FileAppend)

	return nil
}
