datacenter  = "marvel-dc"
node_name   = "client"
retry_join  = ["consul-server"]
client_addr = "0.0.0.0"
bind_addr   = "0.0.0.0"
data_dir    = "/var/lib/consul"
log_level   = "debug"

# TLS settings
tls {
  defaults {
    ca_file         = "/consul/certs/consul-agent-ca.pem"
    verify_incoming = false # this is the client auth in mTLS that we were talking about
    verify_outgoing = true  # this is saying all outbound traffics from this server must use TLS
  }
}

auto_encrypt {
  tls = true # this is saying that client will request TLS certificate from server Connect CA
}
