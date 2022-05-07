//go:build !wasm

// file: http_server.go

package main

import (
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func initServer() {
	withGz := gziphandler.GzipHandler(&app.Handler{
		Name:        "statelessminds",
		Description: "Indie startup studio making digital products that question reality.",
		Styles: []string{
			"/web/app.css",
		},
		Scripts: []string{},
	})
	http.Handle("/", withGz)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
