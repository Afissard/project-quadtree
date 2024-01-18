package game

import (
	"encoding/json"
	"log"
	"os"

	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
)

func check(err error) {
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
}

type CharacterPos struct {
	X                   int
	Y                   int
	Orientation         int
	AnimationStep       int
	XInc                int
	YInc                int
	Moving              bool
	Shift               int
	AnimationFrameCount int
}

// Sauvegarde l'état actuel du jeux : map, joueur et orientation du joueur
func (g *Game) SaveGame() {
	filePath := configuration.Global.FilePath + configuration.Global.SaveFile
	// appel de saveMap
	err := g.floor.SaveMap(filePath)
	check(err)
	// Save info jeu dans un json
	pos := CharacterPos{
		X:                   g.character.X,
		Y:                   g.character.Y,
		Orientation:         g.character.Orientation,
		AnimationStep:       g.character.AnimationStep,
		XInc:                g.character.XInc,
		YInc:                g.character.YInc,
		Moving:              g.character.Moving,
		Shift:               g.character.Shift,
		AnimationFrameCount: g.character.AnimationFrameCount,
	}
	b, err := json.Marshal(pos)
	check(err)
	err = os.WriteFile(filePath+".json", b, 0777)
	check(err)
}

func (g *Game) LoadSave() {
	// var sauvegarde string
	// flag.StringVar(&sauvegarde, "save", "../floor-files/autoSave.json", "select pos + orientation")
	// var sauvegardeCam string
	// flag.StringVar(&sauvegardeCam, "saveCam", "../floor-files/autoSave.json", "select cam pos")
	// flag.Parse()

	tempFloorFile := configuration.Global.FloorFile
	configuration.Global.FloorFile = configuration.Global.SaveFile

	g.camera.Load(sauvegardeCam)
	g.character.Load(sauvegarde)
	g.floor.Init()
	g.zoomLevel = 1
	configuration.Global.FloorFile = tempFloorFile
}
