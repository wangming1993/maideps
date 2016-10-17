package main

import (
	"flag"
	"log"
)

var (
	reload bool
	debug  bool
	pkg    string
)

func init() {}

func main() {
	flag.BoolVar(&reload, "reload", false, "Forced to copy files, ignores file exists")
	flag.BoolVar(&debug, "debug", false, "Debug mode, to show more log")
	flag.StringVar(&pkg, "import", "", "Specific one import package name, only find its dependency")

	flag.Parse()

	var pkgs []string
	if "" == pkg {
		pkgs = AllImports(Pwd())
	} else {
		dependencies := Package(pkg).dependencies()
		for _, dep := range dependencies {
			pkgs = append(pkgs, dep.name())
		}
	}

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
