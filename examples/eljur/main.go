package main

import (
	"flag"
)

func main() {
	generateDocs := flag.Bool("docs", false, "generate openapi docs")
	flag.Parse()
	if *generateDocs {
		BuildDocument().SaveAsJson("openapi.json")
	}
}
