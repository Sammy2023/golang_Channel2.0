package main

import (
	"fmt"
	"math/rand"
	"time"
)

// take slice from main
func sum(slice []int){
	sum := 0
	// go thorugh slice elements to add up sum
	for _, v := range slice {
		sum += v
	}
	//print out the sum
	fmt.Println("sum: ", sum)
}

func main() {
	fmt.Println("Enter your number: ")
	var x int
	
	//take user input
	fmt.Scanln(&x)

	//create slice based on user input
	slice := make([]int, x)
	// generate user input amount of random numbers
	for i:= 0; i< x; i++ {
		slice[i] = rand.Intn(100)
	}

	// start the time for time benchmark of gorountine sum calculation
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