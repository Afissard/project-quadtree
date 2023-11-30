package game

import (
	"github.com/Afissard/project-quadtree/configuration"
	//"github.com/Afissard/project-quadtree/gui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		configuration.Global.Gui = !configuration.Global.Gui
	}

	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y)

	// if configuration.Global.Gui {
	// 	gui.LoresIpsum()
	// }

	return nil
}
