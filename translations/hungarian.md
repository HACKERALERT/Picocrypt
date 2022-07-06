<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p> 

Picocrypt egy nagyon apró (tehát <i>Pico</i>), nagyon egyszerű, de mégis nagyon biztonságos adattitkosító eszköz amit a fájljaid megvédésére használhatsz. Egy <i>go-to</i> adattitkosító eszköznek lett kialakítva, ahol a fókusz a biztonságon, egyszerűségen, és megbízhatóságon van. Picocrypt a biztonságos XChaCha20 ciphert és az Argon2id kulcs deriváció funkciót használja, hogy magas szintű biztonságot nyújtson, még három betűs ügynökségek ellen is, mint az NSA. A maximális biztonságra lett tervezve, így nem köt semmilyen kompromisszumot ilyen téren, valamint a Go standard x/crypto moduljaival lett kiépítve. <strong>A magánéleted és az biztonságod támadás alatt áll. Védd a fájljaidat Picocrypttel és szerezd vissza ezeket magabiztossággal.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="318" alt="Picocrypt"></p>

# Finanszírozás
Kérlek támogasd a Picocryptet <a href="https://opencollective.com/picocrypt">Open Collective</a>-en (crypto-t is elfogadunk), hogy elég pénzt gyűjtsünk össze egy lehetséges auditra a Cure53-tól. Mivel ez egy projekt amibe én rengeteg órát fektetek bele és nem keresek vele semmi pénzt, saját magam nem tudok fizetni az auditért. <i>Picocrypt-nak szüksége van támogatásra a közösségétől.</i>

# Letöltések
**Fontos**: Van egy elavult és hasznavehetetlen abandonware az interneten, amit PicoCrypt-nak hívnak és ami legutóbb 2005-ben lett frissítve. PicoCrypt nem áll kapcsolatban semmilyen formában a Picocrypt-tel (ezzel a projekttel). Győződj meg róla, hogy csak ebből a repository-ból töltöd le a Picocryptet, hogy biztosan az eredeti és backdoor mentes Picocryptet szerzed be.

## Windows
A Picocrypt Windowshoz olyan egyszerű amennyire az csak lehet. Ahhoz, hogy letöltsd a legfrissebb, egyedülálló és hordozható futtatóprogramot Windowshoz, kattints <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.27/Picocrypt.exe">ide</a>. Ha a Windows Defender, vagy a vírusírtód vírusként jelöli meg a Picocryptet, kérlek vedd ki a részed azzal, hogy benyújtod, mint hibás találatot, így segítve mindannyiunknak.

## macOS
Picocrypt a macOS-hez szintén egyszerű. Töltsd le a Picocrypt-et <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.27/Picocrypt.app.zip">itt</a>, csomagold ki a zip fájlt, és futtasd Picocryptet, ami benne van. Ha nem tudod megnyitni a Picocryptet, mert nem egy igazolt fejlesztőtől van, jobb-klikkelj a Picocryptre és válaszd ki a "Megnyitás" lehetőséget. Ha ezután is megkapod a figyelmeztetést, jobb-klikkelj a Picocryptre és kattints a "Megnyitás" lehetőségre ismét és ezután el kéne tudnod indítani a Picocryptet.

## Linux
Több mód van a Picocrypt használatának Linuxon. A javasolt módja a Picocrypt .deb fájlból való telepítése <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.27/Picocrypt.deb">innen</a> (Debian 11+ és Ubuntu 20+). Ha a .deb nem megfelelő vagy nem egy Debian alapú disztrót használsz, nyugodtan használd az AppImage-et <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.27/Picocrypt.AppImage">innen</a>. Ha egyik opció sem működik fentről, telepítheted Picocryptot akár Snapcraftról is, aminek az összes disztrón működnie kéne. Az instrukciókat Snapcrafthoz megtalálod <a href="https://snapcraft.io/picocrypt">itt</a>.

