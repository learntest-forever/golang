package main

import (
	"database/sql"
	 _"github.com/go-sql-driver/mysql"
	"fmt"
)

func main(){
	db,err := sql.Open("mysql","root:@tcp(10.226.133.68:3306)/db_dns_controller")
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
	rows,err := db.Query("select * from dns_domain limit 1;")
	defer rows.Close()
	fmt.Printf("the rows is : %v\n",rows)
	for rows.Next() {
		columns, _ := rows.Columns()
		fmt.Printf("the columns is : %v\n",columns)
		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
			fmt.Printf("scan is : %v\n",scanArgs[i])
			fmt.Printf("valu is : %v\n",values[i])
		}
		fmt.Printf("the scan mmmmm %v\n",values)
		err = rows.Scan(scanArgs...)
		fmt.Printf("the scan nnnnn %v\n",values)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	rows.Close()
}