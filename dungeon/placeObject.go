package dungeon

import (
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/quadtree"
)

/*
getSpawn détermine et sécurise en plaçant une tuile de sol, le spawn du joueur
*/
func (b *bspTree) getSpawn(quad *quadtree.Quadtree) (spawnX, spawnY int) {
	spawnRoom := b.root.getRandomRoom()
	spawnX = spawnRoom.topX + int(spawnRoom.width/2)
	spawnY = spawnRoom.topY + int(spawnRoom.height/2)
	quad.AddNode(spawnX, spawnY, 1) // Assure un spawn sans danger
	return spawnX, spawnY
}

/*
placeObjet sert à placer les différente tuiles interactives sur la map (excepté les piège qui sont déjà créer)
*/
func (b *bspTree) placeObject(quad *quadtree.Quadtree) (listPortal []Portal) {
	// place une sortie
	for i := 1; i <= randomIntBetween(1, 4); i++ {
		room := b.root.getRandomRoom()
		posX := randomIntBetween(room.topX+1, room.topX+room.width-1)
		posY := randomIntBetween(room.topY+1, room.topY+room.height-1)
		quad.AddNode(posX, posY, 4)
	}

	// place un portail
	for i := 1; i <= randomIntBetween(1, 8); i++ {
		portal1 := b.createPortal(quad)
		portal2 := b.createPortal(quad)
		portal1.LinkedPortail = &portal2
		portal2.LinkedPortail = &portal1
		listPortal = append(listPortal, portal1)
		listPortal = append(listPortal, portal2)
	}

	// place un Graal
	if randomIntBetween(0, 10) == 0 {
		room := b.root.getRandomRoom()
		posX := randomIntBetween(room.topX+1, room.topX+room.width-1)
		posY := randomIntBetween(room.topY+1, room.topY+room.height-1)
		quad.AddNode(posX, posY, 5)
	}

	return listPortal
}

/*
getRandomRoom retourne une salle aléatoire dans le donjon
*/
func (n *bspNode) getRandomRoom() (room *bspNode) {
	if n.nodeLeft != nil && n.nodeRight != nil {
		if randomIntBetween(0, 1) == 1 {
			return n.nodeLeft.getRandomRoom()
		} else {
			return n.nodeRight.getRandomRoom()
		}
	}
	return n
}

/*
Fonction pour créer un portail (excepté la liaison à un autre portail)
*/
func (b *bspTree) createPortal(quad *quadtree.Quadtree) (p Portal) {
	room := b.root.getRandomRoom()
	p.PosX = randomIntBetween(room.topX+1, room.topX+room.width-1)
	p.PosY = randomIntBetween(room.topY+1, room.topY+room.height-1)
	p.ExitX = p.PosX
	p.ExitY = p.PosY
	for y := p.PosY - 1; y <= p.PosY+1; y++ {
		for x := p.PosX - 1; x <= p.PosX+1; x++ {
			if x != p.PosX && y != p.PosY {
				if quad.GetTileValue(x, y) == 1 {
					p.ExitX = x
					p.ExitY = y
				}
			}
		}
	}
	quad.AddNode(p.PosX, p.PosY, 3)
	return p
}
