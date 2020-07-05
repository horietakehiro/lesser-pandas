package core_test

import (
	"lpandas/helper"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
	// "github.com/golang/example/stringutil"

	"lpandas/core"
)
var correctColumns = []string{"PassengerId", "Survived", "Pclass", "Name",
								"Sex", "Age", "SibSp", "Parch", 
								"Ticket", "Fare", "Cabin", "Embarked"}


var correctCount = helper.NumpythonicFloatArray{
	891, 891, 891, 891,
	891, 714, 891, 891,
	891, 891, 204, 889,
}
var correctMean = helper.NumpythonicFloatArray{
	446, 0.384, 2.309, math.NaN(),
	math.NaN(), 29.699, 0.523, 0.382,
	math.NaN(), 32.204, math.NaN(), math.NaN(),
}
var correctStd = helper.NumpythonicFloatArray{
	257.209, 0.486, 0.836, math.NaN(),
	math.NaN(), 14.516, 1.102, 0.806,
	math.NaN(), 49.666, math.NaN(), math.NaN(),
}

var correctMin = helper.NumpythonicFloatArray{
	1.000, 0.000, 1.000, math.NaN(),
	math.NaN(), 0.420, 0.000, 0.000,
	math.NaN(), 0.000, math.NaN(), math.NaN(),
}

var correctFirstQuetaile = helper.NumpythonicFloatArray{
	223.500, 0.000, 2.000, math.NaN(),
	math.NaN(), 20.125, 0.000, 0.000,
	math.NaN(), 7.910, math.NaN(), math.NaN(),
}
var correctMedian = helper.NumpythonicFloatArray{
	446.000, 0.000, 3.000, math.NaN(),
	math.NaN(), 28.000, 0.000, 0.000,
	math.NaN(), 14.454, math.NaN(), math.NaN(),
}
var correctThirdQuatile = helper.NumpythonicFloatArray{
	668.500, 1.000, 3.000, math.NaN(),
	math.NaN(), 38.000, 1.000, 0.000,
	math.NaN(), 31.000, math.NaN(), math.NaN(),
}
var correctMax = helper.NumpythonicFloatArray{
	891.000, 1.000, 3.000, math.NaN(),
	math.NaN(), 80.000, 8.000, 6.000,
	math.NaN(), 512.329, math.NaN(), math.NaN(),
}

var correctSum = helper.NumpythonicFloatArray{
	397386.000, 342.000, 2057.000, math.NaN(),
	math.NaN(), 21205.170, 466.000, 340.000,
	math.NaN(), 28693.949, math.NaN(), math.NaN(),
}

var correctUnique = helper.NumpythonicFloatArray{
	891.000, 2.000, 3.000, 891,
	2, 88.000, 7.000, 7.000, 
	681, 248.000, 147, 3,
}
var correctFreq = helper.NumpythonicFloatArray{
	1.000, 549.000, 491.000, 1,
	577, 30.000, 608.000, 678.000, 
	7, 43.000, 4, 644,
}
var correctTop = helper.NumpythonicStringArray{
	"1.000", "0.000", "3.000", "Abbing, Mr. Anthony",
	"male", "24.000", "0.000", "0.000", 
	"1601", "8.050", "B96 B98", "S",
}

const filePath = "../test_datas/titanic-sample.csv"


func TestDataFrame_ReadCsv(t *testing.T)  {
	df := PrepareDataFrame4Test()

	columns := correctColumns
	shape := []int{891, 12}
	index := make([]string, shape[0])
	for i := 0; i < shape[0]; i++ {
		index[i] = fmt.Sprintf("%d", i)
	}
	assert.ElementsMatch(t, columns, df.Columns)
	assert.ElementsMatch(t, index, df.Index)
	assert.ElementsMatch(t, shape, df.Shape)
	for _, col := range columns {
		if df.Values[col].Dtype == "float64" {
			values, _ := df.Values[col].Values.(helper.NumpythonicFloatArray)
			assert.Equal(t, shape[0], len(values))
		} else if df.Values[col].Dtype == "string" {
			values, _ := df.Values[col].Values.(helper.NumpythonicStringArray)
			assert.Equal(t, shape[0], len(values))
		}
	}

}

func TestDataFrame_Count(t *testing.T) {
	df := PrepareDataFrame4Test()

	sr := df.Count()

	assert.Equal(t, "count", sr.Name)
	assert.ElementsMatch(t, correctColumns, sr.Index)

	retValues, _ := sr.Values.(helper.NumpythonicFloatArray)
	for i, val := range correctCount {
		retVal := retValues[i]
		if math.IsNaN(val) {
			assert.True(t, math.IsNaN(retVal))
		} else {
			assert.Equal(t, 
				fmt.Sprintf("%.3f", val), 
				fmt.Sprintf("%.3f", retVal))
		}
	}
}


