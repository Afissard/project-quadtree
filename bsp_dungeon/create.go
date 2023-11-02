package bsp_dungeon

import (
	"fmt"
	"math/rand"
)

/*
Créé un arbre BSP ainsi que son contenu
soit des salle et des couloir (et autres ajouts)

Arguments :
seed 			: seed de génération pour cet étage
nbrDiv 			: nombre de division de l'espace exécuté par le programme (soit nombre de salle * 2)
width, height 	: dimensions de l'étage
*/
func createLevel(seed, nbrDiv, width, height int) (level BSP_tree) {
	level = BSP_tree{
		width:  width,
		height: height,
		alea:   rand.New(rand.NewSource(int64(seed))), // créé la seed à partir du int donné
	}
	level.root = createNode(level.alea, nil, nbrDiv, level.width, level.height, 0, 0)
	// ajout des pièces et couloir
	return level
}

/*
Fonction récursive pour générer procéduralement
un donjon avec un algo de Binary Space Partition
*/
func createNode(alea *rand.Rand, parent *node, nbrDivLeft, width, height, topLeftX, topLeftY int) (currentNode *node) {
	currentNode = &node{
		parent:   parent,
		width:    width,
		height:   height,
		topLeftX: topLeftX,
		topLeftY: topLeftY,
	}

	if nbrDivLeft > 0 {
		if alea.Intn(2) == 1 {
			width1 := alea.Intn(width/2 + 1)
			width2 := width - width1
			x2 := topLeftX + width1
			x1 := topLeftX
			currentNode.children = append(currentNode.children, createNode(alea, currentNode, nbrDivLeft-1, width1, height, x1, topLeftY))
			currentNode.children = append(currentNode.children, createNode(alea, currentNode, nbrDivLeft-1, width2, height, x2, topLeftY))
		} else {
			height1 := alea.Intn(height/2 + 1)
			height2 := height - height1
			y2 := topLeftY + height1
			y1 := topLeftY
			currentNode.children = append(currentNode.children, createNode(alea, currentNode, nbrDivLeft-1, width, height1, topLeftX, y1))
			currentNode.children = append(currentNode.children, createNode(alea, currentNode, nbrDivLeft-1, width, height2, topLeftX, y2))
		}
	}
	fmt.Println(currentNode)
	return currentNode
}

/*
Fonction qui parcourt les nodes générées :
- ajoute des salle au node feuille
- créé un couloir entre les nodes enfant
*/
func addRoom(alea *rand.Rand, currentNode *node) {
	if len(currentNode.children) >= 0 { // couloir
		// TODO: find algo to create corridor (need math)
		currentNode.content = &room{
			isRoom: false,
		}
		addRoom(alea, currentNode.children[0])
		addRoom(alea, currentNode.children[1])
	} else { // salle
		currentNode.content = &room{
			isRoom: true,
		}
	}
}
