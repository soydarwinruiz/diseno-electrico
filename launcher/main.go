// Lanzador para Windows: abre el Diseño Eléctrico en una ventana propia de Microsoft Edge
// (modo aplicación, sin barra de direcciones). El HTML va incrustado dentro del .exe.
//
// Actualización automática: al abrir, el programa arranca de inmediato con la mejor versión
// disponible y, en segundo plano, consulta las Releases de GitHub. Si hay una versión más nueva,
// descarga el ZIP, extrae DisenoElectrico.html y lo deja listo para la próxima vez que se abra.
// Para desactivarlo, cree el archivo "no-actualizar.txt" en %LOCALAPPDATA%\DisenoElectrico.
package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed DisenoElectrico.html
var page []byte

const defaultReleaseURL = "https://api.github.com/repos/soydarwinruiz/diseno-electrico/releases/latest"

var verRe = regexp.MustCompile(`const APP_VERSION='([0-9.]+)'`)

func versionOf(html []byte) string {
	m := verRe.FindSubmatch(html)
	if m == nil {
		return "0"
	}
	return string(m[1])
}

// newer devuelve true si a es una versión mayor que b (formato 1.2.3).
func newer(a, b string) bool {
	pa, pb := strings.Split(strings.TrimPrefix(a, "v"), "."), strings.Split(strings.TrimPrefix(b, "v"), ".")
	for i := 0; i < 3; i++ {
		var x, y int
		if i < len(pa) {
			x, _ = strconv.Atoi(pa[i])
		}
		if i < len(pb) {
			y, _ = strconv.Atoi(pb[i])
		}
		if x != y {
			return x > y
		}
	}
	return false
}

func readText(p string) string {
	b, err := os.ReadFile(p)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}

// prepare escribe el HTML activo y devuelve su ruta, su versión y la versión con la que se abrió la vez anterior.
func prepare(dir string) (string, string, string) {
	_ = os.MkdirAll(dir, 0o755)
	file := filepath.Join(dir, "DisenoElectrico.html")
	embedded := versionOf(page)
	cached := readText(filepath.Join(dir, "version.txt"))
	active := embedded
	if cached != "" && newer(cached, embedded) {
		if b, err := os.ReadFile(file); err == nil && versionOf(b) == cached {
			active = cached // usar la actualización descargada antes
		}
	}
	if active == embedded {
		_ = os.WriteFile(file, page, 0o644)
		_ = os.WriteFile(filepath.Join(dir, "version.txt"), []byte(embedded), 0o644)
	}
	last := readText(filepath.Join(dir, "ultima-apertura.txt"))
	_ = os.WriteFile(filepath.Join(dir, "ultima-apertura.txt"), []byte(active), 0o644)
	return file, active, last
}

func launch(file, base, updatedTo string) {
	u := &url.URL{Scheme: "file", Path: "/" + filepath.ToSlash(file)}
	q := url.Values{"exe": {"1"}}
	if updatedTo != "" {
		q.Set("actualizado", updatedTo)
	}
	u.RawQuery = q.Encode()
	browsers := []string{
		filepath.Join(os.Getenv("ProgramFiles(x86)"), `Microsoft\Edge\Application\msedge.exe`),
		filepath.Join(os.Getenv("ProgramFiles"), `Microsoft\Edge\Application\msedge.exe`),
		filepath.Join(base, `Microsoft\Edge\Application\msedge.exe`),
		filepath.Join(os.Getenv("ProgramFiles"), `Google\Chrome\Application\chrome.exe`),
		filepath.Join(os.Getenv("ProgramFiles(x86)"), `Google\Chrome\Application\chrome.exe`),
		filepath.Join(base, `Google\Chrome\Application\chrome.exe`),
	}
	for _, b := range browsers {
		if st, err := os.Stat(b); err == nil && !st.IsDir() {
			if exec.Command(b, "--app="+u.String(), "--window-size=1440,900").Start() == nil {
				return
			}
		}
	}
	_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", file).Start()
}

type release struct {
	Tag    string `json:"tag_name"`
	Assets []struct {
		Name string `json:"name"`
		URL  string `json:"browser_download_url"`
	} `json:"assets"`
}

// allowedURL limita las descargas a GitHub (o al servidor de prueba indicado en DE_UPDATE_URL).
func allowedURL(u string) bool {
	p, err := url.Parse(u)
	if err != nil {
		return false
	}
	if t := os.Getenv("DE_UPDATE_URL"); t != "" {
		if q, e := url.Parse(t); e == nil && q.Host == p.Host {
			return true
		}
	}
	return p.Scheme == "https" && (p.Host == "github.com" || strings.HasSuffix(p.Host, ".githubusercontent.com"))
}

