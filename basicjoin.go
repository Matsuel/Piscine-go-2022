package piscine

func BasicJoin(elems []string) string {
	var rep string
	for _, el := range elems {
		rep += string(el)
	}
	return rep
}
