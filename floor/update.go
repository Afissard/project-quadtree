package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	switch configuration.Global.FloorKind {
	case gridFloor:
		f.updateGridFloor(camXPos, camYPos)
	case fromFileFloor:
		f.updateFromFileFloor(camXPos, camYPos)
	case quadTreeFloor:
		f.updateQuadtreeFloor(camXPos, camYPos)
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(camXPos, camYPos int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absCamX := camXPos
			if absCamX < 0 {
				absCamX = -absCamX
			}
			absCamY := camYPos
			if absCamY < 0 {
				absCamY = -absCamY
			}
			f.content[y][x] = ((x + absCamX%2) + (y + absCamY%2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
func (f *Floor) updateFromFileFloor(camXPos, camYPos int) {
	/*
		En entré récupère une structure floor et les coordonnée
		de la caméra.
		En sortie, la région de la map afficher (Floor.content)
		par la caméra, soit une partie de Floor.fullContent, centré
		sur les coordonné de la camera (carré de 9*9 tuile). (voir pdf)
		#TODO :
		-> retourné une région de la map (Floor.content)
		-> fonction test ?
	*/
	for y := 0; y < len(f.fullContent) || y < configuration.Global.NumTileY; y++ {
		for x := 0; x < len(f.fullContent[y]) || x < configuration.Global.NumTileX; x++ {
			// les coordonnées relatives doivent être centrées sur la camera
			tileY := 0
			tileX := 0
			//tileY = len(f.fullContent) - configuration.Global.NumTileY - camYPos + y
			//tileX = len(f.fullContent[y]) - configuration.Global.NumTileX - camXPos + x

			if (tileY < 0 || tileY > len(f.fullContent)) || (tileX < 0 || tileX > len(f.fullContent[y])) { // case vide
				f.content[y][x] = -1
			} else {
				f.content[y][x] = f.fullContent[tileY][tileX]
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
