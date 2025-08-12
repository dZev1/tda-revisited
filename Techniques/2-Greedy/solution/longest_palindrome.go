package solution

/*
	Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

	Letters are case sensitive, for example, "Aa" is not considered a palindrome.
*/

func LongestPalindrome(s string) int {
	m := make(map[rune]int)
	res := 0

	for _, l := range s {
		m[l]++
		if m[l]%2 == 0 {
			res += 2
		}
	}

	oddAdded := false

	for _, value := range m {
		if value % 2 == 1 && !oddAdded {
			res++
			oddAdded = true
		}
	}

	return res
}
