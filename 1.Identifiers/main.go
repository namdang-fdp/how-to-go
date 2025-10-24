package main

import (
	"fmt"

	"identifiers/export"
)

// who is your daddy
// exported variable
var ExportedVariable = "Hello, World!"

func main() {
	fmt.Println(ExportedVariable)
	fmt.Println(export.AnotherExportedVariable)
}
