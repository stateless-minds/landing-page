package main

import (
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type home struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "pubsub World!" is displayed as a heading.
func (h *home) Render() app.UI {
	return app.Div().Class("stack").Body(
		app.Div().Class("content card").Body(
			app.H1().Body(
				app.B().Body(
					app.Span().ID("desktop-prompt").Body(
						app.Text("Stateless"),
					),
					app.Span().ID("mobile-prompt").Body(
						app.Text("Stateless"),
					),
					app.Text(" Minds"),
				),
			),
			app.P().Body(
				app.Text("We question reality."),
			),
		),
		app.Div().Class("padding card"),
		app.Div().Class("border card"),
		app.Div().Class("background card"),
		app.Div().Class("box-shadow card"),
	)
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &home{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".

	withGz := gziphandler.GzipHandler(&app.Handler{
		Name:        "statelessminds",
		Description: "Indie startup studio making digital products that question reality.",
		Styles: []string{
			"/web/app.css",
		},
		Scripts: []string{},
	})
	http.Handle("/", withGz)

	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}
