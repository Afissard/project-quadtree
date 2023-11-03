package bsp_dungeon

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
func (level BSP_tree) toMapFile(fileName string) {
	/*
		Algo :
		- créer un fichier map
		- parcourt toute les salle/couloir dans level.roomList
			et print le contenu dans le fichier map
		- ferme et sauvegarde le ficher map (créer json pour
			plus de data plus tard ...)
	*/
	mapContent := [][]int{}
	// rempli mapContent avec une valeur par défaut
	for y := 0; y < level.height; y++ {
		line := []int{}
		for x := 0; x < level.width; x++ {
			defaultVal := 1
			line = append(line, defaultVal)
		}
		mapContent = append(mapContent, line)
	}

	// réécrit certaine valeur de mapContent : salle et couloir
	for i := 0; i < len(level.roomList); i++ {
		currentRoom := level.roomList[i]
		//fmt.Println(currentRoom.topLeftX, currentRoom.topLeftY)
		for y := currentRoom.topLeftY; y < currentRoom.height; y++ { //BUG: x and y doesn't take the right value
			fmt.Println(currentRoom.topLeftY, y) //currentRoom.topLeftY/X est remis à 0 ?!
			for x := currentRoom.topLeftX; x < currentRoom.width; x++ {
				mapContent[y][x] = currentRoom.floorContent
				//fmt.Println(x, y)
			}
		}
	}

	// créé et écrit dans mapFile
	mapFile, err := os.Create("../floor-files/map") //path depuis exécutable
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
