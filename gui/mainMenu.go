package gui

import (
	imgui "github.com/gabstv/cimgui-go"
	ebimgui "github.com/gabstv/ebiten-imgui/v3"
)

func (gui *WinGui) MainMenu() {
	// Init
	ebimgui.Update(1.0 / 60.0)
	ebimgui.BeginFrame()

	// Window
	if imgui.Button("Poem") { // Buttons return true when clicked (most widgets return true when edited/activated)
		gui.Poem = !gui.Poem
	}

	//End
	ebimgui.EndFrame()
}
