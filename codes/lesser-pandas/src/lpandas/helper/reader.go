package helper

import (
	"os"
	"encoding/csv"
	"strconv"
	"io"

	"lpandas"
)


// ReadCsv read from file and return DataFrame struct.
// First row in file is supposed to be columnes([]string).
// Rest of all rows are supposed to be numeric data([][]float64).
func ReadCsv(filePath string) lpandas.DataFrame {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)
	df := lpandas.DataFrame{}
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

	return df

}