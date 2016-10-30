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

func (this *Maideps) Add() error {
	for _, pkg := range this.Pkgs {
		src := filepath.Join(GopathSrc(), pkg)
		dst := filepath.Join(Pwd(), GO_VENDOR, pkg)
		if !this.Reload {
			if InVendor(pkg) {
				//Already exist in vendor folder, ignore
				if this.Debug {
					log.Printf("package: %s already exist in vendor folder, ignore \n", pkg)
				}
				continue
			}
		} else {
			rmdir(dst)
		}
		err := RewriteDir(src, dst)
		if err != nil {
			log.Printf("rewrite dir failed, error with: %v", err)
		}
	}
	return nil
}

func (this *Maideps) Delete() error {
	for _, pkg := range this.Pkgs {
		dst := filepath.Join(Pwd(), GO_VENDOR, pkg)
		err := rmdir(dst)
		if err != nil {
			if this.Debug {
				log.Printf("rewrite dir failed, error with: %v", err)
			}
			return err
		}
	}
	return nil
}
