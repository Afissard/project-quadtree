package dungeon

/*
Initialise et créer l'arbre BSP
*/
func (b *bspTree) makeTree(width, height int) {
	b.width = width
	b.height = height
	b.topX = 0
	b.topY = 0
	b.root = &bspNode{
		width:  b.width,
		height: b.height,
		topX:   b.topX,
		topY:   b.topY,
	}
	b.root.bspTreeGen()
}

/*
Créer récursivement un arbre binaire représentant le donjon
*/
func (n *bspNode) bspTreeGen() {
	/*
		Si l'espace peut être partitionné (rectangle aux côtés de dimension supérieur à 16):
			1. le coupe aléatoirement en deux, soit sur la largeur soit sur la longueur
			2. appel récursif
	*/
	if n.width >= 16 && n.height >= 16 {
		if randomIntBetween(0, 1) == 1 {
			n.widthSpacePartition()
		} else {
			n.heightSpacePartition()
		}
	} else if n.width >= 16 {
		n.widthSpacePartition()
	} else if n.height >= 16 {
		n.heightSpacePartition()
	}
}

/*
Sert à la construction de l'arbre binaire en opérant la partition sur l'axe X
*/
func (n *bspNode) widthSpacePartition() {
	// initialisation des variables
	height1, height2 := n.height, n.height
	y1, y2 := n.topY, n.topY
	// partition de la node sur la largeur
	width1 := randomIntBetween(int(float32(n.width)*0.40), int(float32(n.width)*0.60))
	width2 := n.width - width1
	x1 := n.topX
	x2 := n.topX + width1
	// création des nodes enfant et appel récursif de la génération
	n.nodeLeft = &bspNode{width: width1, height: height1, topX: x1, topY: y1}
	n.nodeRight = &bspNode{width: width2, height: height2, topX: x2, topY: y2}
	n.nodeLeft.bspTreeGen()
	n.nodeRight.bspTreeGen()
}

/*
Sert à la construction de l'arbre binaire en opérant la partition sur l'axe Y
*/
func (n *bspNode) heightSpacePartition() {
	// initialisation des variables
	width1, width2 := n.width, n.width
	x1, x2 := n.topX, n.topX
	// partition de la node sur la hauteur
	height1 := randomIntBetween(int(float32(n.height)*0.40), int(float32(n.height)*0.60))
	height2 := n.height - height1
	y1 := n.topY
	y2 := n.topY + height1
	// création des nodes enfant et appel récursif de la génération
	n.nodeLeft = &bspNode{width: width1, height: height1, topX: x1, topY: y1}
	n.nodeRight = &bspNode{width: width2, height: height2, topX: x2, topY: y2}
	n.nodeLeft.bspTreeGen()
	n.nodeRight.bspTreeGen()
}
