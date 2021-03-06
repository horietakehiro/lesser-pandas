package core_test

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
	"os"

	"lpandas/helper"
	"lpandas/core"
)


// use "PassangerId" and "Embarked" columns for this tests.
const numeric = "PassengerId"
const str = "Embarked"


func TestSeries_Count_numeric(t *testing.T) {
	correct := float64(891)
	sr := PrepareSeries4Test(numeric)
	
	assert.Equal(t, correct, sr.Count())
}

func TestSeries_Count_string(t *testing.T) {
	correct := float64(889)
	sr := PrepareSeries4Test(str)
	
	assert.Equal(t, correct, sr.Count())
}



func TestSeries_Mean_numeric(t *testing.T) {
	correct := float64(446)
	sr := PrepareSeries4Test(numeric)
	
	assert.Equal(t, correct, sr.Mean())
}

func TestSeries_Mean_string(t *testing.T) {
	// correct := math.NaN()
	sr := PrepareSeries4Test(str)
	
	// assert.Equal(t, correct, sr.Mean())
	assert.True(t, math.IsNaN(sr.Mean()))
}


func TestSeries_Std_numeric(t *testing.T) {
	correct := float64(257.209)

	nMinus1 := false
	sr := PrepareSeries4Test(numeric)
	
	assert.Equal(t, 
		fmt.Sprintf("%.3f", correct),
		fmt.Sprintf("%.3f", sr.Std(nMinus1)))
}

func TestSeries_Std_string(t *testing.T) {
	nMinus1 := false
	sr := PrepareSeries4Test(str)
	
	assert.True(t, math.IsNaN(sr.Std(nMinus1)))
}


func TestSeries_Min_numeric(t *testing.T) {
	correct := float64(1)

	sr := PrepareSeries4Test(numeric)
	
	assert.Equal(t, 
		fmt.Sprintf("%.3f", correct),
		fmt.Sprintf("%.3f", sr.Min()))
}

func TestSeries_Min_string(t *testing.T) {
	sr := PrepareSeries4Test(str)
	
	assert.True(t, math.IsNaN(sr.Min()))
}


func TestSeries_Median_numeric(t *testing.T) {
	correct := float64(446.000)
	location := 0.5

	sr := PrepareSeries4Test(numeric)
	
	assert.Equal(t, 
		fmt.Sprintf("%.3f", correct),
		fmt.Sprintf("%.3f", sr.Percentile(location)))
}

func TestSeries_Median_string(t *testing.T) {
	sr := PrepareSeries4Test(str)
	location := 0.5
	
	assert.True(t, math.IsNaN(sr.Percentile(location)))
}


func TestSeries_Max_numeric(t *testing.T) {
	correct := float64(891.000)

	sr := PrepareSeries4Test(numeric)
	
	assert.Equal(t, 
		fmt.Sprintf("%.3f", correct),
		fmt.Sprintf("%.3f", sr.Max()))
}

func TestSeries_Max_string(t *testing.T) {
	sr := PrepareSeries4Test(str)
	
	assert.True(t, math.IsNaN(sr.Max()))
}


func TestSeries_Sum_numeric(t *testing.T) {
	correct := float64(397386.000)

	sr := PrepareSeries4Test(numeric)
	
	assert.Equal(t, 
		fmt.Sprintf("%.3f", correct),
		fmt.Sprintf("%.3f", sr.Sum()))
}

func TestSeries_Sum_string(t *testing.T) {
	sr := PrepareSeries4Test(str)
	
	assert.True(t, math.IsNaN(sr.Sum()))
}



func TestSeries_Unique_numeric(t *testing.T) {
	correct := float64(891.000)

	sr := PrepareSeries4Test(numeric)
	
	assert.Equal(t, 
		fmt.Sprintf("%.3f", correct),
		fmt.Sprintf("%.3f", sr.Unique()))
}

func TestSeries_Unique_string(t *testing.T) {
	correct := float64(3)
	sr := PrepareSeries4Test(str)
	
	assert.Equal(t, 
		fmt.Sprintf("%.3f", correct),
		fmt.Sprintf("%.3f", sr.Unique()))
}


func TestSeries_Freq_numeric(t *testing.T) {
	correct := float64(1.000)

	sr := PrepareSeries4Test(numeric)
	
	assert.Equal(t, 
		fmt.Sprintf("%.3f", correct),
		fmt.Sprintf("%.3f", sr.Freq()))
}

func TestSeries_Freq_string(t *testing.T) {
	correct := float64(644)
	sr := PrepareSeries4Test(str)
	
	assert.Equal(t, 
		fmt.Sprintf("%.3f", correct),
		fmt.Sprintf("%.3f", sr.Freq()))
}


func TestSeries_Top_numeric(t *testing.T) {
	correct := "1.000"

	sr := PrepareSeries4Test(numeric)
	
	assert.Equal(t, correct, sr.Top())
}

