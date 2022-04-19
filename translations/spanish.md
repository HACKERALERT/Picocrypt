<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p>

Picocrypt es una herramienta de cifrado muy pequeña (de ahí <i>Pico</i>), muy simple pero muy segura que puede usar para proteger sus archivos. Está diseñada para ser la herramienta <i>de referencia</i> para el cifrado, con un enfoque en la seguridad, la simplicidad y la confiabilidad. Picocrypt utiliza el cifrado seguro XChaCha20 y la función de derivación de clave Argon2id para proporcionar un alto nivel de seguridad, incluso de agencias de tres letras como la NSA. Está diseñada para la máxima seguridad, sin comprometerla en absolutamente nada, y está construida con los módulos x/crypto estándar de Go. <strong>Su privacidad y seguridad están bajo ataque. Recuperelas con confianza protegiendo tus archivos con Picocrypt.

<p align="center"><img align="center" src="/images/screenshot.png" width="318" alt="Picocrypt"></p>

# Financiación
Por favor done a Picocrypt en <a href="https://opencollective.com/picocrypt">Open Collective</a> (se aceptan criptomonedas) para recaudar dinero para una posible auditoría de Cure53. Debido a que este es un proyecto en el que dedico muchas horas y no gano dinero, no puedo pagar una auditoría yo mismo. <i>Picocrypt necesita el apoyo de su comunidad.</i>
  
# Descargas
**Importante**: Existe un software ya abandonado, desactualizado e inservible llamado PicoCrypt en Internet, que fue actualizado por última vez en 2005. PicoCrypt no está relacionado de ninguna manera con Picocrypt (este proyecto). Asegúrese que únicamente descarga Picocrypt desde este repositorio para asegurarse de obtener el Picocrypt auténtico y sin puertas traseras.
  
## Windows
Picocrypt para Windows es tan simple como parece. Para descargar el ejecutable más reciente, independiente y portable para Windows, haga clic <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.26/Picocrypt.exe">aquí</a>. Si Windows Defender o su antivirus marca Picocrypt como un virus, por favor haga su parte y repoórtelo como un falso positivo para el bien de todos.
  
## macOS
Picocrypt para macOS también es muy simple. Descargue Picocrypt <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.26/Picocrypt.app.zip">aquí</a>, extraiga el archivo zip, y ejecute el archivo Picocrypt que se encuentra dentro. Si no puede abrir Picocrypt porque no procede de un desarrollador verificado, haga clic derecho con el ratón en Picocrypt y pulse "Abrir". Si todavía aparece la alerta, haga clic derecho con el ratón en Picocrypt y pulse "Abrir" de nuevo y ya debería poder arrancar Picocrypt.
  
## Linux
Hay una imagen Snap disponible para Linux. Asumiendo que estás en un sistema basado en Debian, un simple `apt install snapd` y `snap install picocrypt` será suficiente. Para otras distribuciones como Fedora, puede consultar las instrucciones detalladas que están disponibles en https://snapcraft.io/picocrypt. Debido a la complejidad de las dependencias y los enlaces estáticos, no distribuyo binarios independientes .deb o .rpm porque no serían confiables y no valdrían la pena. Snapcraft administra todas las dependencias y ejecutables automáticamente y es la forma recomendada de ejecutar Picocrypt en cualquier distribución importante de Linux. Además, Snapcraft brinda mejor seguridad y contenedorización que Flatpak y AppImage, lo cual es importante para una herramienta de cifrado como Picocrypt. Si prefiere no tratar con Canonical, recuerde que compilar desde el código fuente siempre es una opción.

## Paquetes paranoicos
Paranoid Pack es un archivo comprimido que contiene ejecutables para cada versión de Picocrypt que se haya lanzado para Windows, macOS y Linux. Siempre que lo tenga almacenado en un lugar al que pueda acceder, podrá abrirlo y usar cualquier versión de Picocrypt en caso de que este repositorio desaparezca misteriosamente o se queme todo Internet. Piense en ello como un depósito de semillas para Picocrypt. Siempre que una persona tenga el Paranoid Pack a su alcance, puede compartirlo con el resto del mundo y mantener Picocrypt en funcionamiento en casos de eventos catastróficos como el cierre repentino de GitHub o la captura de la NSA (por si acaso, ¿entiende?) . La mejor manera de garantizar que se pueda acceder a Picocrypt dentro de muchas décadas es mantener un Paranoid Pack en un lugar seguro. Entonces, si le preocupa no poder acceder a Picocrypt en el futuro, aquí está su solución. Simplemente diríjase a la pestaña _Releases_ (Publicaciones) y obtenga una copia.
  
