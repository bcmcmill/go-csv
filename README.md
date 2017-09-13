CSV
===
[![Build Status](https://secure.travis-ci.org/JensRantil/go-csv.png?branch=master)](http://travis-ci.org/JensRantil/go-csv) [![Go Report Card](https://goreportcard.com/badge/github.com/JensRantil/go-csv)](https://goreportcard.com/report/github.com/JensRantil/go-csv) [![GoDoc](https://godoc.org/github.com/JensRantil/go-csv?status.svg)](https://godoc.org/github.com/JensRantil/go-csv)

A Go [CSV](https://en.wikipedia.org/wiki/Comma-separated_values) implementation
inspired by [Python's CSV module](https://docs.python.org/2/library/csv.html).
It supports various CSV dialects (see below) and is fully backward compatible
with the [`encoding/csv`](http://golang.org/pkg/encoding/csv/) package in the
Go standard library.

Examples
--------

Writing
~~~~~~~
Here's a basic writing example::

    f, err := os.Create("output.csv")
    checkError(err)
    defer func() {
      err := f.Close()
      checkError(err)
    }
    w := NewWriter(f)
    w.Write([]string{
      "a",
      "b",
      "c",
    })
    w.Flush()
    // output.csv will now contains the line "a b c" with a trailing newline.

Reading
~~~~~~~
Here's a basic reading example::

    f, err := os.Open('myfile.csv')
    checkError(err)
    defer func() {
      err := f.Close()
      checkError(err)
    }

    r := NewReader(f)
    for {
      fields, err := r.Read()
      if err == io.EOF {
        break
      }
      checkOtherErrors(err)
      handleFields(fields)
    }


To automatically detects the CSV delimiter conforming to the specifications outlined on the on the [Wikipedia article][csv]. Looking through many CSV libraries code and discussion on the stackoverflow, finding that their CSV delimiter detection is limited or incomplete or containing many unneeded features. Hoping this can people solve the CSV delimiter detection problem without importing extra overhead.

[csv]: http://en.wikipedia.org/wiki/Comma-separated_values

## Usage

    package main
    
    import (
    	"github.com/eltorocorp/go-csv/detector"
    	"os"
    	"fmt"
    )
    
    func main()  {
    	detector := detector.New()
    
    	file, err := os.OpenFile("example.csv", os.O_RDONLY, os.ModePerm)
    	if err != nil {
    		os.Exit(1)
    	}
    	defer file.Close()
    
    	delimiters := detector.DetectDelimiter(file, '"')
    	fmt.Println(delimiters)
    }

CSV dialects
------------
To modify CSV dialect, have a look at `csv.Dialect`,
`csv.NewDialectWriter(...)` and `csv.NewDialectReader(...)`. It supports
changing:

* separator/delimiter.
* quoting modes:
  * Always quote.
  * Never quote.
  * Quote when needed (minimal quoting).
  * Quote all non-numerical fields.
* line terminator.
* how quote character escaping should be done - using double escape, or using a
  custom escape character.

Have a look at [the
documentation](http://godoc.org/github.com/JensRantil/go-csv) `csv_test.go` for
example on how to use these. All values above have sane defaults (that makes
the module behave the same as the `csv` module in the Go standard library).

Documentation
-------------
Package documentation can be found
[here](http://godoc.org/github.com/JensRantil/go-csv)



