<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p> 

A Picocrypt egy nagyon apró (lásd <i>Pico</i>), nagyon egyszerű, de mégis nagyon biztonságos adattitkosító eszköz amit hasznáhatsz a fájljaid megvédésére, checksum generálásra, és sok minden másra is. Egy <i>go-to</i> adattitkosító eszköznek lett tervezve, amivel a fókusz a biztonnságon, egyszerűségen, és megbízhatóságon van. A Picocrypt a biztonságos XChaCha20 ciphert és az SHA3 hash funkciót használja, hogy magas szintű biztonságot nyújtson, még három-betűs ügynökségek ellen is, mint az NSA. A maximum biztonságra lett tervezve, így biztonásági téren nem köt semmilyen kompromisszumot, és a Go standard x/crypto moduljaival lett kiépítve. <strong>A magánéleted és az adatbiztonságod támadás alatt áll. Szerezd vissza ezeket magabiztosan úgy, hogy a Picocrypt-tel véded a fájljaidat.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="384" alt="Picocrypt"></p>

# Finanszírozás
Kérlek támogasd a Picocryptet <a href="https://opencollective.com/picocrypt">Open Collective</a>-en (crypto-t is elfogadunk), hogy elég pénzt gyűjtsünk össze egy lehetséges auditra a Cure53-tól. Mivel ez egy olyan projekt amibe én rengeteg órát fektetek bele és nem keresek vele semmi pénzt, saját magam nem tudok fizetni az auditért. <i>Picocrypt-nak szüksége van támogatásra a közösségétől.</i>

# Letöltések
<strong>Fontos:</strong> Van egy elavult és hasznavehetetlen abandonware az interneten, amit PicoCrypt-nak hívnak és ami legutóbb 2005-ben lett frissítve. PicoCrypt nem áll kapcsolatban semmilyen formában a Picocrypt-tel (ezzel a projekttel). Győződj meg róla, hogy csak ebből a repository-ból töltöd le a Picocryptet, hogy biztosan az eredeti és backdoor mentes Picocrypted szerzed meg.

## Windows
A Picocrypt Windows-ra olyan egyszerű amennyire csak az lehet. Ahhoz, hogy letöltsd a legfrissebb, egyedülálló és hordozható futtatóprogramot, kattints <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.21/Picocrypt.exe">ide</a>. Ha a Windows Defender, vagy a vírusírtód vírusként jelöli meg a Picocryptet, kérlek vedd ki a részed azzal, hogy beadod, mint hibás találat, így segítve mindannyiunknak.

## macOS
Picocrypt a macOS-en szintén egyszerű. Töltsd le a Picocrypt-et <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.21/Picocrypt.app.zip">itt</a>, csomagold ki a zip fájlt, és futtasd a Picocryptet ami benne van. H a nem tudod megnyitni a Picocryptet, mert nem egy igazolt fejlesztőtől van, jobb-klikkelj a Picocryptre és válaszd ki a "Megnyitás" lehetőséget. Ha ez után is megkapod a figyelmeztetést, jobb-klikkelj a Picocryptre és kattints a "megnyitás" lehetőségre ismét és ezután el kéne tudnod indítani a Picocryptet.

## Linux
Linuxhoz egy Snap elérhető. Feltéve, hogy egy Debian alapú rendszeren vagy, egy egyszerű `apt install snapd` és `snap install picocrypt` elég lesz. Más disztrókhoz, mint a Fedora, részletes utasítások elérhetőek a https://snapcraft.io/picocrypt linken. A dependency-k komplexitása miatt és statikus linkelés miatt, én nem készítek .deb és .rpb binary-kat, mivel megbízhatatlanok lennének és nem éri meg a bonyodalmakat. A Snapcraft autómatikusan kezel minden dependency-t és futtatót és ez az ajánlott módja a Picocrypt futtatásának bármely fő Linux disztibúción. Ezen felül, a Snapcraft jobb biztonságot és elkülönítést nyújt, mint a Flatpak-ek és AppImage-ek, amik fontosak egy titkosító eszköznek mint a Picocrypt. Ha nem szeretnél a Canonical-lal vesződni, emlékeztetlek, hogy a forrásból buildelés mindig lehetséges.

