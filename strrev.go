package piscine

func StrRev(s string) string {
	var word string
	word_list := []string{}
	for _, character := range s {
		word_list = append(word_list, string(character))
	}
	for i := len(word_list) - 1; len(word_list) != 0; i-- {
		word += word_list[i]
		word_list = word_list[:len(word_list)-1]
	}
	return word
}

// Stirng c'est comme un tebaleau donc string[-1]= tableau[-1]
