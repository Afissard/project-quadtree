package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	/*
		En entré : la map (liste 2D)
		En sortie : un quadtree qui stocke la map

		Idée algo (user de fonctions récursives):
		1. déterminer le nombre de nodes quadtree à lié -> nombre total "n" de nodes pour cette couche.
		3. créer des nodes pour rattacher les node précédemment créés (ou aux données de la map à itération n°1).
		4. si couche 0 atteint (1 seule node à attacher) alors attacher la node restante à la racine.
	*/
	emptiness := false
	if len(floorContent) <= 0 {
		emptiness = true
	} else if len(floorContent[0]) <= 0 {
		emptiness = true
	}

	if !emptiness {
		// création des node pour la couche "feuille"
		nodesList := [][]node{}
		for y := 0; y < len(floorContent); y++ {
			nodesLine := []node{}
			for x := 0; x < len(floorContent[y]); x++ {
				currentNode := node{
					topLeftX: x,
					topLeftY: y,
					width:    1,
					height:   1,
					content:  floorContent[y][x],
				}
				// ligne de nodes
				nodesLine = append(nodesLine, currentNode)
			}
			// tableau 2D de nodes
			nodesList = append(nodesList, nodesLine)
		}
		q = createNodesLayer(nodesList)
	}
	return q
}

func createNodesLayer(nodesList [][]node) (q Quadtree) {
	/*
		Sert à makeFromArray, associe les node entre elles dans un arbre récursif
	*/
	// calcul le nombre de nodes (en supposant que tout len(floorContent[x] = len(floorContent[0]) -> map rectangulaire)
	var nbrNodes float32 = (float32(len(nodesList) * len(nodesList[0]))) / 4

	if nbrNodes <= float32(1/4) { // on sait jamais
		panic("nbrNodes <=0 ")
	} else if nbrNodes > float32(1.0) { // plus d'une node -> création d'une couche
		newNodesList := [][]node{}
		for y := 0; y < len(nodesList); y += 2 { // +2 au lieu de +1 car une node couvre 4 nodes
			nodesLine := []node{}
			for x := 0; x < len(nodesList[y]); x += 2 {
				currentNode := node{
					topLeftX: x,
					topLeftY: y,
					width:    nodesList[y][x].width * 2, // TODO: attention aux potentiels tuiles vide ?
					height:   nodesList[y][x].height * 2,
				}

				// lie les nodes, en faisant attention au potentiel nodes inexistante
				currentNode.topLeftNode = &nodesList[y][x]
				if x+1 > len(nodesList[y]) {
					currentNode.topRightNode = &nodesList[y][x+1]
				}
				if y+1 > len(nodesList) {
					currentNode.bottomLeftNode = &nodesList[y+1][x]
					if x+1 > len(nodesList[y]) {
						currentNode.bottomRightNode = &nodesList[y+1][x+1]
					}
				}
				// ligne de nodes
				nodesLine = append(nodesLine, currentNode)
			}
			// tableau 2D de nodes
			nodesList = append(nodesList, nodesLine)
		}
		q = createNodesLayer(newNodesList)
	} else { // si une seule node existe
		q.width = nodesList[0][0].width
		q.height = nodesList[0][0].height
		q.root = &nodesList[0][0]
	}
	return q
}
