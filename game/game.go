package game

import (
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/camera"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/character"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/floor"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/ui"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/dungeon"
)

// Game est le type permettant de représenter les données du jeu.
// Aucun champs n'est exporté pour le moment.
//
// Les champs non exportés sont :
//   - camera : la représentation de la caméra
//   - floor : la représentation du terrain
//   - character : la représentation du personnage
type Game struct {
	camera    camera.Camera
	floor     floor.Floor
	character character.Character
	ui        ui.UI

	state        int
	zoomLevel    int
	needNewLevel bool
	listPortal []dungeon.Portal
}

/*
Liste des state du jeu
*/
const (
	mainMenu int = iota // commence à 0, les constantes suivantes auront la valeur d'un int + "n° déclaration"
	inGame
	gameOver
	endScreen
)

var sauvegardeCam string
var sauvegarde string
