package main

import (
	"fmt"
	"math/rand"
)

func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func createHugeString(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}

func main() {
	fmt.Println(someFunc())
}
