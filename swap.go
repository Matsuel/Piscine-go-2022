package piscine

func Swap(a *int, b *int) {
	tmp_a := *a
	tmp_b := *b
	*a = tmp_b
	*b = tmp_a
}
