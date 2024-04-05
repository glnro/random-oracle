#!/bin/bash
set -eux

pkill -f appd &> /dev/null || true
