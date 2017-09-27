package main

import (
	s "strings"
	"fmt"
	"reflect"

	"runtime"
	strings "strings"
	"unicode"

)
func main() {

	fmt.Println("Contains:  ", s.Contains("test", "es"))
	fmt.Println("Count:     ", s.Count("test", "t"))
	fmt.Println("HasPrefix: ", s.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", s.HasSuffix("test", "st"))
	fmt.Println("Index:     ", s.Index("test", "e"))
	fmt.Println("Join:      ", s.Join([]string{"a", "b"}, "-"))
	fmt.Println("Join:      ", s.Join([]string{"a", "b"}, ""))
	fmt.Println("Repeat:    ", s.Repeat("a", 5))
	fmt.Println("Replace:   ", s.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", s.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", s.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", s.ToLower("TEST"))
	fmt.Println("ToUpper:   ", s.ToUpper("test"))
	fmt.Println()
	fmt.Println("Len: ", len("hello"))
	fmt.Println("Char:", "hello"[1])
	fmt.Println("Char:", "hello"[0])

	str := fmt.Sprintf("%s%s", "abc", "def")
	fmt.Println(str[:])
	fmt.Println(str)

	array := s.FieldsFunc("my, hello, world ", split)
	fmt.Println(array)

	var hello = "hello"

	// 104 is the ascii code of char 'h'
	fmt.Println(hello[0]) // 104
	fmt.Println(reflect.TypeOf(hello[0])) // uint8 (byte)

	fmt.Println(hello == "hello")    // true
	fmt.Println("helloa" > "helloWorld") // false

	s2 := "éक्षिaπ汉字"
	for i, rn := range s2 {
		fmt.Printf("%d: 0x%xd %s \n", i, rn, string(rn))
	}

	var strtest = "hello world ........"
	var inf interface{}
	inf = strtest
	fmt.Println(inf)

	ss,ok := inf.(string)
	if ok {
		fmt.Println(ss)
	}

	stringTest := "hello world!"
	for i,b := range []byte(stringTest){
		fmt.Println(i,b,string(b))
	}

	key := []byte{'k','e','v'}
	m := map[string]string{}
	m = make(map[string]string)
	m[string(key)] = "value"
	fmt.Println(m)

	var s1 string
	var x1 = []byte{1024: 'x','1','2'}
	var y1 = []byte{1024: 'y'}
	s1 = string(x1) + string(y1)
	fmt.Println(s1)

	bs := make([]byte, 1 << 20) // (1 << 20) is 1M
	bs = []byte{'0','1','2'}
	fmt.Println(string(bs))

	runtime.GC()

	hello1 := []byte{'h','e','l','l','o',' '}
	world1 := "world"
	helloworld := append(hello1, world1...)
	fmt.Println(string(helloworld))

	helloworld2 := make([]byte, len(hello1)+len(world1))
	n := copy(helloworld2,hello1)
	fmt.Println(string(helloworld2))
	copy(helloworld2[n:],world1)
	fmt.Println(string(helloworld2))


	fmt.Println(s.ToTitle("myTESTKKxxxsss"))

	fmt.Println(s.EqualFold("hello","hellos"))
	fmt.Println("+++")
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println(strings.Index("chicken", "ken"))

	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))

	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("gophers", 'h'))
	fmt.Println(strings.IndexByte("golang", 'x'))

	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))

	fmt.Println("ba" + strings.Repeat("na", 2))

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", -1))

	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))

	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))

	var s3 = "Hello, goodbye, etc!"
	s3 = strings.TrimSuffix(s3, "goodbye, etc!")
	fmt.Println(s3)

	var s4 = "Goodbye,, world!"
	s4 = strings.TrimPrefix(s4, "Goodbye,")
	fmt.Println(s4)

	f1 := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("__  Afoo1;bar2,baz3...", f1))
}

func split(s rune) bool {
	if s == ',' {
		return true
	}
	return false
}
