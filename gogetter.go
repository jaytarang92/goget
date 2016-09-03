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
	_, body, _ := gorequest.New().Get(goURL).End()
	return body
}

//SaveIt writes the response from GoTo to a file (download_out)
func SaveIt(bodyresponse string, downloadOut string, checksum string) {
	//convert to bytes or ioutil complains
	byteit := []byte(bodyresponse)
	if len(checksum) != 0 {
		//assign to hash to output as string
		hash := md5.Sum(byteit)
		strhash := fmt.Sprintf("%x", hash)
		//compare md5
		if checksum == strhash {
			fmt.Printf("Checksums Match! Saving to %s\n", downloadOut)
		} else {
			fmt.Printf("\tChecksums don't match exiting!\n")
			fmt.Printf("\tPassed:\t%s\n\tActual:\t%x\n", checksum, hash)
			os.Exit(1)
		}
	}
	// writes the output file with read/write permissions
	ioutil.WriteFile(downloadOut, byteit, 0644)
}
