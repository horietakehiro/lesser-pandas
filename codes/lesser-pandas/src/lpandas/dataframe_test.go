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

	numericColumns := []string{"PassengerId","Survived","Pclass", "Age", "SibSp", "Parch", "Fare"}
	stringColumns := []string{"Name", "Sex", "Ticket", "Cabin", "Embarked"}

	assert.ElementsMatch(t, correctCsvDataNumericSahpe, csvData.NumericShape,
		"Expected csvData's Numericshape is %s but acutually %s", correctCsvDataNumericSahpe, csvData.NumericShape,
	)
	assert.ElementsMatch(t, correctCsvDataStringSahpe, csvData.StringShape,
		"Expected csvData's Stringcshape is %s but acutually %s", correctCsvDataStringSahpe, csvData.StringShape,
	)

	assert.ElementsMatch(t, stringColumns, csvData.StringColumns)
	assert.ElementsMatch(t, numericColumns, csvData.NumericColumns)

}

func ExampleDataFrame_Info() {
	csvData := PrepareDataFrame4Test()
	
	csvData.Info()
	// Output:
	// RangeIndex: 891 entries, 0 to 890
	// Data columns (total 12 columns):
	// ===== Numeric columns (total 7 columns) =====
	// name,non-null,null,dtype
	// PassengerId,891,0,float64
	// Survived,891,0,float64
	// Pclass,891,0,float64
	// Age,714,177,float64
	// SibSp,891,0,float64
	// Parch,891,0,float64
	// Fare,891,0,float64
	//
	// ===== String columns (total 5 columns) =====
	// name,non-null,null,dtype
	// Name,891,0,string
	// Sex,891,0,string
	// Ticket,891,0,string
	// Cabin,204,687,string
	// Embarked,889,2,string
	

}

func ExampleDataFrame_Describe() {
	csvData := PrepareDataFrame4Test()
	csvData.Describe()
	// Output:
	// === Numeric columns (total 7 columns) =====
	// metric,PassengerId,Survived,Pclass,Age,SibSp,Parch,Fare
	// count,891.000,891.000,891.000,714.000,891.000,891.000,891.000
	// mean,446.000,0.384,2.309,29.699,0.523,0.382,32.204
	// std,257.209,0.486,0.836,14.516,1.102,0.806,49.666
	// min,1.000,0.000,1.000,0.420,0.000,0.000,0.000   
	// 25%,223.500,0.000,2.000,20.125,0.000,0.000,7.910
	// 50%,446.000,0.000,3.000,28.000,0.000,0.000,14.454
	// 75%,668.500,1.000,3.000,38.000,1.000,0.000,31.000
	// max,891.000,1.000,3.000,80.000,8.000,6.000,512.329
	// sum,397386.000,342.000,2057.000,21205.170,466.000,340.000,28693.949
	//
	// === String columns (total 5 columns) =====
	// metric,Name,Sex,Ticket,Cabin,Embarked   
	// count,891,891,891,204,889
	// unique,891,2,681,147,3  
	// freq,1,577,7,4,644
	// top,Abbing, Mr. Anthony,male,1601,B96 B98,S

}

func PrepareDataFrame4Test() lpandas.DataFrame {

	csvData := lpandas.DataFrame{}
	csvData.ReadCsv(filePath)
	return csvData
}