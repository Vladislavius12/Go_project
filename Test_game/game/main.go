package main

import (
	"fmt"
)

func main() {
	count := 1
	for i := 0; i < count; i++ {

		var tttb [9]string
		tttb[0] = " "
		tttb[1] = " "
		tttb[2] = " "
		tttb[3] = " "
		tttb[4] = " "
		tttb[5] = " "
		tttb[6] = " "
		tttb[7] = " "
		tttb[8] = " "

		Legendprinter()

		a := Step1(tttb, "X")
		Bprinter(a)
		b := Step1(a, "O")
		Bprinter(b)
		c := Step1(b, "X")
		Bprinter(c)
		d := Step1(c, "O")
		Bprinter(d)
		e := Step1(d, "X")
		Bprinter(e)
		Legendprinter()
		be := Boardwinchecker(e)
		if be == "nowinner" {
			f := Step1(e, "O")
			Bprinter(f)
			bf := Boardwinchecker(f)
			if bf == "nowinner" {
				Bprinter(f)
				g := Step1(f, "X")
				bg := Boardwinchecker(g)
				if bg == "nowinner" {
					Bprinter(g)
					h := Step1(g, "O")
					bh := Boardwinchecker(h)
					if bh == "nowinner" {
						Bprinter(h)
						i := Step1(h, "X")
						bi := Boardwinchecker(i)
						if bi == "nowinner" {
							Legendprinter()
							lastmessage := "CAT's GAME. (It was a draw.)"
							Bprinter(i)
							fmt.Println(lastmessage)
						} else {
							Bprinter(i)
							fmt.Println(bi)
						}
					} else {
						Bprinter(h)
						fmt.Println(bh)
					}
				} else {
					Bprinter(g)
					fmt.Println(bg)
				}
			} else {
				Bprinter(f)
				fmt.Println(bf)
			}
		} else {
			Bprinter(e)
			fmt.Println(be)
		}
	}
}
