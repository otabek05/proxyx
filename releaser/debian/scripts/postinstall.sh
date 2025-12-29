#!/bin/sh
set -e

# Create config directory
mkdir -p /etc/proxyx

# Copy web assets (installed by nfpm into /etc/proxyx.tmp)
if [ -d /etc/proxyx.web ]; then
    cp -r /etc/proxyx.web/* /etc/proxyx/
    rm -rf /etc/proxyx.web
fi

# Enable and start service
systemctl daemon-reload
systemctl enable proxyx
systemctl restart proxyx
