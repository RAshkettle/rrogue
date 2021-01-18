package main

import (
	"crypto/rand"
	"math/big"
)

//Returns a number between the two numbers inclusive.
func GetRandomBetween(low int, high int) int {
	var randy int = -1
	for {
		randy = GetDiceRoll(high)
		if randy >= low {
			break
		}
	}

	return randy
}

//GetRandomInt returns an integer from 0 to the number - 1
func GetRandomInt(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64())

}

//GetDiceRoll returns an integer from 1 to the number
func GetDiceRoll(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64()) + 1

}
