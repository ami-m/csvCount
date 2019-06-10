package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

type Record []string

func getCsvReader(r io.Reader, p runParams) *csv.Reader {
	res := csv.NewReader(r)
	res.Comma = p.separator
	return res
}

func countRows(r *csv.Reader) int64 {
	var res int64

	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		res++
	}
	return res
}

func main() {
	params := initParams()
	r := getCsvReader(getRawReader(params), params)

	fmt.Println(countRows(r))
}