## Paranoid csomagok
A Paranoid Pack egy tömörített archívum ami futtatóprogramokat tartalmaz az összes kiadott Picocrypt verzióhoz valaha, Windowshoz, macOS-hez, és Linuxhoz. Ameddig ezt olyan helyen tárolod amihez hozzá tudsz férni, képes leszel megnyitni és használni Picocrypt bármely verzióját ha netán ez a repository rejtélyes módon eltűnik vagy leég az egész internet. Gondolj úgy erre mint egy Picocrypt magbunker. Ameddig egy személy elérhető helyen tartja a paranoid csomagot, meg tudja azt osztani a világgal, így életben tartva Picocryptet katasztófikus események esetén mint a GitHub hirtelen leállása, vagy ha elfog engem az NSA (a biztonság kedvéért, érted?). A legjobb módja a Picocrypt akár évtizedek múlti hozzáférhetőségének biztosításának, a Paranoid Pack biztonságos helyen tárolása. Tehát, ha aggódsz a miatt, hogy nem fogsz tudni hozzáférni a Picocrypthez a jövőben, hát, itt a megoldás. Csak menj a Releasek felületre és szerezz be egy másolatot.

# Miért a Picocrypt?
Miért használd a Picocryptet a BitLocker, NordLocker, VeraCrypt, AxCrypt, vagy 7-Zip helyett? Itt van pár ok, hogy miért válaszd a Picocryptet:
<ul>
	<li>Nem úgy mint a NordLocker, BitLocker, AxCrypt, és a legtöbb felhőtárhely szolgáltató, Picocrypt és annak dependenciái teljesen nyílt forráskódúak és auditolhatók. Saját magad is ellenőrizni tudod, hogy nincs-e benne backdoor, vagy hiba.</li>
	<li>A Picocrypt <i>kicsi</i>. Míg a NordLocker 50MB fölött és VeraCrypt 20MB fölött van, addig a Picocrypt csupán csak 2MB, ami körülbelül egy magas felbontású kép mérete. És ez nem minden - A Picocrypt hordozható (nem kell, hogy telepítve legyen) és nem igényel adminisztrátor/root hozzáférést.</li>
	<li>A Picocryptet könnyeb és sokkal produktívabb használni mint a VeraCryptet. Ahhoz, hogy fájlokat titkosíts a VeraCrypt-tel, legalább öt percet el kell töltened egy kötet létrehozásával. A Picocrypt szimpla felhasználói felületével, csak annyit kell tenned, hogy ráhúzod a fájljaidat, beírsz egy jelszót és kiválasztod az Indítást. Az összes komplex folyamat a Picocrypt álltal és saját magán belül van kezelve. Ki mondta, hogy biztonságos titkosítás nem lehet egyszerű?</li>
	<li>A Picocrypt a biztonságra lett tervezve. A 7-Zip egy archíváló program és nem egy titkosító eszköz, ezért a fókusza nem a biztonságra helyeződik. A Picocrypt azonban biztonsággal lett tervezve, mint legfőbb prioritás. A Picocrypt minden egyes része okkal van ott ahol, és bármi ami kihatna a Picocrypt biztonságára el van távolítva. A Picocrypt olyan kriptográfiát alkalmaz amiben megbízhatsz.</li>
	<li>A Picocrypt hitelesíti az adatokat azok védelmén kívül, így meggátolja, hogy hekkerek érzékeny információn rosszindulatú módosításokat hajtsanak végre. Ez hasznos amikor titkosított fájlokat küldesz egy nem biztonságos csatornán keresztül és meg akarsz róla győződni, hogy azok érintetlenül érkeznek meg.</li>
	<li>A Picocrypt aktívan védi a titkosított header adatokat a sérüléstől Reed-Solomon parity byteok hozzáadásával, így ha a kötet header-jének egy része (ami fontos kriptográfiai komponenseket tartalmaz) sérül (pl. merevlemez bit rot), attól a Picocrypt még mindig nagy valószínűséggel helyre tudja állítani a headert és vissza tudja fejteni az adatokat. Picocrypt ezen kívül encode-olni tudja a teljes kötetet Reed-Solomonnal, hogy megelőzzön bármilyen adatsérülést ami a fontos fájljaidat érintené.</li>
</ul>

# Összehasonlítás

A Picocrypt így hasonlítható össze más népszerű titkosítási eszközökkel.