func TestDataFrame_Mean(t *testing.T) {
	df := PrepareDataFrame4Test()

	sr := df.Mean()

	assert.Equal(t, "mean", sr.Name)
	assert.ElementsMatch(t, correctColumns, sr.Index)

	retValues, _ := sr.Values.(helper.NumpythonicFloatArray)
	for i, val := range correctMean {
		retVal := retValues[i]
		if math.IsNaN(val) {
			assert.True(t, math.IsNaN(retVal))
		} else {
			assert.Equal(t, 
				fmt.Sprintf("%.3f", val), 
				fmt.Sprintf("%.3f", retVal))
		}
	}
}


func TestDataFrame_Std(t *testing.T) {
	df := PrepareDataFrame4Test()

	nMinus1 := false
	sr := df.Std(nMinus1)

	assert.Equal(t, "std", sr.Name)
	assert.ElementsMatch(t, correctColumns, sr.Index)

	retValues, _ := sr.Values.(helper.NumpythonicFloatArray)
	for i, val := range correctStd {
		retVal := retValues[i]
		if math.IsNaN(val) {
			assert.True(t, math.IsNaN(retVal))
		} else {
			assert.Equal(t, 
				fmt.Sprintf("%.3f", val), 
				fmt.Sprintf("%.3f", retVal))
		}
	}
}



func TestDataFrame_Min(t *testing.T) {
	df := PrepareDataFrame4Test()

	sr := df.Min()

	assert.Equal(t, "min", sr.Name)
	assert.ElementsMatch(t, correctColumns, sr.Index)

	retValues, _ := sr.Values.(helper.NumpythonicFloatArray)
	for i, val := range correctMin {
		retVal := retValues[i]
		if math.IsNaN(val) {
			assert.True(t, math.IsNaN(retVal))
		} else {
			assert.Equal(t, 
				fmt.Sprintf("%.3f", val), 
				fmt.Sprintf("%.3f", retVal))
		}
	}
}


func TestDataFrame_Median(t *testing.T) {
	df := PrepareDataFrame4Test()

	location := 0.5
	sr := df.Percentile(location)

	assert.Equal(t, "50.0%", sr.Name)
	assert.ElementsMatch(t, correctColumns, sr.Index)

	retValues, _ := sr.Values.(helper.NumpythonicFloatArray)
	for i, val := range correctMedian {
		retVal := retValues[i]
		if math.IsNaN(val) {
			assert.True(t, math.IsNaN(retVal))
		} else {
			assert.Equal(t, 
				fmt.Sprintf("%.3f", val), 
				fmt.Sprintf("%.3f", retVal))
		}
	}
}

func TestDataFrame_Max(t *testing.T) {
	df := PrepareDataFrame4Test()

	sr := df.Max()

	assert.Equal(t, "max", sr.Name)
	assert.ElementsMatch(t, correctColumns, sr.Index)

	retValues, _ := sr.Values.(helper.NumpythonicFloatArray)
	for i, val := range correctMax {
		retVal := retValues[i]
		if math.IsNaN(val) {
			assert.True(t, math.IsNaN(retVal))
		} else {
			assert.Equal(t, 
				fmt.Sprintf("%.3f", val), 
				fmt.Sprintf("%.3f", retVal))
		}
	}
}

func TestDataFrame_Sum(t *testing.T) {
	df := PrepareDataFrame4Test()

	sr := df.Sum()

	assert.Equal(t, "sum", sr.Name)
	assert.ElementsMatch(t, correctColumns, sr.Index)

	retValues, _ := sr.Values.(helper.NumpythonicFloatArray)
	for i, val := range correctSum {
		retVal := retValues[i]
		if math.IsNaN(val) {
			assert.True(t, math.IsNaN(retVal))
		} else {
			assert.Equal(t, 
				fmt.Sprintf("%.3f", val), 
				fmt.Sprintf("%.3f", retVal))
		}
	}
}


func TestDataFrame_Unique(t *testing.T) {
	df := PrepareDataFrame4Test()

	sr := df.Unique()

	assert.Equal(t, "unique", sr.Name)
	assert.ElementsMatch(t, correctColumns, sr.Index)

	retValues, _ := sr.Values.(helper.NumpythonicFloatArray)
	for i, val := range correctUnique {
		retVal := retValues[i]
		if math.IsNaN(val) {
			assert.True(t, math.IsNaN(retVal))
		} else {
			assert.Equal(t, 
				fmt.Sprintf("%.3f", val), 
				fmt.Sprintf("%.3f", retVal))
		}
	}
}