func download(client *http.Client, u string, max int64) ([]byte, error) {
	if !allowedURL(u) {
		return nil, errors.New("dirección de descarga no permitida: " + u)
	}
	r, err := client.Get(u)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return nil, errors.New("descarga " + r.Status)
	}
	return io.ReadAll(io.LimitReader(r.Body, max))
}

// verify comprueba que el programa descargado esté completo y coincida con la huella SHA-256 publicada.
func verify(html []byte, tag, sum string) error {
	if versionOf(html) != tag {
		return errors.New("la versión del archivo no coincide")
	}
	if !bytes.Contains(html, []byte("</html>")) {
		return errors.New("el archivo descargado está incompleto")
	}
	if sum == "" {
		if !newer("2.1.0", tag) {
			return errors.New("la versión no publica su huella SHA-256")
		}
		return nil
	}
	h := sha256.Sum256(html)
	if !strings.EqualFold(hex.EncodeToString(h[:]), sum) {
		return errors.New("la huella SHA-256 no coincide: descarga dañada o alterada")
	}
	return nil
}

// update descarga la versión nueva si existe y la deja en dir. Devuelve la versión instalada o "".
func update(dir, active string) (string, error) {
	if _, err := os.Stat(filepath.Join(dir, "no-actualizar.txt")); err == nil {
		return "", nil
	}
	api := os.Getenv("DE_UPDATE_URL")
	if api == "" {
		api = defaultReleaseURL
	}
	client := &http.Client{Timeout: 15 * time.Second}
	req, _ := http.NewRequest("GET", api, nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "DisenoElectrico-actualizador")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", errors.New("respuesta " + res.Status)
	}
	var rel release
	if err := json.NewDecoder(io.LimitReader(res.Body, 1<<20)).Decode(&rel); err != nil {
		return "", err
	}
	tag := strings.TrimPrefix(rel.Tag, "v")
	if !newer(tag, active) {
		return "", nil
	}
	zipURL, sumURL := "", ""
	for _, a := range rel.Assets {
		if strings.HasSuffix(a.Name, "-Windows.zip") {
			zipURL = a.URL
		}
		if strings.HasSuffix(a.Name, ".sha256") {
			sumURL = a.URL
		}
	}
	if zipURL == "" {
		return "", errors.New("la release no tiene ZIP")
	}
	client.Timeout = 120 * time.Second
	sum := ""
	if sumURL != "" {
		b, err := download(client, sumURL, 4096)
		if err != nil {
			return "", err
		}
		f := strings.Fields(string(b))
		if len(f) == 0 || len(f[0]) != 64 {
			return "", errors.New("archivo de huella inválido")
		}
		sum = f[0]
	}
	data, err := download(client, zipURL, 64<<20)
	if err != nil {
		return "", err
	}
	z, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return "", err
	}
	for _, f := range z.File {
		if strings.HasSuffix(f.Name, "DisenoElectrico.html") {
			rc, err := f.Open()
			if err != nil {
				return "", err
			}
			html, err := io.ReadAll(io.LimitReader(rc, 32<<20))
			rc.Close()
			if err != nil {
				return "", err
			}
			if err := verify(html, tag, sum); err != nil {
				return "", err
			}
			tmp := filepath.Join(dir, "DisenoElectrico.html.nuevo")
			if err := os.WriteFile(tmp, html, 0o644); err != nil {
				return "", err
			}
			if err := os.Rename(tmp, filepath.Join(dir, "DisenoElectrico.html")); err != nil {
				return "", err
			}
			_ = os.WriteFile(filepath.Join(dir, "version.txt"), []byte(tag), 0o644)
			return tag, nil
		}
	}
	return "", errors.New("el ZIP no contiene el programa")
}

func main() {
	base := os.Getenv("LOCALAPPDATA")
	if base == "" {
		base = os.TempDir()
	}
	dir := filepath.Join(base, "DisenoElectrico")
	file, active, last := prepare(dir)
	updated := ""
	if last != "" && newer(active, last) {
		updated = active
	}
	if os.Getenv("DE_NO_LAUNCH") == "" {
		launch(file, base, updated)
	}
	// Buscar la versión nueva en segundo plano; quedará lista para la próxima apertura.
	if v, err := update(dir, active); err != nil {
		_ = os.WriteFile(filepath.Join(dir, "actualizacion-error.txt"), []byte(time.Now().Format(time.RFC3339)+" "+err.Error()), 0o644)
	} else if v != "" {
		_ = os.WriteFile(filepath.Join(dir, "actualizacion.txt"), []byte(time.Now().Format(time.RFC3339)+" descargada v"+v), 0o644)
	}
}
