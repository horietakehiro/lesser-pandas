package lpandas

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
	"strings"
	"io"
	"math"
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
	counts := make([]float64, df.Shape[1])
	sums := make([]float64, df.Shape[1])
	means := make([]float64, df.Shape[1])
	maxes := make([]float64, df.Shape[1])
	mins := make([]float64, df.Shape[1])
	stds := make([]float64, df.Shape[1])


	for i := 0; i < df.Shape[1]; i++ {
		// count := float64(0)
		mean := float64(0)
		sum := float64(0)
		max := float64(0)
		min := float64(0)
		std := float64(0)

		for j := 0; j < df.Shape[0]; j++ {
			val := df.Rows[j][i]
			sum += val
			// initialize max, min
			if j == 0 {
				max = val
				min = val
			}

			if max < val {
				max = val
			}
			if min > val {
				min = val
			}
		}
		// calc mean using sum
		mean = sum / float64(df.Shape[0])

		// calc standard deviation using mean
		sigmaSquared := float64(0)
		for j := 0; j < df.Shape[0]; j++ {
			sigmaSquared += math.Pow(df.Rows[j][i] - mean, 2)
		}
		std = math.Sqrt(sigmaSquared / float64(df.Shape[0]))

		counts[i] = float64(df.Shape[0])
		sums[i] = sum
		means[i] = mean
		maxes[i] = max
		mins[i] = min
		stds[i] = std
	}
	stduutDescribe(df, counts, sums, means, maxes, mins, stds)
}

func stduutDescribe(df *DataFrame, counts, sums, means, maxes, mins, stds []float64) {
	strCounts := make([]string, len(counts))
	strSums := make([]string, len(sums))
	strMeans := make([]string, len(means))
	strMaxes := make([]string, len(maxes))
	strMins := make([]string, len(mins))
	strStds := make([]string, len(stds))


	for i := 0; i < df.Shape[0]; i++ {
		strCounts[i] = fmt.Sprintf("%.3f", counts[i])
		strSums[i] = fmt.Sprintf("%.3f", sums[i])
		strMeans[i] = fmt.Sprintf("%.3f", means[i])
		strMaxes[i] = fmt.Sprintf("%.3f", maxes[i])
		strMins[i] = fmt.Sprintf("%.3f", mins[i])
		strStds[i] = fmt.Sprintf("%.3f", stds[i])
	}

	fmt.Printf("metric,%s\n",strings.Join(df.Columns, ","))
	fmt.Printf("count,%s\n",strings.Join(strCounts, ","))
	fmt.Printf("sum,%s\n",strings.Join(strSums, ","))
	fmt.Printf("mean,%s\n",strings.Join(strMeans, ","))
	fmt.Printf("max,%s\n",strings.Join(strMaxes, ","))
	fmt.Printf("min,%s\n",strings.Join(strMins, ","))	
	fmt.Printf("std,%s\n",strings.Join(strStds, ","))

}
