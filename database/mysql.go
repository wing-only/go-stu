package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//连接数据库
func mysqltest1() (db *sql.DB) {
	db, _ = sql.Open("mysql", "admin:admin@(127.0.0.1:3306)/test")
	defer db.Close()

	err := db.Ping()

	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	return
}

//
func mysqltest2() {
	db := mysqltest1()

	sql := "insert into people(person_id, first_name, last_name) values(1001, 'wing', 'wang')"
	result, _ := db.Exec(sql)
	affected, _ := result.RowsAffected()
	fmt.Println("受影响行数:", affected)
}

//
func mysqltest3() {

}

//
func mysqltest4() {

}

//
func mysqltest5() {

}

//
func mysqltest6() {

}

//
func mysqltest7() {

}

func main() {
	//mysqltest1()
	mysqltest2()
}
