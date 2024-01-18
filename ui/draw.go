package ui

import "github.com/hajimehoshi/ebiten/v2"

/*
Fonction Draw (fichier draw.go)

Entrées: screen -> image représentant l'écran affiché
Sorties: screen mis à jour avec les ui dessiné dessus

Est appelé dans game.Draw
Affiche tout les interface activé, voir la structure UI.
*/
func (ui *UI) Draw(screen *ebiten.Image) {
	if ui.MainMenu.IsActive {
		ui.drawMainMenu(screen)
	} else if ui.EndScreen.IsActive {
		ui.drawEndScreen(screen)
	} else { // UI "overlay" en jeu
		if ui.GameOver.IsActive {
			ui.drawGameOver(screen)
		} else if ui.DungeonLevel.IsActive {
			ui.drawDungeonLevel(screen)
		}
	}
}
