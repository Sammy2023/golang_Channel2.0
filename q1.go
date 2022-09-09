package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(slice []int,  c chan int){
	sum := 0
	// Summing the slice
	for _, v := range slice {
		sum += v
	}
	fmt.Println("sum: ", sum)
	c <- sum
}

func main() {
	fmt.Println("Enter your number: ")
	var x int
	c := make(chan int)
	// Get x from command line
	fmt.Scanln(&x)

	slice := make([]int, x)

	for i:= 0; i< x; i++ {
		slice[i] = rand.Intn(100)
	}

	start := time.Now()
	go sum(slice, c)
	elapsed := time.Since(start)
	fmt.Println("Time for thread1 to sum up / print slice %s", elapsed)


	start = time.Now()
	go sum(slice, c)
	elapsed = time.Since(start)
	fmt.Println("Time took thread2 to sum up / print slice %s", elapsed)

	a, b := <-c, <-c
	fmt.Println(a, b)
}