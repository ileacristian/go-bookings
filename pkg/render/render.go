package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/ileacristian/go-bookings/pkg/config"
	"github.com/ileacristian/go-bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {

	return templateData
}

func RenderTemplate(w http.ResponseWriter, tempateName string, tempateData *models.TemplateData) {
	var cache map[string]*template.Template
	if app.UseCache {
		cache = app.TemplateCache
	} else {
		cache, _ = CreateTemplateCache()
	}

	template, ok := cache[tempateName]
	if !ok {
		log.Fatal("could not find template in cache")
	}

	tempateData = AddDefaultData(tempateData)

	tempBuffer := new(bytes.Buffer)
	_ = template.Execute(tempBuffer, tempateData)

	w.Write(tempBuffer.Bytes())
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		templateName := filepath.Base(page)
		template, err := template.New(templateName).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			template, err = template.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[templateName] = template
	}

	log.Println("successfully created templates cache")

	return cache, nil
}
