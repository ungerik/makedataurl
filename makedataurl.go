package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ogier/pflag"
	"github.com/vincent-petithory/dataurl"
)

var (
	outFile string
	serveAt string
)

const indexHTML = `<!DOCTYPE html><html>
<head><meta charset="utf-8"/><title>%s</title></head>
<body>
<h1>%s</h1>
<img src="%s"/>
<hr/>
<code style="word-wrap:break-word">%s</code>
`

func main() {
	pflag.StringVar(&outFile, "out", "", "save data-URL to file")
	pflag.StringVar(&serveAt, "serve", "", "serve HTML preview at this HTTP address")

	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of makedataurl <filename>")
		pflag.PrintDefaults()
	}

	pflag.Parse()

	if len(pflag.Args()) == 0 {
		pflag.Usage()
		return
	}

	filename := pflag.Args()[0]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	url := dataurl.EncodeBytes(data)

	if outFile != "" {
		f, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0660)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(f, url)
		f.Close()
	} else if serveAt == "" {
		fmt.Println(url)
	}

	if serveAt != "" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, indexHTML, filename, filename, url, url)
		})

		fmt.Println("serving at:", serveAt)
		err = http.ListenAndServe(serveAt, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
