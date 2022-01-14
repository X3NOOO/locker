package main

import (
	"fmt"
	"log"
	//"bufio"
	"os"
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
		if(fileExists(filename)){
			switch os.Args[1] {
			case "lock":
				lock(filename);
				break;
			case "unlock":
				unlock(filename);
				break;
			default:
				help();
				break;
			}
		}else{
			log.Fatalf("%s does not exist or it's not a file", filename);
		}
	}else{
		help();
	}
	
}