package main

import (
	"flag"
	"fmt"
	"github.com/jaytarang92/goget"
	"os"
)

type cmdline struct {
	url string
	dst string
	md5 string
	sha256 string
}

var arg cmdline


// gogetter -url=http://google.com -dst=/Users/J_Hack92/Desktop/index.html -md5sum=1234567890123456789
// gogetter -url=https://storage.googleapis.com/golang/go1.7.src.tar.gz -dst=go.src -sha256=72680c16ba0891fcf2ccf46d0f809e4ecf47bbf889f5d884ccb54c5e9a17e1c0
func main() {
	// arguments accepted
	md5Check := flag.String("md5sum", "", "If md5 is passed it will check it before saving file.")
	sha256chk := flag.String("sha256", "", "Same a md5 but using sha256")
	url := flag.String("url", "", "url to file")
	dest := flag.String("dst", "", "the output filename")
	flag.Parse()
	arg.url = *url
	arg.dst = *dest
	arg.md5 = *md5Check
	arg.sha256 = *sha256chk
  if fmt.Sprintf("%s", arg.md5) != "" && fmt.Sprintf("%s", arg.sha256) != "" {
    fmt.Println("\tPleae choose 1. Either sha256 or md5.\n")
    os.Exit(1)
  }
  if arg.md5 != "" {
    md5only()
	} else if arg.sha256 != "" {
    sha256only()
  } else {
    none()
  }
}

func none() {
	gogetter.SaveIt(gogetter.GoTo(arg.url), arg.dst)
}

func md5only()  {
	hash := fmt.Sprintf("%s", arg.md5)
	input := gogetter.GoTo(arg.url)
	gogetter.SaveIt(gogetter.HashCheck(input, hash, gogetter.Hash2str(input, hash, "md5")), arg.dst)
}

func sha256only()  {
	hash := fmt.Sprintf("%s", arg.sha256)
	input := gogetter.GoTo(arg.url)
	gogetter.SaveIt(gogetter.HashCheck(input, hash, gogetter.Hash2str(input, hash, "sha256")), arg.dst)
}
