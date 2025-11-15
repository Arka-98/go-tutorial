package problems

import "fmt"

/*
Implements star pattern
*/
func StarPattern(height int) {
	for i := range height {
		str := ""

		for range i + 1 {
			str += "* "
		}

		fmt.Println(str)
	}
}
