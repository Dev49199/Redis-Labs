# Redis + Go Hands-on Labs 🚀

This repository contains **small, practical Redis labs written in Go**.
The goal is to understand Redis concepts through **tiny, focused examples**
instead of building a large project.

Each lab demonstrates **one Redis concept** that you can modify and run locally.

---

## 🧠 What you will practice

- Redis Strings (atomic counters, TTL)
- Redis Hashes (object storage)
- Redis Sets (online users)
- Redis Sorted Sets (leaderboards)
- Redis anti-pattern awareness
- Redis behavior under small failure scenarios

---

## ⚙️ Prerequisites

- Go 1.21+
- Docker

> If Docker is not installed, the setup script will guide you.

---

## 🚀 Setup (One command)

```bash
git clone https://github.com/<your-username>/redis-go-labs.git
cd redis-go-labs
./setup.sh
