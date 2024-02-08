package main

import (
	"jsCssVersioner/internal"
	"jsCssVersioner/lib"
	"strconv"
	"strings"
	"time"
)

func main() {
	config := internal.ParseCliParams()
	if len(config.FileName) < 1 {
		panic("not set filename, please exec command from key -f , example ./main -f ./index.html")
	}
	fileContent := lib.ReadFile(config.FileName)
	version := strconv.FormatInt(time.Now().Unix(), 10)
	jsVersion1 := ".js?v=" + version + "\""
	cssVersion1 := ".css?v=" + version + "\""
	jsVersion2 := ".js?v=" + version + "'"
	cssVersion2 := ".css?v=" + version + "'"
	newFileContent := strings.ReplaceAll(fileContent, ".js\"", jsVersion1)
	newFileContent = strings.ReplaceAll(newFileContent, ".css\"", cssVersion1)
	newFileContent = strings.ReplaceAll(newFileContent, ".js'", jsVersion2)
	newFileContent = strings.ReplaceAll(newFileContent, ".css'", cssVersion2)
	lib.WriteFile(newFileContent, config.OutputFilename)
	println("version " + version + " it has been added to the including of js and css files in " + config.OutputFilename)

	println(config.FileName)
}
