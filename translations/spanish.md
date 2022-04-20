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
| Gratuito       |✅ Sí             |✅ Sí           |✅ Sí           |🟧 Parcialmente |🟧 Parcialmente   |🟧 Parcialmente|
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
	
# Características
Picocrypt es una herramienta muy simple, y la mayoría de los usuarios comprenderán intuitivamente cómo usarla en unos segundos. En un nivel básico, simplemente arrastrar y soltar sus archivos, ingresar una contraseña y presionar Iniciar es todo lo que necesita para cifrar sus archivos. Bastante simple, ¿verdad?

Si bien es simple, Picocrypt también se esfuerza por ser poderoso en manos de usuarios expertos y avanzados. Por lo tanto, hay algunas opciones adicionales que puede utilizar para satisfacer sus necesidades.
	<li><strong>Generador de contraseña</strong>: Picocrypt proporciona un generador de contraseñas seguras que puede usar para crear contraseñas criptográficamente seguras. Puede personalizar la longitud de la contraseña, así como los tipos de caracteres que se incluirán.</li>
	<li><strong>Comentarios</strong>: Úselo para almacenar notas, información y texto junto con el archivo (no se cifrará). Por ejemplo, puede poner una descripción del archivo que está cifrando antes de enviárselo a alguien. Cuando la persona a la que se lo poga el archivo en Picocrypt, su descripción se mostrará a esa persona.</li>
	<li><strong>Archivos de claves</strong>: Picocrypt admite el uso de archivos de claves como una forma adicional de autenticación (o la única forma de autenticación). No solo puede usar varios archivos de claves, sino que también puede requerir que esté presente el orden correcto de los archivos de claves para que se produzca un descifrado exitoso. Un caso de uso particularmente bueno de varios archivos de claves es la creación de un volumen compartido, donde cada persona tiene un archivo de claves, y todos ellos (y sus archivos de claves) deben estar presentes para descifrar el volumen compartido.</li>
	<li><strong>Modo paranóico mode</strong>: El uso de este modo cifrará sus datos con XChaCha20 y Serpent en forma de cascada, y utilizará HMAC-SHA3 para autenticar los datos en lugar de BLAKE2b. Esto se recomienda para proteger archivos de alto secreto y proporciona el nivel más alto de seguridad práctica posible. Para que un delincuente informático pueda descifrar sus datos cifrados, tanto el cifrado XChaCha20 como el cifrado Serpent deben estar rotos, suponiendo que haya elegido una buena contraseña. Es seguro decir que en este modo, sus archivos son imposibles de descifrar.</li>
	<li><strong>Reed-Solomon</strong>: Esta función es muy útil si planea archivar datos importantes en un proveedor de la nube o en un medio externo durante mucho tiempo. Si está marcado, Picocrypt utilizará el código de corrección de errores Reed-Solomon para agregar 8 bytes adicionales por cada 128 bytes para evitar la corrupción de archivos. Esto significa que hasta ~3% de su archivo puede corromperse y Picocrypt aún podrá corregir los errores y descifrar sus archivos sin corrupción. Por supuesto, si su archivo se corrompe mucho (por ejemplo, se le cayó el disco duro), Picocrypt no podrá recuperar sus archivos por completo, pero hará todo lo posible para recuperar lo que pueda. Tenga en cuenta que esta opción ralentizará considerablemente el cifrado y el descifrado.</li>
	<li><strong>Forzar descifrado</strong>: Picocrypt verifica automáticamente la integridad del archivo al descifrarlo. Si el archivo se ha modificado o está dañado, Picocrypt eliminará automáticamente la salida para la seguridad del usuario. Si desea anular estas medidas de seguridad, marque esta opción. Además, si esta opción está marcada y se utilizó la función Reed-Solomon en el volumen cifrado, Picocrypt intentará recuperar la mayor cantidad posible del archivo durante el descifrado.</li>
	<li><strong>Dividir archivos en fragmentos</strong>: ¿No tiene ganas de lidiar con archivos gigantescos? ¡No hay problema! Con Picocrypt, puede optar por dividir su archivo de salida en fragmentos de tamaño personalizado, de modo que los archivos grandes sean más manejables y más fáciles de cargar a los proveedores de la nube. Simplemente elija una unidad (KiB, MiB, GiB o TiB) e ingrese el tamaño de porción deseado para esa unidad. Para descifrar los fragmentos, simplemente arrastre uno de ellos a Picocrypt y los fragmentos se recombinarán automáticamente durante el descifrado.</li>
</ul>

