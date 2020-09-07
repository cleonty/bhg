package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	lua "github.com/yuin/gopher-lua"
)

const PluginsDir = "plugins"

func main() {
	var (
		l     *lua.LState
		files []os.FileInfo
		err   error
		f     string
	)
	l = lua.NewState()
	defer l.Close()
	register(l)
	if files, err = ioutil.ReadDir(PluginsDir); err != nil {
		log.Fatalln(err)
	}
	for idx := range files {
		fmt.Println("Found plugin: " + files[idx].Name())
		f = fmt.Sprintf("%s/%s", PluginsDir, files[idx].Name())
		if err := l.DoFile(f); err != nil {
			log.Fatalln(err)
		}
	}
}
