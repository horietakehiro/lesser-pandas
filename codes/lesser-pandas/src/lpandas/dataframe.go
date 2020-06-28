package lpandas

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
	"reflect"
	"sort"
	"strings"
	// "io"
	"math"

	"lpandas/helper"
)



// DataFrame is a struct for storing structured data, like csv.
// DataFrame has two types of columns and rows : Numeric / String
type DataFrame struct {

	Numeric map[string][]float64
	NumericShape [2]int

	String map[string][]string
	StringShape [2]int
	
}

// GetShape get the dataframe's shape [r, c]
func (df *DataFrame) GetShape() {
	for _, v := range df.Numeric {
		df.NumericShape[0] = len(v)
		break
	}
	df.NumericShape[1] = len(df.Numeric)

	for _, v := range df.String {
		df.StringShape[0] = len(v)
		break
	}
	df.StringShape[1] = len(df.String)
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
	df.String = map[string][]string{}
	df.Numeric = map[string][]float64{}

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
					df.String[col] = []string{}
					isNumeric = false
					break
				}
			}
		}
		if isNumeric {
			df.Numeric[col] = []float64{}
		}
	}
	stringCols := []string{}
	for k := range df.String {
		stringCols = append(stringCols, k)
	}

	// secondlly, store rows in String Columns as raw,
	// and store rows in Numeric Columns with converting into float
	allRowLength := len(allRows) - 1 // exclude colum row
	for i := 0; i < len(df.String) + len(df.Numeric); i++ {
		col := allRows[0][i]
		
		if helper.PythonicStrIfInList(col, stringCols) {

			tmpStringRows := make([]string, allRowLength)
			for j := 0; j < allRowLength; j++ {
				tmpStringRows[j] = allRows[j+1][i]
			}
			df.String[col] = tmpStringRows

		} else {

			tmpNumericRows := make([]float64, allRowLength)
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
	// count the number of null rows in each columns.
	// We handle math.NaN in NumericRows and "" in StringRows as null value
	numericNullCounts := make(map[string]int, df.NumericShape[1])
	stringNullCounts := make(map[string]int, df.StringShape[1])


	// initiate Nullounts maps
	for col := range df.Numeric {
		numericNullCounts[col] = 0
	}
	for col := range df.String {
		stringNullCounts[col] = 0
	}

	// count null rows 
	for key, values := range df.Numeric {
		for _, val := range values {
			if math.IsNaN(val) {
				numericNullCounts[key]++
			}
		}
	}
	for key, values := range df.String {
		for _, val := range values {
			if val == "" {
				stringNullCounts[key]++
			}
		}
	}

	stdoutInfo(df, numericNullCounts, stringNullCounts)


}


func stdoutInfo(df *DataFrame, numericNullCounts, stringNullCounts map[string]int) {

	orderedNumericKeys := []string{}
	orderedStringKeys := []string{}
	for key := range numericNullCounts {
		orderedNumericKeys = append(orderedNumericKeys, key)
	}
	sort.Strings(orderedNumericKeys)

	for key := range stringNullCounts {
		orderedStringKeys = append(orderedStringKeys, key)
	}
	sort.Strings(orderedStringKeys)

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
	for _, key := range orderedNumericKeys {
		fmt.Printf("%s,%d,%d,%s\n",
			key, df.NumericShape[0] - numericNullCounts[key], 
			numericNullCounts[key], reflect.TypeOf(df.Numeric[key][0]))
	}
	
	fmt.Println("")

	// String
	fmt.Printf("===== String columns (total %d columns) =====\n",
		df.StringShape[1])
	fmt.Println("name,non-null,null,dtype")
	for _, key := range orderedStringKeys {
		fmt.Printf("%s,%d,%d,%s\n",
			key, df.StringShape[0] - stringNullCounts[key], 
			stringNullCounts[key], reflect.TypeOf(df.String[key][0]))
	}

	
}

// Describe stdout the dataframe's statistical description of each columns.
// NumericColumns : count, mean, std, min, 25%, 50%, 75%, max, sum
// StringColumns : count, unique, top, freq
func (df *DataFrame) Describe() {

	numericStatistics := calclNumericStatistics(df)
	stringIntStatistics, stringStrStatistics := calcStringStatistics(df)

	stdoutDescribe(df, 
		numericStatistics, 
		stringIntStatistics, stringStrStatistics)

}

func calclNumericStatistics(df *DataFrame) map[string]map[string]float64 {
	// e.g. {"count" : {"Age" : 100, "..." : ...}, ...}
	numericStatistics := map[string]map[string]float64{
		"count" : map[string]float64{},
		"mean" : map[string]float64{},
		"std" : map[string]float64{},
		"min" : map[string]float64{},
		"25%" : map[string]float64{},
		"50%" : map[string]float64{},
		"75%" : map[string]float64{},
		"max" : map[string]float64{},
		"sum" : map[string]float64{},
	}

	// initiate maps 
	for _, val := range numericStatistics {
		for key := range df.Numeric {
			val[key] = 0.0
		}
	}

	for key, values := range df.Numeric {
		// firstlly, sort acsendingly
		sort.Sort(helper.AcsendingSort(values))

		calcFlag := true
		for i, val := range df.Numeric[key] {
			if math.IsNaN(val) {
				continue
			}

			// sum
			numericStatistics["sum"][key] += val

			// min, count
			if calcFlag {
				numericStatistics["min"][key] = val
				numericStatistics["count"][key] = float64(df.NumericShape[0] - i)
				calcFlag = false
				continue
			}

		}
		// max
		numericStatistics["max"][key] = df.Numeric[key][df.NumericShape[0]-1]
		// mean
		numericStatistics["mean"][key] = numericStatistics["sum"][key] / numericStatistics["count"][key]
		// std
		sigmaSquared := float64(0)
		for _, val := range df.Numeric[key] {
			if !math.IsNaN(val) {
				sigmaSquared += math.Pow(val -
								 numericStatistics["mean"][key], 2)
			}
		}
		numericStatistics["std"][key] = math.Sqrt(
			sigmaSquared / numericStatistics["count"][key])

		// percentiles
		numericStatistics["25%"][key] = calcPercentile(df.Numeric[key], 0.25)
		numericStatistics["50%"][key] = calcPercentile(df.Numeric[key], 0.50)
		numericStatistics["75%"][key] = calcPercentile(df.Numeric[key], 0.75)
		
	}

	return numericStatistics

}

func calcStringStatistics(df *DataFrame) (map[string]map[string]int, map[string]map[string]string) {
	stringIntStatistics := map[string]map[string]int{
		"count" : map[string]int{},
		"unique" : map[string]int{},
		"freq" : map[string]int{},
	}
	stringStrStatistics := map[string]map[string]string {
		"top" : map[string]string{},
	}

	// initiate maps 
	for key := range df.String {
		for _, val := range stringIntStatistics {
			val[key] = 0
		}
		for _, val := range stringStrStatistics {
			val[key] = ""
		}
	}

	for key, values := range df.String {
		sort.Strings(values)
		calcFlag := true
		tmpCounter := map[string]int{}

		for i, val := range df.String[key] {
			if val == "" {
				continue
			}

			// count (exclude empty entry : "" )
			if calcFlag {
				stringIntStatistics["count"][key] = df.StringShape[0] - i
				calcFlag = false
			}

			// for unique, top, freq
			tmpCounter[val]++

		}

		// unique, top, frep

		stringIntStatistics["unique"][key] = len(tmpCounter)
		mostCommonKeys, mostCommonValues := helper.PythonicStrCounterMostCommon(
			tmpCounter, 1)
		stringIntStatistics["freq"][key] = mostCommonValues[0]

		stringStrStatistics["top"][key] = mostCommonKeys[0]


	}

	return stringIntStatistics, stringStrStatistics

}

func calcPercentile(values []float64, percentile float64) float64 {
	nonNullValues := []float64{}
	for _, val := range values {
		if !math.IsNaN(val) {
			nonNullValues = append(nonNullValues, val)
		}
	}
	N := len(nonNullValues) - 1 // this is not a length, but distance from head to tail
	p := float64(N) * percentile
	q := int(math.Floor(p))
	r := p - float64(q)
	D := nonNullValues[q] + (nonNullValues[q+1] - nonNullValues[q]) * r // linear interporation
	return D
}

func stdoutDescribe(df *DataFrame, 
	numericStatistics map[string]map[string]float64,
	stringIntStatistics map[string]map[string]int,
	stringStrStatistics map[string]map[string]string) {

	orderedNumericMetric := []string{"count", "mean", "std", "min", "25%", "50%", "75%", "max", "sum"}
	orderedNumericColumns := []string{}
	for key := range df.Numeric {
		orderedNumericColumns = append(orderedNumericColumns, key)
	}
	sort.Strings(orderedNumericColumns)

	orderedStringIntMetric := []string{"count", "unique", "freq"}
	orderedStringStrMetric := []string{"top"}
	orderedStringColumns := []string{}
	for key := range df.String {
		orderedStringColumns = append(orderedStringColumns, key)
	}
	sort.Strings(orderedStringColumns)



	fmt.Printf("=== Numeric columns (total %d columns) =====\n", len(orderedNumericColumns))
	fmt.Printf("metric,%s\n", strings.Join(orderedNumericColumns, ","))
	for _, key := range orderedNumericMetric {
		stdoutString := key + ","
		for j, col := range orderedNumericColumns {
			stdoutString += fmt.Sprintf("%.3f", numericStatistics[key][col])
			if j < len(orderedNumericColumns) - 1 {
				stdoutString += ","
			}
		}
		fmt.Println(stdoutString)
	}

	fmt.Println("")

	fmt.Printf("=== String columns (total %d columns) =====\n", len(orderedStringColumns))
	fmt.Printf("metric,%s\n", strings.Join(orderedStringColumns, ","))
	for _, key := range orderedStringIntMetric {
		stdoutString := key + ","
		for j, col := range orderedStringColumns {
			stdoutString += fmt.Sprintf("%d", stringIntStatistics[key][col])
			if j < len(orderedStringColumns) - 1 {
				stdoutString += ","
			}
		}
		fmt.Println(stdoutString)
	}
	for _, key := range orderedStringStrMetric {
		stdoutString := key + ","
		for j, col := range orderedStringColumns {
			stdoutString += fmt.Sprintf("%s", stringStrStatistics[key][col])
			if j < len(orderedStringColumns) - 1 {
				stdoutString += ","
			}
		}
		fmt.Println(stdoutString)
	}

}
