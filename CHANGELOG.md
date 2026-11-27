# Historial de versiones

Todas las versiones se pueden descargar en la sección **Releases** del repositorio.

## v2.1.0 — Solidez: seguridad, respaldos y pruebas automáticas
- **Archivos seguros**: todo proyecto que se abre (archivo, Mis proyectos o versión guardada) se revisa antes de usarlo. Un archivo preparado con código malicioso ya no puede ejecutar nada, y un archivo dañado o incompleto no se abre y **deja intacto el proyecto actual**, con un mensaje claro de qué falló.
- **Imagen de fondo liviana**: se reduce sola a máximo 2000 px y se guarda aparte del proyecto, así una foto grande ya no llena el almacenamiento ni impide guardar. Los proyectos antiguos con imagen se convierten solos.
- **🛟 Respaldo completo**: en 📂 Mis proyectos, «Respaldar todo» descarga un solo archivo con todos los proyectos, sus versiones, imágenes y precios; «Restaurar respaldo» los recupera (también en otro computador). Opción de **copia diaria automática a una carpeta** y recordatorio si pasa más de una semana sin respaldo.
- **Actualizaciones verificadas**: el `.exe` comprueba la huella SHA-256 de la versión nueva, que esté completa y que venga de GitHub antes de usarla.
- **GFCI que se dispara**: en la simulación, el botón **TEST** del tomacorriente GFCI lo dispara (luz roja, sin energía) y **RESET** lo restablece; el diagnóstico lo detecta. Con **taco diferencial**, una falla a tierra la corta el diferencial, con su explicación.
- El diseño automático de toda la casa agrega y conecta la **varilla de puesta a tierra** si falta.
- **Recorrido de bienvenida** de 8 pasos la primera vez (y en Ayuda → «Recorrido rápido»). La Ayuda explica la vista 3D, el GFCI y los respaldos.
- La barra de herramientas ya **no se parte en dos líneas** en portátiles (1280–1366 px): los botones muestran solo el ícono y el nombre aparece al pasar el ratón. Se unificó el término **taco** (antes se mezclaba con «breaker»).
- **20 pruebas automáticas** en el repositorio (`tests/`) y pruebas del actualizador; GitHub las ejecuta en Linux y prueba el `.exe` en una máquina **Windows real** antes de publicar cada versión. Si algo falla, la versión no se publica.

## v2.0.0 — Vista 3D y recorrido caminando
- **🧊 Vista 3D** (botón 3D o tecla 3): la casa en 3D con muros, vanos de puertas y ventanas, pisos por espacio, muebles y cada aparato eléctrico a su altura real (tomas a 0,30 m, interruptores a 1,20 m, luces en el techo, tablero a 1,45 m).
- **Luces reales**: las luces encendidas en la simulación iluminan la casa; modo **🌙 Noche** para verlas mejor. Los ventiladores y extractores giran.
- **🚶 Caminar** en primera persona por la casa (ratón para mirar, W A S D o flechas), sin atravesar muros, y **accionar interruptores, tacos y tomacorrientes con un clic** dentro del 3D.
- **Girar**, desplazar y acercar la casa; opción de **muros bajos** para ver la distribución desde arriba.
- Usa la librería three.js (licencia MIT), incluida dentro del programa: funciona sin internet.

## v1.9.0 — Cables por los muros, presupuesto y actualizaciones automáticas
- **Cables por los muros**: la conexión automática y los cables nuevos siguen los muros, como en una obra real (idea adaptada de *house-planner*, MIT). Los metros de cable, la tubería y la caída de tensión salen más exactos. Botón para rehacer todas las rutas.
- **💲 Presupuesto** (pestaña Costos): precios de ejemplo editables que se recuerdan para todos los proyectos, mano de obra, misceláneos, IVA, casillas de «comprado» con el valor pagado, exportación a Excel y cotización para imprimir.
- **Actualizaciones automáticas**: el programa avisa cuando hay una versión nueva y muestra sus novedades. Con el `.exe`, la versión nueva se descarga sola en segundo plano y se usa la próxima vez que se abra.
- **Plano de fondo**: calibración exacta con dos puntos sobre una medida conocida y opción para mover la imagen.
- **Versiones dentro de un proyecto**: guardar y restaurar momentos del diseño.
- **Resumen de áreas** y plantillas nuevas de **1 alcoba** y **4 alcobas**.

