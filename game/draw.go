package game

import (
	"fmt"
	"image/color"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Draw permet d'afficher à l'écran tous les éléments du jeu
// (le sol, le personnage, les éventuelles informations de debug).
// Il faut faire attention à l'ordre d'affichage pour éviter d'avoir
// des éléments qui en cachent d'autres.
func (g *Game) Draw(screen *ebiten.Image) {
	g.floor.Draw(screen)
	g.character.Draw(screen, g.camera.X, g.camera.Y)

	if configuration.Global.DebugMode {
		g.drawDebug(screen)
	}
}

// drawDebug se charge d'afficher les informations de debug si
// l'utilisateur le demande (positions absolues du personnage
// et de la caméra, grille avec les coordonnées, etc).
func (g Game) drawDebug(screen *ebiten.Image) {

	gridColor := color.NRGBA{R: 255, G: 255, B: 255, A: 63}
	gridHoverColor := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	gridLineSize := 2
	cameraColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	cameraLineSize := 1

	mouseX, mouseY := ebiten.CursorPosition()

	xMaxPos := configuration.Global.ScreenWidth
	yMaxPos := configuration.Global.ScreenHeight

	for x := 0; x < configuration.Global.NumTileX; x++ {
		xGeneralPos := x*configuration.Global.TileSize + configuration.Global.TileSize/2
		xPos := float32(xGeneralPos)

		lineColor := gridColor
		if mouseX+1 >= xGeneralPos && mouseX+1 <= xGeneralPos+gridLineSize {
			lineColor = gridHoverColor
		}

		vector.StrokeLine(screen, xPos, 0, xPos, float32(yMaxPos), float32(gridLineSize), lineColor, false)

		xPrintValue := g.camera.X + x - configuration.Global.ScreenCenterTileX
		xPrint := fmt.Sprint(xPrintValue)
		if len(xPrint) <= (2*configuration.Global.TileSize)/16 || (xPrintValue > 0 && xPrintValue%2 == 0) || (xPrintValue < 0 && (-xPrintValue)%2 == 0) {
			xTextPos := xGeneralPos - 3*len(xPrint) - 1
			ebitenutil.DebugPrintAt(screen, xPrint, xTextPos, yMaxPos)
		}
	}

	for y := 0; y < configuration.Global.NumTileY; y++ {
		yGeneralPos := y*configuration.Global.TileSize + configuration.Global.TileSize/2
		yPos := float32(yGeneralPos)

		lineColor := gridColor
		if mouseY+1 >= yGeneralPos && mouseY+1 <= yGeneralPos+gridLineSize {
			lineColor = gridHoverColor
		}

		vector.StrokeLine(screen, 0, yPos, float32(xMaxPos), yPos, float32(gridLineSize), lineColor, false)

		yPrint := fmt.Sprint(g.camera.Y + y - configuration.Global.ScreenCenterTileY)
		xTextPos := xMaxPos + 1
		yTextPos := yGeneralPos - 8
		ebitenutil.DebugPrintAt(screen, yPrint, xTextPos, yTextPos)
	}

	vector.StrokeRect(screen, float32(configuration.Global.ScreenCenterTileX*configuration.Global.TileSize), float32(configuration.Global.ScreenCenterTileY*configuration.Global.TileSize), float32(configuration.Global.TileSize+1), float32(configuration.Global.TileSize+1), float32(cameraLineSize), cameraColor, false)

	ySpace := 16
	ebitenutil.DebugPrintAt(screen, "Camera:", xMaxPos+2*configuration.Global.TileSize, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("(", g.camera.X, ",", g.camera.Y, ")"), xMaxPos+2*configuration.Global.TileSize+configuration.Global.TileSize/2, ySpace)

	ebitenutil.DebugPrintAt(screen, "Character:", xMaxPos+2*configuration.Global.TileSize, 3*ySpace)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("(", g.character.X, ",", g.character.Y, ")"), xMaxPos+2*configuration.Global.TileSize+configuration.Global.TileSize/2, 4*ySpace)
}
