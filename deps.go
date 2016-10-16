package main

import (
	"log"
	"path/filepath"
)

type Maideps struct {
	Pkgs   []string
	Debug  bool
	Reload bool
}

func NewMaideps(imports []string, debug, reload bool) *Maideps {
	return &Maideps{
		Pkgs:   imports,
		Debug:  debug,
		Reload: reload,
	}
}

func (this *Maideps) AddToVendor() error {
	for _, pkg := range this.Pkgs {
		if !this.Reload {
			if InVendor(pkg) && this.Debug {
				//Already exist in vendor folder, ignore
				log.Printf("package: %s already exist in vendor folder, ignore \n", pkg)
				continue
			}
		}
		src := filepath.Join(GopathSrc(), pkg)
		dst := filepath.Join(Pwd(), GO_VENDOR, pkg)
		err := RewriteDir(src, dst)
		if err != nil {
			log.Printf("rewrite dir failed, error with: %v", err)
		}
	}
	return nil
}
