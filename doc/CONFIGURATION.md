# ProxyX Configuration Guide

This document explains how to configure **ProxyX** using a single YAML file.
All fields are shown in YAML format with inline comments for clarity.

---

## 📄 Full Example Configuration

```yaml
apiVersion: proxyx.io/v1        # Configuration schema version
kind: ProxyConfig               # Resource type

metadata:
  name: local-proxy              # Human-readable configuration name
  namespace: default             # Optional logical grouping

spec:
  domain: localhost              # Domain this configuration applies to

  tls:
    certFile: /path/to/certs/server.crt   # TLS certificate file path
    keyFile:  /path/to/certs/server.key   # TLS private key file path

  rateLimit:
    requests: 1000               # Maximum number of requests allowed in the window
    windowSeconds: 2             # Sliding window size in seconds

  routes:
    - name: api-route             # Route identifier
      path: /api/v1/**            # URL path pattern
      type: ReverseProxy          # Route type
      reverseProxy:
        servers:
          - url: http://127.0.0.1:8080   # Backend server
          - url: http://127.0.0.1:8081   # Backend server

    - name: static-files
      path: /**                   # Match all remaining paths
      type: Static
      static:
        root: /var/www/app        # Directory to serve static files from

    - name: websocket-route
      path: /ws/**                # WebSocket endpoint path
      type: Websocket
      websocket:
        url: ws://127.0.0.1:9000/ws  # WebSocket backend URL

```

## Metadata

```yaml
metadata:
  name: example-name       # Used for identification and CLI operations
  namespace: default       # Logical grouping for multiple configurations

```

* **name** uniquely identifies the configuration within the system.
* **namespace** allows grouping multiple configurations logically.


### Used by CLI commands:

```yaml
proxyx delete <name>
proxyx configs
```
---

## 🌐 Domain

```yaml
spec:
  domain: example.com      # Incoming Host header must match this domain
```
* Each domain is isolated.
* Rate limiting and routing are applied per domain.
* Supports hosting multiple domains on the same ProxyX instance.

---

## 🔐 TLS Configuration

```yaml
tls:
  certFile: /path/cert.pem # TLS certificate file
  keyFile: /path/key.pem   # TLS private key file
```
* Enables HTTPS for the domain.
* Centralized TLS management for static, reverse proxy, and WebSocket routes.
* Supports HTTP/2 and secure WebSocket connections.

---

## 🚦 Rate Limiter
```yaml
rateLimit:
  requests: 1000           # Max requests allowed in the window
  windowSeconds: 2         # Time window in seconds
```
* Rate limiting is applied per domain, across all routes.
* Prevents abuse and overload of backend servers.
* Ensures fair resource allocation among clients.

### How it works

* Counts incoming requests for the domain.
* Rejects requests exceeding the limit in the given window.
* Sliding window algorithm is used for accuracy and smooth enforcement.

### Why it is needed

* Protects backends from spikes or malicious traffic.
* Ensures consistent QoS for multiple routes under the same domain.

---

## 🧭 Routes

Routes define how requests are processed.

```yaml
routes:
  - name: api-route
    path: /api/**
    type: ReverseProxy
```

### Route Fields

* **name**: Unique route identifier.

* **path**: URL path pattern to match requests.

* **type**: One of `Static`, `ReverseProxy`, or `Websocket`.


### 📁 Static Route
```yaml
type: Static
static:
  root: /var/www/html      # Filesystem path to serve files from
```
* Serves directly from disk.
* Fastest response, minimal overhead.
* Supports recursive path matching using `/**`.

### 🔁 Reverse Proxy Route

```yaml
type: ReverseProxy
reverseProxy:
  servers:
    - url: http://localhost:8080  # Backend server
    - url: http://localhost:8081  # Backend server
```

* Forwards HTTP requests to backend servers.
* Round-robin selection among healthy backends.
* Distributes traffic evenly to multiple services.

---

### 🔌 WebSocket Route
```yaml
type: Websocket
websocket:
  url: ws://localhost:9000/ws     # Persistent WebSocket backend
```

* Proxies WebSocket connections with full duplex streaming.
* Supports real-time applications like chat or dashboards.
* Bridges TCP streams between client and backend.

---

## 🧠 Route Matching Rules
```yaml
path: /**        # Matches all paths
path: /api/**    # Matches only /api/*
```

* Routes are evaluated in order from top to bottom.
* WebSocket routes must explicitly use `type: Websocket`.
* /** wildcard matches all sub-paths.

✅ Summary

* YAML-first configuration with inline comments.
* Metadata enables safe lifecycle management via CLI.
* RateLimiter protects domains from overload.
* Multiple route types supported: Static, ReverseProxy, WebSocket.
* TLS termination handled per domain.