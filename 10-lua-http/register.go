package main

import lua "github.com/yuin/gopher-lua"

const LuaHttpTypeName = "http"

func register(l *lua.LState) {
	mt := l.NewTypeMetatable(LuaHttpTypeName)
	l.SetGlobal("http", mt)
	// static attributes
	l.SetField(mt, "head", l.NewFunction(head))
	l.SetField(mt, "get", l.NewFunction(get))
}
