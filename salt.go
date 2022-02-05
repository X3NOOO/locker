/*
*locker (https://github.com/X3NOOO/locker)
*Copyright (C) 2022 X3NO <X3NO@disroot.org> [https://X3NO.ct8.pl] [https://github.com/X3NOOO]
*
*This program is free software: you can redistribute it and/or modify
*it under the terms of the GNU General Public License as published by
*the Free Software Foundation, either version 3 of the License, or
*(at your option) any later version.
*
*This program is distributed in the hope that it will be useful,
*but WITHOUT ANY WARRANTY; without even the implied warranty of
*MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*GNU General Public License for more details.
*
*You should have received a copy of the GNU General Public License
*along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
