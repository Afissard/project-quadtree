package floor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (f *Floor) SaveMap(fileName string) (err error) {
	// Récupération de la map actuelle
	// fmt.Println(f.quadtreeContent.root, f.quadtreeContent.width, f.quadtreeContent.height) // doesn't work
	f.quadtreeContent.GetContent(0, 0, f.fullContent) // BUG : vide
	fmt.Println("fullContent : ", f.fullContent)

	// Création du fichier de sauvegarde
	_, errFileExist := os.Stat("../floor-files/" + fileName) // test l'existence du fichier save
	if errFileExist == nil {
		err := os.Remove("../floor-files/" + fileName) // si il existe : le supprime pour le recréer (évite bug de réécriture)
		check(err)
	}
	mapFile, err := os.Create("../floor-files/" + fileName) // Créer le fichier save
	check(err)
	defer mapFile.Close()

	// Écriture du fichier de sauvegarde
	w := bufio.NewWriter(mapFile)
	bytesWritten := 0
	for y := 0; y < len(f.fullContent); y++ {
		//TODO: add modification to convert mapContent value to its equivalent in ASCII
		ny, err := w.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(f.fullContent[y])), ""), "[]") + "\n")
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
