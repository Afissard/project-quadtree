package floor

import (
	"image"

	// "github.com/Afissard/project-quadtree/assets"
	// "github.com/Afissard/project-quadtree/configuration"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/assets"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).

// func (f Floor) Draw(screen *ebiten.Image) {
// 	// for y ù1; y < len(f.content)-1; y++ {
// 	// 	for x := 1; x < len(f.content[y])-1; x++ {
// 	// 		if f.content[y][x] != -1 {
// 	// 			op := &ebiten.DrawImageOptions{}
// 	// 			op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize))

// 	// 			shiftX := f.content[y][x] * configuration.Global.TileSize

// 	// 			currentFloorImage.DrawImage(
// 	// 				assets.FloorImage.SubImage(
// 	// 					image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
// 	// 				).(*ebiten.Image), op)
// 	// 		}
// 	// 	}
// 	// }
// 	op := &ebiten.DrawImageOptions{}
// 	currentFloorImage := f.CreateFloorImage()
// 	// if ebiten.IsKeyPressed(ebiten.KeyRight) {
// 	// 	xPos += f.animationStep
// 	// } else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
// 	// 	xPos -= f.animationStep
// 	// } else if ebiten.IsKeyPressed(ebiten.KeyUp) {
// 	// 	yPos += f.animationStep
// 	// } else if ebiten.IsKeyPressed(ebiten.KeyDown) {
// 	// 	yPos -= f.animationStep
// 	// }
// 	// op.GeoM.Translate(float64(xPos), float64(yPos))

// 	screen.DrawImage(currentFloorImage, op)
// }

//CreateFloorImage (currentFloorImage *ebiten.Image)

func (f Floor) Draw(screen *ebiten.Image) {
	for y := range f.content {
		for x := range f.content[y] {
			if f.content[y][x] != -1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize))

				shiftX := f.content[y][x] * configuration.Global.TileSize

				screen.DrawImage(assets.FloorImage.SubImage(
					image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
				).(*ebiten.Image), op)
			}
		}
	}
}

func (f *Floor) GetFloorImage() (currentFloorImage *ebiten.Image) {
	currentFloorImage = ebiten.NewImage(configuration.Global.NumTileX*configuration.Global.TileSize, configuration.Global.NumTileX*configuration.Global.TileSize)
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			if f.content[y][x] != -1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize))

				shiftX := f.content[y][x] * configuration.Global.TileSize

				currentFloorImage.DrawImage(assets.FloorImage.SubImage(
					image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
				).(*ebiten.Image), op)
			}
		}
	}
	return currentFloorImage
}
