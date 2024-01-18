package assets

import (
	"bytes"
	"image"
	"log"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
	"golang.org/x/image/font"
)

//go:embed tileSheet/floor.png
var floorBytes []byte

//go:embed tileSheet/dungeonFloor.png
var dungeonFloorBytes []byte

// FloorImage contient une version compatible avec Ebitengine de l'image
// qui contient les différents éléments qui peuvent s'afficher au sol
// (herbe, sable, etc).
// Dans la version du projet qui vous est fournie, ces éléments sont des
// carrés de 16 pixels de côté. Vous pourrez changer cela si vous le voulez.
var FloorImage *ebiten.Image

//go:embed spriteSheet/character.png
var characterBytes []byte

// CharacterImage contient une version compatible avec Ebitengine de
// l'image qui contient les différentes étapes de l'animation du
// personnage.
// Dans la version du projet qui vous est fournie, ce personnage tient
// dans un carré de 16 pixels de côté. Vous pourrez changer cela si vous
// le voulez.
var CharacterImage *ebiten.Image

var (
	//go:embed fonts/BigBlueTerm437NerdFont-Regular.ttf
	BigBlueTerm []byte
	// Police de caractère : BigBlueTerm NerdFont
	FontBigBlueTerm font.Face
)

// Load est la fonction en charge de transformer, à l'exécution du programme,
// les images du jeu en structures de données compatibles avec Ebitengine.
// Ces structures de données sont stockées dans les variables définies ci-dessus.
func Load() {
	if configuration.Global.FloorKind == 3 {
		FloorImage = imgLoad(dungeonFloorBytes)
	} else {
		FloorImage = imgLoad(floorBytes)
	}

	CharacterImage = imgLoad(characterBytes)

	FontBigBlueTerm = fontLoad(BigBlueTerm, 12, 72)
}

// Fonction pour convertir les bytes d'une image au format d'image supporté par ebiten.
func imgLoad(imgBytes []byte) (img *ebiten.Image) {
	decoded, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(decoded)
}
