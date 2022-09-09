package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(slice []int){
	sum := 0
	// Summing the slice
	for _, v := range slice {
		sum += v
	}
	fmt.Println("sum: ", sum)
}

func main() {
	fmt.Println("Enter your number: ")
	
	// Get x from command line
	fmt.Scanln(&x)

	slice := make([]int, x)

	for i:= 0; i< x; i++ {
		slice[i] = rand.Intn(100)
	}

	start := time.Now()
	go sum(slice)
	elapsed := time.Since(start)
	fmt.Println("Time for thread1 to sum up / print slice %s", elapsed)


	start = time.Now()
	go sum(slice)
	elapsed = time.Since(start)
	fmt.Println("Time took thread2 to sum up / print slice %s", elapsed)

	fmt.Scanln()
}