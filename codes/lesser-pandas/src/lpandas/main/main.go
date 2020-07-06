package main

import (
	"flag"
	"strings"
	"fmt"
	"strconv"


	"lpandas/core"
)

// Session manages the command execution.
type Session struct {
	Args map[string]string
}

var (
	argFile = flag.String("file", "", "file to to be read")
	argFormat = flag.String("format", "", "stdout format : csv|pretty")
	argMethod = flag.String("method", "", "method to be executed")
)

// ParseArgs parse and store command line arguments
func (sess *Session) ParseArgs() {
	sess.Args = map[string]string{}
	flag.Parse()
	sess.Args["file"] = *argFile
	sess.Args["method"] = *argMethod
	sess.Args["format"] = *argFormat

}

func main() {

	sess := Session{}
	sess.ParseArgs()


	df := core.DataFrame{}
	df.ReadCsv(sess.Args["file"])


	methods := strings.Split(sess.Args["method"], " ")
	for _, method := range methods {
		ExecMethod(df, method, sess.Args["format"])
	}

}

// ExecMethod execute method and stdout its result.
func ExecMethod(df core.DataFrame, method, format string) {
	args := strings.Split(method, ",")
	method = args[0]
	
	switch method {
	case "info":
		ret := df.Info()
		ret.Display(format)
	case "describe":
		ret := df.Describe()
		ret.Display(format)
	case "count":
		ret := df.Count()
		ret.Display(format)

	case "sum":
		ret := df.Sum()
		ret.Display(format)
	case "mean":
		ret := df.Mean()
		ret.Display(format)
	case "min":
		ret := df.Min()
		ret.Display(format)

	case "max":
		ret := df.Max()
		ret.Display(format)

	case "std":
		var nMinus1 bool
		var err error
		if len(args) == 2 {
			nMinus1, err = strconv.ParseBool(args[1])
			if err != nil {
				panic(err)
			}
		} else {
			nMinus1 = true
		}
		ret := df.Std(nMinus1)
		ret.Display(format)
	case "percentile":
		var location float64
		var err error
		if len(args) >= 2 {
			location, err = strconv.ParseFloat(strings.Join(args[1:], "."), 64)
			if err != nil {
				panic(err)
			}
		} else {
			location = 0.5
		}
		ret := df.Percentile(location)
		ret.Display(format)

	}

	fmt.Println("")
}


// Main is for testing main function.
func Main() {
	main()
}

// func main() {
// 	filePath := "../test_datas/titanic-sample.csv"
// 	df := core.DataFrame{}
// 	df.ReadCsv(filePath)

// 	df.Display("pretty")
// }
