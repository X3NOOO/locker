package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/aes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	// "strconv"
	// "encoding/json"
	// "io/ioutil"
)

func removePadding(data []byte) []byte {
	length := len(data)
	padLen := int(data[length-1])

	if length < padLen {
		log.Fatal("invalid unpadding length")
	}
	return (data[:(length - padLen)])
}

func decrypt(filename string, password []byte) {
	encrypted, err := ioutil.ReadFile(filename)
	//verification of signature
	if string(encrypted[:len(signature)]) != signature {
		log.Fatal("Signature is invalid")
	}
	encrypted = encrypted[len(signature):]
	if err != nil {
		log.Fatal(err)
	}
	log.Print("file data: ", string(encrypted))

	aesDec, err := aes.NewCipher(password)
	if err != nil {
		log.Fatal(err)
	}
	if len(encrypted)%aesDec.BlockSize() != 0 {
		log.Fatal("invalid encrypted data")
	}
	if len(encrypted) < 1 {
		log.Fatal("encrypted data cannot be < 1")
	}
	//decrypt data
	var decrypted []byte
	tmpData := make([]byte, aesDec.BlockSize())

	for index := 0; index < len(encrypted); index += aesDec.BlockSize() {
		aesDec.Decrypt(tmpData, encrypted[index:index+aesDec.BlockSize()])
		decrypted = append(decrypted, tmpData...)
	}
	decrypted = removePadding(decrypted)
	//untar data
	// decrypted =
	f, err := os.OpenFile(string(filename+".locker.unlock"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//write data
	if _, err = f.WriteString(string(decrypted)); err != nil {
		log.Fatal(err)
	}
}

func untar(filename string) {
	tarFile, err := os.Open(filename)
	if(err != nil){
		log.Fatal("error while opening decrypted file: ", err)
	}
	gr, err := gzip.NewReader(tarFile)
	if(err != nil){
		log.Fatal("error while creating gzip reader: ", err)
	}
	tr := tar.NewReader(gr)

	for true {
		header, err := tr.Next()
		if(err == io.EOF){
			break
		}
		if(err != nil){
			log.Fatal("error while getting tar header: ", err)
		}

		comp := header.Name

		if(comp == "" || strings.Contains(comp, `\`) || strings.HasPrefix(comp, "/") || strings.Contains(comp, "../")){
			fmt.Println("WARNING: tared file contains invalid path: " + comp)
			fmt.Println("Would you like to unpack data anyway? (y/N): ")
			var out string
			fmt.Scanln(&out)
			if(out != "y"){
				log.Fatal("exited program to avoid unpacking unwanted file")
			}
		}
		comp = filepath.ToSlash(comp)

		if(header.Typeflag == tar.TypeDir){
			err := os.MkdirAll(comp, 0744)
			if(err != nil){
				log.Fatal("error while unpacking directory: ", err)
			}
		}else if(header.Typeflag == tar.TypeReg){
			writer, err := os.Create(comp)
			if(err != nil){
				log.Fatal("error while creating file: ", err)
			}

			_, err = io.Copy(writer, tr)
			if(err != nil){
				log.Fatal("error while writing untared file: ", err)
			}
			writer.Close()
			os.Chmod(comp, os.FileMode(header.Mode))
		}
	}
}

func removeEnc(filename string) {
	fmt.Print("Would you like to remove encrypted version? (Y/n): ")
	var out string
	fmt.Scanln(&out)
	if(out != "n"){
		log.Print("removing " + filename)
		addrw(filename)
		os.RemoveAll(filename)
	}
}

func addrw(filename string){
	log.Print("adding rw perms to " + filename)
	os.Chmod(filename, 0644)
}

func unlock(filename string, password []byte) {
	log.SetOutput(ioutil.Discard)
	log.Println("path: " + filename)
	addrw(filename)
	decrypt(filename, password)
	untar(string(filename+".locker.unlock"))
	os.RemoveAll(string(filename+".locker.unlock"))
	removeEnc(filename)
	os.Chmod(filename[:len(".locker")], 0744)
	// getInfo(filename)
}
