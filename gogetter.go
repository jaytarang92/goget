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

package gogetter

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/parnurzeal/gorequest"
)

//GoTo gorequest to open url and get the response and return it
//proxyTunnel is just a proxy url being passed in
func GoTo(goURL string, proxyTunnel string) string {
	fmt.Println("\tReading input...")
	_, body, _ := gorequest.New().
		Proxy(proxyTunnel).
		Get(goURL).
		End()
	return body
}

//Hash2str eturns hash as string
func Hash2str(bodyresponse string, checksum string, checktype string) string {
	var strhash string
	switch checktype {
	case "md5":
		byteit := []byte(bodyresponse)
		hash := md5.Sum(byteit)
		sthash := fmt.Sprintf("%x", hash)
		strhash = sthash
		/* ADDED TO MAIN FOR FASTER RESPONSE | BEFORE STARTING DOWNLOAD
		if len(strhash) != 32 {
			fmt.Println("\tA valid md5 was not passed please try again!")
			os.Exit(1)
		}*/
	case "sha256":
		fmt.Printf("\tCase sha256\n")
		byteit := []byte(bodyresponse)
		hash := sha256.Sum256(byteit)
		sthash := fmt.Sprintf("%x", hash)
		strhash = sthash
		/* ADDED TO MAIN FOR FASTER RESPONSE | BEFORE STARTING DOWNLOAD
		if len(strhash) != 64 {
			fmt.Println("\tA valid sha256 was not passed please try again!")
			os.Exit(1)
		}
		*/
	}
	return strhash
}
func HashCheck(bodyresponse string, checksum string, strhash string) string {
	//compare file hash vs argument passed in
	if checksum == strhash {
		fmt.Printf("\tChecksums Match!\n\tHash:\t%s", checksum)
	} else {
		fmt.Printf("\tChecksums don't match exiting!\n")
		fmt.Printf("\tPassed:\t%s\n\tActual:\t%x\n", checksum, strhash)
		os.Exit(1)
	}
	return bodyresponse
}

//SaveIt writes the response from GoTo to a file (download_out)
func SaveIt(bodyresponse string, downloadOut string) {
	fmt.Printf("\n\tSaving to %s ...\n", downloadOut)
	byteit := []byte(bodyresponse)
	// writes the output file with read/write permissions
	ioutil.WriteFile(downloadOut, byteit, 0644)
}
