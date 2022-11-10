package piscine

func UltimateDivMod(a *int, b *int) {
	tmp_a := *a
	tmp_b := *b
	*a = tmp_a / tmp_b
	*b = tmp_a % tmp_b
}
