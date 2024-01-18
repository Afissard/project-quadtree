package quadtree

/*
Fonction GetContent (fichier get.go)

Entrées: topLeftX / topLeftY int -> coordonnées de la feuille (point de départ de l'itération),
contentHolder [][]int -> vide au départ
Sorties: contentHolder [][]int -> tableau en 2 dimensions d'entiers avec le contenu de chaque tuile

GetContent remplit le tableau contentHolder (qui représente un terrain dont la case le plus en haut
à gauche a pour coordonnées (topLeftX, topLeftY)) à partir du quadtree q.
GetContent permet de remonter une branche de l'arbre jusqu'a ce que node.content != nil
	- si coordonnées dans contentHolder alors append
	- sinon échec
Si toutes les branches ont été explorées et contentHolder pas remplit,
alors le remplir de tuiles vide (tuile par default)
*/
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	for y := 0; y < len(contentHolder); y++ {
		for x := 0; x < len(contentHolder[y]); x++ {
			targetY, targetX := y+topLeftY, x+topLeftX
			contentHolder[y][x] = q.Seek(targetX, targetY).content
		}
	}
}
