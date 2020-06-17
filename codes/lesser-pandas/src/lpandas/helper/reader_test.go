package helper_test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"lpandas/helper"
	"lpandas"
)

func TestReadCsv(t *testing.T)  {
	const file_path = "../test_datas/sample.csv"
	correctCsvData := lpandas.DataFrame{
		Columns : []string{"f1", "f2", "f3"},
		Rows : [][]float64{{float64(1), float64(2), float64(3)},
							{float64(0.5), float64(10), float64(0)}},
	}

	csvData := helper.ReadCsv(file_path)
	assert.ElementsMatch(t, csvData.Columns, correctCsvData.Columns,
		"Expected csvData columnes are %s, but actually %s", correctCsvData.Columns, csvData.Columns,
	)
	
}