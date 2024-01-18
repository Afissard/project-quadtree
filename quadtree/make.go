package quadtree

/*
Fonction MakeFromArray (fichier make.go)

Entrées: floorContent -> un tableau en 2 dimensions d'entiers
Sorties:  q Quadtree -> une structure de l'arbre finale contenant le terrain

Construit un quadtree à partir d'une tableau de donnée représentant le terrain voulue.
La fonction créer le quadtree de manière récursive, en stockant les différents
types de terrains et leur emplacement dans les feuilles de l'arbre. Si il n'y a pas
de terrain, la valeur est -1 sinon en fonction du type de terrain, un entier associé
à celui-ci.
*/
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	totalHeight := len(floorContent)
	totalWidth := len(floorContent[0])
	for y := 1; y < totalHeight; y++ {
		if len(floorContent[y]) > totalWidth { //détermine la largeur maximale
			totalWidth = len(floorContent[y])
		}
	}

	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1, // contenu à -1 par défaut
	}

	q = Quadtree{
		Width:  totalWidth,
		Height: totalHeight,
		root:   &rootNode,
	}

	for y := 0; y < len(floorContent); y++ {
		for x := 0; x < len(floorContent[y]); x++ {
			q.AddNode(x, y, floorContent[y][x])
		}
	}
	return q
}
