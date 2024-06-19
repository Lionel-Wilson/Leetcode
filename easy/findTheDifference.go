package easy

func FindTheDifference(s string, t string) byte {
	// Create an array to count characters (assuming ASCII)
	/*
		In golang an initialized array with no values looks like this:
		[0 0 0 0 0]

		Which means that it can be used as a frequency counter for ASCII chracters as done below.
		remember s[i] returns the byte of a character. In programming languages the byte is the characeters unicode
		So for example, "a"'s unicode number is 97. So s[i] = 97s
		count[s[i]]++ just means the 0 at index 97 in the array will increment by 1. meaning there is 1 "a" in the string.
		And so on.
		count[t[i]]-- decrements each characters counter. meaning the extra added character will cause it's counter to become minus.
		Since there's one more than before.

	*/
	var count [256]int

	// Count each character in s
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	// Count each character in t
	for i := 0; i < len(t); i++ {
		count[t[i]]--
	}

	// The character with a count of -1 is the extra character
	for i := 0; i < len(count); i++ {
		if count[i] == -1 {
			return byte(i)
		}
	}

	// Return 0 if no difference is found (should not happen in valid input)
	return 0
}

/*
MY ORIGINAL SOLUTION(4ms run time). THE ABOVE IS CHATGPT IMPROVEMENT(2ms run time).

	func findTheDifference(s string, t string) byte {
		var sCharacterCountMap = make(map[string]int)
		var tCharacterCountMap = make(map[string]int)

		if len(s) == 0 {
			return t[0]
		}
		for i := 0; i < len(s); i++ {
			_, exists := sCharacterCountMap[string(s[i])]
			var count int = 0

			if !exists {
				for j := 0; j < len(s); j++ {
					if strings.EqualFold(string(s[i]), string(s[j])) {
						count += 1
						sCharacterCountMap[string(s[i])] = count
					}
				}
			}
		}
		for l := 0; l < len(t); l++ {
			_, exists := tCharacterCountMap[string(t[l])]
			var count int = 0

			if !exists {
				for m := 0; m < len(t); m++ {
					if strings.EqualFold(string(t[l]), string(t[m])) {
						count += 1
						tCharacterCountMap[string(t[l])] = count
					}
				}
			}
		}

		for k, v := range tCharacterCountMap {
			characterCount, exists := sCharacterCountMap[k]

			if characterCount != v || !exists {
				return k[0]
			}
		}

		return t[0]
	}
*/
