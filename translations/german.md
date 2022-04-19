<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p>

Picocrypt ist ein sehr kleines (daher <i>Pico</i>), sehr einfaches, aber sehr sicheres Verschlüsselungsprogramm, mit dem Sie Ihre Dateien schützen können. Es wurde entwickelt, um die erste Wahl für die Verschlüsselung zu sein, mit einem Schwerpunkt auf Sicherheit, Einfachheit und Zuverlässigkeit. Picocrypt verwendet die sichere XChaCha20-Chiffre und die Argon2id-Schlüsselableitungsfunktion, um ein hohes Maß an Sicherheit zu bieten, selbst vor Agenturen mit drei Buchstaben wie der NSA. Picocrypt ist auf maximale Sicherheit ausgelegt, ohne Kompromisse bei der Sicherheit einzugehen, und wird mit den Standard x/crypto-Modulen von Go gebaut. <strong>Ihre Privatsphäre und Sicherheit werden angegriffen. Holen Sie sich mit Zuversicht zurück, indem Sie Ihre Dateien mit Picocrypt schützen.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="318" alt="Picocrypt"></p>

# Finanzierung

Bitte spenden Sie für Picocrypt auf <a href="https://opencollective.com/picocrypt">Open Collective</a> (Kryptowährungen werden akzeptiert), um Geld für einen mögliche Sicherheitsaudit durch Cure53 zu sammeln. Da dies ein Projekt ist, in das ich viele Stunden investiere und mit dem ich kein Geld verdiene, kann ich einen Sicherheitsaudit nicht selbst bezahlen. <i>Picocrypt braucht die Unterstützung der Community.</i>

# Downloads

**Wichtig**: Es gibt ein veraltetes und nutzloses Stück Abandonware namens PicoCrypt im Internet, das zuletzt 2005 aktualisiert wurde. PicoCrypt ist in keiner Weise mit Picocrypt (diesem Projekt) verwandt. Stellen Sie sicher, dass Sie Picocrypt nur von diesem Repository herunterladen, um sicherzustellen, dass Sie das authentische und Backdoor-freie Picocrypt erhalten.

## Windows

Picocrypt für Windows ist so einfach, wie es nur geht. Klicken Sie <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.26/Picocrypt.exe">hier</a>, um die neueste, eigenständige, portabel ausführbare Datei für Windows herunterzuladen. Wenn Windows Defender oder Ihr Antivirenprogramm Picocrypt als Virus erkennt, tun Sie bitte Ihren Teil und melden Sie es als Fehlalarm, damit alle davon profitieren können.

## macOS

Picocrypt für macOS ist ebenfalls sehr einfach. Laden Sie Picocrypt <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.26/Picocrypt.app.zip">hier</a> herunter, entpacken Sie die Zip-Datei und starten Sie das darin enthaltene Picocrypt. Wenn Sie Picocrypt nicht öffnen können, weil es nicht von einem verifizierten Entwickler stammt, klicken Sie mit der rechten Maustaste auf Picocrypt und wählen Sie "Öffnen". Wenn Sie immer noch die Warnung erhalten, klicken Sie mit der rechten Maustaste auf Picocrypt und klicken Sie erneut auf "Öffnen" und Sie sollten in der Lage sein, Picocrypt zu starten.

## Linux

Ein Snap ist für Linux verfügbar. Wenn Sie mit einem Debian-basierten System arbeiten, genügt ein einfaches `apt install snapd` und `snap install picocrypt`. Für andere Distributionen wie Fedora sind detaillierte Anweisungen unter https://snapcraft.io/picocrypt verfügbar. Aufgrund der Komplexität der Abhängigkeiten und des statischen Linkings verteile ich keine eigenständigen .deb- oder .rpm-Binärdateien, da diese unzuverlässig wären und den Aufwand nicht wert wären. Snapcraft verwaltet alle Abhängigkeiten und Laufzeiten automatisch und ist der empfohlene Weg, um Picocrypt auf jeder größeren Linux-Distribution auszuführen. Außerdem bietet Snapcraft eine bessere Sicherheit und Containerisierung als Flatpaks und AppImages, was für ein Verschlüsselungstool wie Picocrypt wichtig ist. Wenn Sie es vorziehen, sich nicht mit Canonical zu befassen, denken Sie daran, dass das Kompilieren aus dem Quellcode immer eine Option ist.

## Paranoid Packs

