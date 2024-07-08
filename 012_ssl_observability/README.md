# What Insights Can eBPF Provide into Real-Time SSL/TLS Encrypted Traffic and How?

## Prerequisites

Install dependencies using:
```
sudo apt install libbpf-dev llvm clang linux-tools-common gcc-multilib
```

## Run it

- Run eBPF Program 

```
go generate
go build
sudo ./ssl-ebpf
```

- Run HTTPS Server in `/test` directory:

```
go run main.go
```

- Then you can make requests using `curl`:

```
curl -X GET https://localhost:4445 --insecure --http1.1
curl -X POST https://localhost:4445 --insecure --http1.1
curl -X PUT https://localhost:4445 --insecure --http1.1
curl -X PATCH https://localhost:4445 --insecure --http1.1
curl -X DELETE https://localhost:4445 --insecure --http1.1
curl -X CONNECT https://localhost:4445 --insecure --http1.1
curl -X OPTIONS https://localhost:4445 --insecure --http1.1
curl -X TRACE https://localhost:4445 --insecure --http1.1
```
