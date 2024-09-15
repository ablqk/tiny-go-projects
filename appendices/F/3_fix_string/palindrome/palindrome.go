package palindrome

import (
	"fmt"
	"strconv"
)

// IsPalindromeNumber returns if an integer is readable in both ways.
// Only unsigned int are supported.
// "1221" is a palindrome
// "01" is not a palindrome
// "-1" is not a palindrome
func IsPalindromeNumber(s string) (bool, error) {
	// check if it's a number, e.g.: "ң\xa3" is not a number
	_, err := strconv.Atoi(s)
	if err != nil {
		// V2 returns an error
		return false, fmt.Errorf("not a number: %s", s)
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false, nil
		}
	}

	return true, nil
}
