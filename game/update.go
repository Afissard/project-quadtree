package game

import (
	"github.com/Afissard/project-quadtree/configuration"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {
	// Main Menu
	if inpututil.IsKeyJustPressed(ebiten.KeyG){
		g.state = inGame
	}
	// In game
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.SaveGame("autoSave")
		// TODO: si debug : montré sauvegarde
	}

	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y)

	return nil
}
