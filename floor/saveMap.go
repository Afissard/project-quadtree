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
	// créé et écrit dans mapFile
	mapFile, err := os.Create("../floor-files/" + fileName) //path depuis exécutable
	check(err)
	defer mapFile.Close()

	w := bufio.NewWriter(mapFile)

	bytesWritten := 0
	for y := 0; y < len(f.fullContent); y++ {
		//TODO: add modification to convert mapContent value to its equivalent in ASCII
		ny, err := w.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(f.fullContent[y])), ""), "[]") + "\n")
		check(err)
		bytesWritten += ny
	}
	// fmt.Printf("wrote %d bytes\n", bytesWritten)
	err = w.Flush()
	check(err)
	return err
}
