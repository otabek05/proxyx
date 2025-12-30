#!/bin/sh
set -e

launchctl bootout system/com.proxyx.service 2>/dev/null || true