# ¿Por qué Picocrypt?
¿Por qué debería utilizar Picocrypt en vez de BitLocker, NordLocker, VeraCrypt, AxCrypt, or 7-Zip? Aquí tienes algunas razones por las que debería escoger Picocrypt:
<ul>
	<li>A diferencia de NordLocker, BitLocker, AxCrypt, y la mayoría de proveedores de almacenamiento en la nube, Picocrypt y sus depedencias son totalmente de código abierto y auditables. Puede verificar usted mismo que no existen puertas traseras o fallos.</li>
	<li>Picocrypt es <i>minúsculo</i>. Mientras que NordLocker supera los 50MB de tamaño y VeraCrypt está sobre 20MB, Picocrypt se sitúa en solo 2MB, aproximadamente el tamaño de una fotografía de resolución media. Y no solo eso, Picocrypt es portable (no necesita ser instalado) y no necesita de privilegios de administrador o root.</li>
	<li>Picocrypt es más sencillo y más productivo de utilizar que VeraCrypt. Para cifrar archivos con VeraCrypt, tienes que pasar al menos cinco minutos estableciendo un volúmen. Con la interfaz simple de Picocrypt, todo lo que necesitas hacer es arrastrar y soltar tus archivos, introducir una contraseña y pulsar Iniciar. Todos los procedimientos complejos son manejados de manera interna por Picocrypt. ¿Quién dijo que el cifrado seguro no puede ser simple?</li>
	<li>Picocrypt está diseñado para la seguridad. 7-Zip es una utilidad de archivos y no una herramienta de cifrado, así que su foco principal no está en la seguridad. Picocrypt, sin embargo, está creado con la seguridad como el número uno de sus prioridades. Cada parte de Picocrypt existe por una razón y cualquier cosa que pudiera tener un impacto en la seguridad de Picocrypt es eliminada. Picocrypt está construido con cifrado en el que puede confiar.</li>
	<li>Picocrypt autentica los datos además de protegerlos, evitando que los delincuentes informáticos modifiquen datos confidenciales de forma malintencionada. Esto es útil cuando envía archivos cifrados a través de un canal inseguro y quiere asegurarse de que lleguen intactos.</li>
	<li>Picocrypt protege activamente los datos del cifrados del encabezado contra la corrupción al agregar bytes de paridad Reed-Solomon adicionales, por lo que si parte del encabezado de un volumen (que contiene componentes criptográficos importantes) se corrompe (por ejemplo, se pudre el bit del disco duro), Picocrypt aún puede recuperar el encabezado y descifrar sus datos con una alta tasa de éxito. Picocrypt también puede codificar todo el volumen con Reed-Solomon para evitar que sus archivos importantes se dañen.</li>
</ul>
	
# Comparativa
Así es como Picocrypt se compara con otras herramientas de cifrado populares.
	
|                | Picocrypt        | VeraCrypt       | 7-Zip (GUI)    | NordLocker     | BitLocker        | AxCrypt        |
| -------------- | ---------------- | --------------- | -------------- | -------------- | ---------------- | -------------- |
| Libre          |✅ Sí             |✅ Sí           |✅ Sí           |🟧 Parcialmente |🟧 Parcialmente   |🟧 Parcialmente|
| Código abierto |✅ GPLv3          |✅ Multi        |✅ LGPL         |❌ No           |❌ No             |❌ No          |
| Multiplataforma|✅ Sí             |✅ Sí           |❌ No           |❌ No           |❌ No             |❌ No          |
| Tamaño         |✅ 2MB            |❌ 20MB         |✅ 2MB          |❌ 60MB         |✅ Incluido       |🟧 8MB         |
| Portable       |✅ Sí             |✅ Sí           |❌ No           |❌ No           |✅ Sí             |✅ Sí          |
| Permisos       |✅ Ninguno        |❌ Admin        |❌ Admin        |❌ Admin        |❌ Admin          |❌ Admin       |
| Facilidad de uso|✅ Fácil          |❌ Hard         |✅ Fácil        |🟧 Medio        |🟧 Medio          |✅ Fácil       |
| Derivac. de clave|✅ Argon2         |🆗 PBKDF2       |❌ SHA256       |✅ Argon2       |❓ Desconocido     |🆗 PBKDF2      |
| Integridad datos |✅ Siempre        |❌ No           |❌ No           |✅ Siempre      |❓ Desconocido     |✅ Siempre     |
| Reed-Solomon   |✅ Sí             |❌ No           |❌ No           |❌ No           |❌ No              |❌ No         |
| Compresión     |✅ Sí             |❌ No           |✅ Sí           |❌ No           |✅ Sí             |✅ Sí          |
| Telemetría     |✅ Ninguna        |✅ Ninguna      |✅ Ninguna      |❌ Analíticas   |❓ Desconocido     |❌ Cuentas     |
| Auditado       |🟧 Plaificado     |✅ Yes          |❌ No           |❓ Desconocido   |❓ Desconocido     |❌ No          |
