package main

import (
	"fmt"
	"log"

	//"bufio"
	"crypto/sha256"
	"os"
	// "golang.org/x/term"	//TODO uncomment
	// "syscall"			//
)

func help() {
	//fmt.Println(os.Args[0])
	fmt.Println(hello)
	fmt.Println("Usage:\n\t" + os.Args[0] + " <option> [arguments] file")
	fmt.Println("\nOptions:\n\thelp\tDisplay this message\n\tlock\tLock directory/file\n\tunlock\tUnlock file/directory\n\tlicense\tDisplay license")
	fmt.Println("\nArguments:\n\t--debug <true/false>\tShow debug messages\n\t--default-config\tUse default config")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	//if option is passed do something, if not return help
	if len(os.Args) > 2 {
		//filename = last args
		var filename string = os.Args[len(os.Args)-1]
		//check if filename is existing file
		log.Print("Checking if file exists")
		if _, err := os.Stat(filename); !os.IsNotExist(err) {
			log.Print("Getting password")
			//get password
			// fmt.Printf("Enter password to %s: ", filename)													//
			// userpass, err := term.ReadPassword(int(syscall.Stdin));											//
			// if(err != nil){																					//
			// 	log.Fatal(err)																					//TODO uncomment for release
			// }																								//
			// combinate userpass with salt to make file ununlockable without locker even if user know password	and use it in md5 form
			// var passwordString string = string(userpass) + salt;												//
			// var password = sha356.Sum256([]byte(passwordString));															//
			password32 := sha256.Sum256([]byte("haslo"))
			password := []byte(password32[:])

			// fmt.Print(password)
			switch os.Args[1] {
			case "lock":
				log.Print("Going to lock func 1")
				lock(filename, password)
				break
			case "unlock":
				log.Print("Going to unlock func 2")
				unlock(filename, password)
				break
			default:
				log.Print("Going to help func 3")
				help()
				break
			}
		} else {
			log.Fatalf("%s does not exist or it's not a file", filename)
		}
	} else if len(os.Args) == 1 {
		fmt.Println(hello) 
	} else if(os.Args[1] == "license"){
		fmt.Println(license)
	} else {
		log.Print("Going to help func 4")
		help()
	}

}
