package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Character interface {
	name() string
	choice() string
	count(int)
}

func Mapper(letter string) int {
	if strings.TrimRight(letter, "\n") == "7" {
		return 0
	}
	if strings.TrimRight(letter, "\n") == "8" {
		return 1
	}
	if strings.TrimRight(letter, "\n") == "9" {
		return 2
	}
	if strings.TrimRight(letter, "\n") == "4" {
		return 3
	}
	if strings.TrimRight(letter, "\n") == "5" {
		return 4
	}
	if strings.TrimRight(letter, "\n") == "6" {
		return 5
	}
	if strings.TrimRight(letter, "\n") == "1" {
		return 6
	}
	if strings.TrimRight(letter, "\n") == "2" {
		return 7
	}
	if strings.TrimRight(letter, "\n") == "3" {
		return 8
	}
	return 99
}

func Bprinter(board [9]string) {
	fmt.Println("The board now looks like this:")
	fmt.Println("-------")
	fmt.Println("|" + board[0] + "|" + board[1] + "|" + board[2] + "|")
	fmt.Println("|" + board[3] + "|" + board[4] + "|" + board[5] + "|")
	fmt.Println("|" + board[6] + "|" + board[7] + "|" + board[8] + "|")
	fmt.Println("-------")
	fmt.Println("")
}

func Legendprinter() {
	fmt.Println("Here is the legend for you to make a mark:")
	fmt.Println("_______")
	fmt.Println("|7|8|9|")
	fmt.Println("|4|5|6|")
	fmt.Println("|1|2|3|")
	fmt.Println("-------")
}

func Step1(board [9]string, player string) [9]string {
	boardpos := Requestor(player)
	newboard := Marker(board, player, boardpos)
	tflag := Islistsame(board, newboard)
	if tflag {
		fmt.Println("No change was made. The same player must make a valid selection. Player " + player + " please try again.")
		newboard = Step1(board, player)
	}
	return newboard
}

func Boardwinchecker(board [9]string) string {
	lastmessage := "nowinner"
	if board[0] == board[4] { // check the diagonal from left top to right bottom corner
		if board[4] == board[8] {
			if board[4] != " " {
				lastmessage = "Player " + board[0] + " wins!"
			}
		}
	}
	if board[2] == board[4] { // check the diagonal from right top to left bottom corner
		if board[4] == board[6] {
			if board[6] != " " {
				lastmessage = "Player " + board[2] + " wins!"
			}
		}
	}
	if board[0] == board[3] { // check the left column
		if board[3] == board[6] {
			if board[6] != " " {
				lastmessage = "Player " + board[0] + " wins!"
			}
		}
	}
	if board[0] == board[1] { // check the top row
		if board[1] == board[2] {
			if board[0] != " " {
				lastmessage = "Player " + board[0] + " wins!"
			}
		}
	}
	if board[2] == board[5] { // check right column
		if board[5] == board[8] {
			if board[2] != " " {
				lastmessage = "Player " + board[2] + " wins!"
			}
		}
	}
	if board[6] == board[7] { // check bottom row
		if board[7] == board[8] {
			if board[8] != " " {
				lastmessage = "Player " + board[6] + " wins!"
			}
		}
	}
	if board[1] == board[4] { // check middle column
		if board[4] == board[7] {
			if board[1] != " " {
				lastmessage = "Player " + board[1] + " wins!"
			}
		}
	}
	if board[3] == board[4] { // check middle row
		if board[4] == board[5] {
			if board[3] != " " {
				lastmessage = "Player " + board[3] + " wins!"
			}
		}
	}
	return lastmessage
}

func Marker(board [9]string, player string, position string) [9]string {
	x := Mapper(position)
	if board[x] != " " {
		fmt.Println("That square has already been marked.  Please try again.")
	} else {
		board[x] = player
	}
	return board
}

func Requestor(player string) string {

	fmt.Println("Player", player, "it is your turn.")
	fmt.Println("Player ", player, " choose a cell:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if strings.TrimRight(text, "\n") != "7" {
		if strings.TrimRight(text, "\n") != "8" {
			if strings.TrimRight(text, "\n") != "9" {
				if strings.TrimRight(text, "\n") != "4" {
					if strings.TrimRight(text, "\n") != "5" {
						if strings.TrimRight(text, "\n") != "6" {
							if strings.TrimRight(text, "\n") != "1" {
								if strings.TrimRight(text, "\n") != "2" {
									if strings.TrimRight(text, "\n") != "3" {
										fmt.Println("*** WARNING ***")
										fmt.Println("That was an invalid entry. Please try again player " + player)
										fmt.Println(" ")
										text = Requestor(player)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return text
}

func Islistsame(board1 [9]string, board2 [9]string) bool {
	var j = true
	for i := 0; i < 9; i++ {
		if board1[i] == board2[i] {
		} else {
			j = false
		}
	}
	return j
}
