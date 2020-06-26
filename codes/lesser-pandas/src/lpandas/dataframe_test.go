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