package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateNumb() int {
	var number int
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	number = generator.Intn(11)
	return number
}
func GenerateRow() rune {
	chars := []rune("💔⭕❌")
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	c := chars[generator.Intn(len(chars))]
	//number = generator.Intn(101)
	return c
}
func main() {
	var stavka, yourmoney float64
	var num, rightnumb int
	var choiceNull, choiceFirst string

	a := GenerateRow()
	b := GenerateRow()
	c := GenerateRow()
	rightnumb = GenerateNumb()
	fmt.Print("Choose game(1/2): (Roulette/3 in row) ")
	fmt.Scan(&choiceFirst)

	if choiceFirst == "1" {
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
			fmt.Print("Do you want continue? y/n\n")
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
			rightnumb = GenerateNumb()
			fmt.Println("The number dropped out: ", rightnumb)
			if rightnumb == num {
				yourmoney = yourmoney + stavka
				fmt.Println("You have won: ", stavka*2)
			}
			if rightnumb != num {
				yourmoney = yourmoney - stavka*2
				fmt.Println("You have lost: ", stavka*2)
			}
			if yourmoney < 0 && yourmoney > -1000 {
				fmt.Println("You have not enough money: ", stavka*2)
			} else if yourmoney > -1000 && stavka < -1000000 {
				fmt.Print("You have lost your car\n")
			} else if yourmoney < -1000000 {
				fmt.Print("You have lost your life. Collectors will come for you in 30 minutes\n")
			}

		}
	} else if choiceFirst == "2" {
		fmt.Print("How much money do you have: ")
		fmt.Scan(&yourmoney)
		fmt.Print("Enter Your Bid: ")
		fmt.Scan(&stavka)
		fmt.Print("The drum is spinning\n")
		fmt.Println(string(a))
		fmt.Println(string(b))
		fmt.Println(string(c))
		if yourmoney < stavka && a != b && a != c {
			yourmoney = -100
			fmt.Print("You're going to die now. Good night\n")
		}
		if a == b && a == c {
			yourmoney = yourmoney + stavka
			fmt.Println("You have won: ", stavka*2)
		}
		for yourmoney > 0 {
			fmt.Print("Do you want continue? y/n\n")
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

			fmt.Print("The drum is spinning\n")
			a = GenerateRow()
			b = GenerateRow()
			c = GenerateRow()
			fmt.Println(string(a))
			fmt.Println(string(b))
			fmt.Println(string(c))
			if a == b && a == c {
				yourmoney = yourmoney + stavka
				fmt.Println("You have won: ", stavka*2)
			}
			if a != b || a != c {
				yourmoney = yourmoney - stavka*2
				fmt.Println("You have lost: ", stavka*2)
			}
			if yourmoney < 0 && yourmoney > -1000 {
				fmt.Println("You have not enough money: ", stavka*2)
			} else if yourmoney > -1000 && stavka < -1000000 {
				fmt.Print("You have lost your car\n")
			} else if yourmoney < -1000000 {
				fmt.Print("You have lost your life. Collectors will come for you in 30 minutes\n")
			}

		}
	}

}
