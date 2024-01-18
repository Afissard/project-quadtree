package quadtree

import (
	"log"
	"reflect"
	"testing"
)

/*
Test la création d'un quadtree à partir de chunk
*/
func Test_newChunk_n1(t *testing.T) {
	log.Println("Test newChunk n°1: Création d'un simple chunk")

	// initialisation
	wantedContent := 1
	chunk := make([][]int, CHUNK_SIZE)
	wantedChunk := make([][]int, CHUNK_SIZE)

	for y := 0; y < len(wantedChunk); y++ {
		chunk[y] = make([]int, CHUNK_SIZE)
		wantedChunk[y] = make([]int, CHUNK_SIZE)
		for x := 0; x < len(wantedChunk[y]); x++ {
			wantedChunk[y][x] = wantedContent
		}
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
		Width:  totalWidth,
		Height: totalHeight,
		root:   &rootNode,
	}

	// fonctions
	quad.CreateSimpleChunk(0, 0, wantedContent)
	

	quad.GetContent(0, 0, chunk)

	// test
	if !reflect.DeepEqual(wantedChunk, chunk) {
		t.Fatalf("\nerror : \n%v \n!= \n%v", chunk, wantedChunk)
	}
}

func Test_newChunk_n2(t *testing.T) {
	log.Println("Test newChunk n°1: Ajout d'un chunk x16 y16")

	// initialisation
	totalWidth, totalHeight := 2, 2
	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1, // contenu à -1 par défaut
	}

	quad := Quadtree{
		Width:  totalWidth,
		Height: totalHeight,
		root:   &rootNode,
	}

	// fonction
	quad.CreateSimpleChunk(CHUNK_SIZE*1, CHUNK_SIZE*1, 2)
}

func Test_newChunk_n3(t *testing.T) {
	log.Println("Test newChunk n°1: Ajout d'un chunk x16 y0")

	// initialisation
	totalWidth, totalHeight := 2, 2
	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1, // contenu à -1 par défaut
	}

	quad := Quadtree{
		Width:  totalWidth,
		Height: totalHeight,
		root:   &rootNode,
	}

	// fonction
	quad.CreateSimpleChunk(CHUNK_SIZE*1, CHUNK_SIZE*0, 3)
	
}

func Test_newChunk_n4(t *testing.T) {
	log.Println("Test newChunk n°1: Ajout d'un chunk x0 y16")

	// initialisation
	totalWidth, totalHeight := 2, 2
	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1, // contenu à -1 par défaut
	}

	quad := Quadtree{
		Width:  totalWidth,
		Height: totalHeight,
		root:   &rootNode,
	}

	// fonction
	quad.CreateSimpleChunk(CHUNK_SIZE*0, CHUNK_SIZE*1, 4)
}

func Test_newChunk_n5(t *testing.T) {
	log.Println("Test newChunk n°1: Ajout d'un chunk x-16 y-16")

	// initialisation
	totalWidth, totalHeight := 2, 2
	rootNode := node{
		topLeftX: 0,
		topLeftY: 0,
		width:    totalWidth,
		height:   totalHeight,
		content:  -1, // contenu à -1 par défaut
	}

	quad := Quadtree{
		Width:  totalWidth,
		Height: totalHeight,
		root:   &rootNode,
	}

	// fonction
	quad.CreateSimpleChunk(CHUNK_SIZE*(-1), CHUNK_SIZE*(-1), 5)
}