package main

import (
	"flag"
	"log"
	"os"
)

var (
	defaultSrc string
	reload     bool
	debug      bool
)

func init() {
	defaultSrc, _ = os.Getwd()
}

func main() {
	var path string
	flag.StringVar(&path, "path", defaultSrc, "Path to parse")
	flag.BoolVar(&reload, "reload", false, "Forced to copy files, ignores file exists")
	flag.BoolVar(&debug, "debug", false, "Debug mode, to show more log")

	flag.Parse()

	pkgs := AllImports(path)

	log.SetPrefix("[Maideps]")
	if debug {
		log.Println("-----------------------")
		log.Println("Get all dependencies:")
		for _, pkg := range pkgs {
			log.Println(pkg)
		}
		log.Println("-----------------------")
	}

	maideps := NewMaideps(pkgs, debug, reload)
	maideps.AddToVendor()
}
