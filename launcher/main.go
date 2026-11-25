// Lanzador para Windows: abre el Diseño Eléctrico en una ventana propia de Microsoft Edge
// (modo aplicación, sin barra de direcciones). El HTML va incrustado dentro del .exe.
package main

import (
	_ "embed"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed DisenoElectrico.html
var page []byte

func main() {
	base := os.Getenv("LOCALAPPDATA")
	if base == "" {
		base = os.TempDir()
	}
	dir := filepath.Join(base, "DisenoElectrico")
	_ = os.MkdirAll(dir, 0o755)
	file := filepath.Join(dir, "DisenoElectrico.html")
	_ = os.WriteFile(file, page, 0o644)

	u := (&url.URL{Scheme: "file", Path: "/" + filepath.ToSlash(file)}).String()
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
			if exec.Command(b, "--app="+u, "--window-size=1440,900").Start() == nil {
				return
			}
		}
	}
	// Último recurso: navegador predeterminado
	_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", file).Start()
}
