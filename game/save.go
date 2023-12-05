package game

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Sauvegarde l'état actuel du jeux : map, joueur et autre (work in progress)
func (g *Game) SaveGame(fileName string) {
	// appel de saveMap
	err := g.floor.SaveMap(fileName)
	check(err)

	// TODO: sauvegarde du joueur (position et autres) dans un json
}

// TODO: LoadFromSave
