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

Then let's run our HTTPS server

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

# mTLS setup

So far we have seen how TLS works. Let's look at mTLS. `m` here stands for Mutual. In TLS setup, only the client authenticates the server. By verifying the server certificate with its CA store.

With mTLS, the server will request client to provide and verify the certificate as well.

Let's regenerate the certs

```
chmod +x script.sh
./script.sh
```

And run our updated Go program

```
go run main.go
```

If we verify with previous curl command used in [TLS setup](#tls-setup), it not gonna work now

```
curl --cacert ./certs/ca.pem https://0:8443

curl: (35) error:1401E412:SSL routines:CONNECT_CR_FINISHED:sslv3 alert bad certificate
```

We will have to supply client certificate for curl to use

```
curl --http1.1 --cert ./certs/client.pem --key ./certs/client-key.pem --cacert ./certs/ca.pem https://0:8443

This is the way!
```

# Learning Point

When troublehshoot issue related to TLS, let's ask these questions:
- Who is the client?
- Who is the server?
- Where is the CA certificate? Server certificate? Client certificate?