Das Paranoid Pack ist ein komprimiertes Archiv, das ausführbare Dateien für jede Version von Picocrypt enthält, die jemals für Windows, macOS und Linux veröffentlicht wurde. Solange Sie es an einem Ort aufbewahren, auf den Sie zugreifen können, sind Sie in der Lage, es zu öffnen und jede Version von Picocrypt zu verwenden, falls dieses Repository auf mysteriöse Weise verschwindet oder das gesamte Internet abbrennt. Betrachten Sie es als eine Art Saatgut-Tresor für Picocrypt. Solange eine Person das Paranoid Pack in Reichweite hat, kann sie es mit dem Rest der Welt teilen und Picocrypt im Falle von katastrophalen Ereignissen wie dem plötzlichen Abschalten von GitHub oder der Verhaftung durch die NSA funktionsfähig halten (nur für den Fall der Fälle, weißt du?). Der beste Weg, um sicherzustellen, dass Picocrypt auch in vielen Jahrzehnten noch zugänglich ist, besteht darin, ein Paranoid Pack an einem sicheren Ort aufzubewahren. Wenn Sie sich also Sorgen machen, dass Sie in Zukunft nicht mehr auf Picocrypt zugreifen können, dann haben wir hier die Lösung für Sie. Gehen Sie in die Release Sektion und holen Sie sich ihre Kopie.

# Warum PicoCrypt?

Warum sollten Sie Picocrypt anstelle von BitLocker, NordLocker, VeraCrypt, AxCrypt oder 7-Zip verwenden? Hier sind ein paar Gründe, warum Sie Picocrypt wählen sollten:

<ul>
  <li>Im Gegensatz zu NordLocker, BitLocker, AxCrypt und den meisten Cloud-Speicheranbietern sind Picocrypt und seine Abhängigkeiten vollständig quelloffen und überprüfbar. Sie können sich selbst davon überzeugen, dass es keine Hintertüren oder Schwachstellen gibt.</li>
  <li>Picocrypt ist <i>winzig</i>. Während NordLocker mehr als 50 MB und VeraCrypt mehr als 20 MB groß sind, ist Picocrypt nur 2 MB groß, etwa so groß wie ein Foto mittlerer Auflösung. Und das ist noch nicht alles - Picocrypt ist portabel (es muss nicht installiert werden) und erfordert keine Administrator-/Root-Rechte.</li>
  <li>Picocrypt ist einfacher und produktiver zu verwenden als VeraCrypt. Um Dateien mit VeraCrypt zu verschlüsseln, müssten Sie mindestens fünf Minuten damit verbringen, ein Volume einzurichten. Mit der einfachen Benutzeroberfläche von Picocrypt müssen Sie nur Ihre Dateien auf das Fenster ziehen und ablegen, ein Passwort eingeben und auf Start klicken. Alle komplexen Vorgänge werden von Picocrypt intern abgewickelt. Wer sagt denn, dass sichere Verschlüsselung nicht einfach sein kann?</li>
  <li>Picocrypt ist auf Sicherheit ausgelegt. 7-Zip ist ein Archivierungsprogramm und kein Verschlüsselungstool, daher liegt sein Schwerpunkt nicht auf Sicherheit. Picocrypt hingegen wurde mit Sicherheit als oberster Priorität entwickelt. Jeder Teil von Picocrypt existiert aus einem bestimmten Grund und alles, was die Sicherheit von Picocrypt beeinträchtigen könnte, wurde entfernt. Picocrypt ist mit einer Kryptographie aufgebaut, der Sie vertrauen können.</li>
  <li>Picocrypt schützt die Daten nicht nur, sondern authentifiziert sie auch und verhindert so, dass Hacker sensible Daten böswillig verändern. Dies ist nützlich, wenn Sie verschlüsselte Dateien über einen unsicheren Kanal versenden und sicher sein wollen, dass sie unangetastet ankommen.</li>
  <li>Picocrypt schützt verschlüsselte Header-Daten aktiv vor Beschädigung, indem es zusätzliche Reed-Solomon-Paritätsbytes hinzufügt. Wenn also ein Teil des Headers eines Datenträgers (der wichtige kryptografische Komponenten enthält) beschädigt wird (z. B. durch Bitfäule auf der Festplatte), kann Picocrypt den Header trotzdem wiederherstellen und Ihre Daten mit einer hohen Erfolgsquote entschlüsseln. Picocrypt kann auch den gesamten Datenträger mit Reed-Solomon verschlüsseln, um eine Beschädigung Ihrer wichtigen Dateien zu verhindern.</li>
  </ul>

