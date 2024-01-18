package floor

import (
	// "github.com/Afissard/project-quadtree/configuration"

	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos, zoomVal int) {
	camXPos, camYPos = f.zoom(camXPos, camYPos, zoomVal)
	switch configuration.Global.FloorKind {
	case gridFloor:
		f.updateGridFloor(camXPos, camYPos)
	case fromFileFloor:
		f.updateFromFileFloor(camXPos, camYPos)
	case quadTreeFloor:
		f.updateQuadtreeFloor(camXPos, camYPos)
	case bspDungeonFloor:
		f.updateQuadtreeFloor(camXPos, camYPos)
	case infiniteWorld:
		f.QuadtreeContent.UpdateInfiniteWorld(camXPos, camYPos)
		f.updateQuadtreeFloor(camXPos, camYPos)
	}
}

func (f *Floor) zoom(camXPos, camYPos, zoomVal int) (newCamXPos, newCamYPos int) {
	configuration.Global.NumTileY = configuration.Global.OriginNumTileY * zoomVal
	configuration.Global.NumTileX = configuration.Global.OriginNumTileX * zoomVal
	configuration.Global.ScreenCenterTileY = configuration.Global.NumTileY / 2
	configuration.Global.ScreenCenterTileX = configuration.Global.NumTileX / 2
	if configuration.Global.NumTileY != len(f.content) {
		f.content = make([][]int, configuration.Global.NumTileY)
		for y := 0; y < len(f.content); y++ {
			f.content[y] = make([]int, configuration.Global.NumTileX)
		}
	}
	return camXPos, camYPos
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
		En entrée récupère une structure floor et les coordonnées
		de la caméra.
		En sortie, la région de la map afficher (Floor.content)
		par la caméra, soit une partie de Floor.fullContent, centrée
		sur les coordonnées de la caméra (carré de 9*9 tuile). (voir pdf)
	*/
	for y := 0; y < configuration.Global.NumTileY; y++ {
		for x := 0; x < configuration.Global.NumTileX; x++ {
			// Calcule les coordonnées des tuiles à afficher par rapport aux coordonnée de la caméra
			tileY := camYPos - configuration.Global.ScreenCenterTileY + y
			tileX := camXPos - configuration.Global.ScreenCenterTileX + x

			// Test tuile vide ou non
			if (tileY < 0) || (tileY >= len(f.fullContent)) || // si tileY hors porté
				(tileX < 0) || (tileX >= len(f.fullContent[tileY])) { // si tileX hors porté
				f.content[y][x] = -1 // case vide
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
	f.QuadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
