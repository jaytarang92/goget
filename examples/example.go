package main

import (
	"flag"
	"fmt"

	"github.com/jaytarang92/goget"
)

// gogetter -url=http://google.com -dst=/Users/J_Hack92/Desktop/index.html -md5sum=1234567890123456789

func main() {
	// arguments accepted
	md5Check := flag.String("md5sum", "", "If md5 is passed it will check it before saving file.")
	url := flag.String("url", "", "url to file")
	dest := flag.String("dst", "", "the output filename")
	flag.Parse()
	//Sprintf formats md5 check t string
	if fmt.Sprintf("%s", *md5Check) == "" {
		////run gogetter to save the file no hash check
		gogetter.SaveIt(gogetter.GoTo(*url), *dest)
	} else {
		//run gogetter to save the file with hash check
		gogetter.SaveIt(gogetter.HashCheck(gogetter.GoTo(*url), *md5Check), *dest)
	}
}
