package test

import (
	"flag"
	"log"
	"os"
	"testing"
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


func TestMMM(t *testing.T) {
	log.Println("[Test1] running ", *wordPtr)
	t.Logf("%s","Test Main")
}

//测试文件的不能test开头，但是必须算是以test结尾


/*
func TestMain(m *testing.M) {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "/tmp")
	flag.Set("v", "3")
	flag.Parse()

	ret := m.Run()
	os.Exit(ret)
}

*/