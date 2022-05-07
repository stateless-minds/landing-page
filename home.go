package main

import (
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
