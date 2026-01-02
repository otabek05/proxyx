# ProxyX

ProxyX is a high‑performance, configuration‑driven reverse proxy and static file server written in **Go**, inspired by **Nginx**. It provides:

* ✅ Reverse Proxy
* ✅ Static File Serving
* ✅ TLS/HTTPS with Certbot
* ✅ Load Balancing (Round‑Robin)
* ✅ Health Checking
* ✅ **Per‑Domain Rate Limiting**
* ✅ Declarative YAML Configuration
* ✅ Powerful Interactive CLI Tool

---


## Installation

ProxyX supports the following platforms:

- **Linux**: Debian-based (`.deb`) and RPM-based (`.rpm`) distributions  
- **Darwin**: macOS

For detailed installation instructions, see the [Installation Guide](docs/INSTALL.md).

---

## ProxyX Configuration Guide

### ProxyX uses a Kubernetes-style YAML configuration format
### to define domains, TLS, rate limits, and routes declaratively.
For full command reference and usage examples, see: [docs/CLI.md](docs/CONFIGURATION.md)

---

## TLS & HTTPS with Certbot

ProxyX integrates with **Certbot** to automatically issue and manage Let's Encrypt TLS certificates.

### ✅ Requirements

You **must install Certbot manually**:

```bash
sudo dnf install certbot   # RHEL / Amazon Linux
sudo apt install certbot   # Ubuntu / Debian
```

---

### ✅ Interactive Certificate Issuance

```bash
sudo proxyx certs
```

ProxyX will **prompt interactively**:

* ✅ Domain name
* ✅ Email address for Let's Encrypt

Then ProxyX will:

* Request the certificate
* Store the cert & key
* Automatically wire it into your configuration

---
## CLI Tool

ProxyX includes a full lifecycle management CLI.  
For full command reference and usage examples, see: [docs/CLI.md](docs/CLI.md)

---

## System Service & Ports

ProxyX automatically installs itself as a **Linux system service (`proxyx.service`)** and is designed to run as a **production-grade daemon**.

### ✅ Service Features

* ✅ Runs as `proxyx` system service
* ✅ Automatically starts on system boot
* ✅ Automatically restarts if the server turns off/on
* ✅ Automatically restarts on crash or failure

### ✅ Network Ports

* ✅ **Port 80** → HTTP traffic
* ✅ **Port 443** → HTTPS (TLS via Certbot)

> ⚠️ ProxyX requires **root (sudo)** access to bind to ports **80 and 443**.

---

## Architecture Overview

* Go **fasthttp** server for high-performance, concurrent HTTP handling
* Custom YAML parser for flexible configuration
* Reverse proxy engine
* Health checker
* Certbot shell integration
* Middleware pipeline:
  * Request Logger
  * Per‑Domain Rate Limiter
  * Load Balancer
  * Health Checker


## Use Cases

* API Gateway
* Static website hosting
* Internal microservice router
* Development reverse proxy
* Production HTTPS entrypoint
---

## Security Features

* HTTPS with Let's Encrypt
* Per‑domain rate limiting
* Backend health validation
* Mandatory TLS for production

---

## Author

Developed by **Otabek** — Go Backend Developer

---

## 📄 License

MIT License
