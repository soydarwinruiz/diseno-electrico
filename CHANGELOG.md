# Historial de versiones

Todas las versiones se pueden descargar en la sección **Releases** del repositorio.

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
