package lpandas

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
	"reflect"
	// "strings"
	// "io"
	"math"

	"lpandas/helper"
)

// DataFrame is a struct for storing structured data, like csv.
// DataFrame has two types of columns and rows : Numeric / String
type DataFrame struct {
	NumericColumns []string
	NumericRows [][]float64
	NumericShape [2]int

	StringColumns []string
	StringRows [][]string
	StringShape [2]int
	
}

// GetShape get the dataframe's shape [r, c]
func (df *DataFrame) GetShape() {
	df.NumericShape[0] = len(df.NumericRows)
	df.NumericShape[1] = len(df.NumericColumns)

	df.StringShape[0] = len(df.StringRows)
	df.StringShape[1] = len(df.StringColumns)
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
	allRows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	// firstly, check whether all rows in each columns can be converted into float or not.
	for i, col := range allRows[0] {
		isNumeric := true
		for j, row := range allRows {
			if j == 0 {
				// first row is columns
				continue
			}
			if row[i] == "" {
				// "" can be converted into NaN
				continue
			} else {
				_, err := strconv.ParseFloat(row[i], 64)
				if err != nil {
					df.StringColumns = append(df.StringColumns, col)
					isNumeric = false
					break
				}
			}
		}
		if isNumeric {
			df.NumericColumns = append(df.NumericColumns, col)
		}
	}

	// secondlly, store rows in StringColumns as raw
	// store rows in NumericColumns with converting into float
	for i, row := range allRows {
		if i == 0 {
			// first row is columns
			continue
		}
		tmpNumericRow := make([]float64, len(df.NumericColumns))
		tmpStringRow := make([]string, len(df.StringColumns))
		tmpNumericIndex := int(0)
		tmpStringIndex := int(0)
		for j, r := range row {
			if helper.PythonicStrIfInList(allRows[0][j], df.StringColumns) {
				tmpStringRow[tmpStringIndex] = r
				tmpStringIndex++
			} else {
				rFloated, err := strconv.ParseFloat(r, 64)
				if err != nil {
					rFloated = math.NaN()
				}
				tmpNumericRow[tmpNumericIndex] = rFloated
				tmpNumericIndex++
			}
		}
		df.StringRows = append(df.StringRows, tmpStringRow)
		df.NumericRows = append(df.NumericRows, tmpNumericRow)
	}

	// finally, store shapes of numeric and string rows
	df.GetShape()

}

// Info stdout an basic infomation of the DataFrame.
func (df *DataFrame) Info() {
	// count the number of null rows in each columns.
	// We handle math.NaN in NumericRows and "" in StringRows as null value
	numericNullCounts := make(map[string]int, len(df.NumericColumns))
	stringNullCounts := make(map[string]int, len(df.StringColumns))

	// initiate Nullounts maps
	for _, col := range df.NumericColumns {
		numericNullCounts[col] = 0
	}
	for _, col := range df.StringColumns {
		stringNullCounts[col] = 0
	}

	// count null rows 
	for _, row := range df.NumericRows {
		for j, col := range df.NumericColumns {
			if math.IsNaN(row[j]) {
				numericNullCounts[col]++
			}
		}
	}
	for _, row := range df.StringRows {
		for j, col := range df.StringColumns {
			if row[j] == "" {
				stringNullCounts[col]++
			}
		}
	}

	stdoutInfo(df, numericNullCounts, stringNullCounts)


}


