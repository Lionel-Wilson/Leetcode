package easy

func MergeAlternately(word1 string, word2 string) string {
	var word1Length int = len(word1)
	var word2Length int = len(word2)
	var stoppingPoint int
	var numberOfLoops int
	var finalString string = ""

	if word1Length > word2Length {
		stoppingPoint = word2Length - 1
		numberOfLoops = word1Length

		for i := 0; i < numberOfLoops; i++ {
			if i > stoppingPoint {
				finalString += string(word1[i:])
				break
			} else {
				finalString += string(word1[i]) + string(word2[i])
			}
		}

	} else if word1Length < word2Length {
		stoppingPoint = word1Length - 1
		numberOfLoops = word2Length

		for i := 0; i < numberOfLoops; i++ {
			if i > stoppingPoint {
				finalString += string(word2[i:])
				break

			} else {
				finalString += string(word1[i]) + string(word2[i])
			}
		}
	} else {

		for i := 0; i < word1Length; i++ {
			finalString += string(word1[i]) + string(word2[i])
		}
	}

	return finalString
}

/* IMPROVED VERSION FROM CHATGPT THAT IS MORE TIME & MEMORY EFFECIENT:
func mergeAlternately(word1 string, word2 string) string {
	var builder strings.Builder
	word1Length := len(word1)
	word2Length := len(word2)
	minLength := word1Length
	if word2Length < minLength {
		minLength = word2Length
	}

	// Append alternate characters from both words
	for i := 0; i < minLength; i++ {
		builder.WriteByte(word1[i])
		builder.WriteByte(word2[i])
	}

	// Append the remaining characters from the longer word
	if word1Length > minLength {
		builder.WriteString(word1[minLength:])
	} else if word2Length > minLength {
		builder.WriteString(word2[minLength:])
	}

	return builder.String()
}


func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}
*/
