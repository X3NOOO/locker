package main

import (
	// "bytes"
	// "encoding/json"
	// "fmt"

	// "io/fs"
	"archive/tar"
	"compress/gzip"
	"io"

	"bufio"
	"crypto/aes"
	"crypto/sha1"
	"fmt"

	// "io"
	"path/filepath"
	"strconv"

	// "io/fs"
	"bytes"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	// "path/filepath"
)

func getMod(filename string) fs.FileMode {
	//get file's current mode for restoring in while unlocking
	var stats, err = os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	oldMode := stats.Mode().Perm()
	log.Printf("filemode: %s", oldMode)

	return (oldMode)
}

func padding(data []byte, blockSize int) []byte {
	paddingLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	return append(data, padding...)
}

func tarData(filename string) string {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal("error while getting file info: ", err)
	}
	tarName := string(filename + ".tar.gz.locker")
	tarFile, err := os.Create(tarName)
	if err != nil {
		log.Fatal("error while creating tar file: ", err)
	}
	defer tarFile.Close()
	//create tar ang gzip writers
	log.Print("crating gzip and tar writers")
	gw := gzip.NewWriter(tarFile)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	//if filename is a file
	if fi.Mode().IsRegular() {
		log.Print(filename, " is a file")
		header, err := tar.FileInfoHeader(fi, filename)
		if err != nil {
			log.Fatal("error while getting tar header: ", err)
		}
		err = tw.WriteHeader(header)
		if err != nil {
			log.Fatal("error while writing tar header: ", err)
		}
		tmp, err := os.Open(filename)
		if err != nil {
			log.Fatal("error while opening file for writing tar: ", err)
		}
		_, err = io.Copy(tw, tmp)
		if err != nil {
			log.Fatal("error while copying file to tar: ", err)
		}
	} else if fi.Mode().IsDir() {
		log.Print(filename, " is a directory")
		filepath.Walk(filename, func(file string, fi os.FileInfo, err error) error {
			header, err := tar.FileInfoHeader(fi, file)
			if err != nil {
				log.Fatal("error while getting tar header: ", err)
			}
			header.Name = filepath.ToSlash(file)
			log.Print("filepath to file in tar: ", header.Name)
			err = tw.WriteHeader(header)
			if err != nil {
				log.Fatal("error while writing header: ", err)
			}
			if !fi.IsDir() {
				data, err := os.Open(file)
				if err != nil {
					log.Fatal("error while opening file: ", err)
				}
				_, err = io.Copy(tw, data)
				if err != nil {
					log.Fatal("error while copying file: ", err)
				}
			}
			return (nil)
		})
	} else {
		log.Fatal("cannot recognize file type")
	}
	return (tarName)
}

func pause() {
	//pause program
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func encrypt(filename string, key []byte) {
	// fi, err := os.Stat(filename)
	// if(err != nil){
	// log.Fatal("error while getting file info: ", err)
	// }
	//read file to file
	//log.Printf("Reading %s to file", filename)
	// file, err := ioutil.ReadFile(filename)
	// if(err != nil){
	// 	log.Fatal("error while reading file: ", err)
	// }
	// log.Print("file data: ", string(file))
	///tar file
	//create tarball

	//read file from filename as fileData, tar data, encrypt it with aes with password key
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("error while creating aescipher: ", err)
	}
	file, err := ioutil.ReadFile(filename)
	// log.Printf("tared file: %x", file)
	if err != nil {
		log.Fatal("error while reading tared file: ", err)
	}
	// fileData = padded file
	fileData := padding(file, aesBlock.BlockSize())
	file = nil //free the memory
	// log.Printf("padded file: %x", fileData)

	// var encrypted []byte
	tmpData := make([]byte, aesBlock.BlockSize())

	for i := 1; i <= verifyData; i++ {
		for index := 0; index < len(fileData); index += aesBlock.BlockSize() {
			aesBlock.Encrypt(tmpData, fileData[index:index+aesBlock.BlockSize()])
			f, err := os.OpenFile(string(filename+"."+strconv.Itoa(i)), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			//encrypted = append(encrypted, tmpData...)
			//write encrypted block to {filename}.locker.{i}

			//write signature
			if index == 0 {
				if _, err = f.WriteString(signature); err != nil {
					log.Fatal(err)
				}
			}

			//write data
			if _, err = f.WriteString(string(tmpData)); err != nil {
				log.Fatal(err)
			}
		}
	}
	// log.Print("encrypted: ", string(encrypted))
	// return(encrypted)
}

// func write(filename string, encrypted []byte){
//// write data to filename.locker, do this second time to , if
// data = append([]byte(signature), encrypted)
// encrypted = nil; //free the memory
// }

func verifyEnc(filename string) {
	var oldHash [20]byte
	for i := 1; i <= verifyData; i++ {
		file, err := ioutil.ReadFile(filename + "." + strconv.Itoa(i))
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		tmp := sha1.Sum(file)
		if i == 1 {
			oldHash = tmp
			} else {
				if oldHash != tmp {
					log.Fatal("error while verification of encrypted files: files are different")
				}
			}
		}
		log.Print("output files are the same")
	}
	
func removeCopies(filename string, start bool, originalName string) {
	log.Print("originalName: " + originalName)
	if start {
		for i := 2; i <= verifyData; i++ {
			err := os.Remove(filename + "." + strconv.Itoa(i))
			if err != nil {
				log.Fatal("error while removing copies: ", err)
			}
			os.Rename(filename+".1", filename+".locker")
		}
	}

	fmt.Print("Would you like to replace original file with locker version (not recommended in case of folders)? (y/N): ")
	var out string
	fmt.Scanln(&out)
	log.Print("replace out: ", out)
	// fmt.Print(filename)
	if string(out) == "y" || string(out) == "Y"{
		// os.Remove(filename)
		os.RemoveAll(originalName)
		os.Rename(filename+".locker", originalName)
		changeMod(originalName)
	} else {
		os.Rename(filename+".locker", originalName+".locker")
		changeMod(filename + ".locker")
	
		fmt.Print("Would you like to remove original file? (Y/n): ")
		fmt.Scanln(&out)
		if(out != "n"){
			log.Print("removing " + originalName)
			os.Chmod(originalName, 0777)
			os.RemoveAll(originalName)
		}
		changeMod(originalName + ".locker")
	}
	os.RemoveAll(filename)
}

func changeMod(filename string) {
	log.Print("Changing mod to 000: " + filename)
	os.Chmod(filename, 000)
}

func lock(filename string, password []byte) {
	log.Println("path: " + filename)

	tarName := tarData(filename)
	encrypt(tarName, password)

	if verifyData > 1 {
		verifyEnc(tarName)
		removeCopies(tarName, true, filename)
	} else {
		removeCopies(tarName, false, filename)
	}
}
