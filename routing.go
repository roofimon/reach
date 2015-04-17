package reach

import "github.com/ant0ine/go-json-rest/rest"

var All = &rest.Route{HttpMethod: "GET", PathExp: "/contact", Func: handler.all}
