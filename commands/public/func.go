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

	// Fetch data
	for rows.Next() {
		for i := range cols {
			findCase := regularizer.FindAllString(callback(i), -1)
			for _, s := range findCase {
				fmt.Println(s)
			}
		}
	}
	return nil
}
