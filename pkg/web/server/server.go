package server

import (
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/schlucht/gorecipe/pkg/web/config"
	"github.com/schlucht/gorecipe/pkg/web/handler"
	"github.com/schlucht/gorecipe/pkg/web/render"
)

type Server struct {
	Port string
}

var app config.AppConfig
var session *scs.SessionManager

func (m *Server) LoadServer() {
	m.Port = ":8080"

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCatch = tc
	app.UseCatch = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplate(&app)
	fmt.Println("Der Server läuft auf PORT:", m.Port)
	srv := &http.Server{
		Addr:    m.Port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
