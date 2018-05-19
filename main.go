package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

func main() {
	os.Exit(Run(os.Args[1:]))
}
