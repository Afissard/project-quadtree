/*
Algo 1, défectueux pour map supérieur à 6*6:
1. créer la couche "feuille" de l'arbre quadtree : chaque node récupère le contenu de floorContent
2. Boucle récursive :
	- si couche 0 atteint (1 seule node à attacher) alors attacher la node restante à la racine.
	- sinon créer des nodes pour rattacher les node précédemment créés (ou aux données de la map à itération n°1)
*/
/*
func createNodesLayer(nodesList [][]node) (q Quadtree) {
	//Sert à makeFromArray, associe les node entre elles dans un arbre récursif
	//BUG: ne marche pas avec grande map...
	fmt.Println("\n---New Call---")
	// test racine de quadtree : si une seule node existe
	if len(nodesList) == 1 && len(nodesList[0]) == 1 {
		q.width = nodesList[0][0].width
		q.height = nodesList[0][0].height
		q.root = &nodesList[0][0]
		return q
	}
	// else implicite : calcul d'une nouvelle couche
	newNodesList := [][]node{}
	for y := 0; y < len(nodesList); y += 2 { // +2 au lieu de +1 car une node couvre (2*2=4) nodes
		nodesLine := []node{}
		for x := 0; x < len(nodesList[y]); x += 2 {
			currentNode := node{
				topLeftX: x,
				topLeftY: y,
				width:    nodesList[y][x].width * 2, // TODO: attention aux potentiels tuiles vide ?
				height:   nodesList[y][x].height * 2,
			}

			// lie les nodes, en faisant attention aux potentiels nodes inexistante
			//FIXME: ne vas pas au delà de x6 y6

			currentNode.topLeftNode = &nodesList[y][x]
			if x+1 < len(nodesList[y]) {
				currentNode.topRightNode = &nodesList[y][x+1]
			}
			if y+1 < len(nodesList) {
				currentNode.bottomLeftNode = &nodesList[y+1][x]
				if x+1 < len(nodesList[y]) {
					currentNode.bottomRightNode = &nodesList[y+1][x+1]
				}
			}
			// ligne de nodes
			nodesLine = append(nodesLine, currentNode)
		}
		// tableau 2D de nodes
		newNodesList = append(newNodesList, nodesLine)
	}
	fmt.Println("\nListe:")
	for y := 0; y < len(newNodesList); y++ {
		fmt.Println("line :", y)
		for x := 0; x < len(newNodesList[y]); x++ {
			fmt.Println(newNodesList[y][x])
		}
	}

	q = createNodesLayer(newNodesList)
	return q
}
*/

