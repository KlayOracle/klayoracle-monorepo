## Create a Self-Signed Root CA

```shell
openssl req -x509                                                   \
    -newkey rsa:4096                                                \
    -nodes                                                          \
    -days 3650                                                      \
    -keyout origineum_root_key.pem                                  \
    -out origineum_root_cert.pem                                    \
    -subj /C=AE/ST=DXB/L=DWTC/O=Origineum-FZCO/CN=origineum.com/    \
    -config ./openssl.cnf                                           \
    -extensions ca                                                  \
    -sha256
```

## Generate RSA Private key for CSR to extract Public Key

```shell
openssl genrsa -out origineum_node_key.pem 4096
```

## Create Certificate Signing Request (CSR)

```shell
openssl req -new                                                \
  -key origineum_node_key.pem                                   \
  -days 3650                                                    \
  -out origineum_csr.pem                                        \
  -subj /C=AE/ST=DXB/L=DWTC/O=Origineum-FZCO/CN=origineum.com/  \
  -config ./openssl.cnf                                         \
  -reqexts server
```

## Generate Self Signed Certificate With Root CA

```shell
openssl x509 -req                       \
  -in origineum_csr.pem                 \
  -CAkey origineum_root_key.pem         \
  -CA origineum_root_cert.pem           \
  -days 3650                            \
  -set_serial 1000                      \
  -out origineum_node_cert.pem          \
  -extfile ./openssl.cnf                \
  -extensions server                    \
  -sha256
```

## Resource
- [Create a Self Signed Certificate with OpenSSL](https://www.baeldung.com/openssl-self-signed-cert)
- [OpenSSL Quick Reference Guide](https://www.digicert.com/kb/ssl-support/openssl-quick-reference-guide.htm)
- [OpenSSL Cookbook](https://www.feistyduck.com/books/openssl-cookbook/)