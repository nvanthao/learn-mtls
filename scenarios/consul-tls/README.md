# Objective

Understand TLS setup in Consul

# Steps

- Go to `/certs` directory
- Generate CA certificate

```
consul tls ca create
```

- Generate server certificate

```
consul tls cert create -server -dc marvel-dc
```

- Bring up the cluster

```
docker compose up -d
```

- Capturing traffic on client container

```
docker run --rm --net container:ptfe_atlas nicolaka/netshoot tcpdump -w - > packets.pcap
```