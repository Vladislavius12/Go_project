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
	var stavka, yourmoney float64
	var num, rightnumb int
	var choiceNull string
	rightnumb = GenerateNumb()
	fmt.Print("How much money do you have: ")
	fmt.Scan(&yourmoney)
	fmt.Print("Enter Your Bid: ")
	fmt.Scan(&stavka)
	fmt.Print("Select the integer you want to bet on from 0 to 100: ")
	fmt.Scan(&num)
	fmt.Print("The drum is spinning\n")
	fmt.Println("The number dropped out: ", rightnumb)

	if yourmoney < stavka && num != rightnumb {
		yourmoney = -100
		fmt.Print("You're going to die now. Good night\n")
	}
	if rightnumb == num {
		yourmoney = yourmoney + stavka
		fmt.Println("You have won: ", stavka*2)
	}
	for yourmoney > 0 {
		fmt.Print("Do you want continue?\n")
		fmt.Scan(&choiceNull)
		switch choiceNull {
		case "y":
			fmt.Println("You are welcome")
			break
		case "n":
			return
		default:
			fmt.Println("I didn't understand whether we were playing or not")
			continue
		}
		fmt.Print("Enter Your Bid: ")
		fmt.Scan(&stavka)
		fmt.Print("Select the integer you want to bet on from 0 to 100: ")
		fmt.Scan(&num)
		fmt.Print("The drum is spinning\n")
		fmt.Println("The number dropped out: ", rightnumb)
		if rightnumb == num {
			yourmoney = yourmoney + stavka
			fmt.Println("You have won: ", stavka*2)
		}
		if rightnumb != num {
			yourmoney = yourmoney - stavka*2
			fmt.Println("You have lost: ", stavka*2)
		}
		if yourmoney < 0 {
			fmt.Println("You have not enough: ", stavka*2)
		} else if yourmoney < -1000 && stavka < -1000000 {
			fmt.Print("You have lost your car\n")
		} else if stavka < -1000000 {
			fmt.Print("You have lost your life. Collectors will come for you in 30 minutes\n")
		}

	}

}
