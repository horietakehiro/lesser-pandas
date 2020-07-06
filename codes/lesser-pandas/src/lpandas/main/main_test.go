package main_test

import (
	"flag"
	// "strings"

	"testing"
	"github.com/stretchr/testify/assert"

	// "lpandas/core"
	"lpandas/main"
)

const filePath = "../test_datas/titanic-sample.csv"

func TestParseArgs(t *testing.T) {
	args := map[string]string{
		"file" : filePath,
		"method" : "describeinfo info ",
		"format" : "csv",
	}
	for k, v := range args {
		flag.CommandLine.Set(k, v)
	}

	sess := main.Session{}
	sess.ParseArgs()

	for k, v := range args {
		assert.Equal(t, v, sess.Args[k])
	}

}

func ExampleExecMethod() {
	args := map[string]string{
		"file" : filePath,
		"method" : "info count",
		// "method" : "info",
		"format" : "csv",
	}
	for k, v := range args {
		flag.CommandLine.Set(k, v)
	}

	// sess := main.Session{}
	// sess.ParseArgs()
	// df := core.DataFrame{}
	// df.ReadCsv(sess.Args["file"])
	// for _, method := range strings.Split(sess.Args["method"], " ") {
	// 	main.ExecMethod(df, method, sess.Args["format"])
	// }
	main.Main()
	// output:
	// index,non-null,null,dtype
	// PassengerId,891.000,0.000,float64
	// Survived,891.000,0.000,float64
	// Pclass,891.000,0.000,float64
	// Name,891.000,0.000,string
	// Sex,891.000,0.000,string
	// Age,714.000,177.000,float64
	// SibSp,891.000,0.000,float64
	// Parch,891.000,0.000,float64
	// Ticket,891.000,0.000,string
	// Fare,891.000,0.000,float64
	// Cabin,204.000,687.000,string
	// Embarked,889.000,2.000,string
	//
	// index,count
	// PassengerId,891.000
	// Survived,891.000
	// Pclass,891.000
	// Name,891.000
	// Sex,891.000
	// Age,714.000
	// SibSp,891.000
	// Parch,891.000
	// Ticket,891.000
	// Fare,891.000
	// Cabin,204.000
	// Embarked,889.000

}