# Miért a Picocrypt?
Miért használd a Picocryptet a BitLocker, NordLocker, VeraCrypt, AxCrypt, vagy 7-Zip helyett? Itt egy pár ok, hogy miért válaszd a Picocryptet:
<ul>
	<li>Nem úgy mint a NordLocker, BitLocker, AxCrypt, és a legtöbb felhőtárhely szolgáltató, Picocrypt és annak dependency-jei teljesen nyílt forráskódúak és auditolhatók. Saját magad is ellenőrizni tudod, hogy nincs-e benne beckdoor, vagy hiba.</li>
	<li>A Picocrypt <i>kicsi</i>. Míg a NordLocker 100MB fölött és VeraCrypt 30MB fölött van, addig a Picocrypt csupán csak 3MB, ami körülbelül egy magas felbontású kép mérete. És ez nem minden - A Picocrypt hordozható (nem kell, hogy telepítve legyen) és nem szükségel adminisztrátori/root hozzáférést.</li>
	<li>A Picocryptet könnyeb használni és sokkal produktívabb mint a VeraCrypt. Ahhoz, hogy fájlokat titkosíts a VeraCrypt-tel, legalább öt percet kell eltöltened azzal, hogy létrehozz egy kötetet. A Picocrypt szimpla felhasználói felületével, csak annyit kell csinálnod, hogy behúzod a fájljaidat, beírsz egy jelszót és kiválasztod az Indítást. Az összes komplex folyamat a Picocrypt álltal és azon belül van kezelve. Ki mondta, hogy a biztonságos titkosítás nem lehet egyszerű?</li>
	<li>A Picocrypt a biztonságra lett tervezve. A 7-Zip egy archíváló segédprogram és nem egy titkosító eszköz, ezért a fókusza nem a biztonságra helyeződik. Picocrypt is an archive utility and not an encryption tool, so its focus is not on security. A Picocrypt mindamellett biztonsággal - mint legfő prioritás - lett tervezve. A Picocrypt minden egyes része okkal van ott ahol van és bármi ami kihatna Picocrypt biztonságára el van távolítva. A Picocrypt olyan kriptográfiát tartalmaz amiben megbízhatsz.</li>
	<li>A Picocrypt hitelesíti az adatokat azok védelmén kívül, így meggátolja, hogy hekkerek rosszindulatú módosításokat hajtsanak végre érzékeny információn. Ez hasznos amikor egy nem biztonságos csatornán keresztül küldesz titkosított fájlokat és meg akarsz róla győződni, hogy atok érintetlenül érkeznek meg. Picocrypt HMAC-SHA3-at használ hitelesítésre, ami egy magas biztonságú hash funkció egy jól ismert felépítéssel.</li>
	<li>A Picocrypt aktívan védi a titkosított header adatokat az adatsérüléstől Reed-Solomon paritás byteok hozzáadásával, így ha a kötet header-jének egy része (ami fontos kriptográfiai komponenseket tartalmaz) sérül (pl. merevlemez byte-rot), attól a Picocrypt még mindig helyre tudja állítani a headert és vissza tudja fejteni az adatokat nagy sikerrel. Picocrypt ezen kívül encode-olni tudja a teljes kötetet Reed-Solomonnal, hogy megelőzzön bármilyen korruptálódást ami a fontos fájljaidat veszélyeztetné.</li>
</ul>

Még mindig nem vagy meggyőzve? Itt lent olvashatsz még több okot, hogy hogyan tűnik Picocrypt a többi közül.

# Tulajdonságok
A Picocrypt egy nagyon egyszerű eszköz és a legtöbb felhasználó ösztönösen rá fog jönni, hogy hogyan kell használni pár másodpercen belül. Lap szinten a fájljaid behúzása, egy jelszó beírása és az Indítás kiválasztása minden amire szükség van a fájljaid titkosítására. Elég egyszerű, ugye?

