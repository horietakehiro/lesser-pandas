package core


import (
	// "fmt"
	"strconv"
	"math"

	"lpandas/helper"
)



// Series wraps the array-like objects and add some metadatas.
type Series struct {
	Dtype string
	Values interface{} // Values store the Values with various Dtypes
	Name string
	Index []string
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


// Count returns the number of valid values in the series.
// if the dtype of series is string, returns math.NaN()
func (sr Series) Count() float64 {
	var count float64
	if sr.Dtype == "string" {
		return math.NaN()
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			count = float64(values.Count())
		}
	}

	return count
}