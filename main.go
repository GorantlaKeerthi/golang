package main

import "training/main/work"

func main() {
	/*var s []int
	for {
		var a string
		fmt.Println("if you want continue then press Y/y, to exit: q/Q ")
		fmt.Scan(&a)
		if a == "q" || a == "Q" {
			break

		} else if a == "y" || a == "Y" {
			var d int
			fmt.Println("enter a number")
			fmt.Scan(&d)
			s = append(s, d)
		}

	}
	fmt.Println("slice values: ", s)
	*/
	/*var rows int
	var k int = 0
	fmt.Print("Enter number of rows :")
	fmt.Scan(&rows)
	for i := 1; i <= rows; i++ {
		k = 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print(" ", k, " ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}*/
	/*var rows int
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
	}*/

	/* type employee struct {
		EmployeeId   int
		Employeename string
		Age          int
		Country      string
	}
	fmt.Println(employee{1100, "john", 25, "USA"}) */
	work.Infiniteloop()
	//work.Pyramid()
	//work.Concat()
	//work.Sort()
}
