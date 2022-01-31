package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
)

func main(){
	size, _ := strconv.Atoi(os.Args[1])
	// fmt.Print(size)

	x := make([]byte, size)

	rand.Read(x)
	fmt.Print(x)
}