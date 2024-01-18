package character

import (
	// "github.com/Afissard/project-quadtree/configuration"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Character) Update(blocking [4]bool) {
	// if !c.Moving {
	// 	if ebiten.IsKeyPressed(ebiten.KeyRight) {
	// 		c.Orientation = orientedRight
	// 		if !blocking[1] {
	// 			c.XInc = 1
	// 			c.Moving = true
	// 		}
	// 	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
	// 		c.Orientation = orientedLeft
	// 		if !blocking[3] {
	// 			c.XInc = -1
	// 			c.Moving = true
	// 		}
	// 	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
	// 		c.Orientation = orientedUp
	// 		if !blocking[0] {
	// 			c.YInc = -1
	// 			c.Moving = true
	// 		}
	// 	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
	// 		c.Orientation = orientedDown
	// 		if !blocking[2] {
	// 			c.YInc = 1
	// 			c.Moving = true
	// 		}
	// 	}
	// } else {
	// 	c.AnimationFrameCount++
	// 	if c.AnimationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
	// 		c.AnimationFrameCount = 0
	// 		shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
	// 		c.Shift += shiftStep
	// 		c.AnimationStep = -c.AnimationStep
	// 		if c.Shift > configuration.Global.TileSize-shiftStep {
	// 			c.Shift = 0
	// 			c.Moving = false
	// 			c.X += c.XInc
	// 			c.Y += c.YInc
	// 			c.XInc = 0
	// 			c.YInc = 0
	// 		}
	// 	}
	// }
	if !c.Moving {
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			c.Orientation = orientedRight
			if !blocking[1] {
				c.XInc = 1
				c.Moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			c.Orientation = orientedLeft
			if !blocking[3] {
				c.XInc = -1
				c.Moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			c.Orientation = orientedUp
			if !blocking[0] {
				c.YInc = -1
				c.Moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			c.Orientation = orientedDown
			if !blocking[2] {
				c.YInc = 1
				c.Moving = true
			}
		}
	} else {
		c.AnimationFrameCount++
		if c.AnimationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			c.AnimationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			c.Shift += shiftStep
			c.AnimationStep = -c.AnimationStep
			if c.Shift > configuration.Global.TileSize-shiftStep {
				c.Shift = 0
				c.Moving = false
				c.X += c.XInc
				c.Y += c.YInc
				c.XInc = 0
				c.YInc = 0
			}
		}
	}
}
