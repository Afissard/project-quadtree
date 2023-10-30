package quadtree

import "fmt"

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du quadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	/*
		En entré : contentHolder, les coordonnées de départ (pour le remplir)
		En sortie : contentHolder remplis au maximum avec le contenu de q (Quadtree)
		en commençant au coordonnées données (contentHolder est une copie, donc pas
		besoin de le retourné, il est directement modifié à l’exécution)

		Algo :
		1. Remonte une branche de l'arbre jusqu'a ce que node.content != nil
			- si coordonnées dans contentHolder alors append
			- sinon échec
		2. Si toutes les branches ont été explorées et contentHolder pas remplit,
			alors le remplir de tuiles vide ?

		TODO: quand MakeFromArray optimisera les quadtree :
		Adapté l'algo pour prendre en compte les modifications réalisé sur la
		structure des quadtree.
	*/
	getNodeContent(q.root, topLeftX, topLeftY, contentHolder)
	// TODO : ajouter cases vide ?
	fmt.Println("\nexit with", contentHolder)
}

type nodeStalker struct {
	previousNode     *node
	currentNode      *node // permet d’accéder aux coordonnée de cette node
	tl, tr, bl, br   bool  // branche déjà emprunté : true, sinon false
	previousTrackPos int
	currentTrackPos  int
}

func nodesTrack(track []nodeStalker, trackPos int, forward bool) (goToNode *node) {
	if forward { // en avant
		if !track[trackPos].tl && track[trackPos].currentNode.topLeftNode != nil { // go top left
			track = append(track, nodeStalker{
				previousNode:     track[trackPos].currentNode,
				currentNode:      track[trackPos].currentNode.topLeftNode,
				previousTrackPos: trackPos,
				currentTrackPos:  len(track) - 1,
			})
			track[trackPos].tl = true // bloque l'accès
			return track[trackPos].currentNode

		} else if !track[trackPos].tr && track[trackPos].currentNode.topRightNode != nil { // go top right
			track = append(track, nodeStalker{
				previousNode:     track[trackPos].currentNode,
				currentNode:      track[trackPos].currentNode.topRightNode,
				previousTrackPos: trackPos,
				currentTrackPos:  len(track) - 1})
			track[trackPos].tr = true // bloque l'accès
			return track[trackPos].currentNode

		} else if !track[trackPos].bl && track[trackPos].currentNode.bottomLeftNode != nil { // go bottom left
			track = append(track, nodeStalker{
				previousNode:     track[trackPos].currentNode,
				currentNode:      track[trackPos].currentNode.bottomLeftNode,
				previousTrackPos: trackPos,
				currentTrackPos:  len(track) - 1})
			track[trackPos].bl = true // bloque l'accès
			return track[trackPos].currentNode

		} else if !track[trackPos].br && track[trackPos].currentNode.bottomRightNode != nil { // go bottom right
			track = append(track, nodeStalker{
				previousNode:     track[trackPos].currentNode,
				currentNode:      track[trackPos].currentNode.bottomRightNode,
				previousTrackPos: trackPos,
				currentTrackPos:  len(track) - 1})
			track[trackPos].br = true // bloque l'accès
			return track[trackPos].currentNode

		} else { // en avant pas possible alors en arrière
			return nodesTrack(track, trackPos, false)
		}
	} else { // en arrière
		return track[trackPos].previousNode
	}

}

func getNodeContent(currentNode *node, topLeftX, topLeftY int, contentHolder [][]int) {
	/*
		Algo :
		Remonte une branche de l'arbre jusqu'a ce que node.content != nil
		- si coordonnées dans contentHolder alors append
		- sinon échec
	*/
	fmt.Println("\nnode x:", currentNode.topLeftX, "y:", currentNode.topLeftY)

	if currentNode.topLeftNode == nil { // si il n'y a pas de node attaché au bout
		fmt.Println("try append at x:", currentNode.topLeftX, "y:", currentNode.topLeftY)
		if currentNode.topLeftY >= topLeftY &&
			(currentNode.topLeftY-topLeftY) < len(contentHolder) &&
			currentNode.topLeftX >= topLeftX &&
			(currentNode.topLeftX-topLeftX) < len(contentHolder[0]) { // si node dans coordonnée de contentHolder (supposé rectangulaire)

			// alors ajout du contenu de la node au bon endroit dans contentHolder
			contentHolder[currentNode.topLeftY-topLeftY][currentNode.topLeftX-topLeftX] = currentNode.content
			// HERE: retour en arrière vers les autres branches
		}
	} else { // si il y a des nodes attaché, alors, continué de remonter
		/*
			FIXME: pas de retour en arrière dans les branches ...
			Solutions :
			- programmation "dynamique" : array "Petit Poucet" ?
			- algo backward ? (on doit trouver tout les résultats, pas qu'un seul)
		*/
		if currentNode.topLeftNode != nil {
			fmt.Println("go to node x:", currentNode.topLeftNode.topLeftX, "y:", currentNode.topLeftNode.topLeftY)
			getNodeContent(currentNode.topLeftNode, topLeftX, topLeftY, contentHolder)
		}
		if currentNode.topRightNode != nil {
			fmt.Println("go to node x:", currentNode.topRightNode.topLeftX, "y:", currentNode.topRightNode.topLeftY)
			getNodeContent(currentNode.topRightNode, topLeftX, topLeftY, contentHolder)
		}
		if currentNode.bottomLeftNode != nil {
			fmt.Println("go to node x:", currentNode.bottomLeftNode.topLeftX, "y:", currentNode.bottomLeftNode.topLeftY)
			getNodeContent(currentNode.bottomLeftNode, topLeftX, topLeftY, contentHolder)
		}
		if currentNode.bottomRightNode != nil {
			fmt.Println("go to node x:", currentNode.bottomRightNode.topLeftX, "y:", currentNode.bottomRightNode.topLeftY)
			getNodeContent(currentNode.bottomRightNode, topLeftX, topLeftY, contentHolder)
		}
	}
}
