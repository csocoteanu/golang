package main 

import (
	"fmt"
)

type TernarySearchTree struct {
	Value uint8
	IsEndOfString bool
	Left  *TernarySearchTree
	Right *TernarySearchTree
	Down  *TernarySearchTree
}

func insert(t *TernarySearchTree, inputString string, index int) *TernarySearchTree {

	fmt.Printf("--- enter insert: %c - %d\n", inputString[index], index)

	if (t == nil) {

		fmt.Printf("--- creating new node: %c - %d\n", inputString[index], index)

		t = &TernarySearchTree{inputString[index], false, nil, nil, nil}
	}

	if (inputString[index] < t.Value) {

		fmt.Printf("--- lower then: %c - %d < %c\n", inputString[index], index, t.Value)

		t.Left = insert(t.Left, inputString, index)
	}
	if (inputString[index] > t.Value) {

		fmt.Printf("--- greater then: %c - %d > %c\n", inputString[index], index, t.Value)

		t.Right = insert(t.Right, inputString, index)
	} else {

		if (index + 1 < len(inputString)) {
			fmt.Printf("--- equal: %c - %d == %c\n", inputString[index], index, t.Value)
			t.Down = insert(t.Down, inputString, index + 1)
		} else {
			fmt.Println("====== Reaching end of string")
			t.IsEndOfString = true	
		}
	}

	return t
}

func Print(t *TernarySearchTree, storedWord []uint8, index uint8) {

	if (t == nil) {
		return
	}

	Print(t.Left, storedWord, index)

	storedWord[index] = t.Value

	if (t.IsEndOfString) {
		storedWord[index + 1] = 0
		fmt.Println(string(storedWord))
	}

	Print(t.Down, storedWord, index + 1)

	Print(t.Right, storedWord, index)	
}

func HasWord(t *TernarySearchTree, searchString string, index int) bool {

	if (t == nil) {
		return false
	}

	if (searchString[index] < t.Value) {
		return HasWord(t.Left, searchString, index)
	}
	if (searchString[index] > t.Value) {
		return HasWord(t.Right, searchString, index)
	} else {

		if (index + 1 < len(searchString)) {
			return HasWord(t.Down, searchString, index + 1)
		}

		return t.IsEndOfString
	}
}

func PrintSimilarWords(t *TernarySearchTree, searchPattern string, index int, matchedWord []uint8, matchedWordIndex int) {

	if (t == nil) {
		return
	}

	if (searchPattern[index] < t.Value) {
		PrintSimilarWords(t.Left, searchPattern, index, matchedWord, matchedWordIndex)
	}
	if (searchPattern[index] > t.Value) {
		PrintSimilarWords(t.Right, searchPattern, index, matchedWord, matchedWordIndex)
	} else {
		matchedWord[matchedWordIndex] = searchPattern[index]

		if (index + 1 < len(searchPattern)) {
			PrintSimilarWords(t.Down, searchPattern, index + 1, matchedWord, matchedWordIndex + 1)
		} else {
			fmt.Printf("All Words ar starting with: %s",string(matchedWord))
			// TODO:  BFS here
		}
	}
}

func main() {

	words := []string{"ana", "analog", "anapoda", "bogdan"}
	var tree *TernarySearchTree

	for _, word := range words {
		tree = insert(tree, word, 0)
		fmt.Println()
		fmt.Println()
	}

	fmt.Println("---- Printing tree ------>")
	storedWord := make([]uint8, 100)
	Print(tree, storedWord, 0)

	for _, word := range words {
		result := HasWord(tree, word, 0)
		fmt.Printf("%s is present in the dictionary: %d\n", word, result)
	}

	missingWords := []string{"anapoga", "bog og"}
	for _, word := range missingWords {
		result := HasWord(tree, word, 0)
		fmt.Printf("%s is present in the dictionary: %d\n", word, result)
	}

	possibleWords := make([]uint8, 100)
	fmt.Println("----------------------------------------------------------")
	PrintSimilarWords(tree, "ana", 0, possibleWords, 0)
}
