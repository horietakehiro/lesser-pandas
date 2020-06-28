package helper

import (
	"sort"
	"math"
	"github.com/aclements/go-gg/generic/slice"

)

// NumpythonicFloatArray can be used like python's numpy float array.
type NumpythonicFloatArray []float64

// NumpythonicStringArray can be used like python's numpy string array.
// Or, implements some utility function which can be used in python.
type NumpythonicStringArray []string


func isValid(val float64) bool {
	return !math.IsInf(val, 0) && !math.IsNaN(val)
}

// Count returns array's total numbers of valid values 
func (arr NumpythonicFloatArray) Count() int {
	count := 0
	for _, val := range arr {
		if isValid(val) {
			count++
		}
	}

	return count

}

// Sum returns the arr's sum
func (arr NumpythonicFloatArray) Sum() float64 {

	if len(arr) == 0 {
		return math.NaN()
	}

	sum := float64(0)
	invalidFlag := true
	for _, val := range arr {
		// exclude NaN and InF values
		if isValid(val) {
			sum += val
			invalidFlag = false
		}
	}

	if invalidFlag {
		return math.NaN()
	}
	return sum
}


// Max returns the array's max
func (arr NumpythonicFloatArray) Max() float64 {
	if len(arr) == 0 {
		return math.NaN()
	}

	max := float64(0)
	initFlag := true
	invalidFlag := true
	for _, val := range arr {
		// initiate 
		if initFlag && isValid(val) {
			max = val
			initFlag = false
			invalidFlag = false
		}

		if max < val && isValid(val) {
			max = val
		}

		if invalidFlag {
			return math.NaN()
		}
	}
	return max
}


// Min returns the array's min
func (arr NumpythonicFloatArray) Min() float64 {
	if len(arr) == 0 {
		return math.NaN()
	}

	min := float64(0)
	initFlag := true
	invalidFlag := true
	for _, val := range arr {
		// initiate 
		if initFlag && isValid(val) {
			min = val
			initFlag = false
			invalidFlag = false
		}

		if min > val && isValid(val) {
			min = val
		}

		if invalidFlag {
			return math.NaN()
		}
	}
	return min
}

// Mean returns the array's mean (exclude NaN or Inf values in calculation)
func (arr NumpythonicFloatArray) Mean() float64 {
	validLength := arr.Count()
	if validLength == 0 {
		return math.NaN()
	}
	mean := arr.Sum() / float64(validLength)
	return mean
}


// Std returns the array's std (exclude NaN or Inf values in calculation)
func (arr NumpythonicFloatArray) Std(nMinus1 bool) float64 {
	validLength := arr.Count()
	if validLength == 0 {
		return math.NaN()
	}

	if nMinus1 {
		validLength--
	}
	if validLength == 0 {
		return float64(0.0)
	}

	mean := arr.Mean()
	sigmaSquared := float64(0)
	for _, val := range arr {
		if isValid(val) {
			sigmaSquared += math.Pow(val - mean, 2)
		}
	}
	std := math.Sqrt(sigmaSquared / float64(validLength))

	return std
}

// Sort returns the sorted array
func (arr NumpythonicFloatArray) Sort(desc bool) NumpythonicFloatArray {
	slice.Sort(arr)

	if desc {
		reversedArray := make(NumpythonicFloatArray, len(arr))
		for i, val := range arr {
			reversedArray[(len(arr) - 1) - i] = val
		}
		arr = reversedArray
	}

	return arr
}

// Percentile returns array's values located on the location
func (arr NumpythonicFloatArray) Percentile(location float64) float64 {
	validValues := NumpythonicFloatArray{}
	for _, val := range arr {
		if isValid(val) {
			validValues = append(validValues, val)
		}
	}
	if len(validValues) == 0 {
		return math.NaN()
	}
	if len(validValues) == 1 {
		return validValues[0]
	}

	N := len(validValues) - 1 // this is not a length, but distance from head to tail
	p := float64(N) * location
	q := int(math.Floor(p))
	r := p - float64(q)
	D := validValues[q] + (validValues[q+1] - validValues[q]) * r // linear interporation
	return D

}


// Broadcast returns the array calculated with the given value by the given operation
func (arr NumpythonicFloatArray) Broadcast(operation string, value float64) NumpythonicFloatArray {
	retArray := make(NumpythonicFloatArray, len(arr))

	switch operation {
		case "add" : {
			retArray = arr.add(value)
		}
		case "sub" : {
			retArray = arr.sub(value)
		}
		case "mul" : {
			retArray = arr.mul(value)
		}
		case "div" : {
			retArray = arr.div(value)
		}
		default : {
			retArray = arr
		}

				
	}

	return retArray
}

func (arr NumpythonicFloatArray) add(value float64) NumpythonicFloatArray {
	
	retArray := make(NumpythonicFloatArray, len(arr))
	for i, val := range arr {
		retArray[i] = val + value
	}
	return retArray

}

func (arr NumpythonicFloatArray) sub(value float64) NumpythonicFloatArray {
	
	retArray := make(NumpythonicFloatArray, len(arr))
	for i, val := range arr {
		retArray[i] = val - value
	}
	return retArray

}


func (arr NumpythonicFloatArray) mul(value float64) NumpythonicFloatArray {
	
	retArray := make(NumpythonicFloatArray, len(arr))
	for i, val := range arr {
		retArray[i] = val * value
	}
	return retArray
}

func (arr NumpythonicFloatArray) div(value float64) NumpythonicFloatArray {
	if value == 0 {
		return arr
	}

	retArray := make(NumpythonicFloatArray, len(arr))
	for i, val := range arr {
		retArray[i] = val / value
	}
	return retArray
}



// Counter returns array's each elemnts' frequencey
func (arr NumpythonicStringArray) Counter() map[string]int {
	counter := map[string]int{}

	for _, val := range arr {
		if val != "" {
			counter[val]++
		}
	}

	return counter
}


// MostCommon returns the array's most common elements and its frequency, ordered by ascending.
func (arr NumpythonicStringArray) MostCommon(n int) map[string]int {
	mostCommon := map[string]int{}

	counter := arr.Counter()

	if len(counter) == 0 {
		return mostCommon
	}

	keys := make([]string, len(counter))
	values := make([]int, len(counter))

	orderedKeys := make([]string, len(counter))
	i := 0
	for key := range counter {
		orderedKeys[i] = key
		i++
	}
	sort.Strings(orderedKeys)

	for i, key := range orderedKeys {
		keys[i] = key
		values[i] = counter[key]
	}

	remain := len(counter)
	for i := 0; i < n; i++ {
		if i == remain {
			break
		}
		index := slice.ArgMax(values)
		mostCommon[keys[index]] = values[index]
		values[index] = 0
	}

	return mostCommon
}