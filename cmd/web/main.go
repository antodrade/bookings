package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"bookings/internal/config"
	"bookings/internal/handlers"
	"bookings/internal/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

    app.InProduction = false
	session = scs.New()
	session.Lifetime = 24*time.Hour
	session.Cookie.Persist=true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	

	render.NewTemplate(&app)
    


	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting aplication on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
