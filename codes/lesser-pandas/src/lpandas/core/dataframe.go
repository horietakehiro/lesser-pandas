package core

import (
	"fmt"
	"os"
	"encoding/csv"
	"math"
	"io"
	"strings"

	"lpandas/helper"
)


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


// Count returns the number of valid values of each dataframe's columns as a Series 
func (df DataFrame) Count() (Series) {
	retSeries := Series{
		Name : "count", Index : make([]string, df.Shape[1]), Dtype : "float64"}
	values := make(helper.NumpythonicFloatArray, df.Shape[1])
	for i, col := range df.Columns {
		retSeries.Index[i] = col
		if df.Values[col].Dtype == "string" {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicStringArray); ok {
				values[i] = float64(tmpValues.Count())
			}
		} else {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				values[i] = float64(tmpValues.Count())
			}
		}
	}
	retSeries.Values = values

	return retSeries
}


// Mean returns the mean values of each dataframe's columns as a Series 
// columns with string dtype returns the math.NaN()
func (df DataFrame) Mean() (Series) {
	retSeries := Series{
		Name : "mean", Index : make([]string, df.Shape[1]), Dtype : "float64"}
	values := make(helper.NumpythonicFloatArray, df.Shape[1])
	for i, col := range df.Columns {
		retSeries.Index[i] = col
		if df.Values[col].Dtype == "string" {
			values[i] = math.NaN()
		} else {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				values[i] = tmpValues.Mean()
			}
		}
	}
	retSeries.Values = values

	return retSeries
}


// Std returns the std values of each dataframe's columns as a Series 
// columns with string dtype returns the math.NaN()
func (df DataFrame) Std(nMinus1 bool) (Series) {
	retSeries := Series{
		Name : "std", Index : make([]string, df.Shape[1]), Dtype : "float64"}
	values := make(helper.NumpythonicFloatArray, df.Shape[1])
	for i, col := range df.Columns {
		retSeries.Index[i] = col
		if df.Values[col].Dtype == "string" {
			values[i] = math.NaN()
		} else {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				values[i] = tmpValues.Std(nMinus1)
			}
		}
	}
	retSeries.Values = values

	return retSeries
}


// Min returns the min values of each dataframe's columns as a Series 
// columns with string dtype returns the math.NaN()
func (df DataFrame) Min() (Series) {
	retSeries := Series{
		Name : "min", Index : make([]string, df.Shape[1]), Dtype : "float64"}
	values := make(helper.NumpythonicFloatArray, df.Shape[1])
	for i, col := range df.Columns {
		retSeries.Index[i] = col
		if df.Values[col].Dtype == "string" {
			values[i] = math.NaN()
		} else {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				values[i] = tmpValues.Min()
			}
		}
	}
	retSeries.Values = values

	return retSeries
}


// Percentile returns the percentile values, where specified with 'location',  of each dataframe's columns as a Series 
// columns with string dtype returns the math.NaN()
func (df DataFrame) Percentile(location float64) (Series) {
	retSeries := Series{
		Name : fmt.Sprintf("%.1f%%", 100.0 * location), Index : make([]string, df.Shape[1]), Dtype : "float64"}
	values := make(helper.NumpythonicFloatArray, df.Shape[1])
	for i, col := range df.Columns {
		retSeries.Index[i] = col
		if df.Values[col].Dtype == "string" {
			values[i] = math.NaN()
		} else {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				values[i] = tmpValues.Percentile(location)
			}
		}
	}
	retSeries.Values = values

	return retSeries
}



// Max returns the max values of each dataframe's columns as a Series 
// columns with string dtype returns the math.NaN()
func (df DataFrame) Max() (Series) {
	retSeries := Series{
		Name : "max", Index : make([]string, df.Shape[1]), Dtype : "float64"}
	values := make(helper.NumpythonicFloatArray, df.Shape[1])
	for i, col := range df.Columns {
		retSeries.Index[i] = col
		if df.Values[col].Dtype == "string" {
			values[i] = math.NaN()
		} else {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				values[i] = tmpValues.Max()
			}
		}
	}
	retSeries.Values = values

	return retSeries
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



// Unique returns the number of unique values of each dataframe's columns as a Series 
func (df DataFrame) Unique() (Series) {
	retSeries := Series{
		Name : "unique", Index : make([]string, df.Shape[1]), Dtype : "float64"}
	values := make(helper.NumpythonicFloatArray, df.Shape[1])
	for i, col := range df.Columns {
		retSeries.Index[i] = col
		if df.Values[col].Dtype == "string" {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicStringArray); ok {
				counter := tmpValues.Counter()
				values[i] = float64(len(counter))
			}
		} else {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				counter := tmpValues.Counter()
				values[i] = float64(len(counter))
			}
		}
	}
	retSeries.Values = values

	return retSeries
}


// Freq returns the number of most common values's frequency of each dataframe's columns as a Series 
func (df DataFrame) Freq() (Series) {
	retSeries := Series{
		Name : "freq", Index : make([]string, df.Shape[1]), Dtype : "float64"}
	values := make(helper.NumpythonicFloatArray, df.Shape[1])
	for i, col := range df.Columns {
		retSeries.Index[i] = col
		if df.Values[col].Dtype == "string" {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicStringArray); ok {
				_, val := tmpValues.MostCommon(1)
				values[i] = float64(val[0])
			}
		} else {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				_, val := tmpValues.MostCommon(1)
				values[i] = float64(val[0])
			}
		}
	}
	retSeries.Values = values

	return retSeries
}


