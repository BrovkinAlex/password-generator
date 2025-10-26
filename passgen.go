package main

import (
	"fmt"
	"math/rand"
)

func main() {
	passLen := 16
	segLen := passLen / 4

	uppercaseSym := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseSym := "abcdefghijklmnopqrstuvwxyz"
	numberSym := "1234567890"
	cpecialSym := "!@#$%^&*"

	outPass := make([]rune, passLen)

	fillArr(outPass, 0, segLen, uppercaseSym)
	fillArr(outPass, segLen, segLen*2, lowercaseSym)
	fillArr(outPass, segLen*2, segLen*3, numberSym)
	fillArr(outPass, segLen*3, passLen, cpecialSym)

	shakeArray(outPass)

	fmt.Println(string(outPass))
}

func giveRndSym(syms string) rune {
	rndNum := rand.Intn(len(syms))

	rndRune := syms[rndNum]

	return rune(rndRune)
}

func fillArr(arr []rune, a, b int, src string) {
	for ; a != b; a++ {
		arr[a] = giveRndSym(src)
	}
}

func shakeArray(arr []rune) {
	for i := len(arr) - 1; i > 0; i-- {
		rndNum := rand.Intn(i + 1)
		arr[rndNum], arr[i] = arr[i], arr[rndNum]
	}
}
