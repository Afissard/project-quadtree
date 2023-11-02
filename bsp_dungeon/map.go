package bsp_dungeon

import "fmt"

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
	fmt.Println("map")
}
