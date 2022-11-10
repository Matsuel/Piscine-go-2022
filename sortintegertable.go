package piscine

func SortIntegerTable(table []int) {
	for i := 1; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			tmp_i := table[i]
			if table[i] < table[j] {
				table[i] = table[j]
				table[j] = tmp_i
			}
		}
	}
}