|                | Picocrypt      | VeraCrypt      | 7-Zip (GUI)    | BitLocker      | Cryptomator    | NordLocker     | AxCrypt        |
| -------------- | -------------- | -------------- | -------------- | -------------- | -------------- | -------------- | -------------- |
| Ingyenes       |✅ Igen        |✅ Igen         |✅ Igen        |🟧 Részben      |✅ Igen        |🟧 Részben      |🟧 Részben     |
| Nyílt forráskód|✅ GPLv3       |✅ Multi        |✅ LGPL        |❌ No           |✅ GPLv3       |❌ Nem          |❌ Nem         |
| Cross-Platform |✅ Igen        |✅ Igen         |❌ Nem         |❌ No           |✅ Yes         |❌ Nem          |❌ Nem         |
| Méret          |✅ 2MB         |❌ 20MB         |✅ 2MB         |✅ Tartalmazza  |❌ 50MB        |❌ 60MB         |🟧 8MB         |
| Hordozható     |✅ Igen        |✅ Igen         |❌ Nem         |✅ Igen         |❌ Nem         |❌ Nem          |✅ Igen        |
| Engedélyek     |✅ Semmilyen   |❌ Admin        |❌ Admin       |❌ Admin        |❌ Admin       |❌ Admin        |❌ Admin       |
| Kényelem       |✅ Könnyű      |❌ Nehéz        |✅ Könnyű      |🟧 Közepes      |🟧 Közepes     |🟧 Közepes      |✅ Könnyű      |
| Cipher         |✅ XChaCha20   |✅ AES-256      |✅ AES-256     |🟧 AES-128      |✅ AES-256     |✅ AES-256      |🟧 AES-128     |
| Kulcs deriváció|✅ Argon2      |🆗 PBKDF2       |❌ SHA256      |❓ Ismeretlen   |✅ Scrypt      |✅ Argon2       |🆗 PBKDF2      |
| Adat integritás|✅ Mindig      |❌ Nem          |❌ Nem         |❓ Ismeretlen   |✅ Mindig      |✅ Mindig       |✅ Mindig      |
| Reed-Solomon   |✅ Igen        |❌ Nem          |❌ Nem         |❌ Nem          |❌ Nem         |❌ Nem          |❌ Nem         |
| Tömörítés      |✅ Igen        |❌ Nem          |✅ Igen        |✅ Igen         |❌ Nem         |❌ Nem          |✅ Igen        |
| Telemetria     |✅ Semmilyen   |✅ Semmilyen    |✅ Semmilyen   |❓ Ismeretlen   |✅ Semmilyen   |❌ Analitika    |❌ Fiókok      |
| Auditolt       |🟧 Tervezett   |✅ Igen         |❌ Nem         |❓ Ismeretlen   |✅ Igen        |❓ Ismeretlen   |❌ Nem         |

# Sajátosságok
A Picocrypt egy nagyon egyszerű eszköz és a legtöbb felhasználó pár másodpercen belül ösztönösen rá fog jönni, hogy hogyan kell használni. Alap szinten, a fájljaid egyszerű behúzása, egy jelszó beírása és az Indítás kiválasztása minden amire szükség van a fájljaid titkosításához. Elég egyszerű, ugye?

