package main

import (
	"time"
	"fmt"
	"sync"
	"math"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1)
//	time.Sleep(time.Second)

	t2 := time.Now()
	flag := t2.After(t1)
	fmt.Println(flag)

	t3 := time.Now()
	fmt.Println(t3)
	t4 := time.Now()
	fmt.Println(t4)
	flag = t3.Equal(t4)
	fmt.Println(flag)

	fmt.Println(time.Now().Format(time.ANSIC))
	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println(time.Now().Format(time.RFC3339Nano))
	fmt.Println(time.Now().Format(time.Kitchen))
	layout := "01-02-2006 3.04.05 PM"
	fmt.Println(time.Now().Format(layout))

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println(time.Now().UTC().String())
	fmt.Println(time.Now().Local().String())

	timeDays := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(timeDays)

	start := time.Now()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(start," ",end," ",elapsed)

	/*
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v %s\n", now, "h")
	}
	*/

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	t = time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}

	t, _ = time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")
	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
	}

	for _, d := range trunc {
		fmt.Printf("t.Truncate(%5s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
	}

	t = time.Unix(1362984425, 0)
	nt := t.Format("2006-01-02 15:04:05")
	fmt.Println(nt)

	times := time.Now().Unix()
	fmt.Println(times)
	t = time.Unix(times, 0)
	nt = t.Format("2006-01-02 15:04:05")
	fmt.Println(nt," local: ",t.Local()," utc:",t.UTC().String())

	fmt.Printf("%d\n",uint64(math.MaxUint64))
	fmt.Printf("%d\n",int64(math.MaxInt64))
	t = time.Unix(math.MaxInt64, 0)
	nt = t.Format("2006-01-02 15:04:05")
	fmt.Println(nt,"-----------> local: ",t.Local()," utc:",t.UTC().String())

	ticker := time.NewTicker(time.Millisecond * 500)
	c := make(chan int,3) //num为指定的执行次数
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for t := range ticker.C {
			if len(c) == 3 {
				break
			}
			c<-1
			fmt.Println("Tick at", t)

		}
	}()
	wg.Wait()

	/*
	for {
		timer := time.NewTimer(time.Millisecond*500)
		select {
		case t :=<- timer.C:
			fmt.Println("test Tick at", t)
		}
		timer.Stop()

		//select {
		//case <-time.After(2 * time.Second):
		//	fmt.Println("test time out")
		//}
	}

	timer := time.NewTimer(time.Millisecond*500)
	for {
		select {
		case tt :=<- timer.C:
			fmt.Println("test Tick at", tt)
		case <-time.After(2 * time.Second):
			fmt.Println("test time out")
		}
	}
	*/

	ticker = time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case tt :=<-ticker.C:
			fmt.Println("test 2 Tick at", tt)
		case <-time.After(1 * time.Second):
			fmt.Println("test 2 time out")
		}
		/*
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("test 2 time out")
		}*/
	}
	ticker.Stop()
}
