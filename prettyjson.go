package main

import (
	"os"
	"io"
	"log"
	"flag"
	"fmt"
	"encoding/json"
)

const usage = `usage: prettyjson [-i indent] file

prettyjson pretty-prints JSON input from file. If file is absent, prettyjson
reads from stdin. prettyjson will indent using a tab character by default,
otherwise the string specified by -i.
`

func main() {
	indent := flag.String("i", "\t", "indent string")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
	}
	flag.Parse()
	var r io.Reader
	if flag.NArg() == 0 {
		r = os.Stdin
	} else if flag.NArg() == 1 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		r = file
	} else {
		flag.Usage()
		os.Exit(1)
	}
	var v interface{}
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		log.Fatal(err)
	}
	json, err := json.MarshalIndent(v, "", *indent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}
