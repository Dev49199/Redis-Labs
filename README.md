# Redis + Go Hands-on Labs 🚀

This repository contains **small, practical Redis labs written in Go**.
The goal is to understand Redis concepts through **tiny, focused examples**
instead of building a large project.

Each lab is independent and can be run by combining:
- a shared Redis client file
- one lab file at a time

---

## 🧠 What you will practice

- Redis Strings (atomic counters, TTL)
- Redis Hashes (object-style storage)
- Redis Sets (online user tracking)
- Redis Sorted Sets (leaderboards)
- Basic distributed locking
- Redis Streams (event-style processing)

---

## ⚙️ Prerequisites

- Go 1.21+
- Docker

---

## 🐳 Run Redis locally (Docker)

Start Redis using a single Docker command:

```bash
docker run -d \
  --name redis-go-labs \
  -p 6379:6379 \
  redis:7
```

## Other Commands needed

```bash
docker ps

docker exec -it redis-go-labs redis-cli


go run redisClient.go lab1_counter.go
go run redisClient.go lab2_hash.go
go run redisClient.go lab3_set.go
go run redisClient.go lab6_streams.go
etc..
```
