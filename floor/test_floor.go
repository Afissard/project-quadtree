package floor

import (
	"fmt"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func Test_readFloorFromFile(){
	floorContent := readFloorFromFile("../floor-file/exemple")
	fmt.Println("floorContent", floorContent)
}