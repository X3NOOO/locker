package main

import (
	// "bytes"
	// "encoding/json"
	// "fmt"

	// "io/fs"
	"crypto/aes"
	"crypto/sha1"
	"fmt"
	"strconv"

	// "io/fs"
	"bytes"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func getMod(filename string) fs.FileMode{
	//get file's current mode for restoring in while unlocking
	var stats, err = os.Stat(filename);
	if err != nil {
		log.Fatal(err);
	}
	oldMode := stats.Mode().Perm();
	log.Printf("filemode: %s", oldMode)

	return(oldMode);
}

func padding(data []byte, blockSize int) []byte{
	paddingLen := blockSize - len(data)%blockSize;
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen);
	return append(data, padding...);
}

func encrypt(filename string, key []byte){
	//read file from filename as fileData, encrypt it with aes with password key
	aesBlock, err := aes.NewCipher(key);
	//read file to fileData
	log.Printf("Reading %s to fileData", filename);
    file, err := ioutil.ReadFile(filename);
	if(err != nil){
		log.Fatal(err);
	}
	log.Print("file data: ", string(file));
	mod := getMod(filename);
	// filedata = mod + file
	fileData := padding(file, aesBlock.BlockSize());
	file = nil; //free the memory
	log.Print("filemode + file: ", string(fileData));
	
	if(err != nil){
		log.Fatal(err);
	}
	fileData = padding(fileData, aesBlock.BlockSize());

	// var encrypted []byte;
	tmpData := make([]byte, aesBlock.BlockSize());
	
	for i := 1; i <= verifyData; i++{
		for index := 0; index < len(fileData); index += aesBlock.BlockSize() {
			aesBlock.Encrypt(tmpData, fileData[index:index+aesBlock.BlockSize()]);
			f, err := os.OpenFile(string(filename + ".locker." + strconv.Itoa(i)), os.O_APPEND|os.O_WRONLY|os.O_CREATE, mod);
			if err != nil {
				log.Fatal(err);
			}
			defer f.Close();
			//encrypted = append(encrypted, tmpData...);
			//write encrypted block to {filename}.locker.{i}

			//write signature
			if(index == 0){
				if _,err = f.WriteString(signature); err != nil {
					log.Fatal(err);
				}
			}

			//write data
			if _, err = f.WriteString(string(tmpData)); err != nil {
	    		log.Fatal(err);
			}
		}
	}
	// log.Print("encrypted: ", string(encrypted))
	// return(encrypted);
}

// func write(filename string, encrypted []byte){
	//// write data to filename.locker, do this second time to , if 
	// data = append([]byte(signature), encrypted);
	// encrypted = nil; //free the memory
// }

func verifyEnc(filename string){
	var oldHash [20]byte;
	for i := 1; i <= verifyData; i++{
		file, err := ioutil.ReadFile(filename + ".locker." + strconv.Itoa(i));
		if(err != nil){
			log.Fatalf("unable to read file: %v", err)
		}
		tmp := sha1.Sum(file);
		if(i == 1){
			oldHash = tmp;
		}else{
			if(oldHash != tmp){
				log.Fatal("error while verification of encrypted files: files are different")
			}
		}
	}
	log.Print("output files are the same")
}

func removeCopies(filename string, start bool){
	if(start){
		for i := 2; i <= verifyData; i++{
			err := os.Remove(filename + ".locker." + strconv.Itoa(i));
			if(err != nil){
				log.Fatal("error while removing copies: ", err);
			}
			os.Rename(filename + ".locker.1", filename + ".locker")
		}
	}
	
	fmt.Print("Would you like to replace orginal file with locker version? (Y/n): ");
	var out string;
	fmt.Scanln(&out);
	log.Print("replace out: ", out);
	if(string(out) != "n"){
		os.Remove(filename);
		os.Rename(filename + ".locker", filename);
		changeMod(filename);
	}else{
		changeMod(filename + ".locker");
	}
}

func changeMod(filename string){
	log.Print("Changing mod to 000");
	os.Chmod(filename, 000);
}

func lock(filename string, password []byte) {
	log.Println("path: " + filename);

	encrypt(filename, password);
	
	if(verifyData > 1){
		verifyEnc(filename);
		removeCopies(filename, true);
	}else{
		removeCopies(filename, false);
	}
}
