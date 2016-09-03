package main

import (
	"flag"

	"github.com/jaytarang92/goget"
)

// gogetter -url=http://google.com -dst=/Users/J_Hack92/Desktop/index.html -md5sum=true

func main() {
	// arguments accepted
	md5Check := flag.Bool("md5sum", false, "Set to true if want to display md5")
	url := flag.String("url", "", "url to file")
	dest := flag.String("dst", "", "the output filename")
	flag.Parse()

	//run gogetter to save the file
	gogetter.SaveIt(gogetter.GoTo(*url), *dest, *md5Check)
}
