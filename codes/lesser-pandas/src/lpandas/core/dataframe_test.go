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

const filePath = "../test_datas/titanic-sample.csv"
var correctColumns = []string{"PassengerId", "Survived", "Pclass", "Name",
								"Sex", "Age", "SibSp", "Parch", 
								"Ticket", "Fare", "Cabin", "Embarked"}



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


// func ExampleDataFrame_Info() {
// 	csvData := PrepareDataFrame4Test()
	
// 	csvData.Info()
// 	// Output:
// 	// RangeIndex: 891 entries, 0 to 890
// 	// Data columns (total 12 columns):
// 	// ===== Numeric columns (total 7 columns) =====
// 	// name,non-null,null,dtype
// 	// PassengerId,891,0,float64
// 	// Survived,891,0,float64
// 	// Pclass,891,0,float64
// 	// Age,714,177,float64
// 	// SibSp,891,0,float64
// 	// Parch,891,0,float64
// 	// Fare,891,0,float64
// 	//
// 	// ===== String columns (total 5 columns) =====
// 	// name,non-null,null,dtype
// 	// Name,891,0,string
// 	// Sex,891,0,string
// 	// Ticket,891,0,string
// 	// Cabin,204,687,string
// 	// Embarked,889,2,string
	

// }

// func ExampleDataFrame_Describe() {
// 	csvData := PrepareDataFrame4Test()
// 	csvData.Describe()
// 	// Output:
// 	// === Numeric columns (total 7 columns) =====
// 	// metric,PassengerId,Survived,Pclass,Age,SibSp,Parch,Fare
// 	// count,891.000,891.000,891.000,714.000,891.000,891.000,891.000
// 	// mean,446.000,0.384,2.309,29.699,0.523,0.382,32.204
// 	// std,257.209,0.486,0.836,14.516,1.102,0.806,49.666
// 	// min,1.000,0.000,1.000,0.420,0.000,0.000,0.000   
// 	// 25%,223.500,0.000,2.000,20.125,0.000,0.000,7.910
// 	// 50%,446.000,0.000,3.000,28.000,0.000,0.000,14.454
// 	// 75%,668.500,1.000,3.000,38.000,1.000,0.000,31.000
// 	// max,891.000,1.000,3.000,80.000,8.000,6.000,512.329
// 	// sum,397386.000,342.000,2057.000,21205.170,466.000,340.000,28693.949
// 	//
// 	// === String columns (total 5 columns) =====
// 	// metric,Name,Sex,Ticket,Cabin,Embarked   
// 	// count,891,891,891,204,889
// 	// unique,891,2,681,147,3  
// 	// freq,1,577,7,4,644
// 	// top,Abbing, Mr. Anthony,male,1601,B96 B98,S

// }

func TestDataFrame_Sum(t *testing.T) {
	df := PrepareDataFrame4Test()
	correctSum := helper.NumpythonicFloatArray{
		float64(397386.000), float64(342.000), float64(2057.000), math.NaN(),
		math.NaN(), float64(21205.170), float64(466.000), float64(340.000), 
		math.NaN(), float64(28693.949), math.NaN(), math.NaN(),
	}
	retSum := df.Sum()

	assert.Equal(t, "sum", retSum.Name)
	assert.ElementsMatch(t, correctColumns, retSum.Index)

	retValues, _ := retSum.Values.(helper.NumpythonicFloatArray)
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


func PrepareDataFrame4Test() core.DataFrame {

	csvData := core.DataFrame{}
	csvData.ReadCsv(filePath)
	return csvData
}