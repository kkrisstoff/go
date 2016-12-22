package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get " + otherWord,
	"go " + otherWord,
	"lets " + otherWord,
	otherWord + "hq",
}

func main() {
	//use the current time as a random seed
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	// Scan method tells the scanner to read the next block of bytes from the input
	for s.Scan() {
		// select a random item from the transforms slice
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
