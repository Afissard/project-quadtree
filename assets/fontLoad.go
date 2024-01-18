package assets

import (
	_ "embed"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// fontLoad sert à charger la police d'écriture utilisé par le jeu (excepté celle de debug qui est déjà intégré)
func fontLoad(fontByte []byte, size, dpi float64) (theFont font.Face) {
	var err error
	// f, err := opentype.Parse(gomono.TTF)
	f, err := opentype.Parse(fontByte)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	theFont, err = opentype.NewFace(f, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return theFont
}
