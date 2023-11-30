package gui

import (
	imgui "github.com/gabstv/cimgui-go"
	ebimgui "github.com/gabstv/ebiten-imgui/v3"
)

func LoresIpsum() {
	// Init
	ebimgui.Update(1.0 / 60.0)
	ebimgui.BeginFrame()

	// Window
	imgui.Text("Hello World!")

	// End
	ebimgui.EndFrame()
}
