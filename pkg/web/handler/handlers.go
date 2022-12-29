package handler

import (
	"fmt"
	"net/http"

	"github.com/schlucht/gorecipe/pkg/fhx"
	"github.com/schlucht/gorecipe/pkg/web/config"
	"github.com/schlucht/gorecipe/pkg/web/model"
	"github.com/schlucht/gorecipe/pkg/web/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) ReadFhx(w http.ResponseWriter, r *http.Request) {
	fhx2, _ := fhx.ReadUTF16("fhx/deltaV.fhx")
	reg := make(map[string]string)
	reg["name"] = `BATCH_RECIPE NAME="(?P<s>.*)" T`
	stringMap := make(map[string]string)
	title := fhx.ReadRegex(reg, fhx2)
	stringMap["name"] = fmt.Sprint(title["name"])
	render.RenderTemplate(w, r, "fhx.page.html", &model.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) NewUnit(w http.ResponseWriter, r *http.Request) {
	unit := model.Unit{}
	units := unit.Load()
	unitData := make(map[string]interface{})
	unitData["units"] = units
	render.RenderTemplate(w, r, "form.page.html", &model.TemplateData{
		Data: unitData,
	})
}

func (m *Repository) PostNewUnit(w http.ResponseWriter, r *http.Request) {
	title := r.Form.Get("uptitle")
	w.Write([]byte("Posted the new UP Tilte " + title))
}


