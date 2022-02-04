package main

import (
	"fmt"
	"io/ioutil"
	"log"

	//"bufio"
	"crypto/sha256"
	"os"
	"syscall" //

	"golang.org/x/term" //TODO uncomment
	// "os/exec"
	// "strings"
)


func help() {
	//fmt.Println(os.Args[0])
	fmt.Println(hello)
	fmt.Println("Usage:\n\t" + os.Args[0] + " [arguments] <option> file")
	fmt.Println("\nOptions:\n\thelp\tDisplay this message\n\tlock\tLock directory/file\n\tunlock\tUnlock file/directory\n\tlicense\tDisplay license")
	fmt.Println("\nArguments:\n\t--debug\tShow debug messages")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

var debug bool = false

func main() {
	if(len(os.Args) > 2){
		if(len(os.Args) > 3){
			if(os.Args[1]=="--debug"){
				debug = true
			}
		}
		if(!debug){
			log.SetOutput(ioutil.Discard)
		}
		//filename = last args
		var filename string = os.Args[len(os.Args)-1]
		//check if filename is existing file
		log.Print("Checking if file exists")
		if _, err := os.Stat(filename); !os.IsNotExist(err) {
			log.Print("Getting password")
			//get password
			fmt.Printf("Enter password to %s: ", filename)													//
			userpass, err := term.ReadPassword(int(syscall.Stdin));											//
			if(err != nil){																					//
				fmt.Fprintln(os.Stderr, err)																			//TODO uncomment for release
			}																								//
			fmt.Print("\n")
			// fmt.Print(userpass)
			// combinate userpass with salt to make file ununlockable without locker even if user know password	and use it in md5 form
			var passwordString string = string(userpass) + salt;												//
			var password32 = sha256.Sum256([]byte(passwordString));															//
			// password32 := sha256.Sum256([]byte("haslo"))
			password := []byte(password32[:])

			// fmt.Print(password)
			switch os.Args[len(os.Args)-2] {
			case "lock":
				lock(filename, password)
				break
			case "unlock":
				unlock(filename, password)
				break
			default:
				help()
				break
			}
		} else {
			fmt.Fprintln(os.Stderr, string(filename), "does not exist")
		}
	} else if len(os.Args) == 1 {
		fmt.Print(hello)
		fmt.Print("For more information type \"locker help\"")
	} else if(os.Args[1] == "license"){
		fmt.Println(license)
	} else {
		help()
	}
}