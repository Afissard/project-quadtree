package bsp_dungeon

import "fmt"

/*
Fonction pour créer et remplir le contenu des couloirs
Algo :
Parcourt l'arbre BSP pour créer un couloir entre chaque node voisine :
1. Descend au bout de l'arbre :
- tant que roomContant = nil -> appel récursif pour chacune des
	deux node fille
- sinon : retourner roomContent
2. Création des couloir en remontant l'arbre
- Pour la node donnée en entrée, créer un couloir entre les deux
	roomContent récupéré (retourné récursivement)
*/
func createCorridor(currentNode *node) (content *room, roomList []*room) {
	if len(currentNode.children) != 0 {
		roomA, roomListA := createCorridor(currentNode.children[0])
		roomB, roomListB := createCorridor(currentNode.children[1])
		roomList = append(roomList, roomListA...)
		roomList = append(roomList, roomListB...)

		corridor := room{
			isRoom:       false,
			floorContent: 2,
		}

		//TODO: centré/positionné aléatoirement le couloir + taille ~fixe
		// possibilité de virage dans les couloir si trop large -> Pythagore ?
		if roomA.topLeftX < roomB.topLeftX {
			corridor.topLeftX = roomA.topLeftX
			corridor.width = roomB.topLeftX - roomA.topLeftX
		} else {
			corridor.topLeftX = roomB.topLeftX
			corridor.width = roomA.topLeftX - roomB.topLeftX
		}

		if roomA.topLeftY < roomB.topLeftY {
			corridor.topLeftY = roomA.topLeftY
			corridor.height = roomB.topLeftY - roomA.topLeftY
		} else {
			corridor.topLeftY = roomB.topLeftY
			corridor.height = roomA.topLeftY - roomB.topLeftY
		}

		currentNode.content = &corridor
		roomList = append(roomList, currentNode.content)
		fmt.Println(roomList)
		return &corridor, roomList
	} else { // si au bout de l'arbre
		return currentNode.content, roomList
	}
}
