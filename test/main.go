// package main

// import "piscine"

// func main() {
// 	piscine.IsNegative(1)
// 	piscine.IsNegative(0)
// 	piscine.IsNegative(-1)
// }

// func main() {
// 	piscine.PrintComb()
// }

// func main() {
// 	piscine.PrintComb2()
// }

// package main

// import (
// 	"piscine"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	piscine.PrintNbr(-123)
// 	piscine.PrintNbr(0)
// 	piscine.PrintNbr(123)
// 	z01.PrintRune('\n')
// }

// package main

// import "piscine"

// func main() {
// 	piscine.PrintCombN(1)
// 	piscine.PrintCombN(3)
// 	piscine.PrintCombN(9)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	n := 0
// 	piscine.PointOne(&n)
// 	fmt.Println(n)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 0
// 	b := &a
// 	n := &b
// 	piscine.UltimatePointOne(&n)
// 	fmt.Println(a)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 13
// 	b := 2
// 	var div int
// 	var mod int
// 	piscine.DivMod(a, b, &div, &mod)
// 	fmt.Println(div)
// 	fmt.Println(mod)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 13
// 	b := 2
// 	piscine.UltimateDivMod(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// package main

// import "piscine"

// func main() {
// 	piscine.PrintStr("Hello World!")
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	l := piscine.StrLen("Hello World!")
// 	fmt.Println(l)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 0
// 	b := 1
// 	piscine.Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := "Hello World!"
// 	s = piscine.StrRev(s)
// 	fmt.Println(s)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.BasicAtoi("12345"))
// 	fmt.Println(piscine.BasicAtoi("0000000012345"))
// 	fmt.Println(piscine.BasicAtoi("000000"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.BasicAtoi2("12345"))
// 	fmt.Println(piscine.BasicAtoi2("0000000012345"))
// 	fmt.Println(piscine.BasicAtoi2("012 345"))
// 	fmt.Println(piscine.BasicAtoi2("Hello World!"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Atoi("12345"))
// 	fmt.Println(piscine.Atoi("0000000012345"))
// 	fmt.Println(piscine.Atoi("012 345"))
// 	fmt.Println(piscine.Atoi("Hello World!"))
// 	fmt.Println(piscine.Atoi("+1234"))
// 	fmt.Println(piscine.Atoi("-1234"))
// 	fmt.Println(piscine.Atoi("++1234"))
// 	fmt.Println(piscine.Atoi("--1234"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := []int{5, 4, 3, 2, 1, 0}
// 	piscine.SortIntegerTable(s)
// 	fmt.Println(s)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	arg := 20
// 	fmt.Println(piscine.IterativeFactorial(arg))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	arg := 20
// 	fmt.Println(piscine.RecursiveFactorial(arg))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.IterativePower(4, 3))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.RecursivePower(-4, 3))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	arg1 := 4
// 	fmt.Println(piscine.Fibonacci(arg1))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Sqrt(4))
// 	fmt.Println(piscine.Sqrt(3))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.IsPrime(5))
// 	fmt.Println(piscine.IsPrime(4))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.FindNextPrime(5))
// 	fmt.Println(piscine.FindNextPrime(4))
// }

// package main

// import (
// 	"piscine"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	z01.PrintRune(piscine.FirstRune("Hello!"))
// 	z01.PrintRune(piscine.FirstRune("Salut!"))
// 	z01.PrintRune(piscine.FirstRune("Ola!"))
// 	z01.PrintRune('\n')
// }

// package main

// import (
// 	"piscine"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	z01.PrintRune(piscine.LastRune("Hello!"))
// 	z01.PrintRune(piscine.LastRune("Salut!"))
// 	z01.PrintRune(piscine.LastRune("Ola!"))
// 	z01.PrintRune('\n')
// }

// package main

// import (
// 	"piscine"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	z01.PrintRune(piscine.NRune("Hello!", 3))
// 	z01.PrintRune(piscine.NRune("Salut!", 2))
// 	z01.PrintRune(piscine.NRune("Bye!", -1))
// 	z01.PrintRune(piscine.NRune("Bye!", 5))
// 	z01.PrintRune(piscine.NRune("Ola!", 4))
// 	z01.PrintRune('\n')
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Compare("Hello!", "Hello!"))
// 	fmt.Println(piscine.Compare("Salut!", "lut!"))
// 	fmt.Println(piscine.Compare("Ola!", "Ol"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := "Hello 78 World!    4455 /"
// 	nb := piscine.AlphaCount(s)
// 	fmt.Println(nb)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	// fmt.Println(piscine.Index("Hello!", "l"))
// 	// fmt.Println(piscine.Index("Salut!", "alu"))
// 	// fmt.Println(piscine.Index("Ola!", "hOl"))
// 	// fmt.Println(piscine.Index("Salut!", "ut!"))
// 	fmt.Println(piscine.Index("\\gm^45[*Ol7V[", "5[*Ol7V"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Concat("Hello!", " How are you?"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.IsUpper("HELLO"))
// 	fmt.Println(piscine.IsUpper("HELLO!"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Capitalize("6\"Uu5TdN>xI\"E"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(piscine.Join(toConcat, ":"))
// }

// package main

// import "piscine"

// func main() {
// 	piscine.PrintNbrInOrder(143)
// 	piscine.PrintNbrInOrder(0)
// 	piscine.PrintNbrInOrder(321)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.TrimAtoi("12345"))
// 	fmt.Println(piscine.TrimAtoi("str123ing45"))
// 	fmt.Println(piscine.TrimAtoi("012 345"))
// 	fmt.Println(piscine.TrimAtoi("Hello World!"))
// 	fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
// 	fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
// 	fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
// 	fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.AppendRange(5, 10))
// 	fmt.Println(piscine.AppendRange(10, 5))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.MakeRange(5, 10))
// 	fmt.Println(piscine.MakeRange(10, 5))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	test := []string{"Hello", "how", "are", "you?"}
// 	fmt.Println(piscine.ConcatParams(test))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := "HelloHAhowHAareHAyou?"
// 	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
// }

// package main

// import "piscine"

// func main() {
// 	a := piscine.SplitWhiteSpaces("Hello how are you?")
// 	piscine.PrintWordsTables(a)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	result := piscine.ConvertBase("101011", "01", "0123456789")
// 	fmt.Println(result)
// }

package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func f(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

// package main

// import (
// 	"fmt"

// 	"piscine"
// )

// func PrintList(l *piscine.List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}

// 	fmt.Print(nil, "\n")
// }

// func main() {
// 	link := &piscine.List{}
// 	link2 := &piscine.List{}

// 	fmt.Println("----normal state----")
// 	piscine.ListPushBack(link2, 1)
// 	PrintList(link2)
// 	piscine.ListRemoveIf(link2, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link2)
// 	fmt.Println()

// 	fmt.Println("----normal state----")
// 	piscine.ListPushBack(link, 1)
// 	piscine.ListPushBack(link, "Hello")
// 	piscine.ListPushBack(link, 1)
// 	piscine.ListPushBack(link, "There")
// 	piscine.ListPushBack(link, 1)
// 	piscine.ListPushBack(link, 1)
// 	piscine.ListPushBack(link, "How")
// 	piscine.ListPushBack(link, 1)
// 	piscine.ListPushBack(link, "are")
// 	piscine.ListPushBack(link, "you")
// 	piscine.ListPushBack(link, 1)
// 	PrintList(link)

// 	piscine.ListRemoveIf(link, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link)
// }
