package mysql

import (
	// 导入golang提供的操作数据库的标准接口
	"database/sql"
	"fmt"
	"log"
	"os"

	// 匿名导入mysql的数据库驱动，自行初始化并注册自己到Golang的database/sql上下文中, 因此我们就可以通过 database/sql 包提供的方法访问数据库了.
	_ "github.com/go-sql-driver/mysql"
)

// 创建一个db连接对象
var db *sql.DB

// 初始化
func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	// 设置活跃连接数
	db.SetMaxOpenConns(1000)
	// 数据库连接测试
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}

// DBConn : 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}

// ParseRows : 数据封装
func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

// checkErr : 数据校验
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
