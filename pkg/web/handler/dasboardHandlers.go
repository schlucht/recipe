package handler

import (
	"net/http"

	"github.com/schlucht/gorecipe/pkg/web/model"
	"github.com/schlucht/gorecipe/pkg/web/render"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "dashboard"
	render.RenderTemplate(w, r, "index.page.html", &model.TemplateData{
		StringMap: stringMap,
	})
}
