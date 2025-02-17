package utils

import (
	"html/template"
	"net/http"
)

var templatesGerais *template.Template

func CarregarTemplates(){
	templatesGerais = template.Must(template.ParseGlob("views/*.html"))
}

func Exec(w http.ResponseWriter, template string, dados interface{}){
	templatesGerais.ExecuteTemplate(w, template, dados)
}
