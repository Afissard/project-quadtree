package quadtree

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_newChunk_n1(t *testing.T) {
	fmt.Println("\nTest newChunk n°1:")

	// initialisation
	wantedContent := 2
	chunk := make([][]int, CHUNK_SIZE)
	wantedChunk := make([][]int, CHUNK_SIZE)

	for y := 0; y < len(wantedChunk); y++ {
		chunk[y] = make([]int, CHUNK_SIZE)
		wantedChunk[y] = make([]int, CHUNK_SIZE)
	}

	totalWidth, totalHeight := 64, 64
	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1, // contenu à -1 par défaut
	}

	quad := Quadtree{
		width:  totalWidth,
		height: totalHeight,
		root:   &rootNode,
	}

	for y := 0; y < len(wantedChunk); y++ {
		for x := 0; x < len(wantedChunk[y]); x++ {
			wantedChunk[y][x] = wantedContent
		}
	}

	// fonctions
	quad.CreateChunk(0, 0)
	quad.GetContent(0, 0, chunk)

	//fmt.Println("Created chunk :")
	for y := 0; y < len(chunk); y++ {
		//fmt.Println(y, chunk[y])
	}
	//fmt.Println("Wanted chunk :")
	for y := 0; y < len(wantedChunk); y++ {
		//fmt.Println(y, wantedChunk[y])
	}

	// test
	if !reflect.DeepEqual(wantedChunk, chunk) {
		t.Fatalf("\nerror : \n%v \n!= \n%v", chunk, wantedChunk)
	}
}
