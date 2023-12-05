package floor

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_save_n1(t *testing.T) {
	filepath := "../floor-files/"
	mapFile := filepath + "exemple"
	saveFile := filepath + "autoSave"
	f := Floor{
		fullContent: readFloorFromFile(mapFile),
	}
	f.SaveMap("autoSave")
	savedContent := readFloorFromFile(saveFile)

	// Test
	if !reflect.DeepEqual(f.fullContent, savedContent) {
		t.Fatalf("\nBad save : %d != %d", f.fullContent, savedContent)
	} else {
		fmt.Printf("\nOK")
	}
}
