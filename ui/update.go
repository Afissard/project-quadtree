package ui

/*
Fonction Update (fichier update.go)

Est appelé dans game.Update
Update général des UI actifs
*/
func (ui *UI) Update() {
	if ui.MainMenu.IsActive {
		ui.updateMainMenu()
	} else if ui.EndScreen.IsActive {
		ui.updateEndScreen()
	} else { // UI en jeu
		if ui.GameOver.IsActive {
			ui.updateGameOver()
		} else if ui.DungeonLevel.IsActive {
			ui.updateDungeonLevel()
		}
	}
}
