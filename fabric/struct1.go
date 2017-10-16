package main


import (
"fmt"
"encoding/json"
)
type Person struct {
	FirstName string `json:"first_name"` //FirstName <=> firest_name
	LastName string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
}
func main() {
	json_string := ` { "first_name": "John", "last_name": "Smith" }`
	person := new(Person)
	json.Unmarshal([]byte(json_string), person) //将json数据转为Person Struct
	fmt.Println(person)
	new_json, _ := json.Marshal(person) //将Person Sturct 转为json格式
	fmt.Printf("%s\n", new_json)
}

// *Output*
// &{John Smith }
// {"first_name":"John","last_name":"Smith"}

/*

从上面的代码可看出，结构体间的数据转换可以不用严格遵循一对一的转换，例如Person Struct 中定义的MiddleName，如果json数据定义中无此字段，可以在StructTag中加入”omitempty”, 标识该字段的数据可忽略。

http://wangwei.info/golang-struct-tag/
 */