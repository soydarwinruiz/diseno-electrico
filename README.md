# ⚡ Diseño Eléctrico Residencial (RETIE / NTC 2050)

Programa sencillo para Windows para diseñar la instalación eléctrica de una casa, con aparatos dibujados de forma realista en lugar de símbolos.

![Casa de ejemplo en modo simulación](docs/captura-simulacion.png)

## Descargar y usar (Windows 10/11)

1. Descargue [`dist/DisenoElectrico-Windows.zip`](dist/DisenoElectrico-Windows.zip) y descomprímalo.
2. Doble clic en **DisenoElectrico.exe** (se abre en su propia ventana usando Microsoft Edge).
   - Si aparece "Windows protegió su PC": **Más información → Ejecutar de todas formas** (el programa no tiene firma digital).
   - Alternativa: abrir `DisenoElectrico.html` en el navegador.

No requiere internet ni instalación.

## Funciones

- **Plano base:** muros, espacios (alcoba, sala, cocina, baño, lavadero…), puertas, ventanas e imagen de fondo del plano.
- **Aparatos realistas:** bombillo, plafón LED, aplique, ventilador, interruptor sencillo/doble/triple, conmutable (3 vías), cruce (4 vías), interruptor + tomacorriente, tomacorriente doble, GFCI, 240 V, ducha eléctrica, timbre, empalme, varilla de puesta a tierra.
- **Cables de colores y calibres AWG**, conectados a los bornes de cada aparato (fase, neutro, tierra).
- **Tablero** con breakers configurables (polos, amperaje, calibre, fase, diferencial).
- **Conexión automática:** luces desde 1, 2 o 3+ sitios; tomacorrientes encadenados.
- **▶ Simulación:** accione interruptores y breakers; detecta cortocircuitos.
- **Revisión RETIE / NTC 2050**, cuadro de cargas, balance de fases y lista de materiales (exportable a Excel).
- Guardar/abrir proyectos, exportar imagen, imprimir / PDF.

### Facilidad de uso
- Interruptores y tomacorrientes **se pegan al muro** y miran hacia el espacio (Alt: ubicación libre).
- **Clic derecho** para menú contextual: rotar, duplicar, asignar circuito, cablear…
- **Ver un solo circuito** (selector *Ver:* o 👁 en la pestaña Tablero).
- **Cablear todo lo que falta** con un clic, espacio por espacio.
- Paleta con **búsqueda** y **arrastrar y soltar**; lista de **primeros pasos**.
- Ctrl+C / Ctrl+V / Ctrl+D, flechas para mover fino, Ctrl+Z / Ctrl+Y.

![Ver un solo circuito](docs/ver-circuito.png)

## Estructura

| Carpeta | Contenido |
|---|---|
| `app/DisenoElectrico.html` | La aplicación completa (HTML + JavaScript, un solo archivo) |
| `launcher/` | Lanzador en Go que incrusta el HTML y lo abre en una ventana de Edge |
| `dist/` | ZIP listo para Windows |
| `build.sh` | Compila `dist/DisenoElectrico.exe` (requiere Go) |

## Aviso

Herramienta de apoyo para diseño y aprendizaje. El diseño definitivo debe ser revisado y firmado por un profesional competente con matrícula, conforme al RETIE y la NTC 2050 vigentes.
