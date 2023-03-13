package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	//"log"
)

//Appconfig holds the application config
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfogLog *log.Logger
	InProduction bool
	Session *scs.SessionManager
	

}
