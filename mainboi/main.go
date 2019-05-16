package main

import (
	"log"
	"strconv"
)

func main() {
	closer := closeIt("1" +
		"")
	yoland := closer()
	log.Print(yoland)
}

func closeIt(lmao string) func() int {
	x := lmao
	return func() int {
		num, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal("That's not an int!")
		}
		return num
	}
}