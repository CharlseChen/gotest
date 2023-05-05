package main

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	p := widgets.NewParagraph()
	p.Title = "Hello World!"
	p.Text = "Press Q to quit."
	p.SetRect(0, 0, 25, 5)
	p.Border = true

	termui.Render(p)

	// Listen for keyboard events
	termuiEvents := termui.PollEvents()
	for {
		select {
		case e := <-termuiEvents:
			if e.Type == termui.KeyboardEvent && e.ID == "q" {
				return
			}
		}
	}
}