func stdoutInfo(df *DataFrame, numericNullCounts, stringNullCounts map[string]int) {
	// header
	if df.NumericShape[0] != 0 {
		fmt.Printf("RangeIndex: %d entries, %d to %d\n", 
				df.NumericShape[0], 0, df.NumericShape[0] - 1)
	} else {
		fmt.Printf("RangeIndex: %d entries, %d to %d\n", 
				df.StringShape[0], 0, df.StringShape[0] - 1)
	}
	fmt.Printf("Data columns (total %d columns):\n",
		df.NumericShape[1] + df.StringShape[1])

	// Numeric
	fmt.Printf("===== Numeric columns (total %d columns) =====\n",
		df.NumericShape[1])
	fmt.Println("name,non-null,null,dtype")
	for _,col := range df.NumericColumns {
		fmt.Printf("%s,%d,%d,%s\n",
			col, df.NumericShape[0] - numericNullCounts[col], 
			numericNullCounts[col], reflect.TypeOf(df.NumericRows[0][0]))
	}
	
	fmt.Println("")

	// String
	fmt.Printf("===== String columns (total %d columns) =====\n",
		df.StringShape[1])
	fmt.Println("name,non-null,null,dtype")
	for _,col := range df.StringColumns {
		fmt.Printf("%s,%d,%d,%s\n",
			col, df.StringShape[0] - stringNullCounts[col], 
			stringNullCounts[col], reflect.TypeOf(df.StringRows[0][0]))
	}
	
}
// // Describe stdout the dataframe's
// // count, mean, max, min, 25percentile, 50percentile, 75percentile
// // via column-wise
// func (df *DataFrame) Describe() {
// 	counts := make([]float64, df.Shape[1])
// 	sums := make([]float64, df.Shape[1])
// 	means := make([]float64, df.Shape[1])
// 	maxes := make([]float64, df.Shape[1])
// 	mins := make([]float64, df.Shape[1])
// 	stds := make([]float64, df.Shape[1])


// 	for i := 0; i < df.Shape[1]; i++ {
// 		// count := float64(0)
// 		mean := float64(0)
// 		sum := float64(0)
// 		max := float64(0)
// 		min := float64(0)
// 		std := float64(0)

// 		for j := 0; j < df.Shape[0]; j++ {
// 			val := df.Rows[j][i]
// 			sum += val
// 			// initialize max, min
// 			if j == 0 {
// 				max = val
// 				min = val
// 			}

// 			if max < val {
// 				max = val
// 			}
// 			if min > val {
// 				min = val
// 			}
// 		}
// 		// calc mean using sum
// 		mean = sum / float64(df.Shape[0])

// 		// calc standard deviation using mean
// 		sigmaSquared := float64(0)
// 		for j := 0; j < df.Shape[0]; j++ {
// 			sigmaSquared += math.Pow(df.Rows[j][i] - mean, 2)
// 		}
// 		std = math.Sqrt(sigmaSquared / float64(df.Shape[0]))

// 		counts[i] = float64(df.Shape[0])
// 		sums[i] = sum
// 		means[i] = mean
// 		maxes[i] = max
// 		mins[i] = min
// 		stds[i] = std
// 	}
// 	stduutDescribe(df, counts, sums, means, maxes, mins, stds)
// }

// func stduutDescribe(df *DataFrame, counts, sums, means, maxes, mins, stds []float64) {
// 	strCounts := make([]string, len(counts))
// 	strSums := make([]string, len(sums))
// 	strMeans := make([]string, len(means))
// 	strMaxes := make([]string, len(maxes))
// 	strMins := make([]string, len(mins))
// 	strStds := make([]string, len(stds))


// 	for i := 0; i < df.Shape[0]; i++ {
// 		strCounts[i] = fmt.Sprintf("%.3f", counts[i])
// 		strSums[i] = fmt.Sprintf("%.3f", sums[i])
// 		strMeans[i] = fmt.Sprintf("%.3f", means[i])
// 		strMaxes[i] = fmt.Sprintf("%.3f", maxes[i])
// 		strMins[i] = fmt.Sprintf("%.3f", mins[i])
// 		strStds[i] = fmt.Sprintf("%.3f", stds[i])
// 	}

// 	fmt.Printf("metric,%s\n",strings.Join(df.Columns, ","))
// 	fmt.Printf("count,%s\n",strings.Join(strCounts, ","))
// 	fmt.Printf("sum,%s\n",strings.Join(strSums, ","))
// 	fmt.Printf("mean,%s\n",strings.Join(strMeans, ","))
// 	fmt.Printf("max,%s\n",strings.Join(strMaxes, ","))
// 	fmt.Printf("min,%s\n",strings.Join(strMins, ","))	
// 	fmt.Printf("std,%s\n",strings.Join(strStds, ","))

// }
