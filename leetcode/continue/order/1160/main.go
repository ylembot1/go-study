package main

func countCharacters(words []string, chars string) int {
	res := 0

	for _, word := range words {
		charsMap := make(map[rune]int)
		for _, char := range chars {
			charsMap[char]++
		}

		flag := true
		for _, char := range word {
			charsMap[char]--

			if charsMap[char] < 0 {
				flag = false
				break
			}
		}

		if flag {
			res += len(word)
		}
	}
	return res
}
