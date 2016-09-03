package gogetter

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/parnurzeal/gorequest"
)

//GoTo gorequest to open url and get the response and return it
func GoTo(goURL string) string {
	fmt.Println("\tReading input...")
	_, body, _ := gorequest.New().Get(goURL).End()
	return body
}

//HashCheck checks hashes
func HashCheck(bodyresponse string, checksum string) string {
	byteit := []byte(bodyresponse)
	if len(checksum) != 0 {
		//assign to hash to output as string
		hash := md5.Sum(byteit)
		strhash := fmt.Sprintf("%x", hash)
		//compare md5
		if checksum == strhash {
			fmt.Printf("\tChecksums Match!\n\tHash:\t%s", checksum)
		} else {
			fmt.Printf("\tChecksums don't match exiting!\n")
			fmt.Printf("\tPassed:\t%s\n\tActual:\t%x\n", checksum, hash)
			os.Exit(1)
		}
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