Bár egyszerű, Picocrypt ezen felül törekszik arra, hogy rendkívül hatékony legyen tapasztalt és haladó felhasználók kezében. Így van pár további opció amit használhatsz a szükségleteidnek megfelelően.
<ul>
	<li><strong>Jelszó generátor</strong>: Picocrypt provides a secure password generator that you can use to create cryptographically secure passwords. You can customize the password length, as well as the types of characters to include.</li>
	<li><strong>File metadata</strong>: Use this to store notes, information, and text along with the file (it won't be encrypted). For example, you can put a description of the file you're encrypting before sending it to someone. When the person you sent it to drops the file into Picocrypt, your description will be shown to that person.</li>
	<li><strong>Keyfiles</strong>: Picocrypt supports the use of keyfiles as an additional form of authentication. Not only can you use multiple keyfiles, but you can also require the correct order of keyfiles to be present, for a successful decryption to occur. A particularly good use case of multiple keyfiles is creating a shared volume, where each person holds a keyfile, and all of them (and their keyfiles) must be present in order to decrypt the shared volume.</li>
	<li><strong>Fast mode</strong>: Using this mode will greatly speed up encryption/decryption. In this mode, BLAKE2b will be used to authenticate data instead of SHA3, and Argon2 parameters will be lowered. Doing this provides higher speeds, but at a lower security margin. If all you need to do is encrypt some low-sensitivity files, this option can be a useful and performant choice.</li>
	<li><strong>Paranoid mode</strong>: Using this mode will encrypt your data with both XChaCha20 and Serpent in a cascade fashion. This is recommended for protecting top-secret files and provides the highest level of practical security attainable. In order for a hacker to crack your encrypted data, both the XChaCha20 cipher and the Serpent cipher must be broken, assuming you've chosen a good password.</li>
	<li><strong>Prevent corruption using Reed-Solomon</strong>: This feature is very useful if you are planning to archive important data on a cloud provider or external medium for a long time. If checked, Picocrypt will use the Reed-Solomon error correction code to add 8 extra bytes for every 128 bytes to prevent file corruption. This means that up to ~3% of your file can corrupt and Picocrypt will still be able to correct the errors and decrypt your files with no corruption. Of course, if your file corrupts very badly (e.g., you dropped your hard drive), Picocrypt won't be able to fully recover your files, but it will try its best to recover what it can. Note that this option will slow down encryption and decryption considerably.</li>
	<li><strong>Keep decrypted output even if it's corrupted or modified</strong>: Picocrypt automatically checks for integrity upon decryption. If the file has been modified or is corrupted, Picocrypt will automatically delete the output for the user's safety. If you want to keep the corrupted or modified data after decryption, check this option. Also, if this option is checked and the Reed-Solomon feature was used on the encrypted file, Picocrypt will attempt to recover as much of the file as possible during decryption.</li>
	<li><strong>Split files into chunks</strong>: Don't feel like dealing with gargantuan files? No worries! With Picocrypt, you can choose to split your output file into custom-sized chunks, so large files can become more manageable and easier to upload to cloud providers. Simply choose a unit (KiB, MiB, or GiB) and enter your desired number for that unit. To decrypt the chunks, simply drag one of them into Picocrypt, and the chunks will be automatically recombined during decryption.</li>
</ul>

In addition to these comprehensive options for encryption and decryption, Picocrypt also provides a checksum generator for validating the integrity of sensitive files, which supports numerous hash functions like MD5, BLAKE2, and SHA3.

# Security
For more information on how Picocrypt handles cryptography, see <a href="Internals.md">Internals</a> for the technical details. If you're worried about the safety of me or this project, let me assure you that this repository won't be hijacked or backdoored. I have 2FA (TOTP) enabled on all accounts with a tie to Picocrypt (GitHub, Google, Reddit, Ubuntu One/Snapcraft, Discord, etc.), in addition to full-disk encryption on all of my portable devices. For further hardening, Picocrypt uses my isolated forks of dependencies and I fetch upstream only when I have taken a look at the changes and believe that there aren't any security issues. This means that if a dependency gets hacked or deleted by the author, Picocrypt will be using my fork of it and remain completely unaffected. You can feel confident about using Picocrypt.

# Community
Here are some places where you can stay up to date with Picocrypt and get involved:
<ul>
	<li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
	<li><a href="https://discord.gg/8QM4A2caxH">Discord</a></li>
</ul>
I highly recommend you join Picocrypt's subreddit because all updates and polls will be posted there. Remember to only trust these social networks and be aware of hackers that might try to impersonate me. I will never ask you for your password, and anyone who does is not me. I will never tell you to download a file from a suspicious link, and anyone who does is not me.

# Stargazers
How's Picocrypt doing? Take a look below to find out.
[![Stargazers over time](https://starchart.cc/HACKERALERT/Picocrypt.svg)](https://starchart.cc/HACKERALERT/Picocrypt)

# Donations
If you find Picocrypt useful, please consider tipping my <a href="https://paypal.me/evanyiwensu">PayPal</a>. I'm providing this software completely free of charge, and would love to have some supporters that will motivate me to continue my work on Picocrypt.

# Thank You's
A thank you from the bottom of my heart to the people on Open Collective who have made a significant contribution:
<ul>
	<li>jp26 ($50)</li>
	<li>Tybbs ($10)</li>
	<li>N. Chin ($10)</li>
	<li>Manjot ($10)</li>
	<li>Phil P. ($10)</li>
</ul>
You are the people who inspire me to work on Picocrypt and provide it free of charge to everyone!

Also, a huge thanks to the following list of five people, who were the first to donate and support Picocrypt:
<ul>
	<li>W.Graham</li>
	<li>N. Chin</li>
	<li>Manjot</li>
	<li>Phil P.</li>
	<li>E. Zahard</li>
</ul>

As well, a great thanks to these people, who have helped translate Picocrypt and make it more accessible to the world:
<ul>
	<li>@umitseyhan75 for Turkish</li>
	<li>@digitalblossom for German</li>
	<li>@zeeaall for Brazilian Portuguese</li>
	<li>@kurpau for Lithuanian</li>
	<li>u/francirc for Spanish</li>
	<li>yn for Russian</li>
</ul>

Finally, thanks to these people for helping me out when needed:
<ul>
	<li>Fuderal on Discord for helping me setup a Discord server</li>
	<li>u/greenreddits for constant feedback and support</li>
	<li>u/Tall_Escape for helping me test Picocrypt</li>
	<li>u/NSABackdoors for doing plenty of testing</li>
</ul>
