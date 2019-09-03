package main

import (
	"fmt"
	
)

func main() {
	var liste []int
	for i := 0; i < 10; i++ {
		liste = append(liste, i)
	}
	fmt.Println(liste)
}
