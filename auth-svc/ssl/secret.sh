kubectl create secret generic auth-rsa-key --from-file=privkey.pem --from-file=pubkey.pem -n istio-demo
