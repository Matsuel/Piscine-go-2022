package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func numberOfBytes(args []string) (int, []string) {
	n := len(args)
	nbytes := 0
	var files []string
	for i, v := range args {
		var err error
		_, err = strconv.Atoi(v)
		if v == "-c" {
			if i >= n-1 {
				fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.")
				os.Exit(1)
			}
			arg := args[i+1]

			nbytes, err = strconv.Atoi(arg)

			if err != nil {
				fmt.Printf("tail: invalid number of bytes: %s\n", arg)
				os.Exit(1)
			}
			continue
		}

		if err != nil {
			files = append(files, v)
		}

	}
	return nbytes, files
}

func fileSize(fi *os.File) int64 {
	fil, err := fi.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return fil.Size()
}

func main() {
	n := len(os.Args)
	if n < 4 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	nbytes, files := numberOfBytes(os.Args[1:])
	printName := len(files) > 1
	for j, f := range files {
		fi, err := os.Open(f)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", f)
			os.Exit(1)
		}
		if printName {
			fmt.Printf("==> %s <==\n", f)
		}
		read := make([]byte, int(nbytes))
		_, er := fi.ReadAt(read, fileSize(fi)-int64(nbytes))
		if er != nil {
			fmt.Println(er.Error())
		}
		for _, c := range read {
			z01.PrintRune(rune(c))
		}
		if j < len(files)-1 {
			z01.PrintRune('\n')
		}
		fi.Close()
	}
}

func Atoi(s string) int {
	atoi := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			atoi = atoi*10 + int(ch-'0')
		} else {
			return 0
		}
	}
	return atoi
}

func Itoa(i int) string {
	itoa := ""
	for i != 0 {
		ch := i % 10
		i /= 10
		itoa += string(ch + '0')
	}
	res := ""
	for i := len(itoa) - 1; i >= 0; i-- {
		res += string(itoa[i])
	}
	return res
}
