package reach

import "github.com/ant0ine/go-json-rest/rest"

type Info struct {
	Name  string
	Email string
}

var p Provider

var All = &rest.Route{HttpMethod: "GET", PathExp: "/contact", Func: all}

func all(w rest.ResponseWriter, r *rest.Request) {
	results, err := p.findAll()
	if err != nil {
		w.WriteJson(err)
	}
	w.WriteJson(results)
}
