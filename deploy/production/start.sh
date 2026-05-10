#!/bin/sh

# start backend
/app/homeledger &

# start next standalone
cd /app/frontend
HOSTNAME=0.0.0.0 PORT=3000 node server.js &

# start caddy
caddy run --config /app/Caddyfile
