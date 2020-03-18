package main

import "fmt"

func evenOdd(n int) {
	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

func main() {
	var items = make([]int, 10)
	for i := 0; i < 10; i++ {
		items[i] = i
	}
	for _, item := range items {
		evenOdd(item)
	}
}
