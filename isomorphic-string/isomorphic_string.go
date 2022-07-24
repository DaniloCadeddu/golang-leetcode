package isomorphicstring

// https://leetcode.com/problems/isomorphic-strings/

func IsIsomorphic(s string, t string) bool {

	var hashMap = make(map[rune]rune)
	var newS string
	var currentChar rune = 1

	for _, char := range s {
		if hashMap[char] == 0 {
			hashMap[char] = currentChar
			currentChar++
		}
		newS += string(hashMap[char])
	}

	hashMap = make(map[rune]rune)
	var newT string
	currentChar = 1

	for _, char := range t {
		if hashMap[char] == 0 {
			hashMap[char] = currentChar
			currentChar++
		}
		newT += string(hashMap[char])
	}

	return newS == newT
}
