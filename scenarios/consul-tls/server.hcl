datacenter  = "marvel-dc"
server      = true
bootstrap   = true
node_name   = "server"
client_addr = "0.0.0.0"
bind_addr   = "0.0.0.0"
data_dir    = "/var/lib/consul"
log_level   = "debug"

# TLS settings
tls {
  defaults {
    ca_file         = "/consul/certs/consul-agent-ca.pem"
    cert_file       = "/consul/certs/marvel-dc-server-consul-0.pem"
    key_file        = "/consul/certs/marvel-dc-server-consul-0-key.pem"
    verify_incoming = true # this is the client auth in mTLS that we were talking about
    verify_outgoing = true # this is saying all outbound traffics from this server must use TLS
  }
}

auto_encrypt {
  allow_tls = true # this is saying that server can return a TLS certificate from Connect CA to the client
}
