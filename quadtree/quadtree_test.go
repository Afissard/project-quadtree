package quadtree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
func Test_MakeFromArray(t *testing.T) {
	// n'est pas un véritable test
	// -> se fait testé avec GuetContent
	fmt.Println("Test MakeFromArray : ")
	mapContent := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	q := MakeFromArray(mapContent)
	if q.root == nil {
		t.Fatal("floorContent est vide")
	} else {
		fmt.Println("ok")

		// fmt.Println("q :", q)
		// fmt.Println("root :", q.root)

		// fmt.Println("layer 1, tl :", q.root.topLeftNode)
		// fmt.Println("layer 1, tr :", q.root.topRightNode)
		// fmt.Println("layer 1, bl :", q.root.bottomLeftNode)
		// fmt.Println("layer 1, br :", q.root.bottomRightNode)
	}
}
*/

func Test_GetContent_n1(t *testing.T) {
	/*
		Test la validité de MakeFromArray et GetContent
	*/
	fmt.Println("\nTest GetContent (n°1):")
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
		t.Fatalf("mapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("ok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

func Test_GetContent_n2(t *testing.T) {
	/*
		Test la validité de MakeFromArray et GetContent
	*/
	fmt.Println("\nTest GetContent (n°2):")
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
		t.Fatalf("mapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("ok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}
