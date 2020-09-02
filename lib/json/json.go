package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
)

func jsontest1() {
	person := Person{
		Id:   1001,
		Name: "张三",
		Age: sql.NullInt64{
			Int64: 33,
			Valid: true,
		},
		Hobby: sql.NullString{String: "游泳", Valid: true},
	}

	marshal, err := json.Marshal(&person)
	fmt.Println(err)
	fmt.Println(string(marshal))
}

func jsontest2() {
	str := `{"Id":1001,"Name":"张三","Age": 34,"Hobby":[{"id":1001,"name":"fuck"},{"id":1002,"name":"fuck is fuck"}]}`

	person := Person{}
	json.Unmarshal([]byte(str), &person)

	fmt.Println(person)

	value, dataType, offset, err := jsonparser.Get([]byte(str), "Hobby")

	fmt.Println(string(value))
	fmt.Println(dataType)
	fmt.Println(offset)
	fmt.Println(err)

}

func main() {
	jsontest1()
	jsontest2()

}
