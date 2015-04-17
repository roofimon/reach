package reach

import (
	"log"

	"gopkg.in/mgo.v2"
)

type Provider interface {
	findAll() ([]Info, error)
}

type Memory struct{}

func (m *Memory) findAll() ([]Info, error) {
	return []Info{
		Info{Name: "roof", Email: "roofimon@gmail.com"},
		Info{Name: "foor", Email: "foorimon@gmail.com"},
	}, nil
}

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
