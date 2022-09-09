package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// Get x from command line
	fmt.Println("Enter your number: ")
	var x int
	fmt.Scanln(&x)

	slice := make([]int, x)

	// Summing the slice
	for i:= 0; i< x; i++ {
		slice[i] = rand.Intn(100)
	}

	start := time.Now()
	sort.Slice(slice, func(i,j int) bool{
		return slice[i] < slice[j]
	})
	elapsed := time.Since(start)
	fmt.Println("Time took to sort slice using slice sort", elapsed)
	
	start = time.Now()
	sort.SliceStable(slice, func(i,j int) bool{
		return slice[i] < slice[j]
	})
	elapsed = time.Since(start)
	fmt.Println("Time took to sort slice using slice stable sort", elapsed)

}