package core

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
	// "reflect"
	// "strings"
	"math"
	"io"

	"lpandas/helper"
)


// Series wraps the array-like objects and add some metadatas.
type Series struct {
	Dtype string
	Values interface{} // Values store the Values with various Dtypes
	Name string
	Index []string
}
// DataFrame consists of the set of serieses.
type DataFrame struct {
	Columns []string // for keeping the originam order
	Index []string // for keeping the original order
	Values map[string]*Series
	Shape [2]int // [length of rows, length of columns]
}


// GetShape get the dataframe's shape [r, c]
func (df *DataFrame) GetShape() {
	df.Shape[0] = len(df.Index)
	df.Shape[1] = len(df.Columns)
}

// ReadCsv read csv file and store its first row as Columns 
// and rest of rows as Values.
func (df *DataFrame) ReadCsv(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	// read each rows and store each elements into each serieses.
	var serieses []Series
	var tmpStringArray [][]string
	index := -1
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// store first row as Columns
		// determine the number of series obj to create
		// link each seriese with the dataframe
		if index == -1 {
			df.Columns = row
			df.Values = map[string]*Series{}
			serieses = make([]Series, len(row))
			tmpStringArray = make([][]string, len(row))
			for i, r := range row {
				serieses[i].Name = r
				df.Values[r] = &serieses[i]
			}
			index++
			continue
		}

		// store rest of all rows as string
		for i, r := range row {
			// serieses[i].Values = append(serieses[i].Values, r)
			tmpStringArray[i] = append(tmpStringArray[i], r)
			serieses[i].Index = append(serieses[i].Index, fmt.Sprintf("%d", index))
		}
		// store Index as string
		df.Index = append(df.Index, fmt.Sprintf("%d", index))
		index++
	}

	// converting series's Values into NumpythonicType
	for i := range serieses {
		serieses[i].Values = tmpStringArray[i]
		// type assertion must be neccessary to use Series.Values as type of []string
		if _, ok := serieses[i].Values.([]string); ok {
			serieses[i] = serieses[i].asNumpythonicType()
		}
	}

	df.GetShape()

}

// convert the series into either NumpythonicStringArray or NumpythonicFloatArray
func (sr Series) asNumpythonicType() Series {
	var newSeries Series
	var floatArray helper.NumpythonicFloatArray
	stringArray, ok := sr.Values.([]string)
	if ok {
		floatArray = make(helper.NumpythonicFloatArray, len(stringArray))
	}
	stringFlag := false
	for i, val := range stringArray {
		if val == "" {
			floatArray[i] = math.NaN()
			continue
		}
		
		float, err := strconv.ParseFloat(val, 64)
		if err != nil {
			stringFlag = true
			break
		}
		floatArray[i] = float
	}

	if stringFlag {
		// convert []string into helper.NumpythonicStringArray
		newStringArray := make(helper.NumpythonicStringArray, len(stringArray))
		for i, val := range stringArray {
			newStringArray[i] = val
		}
		newSeries = Series{
			Name : sr.Name, 
			Index : sr.Index, 
			Dtype : "string", 
			Values : newStringArray,
		}
	} else {
		newSeries = Series{
			Name : sr.Name, 
			Index : sr.Index, 
			Dtype : "float64", 
			Values : floatArray,
		}	
	}
	return newSeries
}


// Sum returns the sum values of each dataframe's columns as a Series 
// columns with string dtype returns the math.NaN()
func (df DataFrame) Sum() (Series) {
	retSeries := Series{
		Name : "sum", Index : make([]string, df.Shape[1]), Dtype : "float64"}
	values := make(helper.NumpythonicFloatArray, df.Shape[1])
	for i, col := range df.Columns {
		retSeries.Index[i] = col
		if df.Values[col].Dtype == "string" {
			values[i] = math.NaN()
		} else {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				values[i] = tmpValues.Sum()
			}
		}
	}
	retSeries.Values = values

	return retSeries
}

