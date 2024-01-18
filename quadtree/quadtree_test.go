package quadtree

import (
	"log"
	"reflect"
	"testing"
)

const (
	showResult bool = false
)

/*
Test la validité de MakeFromArray et GetContent pour un
quadtree uniforme, celui-ci se fait donc optimisé à 100%
*/
func Test_GetContent_n0(t *testing.T) {
	log.Printf("Test GetContent n°0: Quadtree Uniforme")
	// Initialisation
	mapContent := [][]int{}
	contentHolder := [][]int{}
	maxY, maxX := 17, 17
	for y := 0; y < maxY; y++ {
		lineMap, lineHolder := []int{}, []int{}
		for x := 0; x < maxX; x++ {
			lineMap = append(lineMap, 4)
			lineHolder = append(lineHolder, -1)
		}
		mapContent = append(mapContent, lineMap)
		contentHolder = append(contentHolder, lineHolder)
	}

	// Fonctions
	q := MakeFromArray(mapContent)
	q.GetContent(0, 0, contentHolder)

	// Test
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else if showResult {
		log.Printf("ok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

/*
Test la validité de MakeFromArray et GetContent pour un
petit quadtree de taille 4*4
*/
func Test_GetContent_n1(t *testing.T) {
	log.Printf("Test GetContent n°1: Petit Quadtree Carré")
	// Initialisation
	mapContent := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	contentHolder := [][]int{}
	for y := 0; y < len(mapContent); y++ {
		line := []int{}
		for x := 0; x < len(mapContent[y]); x++ {
			line = append(line, -1)
		}
		contentHolder = append(contentHolder, line)
	}

	// Fonctions
	q := MakeFromArray(mapContent)
	q.GetContent(0, 0, contentHolder)

	// Test
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else if showResult {
		log.Printf("ok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

/*
Test la validité de MakeFromArray et GetContent pour un
grand quadtree de taille 8*8
*/
func Test_GetContent_n2(t *testing.T) {
	log.Printf("Test GetContent n°2: Grand Quadtree Carré")
	// Initialisation
	mapContent := [][]int{
		{1, 1, 3, 4, 1, 1, 3, 4},
		{1, 1, 4, 3, 1, 1, 4, 3},
		{0, 0, 2, 2, 0, 0, 2, 2},
		{0, 0, 2, 2, 0, 0, 2, 2},
		{1, 1, 3, 4, 1, 1, 3, 4},
		{1, 1, 4, 3, 1, 1, 4, 3},
		{0, 0, 2, 2, 0, 0, 2, 2},
		{0, 0, 2, 2, 0, 0, 2, 2},
	}
	contentHolder := [][]int{}
	for y := 0; y < len(mapContent); y++ {
		line := []int{}
		for x := 0; x < len(mapContent[y]); x++ {
			line = append(line, -1)
		}
		contentHolder = append(contentHolder, line)
	}

	// Fonction
	q := MakeFromArray(mapContent)
	q.GetContent(0, 0, contentHolder)

	// Test
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else if showResult {
		log.Printf("ok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

/*
Test la validité de MakeFromArray et GetContent pour
un quadtree rectangulaire
*/
func Test_GetContent_n3(t *testing.T) {
	log.Printf("Test GetContent n°3: Quadtree Rectangulaire")
	// Initialisation
	mapContent := [][]int{
		{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		{1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 2, 2, 1},
		{0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0},
		{1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1},
		{0, 2, 0, 2, 0, 2, 2, 1, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0},
		{1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 0, 2, 0, 1, 2, 2, 2, 1},
		{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
	}
	contentHolder := [][]int{}
	for y := 0; y < len(mapContent); y++ {
		line := []int{}
		for x := 0; x < len(mapContent[y]); x++ {
			line = append(line, -1)
		}
		contentHolder = append(contentHolder, line)
	}

	// Fonctions
	q := MakeFromArray(mapContent)
	q.GetContent(0, 0, contentHolder)

	// Test
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else if showResult {
		log.Printf("ok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

/*
Test la validité de MakeFromArray et GetContent pour un
quadtree exotique
*/
func Test_GetContent_n4(t *testing.T) {
	log.Printf("Test GetContent n°4: Quadtree Exotique")
	// Initialisation
	mapContent := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3, 0},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	contentHolder := [][]int{}
	for y := 0; y < len(mapContent); y++ {
		line := []int{}
		for x := 0; x < len(mapContent[y]); x++ {
			line = append(line, -1)
		}
		contentHolder = append(contentHolder, line)
	}

	// Fonctions
	q := MakeFromArray(mapContent)
	q.GetContent(0, 0, contentHolder)

	// Test
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else if showResult {
		log.Printf("ok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}
