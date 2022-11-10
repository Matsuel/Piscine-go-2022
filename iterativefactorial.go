package piscine

func IterativeFactorial(nb int) int {
	rep := 1
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 1 {
		return 1
	} else {
		for i := nb; i > 0; i-- {
			rep *= i
		}
	}
	return rep
}
