package quadtree

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_newChunk_n1(t *testing.T) {
	fmt.Println("\nTest newChunk n°1:")

	// initialisation
	wantedChunk := [][]int{}
	chunk := [][]int{}

	totalWidth, totalHeight := 64, 64
	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1, // contenu à -1 par défaut
	}

	newChunk := Quadtree{
		width:  totalWidth,
		height: totalHeight,
		root:   &rootNode,
	}

	for y := 0; y < CHUNK_SIZE; y++ {
		line := []int{}
		for x := 0; x < CHUNK_SIZE; x++ {
			line = append(line, 4)
		}
		wantedChunk = append(wantedChunk, line)
	}

	// fonctions
	newChunk.createChunk(32, 16, 2)
	newChunk.GetContent(32, 16, chunk)

	for y := 0; y < len(chunk); y++ {
		fmt.Println(chunk[y])
	}
	fmt.Println("here")

	// test
	if !reflect.DeepEqual(wantedChunk, newChunk) {
		t.Fatalf("\nerror %v", newChunk)
	}
}
