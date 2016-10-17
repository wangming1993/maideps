package main

import (
	"log"
	"path/filepath"
)

type Package string

func (this Package) dependencies() (packages []Package) {
	if this.IsStandard() {
		return
	}
	if !this.IsInternal() {
		packages = append(packages, this)
	}

	if this.InVendor() {
		absPath := filepath.Join(Pwd(), GO_VENDOR, this.name())
		vendorPkgs := AllImports(absPath)
		for _, name := range vendorPkgs {
			packages = append(packages, Package(name))
		}
		return
	}
	if this.InGopath() {
		gopathPkgs := AllImports(filepath.Join(GopathSrc(), this.name()))
		for _, name := range gopathPkgs {
			packages = append(packages, Package(name))
		}
		return
	}
	log.Fatalf("Missed package: %s, please use go get %s\n", this.name(), this.name())
	return
}

func (this Package) name() string {
	return string(this)
}

func (this Package) IsStandard() bool {
	return IsStandardPkg(this.name())
}

func (this Package) IsInternal() bool {
	return isInternal(this.name())
}

func (this Package) InVendor() bool {
	return InVendor(this.name())
}

func (this Package) InGopath() bool {
	return InGopath(this.name())
}
