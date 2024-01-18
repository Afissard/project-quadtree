package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
)

/*
GameOverUI est l'interface affiché si le joueur perd.
*/
type GameOverUI struct {
	// Commun aux structure d'UI
	IsActive   bool
	TimeShowed int
	SourceMort int
	message    string
}

/*
Fonction updatePoem (fichier gameOver.go)

Est appelé en jeu si le joueur perd.
*/
func (ui *UI) updateGameOver() {
	if ui.GameOver.TimeShowed > 0 {
		ui.GameOver.TimeShowed--
	}
	if ui.GameOver.SourceMort == -1 {
		ui.GameOver.message = "Vous êtes désormais dans le néant."
	} else if ui.GameOver.SourceMort == 2 {
		ui.GameOver.message = "Vous êtes mort d'un piège."
	} else if ui.GameOver.SourceMort == 9 {
		ui.GameOver.message = "Vous êtes tombé dans un trou."
	}
}

/*
Fonction drawMainMenu (fichier gameOver.go)

Entrées: screen -> image représentant l'écran affiché
Sorties: screen mis à jour avec l'écran de Game Over dessus
*/
func (ui *UI) drawGameOver(screen *ebiten.Image) {
	text.Draw(screen, "Game Over", ui.textFont, configuration.Global.ScreenWidth/2-32, configuration.Global.ScreenHeight/2-8, ui.textSpecialColor)
	text.Draw(screen, ui.GameOver.message, ui.textFont, configuration.Global.ScreenWidth/2-110, configuration.Global.ScreenHeight/2+8, ui.textDeath)
}
