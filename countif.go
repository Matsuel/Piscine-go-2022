package piscine

func CountIf(f func(string) bool, tab []string) int {
	rep := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) == true {
			rep += 1
		}
	}
	return rep
}
