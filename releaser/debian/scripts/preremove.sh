#!/bin/sh
set -e

systemctl stop proxyx || true
systemctl disable proxyx || true

rm -f /etc/systemd/system/proxyx.service
rm -f /usr/local/bin/proxyx
rm -rf /etc/proxyx

systemctl daemon-reload
