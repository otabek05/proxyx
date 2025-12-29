#!/bin/sh
set -e

systemctl stop proxyx 2>/dev/null || true
