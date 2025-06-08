package main

import (
	"flag"
)

func main() {
	generateDocs := flag.Bool("docs", false, "generate openapi docs")
	flag.Parse()
	if *generateDocs {
		doc := BuildDocument()
		doc.SaveAsJson("openapi.json")
		doc.SaveAsYaml("openapi.yaml")
	}
}
