package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jaytarang92/goget"
)

// struct for the cmdline arguments
//other wise the variables defined in main() will not be read
type cmdline struct {
	url    string
	dst    string
	md5    string
	sha256 string
}

// makes the cmdline arguments in scope
var arg cmdline

// gogetter -url=http://www.7-zip.org/a/7z1602-src.7z -dst=7zip.7z -md5=8523200928a577cd1747d8575c4ecacf
// gogetter -url=https://storage.googleapis.com/golang/go1.7.src.tar.gz -dst=go.src.tar.gz -sha256=72680c16ba0891fcf2ccf46d0f809e4ecf47bbf889f5d884ccb54c5e9a17e1c0
func main() {
	// arguments accepted
	md5Check := flag.String("md5", "", "If md5 is passed it will check it before saving file.")
	sha256chk := flag.String("sha256", "", "Same a md5 but using sha256")
	url := flag.String("url", "", "url to file")
	dest := flag.String("dst", "", "the output filename")
	flag.Parse()
	//define values to the arg struct(cmdline)
	arg.url = *url
	arg.dst = *dest
	arg.md5 = *md5Check
	arg.sha256 = *sha256chk
	// use Sprintf to format to %x to a %s
	//if both md5 and sha256 are not empty print message
	if fmt.Sprintf("%s", arg.md5) != "" && fmt.Sprintf("%s", arg.sha256) != "" {
		fmt.Println("\tPlesae choose 1. Either sha256 or md5.\n")
		os.Exit(1)
	}
	//if md5 is not empty
	if arg.md5 != "" {
		if len(arg.md5) != 32 {
			fmt.Printf("\tPlease pass a valid md5 value!\tLength:%d\n", len(arg.md5))
			os.Exit(1)
		}
		md5only()
		//if sha256 not empty
	} else if arg.sha256 != "" {
		if len(arg.sha256) != 64 {
			fmt.Printf("\tPlease pass a valid sha256 value!\tLength:%d\n", len(arg.sha256))
			os.Exit(1)
		}
		sha256only()
		//if no checksum is passed
	} else {
		none()
	}
}

// just downloads it
func none() {
	gogetter.SaveIt(gogetter.GoTo(arg.url), arg.dst)
}

// needs a md5sum to be passed for verification
func md5only() {
	hash := fmt.Sprintf("%s", arg.md5)
	input := gogetter.GoTo(arg.url)
	gogetter.SaveIt(gogetter.HashCheck(input, hash, gogetter.Hash2str(input, hash, "md5")), arg.dst)
}

// needs a sha256 passed for verification
func sha256only() {
	hash := fmt.Sprintf("%s", arg.sha256)
	input := gogetter.GoTo(arg.url)
	gogetter.SaveIt(gogetter.HashCheck(input, hash, gogetter.Hash2str(input, hash, "sha256")), arg.dst)
}
