package reach

import (
	"log"
	"fmt"
	"gopkg.in/mgo.v2"
)

type Mongo struct{}

func (m *Mongo) findAll() (results []Info, err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	defer func() {
	        if r := recover(); r != nil {
	            fmt.Println("Recovered in f", r)
                        var ok bool
                        err, ok = r.(error)
                        if !ok {
                                err = fmt.Errorf("Dial: %v", r)
                        }
	        }
    	}()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	err = c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
