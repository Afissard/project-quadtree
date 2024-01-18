package character

import (
	"image"

	// "github.com/Afissard/project-quadtree/assets"
	// "github.com/Afissard/project-quadtree/configuration"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/assets"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw permet d'afficher le personnage dans une *ebiten.Image
// (en pratique, celle qui représente la fenêtre de jeu) en
// fonction des caractéristiques du personnage (position, orientation,
// étape d'animation, etc) et de la position de la caméra (le personnage
// est affiché relativement à la caméra).
func (c Character) Draw(screen *ebiten.Image, camX, camY int) {
	// xShift := 0
	// yShift := 0

	// xTileForDisplay := c.X - camX + configuration.Global.ScreenCenterTileX
	// yTileForDisplay := c.Y - camY + configuration.Global.ScreenCenterTileY
	// xPos := (xTileForDisplay)*configuration.Global.TileSize + xShift
	// yPos := (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 + yShift

	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(float64(xPos), float64(yPos))

	// shiftX := configuration.Global.TileSize
	// if c.Moving {
	// 	shiftX += c.AnimationStep * configuration.Global.TileSize
	// }
	// shiftY := c.Orientation * configuration.Global.TileSize

	// screen.DrawImage(assets.CharacterImage.SubImage(
	// 	image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
	// ).(*ebiten.Image), op)
	xShift := 0
	yShift := 0
	switch c.Orientation {
	case orientedDown:
		yShift = c.Shift
	case orientedUp:
		yShift = -c.Shift
	case orientedLeft:
		xShift = -c.Shift
	case orientedRight:
		xShift = c.Shift
	}

	xTileForDisplay := c.X - camX + configuration.Global.ScreenCenterTileX
	yTileForDisplay := c.Y - camY + configuration.Global.ScreenCenterTileY
	xPos := (xTileForDisplay)*configuration.Global.TileSize + xShift
	yPos := (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 + yShift

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xPos), float64(yPos))

	shiftX := configuration.Global.TileSize
	if c.Moving {
		shiftX += c.AnimationStep * configuration.Global.TileSize
	}
	shiftY := c.Orientation * configuration.Global.TileSize

	screen.DrawImage(assets.CharacterImage.SubImage(
		image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
	).(*ebiten.Image), op)
}
