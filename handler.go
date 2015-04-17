package reach

import "github.com/ant0ine/go-json-rest/rest"

var handler Handler

type Handler struct {
	provider Provider
}

func (h *Handler) all(w rest.ResponseWriter, r *rest.Request) {
	results, err := h.provider.findAll()
	if err != nil {
		w.WriteJson(err)
	}
	w.WriteJson(results)
}
