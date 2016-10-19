package main

import (
	"go/build"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func listFiles(src, pattern string) []string {
	files, err := filepath.Glob(filepath.Join(src, pattern))
	if nil != err {
		log.Println(err)
		return nil
	}
	return files
}

func fileName(file string) string {
	return filepath.Base(file)
}

func dotPackage() (*build.Package, error) {
	dir, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}
	return build.ImportDir(dir, build.FindOnly)
}

func GetImports(file string) []string {
	//ignore test file
	testFile := strings.HasSuffix(file, "_test.go")
	if testFile {
		//log.Println("test file", file)
		return nil
	}

	pf, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.ImportsOnly|parser.ParseComments)
	if err != nil {
		log.Println(err)
		return nil
	}

	var pkgs []string
	for _, is := range pf.Imports {
		name, err := strconv.Unquote(is.Path.Value)
		if err != nil {
			log.Println("error:", err.Error())
			return nil
		}
		if IsStandardPkg(name) {
			continue
		}
		pkgs = append(pkgs, name)
	}
	return pkgs
}

func AllImports(dir string) []string {
	files := listFiles(dir, "*.go")
	var pkgs []string
	for _, file := range files {
		imports := GetImports(file)
		pkgs = append(pkgs, imports...)
	}
	log.Println("pkgs", pkgs)
	pkgs = GetRecursions(uniq(pkgs))
	return uniq(pkgs)
}

func GetRecursions(pkgs []string) []string {
	log.Println("recursions:", pkgs)
	var recursionPkgs []string
	for _, pkg := range pkgs {
		log.Println("range ", pkg)
		dependencies := Package(pkg).dependencies()
		if len(dependencies) > 0 {
			for _, dep := range dependencies {
				log.Println("dep:", dep.name())
				recursionPkgs = append(recursionPkgs, dep.name())
			}
		}
	}
	return uniq(recursionPkgs)
}

func isInternal(pkg string) bool {
	dp, err := dotPackage()
	if err != nil {
		return false
	}
	return strings.HasPrefix(pkg, dp.ImportPath)
}

func RewriteDir(src, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, "_test.go") {
			//ignore test fo file
			continue
		}
		if name == "README.md" || name == "LICENSE" || strings.HasSuffix(name, ".go") {
			err := RewriteFile(filepath.Join(src, name), filepath.Join(dst, name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RewriteFile(src, dst string) error {
	dir := fileDir(dst)
	if !exists(dir) {
		mkdir(dir)
	}
	return CopyFile(dst, src, 0755)
}

func fileDir(file string) string {
	return filepath.Dir(file)
}

func mkdir(dir string) error {
	return os.MkdirAll(dir, 0777)
}
func rmdir(dir string) error {
	return os.RemoveAll(dir)
}

// CopyFile copies the contents from src to dst using io.Copy.
// If dst does not exist, CopyFile creates it with permissions perm;
// otherwise CopyFile truncates it before writing.
func CopyFile(dst, src string, perm os.FileMode) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()
	_, err = io.Copy(out, in)
	return
}
