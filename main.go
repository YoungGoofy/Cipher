package main

import "fmt"

func main() {
	plainText := "have a nice day bro"
	keyword := "golang"
	for i := 0; i < len([]rune(plainText)); i++ {
		pl := plainText[i]
		if pl >= 'a' && pl <= 'z' {
			kw := keyword[i % len([]rune(keyword))]
			pl = (pl + kw) % 26 + 'a'
		}
		fmt.Printf("%c", pl)
	}
}
