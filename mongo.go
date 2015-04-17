package reach

import (
	"log"

	"gopkg.in/mgo.v2"
)

type Mongo struct{}

func (m *Mongo) findAll() ([]Info, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	results := []Info{}
	err = c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results, nil
}
