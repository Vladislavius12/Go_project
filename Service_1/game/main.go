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

		a := step1(tttb, "X")
		bprinter(a)
		b := step1(a, "O")
		bprinter(b)
		c := step1(b, "X")
		bprinter(c)
		d := step1(c, "O")
		bprinter(d)
		e := step1(d, "X")
		bprinter(e)
		Legendprinter()
		be := boardwinchecker(e)
		if be == "nowinner" {
			f := step1(e, "O")
			bprinter(f)
			bf := boardwinchecker(f)
			if bf == "nowinner" {
				bprinter(f)
				g := step1(f, "X")
				bg := boardwinchecker(g)
				if bg == "nowinner" {
					bprinter(g)
					h := step1(g, "O")
					bh := boardwinchecker(h)
					if bh == "nowinner" {
						bprinter(h)
						i := step1(h, "X")
						bi := boardwinchecker(i)
						if bi == "nowinner" {
							Legendprinter()
							lastmessage := "CAT's GAME. (It was a draw.)"
							bprinter(i)
							fmt.Println(lastmessage)
						} else {
							bprinter(i)
							fmt.Println(bi)
						}
					} else {
						bprinter(h)
						fmt.Println(bh)
					}
				} else {
					bprinter(g)
					fmt.Println(bg)
				}
			} else {
				bprinter(f)
				fmt.Println(bf)
			}
		} else {
			bprinter(e)
			fmt.Println(be)
		}
	}
}
