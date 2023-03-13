package models

//TemplateData holds data sent from handlers to templates
type TemplateData struct{
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]string
	Data      map[string]interface{}
	CSFRToken string
	Flash     string
	Warning   string
	Error     string
}