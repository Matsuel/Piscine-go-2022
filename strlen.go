package piscine

func StrLen(s string) int {
	cpt := 0
	l := []rune{}
	for _, i := range s {
		l = append(l, i)
		cpt = len(l)
	}
	return cpt
}
