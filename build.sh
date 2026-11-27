#!/usr/bin/env bash
# Compila el lanzador de Windows (DisenoElectrico.exe), arma el ZIP de distribución
# y escribe la huella SHA-256 del programa (la comprueba la actualización automática).
set -euo pipefail
cd "$(dirname "$0")"
V=$(grep -oP "const APP_VERSION='\K[0-9.]+" app/DisenoElectrico.html)
mkdir -p dist
cp app/DisenoElectrico.html launcher/
(cd launcher && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-H windowsgui -s -w" -o ../dist/DisenoElectrico.exe .)
rm launcher/DisenoElectrico.html
T=$(mktemp -d); trap 'rm -rf "$T"' EXIT
mkdir "$T/DisenoElectrico"
cp dist/DisenoElectrico.exe app/DisenoElectrico.html "$T/DisenoElectrico/"
sed "s/@VERSION@/$V/" launcher/LEEME.txt | sed 's/$/\r/' > "$T/DisenoElectrico/LEEME.txt"
rm -f dist/DisenoElectrico-Windows.zip
(cd "$T" && zip -qr -X "$OLDPWD/dist/DisenoElectrico-Windows.zip" DisenoElectrico)
(cd app && sha256sum DisenoElectrico.html) > dist/DisenoElectrico.html.sha256
echo "Listo v$V: dist/DisenoElectrico.exe, dist/DisenoElectrico-Windows.zip, dist/DisenoElectrico.html.sha256"
