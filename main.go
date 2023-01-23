package main

//Задание 4. Предыдущее занятие на if

import (
	"fmt"
)

func main() {
	var a, b, c int = 0, 0, 0
	limit1 := 10
	limit2 := 100
	limit3 := 1000

	for {
		fmt.Println(a, b, c)
		if c == limit3 {
			break
		}
		c += 1
		if b == limit2 {
			continue
		}
		b += 1
		if a == limit1 {
			continue
		}
		a += 1
	}
}
