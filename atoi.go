package piscine

func Atoi(s string) int {
	atoi := 0
	n := 1
	var isnegative bool
	if s == "" {
		return 0
	}
	if s[0] == '-' {
		isnegative = true
	} else {
		isnegative = false
	}
	slice := []rune{}
	for _, ch := range s {
		if ch != '0' && ch != '1' && ch != '2' && ch != '3' && ch != '4' && ch != '5' && ch != '6' && ch != '7' && ch != '8' && ch != '9' && ch != '+' && ch != '-' {
			return 0
		} else {
			slice = append(slice, rune(ch))
		}
	}
	if slice[0] == '-' || slice[0] == '+' {
		slice = slice[1:]
	}
	for i := len(slice) - 1; len(slice) != 0; i-- {
		if slice[i] == '-' || slice[i] == '+' {
			return 0
		}
		atoi += int(slice[i]-48) * n
		slice = slice[:len(slice)-1]
		n *= 10
	}
	if isnegative == true {
		atoi *= -1
	}
	return atoi
}

// package piscine

// func Atoi(s string) int {
// 	rep := 0
// 	coef_mul := 1
// 	n := 1
// 	atoi := []rune{}
// 	if s == "" {
// 		return 0
// 	}
// 	for _, character := range s {
// 		if len(s) == 1 {
// 			if character == '-' || character == '+' {
// 				return 0
// 			}
// 		} else if len(s) == 0 {
// 			return 0
// 		} else if character == '-' && s[1] != '-' {
// 			coef_mul = -1
// 		} else if character == '+' && s[1] != '+' {
// 			coef_mul = 1
// 		} else if character == '+' && s[1] == '+' || character == '-' && s[1] == '-' {
// 			return 0
// 		} else if character != '0' && character != '1' && character != '2' && character != '3' && character != '4' && character != '5' && character != '6' && character != '7' && character != '8' && character != '9' {
// 			return 0
// 		} else {
// 			atoi = append(atoi, rune(character))
// 		}
// 	}
// 	for i := len(atoi) - 1; len(atoi) != 0; i-- {
// 		rep += int(atoi[i]-48) * n
// 		atoi = atoi[:len(atoi)-1]
// 		n *= 10
// 	}
// 	return (rep * coef_mul)
// }
