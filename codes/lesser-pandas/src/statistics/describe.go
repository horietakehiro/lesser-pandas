package main

import (
	"io"
	"os"
	"encoding/csv"
	"fmt"
)

func main()  {
	reader, _ := os.Open("sample.csv")
	r := csv.NewReader(reader)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line)

	}
}
