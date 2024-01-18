package dungeon

import (
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/quadtree"
)

/*
Fonction pour créer les salles et couloir dans le quadtree
*/
func (n *bspNode) build(quad *quadtree.Quadtree) {
	if n.nodeLeft != nil && n.nodeRight != nil {
		n.nodeLeft.build(quad)
		n.nodeRight.build(quad)
		n.buildCorridor(quad) // relie les deux branches enfant
	} else {
		n.buildRoom(quad)
		// suppression du surplus de branche potentiel (rare cas dù à l'aléatoire)
		if n.nodeLeft != nil {
			n.nodeLeft = nil
		} else if n.nodeRight != nil {
			n.nodeRight = nil
		}
	}
}

/*
Fonction pour construire une salle dans le quadtree
*/
func (n *bspNode) buildRoom(quad *quadtree.Quadtree) {
	/*
		Création d'une salle :
		création de la salle dans la node (modification des coordonnées et dimensions)
		création de la salle dans le quadtree
	*/
	// Construction de la salle dans la node
	currentWidth := n.width
	currentHeight := n.height
	// dimension de la salle (6 est la taille minimal)
	n.width = randomIntBetween(4, n.width-2)
	n.height = randomIntBetween(4, n.height-2)
	// position de la salle
	n.topX = randomIntBetween(1, currentWidth-n.width-1) + n.topX - 1
	n.topY = randomIntBetween(1, currentHeight-n.height-1) + n.topY - 1

	// Construction de la salle dans le quadtree
	for y := n.topY; y <= n.topY+n.height; y++ {
		for x := n.topX; x <= n.topX+n.width; x++ {
			if y == n.topY || y == n.topY+n.height || x == n.topX || x == n.topX+n.width {
				quad.AddNode(x, y, 0)
			} else {
				if randomIntBetween(0, 10) == 0 {
					if randomIntBetween(0, 3) == 3 {
						// Obstacle
						obstacle := randomIntBetween(6, 8)
						quad.AddNode(x, y, obstacle)
					} else {
						// Tile mort
						if randomIntBetween(0, 5) == 5 {
							quad.AddNode(x, y, 9)
						} else {
							quad.AddNode(x, y, 2)
						}
					}
				} else {
					quad.AddNode(x, y, 1)
				}
			}
		}
	}
}

/*
Fonction pour construire un couloir entre deux salles dans le quadtree
*/
func (n *bspNode) buildCorridor(quad *quadtree.Quadtree) {
	/*
		Création de couloirs :
		choisie dans les deux node enfants à lier, deux salle voisine
		création du couloir dans le quadtree à partir des deux node trouvé précédemment
			1. trouvé les coté à relier de chaque salle
			2. tracé un couloir entre les deux (taille 2 à 4)
				- cas simple : une ligne droite peut être tracé
				- cas complexe : il faut faire un virage (face pas assez en face)
					-> virage d'ajustement
					-> virage pour relier une autre face plus proche
	*/
	// Construction du couloir à partir des nodes
	// recherche de deux salle voisine (situé au milieu de la node n)
	entryRoom := n.nodeLeft.searchRoom(n.topX+(n.width/2), n.topY+(n.height/2))
	exitRoom := n.nodeRight.searchRoom(n.topX+(n.width/2), n.topY+(n.height/2))

	if distanceBetween(entryRoom.topX, exitRoom.topX) > distanceBetween(entryRoom.topY, exitRoom.topY) {
		// connection par l'axe X
		if entryRoom.topX > exitRoom.topX {
			entryRoom, exitRoom = exitRoom, entryRoom // inversion des salles
		}
		/*
			Algo :
				1. Détermine la longueur de la ligne sur l'axe X
				2. Coupe en deux cette ligne
				3. Dessine la première partie attaché à la salle "entré" et la seconde à la salle "sortie"
				4. Sur l'axe Y, relie les deux morceau de couloirs
		*/

		corridorHalfLength := (exitRoom.topX - entryRoom.topX + entryRoom.width) / 2
		startPoint := entryRoom.topX + entryRoom.width
		midEntry := entryRoom.topY + int(entryRoom.height/2)
		endPoint := exitRoom.topX
		midExit := exitRoom.topY + int(exitRoom.height/2)
		traceCorridorBetween(quad, startPoint, midEntry, corridorHalfLength, midEntry)
		traceCorridorBetween(quad, corridorHalfLength, midExit, endPoint, midExit)
		traceCorridorBetween(quad, corridorHalfLength, midEntry, corridorHalfLength, midExit)
	} else {
		// connection par l'axe Y
		if entryRoom.topY > exitRoom.topY {
			entryRoom, exitRoom = exitRoom, entryRoom // inversion des salles
		}
		// Même algo que pour X, mais inversé pour l'axe Y

		corridorHalfLength := (exitRoom.topY - entryRoom.topY + entryRoom.height) / 2
		startPoint := entryRoom.topY + entryRoom.height
		midEntry := entryRoom.topX + int(entryRoom.width/2)
		endPoint := exitRoom.topY
		midExit := exitRoom.topX + int(exitRoom.width/2)
		traceCorridorBetween(quad, midEntry, startPoint, midEntry, corridorHalfLength)
		traceCorridorBetween(quad, midExit, corridorHalfLength, midExit, endPoint)
		traceCorridorBetween(quad, midEntry, corridorHalfLength, midExit, corridorHalfLength)
	}
}

/*
Créer un couloir entre deux node dans le quadtree
*/
func traceCorridorBetween(quad *quadtree.Quadtree, minX, minY, maxX, maxY int) {
	/*
	   Créer un couloir entre les deux point donnée
	*/
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			placeTile(quad, x, y, 1)
			placeTile(quad, x-1, y, 0)
			placeTile(quad, x+1, y, 0)
			placeTile(quad, x, y-1, 0)
			placeTile(quad, x, y+1, 0)

		}
	}
}

/*
Place une tuile dans le quadtree aux coordonnée fournies en paramètres
*/
func placeTile(quad *quadtree.Quadtree, x, y, val int) {
	currentTileVal := quad.GetTileValue(x, y)
	if currentTileVal == -1 || currentTileVal == 0 {
		quad.AddNode(x, y, val)
	}
}
