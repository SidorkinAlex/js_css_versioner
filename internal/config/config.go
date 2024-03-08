package config

import "flag"

type Config struct {
	FileName       string
	OutputFilename string
}

func ParseCliParams() Config {
	var filename string
	var outputFilename string

	flag.StringVar(&filename, "f", "", "")
	flag.StringVar(&outputFilename, "o", "", "")
	flag.Parse()

	config := Config{FileName: filename}
	if len(outputFilename) < 1 {
		config.OutputFilename = filename
	} else {
		config.OutputFilename = outputFilename
	}
	return config
}
