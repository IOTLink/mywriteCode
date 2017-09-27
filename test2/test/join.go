package main

import(
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	strs = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}
)

func TestStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(strs, "")
	}
}

func TestStringsPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < len(strs); j++ {
			s += strs[j]
		}
	}
}

func TestBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		for j := 0; j < len(strs); j++ {
			b.WriteString(strs[j])
		}
	}
}
func TestByte(b *testing.B){
	var byteArray []byte
	for i :=0; i< b.N; i++ {
		for j := 0; j < len(strs); j++ {
			byteArray = append(byteArray,[]byte(strs[j])...)
		}
	}
	_ = string(byteArray)
}

func TestSprintf(b *testing.B){
	var byteArray string
	for i :=0; i< b.N; i++ {
		for j := 0; j < len(strs); j++ {
			byteArray += fmt.Sprintf("%s",strs[j])
		}
	}
}

func main() {
	fmt.Println("strings.Join:")
	fmt.Println(testing.Benchmark(TestStringsJoin))
	fmt.Println("bytes.Buffer:")
	fmt.Println(testing.Benchmark(TestBytesBuffer))

	fmt.Println("bytes.bytes:")
	fmt.Println(testing.Benchmark(TestByte))

	fmt.Println("bytes.TestSprintf:")
	fmt.Println(testing.Benchmark(TestSprintf))


	fmt.Println("+:")
	fmt.Println(testing.Benchmark(TestStringsPlus))
}
