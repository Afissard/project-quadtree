package camera

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
)

// Affiche l'image du floor et du joueur par rapport aux paramètre de la camera, y applique notamment le zoom 
func (c *Camera) Draw(screen, floorImage *ebiten.Image) {
	camView := ebiten.NewImage(configuration.Global.NumTileX*configuration.Global.TileSize, configuration.Global.NumTileY*configuration.Global.TileSize)

	// if configuration.Global.DebugMode {
	// 	camView.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 50}) // debug
	// }

	halfX := camView.Bounds().Dx() / 2
	halfY := camView.Bounds().Dy() / 2
	halfSizeFloorX := floorImage.Bounds().Dx() / 2
	halfSizeFloorY := floorImage.Bounds().Dy() / 2
	r := image.Rect(halfSizeFloorX-halfX, halfSizeFloorY-halfY, halfSizeFloorX+halfX, halfSizeFloorY+halfY)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1/float64(configuration.Global.NumTileX/configuration.Global.OriginNumTileX), 1/float64(configuration.Global.NumTileY/configuration.Global.OriginNumTileY))
	camView.DrawImage(floorImage.SubImage(r).(*ebiten.Image), op)

	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(camView, op)
}
