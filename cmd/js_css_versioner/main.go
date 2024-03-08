package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SidorkinAlex/js_css_versioner/internal/config"
	"github.com/SidorkinAlex/js_css_versioner/internal/models"
	"github.com/SidorkinAlex/js_css_versioner/internal/version_replace"
)

func main() {

	log.Println("starting js css versioner...")
	config := config.ParseCliParams()
	if len(config.FileName) == 0 {
		log.Fatalln("not set filename, please exec command from key -f , example ./main -f ./index.html")
	}

	repl, err := version_replace.New(models.VersionedExtensionCSS, models.VersionedExtensionJS)
	if err != nil {
		log.Fatalln("cannot init version replace service")
	}

	tmpFileName := fmt.Sprintf("%s.tmp", config.OutputFilename)
	if errRpl := replaceFileContent(repl, config.FileName, tmpFileName); errRpl != nil {
		log.Fatalf("replace file content failed: %s\n", errRpl)
	}

	if config.FileName == config.OutputFilename {
		if errRm := os.Remove(config.FileName); errRm != nil {
			log.Fatalf("failed to remove original file: %s\n", config.FileName)
		}
	}

	if errRn := os.Rename(tmpFileName, config.OutputFilename); errRn != nil {
		log.Fatalf("failed to rename output file: %s\n", config.OutputFilename)
	}

	log.Println("js css versioner finished")

}

func replaceFileContent(repl *version_replace.VersionReplace, srcFileName string, dstFileName string) error {

	log.Printf("open source file %s", srcFileName)
	sourceFile, errOpen := os.Open(srcFileName)
	if errOpen != nil {
		return fmt.Errorf("open file: %s failed with error: %w", srcFileName, errOpen)
	}
	defer func() {
		if errScr := sourceFile.Close(); errScr != nil {
			log.Printf("source file closed with error: %s\n", errScr)
		}
	}()

	destFile, errCreate := os.Create(dstFileName)
	if errCreate != nil {
		return fmt.Errorf("create destination file: %s failed with error: %w", dstFileName, errCreate)
	}
	defer func() {
		if errDst := destFile.Close(); errDst != nil {
			log.Printf("destination file closed with error: %s\n", errDst)
		}
	}()

	return repl.Execute(sourceFile, destFile, time.Now().Unix())
}
