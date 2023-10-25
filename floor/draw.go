package floor

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
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
