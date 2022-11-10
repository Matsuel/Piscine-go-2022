package piscine

func IterativePower(nb int, power int) int {
	rep := nb
	if power < 0 {
		return 0
	} else if power == 0 && nb == 0 {
		return 1
	}
	for i := power; i > 1; i-- {
		rep *= nb
	}
	return rep
}
