#!/bin/sh

#fix for: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0
export GODEBUG="x509ignoreCN=0"
go run main.go