package main

import (
	"fmt"
	"io/ioutil"
	//"os"
	//"path/filepath"
)


func main() {
	files, _ := ioutil.ReadDir("/home/liuhy/work/src/fabric-ca-demo/modelv4/app/enroll_user_peerorg1")
	//files, _ := ioutil.ReadDir("/")
	for _, f := range files {
		fmt.Println(f.Name(),f.Size())
	}
	fmt.Println(len(files))
}