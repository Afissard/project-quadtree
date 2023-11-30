package gui

import (
	imgui "github.com/gabstv/cimgui-go"
	ebimgui "github.com/gabstv/ebiten-imgui/v3"
)

func LoresIpsum() {
	ebimgui.Update(1.0 / 60.0)
	ebimgui.BeginFrame()
	imgui.Text("Hello World!")
	ebimgui.EndFrame()
}
