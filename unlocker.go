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
		fmt.Fprintln(os.Stderr, "invalid unpadding length")
	}
	return (data[:(length - padLen)])
}

func decrypt(filename string, password []byte) {
	encrypted, err := ioutil.ReadFile(filename)
	//verification of signature
	if string(encrypted[:len(signature)]) != signature {
		fmt.Fprintln(os.Stderr, "Signature is invalid")
	}
	encrypted = encrypted[len(signature):]
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	log.Print("file data: ", string(encrypted))

	aesDec, err := aes.NewCipher(password)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if len(encrypted)%aesDec.BlockSize() != 0 {
		fmt.Fprintln(os.Stderr, "invalid encrypted data")
	}
	if len(encrypted) < 1 {
		fmt.Fprintln(os.Stderr, "encrypted data cannot be < 1")
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
		fmt.Fprintln(os.Stderr, err)
	}
	defer f.Close()
	//write data
	if _, err = f.WriteString(string(decrypted)); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func untar(filename string) {
	tarFile, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error while opening decrypted file: ", err)
	}
	gr, err := gzip.NewReader(tarFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error while creating gzip reader: ", err)
	}
	tr := tar.NewReader(gr)

	for true {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "error while getting tar header: ", err)
		}

		comp := header.Name

		if comp == "" || strings.Contains(comp, `\`) || strings.HasPrefix(comp, "/") || strings.Contains(comp, "../") {
			fmt.Println("WARNING: tared file contains invalid path: " + comp)
			fmt.Println("Would you like to unpack data anyway? (y/N): ")
			var out string
			fmt.Scanln(&out)
			if out != "y" {
				fmt.Fprintln(os.Stderr, "exited program to avoid unpacking unwanted file")
			}
		}
		comp = filepath.ToSlash(comp)

		if header.Typeflag == tar.TypeDir {
			err := os.MkdirAll(comp, 0744)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error while unpacking directory: ", err)
			}
		} else if header.Typeflag == tar.TypeReg {
			writer, err := os.Create(comp)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error while creating file: ", err)
			}

			_, err = io.Copy(writer, tr)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error while writing untared file: ", err)
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
	if out != "n" {
		log.Print("removing " + filename)
		addrw(filename)
		os.RemoveAll(filename)
	}
}

func addrw(filename string) {
	log.Print("adding rw perms to " + filename)
	os.Chmod(filename, 0644)
}

func unlock(filename string, password []byte) {
	log.Println("path: " + filename)
	addrw(filename)
	decrypt(filename, password)
	untar(string(filename + ".locker.unlock"))
	os.RemoveAll(string(filename + ".locker.unlock"))
	removeEnc(filename)
	os.Chmod(filename[:len(".locker")], 0744)
	// getInfo(filename)
}
