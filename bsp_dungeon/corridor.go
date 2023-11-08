package bsp_dungeon

import "fmt"

/*
Fonction pour créer et remplire le contenu des couloirs
Algo :
Pacourt l'arbre BSP pour créer un couloir entre chaque node voisine :
1. Descend au bout de l'arbre :
- tant que roomContant = nil ->appel récursif pour chacune des
	deux node fille
- sinon : retourner roomContent
2. Création des couloir en remontant l'arbre
- Pour la node donnée en entrée, créer un couloir entre les deux
	roomContent récupéré (retourné récursivement)
*/
func createCorridor(currentNode *node) (content *room) {
	fmt.Println(currentNode) // juste pour eviter erreur de compilation
	if len(currentNode.children) != 0 {
		roomA := createCorridor(currentNode.children[0])
		roomB := createCorridor(currentNode.children[1])

		corridor := room{
			isRoom:       false,
			floorContent: 0,
		}

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
		return &corridor
	} else { // si au bout de l'arbre
		return currentNode.content
	}
}
