package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rgasc/booking-app/pkg/config"
	"github.com/rgasc/booking-app/pkg/models"
)

var functions = template.FuncMap{}
var App *config.AppConfig

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if App.UseCache {
		tc = App.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	if t, ok := tc[tmpl]; ok {
		buf := new(bytes.Buffer)
		_ = t.Execute(buf, AddDefaultData(td))
		_, err := buf.WriteTo(w)
		if err != nil {
			fmt.Println("error writing template to browser", err)
			return
		}
	} else {
		log.Fatal("could not get template from cache")
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = ts
	}

	return templateCache, nil
}
