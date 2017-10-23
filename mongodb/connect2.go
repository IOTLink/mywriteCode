package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name string
	Phone string
}


func main() {
	dialInfo := &mgo.DialInfo{  //需要认证权限登陆
		Addrs:     []string{"10.10.1.52"},
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  "ca",
		Source:    "",
		Username:  "admin",
		Password:  "adminw",
		PoolLimit: 27017, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if nil != err {
		//panic(err)
		fmt.Println(err.Error())
	} else {
		fmt.Println("connect success to remote mongod")
	}
	defer session.Close()



	//tests
	c := session.DB("ca").C("accout")
	err = c.Insert(&Person{"use1", "+55 53 8116 9639"},
		&Person{"use2", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "use2"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)

}