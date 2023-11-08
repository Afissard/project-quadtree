package bsp_dungeon

import (
	"math/rand"
)

/*
Créé un arbre BSP ainsi que son contenu
soit des salle et des couloir (et autres ajouts)

Arguments :
seed 			: seed de génération pour cet étage
nbrDiv 			: nombre de division de l'espace exécuté par le programme (soit nombre de salle = 2^nbrDiv)
width, height 	: dimensions de l'étage
*/
func CreateLevel(seed string, nbrDiv, width, height int) (level BSP_tree) {
	// seed conversion
	seedRune := []rune(seed)
	seedInt := 0
	for i := 0; i < len(seedRune); i++ {
		currentSeedChar := int(seedRune[i])
		seedInt += currentSeedChar
	}
	// level initialisation
	level = BSP_tree{
		width:  width,
		height: height,
		alea:   rand.New(rand.NewSource(int64(seedInt))), // créé la seed à partir du int donné
	}
	level.root, level.roomList = createNode(level.alea, nil, nbrDiv, level.width, level.height, 0, 0)

	corridorList := []*room{}
	level.root.content, corridorList = createCorridor(level.root) // work in progress
	level.roomList = append(level.roomList, corridorList...)

	return level
}

/*
Fonction récursive pour générer procéduralement
un donjon avec un algo de Binary Space Partition
*/
func createNode(alea *rand.Rand, parent *node, nbrDivLeft, width, height, topLeftX, topLeftY int) (currentNode *node, roomList []*room) {
	// création de la node
	currentNode = &node{
		parent:   parent,
		width:    width,
		height:   height,
		topLeftX: topLeftX,
		topLeftY: topLeftY,
	}
	// BSP
	if nbrDivLeft > 0 {
		if alea.Intn(2) == 1 {
			width1 := alea.Intn(int(float32(width)*0.60-float32(width)*0.40)) + int(float32(width)*0.40) // aléa entre width*25 et width*75
			width2 := width - width1
			x1 := topLeftX
			x2 := topLeftX + width1

			child1, roomList1 := createNode(alea, currentNode, nbrDivLeft-1, width1, height, x1, topLeftY)
			child2, roomList2 := createNode(alea, currentNode, nbrDivLeft-1, width2, height, x2, topLeftY)
			currentNode.children = append(currentNode.children, child1, child2)
			roomList = append(roomList, roomList1...)
			roomList = append(roomList, roomList2...)
		} else {
			height1 := alea.Intn(int(float32(height)*0.60-float32(height)*0.40)) + int(float32(height)*0.40)
			height2 := height - height1
			y1 := topLeftY
			y2 := topLeftY + height1

			child1, roomList1 := createNode(alea, currentNode, nbrDivLeft-1, width, height1, topLeftX, y1)
			child2, roomList2 := createNode(alea, currentNode, nbrDivLeft-1, width, height2, topLeftX, y2)
			currentNode.children = append(currentNode.children, child1, child2)
			roomList = append(roomList, roomList1...)
			roomList = append(roomList, roomList2...)
		}
	}
	currentNode.content = createRoom(currentNode, alea)
	if currentNode.content.isRoom == true {
		roomList = append(roomList, currentNode.content)
	}
	return currentNode, roomList
}