## v1.8.0 — Diseño automático, tubería, multímetro y retos
- **🔌 Tomas automáticas**: por espacio o para toda la casa, según la regla de 1,80 m, sin tapar puertas y con GFCI en baños, cocina y zonas húmedas; quedan conectadas al tablero.
- **💡 Luz con interruptor en un clic**: al poner una luz en un espacio sin interruptor, el programa ubica el interruptor junto a la puerta y lo conecta (con dos puertas usa conmutables). Ctrl+Z quita solo el interruptor.
- **✨ Diseño automático de toda la casa** desde Propiedades del proyecto.
- **Tubería (conduit)** calculada a partir de los cables: diámetro (½", ¾", 1"…) según la ocupación de la NTC 2050, número de cables por tramo, metros, tubos de 3 m y curvas. Se ve en el plano con la casilla «tubería».
- **🔧 Multímetro** (tecla U): mida el voltaje entre dos bornes (120 V, 240 V, 0 V, 12 V) con explicación.
- **🔍 ¿Por qué no prende?**: clic derecho en una luz o toma durante la simulación para ver el diagnóstico paso a paso (falta el neutro, el retorno no llega, el taco está disparado, el interruptor está apagado…).
- **🎯 Retos de práctica**: 6 ejercicios que el programa revisa probando todas las posiciones de los interruptores.
- **Aparatos nuevos**: dimmer, sensor de movimiento, fotocelda, extractor de baño, transformador 120/12 V con timbre/citófono de 12 V, toma de TV y toma de datos. Nuevas guías de sensor y timbre de 12 V.
- Menú **📁 Archivo** que agrupa nuevo, abrir, guardar, exportar e imprimir.

## v1.7.0 — Proyectos, planos profesionales y herramientas
- **📂 Mis proyectos**: cada proyecto se guarda por separado en el equipo. «Nuevo», «Casa de ejemplo» y «Abrir archivo» ya no reemplazan el trabajo actual; se puede abrir, duplicar y eliminar proyectos.
- **Impresión / PDF profesional en 3 hojas** con **rótulo** (proyecto, dirección, propietario, diseñó, matrícula, fecha, hoja y firma) y **convenciones** de aparatos y colores de cable.
- **📐 Diagrama unifilar** generado automáticamente: red, medidor, totalizador, barraje, cada taco con su calibre, carga y fase, y la puesta a tierra.
- **Caída de tensión (ΔV %)** por circuito, calculada con la longitud real de los cables dibujados; aviso cuando supera el 3 %.
- **📏 Medir** distancias en el plano (tecla D), con imán a esquinas y bordes.
- **Notas de texto** en el plano para aclaraciones al instalador.
- **⌨ Atajos de teclado** (tecla ? o F1).
- Datos del rótulo en Propiedades del proyecto; barra de colores más compacta.

## v1.6.0 — Guías, glosario y capacidad del tablero
- **📖 Guías rápidas** de las instalaciones más comunes, con diagrama real, pasos, materiales y un botón para insertarlas en el plano y probarlas: interruptor sencillo, doble, luz desde 2 sitios (conmutables), desde 3 o más sitios (cruce), tomas con polo a tierra, GFCI, interruptor + toma, timbre, ducha/estufa 240 V, puesta a tierra y derivaciones.
- **Glosario** con buscador: fase, neutro, retorno, viajeros, COM, S1/S2, V1/V2, tándem, totalizador, AWG y más.
- Cada aparato explica en Propiedades qué significa cada borne, y tiene un botón a su guía.
- El tablero del plano ahora **cambia según la capacidad elegida**: muestra los espacios libres y avisa si no caben los tacos.
- Número de versión visible en el programa.

## v1.5.0 — Derivaciones de cable
- Con la herramienta Cable, un clic sobre un cable crea un **empalme** y arranca un ramal del mismo color y calibre.
- Clic derecho → «Derivar un cable desde aquí»; al borrar un empalme de dos tramos el cable se vuelve a unir.

## v1.4.0 — Mover espacios con anclaje
- Los espacios se arrastran desde cualquier punto y **se anclan solos** a los vecinos (guías rosadas); llevan su contenido.
- Manijas en lados y esquinas con imán, aviso de superposición, doble clic para renombrar.

## v1.3.0 — Plano fácil y modelos de tablero
- Muros automáticos a partir de los espacios, arrastrar muros, cotas, puertas y ventanas que se insertan en el muro, muebles y plantillas.
- Modelos de tablero (metálico de empotrar, blanco de sobreponer, riel DIN) y tipos de taco (sencillo, tándem, bipolar, diferencial).
- Salidas del tablero por una bornera inferior rotulada.

## v1.2.0 — Tablero realista e interactivo
- Vista del tablero con tapa, frente con tacos e interior; tacos con animación, TRIP y rearme.
- Medidor de energía en vivo, aparatos conectables, disparo por sobrecarga y sonidos.

## v1.1.0 — Facilidad de uso
- Barra compacta, aparatos que se pegan al muro, menú de clic derecho, copiar/pegar, filtro por circuito, cablear todo, primeros pasos.

## v1.0.0 — Versión inicial
- Plano base, aparatos realistas, cables de colores, tablero y circuitos, conexión automática, simulación, revisión RETIE / NTC 2050, cuadro de cargas, materiales y lanzador para Windows.
