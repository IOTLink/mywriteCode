package main


import (
	"fmt"
	//"encoding/json"
	"reflect"
)
type User struct {
	 Name   string "user name" //这引号里面的就是tag
	 Passwd string "user passsword"
}

func test1() {
	user := &User{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}
}

func test2() {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}

func test3() {
	type Int64Struct struct {
		Int64Var int64 `def:"3546343826724305832" help:"int64"`
	}

	s := Int64Struct{}
	st := reflect.TypeOf(s)
	fmt.Println("_________________________________")
	fmt.Println(st)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("def"), field.Tag.Get("help"))

	fmt.Println()
}
func main() {
	test1()
	test2()
	test3()
}

/*

从上面的代码可看出，结构体间的数据转换可以不用严格遵循一对一的转换，例如Person Struct 中定义的MiddleName，如果json数据定义中无此字段，可以在StructTag中加入”omitempty”, 标识该字段的数据可忽略。

http://wangwei.info/golang-struct-tag/
 */