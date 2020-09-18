#!/bin/sh

openssl req -nodes -x509 -newkey rsa:4096 -keyout serverKey.pem -out serverCrt.pem -days 365