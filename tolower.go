package piscine

func ToLower(s string) string {
	var rep string
	for _, ch := range s {
		if IsUpper(string(ch)) {
			rep += string(rune(ch + 32))
		} else {
			rep += string(ch)
		}
	}
	return rep
}
