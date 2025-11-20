package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"
	"time"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type StatusState string

const (
	StateFocus StatusState = "focus"
	StateBreak StatusState = "break"
)

type Status struct {
	State    StatusState `json:"state"`
	IsNotify bool        `json:"is_notify"`

	// Duration left where we hit pause.
	PausePoint *string   `json:"pause_point"`
	EndTime    time.Time `json:"end_time"`
}

func main() {
	statusPath := "/tmp/simo.json"

	if len(os.Args) > 2 {
		fmt.Fprintf(
			os.Stdout,
			"%s\n%s\n",
			"Too much arguments. Only need json file path.",
			"For example: ./simo-gui /tmp/simo.json",
		)

		os.Exit(0)
	} else if len(os.Args) == 2 {
		statusPath = os.Args[1]
	}

	go func() {
		window := new(app.Window)

		window.Option(
			app.Size(
				unit.Dp(400),
				unit.Dp(70),
			),
		)

		err := run(window, statusPath)

		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}

func run(window *app.Window, statusPath string) error {
	var (
		ops   op.Ops
		theme = material.NewTheme()
	)

	var (
		initState = "not started"
		state     = initState
		status    = new(Status)
	)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// For managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define label.
			title := material.H3(
				theme,
				strings.ToUpper(state),
			)

			titleColor := color.NRGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 255,
			}

			title.Color = titleColor
			title.Alignment = text.Middle

			title.Layout(gtx)

			data, err := os.ReadFile(statusPath)

			if !errors.Is(err, os.ErrNotExist) {
				err = json.Unmarshal(data, status)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}

				state = string(status.State)
			} else {
				state = initState
			}

			// Reference:
			// https://jonegil.github.io/gui-with-gio/egg_timer/11_improved_animation.html
			gtx.Execute(op.InvalidateCmd{})

			e.Frame(gtx.Ops)
		}
	}
}
