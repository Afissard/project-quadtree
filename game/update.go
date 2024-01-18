package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() (err error) {
	switch g.state {
	case mainMenu:
		err = g.updateMainMenu()
	case inGame:
		err = g.updateInGame()
	case gameOver:
		err = g.updateGameOverScreen()
	case endScreen:
		err = g.updateEndScreen()
	default:
		log.Fatalf("Game state == %d", g.state)
	}
	return err
}

func (g *Game) updateInGame() error {
	// Inputs
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.state = mainMenu
		g.ui.MainMenu.IsActive = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		return ebiten.Termination
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF2) {
		// zoom in
		if g.zoomLevel+1 <= 4 {
			g.zoomLevel++
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		// zoom out
		if g.zoomLevel-1 >= 1 {
			g.zoomLevel--
		}
	}

	if configuration.Global.FloorKind == 3 {
		currentCharacterTile := g.floor.QuadtreeContent.GetTileValue(g.character.X, g.character.Y)
		// Game Over
		if currentCharacterTile == -1 || currentCharacterTile == 2 || currentCharacterTile == 9 {
			g.state = gameOver
			g.ui.GameOver.IsActive = !g.ui.GameOver.IsActive
			g.ui.GameOver.TimeShowed = 60 * 2
			g.ui.GameOver.SourceMort = currentCharacterTile
			g.ui.DungeonLevel.DungeonLevel = 0
		}
		// Nouvel étage
		if currentCharacterTile == 4 {
			g.newLevel()
		}
		// Victoire
		if currentCharacterTile == 5 {
			g.state = endScreen
			g.ui.EndScreen.IsActive = true
			g.ui.DungeonLevel.DungeonLevel = 0
		}
		if currentCharacterTile == 3 {
			g.teleportPlayer()
		}
	}

	// update du jeu
	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y, g.zoomLevel)
	g.ui.Update()

	return nil
}

func (g *Game) updateMainMenu() error {
	// Input
	if g.ui.MainMenu.OptionChoose {
		switch g.ui.MainMenu.MenuList[g.ui.MainMenu.OptionSelected] {
		case g.ui.MainMenu.MenuList[0]:
			// Jouer
			g.state = inGame
			g.ui.MainMenu.IsActive = false
			if configuration.Global.FloorKind == 3 && g.needNewLevel {
				g.needNewLevel = false
				g.newLevel()
			}
		case g.ui.MainMenu.MenuList[1]:
			// Sauvegarde
			g.SaveGame()
		case g.ui.MainMenu.MenuList[2]:
			g.LoadSave()
			g.state = inGame
			g.ui.MainMenu.IsActive = false
		case g.ui.MainMenu.MenuList[3]:
			// Quitte le jeu sans "erreur"
			return ebiten.Termination
		default:
			log.Panic("Menu choisi hors porté")
		}
	}
	// update de l'ui
	g.ui.Update()
	return nil
}

func (g *Game) updateGameOverScreen() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || g.ui.GameOver.TimeShowed <= 0 {
		g.needNewLevel = true
		g.ui.DungeonLevel.DungeonLevel = 0
		g.state = mainMenu
		g.ui.GameOver.IsActive = false
		g.ui.MainMenu.IsActive = true
	}
	g.ui.Update()
	return nil
}

func (g *Game) updateEndScreen() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.needNewLevel = true
		g.ui.DungeonLevel.DungeonLevel = 0
		g.state = mainMenu
		g.ui.EndScreen.IsActive = false
		g.ui.MainMenu.IsActive = true
	}
	g.ui.Update()
	return nil
}
