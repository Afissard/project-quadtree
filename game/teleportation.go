package game

/*
teleport player téléporte le joueur à la la tuile de sortie du portail sur lequel il se trouve 
*/
func (g *Game) teleportPlayer() {
	for i := 0; i < len(g.listPortal); i++ {
		currentPortal := g.listPortal[i]
		if g.character.X == currentPortal.PosX && g.character.Y == currentPortal.PosY {
			g.character.X = currentPortal.LinkedPortail.ExitX
			g.character.Y = currentPortal.LinkedPortail.ExitY
			g.camera.X = g.character.X
			g.camera.Y = g.character.Y
		}
	}
}
