package dungeon

import (
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/quadtree"
)

/*
Créer la map du donjon dans un nouveau quadtree
*/
func (d *Dungeon) CreateDungeon(width, height int, seed string) (q quadtree.Quadtree, spawnX, spawnY int, listPortal []Portal) {
	setSeed(seed)
	d.bspTree.makeTree(width, height)
	d.quad.Init(d.bspTree.width, d.bspTree.height)
	d.bspTree.root.build(&d.quad)
	spawnX, spawnY = d.bspTree.getSpawn(&d.quad)
	// puisque place object est volontairement placé après, il est possible que le joueur spawn dessus (offre la possibilité de gagné au spawn)
	listPortal = d.bspTree.placeObject(&d.quad)
	return d.quad, spawnX, spawnY, listPortal
}

/*
Fonction d'interface pour créer un donjon dans un quadtree sans utiliser le type Dungeon
(une fois créer la variable Dungeon est donc supprimé)
*/
func NewDungeon(width, height int, seed string) (q quadtree.Quadtree, spawnX, spawnY int, listPortal[]Portal) {
	d := Dungeon{}
	return d.CreateDungeon(width, height, seed)
}
