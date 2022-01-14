package main

import (
	"io/fs"
)

type fileStruct struct {
	Path string
	Mode fs.FileMode
}

var configPath = "./config.json"
