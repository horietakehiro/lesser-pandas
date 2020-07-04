package core_test

import (
	// "lpandas/helper"
	// "fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	// "math"

	"lpandas/core"
)


// use "Age" columns for this tests.


func TestSeries_Count(t *testing.T) {
	correctCount := float64(714)
	sr := PrepareSeries4Test()
	
	assert.Equal(t, correctCount, sr.Count())
}


func PrepareSeries4Test() core.Series {
	filePath := "../test_datas/titanic-sample.csv"

	csvData := core.DataFrame{}
	csvData.ReadCsv(filePath)
	return *csvData.Values["Age"]
}
