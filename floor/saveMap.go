package floor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (f *Floor) SaveMap(fileName string) (err error) {
	mapFile := &os.File{}
	_, errFileExist := os.Stat("../floor-files/" + fileName) // test l'existence du fichier save
	if errFileExist == nil {
		mapFile, err := os.Open("../floor-files/" + fileName) // Ouvre le fichier save pour le réécrire
		check(err)
		defer mapFile.Close()
	} else {
		mapFile, err := os.Create("../floor-files/" + fileName) // Créer le fichier save
		check(err)
		defer mapFile.Close()
	}

	w := bufio.NewWriter(mapFile)

	bytesWritten := 0
	for y := 0; y < len(f.fullContent); y++ {
		//TODO: add modification to convert mapContent value to its equivalent in ASCII
		ny, err := w.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(f.fullContent[y])), ""), "[]") + "\n")
		check(err)
		bytesWritten += ny
	}
	if bytesWritten <= 0 {
		err := fmt.Errorf(fmt.Sprintf("byte écrits : %d", bytesWritten))
		check(err)
	}
	fmt.Printf("wrote %d bytes\n", bytesWritten)
	err = w.Flush()
	check(err)
	return err
}