# Vergleich

Hier sehen Sie, wie Picocrypt im Vergleich zu anderen gängigen Verschlüsselungsprogrammen abschneidet.

|                         | Picocrypt      | VeraCrypt         | 7-Zip (GUI)      | NordLocker       | BitLocker        | AxCrypt           |
| --------------          | -------------- | --------------    | --------------   | --------------   | --------------   | --------------    |
| Kostenlos               |✅ Ja           |✅ Ja             |✅ Ja             |🟧 Teilweise      |🟧 Teilweise      |🟧 Teilweise      |
| Open Source             |✅ GPLv3        |✅ Multi          |✅ LGPL           |❌ Nein           |❌ Nein           |❌ Nein           |
| Plattformübergreifend   |✅ Ja           |✅ Ja             |❌ Nein           |❌ Nein           |❌ Nein           |❌ Nein           |
| Größe                   |✅ 2MB          |❌ 20MB           |✅ 2MB            |❌ 60MB           |✅ Included       |🟧 8MB            |
| Portable                |✅ Ja           |✅ Ja             |❌ Nein           |❌ Nein           |✅ Ja             |✅ Ja             |
| Berechtigungen          |✅ Keine        |❌ Administrator  |❌ Administrator  |❌ Administrator  |❌ Administrator  |❌ Administrator  |
| Benutzerfreundlichkeit  |✅ Leicht       |❌ Schwer         |✅ Leicht         |🟧 Mittel         |🟧 Mittel         |✅ Leicht         |
| Schlüssel Ableitung     |✅ Argon2       |🆗 PBKDF2         |❌ SHA256         |✅ Argon2         |❓ Unbekannt       |🆗 PBKDF2         |
| Datenintegrität         |✅ Always       |❌ Nein           |❌ Nein           |✅ Always         |❓ Unbekannt       |✅ Always         |
| Reed-Solomon            |✅ Ja           |❌ Nein           |❌ Nein           |❌ Nein           |❌ Nein           |❌ Nein           |
| Komprimierung           |✅ Ja           |❌ Nein           |✅ Ja             |❌ Nein           |✅ Ja             |✅ Ja             |
| Telemetrie              |✅ Keine        |✅ Keine          |✅ Keine          |❌ Analytics      |❓ Unbekannt       |❌ Accounts       |
| Geprüft                 |🟧 Geplant      |✅ Ja             |❌ Nein           |❓ Unbekannt       |❓ Unbekannt       |❌ Nein           |

# Eigenschaften

Picocrypt ist ein sehr einfaches Werkzeug, und die meisten Benutzer werden intuitiv in wenigen Sekunden verstehen, wie man es benutzt. Um Ihre Dateien zu verschlüsseln, müssen Sie lediglich Ihre Dateien auf das Fenster ziehen, ablegen, ein Passwort eingeben und auf "Start" klicken. Ziemlich einfach, oder?

Picocrypt ist zwar einfach, will aber in den Händen erfahrener und fortgeschrittener Benutzer auch leistungsstark sein. Daher gibt es einige zusätzliche Optionen, die Sie je nach Ihren Bedürfnissen nutzen können.