// Info returns the number of non-null values, null values and dtypes of each df's columns as a DataFrame
func (df DataFrame) Info() DataFrame {
	retDf := DataFrame{
		Columns : []string{"non-null", "null", "dtype"}, Index : df.Columns,
		Values : map[string]*Series{}, 
	}
	nonNullValues := make(helper.NumpythonicFloatArray, len(df.Columns))
	nullValues := make(helper.NumpythonicFloatArray, len(df.Columns))
	dtypeValues := make(helper.NumpythonicStringArray, len(df.Columns))


	for i, col := range df.Columns {
		if df.Values[col].Dtype == "string" {
			if asserted, ok := df.Values[col].Values.(helper.NumpythonicStringArray); ok {
				nonNullValues[i] = float64(asserted.Count())
				nullValues[i] = float64(len(asserted)) - nonNullValues[i]
				dtypeValues[i] = df.Values[col].Dtype
			}
		} else if df.Values[col].Dtype == "float64" {
			if asserted, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				nonNullValues[i] = float64(asserted.Count())
				nullValues[i] = float64(len(asserted)) - nonNullValues[i]
				dtypeValues[i] = df.Values[col].Dtype
			}
		}
	}

	retDf.Values["non-null"] = &Series{
		Name : "non-null", Index : retDf.Index, Dtype : "float64", Values : nonNullValues,
	}
	retDf.Values["null"] = &Series{
		Name : "null", Index : retDf.Index, Dtype : "float64", Values : nullValues,
	}
	retDf.Values["dtype"] = &Series{
		Name : "dtype", Index : retDf.Index, Dtype : "string", Values : dtypeValues,
	}

	retDf.GetShape()

	return retDf
}

// func convertDtype(sr Series) Series {

// 	// convert the whole series into either string or float64
// 	tmpFloatArray := make([]float64, len(sr.Values))
// 	for i, val := range sr.Values {
// 		if val == "" {
// 			tmpFloatArray[i] = math.NaN()
// 			continue
// 		}
// 		if str, ok := val.(string); ok { // must need type assertion
// 			float, err := strconv.ParseFloat(str, 64)
// 			if err != nil {
// 				// must be a string array
// 				sr.Dtype = "string"
// 				return sr
// 			}

// 			tmpFloatArray[i] = float
// 		}
// 	}
// 	newSeries := Series{Name : sr.Name, Index : sr.Index, Dtype : "float64"}
// 	for _, val := range tmpFloatArray {
// 		newSeries.Values = append(newSeries.Values, val)
// 	}
// 	return newSeries
// }

// // Info stdout an basic infomation of the DataFrame.
// func (df *DataFrame) Info() {

// 	// header
// 	if df.NumericShape[0] != 0 {
// 		fmt.Printf("RangeIndex: %d entries, %d to %d\n", 
// 				df.NumericShape[0], 0, df.NumericShape[0] - 1)
// 	} else {
// 		fmt.Printf("RangeIndex: %d entries, %d to %d\n", 
// 				df.StringShape[0], 0, df.StringShape[0] - 1)
// 	}
// 	fmt.Printf("Data Columns (total %d Columns):\n",
// 		df.NumericShape[1] + df.StringShape[1])

// 	// Numeric
// 	fmt.Printf("===== Numeric Columns (total %d Columns) =====\n",
// 		df.NumericShape[1])
// 	fmt.Println("Name,non-null,null,Dtype")
// 	for _, key := range df.NumericColumns {
// 		validCount := df.Numeric[key].Count()
// 		fmt.Printf("%s,%d,%d,%s\n",
// 			key, 
// 			validCount,
// 			df.NumericShape[0] - validCount, 
// 			reflect.TypeOf(df.Numeric[key][0]))
// 	}
// 	fmt.Println("")

// 	// String
// 	fmt.Printf("===== String Columns (total %d Columns) =====\n",
// 		df.StringShape[1])
// 	fmt.Println("Name,non-null,null,Dtype")
// 	for _, key := range df.StringColumns {
// 		validCount := df.String[key].Count()
// 		fmt.Printf("%s,%d,%d,%s\n",
// 			key, 
// 			validCount, 
// 			df.StringShape[0] - validCount, 
// 			reflect.TypeOf(df.String[key][0]))
// 	}

// }


// // Describe stdout the dataframe's statistical description of each Columns.
// // NumericColumns : count, mean, std, min, 25%, 50%, 75%, max, sum
// // StringColumns : count, unique, top, freq
// func (df *DataFrame) Describe() {

// 	count := "count" + ","
// 	mean := "mean" + ","
// 	std := "std" + ","
// 	min := "min" + ","
// 	fistQuartile := "25%" + ","
// 	median := "50%" + ","
// 	thirdQuartile := "75%" + ","
// 	max := "max" + ","
// 	sum := "sum" + ","

// 	// count := "count" + ","
// 	unique := "unique" + ","
// 	freq := "freq" + ","
// 	top := "top" + ","


