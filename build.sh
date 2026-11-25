#!/usr/bin/env bash
# Compila el lanzador de Windows (DisenoElectrico.exe) y arma el ZIP de distribución.
set -euo pipefail
cd "$(dirname "$0")"
cp app/DisenoElectrico.html launcher/
(cd launcher && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-H windowsgui -s -w" -o ../dist/DisenoElectrico.exe .)
rm launcher/DisenoElectrico.html
echo "Listo: dist/DisenoElectrico.exe"
