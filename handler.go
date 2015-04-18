package reach

import "github.com/ant0ine/go-json-rest/rest"

var handler Handler

func MakeHandler() {
	handler = Handler{provider: &Mongo{}}
}

type Handler struct {
	provider Provider
}

func (h *Handler) all(w rest.ResponseWriter, r *rest.Request) {
	results, err := h.provider.findAll()
	if err != nil {
		w.WriteJson(err)
		return
	}
	w.WriteJson(results)
}
