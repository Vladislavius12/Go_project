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
	fmt.Print("Введите Вашу Ставку: ")
	fmt.Scan(&stavka)
	fmt.Print("Выберете целое число на которое хотите поставить от 0 до 100: ")
	fmt.Scan(&num)
	fmt.Print("Барабан крутиться\n")
	fmt.Println("Выпало число: ", rightnumb)

	if rightnumb == num {
		fmt.Println("Вы выиграли: ", stavka*2)
	} else {
		fmt.Print("Вы проиграли\n")
	}
}
