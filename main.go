package main

import (
	"fmt"
	"strings"
)

func main() {

	// 	input := ["hello", "he", "hellbound", "heaven"]
	// pref := hel

	input := []string{"hello", "he", "hellbound", "heaven"}

	pref := "hel"

	fmt.Println(verify(input, pref))

}

func verify(input []string, pref string) int {

	var count, cheak int

	for i := 0; i < len(input); i++ {

		find := strings.Split(pref, "")
		strTemp := strings.Split(input[i], "")

		count = 0

		for j := 0; j < len(find); j++ {

			if j < len(strTemp) {
				if find[j] == strTemp[j] {
					count++
				}

			}

		}

		if count == len(find) {
			cheak++
		}

	}
	return count
}
