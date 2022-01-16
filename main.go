package main

import (
	"fmt"
	"log"
	//"bufio"
	"os"
	"crypto/sha256"
	// "golang.org/x/term"	//TODO uncomment
	// "syscall"			//
)

func help(){
	//fmt.Println(os.Args[0]);
	fmt.Println("Usage:\n\t" + os.Args[0] + " <option> [arguments] file");
	fmt.Println("\nOptions:\n\thelp\tDisplay this message\n\tlock\tLock directory/file\n\tunlock\tUnlock file/directory");
	fmt.Println("\nArguments:\n\t--debug <true/false>\tShow debug messages\n\t--default-config\tUse default config");
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename);
    if os.IsNotExist(err) {
        return false;
    }
    return !info.IsDir();
}

func main(){
	//if option is passed do something, if not return help
	if(len(os.Args) > 2){
		//filename = last args
		var filename string = os.Args[len(os.Args)-1];
		//check if filename is existing file
		log.Print("Checking if file exists");
		if(fileExists(filename)){
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
			password32 := sha256.Sum256([]byte("haslo"));
			password := []byte(password32[:])
			
			// fmt.Print(password);
			switch os.Args[1] {
			case "lock":
				log.Print("Going to lock func");
				lock(filename, password);
				break;
			case "unlock":
				log.Print("Going to unlock func");
				unlock(filename);
				break;
			default:
				log.Print("Going to help func");
				help();
				break;
			}
		}else{
			log.Fatalf("%s does not exist or it's not a file", filename);
		}
	}else if(len(os.Args) == 1){
		fmt.Println("Locker"); //TODO add here ascii art and other stuff
	}else{
		log.Print("Going to help func");
		help();
	}
	
}