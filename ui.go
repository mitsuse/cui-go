/*
Package "ui" provide A library for building CUI.
*/
package ui

import (
	"github.com/nsf/termbox-go"
)

/*
Run the main loop of application.
*/
func Run() error {
	if err := termbox.Init(); err != nil {
		return err
	}
	defer termbox.Close()

	for {
		event := termbox.PollEvent()

		// TODO: Send event to the focused view.
		_ = event
	}

	return nil
}
