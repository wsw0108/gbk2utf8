package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
)

func main() {
	var inplace bool
	flag.BoolVar(&inplace, "inplace", false, "convert inplace")
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("no input file")
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var r io.Reader
	var w io.Writer
	if inplace {
		b, err1 := io.ReadAll(f)
		f.Close()
		if err1 != nil {
			fmt.Println(err1)
			os.Exit(1)
		}
		r = bytes.NewReader(b)
		f, err1 = os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err1 != nil {
			fmt.Println(err1)
			os.Exit(1)
		}
		w = f
	} else {
		r = f
		w = os.Stdout
	}
	defer f.Close()

	dec := simplifiedchinese.GB18030.NewDecoder()
	enc := unicode.UTF8.NewEncoder()

	r = dec.Reader(r)
	w = enc.Writer(w)
	_, err = io.Copy(w, r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
