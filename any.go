package piscine

func Any(f func(string) bool, a []string) bool {
	rep := false
	for i := 0; i < len(a); i++ {
		if rep == true {
			return rep
		} else {
			rep = f(a[i])
		}
	}
	return rep
}
