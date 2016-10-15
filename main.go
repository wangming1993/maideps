package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	defaultSrc string
)

func init() {
	defaultSrc, _ = os.Getwd()
}

func main() {
	var path string
	flag.StringVar(&path, "path", defaultSrc, "Path to parse")
	fmt.Println(path)

	pkgs := AllImports(path)
	fmt.Println("all packages", pkgs)

	AddToVendor(pkgs)
}
