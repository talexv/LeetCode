func isPalindrome(s string) bool {
    var chars []rune

    for _, char := range s {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            chars = append(chars, unicode.ToLower(char))
        }
    }
    left, right := 0, len(chars) - 1
    for left < right {
        if chars[left] != chars[right] {
            return false
        }
        left++
        right--
    }
    return true
}