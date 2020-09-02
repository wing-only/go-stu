package main

import (
	"database/sql"
	"fmt"
	ffjson "github.com/pquerna/ffjson/ffjson"
)

type Person struct {
	Id    int64
	Name  string
	Age   sql.NullInt64  `json:"Age"`
	Hobby sql.NullString `json:"Hobby"`
}

func ffjsontest1() {

	person := Person{
		Id:   1001,
		Name: "张三",
		Age: sql.NullInt64{
			Int64: 33,
			Valid: true,
		},
		Hobby: sql.NullString{String: "游泳", Valid: true},
	}

	bytes, _ := ffjson.Marshal(&person)
	fmt.Println(string(bytes))

}

func ffjsontest2() {
	str := `{"Id":1001,"Name":"张三","Age": 34,"Hobby":[{"id":1001,"name":"fuck"},{"id":1002,"name":"fuck is fuck"}]}`

	person := Person{}
	ffjson.Unmarshal([]byte(str), &person)

	fmt.Println(person)
}

func main() {
	//ffjsontest1()
	ffjsontest2()
}
