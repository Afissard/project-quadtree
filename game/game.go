package game

import (
	"github.com/Afissard/project-quadtree/camera"
	"github.com/Afissard/project-quadtree/character"
	"github.com/Afissard/project-quadtree/floor"
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

	state int
}
