package core


import (
	"fmt"
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
func (sr Series) Count() float64 {
	var count float64
	if sr.Dtype == "string" {
		if values, ok := sr.Values.(helper.NumpythonicStringArray); ok {
			count = float64(values.Count())
		}
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			count = float64(values.Count())
		}
	}

	return count
}


// Mean returns the mean of the series.
func (sr Series) Mean() float64 {
	var mean float64
	if sr.Dtype == "string" {
		mean = math.NaN()
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			mean = values.Mean()
		}
	}

	return mean
}



// Std returns the std of the series.
func (sr Series) Std(nMinus1 bool) float64 {
	var std float64
	if sr.Dtype == "string" {
		std = math.NaN()
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			std = values.Std(nMinus1)
		}
	}

	return std
}


// Min returns the min of the series.
func (sr Series) Min() float64 {
	var min float64
	if sr.Dtype == "string" {
		min = math.NaN()
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			min = values.Min()
		}
	}

	return min
}



// Percentile returns the percentile at the specified location of the series.
func (sr Series) Percentile(location float64) float64 {
	var percentile float64

	if sr.Dtype == "string" {
		percentile = math.NaN()
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			percentile = values.Percentile(location)
		}
	}

	return percentile
}


// Max returns the max of the series.
func (sr Series) Max() float64 {
	var max float64
	if sr.Dtype == "string" {
		max = math.NaN()
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			max = values.Max()
		}
	}

	return max
}

// Sum returns the sum of the series.
func (sr Series) Sum() float64 {
	var sum float64
	if sr.Dtype == "string" {
		sum = math.NaN()
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			sum = values.Sum()
		}
	}

	return sum
}


// Unique returns the number of unique values of the series.
func (sr Series) Unique() float64 {
	var unique float64
	var counter map[string]int
	if sr.Dtype == "string" {
		if values, ok := sr.Values.(helper.NumpythonicStringArray); ok {
			counter = values.Counter()
		}
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			counter = values.Counter()
		}
	}
	unique = float64(len(counter))

	return unique
}



// Freq returns the most common values' frequency of the series.
func (sr Series) Freq() float64 {
	var freq float64
	if sr.Dtype == "string" {
		if values, ok := sr.Values.(helper.NumpythonicStringArray); ok {
			_, val := values.MostCommon(1)
			freq = float64(val[0])
		}
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			_, val := values.MostCommon(1)
			freq = float64(val[0])
		}
	}

	return freq
}


// Top returns the most common values of the series.
func (sr Series) Top() string {
	var top string
	if sr.Dtype == "string" {
		if values, ok := sr.Values.(helper.NumpythonicStringArray); ok {
			keys, _ := values.MostCommon(1)
			top = keys[0]
		}
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			keys, _ := values.MostCommon(1)
			top = keys[0]
		}
	}

	return top
}


// Info returns the number of non-null values, null values, and the dtype of the series
// The returned series is dtype of string.
func (sr Series) Info() Series {
	var nonNull float64
	var null float64
	var dtype = "string"
	newSeries := Series{
		Name : sr.Name, Index : []string{"non-null", "null", "dtype"},
		Dtype : "string", 
	}
	if sr.Dtype == "string" {
		if values, ok := sr.Values.(helper.NumpythonicStringArray); ok {
			nonNull = float64(values.Count())
			null = float64(len(values)) - nonNull
		}
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			nonNull = float64(values.Count())
			null = float64(len(values)) - nonNull
		}
	}

	stringArray := helper.NumpythonicStringArray{
		fmt.Sprintf("%.3f", nonNull), fmt.Sprintf("%.3f", null), dtype, 
	}

	newSeries.Values = stringArray
	
	return newSeries
}



// Describe returns the basic statistical values of the series
// The returned series is dtype of string.
func (sr Series) Describe() Series {

	describe := map[string]string{
		"count" : "", "mean" : "", "std" : "", "min" : "",
		"25.0%" : "", "50.0%" : "", "75.0%" : "", "max" : "",
		"sum" : "", "unique" : "", "freq" : "", "top" : "",
	}
	index := []string{
		"count", "mean", "std", "min", 
		"25.0%", "50.0%", "75.0%", "max",
		"sum", "unique", "freq", "top",}
	
	newSeries := Series{
		Name : sr.Name, Index : index,
		Dtype : "string", 
	}

	if sr.Dtype == "string" {
		if values, ok := sr.Values.(helper.NumpythonicStringArray); ok {
			describe["count"] = fmt.Sprintf("%.3f", float64(values.Count()))
			counter := values.Counter()
			k, v := values.MostCommon(1)
			describe["unique"] = fmt.Sprintf("%.3f", float64(len(counter)))
			describe["freq"] = fmt.Sprintf("%.3f", float64(v[0]))
			describe["top"] = fmt.Sprintf("%s", k[0])			
		}
	}

	if sr.Dtype == "float64" {
		if values, ok := sr.Values.(helper.NumpythonicFloatArray); ok {
			describe["count"] = fmt.Sprintf("%.3f", float64(values.Count()))
			describe["mean"] = fmt.Sprintf("%.3f", float64(values.Mean()))
			describe["std"] = fmt.Sprintf("%.3f", float64(values.Std(false)))
			describe["min"] = fmt.Sprintf("%.3f", float64(values.Min()))
			describe["25.0%"] = fmt.Sprintf("%.3f", float64(values.Percentile(0.25)))
			describe["50.0%"] = fmt.Sprintf("%.3f", float64(values.Percentile(0.5)))
			describe["75.0%"] = fmt.Sprintf("%.3f", float64(values.Percentile(0.75)))
			describe["max"] = fmt.Sprintf("%.3f", float64(values.Max()))
			describe["sum"] = fmt.Sprintf("%.3f", float64(values.Sum()))
			counter := values.Counter()
			k, v := values.MostCommon(1)
			describe["unique"] = fmt.Sprintf("%.3f", float64(len(counter)))
			describe["freq"] = fmt.Sprintf("%.3f", float64(v[0]))
			describe["top"] = fmt.Sprintf("%s", k[0])	
		}
	}

	stringArray := make(helper.NumpythonicStringArray, len(index))
	for i, val := range index {
		stringArray[i] = describe[val]
	}

	newSeries.Values = stringArray
	
	return newSeries
}