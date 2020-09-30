#!/bin/bash

lib=libstart.a
go build -buildmode=c-archive -o ${lib} main.go
gcc -shared -pthread -o x.dll scratch.c ${lib} -lWinMM -lntdll -lWS2_32
gcc -pthread -o run.exe run.c ${lib} -lWinMM -lntdll -lWS2_32