// Top returns the number of most common values of each dataframe's columns as a Series 
func (df DataFrame) Top() (Series) {
	retSeries := Series{
		Name : "top", Index : make([]string, df.Shape[1]), Dtype : "string"}
	values := make(helper.NumpythonicStringArray, df.Shape[1])
	for i, col := range df.Columns {
		retSeries.Index[i] = col
		if df.Values[col].Dtype == "string" {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicStringArray); ok {
				key, _ := tmpValues.MostCommon(1)
				values[i] = key[0]
			}
		} else {
			if tmpValues, ok := df.Values[col].Values.(helper.NumpythonicFloatArray); ok {
				key, _ := tmpValues.MostCommon(1)
				values[i] = key[0]
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


// Describe returns the basic statistical information of each df's columns as a DataFrame
// returned metrics are : "count", "mean", "std", "min","25.0%", "50.0%", "75.0%", "max","sum", "unique", "freq", "top"
func (df DataFrame) Describe() DataFrame {
	retDf := DataFrame{
		Columns : []string{"count", "mean", "std", "min",
							"25.0%", "50.0%", "75.0%", "max",
							"sum", "unique", "freq", "top"},
		Index : df.Columns,
		Values : map[string]*Series{}, 
	}

	count := df.Count(); retDf.Values["count"] = &count
	mean := df.Mean(); retDf.Values["mean"] = &mean
	std := df.Std(false); retDf.Values["std"] = &std
	min := df.Min(); retDf.Values["min"] = &min
	firstQuatile := df.Percentile(0.25); retDf.Values["25.0%"] = &firstQuatile
	median := df.Percentile(0.50); retDf.Values["50.0%"] = &median
	thirdQuatile := df.Percentile(0.75); retDf.Values["75.0%"] = &thirdQuatile
	max := df.Max(); retDf.Values["max"] = &max
	sum := df.Sum(); retDf.Values["sum"] = &sum

	unique := df.Unique(); retDf.Values["unique"] = &unique
	freq := df.Freq(); retDf.Values["freq"] = &freq
	top := df.Top(); retDf.Values["top"] = &top

	retDf.GetShape()

	return retDf
}

// Display displays the DataFrame with given format(csv|pretty)
func (df DataFrame) Display(format string) {
	switch format {
	case "csv":
		fmt.Printf("index,%s\n", strings.Join(df.Columns, ","))
		for i := 0; i < len(df.Index); i++ {
			var str string
			for j, col := range df.Columns {
				if df.Values[col].Dtype == "float64" {
					values, _ := df.Values[col].Values.(helper.NumpythonicFloatArray)
					str += fmt.Sprintf("%.3f" ,values[i])
				}
				if df.Values[col].Dtype == "string" {
					values, _ := df.Values[col].Values.(helper.NumpythonicStringArray)
					str += fmt.Sprintf("%s" ,values[i])
				}
				if j != len(df.Columns) - 1 {
					str += ","
				}
			}
			fmt.Printf("%s,%s\n", df.Index[i], str)
		}

	case "pretty":
		index := make(helper.NumpythonicStringArray, len(df.Index))
		for i, val := range df.Index {
			index[i] = val
		}
		indexLength := index.MaxLen()
		indexLength = int(math.Max(float64(indexLength), float64(len("index"))))
		valuesLength := map[string]int{}
		for _, col := range df.Columns {
			valuesLength[col] = 0
		}
		for _, col := range df.Columns {
			if df.Values[col].Dtype == "float64" {
				values, _ := df.Values[col].Values.(helper.NumpythonicFloatArray)
				valuesLength[col] = values.MaxLen()
			}
			if df.Values[col].Dtype == "string" {
				values, _ := df.Values[col].Values.(helper.NumpythonicStringArray)
				valuesLength[col] = values.MaxLen()
			}
			valuesLength[col] = int(math.Max(float64(valuesLength[col]), float64(len(df.Values[col].Name))))
		}


		header := make([]string, len(df.Columns) + 1)
		header[0] = helper.PadString("index", " ", indexLength)
		for i, col := range df.Columns {
			header[i+1] = helper.PadString(col, " ", valuesLength[col])
		}
		fmt.Printf("%s |\n", strings.Join(header, " | "))
		
		
		for i := 0; i < len(df.Index); i++ {
			fmtStrings := make([]string, len(df.Columns) + 1)
			fmtStrings[0] = helper.PadString(df.Index[i], " ", indexLength)
			for j, col := range df.Columns {
				if df.Values[col].Dtype == "float64" {
					values, _ :=  df.Values[col].Values.(helper.NumpythonicFloatArray)
					fmtStrings[j+1] = helper.PadString(fmt.Sprintf("%.3f", values[i]), " ", valuesLength[col])
				}
				if df.Values[col].Dtype == "string" {
					values, _ :=  df.Values[col].Values.(helper.NumpythonicStringArray)
					fmtStrings[j+1] = helper.PadString(values[i], " ", valuesLength[col])
				}
			}
			fmt.Printf("%s |\n", strings.Join(fmtStrings, " | "))

			
		}

	default:
		fmt.Println(df)	
	}
}