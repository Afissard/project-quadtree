package quadtree

import "log"

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

/*
Génère et affiche le monde autour du joueur selon l'algorithme de génération sélectionner
Algo:
 1. Détermine les bordure de la camera
    -> si les coordonnées ne sont pas dans le quadtree
    -> détermine les chunk à créer et les génères
*/
func (q *Quadtree) UpdateInfiniteWorld(posX, posY int) {
	var err error
	// Chunk sous le joueur (cas téléportation hors du quadtree)
	topLeftX, topLeftY := getChunkOriginOf(posX, posY)
	if !q.checkIsGenerated(topLeftX, topLeftY) {
		q.CreateSimpleChunk(topLeftX, topLeftY, 2)
		checkErr(err)
	}

	chunkViewField := 1 //paramètre -> variable global ?
	for y := topLeftY - CHUNK_SIZE*chunkViewField; y <= topLeftY+CHUNK_SIZE*chunkViewField; y += CHUNK_SIZE {
		for x := topLeftX - CHUNK_SIZE*chunkViewField; x <= topLeftX+CHUNK_SIZE*chunkViewField; x += CHUNK_SIZE {
			if !q.checkIsGenerated(x, y) {
				content := 4
				q.CreateSimpleChunk(x, y, content)
				checkErr(err)
			}
		}
	}
}

/*
Vérifie si les coordonnées données en entrée existe dans le quadtree (dedans ou dehors)
*/
func (q *Quadtree) checkIsIn(posX, posY int) bool {
	if posX < q.root.topLeftX ||
		posY < q.root.topLeftY ||
		posX >= q.root.topLeftX+q.Width ||
		posY >= q.root.topLeftY+q.Height {
		return false
	}
	return true
}

/*
Vérifie si les coordonnées données correspondent à une tuile généré ou non
*/
func (q *Quadtree) checkIsGenerated(posX, posY int) bool {
	if q.checkIsIn(posX, posY) {
		if q.Seek(posX, posY).content != -1 {
			return true
		}
	}
	return false
}

/*
retourne les coordonnée de l'origine du chunk (en haut à gauche) par rapport à la position donnée
*/
func getChunkOriginOf(posX, posY int) (topLeftX, topLeftY int) {
	topLeftX = int(posX/CHUNK_SIZE) * CHUNK_SIZE
	topLeftY = int(posY/CHUNK_SIZE) * CHUNK_SIZE
	return topLeftX, topLeftY
}
