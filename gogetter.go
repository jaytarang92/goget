package gogetter

import (
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
)

//gorequest to open url and get the response and return it
func GoTo(goURL string) string {
	_, body, _ := gorequest.New().Get(goURL).End()
	return body
}

//writes the response from GoTO to a file that was declared in CreateDownload
func SaveIt(bodyresponse string, download_out string) {
	//convert to bytes or ioutil complains
	byteit := []byte(bodyresponse)
	// writes the output file with read/write permissions
	ioutil.WriteFile(download_out, byteit, 0644)
}
