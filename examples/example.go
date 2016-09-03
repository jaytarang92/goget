package main

import (
	"flag"

	"github.com/jaytarang92/goget"
)

// gogetter -url=http://google.com -dst=/Users/J_Hack92/Desktop/index.html -md5sum=1234567890123456789

func main() {
	// arguments accepted
	md5Check := flag.String("md5sum", "", "If md5 is passed it will check it before saving file.")
	url := flag.String("url", "", "url to file")
	dest := flag.String("dst", "", "the output filename")
	flag.Parse()

	//run gogetter to save the file
	gogetter.SaveIt(gogetter.GoTo(*url), *dest, *md5Check)
}
