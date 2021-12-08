package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateNumb() int {
	var number int
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	number = generator.Intn(101)
	return number
}
func main() {
	var stavka float64
	var num, rightnumb int
	rightnumb = GenerateNumb()
	fmt.Print("Enter Your Bid: ")
	fmt.Scan(&stavka)
	fmt.Print("Select the integer you want to bet on from 0 to 100: ")
	fmt.Scan(&num)
	fmt.Print("The drum is spinning\n")
	fmt.Println("The number dropped out: ", rightnumb)

	if rightnumb == num {
		fmt.Println("You have won: ", stavka*2)
	} else if stavka < 100 {
		fmt.Print("You have lost your money\n")
	} else if stavka > 1000 && stavka < 1000000 {
		fmt.Print("You have lost your car\n")
	} else if stavka > 1000000 {
		fmt.Print("You have lost your life. Collectors will come for you in 30 minutes\n")
	}
}
