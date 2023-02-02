# Objective
Understand TLS with Golang

# Steps
Check out the Git branch accordingly and follow the instructions in the README file

```
git checkout plain-setup
```

# Plain Setup
We gonna write a simple HTTP server that just says Hi!
No setup script is required, just run

```
go run main.go
```

We can open Wireshark, capture `Loopback` interface and use `http` filter 

From another terminal

```
curl 0:8080

This is the way!
```

# TLS Setup
Next we gonna add TLS to our server.

```
git checkout tls-setup
```

We first generate the required certificates for our server. And we will use [cfssl](https://github.com/cloudflare/cfssl) for convenience

```
chmod +x script.sh
./script.sh
```

Then let's run our HTTPs server

```
go run main.go

Starting HTTPS server on port: 8443 ...
```

Let's verify our TLS setup with curl

```
curl 0:8443

Client sent an HTTP request to an HTTPS server
```

```
curl https://0:8443

curl: (60) SSL certificate problem: unable to get local issuer certificate
More details here: https://curl.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not
establish a secure connection to it. To learn more about this situation and
how to fix it, please visit the web page mentioned above.
```

This is understandable as our CA certificate is self-signed and not in the CA trust store that curl knows. We will help to support the CA certificate to curl

```
curl --cacert ./certs/ca.pem https://0:8443

This is the way!
```