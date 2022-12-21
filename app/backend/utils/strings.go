package utils

// Get the substring skipping the second parameter
func SubstrSkiping(str, letter string) string {
	var init int
	var firstLetter string
	if len(letter) >= 1 {
		firstLetter = letter[0:1]
	} else {
		firstLetter = " "
	}
	for i, char := range str {
		if string(char) != firstLetter {
			init = i
			break
		}
	}
	return str[init:]
}

func FirstIndexSkipping(str, letter string) string {
	substr := SubstrSkiping(str, letter)
	if len(substr) >= 1 {
		return substr[0:1]
	}
	return ""
}

func SubstrWithEnd(str, letter string) string {
	var init int
	for i, v := range str {
		if i == 0 {
			continue
		}
		if string(v) == letter {
			init = i
			break
		}
	}
	return str[0 : init+1]
}
