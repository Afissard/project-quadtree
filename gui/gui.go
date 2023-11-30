package gui

// TODO remplacé par un dictionnaire pour stocker toute les fenêtre
type WinGui struct {
	ActiveWin bool
	Main      bool
	Poem      bool
}

func (gui *WinGui) Init() {
	gui.ActiveWin = false
	gui.Main = false
}
