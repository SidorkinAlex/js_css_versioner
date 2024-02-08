package lib

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func WriteFile(mes string, filename string) {

	if !HasFile(filename) {

		file, err := os.Create(filename)

		if err != nil {
			log.Println("Ошибка создания файла " + filename)
			os.Exit(1)
		}
		defer file.Close()
		file.WriteString(mes)
	} else {
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatalln("data not write file" + filename)
		}
		defer file.Close()
		file.WriteString(mes)
	}

}

func AppentdToFile(mes string, filename string) {

	if !HasFile(filename) {

		file, err := os.Create(filename)

		if err != nil {
			log.Println("Ошибка создания файла")
			os.Exit(1)
		}
		defer file.Close()
		file.WriteString(mes)
	} else {
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			log.Fatalln("servis-stoped test.txt")
		}
		defer file.Close()
		file.WriteString(mes)
	}

}

func RewriteFile(mes string, filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(f, mes)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	f.Sync()
}

func HasFile(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func ConcatString(string []string) string {

	strings := string
	buffer := bytes.Buffer{}
	for _, val := range strings {
		buffer.WriteString(val)
	}

	return buffer.String()
}

func GetRootDir() string {
	rootDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return rootDir
}

func ReadFile(filename string) string {
	var fileContent string
	if HasFile(filename) {
		// для более мелких файлов
		fContent, err := ioutil.ReadFile(filename)
		if err != nil {
			panic("file is not readable")
		}
		fileContent = string(fContent)
	} else {
		panic("file " + filename + " not exist")
	}

	return fileContent
}

func RemoveFile(filename string) {
	// Removing file from the directory
	// Using Remove() function
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}
