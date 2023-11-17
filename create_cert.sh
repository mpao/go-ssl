#!/bin/bash

# CA cert
openssl req -new -newkey rsa:2048 -days 3650 -extensions v3_ca -nodes -x509 -sha256 -set_serial 0 \
-keyout cert/localCA.key -out cert/localCA.crt -subj "/CN=RootCA/"
# Certificate Signing Request
openssl req -new -newkey rsa:2048 -nodes -keyout cert/localhost.key -out cert/localhost.csr \
-config req.cnf -extensions v3_req
# Self-Signing
openssl x509 -req -sha256 -CAcreateserial -in cert/localhost.csr -days 3650 -CA cert/localCA.crt \
-CAkey cert/localCA.key -out cert/localhost.crt -extfile <(printf "subjectAltName=DNS:localhost")

# controlla presenza SAN
# openssl x509 -in cert/localhost.crt -noout -text | grep DNS
rm cert/localhost.csr

