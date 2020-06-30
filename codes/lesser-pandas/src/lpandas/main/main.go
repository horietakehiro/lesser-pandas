package main

import (
	"flag"
	"strings"
	"fmt"

	"lpandas/core"
)

// Session manages the command execution.
type Session struct {
	Args map[string]string
}

var (
	argFile = flag.String("file", "", "file to to be read")
	argMethod = flag.String("method", "info", "method to be executed")
)

// ParseArgs parse and store command line arguments
func (sess *Session) ParseArgs() {
	sess.Args = map[string]string{}
	flag.Parse()
	sess.Args["file"] = *argFile
	sess.Args["method"] = *argMethod
}

func main() {

	sess := Session{}
	sess.ParseArgs()

	df := core.DataFrame{}
	df.ReadCsv(sess.Args["file"])

	for _, method := range strings.Split(sess.Args["method"], ",") {
		execMethod(df, method)
	}

}

func execMethod(df core.DataFrame, method string) {
	switch method {
	case "info":
		fmt.Println("========== df.Info() ==========")
		df.Info()
	case "describe":
		fmt.Println("========== df.Describe() ==========")
		df.Describe()
	}

	fmt.Println("")
}


// Main is for testing main function.
func Main() {
	main()
}
