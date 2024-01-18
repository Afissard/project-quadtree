package ui

import (
	"image/color"

	"golang.org/x/image/font"
)

/*
Structure générale pour le paquet UI
*/
type UI struct {
	// paramètres
	textFont         font.Face
	textColor        color.RGBA
	textSpecialColor color.RGBA
	textDeath        color.RGBA

	// ui affiché en jeu
	MainMenu     MainMenuUI
	GameOver     GameOverUI
	DungeonLevel DungeonLevelUI
	EndScreen    EndScreenUI
}
