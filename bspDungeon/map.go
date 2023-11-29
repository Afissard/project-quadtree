package bspDungeon

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

/*
Fonction pour créer un ficher map a partir
du donjon généré précédemment
*/
func (level BSP_tree) ToMap() (mapContent [][]int) {
	/*
		Algo :
		- créer un fichier map
		- parcourt toute les salle/couloir dans level.roomList
			et print le contenu dans le fichier map
		- ferme et sauvegarde le ficher map (créer json pour
			plus de data plus tard ...)
	*/
	mapContent = make([][]int, level.height, level.width)
	// rempli mapContent avec une valeur par défaut
	for y := 0; y < level.height; y++ {
		for x := 0; x < level.width; x++ {
			defaultVal := 1
			mapContent[y][x] = defaultVal
		}
	}

	// réécrit certaine valeur de mapContent : salle et couloir
	for i := 0; i < len(level.roomList); i++ {
		currentRoom := level.roomList[i]
		for y := 0; y < currentRoom.height; y++ {
			for x := 0; x < currentRoom.width; x++ {
				mapContent[y+currentRoom.topLeftY][x+currentRoom.topLeftX] = currentRoom.floorContent
			}
		}
	}
	return mapContent
}

func (level BSP_tree) SaveToMapFile(mapContent [][]int, fileName string) {
	// créé et écrit dans mapFile
	mapFile, err := os.Create("../floor-files/" + fileName) //path depuis exécutable
	check(err)
	defer mapFile.Close()

	w := bufio.NewWriter(mapFile)

	bytesWritten := 0
	for y := 0; y < len(mapContent); y++ {
		//TODO: add modification to convert mapContent value to its equivalent in ASCII
		ny, err := w.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(mapContent[y])), ""), "[]") + "\n")
		check(err)
		bytesWritten += ny
	}
	fmt.Printf("wrote %d bytes\n", bytesWritten)
	w.Flush()
}
