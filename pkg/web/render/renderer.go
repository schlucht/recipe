package render

import (
	"bytes"
	"html/template"
	"log"

	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/schlucht/gorecipe/pkg/web/config"
	"github.com/schlucht/gorecipe/pkg/web/model"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *model.TemplateData, r *http.Request) *model.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *model.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCatch {
		tc = app.TemplateCatch
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r )
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		log.Printf("Fehler beim Einlesen der Page Seiten %v\n", err.Error())
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("Fehler im Template")
			return myCache, err
		}
		layouts, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			log.Printf("Fehler beim Einlesen der Layout Seiten")
			return myCache, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				log.Printf("Fehler beim Parsen der Layout Seiten")
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil

}
