package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	// "io/fs"
	// "io/ioutil"
	"log"
	"os"
)

func dumpInfo(filename string){
	//get file's current mode for restoring in while unlocking
	var stats, err = os.Stat(filename);
	if err != nil {
		log.Fatal(err);
	}
	var oldMode = stats.Mode();

	//dump data to struct for saving it
	file := fileStruct{
		Path: filename,
		Mode: oldMode,
	}
	log.Println("file: ", file);
	//convert fileStruct to json
	jsonFile, err := json.Marshal(file);
	if err != nil {
		log.Fatal(err);
	}
	log.Print("jsonFile: ", string(jsonFile));
	//write json to path in configPath
	log.Print("writing jsonFile at the top of " + filename);
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, oldMode)
	if err != nil {
	    panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(fmt.Sprint(oldMode)); err != nil {
	    panic(err)
}
}

func changeMod(filename string){
	os.Chmod(filename, 000);
}

func lock(filename string) {
	log.Println("path: " + filename);

	dumpInfo(filename);
	//changeMod(filename);
}
