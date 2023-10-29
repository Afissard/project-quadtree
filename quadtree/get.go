package quadtree

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du quadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	/*
		En entré : contentHolder, les coordonnées de départ (pour le remplir)
		En sortie : contentHolder remplis au maximum avec le contenu de q (Quadtree) 
		en commençant au coordonnées données (contentHolder est une copie, donc pas 
		besoin de le retourné, il est directement modifié à l’exécution)

		TODO: quand MakeFromArray optimisera les quadtree :
		Adapté l'algo pour prendre en compte les modifications réalisé sur la 
		structure des quadtree.
	*/
}
