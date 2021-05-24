package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// package-level variable
var db *sql.DB

// package init function
func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/entry_task?charset=utf8")
	if db == nil {
		fmt.Println("open mysql failed.")
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(100)
	err := db.Ping()
	if err != nil {
		fmt.Printf("Cannot link to mysql. please check it out: %s\n", err.Error())
		os.Exit(1)
	}
}

// DBConn expose mysql db connection to out-of-package
func DBConn() *sql.DB {
	return db
}

// RecordRow alias for row data
type RecordRow = map[string]interface{}

// ParseRow parse row records queried from mysql
func ParseRow(rows *sql.Rows) []RecordRow {
	colums, _ := rows.Columns()
	scanArgs := make([]interface{}, len(colums))
	//values := make([]interface{}, len(colums))

	//for j := range values {
	//	scanArgs[j] = &values[j]
	//}

	record := make(RecordRow)
	records := make([]RecordRow, 0)

	for rows.Next() {
		if err := rows.Scan(scanArgs); err != nil {
			log.Fatal(err)
			panic(err)
		}
		for i, colValue := range scanArgs {
			if colValue != nil {
				record[colums[i]] = colValue
			}
		}
		records = append(records, record)
	}
	return records
}
