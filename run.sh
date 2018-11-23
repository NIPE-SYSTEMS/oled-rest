#!/bin/sh

echo "build ..."
# GOARM=6 (Raspberry Pi A, A+, B, B+, Zero)
# GOARM=7 (Raspberry Pi 2, 3)
GOOS=linux GOARCH=arm GOARM=6 go build

echo "deploy ..."
scp oled-rest root@camera:/root

echo "run ..."
ssh -t root@camera /root/oled-rest
