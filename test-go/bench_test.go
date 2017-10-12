package test

import (
	"strconv"
	"testing"
//	"fmt"
)


func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(i) // op
	}
}

/*
?   	mywriteCode/test-go	[no test files]
是因为文件名为：　testb.go 修改 ***_test.go 则运行正常


liuhy@liuhy ~/work/src/mywriteCode/test-go $ go test -run $^ -bench Itoa -benchmem -cpu 2
goos: linux
goarch: amd64
pkg: mywriteCode/test-go
BenchmarkItoa-2   	30000000	        49.7 ns/op	       7 B/op	       0 allocs/op
PASS
ok  	mywriteCode/test-go	1.542s

liuhy@liuhy ~/work/src/mywriteCode/test-go $ go test -bench=.
goos: linux
goarch: amd64
pkg: mywriteCode/test-go
BenchmarkItoa-4   	30000000	        48.9 ns/op
PASS
ok  	mywriteCode/test-go	1.520s
liuhy@liuhy ~/work/src/mywriteCode/test-go $





*/