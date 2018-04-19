package main

import (
	"fmt"
	"time"
)

type ObServer struct { //观察者
	id int64
}

type MyEvent struct {//观察的对象　关注的消息
	Value int64
}

type TopicSubject struct {//观察者管理
	 obServer map[*ObServer]interface{}
}
/*
——目标知道它的观察者。可以有任意多个观察者观察同一个目标；
——提供注册和删除观察者对象的接口。
*/

func (t *TopicSubject)Register(obj *ObServer) {
	t.obServer[obj] = "test1"
}

func (t *TopicSubject) DeRegister(obj *ObServer) {
	delete(t.obServer, obj)
}

func (t *TopicSubject) Notify(e *MyEvent) {
	for k := range t.obServer {
		fmt.Printf("hi %d %d\n", k.id, e.Value)
	}
}

func main() {
	mangerObServer := TopicSubject {
		obServer: map[*ObServer]interface{}{},
	}

	mangerObServer.Register(&ObServer{id:1})
	mangerObServer.Register(&ObServer{id:2})
	mangerObServer.Register(&ObServer{id:3})

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C

	for {
		select {
		case <- stop:
			return
		case t := <-tick:
			mangerObServer.Notify(&MyEvent{Value: t.UnixNano()})
		}
	}
}

/*
Subject（目标）
——目标知道它的观察者。可以有任意多个观察者观察同一个目标；
——提供注册和删除观察者对象的接口。

Observer（观察者）
——为那些在目标发生改变时需获得通知的对象定义一个更新接口。

*/