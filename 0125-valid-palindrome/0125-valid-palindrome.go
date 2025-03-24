package validpalindrome

import "unicode"

// Compute: O(n)
// Memory: O(n)

// func isPalindrome(s string) bool {
// 	var chars []rune
// 	for _, char := range s {
// 		if unicode.IsLetter(char) || unicode.IsDigit(char) {
// 			chars = append(chars, unicode.ToLower(char))
// 		}
// 	}
// 	left, right := 0, len(chars)-1
// 	for left < right {
// 		if chars[left] != chars[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }

// Compute: O(n)
// Memory: O(1)

//nolint:unused // solution LeetCode problem
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
			continue
		}
		if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
			continue
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}
