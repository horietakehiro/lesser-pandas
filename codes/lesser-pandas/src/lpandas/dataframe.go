package lpandas

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
	"io"
)

// DataFrame is a struct for storing structured data, like csv.
type DataFrame struct {
	Columns []string
	Rows [][]float64 // first dim is row-wise, second dim is col-wise
	Shape [2]int
}

// GetShape get the dataframe's shape [r, c]
func (df *DataFrame) GetShape() {
	df.Shape[0] = len(df.Rows[0])
	df.Shape[1] = len(df.Columns)

}

// ReadCsv read csv file and store its first row as columns 
// and rest of rows as values.
func (df *DataFrame) ReadCsv(filePath string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)
	i := 0
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if i == 0 {
			// store first row as columns
			df.Columns = row
			i++
		} else {
			// convert rest of rows into float64
			tmpRow := make([]float64, len(df.Columns))
			for j, r := range row {
				r, err := strconv.ParseFloat(r, 64)
				if err != nil {
					panic(err)
				}
				tmpRow[j] = r
			}
			df.Rows = append(df.Rows, tmpRow)
		}
	}

	// finally get its shape
	df.GetShape()

}

// Describe stdout the dataframe's
// count, mean, max, min, 25percentile, 50percentile, 75percentile
// via column-wise
func (df *DataFrame) Describe() {
	means := make([]float64, df.Shape[1])
	mean := float64(0)
	for i := 0; i < df.Shape[1]; i++ {
		for j := 0; j < df.Shape[0]; j++ {
			mean += df.Rows[j][i]
		}
		means[i] = mean / float64(df.Shape[0])
	}
	fmt.Printf("%v", means[0])
}
