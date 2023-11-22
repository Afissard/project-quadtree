package chunk

import (
	"reflect"
	"testing"
)

func Test_newChunk_n1(t *testing.T) {
	// initialisation
	wantedChunk := [][]int{}
	quadNewChunk := [][]int{}

	for y := 0; y < CHUNK_SIZE; y++ {
		line := []int{}
		for x := 0; x < CHUNK_SIZE; x++ {
			line = append(line, 4)
		}
		wantedChunk = append(wantedChunk, line)
	}

	// fonctions
	newChunk := createChunk(16, 16, 4)
	newChunk.GetContent(0, 0, quadNewChunk)

	// test
	if !reflect.DeepEqual(wantedChunk, newChunk) {
		t.Fatalf("\nerror %d", newChunk)
	}
}
