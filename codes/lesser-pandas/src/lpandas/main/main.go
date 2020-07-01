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
	args := strings.Split(method, ".")
	method = args[0]
	switch method {
	case "info":
		fmt.Println("========== df.Info() ==========")
		df.Info()
	case "describe":
		fmt.Println("========== df.Describe() ==========")
		df.Describe()
	case "sum":
		fmt.Println("========== df.Sum() ==========")
		df.Sum()
	case "mean":
		fmt.Println("========== df.Mean() ==========")
		df.Mean()
	case "min":
		fmt.Println("========== df.Min() ==========")
		df.Min()
	case "max":
		fmt.Println("========== df.Max() ==========")
		df.Max()
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
		fmt.Printf("========== df.Std(%s) ==========\n", strconv.FormatBool(nMinus1))
		df.Std(nMinus1)
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
		fmt.Printf("========== df.Percentile(%.3f) ==========\n", location)
		df.Percentile(location)
	}



	fmt.Println("")
}


// Main is for testing main function.
func Main() {
	main()
}
