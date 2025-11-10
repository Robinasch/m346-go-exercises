package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var eyes = rand.Intn(6) + 1 
	var when = time.Now()

	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		panic(err)
	}
	defer eyesFile.Close()

	logFile, err := os.Create("dice.log")
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")
	fmt.Fprintln(logFile, "the dice was rolled at", when)

	fmt.Println("Results written to eyes.txt and dice.log")
}
