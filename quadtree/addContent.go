package quadtree

import "fmt"

func addContent(currentNode *node, content, targetX, targetY int) {
	//fmt.Printf("x%d y%d node: %v\n", targetX, targetY, currentNode)
	if currentNode.width == 1 && currentNode.height == 1 && currentNode.content == -1 { // assignation de content
		currentNode.content = content
	} else if currentNode.width > 1 || currentNode.height > 1 { // recherche/création d'une branche

		// on s'assure que la taille est pair pour les x/2 dans le calcul des dimension de la node
		tempWidth, tempHeight := currentNode.width, currentNode.height
		if tempWidth%2 != 0 {
			tempWidth++
		}
		if tempHeight%2 != 0 {
			tempHeight++
		}
		//fmt.Printf("width: %d, height: %d\n", tempWidth, tempHeight)

		if targetX < currentNode.topLeftX+tempWidth/2 {
			if targetY < currentNode.topLeftY+tempHeight/2 {
				if currentNode.topLeftNode == nil {
					newNode := node{
						topLeftX: currentNode.topLeftX,
						topLeftY: currentNode.topLeftY,
						width:    tempWidth / 2,
						height:   tempHeight / 2,
						content:  -1,
					}
					currentNode.topLeftNode = &newNode
				}
				addContent(currentNode.topLeftNode, content, targetX, targetY)
			} else {
				if currentNode.bottomLeftNode == nil {
					newNode := node{
						topLeftX: currentNode.topLeftX,
						topLeftY: currentNode.topLeftY + tempHeight/2,
						width:    tempWidth / 2,
						height:   tempHeight / 2,
						content:  -1,
					}
					currentNode.bottomLeftNode = &newNode
				}
				addContent(currentNode.bottomLeftNode, content, targetX, targetY)
			}
		} else {
			if targetY < currentNode.topLeftY+tempHeight/2 {
				if currentNode.topRightNode == nil {
					newNode := node{
						topLeftX: currentNode.topLeftX + tempWidth/2,
						topLeftY: currentNode.topLeftY,
						width:    tempWidth / 2,
						height:   tempHeight / 2,
						content:  -1,
					}
					currentNode.topRightNode = &newNode
				}
				addContent(currentNode.topRightNode, content, targetX, targetY)
			} else {
				if currentNode.bottomRightNode == nil {
					newNode := node{
						topLeftX: currentNode.topLeftX + tempWidth/2,
						topLeftY: currentNode.topLeftY + tempHeight/2,
						width:    tempWidth / 2,
						height:   tempHeight / 2,
						content:  -1,
					}
					currentNode.bottomRightNode = &newNode
				}
				addContent(currentNode.bottomRightNode, content, targetX, targetY)
			}
		}
		// Optimisation du quadtree pour currentNode
		if currentNode.topLeftNode != nil && currentNode.topRightNode != nil &&
			currentNode.bottomLeftNode != nil && currentNode.bottomRightNode != nil {
			// si 4 branches alors on essai d'optimiser
			if currentNode.topLeftNode.content == currentNode.topRightNode.content &&
				currentNode.topLeftNode.content == currentNode.bottomLeftNode.content &&
				currentNode.bottomLeftNode.content == currentNode.bottomRightNode.content &&
				currentNode.topLeftNode.isLeave() && currentNode.topRightNode.isLeave() &&
				currentNode.bottomLeftNode.isLeave() && currentNode.bottomRightNode.isLeave() {
				// mise à jour de node.content -> prend la valeur des nodes enfant
				currentNode.content = currentNode.topLeftNode.content
				// suppression des nodes désormais inutiles
				currentNode.topLeftNode = nil
				currentNode.topRightNode = nil
				currentNode.bottomLeftNode = nil
				currentNode.bottomRightNode = nil
			}
		}
	} else {
		if currentNode.content != -1 {
			err := fmt.Errorf(fmt.Sprintf("réécriture de content de la node : %v", currentNode))
			fmt.Println(err.Error())
		} else {
			panic("-> quadtree malformé !")
		}
	}
}

// test si la node est une feuille
func (n node) isLeave() (answer bool) {
	answer = true
	if n.topLeftNode != nil || n.topRightNode != nil || n.bottomLeftNode != nil || n.bottomRightNode != nil {
		answer = false
	}
	return answer
}
