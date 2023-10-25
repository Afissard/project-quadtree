package floor

import (
	"bufio"
	"os"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind { // INFO: dépend du config.json
	case fromFileFloor: // TODO: utilise d'abord un fichier pour la map avant algo quadtree
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case quadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	/*
		TODO: here
		get the data from floor-files/filename
		convert the data to fit floorContent
	*/
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	file_scan := bufio.NewScanner(file)
	numLigne := 0
	for file_scan.Scan() {
		// TODO changer de ligne dans floorContent[numLine][dataLigne]
		numLigne++
	}

	file.Close()
	return floorContent
}