<ul>
  <li><strong>Passwortgenerator</strong>: Picocrypt bietet einen sicheren Passwortgenerator, mit dem Sie kryptografisch sichere Passwörter erstellen können. Sie können die Länge des Passworts sowie die Art der zu verwendenden Zeichen anpassen.</li>
  <li><strong>Kommentare</strong>: Hier können Sie Notizen, Informationen und Text zusammen mit der Datei speichern (Diese werden nicht verschlüsselt). Sie können zum Beispiel eine Beschreibung der zu verschlüsselnden Datei eingeben, bevor Sie sie an jemanden schicken. Wenn die Person, an die Sie die Datei geschickt haben, diese in Picocrypt ablegt, wird Ihre Beschreibung angezeigt.</li>
  <li><strong>Schlüsseldateien</strong>: Picocrypt unterstützt die Verwendung von Schlüsseldateien als eine zusätzliche Form der Authentifizierung (oder als die einzige Art der Authentifizierung). Sie können nicht nur mehrere Schlüsseldateien verwenden, sondern auch verlangen, dass die richtige Reihenfolge der Schlüsseldateien vorhanden sein muss, damit eine erfolgreiche Entschlüsselung stattfinden kann. Ein besonders guter Anwendungsfall für mehrere Schlüsseldateien ist die Erstellung eines gemeinsamen Datenträgers, bei dem jede Person eine Schlüsseldatei besitzt und alle (und ihre Schlüsseldateien) vorhanden sein müssen, um den gemeinsamen Datenträger zu entschlüsseln.</li>
  <li><strong>Paranoidmodus</strong>: In diesem Modus werden Ihre Daten sowohl mit XChaCha20 als auch mit Serpent kaskadenartig verschlüsselt, und zur Authentifizierung der Daten wird HMAC-SHA3 anstelle von BLAKE2b verwendet. Dies wird für den Schutz streng geheimer Dateien empfohlen und bietet, in der Praxis, die höchste erreichbare Sicherheitsstufe. Damit ein Hacker Ihre verschlüsselten Daten entschlüsseln kann, müssen sowohl die XChaCha20-Chiffre als auch die Serpent-Chiffre geknackt werden, vorausgesetzt, Sie haben ein gutes Passwort gewählt. Man kann mit Sicherheit sagen, dass Ihre Dateien in diesem Modus nicht zu knacken sind.</li>
  <li><strong>Reed-Solomon</strong>: Diese Funktion ist sehr nützlich, wenn Sie vorhaben, wichtige Daten bei einem Cloud-Anbieter oder auf einem externen Medium über einen längeren Zeitraum zu archivieren. Wenn diese Funktion aktiviert ist, verwendet Picocrypt den Reed-Solomon-Fehlerkorrekturcode, um 8 zusätzliche Bytes für jeweils 128 Bytes hinzuzufügen, um eine Beschädigung der Datei zu verhindern. Das bedeutet, dass bis zu ~3 % Ihrer Datei beschädigt sein können und Picocrypt trotzdem in der Lage ist, die Fehler zu korrigieren und Ihre Dateien ohne Beschädigung zu entschlüsseln. Natürlich kann Picocrypt Ihre Dateien nicht vollständig wiederherstellen, wenn sie sehr stark beschädigt sind (z. B. wenn Ihnen die Festplatte heruntergefallen ist), aber es wird sein Bestes tun, um zu retten, was es kann. Beachten Sie, dass diese Option die Ver- und Entschlüsselung erheblich verlangsamt.</li>
  <li><strong>Entschlüsselung erzwingen</strong>: Picocrypt überprüft bei der Entschlüsselung automatisch die Integrität der Datei. Wenn die Datei geändert wurde oder beschädigt ist, löscht Picocrypt die Ausgabe zur Sicherheit des Benutzers automatisch. Wenn Sie die beschädigten oder veränderten Daten nach der Entschlüsselung behalten möchten, aktivieren Sie diese Option. Wenn diese Option aktiviert ist und die Reed-Solomon-Funktion für die verschlüsselte Datei verwendet wurde, versucht Picocrypt außerdem, während der Entschlüsselung so viel wie möglich von der Datei wiederherzustellen.</li>
  <li><strong>Aufteilung von Dateien in Teile (Chunks)</strong>: Sie haben keine Lust, sich mit riesigen Dateien herumzuschlagen? Kein Grund zur Sorge! Mit Picocrypt können Sie Ihre Ausgabedatei in Stücke benutzerdefinierter Größe aufteilen, so dass große Dateien überschaubarer werden und leichter zu Cloud-Anbietern hochgeladen werden können. Wählen Sie einfach eine Chunkgröße (KiB, MiB, GiB or TiB) und geben Sie die gewünschte Anzahl für diese Einheit ein. Um die Chunks zu entschlüsseln, ziehen Sie einfach einen von ihnen in Picocrypt, und die Chunks werden bei der Entschlüsselung automatisch wieder zusammengefügt.</li>
</ul>

# Sicherheit

Weitere Informationen darüber, wie Picocrypt die Kryptographie handhabt, finden Sie <a href="Internals.md">hier</a> in den technischen Details. Wenn Sie sich Sorgen um meine Sicherheit oder die Sicherheit dieses Projekts machen, kann ich Ihnen versichern, dass dieses Repository nicht gekapert oder manipuliert werden kann. Ich habe 2FA (TOTP) auf allen Konten aktiviert, die mit Picocrypt verbunden sind (GitHub, Google, Reddit, Ubuntu One/Snapcraft, Discord, etc.), zusätzlich zu einer vollständigen Festplattenverschlüsselung auf all meinen tragbaren Geräten. Zur weiteren Absicherung verwendet Picocrypt meine isolierten Forks von Abhängigkeiten, und ich hole Upstream nur, wenn ich mir die Änderungen angesehen habe und glaube, dass es keine Sicherheitsprobleme gibt. Das bedeutet, dass, wenn eine Abhängigkeit gehackt oder vom Autor gelöscht wird, Picocrypt meinen Fork davon verwendet und davon völlig unberührt bleibt. Sie können Picocrypt getrost verwenden.

