package handler

import (
	"fmt"
	"net/http"

	"github.com/schlucht/gorecipe/fhx"
	"github.com/schlucht/gorecipe/pkg/web/config"
	"github.com/schlucht/gorecipe/pkg/web/model"
	"github.com/schlucht/gorecipe/pkg/web/render"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

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
	name := fhx.ReadRegex(reg, fhx2)
	fmt.Printf("%v\n", name["name"])
	// w.Write()
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "index.page.html", &model.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["hallo"] = "Welt"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.html", &model.TemplateData{
		StringMap: stringMap,
	})
}
