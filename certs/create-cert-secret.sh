#!/bin/sh

kubectl create -n istio-system secret tls istio-ingressgateway-certs --key istio.demo.key --cert istio.demo.crt

