package quadtree

import "fmt"

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	totalWidth := len(floorContent[0])
	totalHeight := len(floorContent)
	if totalWidth%2 != 0 { // on s'assure que la taille est pair
		totalWidth++
	}
	if totalHeight%2 != 0 {
		totalHeight++
	}
	fmt.Printf("width: %d, height: %d\n", totalWidth, totalHeight)

	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1, // contenu à -1 par défaut
	}

	q = Quadtree{
		width:  totalWidth,
		height: totalHeight,
		root:   &rootNode,
	}
	for y := 0; y < len(floorContent); y++ {
		for x := 0; x < len(floorContent[y]); x++ {
			addContent(&rootNode, floorContent[y][x], x, y)
		}
	}
	q.root.content = optimize(q.root)
	return q
}

func addContent(currentNode *node, content, targetX, targetY int) {
	fmt.Printf("x%d y%d node: %v\n", targetX, targetY, currentNode)
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
