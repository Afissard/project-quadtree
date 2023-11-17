package quadtree

import (
	"fmt"
	"reflect"
	"testing"
)

// Test la validité de MakeFromArray et GetContent
func Test_GetContent_n0(t *testing.T) {
	fmt.Println("\nTest GetContent n°0:")

	// initialisation :
	mapContent := [][]int{}
	contentHolder := [][]int{}
	maxY, maxX := 64, 64
	for y := 0; y <= maxY; y++ {
		lineMap, lineHolder := []int{}, []int{}
		for x := 0; x <= maxX; x++ {
			lineMap = append(lineMap, 4)
			lineHolder = append(lineHolder, -1)
		}
		mapContent = append(mapContent, lineMap)
		contentHolder = append(contentHolder, lineHolder)
	}

	// fonctions :
	q := MakeFromArray(mapContent)
	q.root.content = optimize(q.root)
	q.GetContent(0, 0, contentHolder)

	// tests :
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("\nok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

func Test_GetContent_n1(t *testing.T) {
	/*
		Test la validité de MakeFromArray et GetContent
	*/
	fmt.Println("\nTest GetContent n°1:")
	mapContent := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	contentHolder := [][]int{
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
	}
	q := MakeFromArray(mapContent)
	q.GetContent(0, 0, contentHolder)
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("\nok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

func Test_GetContent_n2(t *testing.T) {
	/*
		Test la validité de MakeFromArray et GetContent
	*/
	fmt.Println("\nTest GetContent n°2:")
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
	contentHolder := [][]int{
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
	}
	q := MakeFromArray(mapContent)
	q.GetContent(0, 0, contentHolder)
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("\nok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

func Test_GetContent_n1_optimized(t *testing.T) {
	/*
		Test la validité de MakeFromArray et GetContent
	*/
	fmt.Println("\nTest GetContent n°1 optimized:")
	mapContent := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	contentHolder := [][]int{}
	for y := 0; y < len(mapContent); y++ {
		line := []int{}
		for x := 0; x < len(mapContent[0]); x++ {
			line = append(line, -1)
		}
		contentHolder = append(contentHolder, line)
	}

	q := MakeFromArray(mapContent)
	q.root.content = optimize(q.root)
	fmt.Println("layer 1", q.root.topLeftNode, q.root.topRightNode, "\n", q.root.bottomLeftNode, q.root.bottomRightNode)
	fmt.Println("layer 2", q.root.topLeftNode.topLeftNode)
	q.GetContent(0, 0, contentHolder)
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("\nok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

func Test_GetContent_n2_optimized(t *testing.T) {
	/*
		Test la validité de MakeFromArray et GetContent
	*/
	fmt.Println("\nTest GetContent n°2 optimized:")
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
		for x := 0; x < len(mapContent[0]); x++ {
			line = append(line, -1)
		}
		contentHolder = append(contentHolder, line)
	}

	q := MakeFromArray(mapContent)
	q.root.content = optimize(q.root)
	q.GetContent(0, 0, contentHolder)
	fmt.Println("root", q.root)
	fmt.Println("layer 1", q.root.topLeftNode, q.root.topRightNode, q.root.bottomLeftNode, q.root.bottomRightNode)
	fmt.Println("layer 2", q.root.topLeftNode.topLeftNode)
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("\nok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

func Test_GetContent_n3_optimized(t *testing.T) {
	/*
		Test la validité de MakeFromArray et GetContent
	*/
	fmt.Println("\nTest GetContent n°3 optimized:")
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
		for x := 0; x < len(mapContent[0]); x++ {
			line = append(line, -1)
		}
		contentHolder = append(contentHolder, line)
	}

	// redimensionnement de floorContent
	totalWidth := len(mapContent[0])
	totalHeight := len(mapContent)
	if totalWidth > totalHeight {
		for y := totalHeight; y < totalWidth; y++ {
			line := []int{}
			for x := 0; x < totalWidth; x++ {
				line = append(line, -1)
			}
			mapContent = append(mapContent, line)
		}
		fmt.Printf("ajout de %d ligne(s)\n", totalWidth-totalHeight)
		totalHeight = totalWidth
	} else if totalWidth < totalHeight {
		for y := 0; y < totalHeight; y++ {
			for x := totalWidth; x < totalHeight; x++ {
				mapContent[y] = append(mapContent[y], -1)
			}
		}
		totalWidth = totalHeight
		fmt.Printf("ajout de %d colone(s)\n", totalHeight-totalWidth)
	}

	// Fonctions
	q := MakeFromArray(mapContent)
	q.root.content = optimize(q.root)
	q.GetContent(0, 0, contentHolder)

	// Test réussite
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("\nok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}