func TestDataFrame_Freq(t *testing.T) {
	df := PrepareDataFrame4Test()

	sr := df.Freq()

	assert.Equal(t, "freq", sr.Name)
	assert.ElementsMatch(t, correctColumns, sr.Index)

	retValues, _ := sr.Values.(helper.NumpythonicFloatArray)

	for i, val := range correctFreq {
		retVal := retValues[i]
		if math.IsNaN(val) {
			assert.True(t, math.IsNaN(retVal))
		} else {
			assert.Equal(t, 
				fmt.Sprintf("%.3f", val), 
				fmt.Sprintf("%.3f", retVal))
		}
	}
}


func TestDataFrame_Top(t *testing.T) {
	df := PrepareDataFrame4Test()

	sr := df.Top()

	assert.Equal(t, "top", sr.Name)
	assert.ElementsMatch(t, correctColumns, sr.Index)

	retValues, _ := sr.Values.(helper.NumpythonicStringArray)

	for i, val := range correctTop {
		retVal := retValues[i]
		assert.Equal(t, val, retVal)
	}
}



func TestDataFrame_Info(t *testing.T) {
	columns := []string{"non-null", "null", "dtype"}
	index := correctColumns
	shape := []int{3, 12}
	correctNonNull := helper.NumpythonicFloatArray{
		float64(891), float64(891), float64(891), float64(891),
		float64(891), float64(714), float64(891), float64(891), 
		float64(891), float64(891), float64(204), float64(889),
	}
	correctNull := helper.NumpythonicFloatArray{
		float64(0), float64(0), float64(0), float64(0),
		float64(0), float64(891 - 714), float64(0), float64(0), 
		float64(0), float64(0), float64(891 - 204), float64(891 - 889),
	}
	correctDtype := helper.NumpythonicStringArray{
		"float64", "float64", "float64", "string",
		"string", "float64", "float64", "float64", 
		"string", "float64", "string", "string",
	}

	df := PrepareDataFrame4Test()
	info := df.Info()

	assert.ElementsMatch(t, columns, info.Columns)
	assert.ElementsMatch(t, index, info.Index)
	assert.ElementsMatch(t, shape, info.Shape)

	for _, col := range info.Columns {
		switch col {
		case "non-null":
			retValues, _ := info.Values[col].Values.(helper.NumpythonicFloatArray)
			assert.ElementsMatch(t, correctNonNull, retValues)
		case "null":
			retValues, _ := info.Values[col].Values.(helper.NumpythonicFloatArray)
			assert.ElementsMatch(t, correctNull, retValues)
		case "dtype":
			retValues, _ := info.Values[col].Values.(helper.NumpythonicStringArray)
			assert.ElementsMatch(t, correctDtype, retValues)
		default:
			assert.Fail(t, "Invalid columns is in the returned DataFrame")
		}
	}
}

func TestDataFrame_Describe(t *testing.T) {
	columns := []string{"count", "mean", "std", "min",
						 "25.0%", "50.0%", "75.0%", "max",
						  "sum", "unique", "freq", "top"}
	index := correctColumns
	shape := []int{12, 12}

	df := PrepareDataFrame4Test()
	describe := df.Describe()

	assert.ElementsMatch(t, columns, describe.Columns)
	assert.ElementsMatch(t, index , describe.Index)
	assert.ElementsMatch(t, shape, describe.Shape)

	for _, col := range describe.Columns {
		switch col {
		case "count":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctCount, retValues)
		case "mean":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctMean, retValues)
		case "std":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctStd, retValues)
		case "min":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctMin, retValues)
		case "25.0%":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctFirstQuetaile, retValues)
		case "50.0%":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctMedian, retValues)
		case "75.0%":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctThirdQuatile, retValues)
		case "max":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctMax, retValues)
		case "sum":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctSum, retValues)
		case "unique":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctUnique, retValues)
		case "freq":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicFloatArray)
			assertFloatArray(t, correctFreq, retValues)
		case "top":
			retValues, _ := describe.Values[col].Values.(helper.NumpythonicStringArray)
			assertStringArray(t, correctTop, retValues)
		default:
			assert.Fail(t, "Invalid columns is in the returned DataFrame")
		}

	}
}


func ExampleDataFrame_Display_csv() {
	df := PrepareDataFrame4Test()
	format := "csv"
	df.Info().Display(format)
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

}





func assertFloatArray(t *testing.T, correct, retValues helper.NumpythonicFloatArray) {
	for i, val := range retValues {
		if math.IsNaN(correct[i]) {
			assert.True(t, math.IsNaN(val))
		} else {
			assert.Equal(t, fmt.Sprintf("%.3f", correct[i]), fmt.Sprintf("%.3f", val))
		}
	}
}

func assertStringArray(t *testing.T, correct, retValues helper.NumpythonicStringArray) {
	for i, val := range retValues {
		assert.Equal(t, correct[i], val)
	}
}


func PrepareDataFrame4Test() core.DataFrame {

	csvData := core.DataFrame{}
	csvData.ReadCsv(filePath)
	return csvData
}
