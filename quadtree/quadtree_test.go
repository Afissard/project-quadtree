package quadtree

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_MakeFromArray(t *testing.T) {
	/*
		TODO: écrire un vrai test
	*/
	mapContent := [][]int{
		{1, 1, 3, 4},
		{1, 1, 4, 3},
		{0, 0, 2, 2},
		{0, 0, 2, 2},
	}
	q := MakeFromArray(mapContent)
	if q.root == nil {
		t.Fatal("floorContent est vide")
	}
	fmt.Println("quadtree : ", q, "\nnode root :", q.root)

	// test "compatibilité avec get"
	contentHolder := [][]int{
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
	}
	q.GetContent(0, 0, contentHolder)
	if reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("mapContent : %d, différent de contentHolder : %d", mapContent, contentHolder)
	}
}
