package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var uppers = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	var lowers = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	var numbers = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var length int = 16
	uFlag := flag.Bool("u", false, "uppercase")
	lFlag := flag.Bool("l", false, "lowercase")
	nFlag := flag.Bool("n", false, "number")
	flag.Parse()

	if len(flag.Args()) > 0 {
		length, _ = strconv.Atoi(flag.Args()[0])
	}
	result := make([]rune, length)
	var list []rune

	if *uFlag {
		list = append(list, uppers...)
	}
	if *lFlag {
		list = append(list, lowers...)
	}
	if *nFlag {
		list = append(list, numbers...)
	}
	if len(list) == 0 {
		list = append(list, uppers...)
		list = append(list, lowers...)
		list = append(list, numbers...)
	}

	diceRange := len(list)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(result); i++ {
		random := rand.Intn(diceRange)
		result[i] = list[random]
	}

	fmt.Println(string(result))
}
