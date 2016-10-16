/*
Package "ui" provide A library for building CUI.
*/
package ui

import (
	"github.com/nsf/termbox-go"
)

/*
Run the main loop of application.

The main loop repeats rendering and event handling.
*/
func Run(renderer Renderer) error {
	if err := termbox.Init(); err != nil {
		return err
	}
	defer termbox.Close()

	for {
		render(renderer)

		event := termbox.PollEvent()

		// TODO: Send event to the focused view.
		_ = event
	}

	return nil
}

func render(renderer Renderer) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	renderer.Render()
	termbox.Flush()
}

/*
Renderer is abstraction of view.
*/
type Renderer interface {
	Render()
}
