#!/usr/bin/env bash

if [ "$EUID" -eq 0 ]; then
  go build main.go
  mv main /usr/bin/Animatic
else
  echo "Este programa deve ser rodado como sudo"
fi
