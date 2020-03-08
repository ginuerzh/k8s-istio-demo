#!/bin/sh

openssl genrsa -out privkey.pem 2048

openssl rsa -in privkey.pem -outform PEM -pubout -out pubkey.pem
