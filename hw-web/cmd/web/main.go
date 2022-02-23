package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davolu/go-lang-snippets/pkg/config"
	"github.com/davolu/go-lang-snippets/pkg/handlers"
	"github.com/davolu/go-lang-snippets/pkg/render"
)

const PORT = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreatTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println((fmt.Sprintf("Starting application on port %s", PORT)))
	_ = http.ListenAndServe(PORT, nil)
}
