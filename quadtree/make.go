package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	/*
		En entré : la map (liste 2D)
		En sortie : un quadtree qui stocke la map

		Idée algo (user de fonctions récursives):
		1. déterminer le nombre de node quadtree au bout de l'arbre -> détermine le nombre total "n" de nodes/couches.
		2. créer des nodes pour stocker les valeurs de la map associé aux coordonnées (couche n).
		3. créer des nodes pour rattacher les node précédemment créés (couches n-couches).
		4. si couche 0 atteint (1 seule node à attacher) alors attacher la node restante à la racine.
	*/
	return
}
