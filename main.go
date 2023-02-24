package main

import (
	"fmt"
	"os"

	"github.com/M-Faheem-Khan/goTemplateEngine-Test/engine"
)

func main() {
	htmlFile, err := os.ReadFile("/workspaces/goTemplateEngine-Test/templates/index.html")
	if err != nil {
		panic(err)
	}

	vars := map[string]string{
		"title":           "POC - Template Engine written in Go",
		"sample_variable": "Hello World - I'm rendered via Golang template engine",
	}

	s := engine.Parse(htmlFile, vars)
	fmt.Println(s)
}

// EOF
