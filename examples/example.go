package main

import (
	"os"

	"github.com/jaytarang92/goget"
)

//struct for the arguments
type ArgVals struct {
	url          string
	download_out string
}

var info ArgVals

func main() {
	info.url = os.Args[1]
	info.download_out = os.Args[2]
	example2()
}

func example1() {
	//Download source/target
	download_source := gogetter.GoTo(info.url)
	// save the target
	gogetter.SaveIt(download_source, info.download_out)
}

//This is a oneliner :) after the arguments are passed in
func example2() {
	//Call func SaveIt with GoTo as a Argument
	gogetter.SaveIt(gogetter.GoTo(info.url), info.download_out)
}
