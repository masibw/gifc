package main

import (
	"github.com/masibw/gifc"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(gifc.Analyzer) }

