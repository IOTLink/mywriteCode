package main

import (
	"fmt"
	"time"
)

type Event struct {
	Value int64
}

type DogWatch interface {
	OnNotify(Event)
}

type MangerWatcher interface {
	Register(DogWatch)
	DeRegister(DogWatch)
	Notify(DogWatch)
}

type eventDog struct {
	id int64
}

func (o *eventDog) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Value)
}


type eventWatch struct {
	servers map[DogWatch]struct{}
}



func (w *eventWatch) Register(d DogWatch) {
	w.servers[d] = struct{}{}
}

func (w *eventWatch) DeRegister(d DogWatch) {
	delete(w.servers, d)
}

func (w *eventWatch) Notify(e Event) {
	for k := range w.servers {
		k.OnNotify(e)
	}
}

func main() {
	//fmt.Println("hello world")
	manger := eventWatch{
		servers: map[DogWatch]struct{}{},
	}

	manger.Register(&eventDog{id: 1})
	manger.Register(&eventDog{id: 2})

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <- stop:
			return
		case t := <-tick:
			manger.Notify(Event{Value: t.UnixNano()})
		}
	}

}

/**
观察者模式的优点在于实现了表示层和数据层的分离
观察者模式的优点在于实现了表示层和数据层的分离，并定义了稳定的更新消息传递机制，类别清晰，抽象了更新接口，使得相同的数据层可以有各种不同的表示层

*/