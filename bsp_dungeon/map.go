package bsp_dungeon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	for y := 0; y <= level.height; y++ {
		line := []int{}
		for x := 0; x <= level.width; x++ {
			defaultVal, _ := strconv.Atoi(" ")
			line = append(line, defaultVal)
		}
		mapContent = append(mapContent, line)
	}

	// réécrit certaine valeur de mapContent : salle et couloir
	for i := 0; i < len(level.roomList); i++ {
		currentRoom := level.roomList[i]
		for y := currentRoom.topLeftY; y <= currentRoom.height; y++ {
			for x := currentRoom.topLeftX; x <= currentRoom.width; x++ {
				mapContent[x][y] = currentRoom.floorContent
			}
		}
	}

	// créé et écrit dans mapFile
	mapFile, err := os.Create("../floor-files/map") //path depuis exécutable
	check(err)
	defer mapFile.Close()

	w := bufio.NewWriter(mapFile)

	n := 0
	for y := 0; y < len(mapContent); y++ {
		fmt.Println(mapContent[y])
		ny, err := w.WriteString(fmt.Sprint(mapContent[y])) // Error somewhere, way to much byte
		check(err)
		n += ny
	}
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()
}
