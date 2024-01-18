package floor

import (
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
)

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {

	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	switch configuration.Global.FloorKind {
	case 3:
		for i := 0; i < len(fk3BlockList); i++ {
			blocking[0] = blocking[0] || relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == fk3BlockList[i]
			blocking[1] = blocking[1] || relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == fk3BlockList[i]
			blocking[2] = blocking[2] || relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == fk3BlockList[i]
			blocking[3] = blocking[3] || relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == fk3BlockList[i]
		}
	default:
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1
	}
	return blocking
}
