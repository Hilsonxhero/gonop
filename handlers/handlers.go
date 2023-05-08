package handlers

import (
	"myapp/data"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/hilsonxhero/napoleon"
)

type Handlers struct {
	App    *napoleon.Napoleon
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	myData := "bar"
	h.App.Session.Put(r.Context(), "foo", myData)
	myValue := h.App.Session.GetString(r.Context(), "foo")
	vars := make(jet.VarMap)
	vars.Set("foo", myValue)
	err := h.App.Render.Page(w, r, "home", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
