package issubsequence

// https://leetcode.com/problems/is-subsequence/

func IsSubsequence(s string, t string) bool {
	currentNumberOfChars := 0
	for _, char := range t {
		if len(s) == currentNumberOfChars {
			return true
		}
		if string(s[currentNumberOfChars]) == string(char) {
			currentNumberOfChars++
		}
	}
	return len(s) == currentNumberOfChars
}

// follow up

func IsSubsequenceFollowUp(ss []string, t string) bool {
	for _, s := range ss {
		if IsSubsequence(s, t) {
			return true
		}
	}
	return false
}