# Community

Hier finden Sie einige Stellen, an denen Sie über Picocrypt auf dem Laufenden bleiben und sich engagieren können:

<ul>
	<li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
	<li><a href="https://discord.gg/8QM4A2caxH">Discord</a></li>
</ul>

Ich empfehle Ihnen dringend, dem Picocrypt-Subreddit beizutreten, da alle Updates und Umfragen dort veröffentlicht werden. Denken Sie daran, nur diesen sozialen Netzwerken zu vertrauen und sich vor Hackern in Acht zu nehmen, die versuchen könnten, sich für mich auszugeben. Ich werde Sie niemals nach Ihrem Passwort fragen, und jeder, der das tut, ist nicht ich. Ich werde Sie niemals auffordern, eine Datei von einem verdächtigen Link herunterzuladen, und jeder, der das tut, ist nicht ich.

# Stargazers(GitHub-Sterne)

Wie läuft es mit Picocrypt? Werfen Sie einen Blick darauf, um es herauszufinden.
![Stargazers Over Time](https://starchart.cc/HACKERALERT/Picocrypt.svg)

# Spenden

Wenn Sie Picocrypt nützlich finden, ziehen Sie bitte in Betracht, mir ein Trinkgeld über mein <a href="https://paypal.me/evanyiwensu">PayPal</a> dazulassen. Ich stelle diese Software völlig kostenlos zur Verfügung und würde mich freuen, einige Unterstützer zu haben, die mich motivieren, meine Arbeit an Picocrypt fortzusetzen.

# Danksagungen

Ich möchte mich von ganzem Herzen bei den Leuten von Open Collective bedanken, die einen wichtigen Beitrag geleistet haben:

<ul>
  <li>YellowNight ($818)</li>
	<li>jp26 ($50)</li>
	<li>guest-116103ad ($50)</li>
	<li>Tybbs ($10)</li>
	<li>N. Chin ($10)</li>
	<li>Manjot ($10)</li>
	<li>Phil P. ($10)</li>
	<li>donor39 (Unterstützer)</li>
	<li>Pokabu (Unterstützer)</li>
</ul>

Ihr seid die Menschen, die mich inspirieren, an Picocrypt zu arbeiten und es jedem kostenlos zur Verfügung zu stellen!

Ein großes Dankeschön geht auch an die folgenden fünf Personen, die als erste gespendet und Picocrypt unterstützt haben:

<ul>
	<li>W.Graham</li>
	<li>N. Chin</li>
	<li>Manjot</li>
	<li>Phil P.</li>
	<li>E. Zahard</li>
</ul>

Ein großer Dank geht auch an diese Leute, die geholfen haben, Picocrypt zu übersetzen und der Welt zugänglich zu machen:

<ul>
	<li>@umitseyhan75 für Türkisch</li>
	<li>@digitalblossom & Pokabu26 für Deutsch</li>
	<li>@zeeaall für Brasilianisches Portugiesisch</li>
	<li>@kurpau für Litauisch</li>
	<li>u/francirc für Spanisch</li>
	<li>yn für Russisch</li>
	<li>@Etim-Orb für Ungarisch</li>
	<li>@Minibus93 für Italienisch</li>
  <li>Michel für Französisch</li>
</ul>

Schließlich möchte ich mich bei diesen Leuten/Organisationen dafür bedanken, dass sie mir im Bedarfsfall geholfen haben:

<ul>
	<li>Fuderal auf Discord für die Hilfe bei der Einrichtung eines Discord-Servers</li>
	<li>u/greenreddits für ständiges Feedback und Unterstützung</li>
	<li>u/Tall_Escape für die Hilfe beim Testen von Picocrypt</li>
	<li>u/NSABackdoors für die Durchführung zahlreicher Tests</li>
	<li>@samuel-lucas6 für Feedback, Vorschläge und Unterstützung</li>
  <li><a href="https://privacytools.io">PrivacyToolsIO</a> für die Aufnahme von Picocrypt</li>
	<li><a href="https://privacyguides.org">PrivacyGuides</a> für die Aufnahme von Picocrypt</li>
</ul>
