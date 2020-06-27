package lpandas_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	// "github.com/golang/example/stringutil"

	"lpandas"
)

const filePath = "./test_datas/titanic-sample.csv"

func TestDataFrame1(t *testing.T)  {
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

func ExampleDataFrameInfo() {
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
// func ExampleDataFrame() {
// 	csvData := PrepareDataFrame4Test()
// 	csvData.Describe()
// 	// Output:
// 	// metric,f1,f2,f3
// 	// count,3.000,3.000,3.000
// 	// sum,2.250,12.000,3.000
// 	// mean,0.750,4.000,1.000
// 	// max,1.000,10.000,3.000
// 	// min,0.500,0.000,0.000
// 	// std,0.204,4.320,1.414
// }

func PrepareDataFrame4Test() lpandas.DataFrame {

	csvData := lpandas.DataFrame{}
	csvData.ReadCsv(filePath)
	return csvData
}