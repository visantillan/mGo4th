package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("ab", "cdefg"))
}

func mergeAlternately2(word1, word2 string) string {

	var altWord string

	if len(word1) == 0 {
		return "word1 should be at least 1 character"
	}

	sWord := getShortWord(word1, word2)

	for i := 0; i <= len(sWord)-1; i++ {
		altWord += string(word1[i]) + string(word2[i])
	}

	return altWord + getBigWord(word1, word2)[len(sWord):]
}

func mergeAlternately(word1, word2 string) string {

	var altWord string

	if len(word1) == 0 {
		return "word1 should be at least 1 character"
	}

	sWord := getShortWord(word1, word2)

	for i := 0; i <= len(sWord)-1; i++ {
		altWord += string(word1[i]) + string(word2[i])
	}

	return altWord + getBigWord(word1, word2)[len(sWord):]
}

func getShortWord(word1, word2 string) string {
	if len(word1) > len(word2) {
		return word2
	}

	return word1
}

func getBigWord(word1, word2 string) string {
	if len(word1) > len(word2) {
		return word1
	}

	return word2
}
