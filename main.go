package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
	"unicode"
)

func main() {

	/* truncate value 14.65478 to 2 precision points */

	b := 14.65478
	fmt.Printf("after truncating the value is %.2f", b)
	/*conversion of string into integer and float*/
	v := "32"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("\n %T, %v\n", s, s)
	}
	f := "32.56"
	if a, err := strconv.ParseFloat(f, -1); err == nil {
		fmt.Printf("\n %T, %v", a, a)
	}
	/*assignment-3 convert string into uppercase*/
	str := "go-lang-training"

	str1 := strings.ToUpper(str)
	str2 := strings.Replace(str1, "-", " ", -1)

	fmt.Printf("\nafter conversion:\n")
	fmt.Printf(str2)
	/*assignment-4 const declaration*/
	const (
		x = 6
		y
		z
	)
	fmt.Printf("\nconst\n")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	/*assignment5- to modify the value of varible using pointers*/
	va := 100
	var ptr *int
	ptr = &va
	fmt.Printf("before changing the value and the address is")
	fmt.Print(va)

	*ptr = 200
	fmt.Printf("\n after changing the value")
	fmt.Print(va)

/*assignment6 program to find the number of letters and numbers in a string*/
	ip="1Ax3 7y Bk"
	 dig =0
	var j int =1
	var sp int=0
	var lc int =0
	var up int =0
	for i, c := range "1Ax3 7y Bk"{
		if unicode.IsDigit(c)
		incremnt(dig)
	    fmt.Println(dig)
	}
}
