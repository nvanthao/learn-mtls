#! /bin/bash

cd ./certs

echo "Generate self-signed CA certificate..."
cfssl genkey -initca csr.json | cfssljson -bare ca

echo "Generate server certificate..."
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -hostname=0.0.0.0 server.json | cfssljson -bare server