Bár egyszerű, Picocrypt törekszik arra is, hogy rendkívül hatékony legyen tapasztalt és haladó felhasználók kezében. Így van pár további opció amit használhatsz a szükségleteidnek megfelelően.
<ul>
	<li><strong>Jelszó generátor</strong>: Picocrypt nyújt egy biztonságos jelszó generátort amit kriptográfiailag biztonságos jelszavak létrehozására használhatsz. Tesztreszabhatod a jelszó hosszát, illetve a használt karaktertípusokat.</li>
	<li><strong>Kommentek</strong>: Ezt használni tudod jegyzetek, információ és szöveg tárolására a fájl mellett (ez nem lesz titkosítva). Például, csatolhatsz egy leírást a fájlról még mielőtt elküldöd valakinek. Amikor a személy akinek küldted a fájlt azt beilleszti a Picocryptbe, a leírásod meg fog jelenni neki.</li>
	<li><strong>Kulcsfájlok</strong>: A Picocrypt támogatja kulcsfájlok használatát, mint a hitelesítés egy másik formáját (vagy az egyetlen formáját). Nem csak több kulcsfájlt használhatsz, hanem igényelheted, hogy egy sikeres visszafejtés érdekében a kulcsfájlok egy helyes sorrendben legyenek megadva. Egy kifejezetten jó használati módja több kulcsfájlnak, egy közös kötet létrehozása, ahol minden személy rendelkezik egy kulcsfájllal és mindannyiuknak (a kulcsfájlokkal együtt) jelen kell lenniük ahhoz, hogy a közös kötetet vissza lehessen fejteni.</li>
	<li><strong>Paranoiás mód</strong>: Ennek a módnak a használata mind XChaCha20-szal mind Serpent-tel titkosítani forgja az adataidat lépcsőzetes módon, és az SHA3 hash funkciót használja az adatok hitelesítéséhez BLAKE2b helyett. Ez a legtitkosabb fájlok védéséhez ajánlott és az elérhető legcélszerűbb gyakorlati biztonságot nyújtja. Ahhoz, hogy egy hekker feltörje a titkosított adataidat, mind az XChaCha20 ciphert, mind Serpent ciphert fel kell törnie, feltéve, hogy egy jó jelszót választottál. Biztosan ki lehet jelenteni, hogy ebben a módban a fájljaidat lehetetlen feltörni.</li>
	<li><strong>Reed-Solomon</strong>: Ez a sajátosság nagyon hasznos, ha fontos adatokat tervezel arhíválni egy felhő szolgáltatón vagy külső médiumon hosszú időre. Kiválasztva Picocrypt a Reed-Solomon hibakorrekciós kódot fogja használni, 8 byteot hozzáadva minden 128 bytehoz, hogy megelőzze a fájlsérülést. Ez azt jelenti, hogy akár ~3%-a a fájlodnak sérülhet és a Picocrypt még mindig helyre tudja hozni a hibákat és vissza tudja fejteni a fájljaidat sérülésmentesen. Természetesen, ha a fájl erősen sérül, (pl. leejtetted a merevlemezedet), Picocrypt nem fogja tudni teljesen helyrehozni azt, de mindent megtesz majd, hogy visszaszerezzen amit csak tud. Érdemes megjegyezni, hogy ez az opció jelentősen le fogja lassítani a titkosítást és visszafejtést.</li>
	<li><strong>Kényszer visszafejtés</strong>: A Picocrypt autómatikusan ellenőrzi a fájl integritást visszafejtéskor. Ha a fájl módosítva lett, vagy sérült, Picocrypt autómatikusan törölni fogja a visszafejtés végeredményét a felhasználó biztonságának érdekében. Ha felül szeretnéd bírálni ezt a védelmet, jelöld be ezt  az opciót. Továbbá, ha ez az opció ki volt választva és a Reed-Solomon használva volt a titkosított könyvtáron, Picocrypt megpróbál majd visszanyerni a fájlból amennyit csak lehet visszafejtés közben.</li>
	<li><strong>Fájlok részekre osztása</strong>: Nincs kedved óriási fájlokkal vesződni? Semmi gond! A Picocrypttal fel tudod osztani a kiadott fájlt egyedi méretű darabokra, így nagyobb fájlok jobban kezelhetővé és könnyebben feltölthetővé válnak felhőszolgáltatókhoz. Egyszerűen válassz ki egy mértékegységet (KiB, MiB, GiB, or TiB) és írd be a kívánt chunk méretet ahhoz az egységhez. Hogy a darabokat visszafejtsd, szimplán húzd be valamelyiket a Picocryptbe, majd a darabok autómatikusan egyesítve lesznek visszafejtéskor.</li>
</ul>

# Biztonság
További információért arról, hogy a Picocrypt hogyan kezel kriptográfiát, lásd a <a href="Internals.md">Működés</a>-t technikai részletekért. Ha a projekt, vagy az én biztonságom miatt aggódsz, hadd biztosítsalak, hogy ez a repository nem lesz eltérítve, vagy backdoorolva. Teljes lemeztitkosításon kívül minden hordozható eszközömön, be van kapcsolva a 2FA (TOTP) kétfrakkos azonosítás az összes fiókon ami a Picocrypthez köthető (GitHub, Google, Reddit, Ubuntu One/Snapcraft, Discord, stb.). További megerősítésként, Picocrypt az én izolált dependency forkjaimat használja és csak akkor fetchelek upstream ha megnéztem a változásokat és úgy hiszem, hogy nincsenek biztonsági problémák. Ez azt jelenti, hogy ha egy dependency-t meghekkelnek, vagy törli a szerző, Picocrypt az én forkomat fogja használni és teljesen érintetlen marad. A Picocryptet magabiztosan hasznáhatod.

