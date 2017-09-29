package main

import (
	"context"
	"fmt"
	"sync"
)

func test1() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func test2()  {
	var wg sync.WaitGroup

	wg.Add(2)

	dst := make(chan int)
	go func() {
		defer wg.Done()
		fmt.Println("hello world")
		n := 1

		for {
			select {
			case dst <- n:
				n++
			}
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("hello world")

		for {
			select {
				case n :=<-dst:
					fmt.Println(n)
			}
		}
	}()
	wg.Wait()

}

func main() {
	//test1()
	test2()
}

