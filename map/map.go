package main

import (
	"fmt"
	"sync"
	"sort"
)

type Vertex struct {
	Lat, Long float64
}


func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	m1 :=  map[string]Vertex{}
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m1["Bell Labs"])


	m2 := make(map[string]Vertex)
	m2["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m2["Bell Labs"])

	var m3 map[string]Vertex
	m3 = make(map[string]Vertex)
	m3["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m3["Bell Labs"])

	for key, value := range m3 {
		fmt.Println("Key:", key, "Value:", value)
	}

	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	for key, value := range commits {
		fmt.Println("Key:", key, "Value:", value)
	}
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}

func test1(){

	type Node struct {
		Next  *Node
		Value interface{}
	}
	var first *Node = nil

	for i := 0; i< 10; i++{
		item := &Node{Next:nil,Value:i}
		if first == nil {
			first = item
		} else {
			next := first
			item.Next = next
			first = item
		}
	}

	visited := make(map[*Node]bool)
	for n := first; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}

	for key, value := range visited {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func test2() {
	type Person struct {
		Name  string
		Likes []string
	}
	var people []*Person
	for i := 0; i<1; i++{
		p := &Person{Name:string(i), Likes:[]string{"s1","s2","s3"}}
		people = append(people,p)

	}

	likes := make(map[string][]*Person)
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}
	for key, value := range likes {
		fmt.Println("Key:", key, "Value:", value)
	}
}
func test3() {
	type Key struct {
		Path, Country string
	}
	hits := make(map[Key]int)

	for i:=0; i<10 ;i++{
		k := Key{Path:string(i),Country:string(i)}
		hits[k] = 0
	}

	for i:=0; i<10 ;i++ {
		k := Key{Path:string(i),Country:string(i)}
		_, ok := hits[k]
		if ok {
			hits[k]++
		}
	}
	for i:=0; i<10 ;i++ {
		k := Key{Path:string(i),Country:string(i)}
		_, ok := hits[k]
		if ok {
			hits[k]++
		}
	}

	for key, value := range hits {
		fmt.Println("test Key:", key.Path, "Value:", value)
	}

	fmt.Println("test a value:", hits[Key{Path:string(1),Country:string(1)}])
	fmt.Println("test a value:", hits[Key{Path:string(11),Country:string(11)}])
}

func test4(){
	var counter = struct{
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n , "map_length:", len(counter.m))

	for key, value := range counter.m {
		fmt.Println("some Key:", key, "Value:", value)
	}


	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()

	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()

	for key, value := range counter.m {
		fmt.Println("some Key:", key, "Value:", value)
	}

}

func test5() {
	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	delete(countryCapitalMap,"France");
	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	//add value
	countryCapitalMap["United States"] = "United States"

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if(ok){
		fmt.Println("Capital of United States is", captial)
	}else {
		fmt.Println("Capital of United States is not present")
	}
}

func test6(){
	var m = map[string]int{
		"unix":         0,
		"python":       1,
		"go":           2,
		"javascript":   3,
		"testing":      4,
		"philosophy":   5,
		"startups":     6,
		"productivity": 7,
		"hn":           8,
		"reddit":       9,
		"C++":          10,
	}
	fmt.Println("len:", len(m))//, "map-cap:", cap(m))

	var array []string
	for i :=0 ; i<10; i++ {
		array = append(array,string(i))
	}
	fmt.Println("string-len:", len(array),"string-cap:", cap(array))

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}