package handler

import (
	"net/http"

	"github.com/schlucht/gorecipe/pkg/web/model"
	"github.com/schlucht/gorecipe/pkg/web/render"
)

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "about"

	render.RenderTemplate(w, r, "about.page.html", &model.TemplateData{
		StringMap: stringMap,
	})
}
