package main

import "fmt"

func main() {
	printLettersWord("Example")
}

func printLettersWord(word string) {
	lettersQauntity := countLettersWord(word)
	fmt.Printf("the word %s has %d letters\n", word, lettersQauntity)

	for i := 0; i < lettersQauntity; i++ {
		fmt.Println(string(word[i]))
	}

}

func countLettersWord(word string) int {
	quantity := 0

	for range word {
		quantity++
	}

	return quantity
}