func TestSeries_Top_string(t *testing.T) {
	correct := "S"
	sr := PrepareSeries4Test(str)
	
	assert.Equal(t, correct, sr.Top())

}



func TestSeries_Info_numeric(t *testing.T) {
	nonNull := "891.000"
	null := "0.000"
	dtype := "string"
	index := []string{"non-null", "null", "dtype"}
	name := numeric

	sr := PrepareSeries4Test(numeric)
	info := sr.Info()

	assert.Equal(t, dtype, info.Dtype)
	assert.Equal(t, name, info.Name)

	values, _ := info.Values.(helper.NumpythonicStringArray)
	for i, col := range index {
		switch col {
		case "non-null":
			assert.Equal(t, nonNull, values[i])
		case "null":
			assert.Equal(t, null, values[i])
		case "dtype":
			assert.Equal(t, dtype, values[i])
		default:
			assert.Fail(t, "Invalid index is in the returned Series")
			
		}
	}

}


func TestSeries_Info_string(t *testing.T) {
	nonNull := "889.000"
	null := "2.000"
	dtype := "string"
	index := []string{"non-null", "null", "dtype"}
	name := str


	sr := PrepareSeries4Test(str)
	info := sr.Info()
	assert.Equal(t, dtype, info.Dtype)
	assert.Equal(t, name, info.Name)

	values, _ := info.Values.(helper.NumpythonicStringArray)
	for i, col := range index {
		switch col {
		case "non-null":
			assert.Equal(t, nonNull, values[i])
		case "null":
			assert.Equal(t, null, values[i])
		case "dtype":
			assert.Equal(t, dtype, values[i])
			
		default:
			assert.Fail(t, "Invalid index is in the returned Series")
		}
	}

}


func TestSeries_Describe_numeric(t *testing.T) {
	correct := map[string]string{
		"count" : "891.000", "mean" : "446.000", "std" : "257.209", "min" : "1.000",
		"25.0%" : "223.500", "50.0%" : "446.000", "75.0%" : "668.500", "max" : "891.000",
		"sum" : "397386.000", "unique" : "891.000", "freq" : "1.000", "top" : "1.000",
	}
	
	dtype := "string"
	index := []string{
		"count", "mean", "std", "min", 
		"25.0%", "50.0%", "75.0%", "max",
		"sum", "unique", "freq", "top",}
	
	name := numeric

	sr := PrepareSeries4Test(numeric)
	describe := sr.Describe()

	assert.Equal(t, dtype, describe.Dtype)
	assert.Equal(t, name, describe.Name)

	values, _ := describe.Values.(helper.NumpythonicStringArray)
	for i, col := range index {
		assert.Equal(t, correct[col], values[i])
	}
}


func TestSeries_Describe_string(t *testing.T) {
	correct := map[string]string{
		"count" : "889.000", "mean" : "", "std" : "", "min" : "",
		"25.0%" : "", "50.0%" : "", "75.0%" : "", "max" : "",
		"sum" : "", "unique" : "3.000", "freq" : "644.000", "top" : "S",
	}
	
	dtype := "string"
	index := []string{
		"count", "mean", "std", "min", 
		"25.0%", "50.0%", "75.0%", "max",
		"sum", "unique", "freq", "top",}
	
	name := str

	sr := PrepareSeries4Test(str)
	describe := sr.Describe()

	assert.Equal(t, dtype, describe.Dtype)
	assert.Equal(t, name, describe.Name)

	values, _ := describe.Values.(helper.NumpythonicStringArray)
	for i, col := range index {
		assert.Equal(t, correct[col], values[i])
	}
}


func ExampleSeries_Display_csv() {
	sr := PrepareSeries4Test(str)
	format := "csv"
	sr.Info().Display(format)
	// output:
	// index,Embarked
	// non-null,889.000
	// null,2.000
	// dtype,string

}

func ExampleSeries_Display_pretty() {
	sr := PrepareSeries4Test(str)
	format := "pretty"
	sr.Info().Display(format)
	// output:
	// index    | Embarked |
	// non-null | 889.000  |
	// null     | 2.000    |
	// dtype    | string   |
	//

}

func TestSeries_ToCsv_withIndex(t *testing.T) {
	sr := PrepareSeries4Test(numeric)
	info := sr.Info()

	out := "./out.csv"
	withIndex := true
	info.ToCsv(out, withIndex)
	defer os.Remove(out)

	_, err := os.Stat(out)
	assert.True(t, err == nil)

	df := core.DataFrame{}
	df.ReadCsv(out)
	
	assert.ElementsMatch(t, []string{"index", "PassengerId"}, df.Columns)


}


func PrepareSeries4Test(col string) core.Series {
	filePath := "../test_datas/titanic-sample.csv"

	csvData := core.DataFrame{}
	csvData.ReadCsv(filePath)
	return *csvData.Values[col]
}
