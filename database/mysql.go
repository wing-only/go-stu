package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
连接数据库
增删改
*/
func mysqltest1() {
	db, _ := sql.Open("mysql", "admin:admin@(127.0.0.1:3306)/test")
	defer db.Close()

	err := db.Ping()

	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	sql := "insert into people(person_id, first_name, last_name) values(1003, 'wing', 'wang')"
	result, _ := db.Exec(sql)
	affected, _ := result.RowsAffected()
	fmt.Println("受影响行数:", affected)
}

/*
预处理
增、删、改
*/
func mysqltest2() {

	db, _ := sql.Open("mysql", "admin:admin@(127.0.0.1)/test")
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("连接数据库失败")
	}

	p := [2][3]string{{"2001", "keyy", "tom"}, {"2002", "jim", "green"}}

	stmt, _ := db.Prepare("insert into people(person_id, first_name, last_name) values(?,?,?)")
	defer stmt.Close()

	for _, s := range p {
		stmt.Exec(s[0], s[1], s[2])
	}

}

/*
单行查询
*/
func mysqltest3() {
	db, _ := sql.Open("mysql", "admin:admin@(127.0.0.1)/test")
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("连接数据库失败")
	}

	var id int64
	var firstName string
	var lastName string
	stmt, _ := db.Prepare("select * from people where person_id = ?")
	defer stmt.Close()

	row := stmt.QueryRow(1002)
	row.Scan(&id, &firstName, &lastName)

	fmt.Println(id, firstName, lastName)
}

/*
多行查询
*/
func mysqltest4() {
	db, _ := sql.Open("mysql", "admin:admin@(127.0.0.1)/test")
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("连接数据库失败")
	}

	var id int64
	var firstName string
	var lastName string
	stmt, _ := db.Prepare("select * from people")
	defer stmt.Close()

	rows, _ := stmt.Query()
	for rows.Next() {
		rows.Scan(&id, &firstName, &lastName)
		fmt.Println(id, firstName, lastName)
	}
}

/*
事务
*/
func mysqltest5() {

}

/*
连接池
*/
func mysqltest6() {

}

/*

 */
func mysqltest7() {

}

func main() {
	//mysqltest1()
	//mysqltest2()
	//mysqltest3()
	mysqltest4()
}
