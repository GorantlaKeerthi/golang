package work

import (
	"fmt"
	"sort"
)

func Infiniteloop() {
	var s []int
	for {
		var a string
		fmt.Println("if you want continue then press Y/y, to exit: q/Q ")
		fmt.Scan(&a)
		if a == "q" || a == "Q" {
			break

		} else if a == "y" || a == "Y" {
			var d int
			fmt.Println("enter a number:")
			fmt.Scan(&d)
			s = append(s, d)
		}

	}
	fmt.Println("slice values: ", s)
}

func Pyramid() {

	var rows int
	var co int
	var space int
	var i int
	var j int
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)
	for i = 0; i < rows; i++ {
		for space = 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for j = 0; j <= i; j++ {
			if j == 0 || i == 0 {
				co = 1
			} else {
				co = co * (i - j + 1) / j
			}
			fmt.Print(" ", co, " ")
		}
		fmt.Print("\n")
	}
}
func Concat() {

	s1 := []int{1, 2, 4, 5}
	fmt.Println(s1)

	s2 := []int{3, 4, 56, 7, 86, 5}
	fmt.Println(s2)
	s3 := make([]int, 2, 3)

	//s1 = append(s1, s2...)
	//fmt.Println(s1)
	//s3 = append(s3, s1...)
	s3 = append(s3, s1...)
	//s3 = Append(s3, s2...)
	fmt.Print(s3)
}

func Sort() {
	x := []int{4, 2, 8, 5, 1, 7}
	fmt.Println("given array of integers", x)

	sort.Ints(x)
	fmt.Println(x)
	y := sort.IntsAreSorted(x)
	fmt.Println("after sorting:")
	fmt.Println("is array sorted", y)

	a := []float64{111.12, 100.021, 1.1, 12.3}
	fmt.Println("given array of float values", a)

	sort.Float64s(a)
	fmt.Println("after sorting:")
	fmt.Println(a)
	b := sort.Float64sAreSorted(a)

	fmt.Println("is array sorted", b)

	c := []string{"abc", "aef", "bcd", "hig"}
	fmt.Println("given array of string values", c)

	sort.Strings(c)
	fmt.Println("after sorting:")
	fmt.Println(c)
	d := sort.StringsAreSorted(c)

	fmt.Println("is string array sorted", d)

}
