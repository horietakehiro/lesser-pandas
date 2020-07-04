package main


import (
	"fmt"

	"lpandas/core"
)

func main() {
	df := core.DataFrame{}
	filePath := "../test_datas/titanic-sample.csv"

	df.ReadCsv(filePath)

	fmt.Println(df.Info())
	fmt.Println(df.Info().Values["non-null"])


}

// type Series struct {
// 	dtype string
// 	values []interface{}
// 	name string
// 	index []string
// }

// type DataFrame struct {
// 	columns []string // for ordering output
// 	index []string // for ordering output
// 	values map[string]*Series
// }


// func (df *DataFrame) readCsv(file string) {
// 	reader, err := os.Open(file)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer reader.Close()

// 	csvReader := csv.NewReader(reader)

// 	// read each lines and store each seriese
// 	initFlag := true
// 	var serieses []Series
// 	index := 0
// 	for {
// 		row, err := csvReader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			panic(err)
// 		}

// 		// store first row as columns
// 		// determine the number of series obj to create
// 		// link each seriese with the dataframe
// 		if initFlag {
// 			df.columns = row
// 			df.values = map[string]*Series{}
// 			serieses = make([]Series, len(row))
// 			for i, r := range row {
// 				serieses[i].name = r
// 				df.values[r] = &serieses[i]
// 			}
// 			initFlag = false
// 			continue
// 		}

// 		// store rest of all rows as string(firstlly)
// 		for i, r := range row {
// 			serieses[i].values = append(serieses[i].values, r)
// 			serieses[i].index = append(serieses[i].index, fmt.Sprintf("%d", index))
// 		}
// 		df.index = append(df.index, fmt.Sprintf("%d", index))
// 		index++
// 	}

// 	// convert each series's values into proper dtypes
// 	for i, sr := range serieses {
// 		serieses[i] = convertDtype(sr)
// 	}

// }

// func convertDtype(sr Series) Series {

// 	// convert the whole series into either string or float64
// 	tmpFloatArray := make([]float64, len(sr.values))
// 	for i, val := range sr.values {
// 		if val == "" {
// 			tmpFloatArray[i] = math.NaN()
// 			continue
// 		}
// 		if str, ok := val.(string); ok { // must need type assertion
// 			float, err := strconv.ParseFloat(str, 64)
// 			if err != nil {
// 				// must be a string array
// 				sr.dtype = "string"
// 				return sr
// 			}

// 			tmpFloatArray[i] = float
// 		}
// 	}
// 	newSeries := Series{name : sr.name, index : sr.index, dtype : "float64"}
// 	for _, val := range tmpFloatArray {
// 		newSeries.values = append(newSeries.values, val)
// 	}
// 	return newSeries
// }

// func main() {
// 	df := DataFrame{}
// 	file := "../test_datas/titanic-sample.csv"

// 	df.readCsv(file)
// 	// fmt.Println(df)
// 	fmt.Println(df.values["Age"].dtype)
// 	fmt.Println(df.values["Sex"].dtype)

// }