package main

import (
	"os"

	"github.com/jaytarang92/goget"
)

func main() {
	// Download output file
	file_out := gogetter.CreateDownload(os.Args[2])
	//Download source/target
	download_source := gogetter.GoTo(os.Args[1])
	// save the target
	gogetter.SaveIt(download_source, file_out)
}
