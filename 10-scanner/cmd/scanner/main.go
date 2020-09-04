package main

import (
	"C"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"plugin"

	"github.com/cleonty/plugin-core/scanner"
)

const PluginsDir = "../../plugins/"

func main() {
	var (
		files []os.FileInfo
		err   error
		p     *plugin.Plugin
		n     plugin.Symbol
		check scanner.Checker
		res   *scanner.Result
	)
	if files, err = ioutil.ReadDir(PluginsDir); err != nil {
		log.Fatalln(err)
	}
	for idx := range files {
		fmt.Println("Found plugin: " + files[idx].Name())
		filePath := filepath.Join(PluginsDir, files[idx].Name())
		if p, err = plugin.Open(filePath); err != nil {
			log.Fatalf("unable to open %s: %v", filePath, err)
		}
		if n, err = p.Lookup("New"); err != nil {
			log.Fatalf("unable to find New() symbol: %v", err)
		}
		newFunc, ok := n.(func() scanner.Checker)
		if !ok {
			log.Fatalln("Plugin entry point is no good. Expecting: func New() scanner.Checker{ ... }")
		}
		check = newFunc()
		res = check.Check("10.0.1.20", 8080)
		if res.Vulnerable {
			log.Println("Host is vulnerable: " + res.Details)
		} else {
			log.Println("Host is NOT vulnerable")
		}
	}
}
