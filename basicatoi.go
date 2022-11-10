package piscine

func BasicAtoi(s string) int {
	rep := 0
	n := 1
	atoi := []rune{}
	for _, character := range s {
		atoi = append(atoi, rune(character))
	}
	for i := len(atoi) - 1; len(atoi) != 0; i-- {
		rep += int(atoi[i]-48) * n
		atoi = atoi[:len(atoi)-1]
		n *= 10
	}
	return rep
}