# Seguridad
Para obtener más información sobre cómo Picocrypt maneja la criptografía, consulte <a href="Internals.md">Internals</a> para más detalles técnicos. Si está preocupado por mi seguridad o por la de este proyecto, permítame asegurarle que este repositorio no será secuestrado ni podrá incluir puertas traseras. Tengo habilitado 2FA (TOTP) en todas las cuentas vinculadas a Picocrypt (GitHub, Google, Reddit, Ubuntu One/Snapcraft, Discord, etc.), además de cifrado de disco completo en todos mis dispositivos portátiles. Para un mayor fortalecimiento, Picocrypt usa mis _forks_ aisladas de dependencias y añado los cambios de los proyectos originales solo cuando he echado un vistazo a los cambios y creo que no hay ningún problema de seguridad. Esto significa que si una dependencia es atacada o el autor elimina una dependencia, Picocrypt usará mi _fork_ y no se verá afectado en absoluto. Puede sentirse seguro al usar Picocrypt.

# Comunidad
Aquí hay algunos lugares donde puede mantenerse al día con Picocrypt y participar:
<ul>
	<li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
	<li><a href="https://discord.gg/8QM4A2caxH">Discord</a></li>
</ul>

Le recomiendo que se una al subreddit de Picocrypt porque todas las actualizaciones y encuestas se publicarán allí. Recuerde confiar solo en estas redes sociales y estar atento a los delincuentes informáticos que podrían intentar hacerse pasar por mí. Nunca le pediré su contraseña, y cualquiera que lo haga no soy yo. Nunca le diré que descargue un archivo de un enlace sospechoso, y cualquiera que lo haga no soy yo.

# _Stargazers_
¿Cómo va Picocrypt? Echa un vistazo a continuación para averiguarlo
![Stargazers Over Time](https://starchart.cc/HACKERALERT/Picocrypt.svg)

# Donaciones
Si encuentra útil Picocrypt, por favor considere en darme una propina en mi <a href="https://paypal.me/evanyiwensu">PayPal</a>. Ofrezco este software completamente gratuito y me encantaría tener algunos donante que me motivarán a continuar mi trabajo en Picocrypt.
	
# Gracias
Un agradecimiento desde el fondo de mi corazón a la gente de Open Collective que ha hecho una contribución significativa:
<ul>
	<li>YellowNight ($818)</li>
	<li>jp26 ($50)</li>
	<li>guest-116103ad ($50)</li>
	<li>Tybbs ($10)</li>
	<li>N. Chin ($10)</li>
	<li>Manjot ($10)</li>
	<li>Phil P. ($10)</li>
	<li>donor39 (backer)</li>
	<li>Pokabu (backer)</li>
</ul>

¡Ustedes son las personas que me inspiran a trabajar en Picocrypt y ofrecerlo de manera gratuita a todos!
	
También, un gran agradecimiento a las siguientes cinco personas que fueron las primeras en donar y apoyar Picocrypt:
<ul>
	<li>W.Graham</li>
	<li>N. Chin</li>
	<li>Manjot</li>
	<li>Phil P.</li>
	<li>E. Zahard</li>
</ul>

También, un gran agradecimiento a estas personas, que han ayudado a traducir Picocrypt y hacerlo más accesible al mundo:
<ul>
	<li>@umitseyhan75 para el turco</li>
	<li>@digitalblossom y @Pokabu26 para el alemán</li>
	<li>@zeeaall para el portugués de Brasil</li>
	<li>@kurpau para el lituano</li>
	<li>u/francirc para el español</li>
	<li>yn para el ruso</li>
	<li>@Etim-Orb para el húngaro</li>
	<li>@Minibus93 para el italiano</li>
	<li>Michel para el fracés</li>
</ul>

Fianlmente, gracias a estas personas/organizaciones que me ayudaron cuando lo necesité:
<ul>
	<li>Fuderal en Discord por ayudarme a configurar un servidor de Discord</li>
	<li>u/greenreddits por el apoyo y comentarios constantes</li>
	<li>u/Tall_Escape por ayudarme a probar Picocrypt</li>
	<li>u/NSABackdoors por hacer un completo trabajo de pruebas</li>
	<li>@samuel-lucas6 por los comentarios, sugerencias y apoyo</li>
	<li><a href="https://privacytools.io">PrivacyToolsIO</a> por incluir en su lista a Picocrypt</li>
	<li><a href="https://privacyguides.org">PrivacyGuides</a> por incluir en su lista a Picocrypt</li>
</ul>
