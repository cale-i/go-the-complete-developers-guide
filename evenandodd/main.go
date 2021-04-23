package main

import "fmt"

func main() {
	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}

// func main() {
// 	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	for _, number := range numbers {
// 		if number%2 == 0 {
// 			fmt.Println(number, "is even")
// 		} else {
// 			fmt.Println(number, "is odd")
// 		}
// 	}
// }
