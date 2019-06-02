#!/bin/bash
PY_VERSION=2
set -x
curl https://raw.githubusercontent.com/LucienShui/Webhook/master/python${PY_VERSION}/server.py > ${PWD}/server.py 
systemctl restart python${PY_VERSION}_webhook

