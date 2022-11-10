package piscine

func Join(strs []string, sep string) string {
	var rep string
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			rep += string(strs[i])
			rep += sep
		} else {
			rep += string(strs[i])
		}
	}
	return rep
}
