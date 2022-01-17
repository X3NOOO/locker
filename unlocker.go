package main

import (
	"log"
	"io/ioutil"
	"crypto/aes"
	"os"
	"fmt"
	// "strconv"
	// "encoding/json"
    // "io/ioutil"
)

func removePadding(data []byte) []byte{
	length := len(data);
	padLen := int(data[length-1]);

	if(length < padLen){
		log.Fatal("invalid unpadding length");
	}
	return(data[:(length - padLen)]);
}

func decrypt(filename string, password []byte){
	encrypted, err := ioutil.ReadFile(filename);
	encrypted = encrypted[len(signature):];
	if(err != nil){
		log.Fatal(err);
	}
	log.Print("file data: ", string(encrypted));

	aesDec, err := aes.NewCipher(password);
	if(err != nil){
		log.Fatal(err);
	}
	if(len(encrypted)%aesDec.BlockSize() != 0){
		log.Fatal("invalid encrypted data");
	}
	if(len(encrypted) < 1){
		log.Fatal("encrypted data cannot be < 1");
	}

	var decrypted []byte;
	tmpData := make([]byte, aesDec.BlockSize())

	for index := 0; index < len(encrypted); index += aesDec.BlockSize() {
		aesDec.Decrypt(tmpData, encrypted[index:index+aesDec.BlockSize()]);
		decrypted = append(decrypted, tmpData...);
	}
	decrypted = removePadding(decrypted);

	f, err := os.OpenFile(string(filename + ".locker.unlock"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644);
	if(err != nil){
		log.Fatal(err);
	}
	defer f.Close();
	//write data
	if _, err = f.WriteString(string(decrypted)); err != nil {
		log.Fatal(err);
	}
}

func removeEnc(filename string){
	fmt.Print("Would you like to remove encrypted version? (Y/n): ");
	var out string;
	fmt.Scanln(&out);
	log.Print("replace out: ", out)
	os.Remove(filename);
	os.Rename(filename + ".locker.unlock", filename);
}

func unlock(filename string, password []byte){
	log.Println("path: " + filename);
	os.Chmod(filename, 0644);
	decrypt(filename, password);
	removeEnc(filename);
	// getInfo(filename);
}