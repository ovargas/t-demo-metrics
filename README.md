# Temporal PoC

A simple experiment with temporal.io

### Requirements

* Docker
* Docker Compose
* Golang 1.19

### How to use it

Pull all required images executing

```shell
docker-compose pull
```

Execute all services

```shell
docker-compose up -d
```

Access the console

* Temporal UI: [http://localhost:8088](http://localhost:8088)
* Prometheus:  [http://localhost:9090](http://localhost:9090)
* Grafana:  [http://localhost:3000](http://localhost:3000)

To create workflows execute

```shell
go run ./scripts/dummy/main.go
```

### Tested environment

* macOS 12.6 with M1 processor
* go version go1.19.1 darwin/arm64
* Docker

```
Client:
 Cloud integration: v1.0.29
 Version:           20.10.17
 API version:       1.41
 Go version:        go1.17.11
 Git commit:        100c701
 Built:             Mon Jun  6 23:04:45 2022
 OS/Arch:           darwin/arm64
 Context:           default
 Experimental:      true

Server: Docker Desktop 4.12.0 (85629)
 Engine:
  Version:          20.10.17
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.17.11
  Git commit:       a89b842
  Built:            Mon Jun  6 23:01:01 2022
  OS/Arch:          linux/arm64
  Experimental:     false
 containerd:
  Version:          1.6.8
  GitCommit:        9cd3357b7fd7218e4aec3eae239db1f68a5a6ec6
 runc:
  Version:          1.1.4
  GitCommit:        v1.1.4-0-g5fd4c4d
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
```

### References

* [https://github.com/temporalio/dashboards](https://github.com/temporalio/dashboards)
* [https://github.com/tsurdilo/my-temporal-dockercompose](https://github.com/tsurdilo/my-temporal-dockercompose)
* [https://github.com/temporalio/samples-go/tree/main/metrics](https://github.com/temporalio/samples-go/tree/main/metrics)
* [https://github.com/temporalio/docker-compose](https://github.com/temporalio/docker-compose)