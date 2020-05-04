# csvCount

counts the rows in a csv file

## Getting Started

Build the executable with `go build`.

`./csvCount --help` will give you this:
```shell script
Usage of csvCount:
  -f string
        path to input file instead of stdin
  -s string
        separator character (defaults to a comma) (default ",")
```
## Examples
* count the number of rows in the calls file:
```shell script
./csvCount -f="./calls.csv"
```

## Built With

* [Golang](https://golang.org/) - The go language
