/*
The MIT License (MIT)

Copyright (c) 2016 Jasmit Tarang

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
	proxy  string
}

// makes the cmdline arguments in scope
var arg cmdline

//define pproxy for printing
var pproxy string

// gogetter -url=http://www.7-zip.org/a/7z1602-src.7z -dst=7zip.7z -md5=8523200928a577cd1747d8575c4ecacf
// gogetter -url=https://storage.googleapis.com/golang/go1.7.src.tar.gz -dst=go.src.tar.gz -sha256=72680c16ba0891fcf2ccf46d0f809e4ecf47bbf889f5d884ccb54c5e9a17e1c0
func main() {
	// arguments accepted
	proxy := flag.String("proxy", "", "Proxy i.e http://company.proxy.com:8080")
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
	arg.proxy = *proxy

	if arg.url == "" || arg.dst == "" {
		fmt.Println("  Insufficient number of arguments passed! Please check if dst and url were passed!")
		flag.PrintDefaults()
	}

	// use Sprintf to format to %x to a %s
	//if both md5 and sha256 are not empty print message
	if fmt.Sprintf("%s", arg.md5) != "" && fmt.Sprintf("%s", arg.sha256) != "" {
		fmt.Println("\tPlesae choose 1. Either sha256 or md5.")
		os.Exit(1)
	}

	//if no proxy is defined then bypass it
	if arg.proxy == "" {
		pproxy = "proxy is undefined"
	} else {
		pproxy = arg.proxy
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
	gogetter.SaveIt(gogetter.GoTo(arg.url, pproxy), arg.dst)
}

// needs a md5sum to be passed for verification
func md5only() {
	hash := fmt.Sprintf("%s", arg.md5)
	input := gogetter.GoTo(arg.url, pproxy)
	gogetter.SaveIt(gogetter.HashCheck(input, hash, gogetter.Hash2str(input, hash, "md5")), arg.dst)
}

// needs a sha256 passed for verification
func sha256only() {
	hash := fmt.Sprintf("%s", arg.sha256)
	input := gogetter.GoTo(arg.url, pproxy)
	gogetter.SaveIt(gogetter.HashCheck(input, hash, gogetter.Hash2str(input, hash, "sha256")), arg.dst)
}
