package controller

import (
	"api/handler"
	"api/model"
	"net/http"
)

// Default is a controller by default.
type Default struct {
	*Base
}

// NewDefault returns an instance of Default.
func NewDefault() *Default {
	d := &Default{
		Base: &Base{
			Routes: model.Routes{},
			Prefix: "/",
		},
	}

	d.Routes = append(d.Routes, model.Route{
		Name:    "Index",
		Method:  "GET",
		Pattern: d.Prefix,
		Func:    d.Index,
	})

	return d
}

// Index HomePage
func (d *Default) Index(w http.ResponseWriter, r *http.Request) {
	handler.SendJSONOk(w, "hello")
}
