package ui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

/*
DungeonLevelUI est l'interface affiché si le joueur change d'étage.
*/
type DungeonLevelUI struct {
	// Commun aux structure d'UI
	IsActive   bool
	TimeShowed int

	DungeonLevel int
}

/*
Fonction updateDungeonLevel (fichier dungeonLevel.go)

Est appelé en jeu si le joueur change d'étage.
*/
func (ui *UI) updateDungeonLevel() {
	if ui.DungeonLevel.TimeShowed > 0 {
		ui.DungeonLevel.TimeShowed--
	} else {
		ui.DungeonLevel.IsActive = false
	}
}

/*
Fonction drawDungeonLevel (fichier dungeonLevel.go)

Entrées: screen -> image représentant l'écran affiché
Sorties: screen mis à jour avec l'écran de changement d'étage dessus
*/
func (ui *UI) drawDungeonLevel(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("Level : %d", ui.DungeonLevel.DungeonLevel), ui.textFont, 16, 16, ui.textColor)
}
