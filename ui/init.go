package ui

import (
	"image/color"

	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/assets"
)

/*
Fonction Init (fichier init.go)

Initialise la structure UI et les variables associés.
*/
func (ui *UI) Init() {
	// paramètres du paquet UI
	ui.textColor = color.RGBA{0xc0, 0xc0, 0xc0, 0xff}
	ui.textDeath = color.RGBA{0xff, 0xa5, 0x00, 0xff}
	ui.textSpecialColor = color.RGBA{0xee, 0x20, 0x20, 0xff}
	ui.textFont = assets.FontBigBlueTerm

	// Main Menu
	ui.MainMenu.IsActive = true
	ui.MainMenu.MenuList = append(ui.MainMenu.MenuList, initMenuList...)

	// Dungeon Level
	ui.DungeonLevel.DungeonLevel = 0
}
