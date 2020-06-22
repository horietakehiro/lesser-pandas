package lpandas_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	// "github.com/golang/example/stringutil"

	"lpandas"
)

func TestDataFrame1(t *testing.T)  {
	csvData := PrepareDataFrame4Test()

	correctCsvData := lpandas.DataFrame{
		Columns : []string{"f1", "f2", "f3"},
		Rows : [][]float64{{float64(1), float64(2), float64(3)},
							{float64(0.5), float64(10), float64(0)}},
	}

	assert.ElementsMatch(t, csvData.Columns, correctCsvData.Columns,
		"Expected csvData columnes are %s, but actually %s", correctCsvData.Columns, csvData.Columns,
	)

}

func TestDataFrame2(t *testing.T)  {
	csvData := PrepareDataFrame4Test()

	correctCsvDataSahpe := []int{3, 3}
	assert.ElementsMatch(t, correctCsvDataSahpe, csvData.Shape,
		"Expected csvData's shape is %s but acutually %s", correctCsvDataSahpe, csvData.Shape,
	)

}

func ExampleDataFrame() {
	csvData := PrepareDataFrame4Test()
	csvData.Describe()
	// Output: 0.75
}

func PrepareDataFrame4Test() lpandas.DataFrame {
	const filePath = "./test_datas/sample.csv"

	csvData := lpandas.DataFrame{}
	csvData.ReadCsv(filePath)
	return csvData
}