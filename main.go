package main

import (
	"flag"
	"time"
)

var (
	inputDir  string
	outputDir string
)

func main() {
	var run bool
	var serve bool
	var clean bool

	flag.BoolVar(&run, "run", true, "Generate the website from the source files")
	flag.BoolVar(&clean, "clean", false, "Remove the previous build before generating the site")
	flag.BoolVar(&serve, "serve", false, "Run a localhost server to view the site")
	flag.StringVar(&inputDir, "input", ".", "Path to the input directory")
	flag.StringVar(&outputDir, "output", "./www", "Path to save the output files")
	flag.Parse()

	l := NewFileLogger("log.log", time.RFC3339)

	if clean {
		// code to delete output Directory and files
		run = true
	}

	if run {
		Run(l)
	}

	if serve {
		Serve(l)
	}

}
