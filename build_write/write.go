package main


import "fmt"
import "flag"

var (
	version = "v0.0.1"
	build   = "not set"
)

func main() {

	var v = flag.Bool("v", false, "display version")

	flag.Parse()

	if *v {
		fmt.Printf("version: %s\n", version)
		fmt.Printf("build  : %s\n", build)
	} else {
		fmt.Println("hello, test version")
	}
}

/*

http://www.cnblogs.com/wang_yb/p/6163326.html

通过 ldflags， 动态修改 build

go build -ldflags "-X main.build=`git rev-parse HEAD`" main.go

 ./main -v
version: v0.0.1
build  : 23af559af14a0b83c35a7f8bd0670a9741b1dc7e
-X main.build 修改 main package 中的 build 变量
`git rev-parse HEAD` 获取当前的 git revision number






liuhy@liuhy ~/work/src/mywriteCode/build_write $ go build -ldflags "-X main.build=`git rev-parse HEAD`" write.go
liuhy@liuhy ~/work/src/mywriteCode/build_write $ ./write -v
version: v0.0.1
build  : 2b5534649a369ed5d442ae07d7ee82d3c2ac7b03


 */