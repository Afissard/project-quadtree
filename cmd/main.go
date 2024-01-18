package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	// "github.com/Afissard/project-quadtree/configuration"
	// "github.com/Afissard/project-quadtree/game"
	// "github.com/Afissard/project-quadtree/assets"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/assets"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/game"
)

func main() {
	var configFileName string
	flag.StringVar(&configFileName, "config", "config.json", "select configuration file")

	configuration.Load(configFileName)

	assets.Load()

	g := &game.Game{}
	g.Init()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("Quadtree")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
