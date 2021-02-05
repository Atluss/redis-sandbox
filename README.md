My Redis Sandbox

This my sandbox for Redis.

### Run Redis

```bash
cd docker
docker-compose up
```


port | desc
---|---
6379 | standard port to access redis (no password by default)

### Run test

```
go get -u ./...
go mod vendor
```

### Links
[Docker Hub](https://hub.docker.com/_/redis/) \
[Docs](https://redis.io/documentation)
