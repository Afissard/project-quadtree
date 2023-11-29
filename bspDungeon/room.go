package bspDungeon

import "math/rand"

/*
Fonction pour créer des salle ou des couloir (node)
- ajoute des salle au node feuille
- créé un couloir entre les nodes enfant (les couloir seront générer plus tard dans le programme)
- le contenu des salle/couloir sera déterminé plus tard aussi
*/
func createRoom(currentNode *node, alea *rand.Rand) (newRoom *room) {
	newRoom = &room{}

	isRoom := true
	if len(currentNode.children) > 0 { // couloir
		newRoom.isRoom = false
	} else {
		newRoom = &room{ // salle
			isRoom:       isRoom,
			width:        alea.Intn(int(float32(currentNode.width)*0.70-float32(currentNode.width)*0.30)) + int(float32(currentNode.width)*0.30),
			height:       alea.Intn(int(float32(currentNode.height)*0.70-float32(currentNode.height)*0.30)) + int(float32(currentNode.height)*0.30),
			floorContent: 0,
		}
		newRoom.topLeftX = alea.Intn(currentNode.width-newRoom.width) + currentNode.topLeftX
		newRoom.topLeftY = alea.Intn(currentNode.height-newRoom.height) + currentNode.topLeftY
	}
	return newRoom
}
