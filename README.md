
A Aria2 library that is used in Yadownloader(A frontend for youtube-dl and Aria2)

### Usage 
start Aria2 
````
package main

import (
	"github.com/shiningw/aria2go/cmd"
)

func main() {
	opts := cmd.NewRunOptions("token123", "6800")
	opts.SetLogFile("/tmp/Aria2.log")
	opts.AddOption("--log-level=info")
	opts.SetInputFile("/tmp/session.log")
	opts.SetSessionFile("/tmp/session.log")
	cmd.StartAria2(opts)
}
````