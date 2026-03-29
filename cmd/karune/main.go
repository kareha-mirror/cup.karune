package main

import (
	"fmt"
	"log"
	"os"

	"tea.kareha.org/lab/karune/internal/fetch"
	"tea.kareha.org/lab/karune/internal/parser"
	"tea.kareha.org/lab/karune/internal/render"
	"tea.kareha.org/lab/karune/internal/util"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s PATH\n", os.Args[0])
		return
	}
	u, err := util.ResolvePath(os.Args[1])
	if err != nil {
		log.Fatalf("failed to resolve path: %v", err)
	}

	htmlStr, err := fetch.Get(u)
	if err != nil {
		log.Fatalf("failed to fetch page: %v", err)
	}

	doc, err := parser.Parse(htmlStr)
	if err != nil {
		log.Fatalf("failed to parse page: %v", err)
	}

	var lines []string
	parser.ExtractText(doc, &lines)

	render.Print(lines)
}
