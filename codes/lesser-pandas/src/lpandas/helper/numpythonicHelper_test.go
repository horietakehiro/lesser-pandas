package helper_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
	"fmt"

	"lpandas/helper"
)

func TestNumpythonicFloatArray_Sum(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN()}
	sum := float64(15)

	assert.Equal(t, sum, arr.Sum())
}

func TestNumpythonicFloatArray_Sum_empty(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{}

	assert.True(t, math.IsNaN(arr.Sum()))
}

func TestNumpythonicFloatArray_Sum_invalid(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{math.NaN(), math.Inf(0)}

	assert.True(t, math.IsNaN(arr.Sum()))
}


func TestNumpythonicFloatArray_Max(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN(), 100.01, 100}
	max := float64(100.01)

	assert.Equal(t, max, arr.Max())
}
func TestNumpythonicFloatArray_Max_empty(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{}

	assert.True(t, math.IsNaN(arr.Max()))
}

func TestNumpythonicFloatArray_Max_invalid(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{math.NaN(), math.Inf(0)}

	assert.True(t, math.IsNaN(arr.Max()))
}



func TestNumpythonicFloatArray_Min(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN(), -100.01, -100, 0}
	min := float64(-100.01)

	assert.Equal(t, min, arr.Min())
}
func TestNumpythonicFloatArray_Min_empty(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{}

	assert.True(t, math.IsNaN(arr.Min()))
}

func TestNumpythonicFloatArray_Min_invalid(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{math.NaN(), math.Inf(0)}

	assert.True(t, math.IsNaN(arr.Min()))
}

func TestNumpythonicFloatArray_Mean(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3, math.NaN(), math.Inf(0), 4}
	mean := float64(2.5)

	assert.Equal(t, mean, arr.Mean())

}

func TestNumpythonicFloatArray_Mean_empty(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{}

	assert.True(t, math.IsNaN(arr.Mean()))
}

func TestNumpythonicFloatArray_Mean_invalid(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{math.NaN(), math.Inf(0)}

	assert.True(t, math.IsNaN(arr.Mean()))
}

func TestNumpythonicFloatArray_Std(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.NaN(), math.Inf(0)}
	std := "1.414"

	assert.Equal(t, std, fmt.Sprintf("%.3f" ,arr.Std(false)))

}

func TestNumpythonicFloatArray_Std_nMinus(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.NaN(), math.Inf(0)}
	std := "1.581"

	assert.Equal(t, std, fmt.Sprintf("%.3f" ,arr.Std(true)))

}


func TestNumpythonicFloatArray_Std_empty(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{}

	assert.True(t, math.IsNaN(arr.Std(true)))
}

func TestNumpythonicFloatArray_Std_invalid(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{math.NaN(), math.Inf(0)}

	assert.True(t, math.IsNaN(arr.Std(true)))
}


func TestNumpythonicFloatArray_Std_singleValue(t *testing.T) {
	
	arr := helper.NumpythonicFloatArray{1,}
	std := float64(0)
	assert.Equal(t, std, arr.Std(true))
}

func TestNumpythonicFloatArray_Sort(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,5,2,3,4, math.NaN(), math.Inf(0)}
	sort := helper.NumpythonicFloatArray{1,2,3,4,5, math.NaN(), math.Inf(0),}
	arr = arr.Sort(false)

	for i := 0; i < len(arr); i++ {
		if math.IsNaN(sort[i]) || math.IsInf(sort[i], 0) {
			assert.True(t, math.IsNaN(arr[i]) || math.IsInf(arr[i], 0))
		} else {
			assert.Equal(t, sort[i], arr[i])
		}
	}
}

func TestNumpythonicFloatArray_Sort_desc(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,5,2,3,4, math.NaN(), math.Inf(0)}
	sort := helper.NumpythonicFloatArray{math.Inf(0), math.NaN(), 5,4,3,2,1}
	arr = arr.Sort(true)

	for i := 0; i < len(arr); i++ {
		if math.IsNaN(sort[i]) || math.IsInf(sort[i], 0) {
			assert.True(t, math.IsNaN(arr[i]) || math.IsInf(arr[i], 0))
		} else {
			assert.Equal(t, sort[i], arr[i])
		}
	}
}

func TestNumpythonicFloatArray_Count(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN()}
	count := 5
	assert.Equal(t, count, arr.Count())
}

func TestNumpythonicFloatArray_Percentile(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5,6,7, math.Inf(0), math.NaN()}
	percentile := float64(2.5)

	assert.Equal(t, percentile, arr.Percentile(0.25))

}

func TestNumpythonicFloatArray_Percentile_empty(t *testing.T) {
	arr := helper.NumpythonicFloatArray{}

	assert.True(t, math.IsNaN(arr.Percentile(0.25)))

}

func TestNumpythonicFloatArray_Percentile_single(t *testing.T) {
	arr := helper.NumpythonicFloatArray{100}
	percentile := float64(100)
	assert.Equal(t, percentile, arr.Percentile(0.25))

}



