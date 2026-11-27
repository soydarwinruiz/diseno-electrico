package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func fakeHTML(v string) []byte {
	return []byte("<!DOCTYPE html><html><script>const APP_VERSION='" + v + "';</script></html>")
}

func zipOf(html []byte) []byte {
	var b bytes.Buffer
	z := zip.NewWriter(&b)
	w, _ := z.Create("DisenoElectrico/DisenoElectrico.html")
	w.Write(html)
	z.Close()
	return b.Bytes()
}

// servidor que imita la API de Releases de GitHub
func fakeServer(t *testing.T, tag string, html []byte, sum string, assetHost string) *httptest.Server {
	var srv *httptest.Server
	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := srv.URL
		if assetHost != "" {
			host = assetHost
		}
		switch r.URL.Path {
		case "/latest":
			assets := fmt.Sprintf(`{"name":"DisenoElectrico-v%s-Windows.zip","browser_download_url":"%s/zip"}`, tag, host)
			if sum != "" {
				assets += fmt.Sprintf(`,{"name":"DisenoElectrico-v%s.sha256","browser_download_url":"%s/sum"}`, tag, host)
			}
			fmt.Fprintf(w, `{"tag_name":"v%s","assets":[%s]}`, tag, assets)
		case "/zip":
			w.Write(zipOf(html))
		case "/sum":
			fmt.Fprintf(w, "%s  DisenoElectrico.html\n", sum)
		default:
			w.WriteHeader(404)
		}
	}))
	t.Cleanup(srv.Close)
	t.Setenv("DE_UPDATE_URL", srv.URL+"/latest")
	return srv
}

func shaHex(b []byte) string { h := sha256.Sum256(b); return hex.EncodeToString(h[:]) }

func TestNewer(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{{"2.1.0", "2.0.0", true}, {"2.0.0", "2.1.0", false}, {"v10.0.0", "9.9.9", true}, {"2.0", "2.0.0", false}, {"2.0.1", "2.0", true}}
	for _, c := range cases {
		if newer(c.a, c.b) != c.want {
			t.Errorf("newer(%s,%s) != %v", c.a, c.b, c.want)
		}
	}
}

func TestPrepareWritesEmbedded(t *testing.T) {
	dir := t.TempDir()
	file, active, _ := prepare(dir)
	b, err := os.ReadFile(file)
	if err != nil || versionOf(b) != active || active == "0" {
		t.Fatalf("prepare: %v active=%s", err, active)
	}
}

func TestUpdateOK(t *testing.T) {
	html := fakeHTML("99.0.0")
	fakeServer(t, "99.0.0", html, shaHex(html), "")
	dir := t.TempDir()
	v, err := update(dir, "2.0.0")
	if err != nil || v != "99.0.0" {
		t.Fatalf("update = %q, %v", v, err)
	}
	b, _ := os.ReadFile(filepath.Join(dir, "DisenoElectrico.html"))
	if !bytes.Equal(b, html) {
		t.Fatal("el HTML no quedó instalado")
	}
	// la próxima apertura usa la versión descargada
	if _, active, _ := prepare(dir); active != "99.0.0" {
		t.Fatalf("prepare usa %s", active)
	}
}

func TestUpdateRejectsBadSum(t *testing.T) {
	html := fakeHTML("99.0.0")
	fakeServer(t, "99.0.0", html, strings.Repeat("0", 64), "")
	dir := t.TempDir()
	if _, err := update(dir, "2.0.0"); err == nil || !strings.Contains(err.Error(), "SHA-256") {
		t.Fatalf("debía rechazar la huella: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "DisenoElectrico.html")); err == nil {
		t.Fatal("no debía instalar nada")
	}
}

func TestUpdateRequiresSumForNewVersions(t *testing.T) {
	html := fakeHTML("99.0.0")
	fakeServer(t, "99.0.0", html, "", "")
	if _, err := update(t.TempDir(), "2.0.0"); err == nil {
		t.Fatal("debía exigir la huella")
	}
}

func TestUpdateRejectsTruncated(t *testing.T) {
	html := []byte("<html><script>const APP_VERSION='99.0.0';")
	fakeServer(t, "99.0.0", html, shaHex(html), "")
	if _, err := update(t.TempDir(), "2.0.0"); err == nil {
		t.Fatal("debía rechazar un archivo incompleto")
	}
}

func TestUpdateRejectsForeignHost(t *testing.T) {
	html := fakeHTML("99.0.0")
	fakeServer(t, "99.0.0", html, shaHex(html), "https://ejemplo-malicioso.invalid")
	if _, err := update(t.TempDir(), "2.0.0"); err == nil || !strings.Contains(err.Error(), "no permitida") {
		t.Fatalf("debía rechazar otro servidor: %v", err)
	}
}

func TestNoUpdateWhenCurrent(t *testing.T) {
	html := fakeHTML("2.0.0")
	fakeServer(t, "2.0.0", html, shaHex(html), "")
	if v, err := update(t.TempDir(), "2.0.0"); v != "" || err != nil {
		t.Fatalf("no debía actualizar: %q %v", v, err)
	}
}

func TestOptOut(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "no-actualizar.txt"), nil, 0o644)
	t.Setenv("DE_UPDATE_URL", "http://127.0.0.1:1/latest")
	if v, err := update(dir, "2.0.0"); v != "" || err != nil {
		t.Fatalf("con no-actualizar.txt no debe consultar: %q %v", v, err)
	}
}
