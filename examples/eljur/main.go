package main

import (
	"backend/models/document"
	"flag"
)

func main() {
	generateDocs := flag.Bool("docs", false, "generate openapi docs")
	flag.Parse()
	if *generateDocs {
		document.BuildDocument().SaveAsJson("openapi.json")
	}
}
