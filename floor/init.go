package floor

import (
	"bufio"
	"os"
	"strconv"

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
	file_scan.Split(bufio.ScanRunes) // scan char par char
	lineContent := []int{}

	for file_scan.Scan() {
		str_content := file_scan.Text() // get content
		if err != nil {
			panic(err)
		}

		if str_content == "\n" { // if new line, do not add it to content
			floorContent = append(floorContent, lineContent)
			lineContent = nil
		} else {
			content, err := strconv.Atoi(str_content) // convert str to int
			if err != nil {
				panic(err)
			}
			lineContent = append(lineContent, content)
		}
	}
	// else implicit for the last line (maybe found a better solution)
	floorContent = append(floorContent, lineContent) // not a good idea

	file.Close()
	return floorContent
}
