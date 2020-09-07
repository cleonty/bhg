package main

import (
	"fmt"
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

func get(l *lua.LState) int {
	var (
		host     string
		port     uint64
		username string
		password string
		path     string
		resp     *http.Response
		err      error
		url      string
		client   *http.Client
		req      *http.Request
	)
	host = l.CheckString(1)
	port = uint64(l.CheckInt64(2))
	username = l.CheckString(3)
	password = l.CheckString(4)
	path = l.CheckString(5)
	url = fmt.Sprintf("http://%s:%d/%s", host, port, path)
	client = new(http.Client)
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		l.Push(lua.LNumber(0))
		l.Push(lua.LBool(false))
		l.Push(lua.LString(fmt.Sprintf("Unable to build GET request: %s", err)))
		return 3
	}
	if username != "" || password != "" {
		// Assume Basic Auth is required since user and/or password is set
		req.SetBasicAuth(username, password)
	}
	if resp, err = client.Do(req); err != nil {
		l.Push(lua.LNumber(0))
		l.Push(lua.LBool(false))
		l.Push(lua.LString(fmt.Sprintf("Unable to send GET request: %s", err)))
		return 3
	}
	l.Push(lua.LNumber(resp.StatusCode))
	l.Push(lua.LBool(false))
	l.Push(lua.LString(""))
	return 3
}
