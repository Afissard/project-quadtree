package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

/*
Main Menu est le menu principal du jeu, plusieurs
option y sont proposé :
  - Jouer
  - Modifier les paramètres
  - Sauvegarder la partie en cours
  - Quitter
*/
type MainMenuUI struct {
	// Commun aux structure d'UI
	IsActive bool

	// paramètre de l'ui
	OptionSelected int
	OptionChoose   bool
	MenuList       []string
}

var (
	initMenuList = []string{
		"Jouer",
		"Sauvegarder",
		"Charger la sauvegarde",
		"Quitter",
	}
)

/*
Fonction updateMainMenu (fichier mainMenu.go)

Est appelé dans UI.Update si le menu principal est ouvert.
Récupère les inputs donnés et met à jour le menu en conséquence.
*/
func (ui *UI) updateMainMenu() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		if ui.MainMenu.OptionSelected > 0 {
			ui.MainMenu.OptionSelected -= 1
		} else {
			ui.MainMenu.OptionSelected = len(ui.MainMenu.MenuList) - 1
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		if ui.MainMenu.OptionSelected < len(ui.MainMenu.MenuList)-1 {
			ui.MainMenu.OptionSelected += 1
		} else {
			ui.MainMenu.OptionSelected = 0
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		ui.MainMenu.OptionChoose = true
	} else {
		// reset to default value
		ui.MainMenu.OptionChoose = false
	}
}

/*
Fonction drawMainMenu (fichier mainMenu.go)

Entrées: screen -> image représentant l'écran affiché
Sorties: screen mis à jour avec le menu principal dessiné dessus

Affiche le menu principal sur l'écran, donnée en entré.
*/
func (ui *UI) drawMainMenu(screen *ebiten.Image) {
	text.Draw(screen, "---<CAELAH>---\n\nPress Up & Down to choose an option", ui.textFont, 16, 16, ui.textColor)
	for i := 0; i < len(ui.MainMenu.MenuList); i++ {
		text.Draw(screen, ui.MainMenu.MenuList[i], ui.textFont, 32, 64+(i*16), ui.textColor)
	}
	text.Draw(screen, "*", ui.textFont, 16, 64+(ui.MainMenu.OptionSelected*16), ui.textColor)
}
