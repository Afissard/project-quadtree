package quadtree

import "fmt"

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	/*
		En entré : la map (liste 2D)
		En sortie : un quadtree qui stocke la map

		Algo 1:
		1. créer la couche "feuille" de l'arbre quadtree : chaque node récupère le contenu de floorContent
		2. Boucle récursive :
			- si couche 0 atteint (1 seule node à attacher) alors attacher la node restante à la racine.
			- sinon créer des nodes pour rattacher les node précédemment créés (ou aux données de la map à itération n°1)

		(Nouvel) Algo 2 :
		créer la racine, taille de floorContent
		Boucle parcours floorContent :
			- Pour chaque élément appel addContent
		addContent :
			- regarde les coordonnée et cherche la branche qui correspond :
				- si au bout d'une branche et taille<1 -> créer une branche -> récursion dans la branche
				- si au bout d'une branche et taille == 1 -> assigne la valeur -> fin
				- sinon erreur

		TODO: optimiser les quadtree (fonction à part ?):
		Si une node représente 4 fois le même contenu, alors "supprimer" les node d'en dessous,
		node.content prend la valeur du contenu des nodes supprimées, le contenu d'en dessous
		est déterminé grace avec node.width et node.height.
	*/
	/*
		// Algo 1
			if len(floorContent) <= 0 || len(floorContent[0]) <= 0 {
				panic("floorContent vide !")
			} else {
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
				//q = createNodesLayer(nodesList)
				q = createNodesLayer(nodesList)
			}
	*/

	// Algo 2
	totalWidth := len(floorContent[0])
	totalHeight := len(floorContent)
	if totalWidth%2 != 0 { // on s'assure que la taille est pair
		totalWidth++
	} else if totalHeight%2 != 0 {
		totalHeight++
	}
	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1 // contenu à -1 par défaut
	}

	q = Quadtree{
		width:  totalWidth,
		height: totalHeight,
		root:   &rootNode
	}
	for y := 0; y < len(floorContent); y++ {
		for x := 0; x < len(floorContent[y]); x++ {
			addContent(&rootNode, floorContent[y][x], x, y)
		}
	}
	return q
}

/*
(Nouvel) Algo addContent :
regarde les coordonnée et cherche la branche qui correspond :
- si au bout d'une branche et taille<1 -> cherche/créé une branche -> récursion dans la branche
- si au bout d'une branche et taille == 1 -> assigne la valeur -> fin (retourne pointeur de la node ?)
- sinon erreur

BUG: Problème avec assignation du contenu et/ou choix des branches
Voir les panic à la fin de la fonctions il servent à garantir l’intégrité
du quadtree.
La partie du code en commentaire est l’ancienne fonction qui marche pour une map de 6*6 max
Hésite pas à rajouté des test dans le fichier quadtree_test.go
*/
func addContent(currentNode *node, content, targetX, targetY int) {
	if currentNode.width == 1 && currentNode.height == 1 && currentNode.content == -1 { // assignation de content
		currentNode.content = content
	} else if currentNode.width > 1 { // recherche/création d'une branche
		if targetX < currentNode.topLeftX+currentNode.width/2 {
			if targetY < currentNode.topLeftY+currentNode.height/2 {
				if currentNode.topLeftNode == nil {
					newNode := node{
						topLeftX: currentNode.topLeftX,
						topLeftY: currentNode.topLeftY,
						width:    currentNode.width / 2,
						height:   currentNode.height / 2,
						content:  -1,
					}
					currentNode.topLeftNode = &newNode
				}
				addContent(currentNode.topLeftNode, content, targetX, targetY)
			} else {
				if currentNode.bottomLeftNode == nil {
					newNode := node{
						topLeftX: currentNode.topLeftX,
						topLeftY: currentNode.topLeftY + currentNode.height/2,
						width:    currentNode.width / 2,
						height:   currentNode.height / 2,
						content:  -1,
					}
					currentNode.bottomLeftNode = &newNode
				}
				addContent(currentNode.bottomLeftNode, content, targetX, targetY)
			}
		} else {
			if targetY < currentNode.topLeftY+currentNode.height/2 {
				if currentNode.topRightNode == nil {
					newNode := node{
						topLeftX: currentNode.topLeftX + currentNode.width/2,
						topLeftY: currentNode.topLeftY,
						width:    currentNode.width / 2,
						height:   currentNode.height / 2,
						content:  -1,
					}
					currentNode.topRightNode = &newNode
				}
				addContent(currentNode.topRightNode, content, targetX, targetY)
			} else {
				if currentNode.bottomRightNode == nil {
					newNode := node{
						topLeftX: currentNode.topLeftX + currentNode.width/2,
						topLeftY: currentNode.topLeftY + currentNode.height/2,
						width:    currentNode.width / 2,
						height:   currentNode.height / 2,
						content:  -1,
					}
					currentNode.bottomRightNode = &newNode
				}
				addContent(currentNode.bottomRightNode, content, targetX, targetY)
			}
		}
	} else {
		if currentNode.content != -1 {
			panic("-> Erreur : réécriture de currentNode.content")
		} else {
			panic("-> quadtree malformé !")
		}
	}
}


