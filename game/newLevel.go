package game

import (
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/dungeon"
)

func (g *Game) newLevel() {
	g.floor.QuadtreeContent, g.character.X, g.character.Y, g.listPortal = dungeon.NewDungeon(configuration.Global.FixedGenWidth, configuration.Global.FixedGenHeight, "fromCurrentSeed")
	if configuration.Global.CameraMode == 1 {
		g.camera.X = g.character.X
		g.camera.Y = g.character.Y
	}
	g.ui.DungeonLevel.DungeonLevel++
	g.ui.DungeonLevel.IsActive = true
	g.ui.DungeonLevel.TimeShowed = 60 * 1
}
