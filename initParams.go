package main

import "flag"

type runParams struct {
	file      string
	separator rune
}

func initParams() runParams {
	var res runParams
	var file string
	var separator string

	flag.StringVar(&file, "f", "", "path to input file instead of stdin")
	flag.StringVar(&separator, "s", ",", "separator character (defaults to a comma)")

	flag.Parse()

	res.file = file

	separatorRunes := []rune(separator)
	res.separator = separatorRunes[0]

	return res
}