# Közösség
Itt van pár hely ahol naprakész maradhatsz a Picocrypttel és részt tudsz venni a közösségben:
<ul>
	<li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
	<li><a href="https://discord.gg/8QM4A2caxH">Discord</a></li>
</ul>

Nagyon tudom ajánlani a Picocrypt subredditjéhez való csatlakozást, mert minden update és szavazás oda lesz kiposztolva. Tartsd észben, hogy csak ezekben a közösségi hálózatokban bízz, és vigyázz a hekkerekkel akik lehet megpróbálnak majd megszemélyesíteni. Én sosem fogok a jelszavad iránt érdeklődni és bárki aki igen, az nem én vagyok. Én sosem mondom majd, hogy tölts le egy fájlt egy gyanús linkről és bárki aki igen, az nem én vagyok.

# Ábrándozók
Mi a helyzet a Picocrypttel? Hogy megtudd nézz körül lejjebb.
![Stargazers Over Time](https://starchart.cc/HACKERALERT/Picocrypt.svg)

# Adományok	
Ha hasznosnak találod a Picocryptet, fontold meg, hogy támogatsz <a href="https://paypal.me/evanyiwensu">PayPalon</a>. Ezt a szoftver teljesen díjmentesen nyújtom és szeretném, ha lenne pár támogatóm aki továbbra is motiválná a munkámat a Picocrypten.

# Köszönetek
Egy köszönet a szívem mélyéről azoknak akik jelentős hozzájárulásukat tették Open Collective-en:
<ul>
	<li>YellowNight ($818)</li>
	<li>evelian ($50)</li>
	<li>jp26 ($50)</li>
	<li>guest-116103ad ($50)</li>
	<li>Tybbs ($10)</li>
	<li>N. Chin ($10)</li>
	<li>Manjot ($10)</li>
	<li>Phil P. ($10)</li>
	<li>donor39 (backer)</li>
	<li>Pokabu (backer)</li>
</ul>

Ti vagytok azok, akik továbbra is inspirálnak arra, hogy dolgozzak a Picocrypten és miattatok tudom továbbra is ingyenesesen nyújtani azt mindenkinek!

Szintén nagy köszönet a következő 5 embernek akik az első adományozói és támogatói voltak a Picocryptnek:
<ul>
	<li>W.Graham</li>
	<li>N. Chin</li>
	<li>Manjot</li>
	<li>Phil P.</li>
	<li>E. Zahard</li>
</ul>

Úgyszintén nagy köszönet ezeknek az embereknek akik segítettek a Picocrypt lefordításában és jobban hozzáférhetőbbé tették azt a világ számára:
<ul>
	<li>@umitseyhan75 Török</li>
	<li>@digitalblossom és @Pokabu26 Német</li>
	<li>@zeeaall Brazili Portugál</li>
	<li>@kurpau Litván</li>
	<li>u/francirc Spanyol</li>
	<li>yn Orosz</li>
	<li>@Etim-Orb Magyar</li>
	<li>@Minibus93 Olasz</li>
	<li>Michel Francia</li>
	<li>@victorhck Spanyol</li>
	<li>@ungespurv Fényesít</li>
</ul>

Végül köszönet ezeknek az embereknek/szervezeteknek amiért segítettek mikor szükségem volt arra:
<ul>
	<li>[ REDACTED ] amiért segített AppImage-et létrehozni a Picocrypthez</li>
	<li>Fuderal Discordon amiért segített egy Discord szerver létrehozásában</li>
	<li>u/greenreddits az állandó visszajelzésért és támogatásért</li>
	<li>u/Tall_Escape mert segített tesztelni Picocryptet</li>
	<li>u/NSABackdoors a rengeteg tesztelésért</li>
	<li>@samuel-lucas6 visszajelzésért, tanácsokért és támogatásért</li>
	<li><a href="https://privacytools.io">PrivacyToolsIO</a> Picocrypt megjelenítéséért</li>
	<li><a href="https://privacyguides.org">PrivacyGuides</a> Picocrypt megjelenítéséért</li>
</ul>