// 	fmt.Printf("=== Numeric Columns (total %d Columns) =====\n", len(df.NumericColumns))
// 	fmt.Printf("metric,%s\n", strings.Join(df.NumericColumns, ","))
// 	for i, col := range df.NumericColumns {
// 		count = df.createNumericResult(i, count, float64(df.Numeric[col].Count()))
// 		mean = df.createNumericResult(i, mean, df.Numeric[col].Mean())
// 		std = df.createNumericResult(i, std, df.Numeric[col].Std(false))
// 		min = df.createNumericResult(i, min, df.Numeric[col].Min())
// 		fistQuartile = df.createNumericResult(i, fistQuartile, df.Numeric[col].Percentile(0.25))
// 		median = df.createNumericResult(i, median, df.Numeric[col].Percentile(0.5))
// 		thirdQuartile = df.createNumericResult(i, thirdQuartile, df.Numeric[col].Percentile(0.75))
// 		max = df.createNumericResult(i, max, df.Numeric[col].Max())
// 		sum = df.createNumericResult(i, sum, df.Numeric[col].Sum())
// 	}
// 	fmt.Println(count)
// 	fmt.Println(mean)
// 	fmt.Println(std)
// 	fmt.Println(min)
// 	fmt.Println(fistQuartile)
// 	fmt.Println(median)
// 	fmt.Println(thirdQuartile)
// 	fmt.Println(max)
// 	fmt.Println(sum)


// 	fmt.Println("")
// 	count = "count" + "," // reuse 

// 	fmt.Printf("=== String Columns (total %d Columns) =====\n", len(df.StringColumns))
// 	fmt.Printf("metric,%s\n", strings.Join(df.StringColumns, ","))

// 	for i, col := range df.StringColumns {
// 		count = df.createStringIntResult(i, count, df.String[col].Count())
// 		unique = df.createStringIntResult(i, unique, len(df.String[col].Counter()))
// 		mostCommon := df.String[col].MostCommon(1)
// 		for key, val := range mostCommon {
// 			freq = df.createStringIntResult(i, freq, val)
// 			top = df.createStringStrResult(i, top, key)	
// 		}
// 	}
// 	fmt.Println(count)
// 	fmt.Println(unique)
// 	fmt.Println(freq)
// 	fmt.Println(top)

// }

// func (df *DataFrame) createNumericResult(i int, metric string, result float64) string {
// 	metric += fmt.Sprintf("%.3f", result)
// 	if i < len(df.NumericColumns) - 1 {
// 		metric += ","
// 	}
// 	return metric
// }

// func (df *DataFrame) createStringIntResult(i int, metric string, result int) string {
// 	metric += fmt.Sprintf("%d", result)
// 	if i < len(df.StringColumns) - 1 {
// 		metric += ","
// 	}
// 	return metric
// }

// func (df *DataFrame) createStringStrResult(i int, metric string, result string) string {
// 	metric += fmt.Sprintf("%s", result)
// 	if i < len(df.StringColumns) - 1 {
// 		metric += ","
// 	}
// 	return metric
// }


// // Sum stdouts the DataFrame's each numeric Columns's sum
// func (df *DataFrame) Sum() string {
// 	ret := ""
// 	for _, col := range df.NumericColumns {
// 		ret += fmt.Sprintf("%s,%.3f\n", col, df.Numeric[col].Sum())
// 	}
// 	return ret
// }

// // Mean stdouts the DataFrame's each numeric Columns's mean
// func (df *DataFrame) Mean() string{
// 	ret := ""
// 	for _, col := range df.NumericColumns {
// 		ret += fmt.Sprintf("%s,%.3f\n", col, df.Numeric[col].Mean())
// 	}
// 	return ret
// }


// // Min stdout the DataFrame's each numeric Columns's min
// func (df *DataFrame) Min() string {
// 	ret := ""
// 	for _, col := range df.NumericColumns {
// 		ret += fmt.Sprintf("%s,%.3f\n", col, df.Numeric[col].Min())
// 	}
// 	return ret
// }

// // Max stdout the DataFrame's each numeric Columns's max
// func (df *DataFrame) Max() string {
// 	ret := ""
// 	for _, col := range df.NumericColumns {
// 		ret += fmt.Sprintf("%s,%.3f\n", col, df.Numeric[col].Max())
// 	}
// 	return ret
// }

// // Std stdout the DataFrame's each numeric Columns's std
// func (df *DataFrame) Std(nMinut1 bool) string {
// 	ret := ""
// 	for _, col := range df.NumericColumns {
// 		ret += fmt.Sprintf("%s,%.3f\n", col, df.Numeric[col].Std(nMinut1))
// 	}
// 	return ret
// }

// // Percentile stdout the DataFrame's each numeric Columns's percentile
// func (df *DataFrame) Percentile(location float64) string {
// 	ret := ""
// 	for _, col := range df.NumericColumns {
// 		ret += fmt.Sprintf("%s,%.3f\n", col, df.Numeric[col].Percentile(location))
// 	}
// 	return ret
// }
