#!/bin/sh

kubectl create configmap auth-grpc-config -n istio-demo --from-literal=grpc-addr=:8000 --from-literal=redis-url=redis://:Winv3HF60f@redis5-master.redis.svc.cluster.local:6379/0 --from-literal=ttl=300s
