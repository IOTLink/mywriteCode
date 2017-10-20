package test

import "testing"

import (
	"flag"
	"log"
	"os"

	"fmt"
)

var wordPtr = flag.String("word", "foo", "a string")

func TestMain(m *testing.M) {


	flag.Parse()
	log.Println("[TestMain] word:", *wordPtr)

	log.Println("[TestMain] run()前")
	exitVal := m.Run()
	log.Println("[TestMain] run()后")


	os.Exit(exitVal)

	fmt.Println("test Main")
}



func TestM_main(t *testing.T) {
	log.Println("[Test1] running ", *wordPtr)
}



func Reverse(str string) string {
	rs := []rune(str)
	len := len(rs)
	var tt []rune

	tt = make([]rune, 0)
	for i := 0; i < len; i++ {
		tt = append(tt, rs[len-i-1])
	}
	return string(tt[0:])
}

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}



/*

go test -v
=== RUN   TestReverse
--- PASS: TestReverse (0.00s)
PASS
ok  	mywriteCode/test-go	0.001s




 */