package floor

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (f *Floor) SaveMap(filePath string) (err error) {
	// Récupération de la map
	switch configuration.Global.FloorKind {
	case 0:
		err := fmt.Errorf(fmt.Sprintf("Error : floorKind = %d, ce mode ne peut être sauvegardé", configuration.Global.FloorKind))
		check(err)
	case 1:
		if len(f.fullContent) <= 0 {
			log.Panic("fullContent vide!")
		}
	default:
		// les quadtree sont utilisé
		f.fullContent = make([][]int, f.QuadtreeContent.Height)
		for y := 0; y < len(f.fullContent); y++ {
			f.fullContent[y] = make([]int, f.QuadtreeContent.Width)
		}
		f.QuadtreeContent.GetContent(0, 0, f.fullContent) // ne supporte pas les coordonné négative
	}

	// Création du fichier de sauvegarde
	_, errFileExist := os.Stat(filePath) // test l'existence du fichier save
	if errFileExist == nil {
		err := os.Remove(filePath) // si il existe : le supprime pour le recréer (évite bug de réécriture)
		check(err)
	}
	mapFile, err := os.Create(filePath) // Créer le fichier save
	check(err)
	defer mapFile.Close()

	// Écriture du fichier de sauvegarde
	w := bufio.NewWriter(mapFile)
	bytesWritten := 0
	for y := 0; y < len(f.fullContent); y++ {
		//TODO: add modification to convert mapContent value to its equivalent in ASCII
		line := ""
		for x := 0; x < len(f.fullContent[y]); x++ {
			if f.fullContent[y][x] == -1 {
				line += " "
			} else {
				line += fmt.Sprintf("%d", f.fullContent[y][x])
			}
		}
		// ny, err := w.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(f.fullContent[y])), ""), "[]") + "\n")
		ny, err := w.WriteString(line + "\n")
		check(err)
		bytesWritten += ny
	}
	if bytesWritten <= 0 {
		err := fmt.Errorf(fmt.Sprintf("Error : %d bytes écrits", bytesWritten))
		check(err)
	}
	// fmt.Printf("wrote %d bytes\n", bytesWritten)
	err = w.Flush()
	check(err)
	return err
}
