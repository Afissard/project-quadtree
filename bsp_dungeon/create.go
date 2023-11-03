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
func createLevel(seed, nbrDiv, width, height int) (level BSP_tree) {
	level = BSP_tree{
		width:  width,
		height: height,
		alea:   rand.New(rand.NewSource(int64(seed))), // créé la seed à partir du int donné
	}
	level.root, level.roomList = createNode(level.alea, nil, nbrDiv, level.width, level.height, 0, 0)
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
			width1 := alea.Intn(width/2 + 1)
			width2 := width - width1
			x2 := topLeftX + width1
			x1 := topLeftX
			child1, roomList1 := createNode(alea, currentNode, nbrDivLeft-1, width1, height, x1, topLeftY)
			child2, roomList2 := createNode(alea, currentNode, nbrDivLeft-1, width2, height, x2, topLeftY)
			currentNode.children = append(currentNode.children, child1, child2)
			roomList = append(roomList, roomList1...)
			roomList = append(roomList, roomList2...)
		} else {
			height1 := alea.Intn(height/2 + 1)
			height2 := height - height1
			y2 := topLeftY + height1
			y1 := topLeftY
			child1, roomList1 := createNode(alea, currentNode, nbrDivLeft-1, width, height1, topLeftX, y1)
			child2, roomList2 := createNode(alea, currentNode, nbrDivLeft-1, width, height2, topLeftX, y2)
			currentNode.children = append(currentNode.children, child1, child2)
			roomList = append(roomList, roomList1...)
			roomList = append(roomList, roomList2...)
		}
	}
	currentNode.content = createRoom(currentNode, alea)
	roomList = append(roomList, currentNode.content)
	return currentNode, roomList
}

/*
Fonction pour créer des salle ou des couloir
- ajoute des salle au node feuille
- créé un couloir entre les nodes enfant
*/
func createRoom(currentNode *node, alea *rand.Rand) (newRoom *room) {
	newRoom = &room{}

	isRoom := true
	if len(currentNode.children) > 0 { // couloir
		newRoom.isRoom = false
	} else {
		newRoom = &room{ // salle
			isRoom: isRoom,
			// topLeftX:     currentNode.topLeftX + alea.Intn(int(currentNode.width/4)+1),
			// topLeftY:     currentNode.topLeftY + alea.Intn(int(currentNode.height/4)+1),
			topLeftX:     currentNode.topLeftX,
			topLeftY:     currentNode.topLeftY,
			width:        4, // TODO: trouvé une équation
			height:       4,
			floorContent: 0,
		}
	}
	return newRoom
}
