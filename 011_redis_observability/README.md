# Redis eBPF
This is a demo code, for showcasing Observability of the RESP protocol using eBPF. This code is inspired by Alaz, Kubernetes eBPF agent, developed by Anteon.

In order to try it out locally:

- Run eBPF program using
  ```
  go generate
  go build
  sudo ./redis
  ```
- Run the Redis Container using
  ```
  docker run --name redis-server -d --memory 4g --cpus 4 -p 6379:6379 redis
  ```
- Run client inside `/test` using 
  ```
  go run client.go
  ```
- In another shell, inspect eBPF program logs using
  ```
  sudo cat /sys/kernel/debug/tracing/trace_pipe
  ```
