package internal

import "flag"

type config struct {
	FileName       string
	OutputFilename string
}

func ParseCliParams() config {
	var filename string
	var outputFilename string

	flag.StringVar(&filename, "f", "", "")
	flag.StringVar(&outputFilename, "o", "", "")
	flag.Parse()

	config := config{FileName: filename}
	if len(outputFilename) < 1 {
		config.OutputFilename = filename
	} else {
		config.OutputFilename = outputFilename
	}
	return config
}
