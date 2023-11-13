package quadtree

import (
	"fmt"
	"reflect"
	"testing"
)

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
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("\nok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
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
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("\nok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}

func Test_GetContent_n1_optimized(t *testing.T) {
	/*
		Test la validité de MakeFromArray et GetContent
	*/
	fmt.Println("\nTest GetContent (n°1 optimized):")
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
	q.root.content = q.root.optimize()
	fmt.Println("layer 1", q.root.topLeftNode, q.root.topRightNode, "\n", q.root.bottomLeftNode, q.root.bottomRightNode)
	fmt.Println("layer 2", q.root.topLeftNode.topLeftNode)
	q.GetContent(0, 0, contentHolder)
	if !reflect.DeepEqual(mapContent, contentHolder) {
		t.Fatalf("\nmapContent : %d,\ndifférent de contentHolder : %d", mapContent, contentHolder)
	} else {
		fmt.Printf("\nok : mapContent : %d == contentHolder : %d\n", mapContent, contentHolder)
	}
}