func TestNumpythonicStringArray_Counter(t *testing.T) {
	arr := helper.NumpythonicStringArray{"a", "", "b", "a", "b", "c"}

	counter := map[string]int{
		"a" : 2, "b" : 2, "c" : 1,
	}

	for key, val := range counter {
		assert.Equal(t, val, arr.Counter()[key])
	}

}

func TestNumpythonicStringArray_Counter_empty(t *testing.T) {
	arr := helper.NumpythonicStringArray{}

	counter := map[string]int{}

	assert.Equal(t, counter, arr.Counter())

}

func TestNumpythonicStringArray_Counter_invalid(t *testing.T) {
	arr := helper.NumpythonicStringArray{"", ""}

	counter := map[string]int{}

	assert.Equal(t, counter, arr.Counter())

}

func TestNumpythonicStringArray_MostCommon(t *testing.T) {
	arr := helper.NumpythonicStringArray{"", "", "", "b", "c", "d", "a", "a", "b"}
	
	mostCommon := map[string]int{"a" : 2, "b" : 2, "c" : 1}

	for key, val := range mostCommon {
		assert.Equal(t, val, arr.MostCommon(3)[key])
	}
}

func TestNumpythonicStringArray_MostCommon_empty(t *testing.T) {
	arr := helper.NumpythonicStringArray{}
	
	mostCommon := map[string]int{}

	assert.Equal(t, mostCommon, arr.MostCommon(3))
}

func TestNumpythonicStringArray_MostCommon_invalid(t *testing.T) {
	arr := helper.NumpythonicStringArray{"", ""}
	
	mostCommon := map[string]int{}

	assert.Equal(t, mostCommon, arr.MostCommon(3))
}

func TestNumpythonicStringArray_MostCommon_overRequest(t *testing.T) {
	arr := helper.NumpythonicStringArray{"", "", "", "b", "c", "d", "a", "a", "b"}
	mostCommon := map[string]int{
		"a" : 2, "b" : 2, "c" : 1, "d" : 1, 
	}

	for key, val := range mostCommon {
		assert.Equal(t, val, arr.MostCommon(10)[key])
	}
}

func TestNumpythonicStringArray_MostCommon_underRequest(t *testing.T) {
	arr := helper.NumpythonicStringArray{"", "", "", "b", "c", "d", "a", "a", "b"}
	mostCommon := map[string]int{"a" : 2,}

	for key, val := range mostCommon {
		assert.Equal(t, val, arr.MostCommon(1)[key])
	}
}

func TestNumpythonicFloatArray_Brodcast_add(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN()}
	broadcast := helper.NumpythonicFloatArray{2,3,4,5,6, math.Inf(0), math.NaN()}

	ret := arr.Broadcast("add", 1)
	for i, val := range broadcast {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			assert.True(t, math.IsInf(ret[i], 0) || math.IsNaN(ret[i]))
		} else {
			assert.Equal(t, val, ret[i])
		}
	}
}


func TestNumpythonicFloatArray_Brodcast_sub(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN()}
	broadcast := helper.NumpythonicFloatArray{0,1,2,3,4, math.Inf(0), math.NaN()}

	ret := arr.Broadcast("sub", 1)
	for i, val := range broadcast {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			assert.True(t, math.IsInf(ret[i], 0) || math.IsNaN(ret[i]))
		} else {
			assert.Equal(t, val, ret[i])
		}
	}
}



func TestNumpythonicFloatArray_Brodcast_mul(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN()}
	broadcast := helper.NumpythonicFloatArray{5,10,15,20,25, math.Inf(0), math.NaN()}

	ret := arr.Broadcast("mul", 5)
	for i, val := range broadcast {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			assert.True(t, math.IsInf(ret[i], 0) || math.IsNaN(ret[i]))
		} else {
			assert.Equal(t, val, ret[i])
		}
	}
}

func TestNumpythonicFloatArray_Brodcast_div(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN()}
	broadcast := helper.NumpythonicFloatArray{0.5,1,1.5,2,2.5, math.Inf(0), math.NaN()}

	ret := arr.Broadcast("div", 2)
	for i, val := range broadcast {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			assert.True(t, math.IsInf(ret[i], 0) || math.IsNaN(ret[i]))
		} else {
			assert.Equal(t, val, ret[i])
		}
	}
}

func TestNumpythonicFloatArray_Brodcast_zeroDevide(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN()}
	broadcast := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN()}

	ret := arr.Broadcast("div", 0)
	for i, val := range broadcast {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			assert.True(t, math.IsInf(ret[i], 0) || math.IsNaN(ret[i]))
		} else {
			assert.Equal(t, val, ret[i])
		}
	}
}

func TestNumpythonicFloatArray_Brodcast_invalid(t *testing.T) {
	arr := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN()}
	broadcast := helper.NumpythonicFloatArray{1,2,3,4,5, math.Inf(0), math.NaN()}

	ret := arr.Broadcast("hoge", 0)
	for i, val := range broadcast {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			assert.True(t, math.IsInf(ret[i], 0) || math.IsNaN(ret[i]))
		} else {
			assert.Equal(t, val, ret[i])
		}
	}
}