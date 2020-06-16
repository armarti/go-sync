#!/usr/bin/env bash
set -x

P1=$$
P2=$$
onexit() { kill $P1; kill $P2; }
trap onexit EXIT

ssh -vvv -N -L 127.0.0.1:8295:cloud.armarti.com:8295 docker@cloud.armarti.com &
P1=$!
ssh -vvv -N -L 127.0.0.1:8000:cloud.armarti.com:8000 docker@cloud.armarti.com &
P2=$!

wait $P1
wait $P2
