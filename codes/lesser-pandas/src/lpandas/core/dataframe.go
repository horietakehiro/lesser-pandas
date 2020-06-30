package core

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
	"reflect"
	"strings"
	"math"

	"lpandas/helper"
)



// DataFrame is a struct for storing structured data, like csv.
// DataFrame has two types of columns and rows : Numeric / String
type DataFrame struct {

	Numeric map[string]helper.NumpythonicFloatArray
	NumericColumns []string
	NumericShape [2]int

	String map[string]helper.NumpythonicStringArray
	StringColumns []string
	StringShape [2]int
	
}

// GetShape get the dataframe's shape [r, c]
func (df *DataFrame) GetShape() {
	for _, v := range df.Numeric {
		df.NumericShape[0] = len(v)
		break
	}
	df.NumericShape[1] = len(df.NumericColumns)

	for _, v := range df.String {
		df.StringShape[0] = len(v)
		break
	}
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

	// firstly, check whether all rows in each columns can be converted into float or not.
	allRows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	// initiate maps for suppress "assignment to entry in nil map" error
	df.Numeric = map[string]helper.NumpythonicFloatArray{}
	df.String = map[string]helper.NumpythonicStringArray{}

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
					df.String[col] = helper.NumpythonicStringArray{}
					isNumeric = false
					break
				}
			}
		}
		if isNumeric {
			df.NumericColumns = append(df.NumericColumns, col)
			df.Numeric[col] = helper.NumpythonicFloatArray{}
		}
	}


	// secondlly, store rows in String Columns as raw,
	// and store rows in Numeric Columns with converting into float
	allRowLength := len(allRows) - 1 // exclude colum row
	for i := 0; i < len(df.String) + len(df.Numeric); i++ {
		col := allRows[0][i]
		
		if helper.PythonicStrIfInList(col, df.StringColumns) {

			tmpStringRows := make(helper.NumpythonicStringArray, allRowLength)
			for j := 0; j < allRowLength; j++ {
				tmpStringRows[j] = allRows[j+1][i]
			}
			df.String[col] = tmpStringRows

		} else {

			tmpNumericRows := make(helper.NumpythonicFloatArray, allRowLength)
			for j := 0; j < allRowLength; j++ {

				tmpFloat, err := strconv.ParseFloat(allRows[j+1][i], 64)
				if err != nil {
					tmpFloat = math.NaN()
				}
				tmpNumericRows[j] = tmpFloat
			}
			df.Numeric[col] = tmpNumericRows
		}
	}
	// finally, store shapes of numeric and string rows
	df.GetShape()

}

// Info stdout an basic infomation of the DataFrame.
func (df *DataFrame) Info() {

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
	for _, key := range df.NumericColumns {
		validCount := df.Numeric[key].Count()
		fmt.Printf("%s,%d,%d,%s\n",
			key, 
			validCount,
			df.NumericShape[0] - validCount, 
			reflect.TypeOf(df.Numeric[key][0]))
	}
	
	fmt.Println("")

	// String
	fmt.Printf("===== String columns (total %d columns) =====\n",
		df.StringShape[1])
	fmt.Println("name,non-null,null,dtype")
	for _, key := range df.StringColumns {
		validCount := df.String[key].Count()
		fmt.Printf("%s,%d,%d,%s\n",
			key, 
			validCount, 
			df.StringShape[0] - validCount, 
			reflect.TypeOf(df.String[key][0]))
	}

}


// Describe stdout the dataframe's statistical description of each columns.
// NumericColumns : count, mean, std, min, 25%, 50%, 75%, max, sum
// StringColumns : count, unique, top, freq
func (df *DataFrame) Describe() {

	count := "count" + ","
	mean := "mean" + ","
	std := "std" + ","
	min := "min" + ","
	fistQuartile := "25%" + ","
	median := "50%" + ","
	thirdQuartile := "75%" + ","
	max := "max" + ","
	sum := "sum" + ","

	// count := "count" + ","
	unique := "unique" + ","
	freq := "freq" + ","
	top := "top" + ","


	fmt.Printf("=== Numeric columns (total %d columns) =====\n", len(df.NumericColumns))
	fmt.Printf("metric,%s\n", strings.Join(df.NumericColumns, ","))
	for i, col := range df.NumericColumns {
		count = df.createNumericResult(i, count, float64(df.Numeric[col].Count()))
		mean = df.createNumericResult(i, mean, df.Numeric[col].Mean())
		std = df.createNumericResult(i, std, df.Numeric[col].Std(false))
		min = df.createNumericResult(i, min, df.Numeric[col].Min())
		fistQuartile = df.createNumericResult(i, fistQuartile, df.Numeric[col].Percentile(0.25))
		median = df.createNumericResult(i, median, df.Numeric[col].Percentile(0.5))
		thirdQuartile = df.createNumericResult(i, thirdQuartile, df.Numeric[col].Percentile(0.75))
		max = df.createNumericResult(i, max, df.Numeric[col].Max())
		sum = df.createNumericResult(i, sum, df.Numeric[col].Sum())
		
	}
	fmt.Println(count)
	fmt.Println(mean)
	fmt.Println(std)
	fmt.Println(min)
	fmt.Println(fistQuartile)
	fmt.Println(median)
	fmt.Println(thirdQuartile)
	fmt.Println(max)
	fmt.Println(sum)


	fmt.Println("")
	count = "count" + "," // reuse 

	fmt.Printf("=== String columns (total %d columns) =====\n", len(df.StringColumns))
	fmt.Printf("metric,%s\n", strings.Join(df.StringColumns, ","))

	for i, col := range df.StringColumns {
		count = df.createStringIntResult(i, count, df.String[col].Count())
		unique = df.createStringIntResult(i, unique, len(df.String[col].Counter()))
		mostCommon := df.String[col].MostCommon(1)
		for key, val := range mostCommon {
			freq = df.createStringIntResult(i, freq, val)
			top = df.createStringStrResult(i, top, key)	
		}
	}
	fmt.Println(count)
	fmt.Println(unique)
	fmt.Println(freq)
	fmt.Println(top)

}

func (df *DataFrame) createNumericResult(i int, metric string, result float64) string {
	metric += fmt.Sprintf("%.3f", result)
	if i < len(df.NumericColumns) - 1 {
		metric += ","
	}
	return metric
}

func (df *DataFrame) createStringIntResult(i int, metric string, result int) string {
	metric += fmt.Sprintf("%d", result)
	if i < len(df.StringColumns) - 1 {
		metric += ","
	}
	return metric
}

func (df *DataFrame) createStringStrResult(i int, metric string, result string) string {
	metric += fmt.Sprintf("%s", result)
	if i < len(df.StringColumns) - 1 {
		metric += ","
	}
	return metric
}


// Sum stdouts the DataFrame's each numeric columns's sum
func (df *DataFrame) Sum() {
	for _, col := range df.NumericColumns {
		fmt.Printf("%s,%.3f\n", col, df.Numeric[col].Sum())
	}
}