package main

import (
	"fmt"
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

func head(l *lua.LState) int {
	var (
		host string
		port uint64
		path string
		resp *http.Response
		err  error
		url  string
	)
	host = l.CheckString(1)
	port = uint64(l.CheckInt64(2))
	path = l.CheckString(3)
	url = fmt.Sprintf("http://%s:%d/%s", host, port, path)
	if resp, err = http.Head(url); err != nil {
		l.Push(lua.LNumber(0))
		l.Push(lua.LBool(false))
		l.Push(lua.LString(fmt.Sprintf("Request failed: %s", err)))
		return 3
	}
	l.Push(lua.LNumber(resp.StatusCode))
	l.Push(lua.LBool(resp.Header.Get("WWW-Authenticate") != ""))
	l.Push(lua.LString(""))
	return 3
}
