package main
import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	/* err
	session, err := mgo.Dial("10.0.0.52:27017")
	if err != nil {
		//panic(err)
		fmt.Println(err.Error())
		return
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	*/
	maxWait := time.Duration(5 * time.Second)
	session, sessionErr := mgo.DialWithTimeout("10.10.1.52:27017", maxWait)
	if sessionErr == nil {
		fmt.Println(" connect to remote mongo instance!")
	} else { // never gets here
		fmt.Println("Unable to connect to remote mongo instance!")
		return
	}



	c := session.DB("ca").C("accout")
	err := c.Insert(&Person{"use", "+55 53 8116 9639"},
		&Person{"user", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}

/*
https://studygolang.com/articles/3485


*/