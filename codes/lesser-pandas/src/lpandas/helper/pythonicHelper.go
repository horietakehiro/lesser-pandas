package helper

import (
	"github.com/aclements/go-gg/generic/slice"
	"math"
	"sort"
)

// AcsendingSort is a customized []float64
type AcsendingSort []float64
func (a AcsendingSort) Len() int {
	return len(a)
}
func (a AcsendingSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a AcsendingSort) Less(i, j int) bool {
	// consider float array includes NaN values
	return math.IsNaN(a[i]) || a[i] < a[j]
}


// PythonicStrIfInList checks the pythonic condition like `if string() in list()`
// and return the result as bool
func PythonicStrIfInList(target string, list []string) bool {
	for _,str := range list {
		if target == str {
			return true
		}
	}
	return false
}

// PythonicStrCounterMostCommon return the pair of (key, val) like python's collections.Counter.most_common(n)
func PythonicStrCounterMostCommon(counter map[string]int, n int) ([]string, []int) {
	keys, values := make([]string, len(counter)), make([]int, len(counter))
	retKeys, retValues := make([]string, n), make([]int, n)

	orderedKeys := make([]string, len(counter))
	i := 0
	for key := range counter {
		orderedKeys[i] = key
		i++
	}
	sort.Strings(orderedKeys)

	for i := 0; i < len(orderedKeys); i++ {
		keys[i] = orderedKeys[i]
		values[i] = counter[orderedKeys[i]]
		
	}

	for i := 0; i < n; i++ {
		index := slice.ArgMax(values)
		retKeys[i] = keys[index]
		retValues[i] = values[index]

		values[index] = 0		
	}

	return retKeys, retValues
}