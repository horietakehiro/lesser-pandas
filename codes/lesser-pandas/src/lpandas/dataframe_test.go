package lpandas_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	// "github.com/golang/example/stringutil"

	"lpandas"
)

const filePath = "./test_datas/titanic-sample.csv"

func TestDataFrame_ReadCsv(t *testing.T)  {
	csvData := PrepareDataFrame4Test()

	correctCsvDataNumericSahpe := []int{891, 7}
	correctCsvDataStringSahpe := []int{891, 5}


	assert.ElementsMatch(t, correctCsvDataNumericSahpe, csvData.NumericShape,
		"Expected csvData's Numericshape is %s but acutually %s", correctCsvDataNumericSahpe, csvData.NumericShape,
	)
	assert.ElementsMatch(t, correctCsvDataStringSahpe, csvData.StringShape,
		"Expected csvData's Stringcshape is %s but acutually %s", correctCsvDataStringSahpe, csvData.StringShape,
	)

}

func ExampleDataFrame_Info() {
	csvData := PrepareDataFrame4Test()
	
	csvData.Info()
	// Output:
	// RangeIndex: 891 entries, 0 to 890
	// Data columns (total 12 columns):
	// ===== Numeric columns (total 7 columns) =====
	// name,non-null,null,dtype
	// Age,714,177,float64
	// Fare,891,0,float64
	// Parch,891,0,float64
	// PassengerId,891,0,float64
	// Pclass,891,0,float64
	// SibSp,891,0,float64
	// Survived,891,0,float64
	//
	// ===== String columns (total 5 columns) =====
	// name,non-null,null,dtype
	// Cabin,204,687,string
	// Embarked,889,2,string
	// Name,891,0,string
	// Sex,891,0,string
	// Ticket,891,0,string
	

}

func ExampleDataFrame_Describe() {
	csvData := PrepareDataFrame4Test()
	csvData.Describe()
	// Output:
	// === Numeric columns (total 7 columns) =====
	// metric,Age,Fare,Parch,PassengerId,Pclass,SibSp,Survived 
	// count,714.000,891.000,891.000,891.000,891.000,891.000,891.000
	// mean,29.699,32.204,0.382,446.000,2.309,0.523,0.384
	// std,14.516,49.666,0.806,257.209,0.836,1.102,0.486
	// min,0.420,0.000,0.000,1.000,1.000,0.000,0.000   
	// 25%,20.125,7.910,0.000,223.500,2.000,0.000,0.000
	// 50%,28.000,14.454,0.000,446.000,3.000,0.000,0.000
	// 75%,38.000,31.000,0.000,668.500,3.000,1.000,1.000
	// max,80.000,512.329,6.000,891.000,3.000,8.000,1.000
	// sum,21205.170,28693.949,340.000,397386.000,2057.000,466.000,342.000
	//
	// === String columns (total 5 columns) =====
	// metric,Cabin,Embarked,Name,Sex,Ticket
	// count,204,889,891,891,891
	// unique,147,3,891,2,681
	// freq,4,644,1,577,7
	// top,B96 B98,S,Abbing, Mr. Anthony,male,1601

}

func PrepareDataFrame4Test() lpandas.DataFrame {

	csvData := lpandas.DataFrame{}
	csvData.ReadCsv(filePath)
	return csvData
}