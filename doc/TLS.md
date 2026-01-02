## TLS & HTTPS with Certbot

ProxyX integrates with **Certbot** to issue and manage **Let's Encrypt TLS certificates** for your domains.  
It provides HTTPS support for **Static, ReverseProxy, and WebSocket routes**.

---

### ✅ Requirements

Before using TLS, you must install Certbot:

```bash
# RHEL / CentOS / Amazon Linux
sudo dnf install certbot

# Ubuntu / Debian
sudo apt install certbot

```

Ensure your system can reach the public internet and that the domain points to this server.

### ✅Certificate Issuance

To issue a certificate for a domain interactively:

```bash
sudo proxyx cert
```

ProxyX will prompt for:
  * Domain Select [number]
  * Email Address for Let's encrypt registration

Afterwards, ProxyX will:
  * Request the certificate
  * Automatically update your configuration to enable HTTPS


---

## Notes & Best Practices
 * Ports **80**(HTTP) and **443** (HTTPS) must be open for issuance and renewal.
 * Use a fixed domain for TLS to avoid Let's Encrypt validation errors.
 * ProxyX supports multiple domains, each with its own certificate.
 * Certificates issued by Let's Encrypt are trusted by all major browsers.