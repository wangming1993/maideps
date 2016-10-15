package main

import (
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
)

const (
	GO_VENDOR = "vendor"
)

func IsStandardPkg(pkgName string) bool {
	if pkgName == "C" {
		pkgName = "runtime/cgo"
	}
	root := GorootSrc()
	absPath := filepath.Join(root, pkgName)
	return exists(absPath)
}

func InVendor(pkgName string) bool {
	current := Pwd()
	absPath := filepath.Join(current, GO_VENDOR, pkgName)
	return exists(absPath)
}

func InGopath(pkgName string) bool {
	gopath := GopathSrc()
	absPath := filepath.Join(gopath, pkgName)
	return exists(absPath)
}

func Pwd() string {
	dir, _ := os.Getwd()
	return dir
}

func GopathSrc() string {
	srcDirs := GoSrcDirs()
	return srcDirs[1]
}

func GoSrcDirs() []string {
	return build.Default.SrcDirs()
}

func GorootSrc() string {
	srcDirs := GoSrcDirs()
	return srcDirs[0]
}

func runCmd(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	buf, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func uniq(a []string) []string {
	var s string
	var i int
	if !sort.StringsAreSorted(a) {
		sort.Strings(a)
	}
	for _, t := range a {
		if t != s {
			a[i] = t
			i++
			s = t
		}
	}
	return a[:i]
}
