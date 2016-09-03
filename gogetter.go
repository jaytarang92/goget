package gogetter

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
)

//gorequest to open url and get the response and return it
func GoTo(goURL string) string {
	_, body, _ := gorequest.New().Get(goURL).End()
	return body
}

//writes the response from GoTO to a file (download_out)
func SaveIt(bodyresponse string, download_out string, checksum bool) {
	//convert to bytes or ioutil complains
	byteit := []byte(bodyresponse)
	if checksum {
		fmt.Printf("\nmd5:\t%x\n\n", md5.Sum(byteit))
	}
	// writes the output file with read/write permissions
	ioutil.WriteFile(download_out, byteit, 0644)
}
