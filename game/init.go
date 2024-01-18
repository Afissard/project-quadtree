package game

import (
	"flag"

	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
	// "gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/dungeon"
)

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.character.Init()
	g.camera.Init()
	g.character.X, g.character.Y, g.listPortal = g.floor.Init()
	if configuration.Global.CameraMode == 1 {
		g.camera.X = g.character.X
		g.camera.Y = g.character.Y
	}
	g.ui.Init()
	g.state = mainMenu
	g.zoomLevel = 1
	g.needNewLevel = false

	// variable nécessaire à la sauvegarde
	flag.StringVar(&sauvegarde, "save", "../floor-files/autoSave.json", "select pos + orientation")
	flag.StringVar(&sauvegardeCam, "saveCam", "../floor-files/autoSave.json", "select cam pos")
	flag.Parse()
}

