package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	// "log"
	"encoding/hex"
)

func main() {
	// log.Print(os.Args)
	size, _ := strconv.Atoi(os.Args[2])
	rn := make([]byte, size)
	rand.Read(rn)
	rn2 := hex.EncodeToString(rn)

	// log.Print(rn2)
	input, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, "var salt string") {
			if os.Args[1] == "l" {
				// log.Print("args 1 l")
				out := "var salt string = \"" + string(rn2) + "\""
				lines[i] = out
				// log.Print(out)
			} else if os.Args[1] == "u" {
				// log.Print("args 1 u")
				out := "var salt string = \"DEFAULTPASSWORD\""
				lines[i] = out
				// log.Print(out)
			}
		}
	}
	output := strings.Join(lines, "\n")
	// log.Print(output)
	err = ioutil.WriteFile(os.Args[3], []byte(output), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
