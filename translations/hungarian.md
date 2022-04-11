<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p> 

A Picocrypt egy nagyon apró (<i>Pico</i>), nagyon egyszerű, de mégis nagyon biztonságos adattitkosító eszköz amit használhatsz a fájljaid megvédésére, checksum generálásra, és sok minden másra is. Egy <i>go-to</i> adattitkosító eszköznek lett kialakítva, ahol a fókusz a biztonságon, egyszerűségen, és megbízhatóságon van. Picocrypt a biztonságos XChaCha20 ciphert és az Argon2 kulcs deriváció funkciót használja, hogy magas szintű biztonságot nyújtson, még három betűs ügynökségek ellen is, mint az NSA. A maximális biztonságra lett tervezve, így nem köt semmilyen kompromisszumot ilyen téren, valamint a Go standard x/crypto moduljaival lett kiépítve. <strong>A magánéleted és az adatbiztonságod támadás alatt áll. Védd a fájljaidat Picocrypttel és szerezd vissza ezeket magabiztossággal.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="384" alt="Picocrypt"></p>

# Finanszírozás
Kérlek támogasd a Picocryptet <a href="https://opencollective.com/picocrypt">Open Collective</a>-en (crypto-t is elfogadunk), hogy elég pénzt gyűjtsünk össze egy lehetséges auditra a Cure53-tól. Mivel ez egy projekt amibe én rengeteg órát fektetek bele és nem keresek vele semmi pénzt, saját magam nem tudok fizetni az auditért. <i>Picocrypt-nak szüksége van támogatásra a közösségétől.</i>

# Letöltések
<strong>Fontos:</strong> Van egy elavult és hasznavehetetlen abandonware az interneten, amit PicoCrypt-nak hívnak és ami legutóbb 2005-ben lett frissítve. PicoCrypt nem áll kapcsolatban semmilyen formában a Picocrypt-tel (ezzel a projekttel). Győződj meg róla, hogy csak ebből a repository-ból töltöd le a Picocryptet, hogy biztosan az eredeti és backdoor mentes Picocryptet szerzed be.

## Windows
A Picocrypt Windows-ra olyan egyszerű amennyire az csak lehet. Ahhoz, hogy letöltsd a legfrissebb, egyedülálló és hordozható futtatóprogramot Windowshoz, kattints <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.24/Picocrypt.exe">ide</a>. Ha a Windows Defender, vagy a vírusírtód vírusként jelöli meg a Picocryptet, kérlek vedd ki a részed azzal, hogy benyújtod, mint hibás találatot, így segítve mindannyiunknak.

## macOS
Picocrypt a macOS-en szintén egyszerű. Töltsd le a Picocrypt-et <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.24/Picocrypt.app.zip">itt</a>, csomagold ki a zip fájlt, és futtasd Picocryptet, ami benne van. Ha nem tudod megnyitni a Picocryptet, mert nem egy igazolt fejlesztőtől van, jobb-klikkelj a Picocryptre és válaszd ki a "Megnyitás" lehetőséget. Ha ezután is megkapod a figyelmeztetést, jobb-klikkelj a Picocryptre és kattints a "Megnyitás" lehetőségre ismét és ezután el kéne tudnod indítani a Picocryptet.

## Linux
Linuxhoz egy Snap elérhető. Feltételezve, hogy egy Debian alapú rendszeren vagy, egy egyszerű `apt install snapd` és `snap install picocrypt` elég lesz. Más disztrókhoz, mint a Fedora, részletes utasítások elérhetőek a https://snapcraft.io/picocrypt linken. A dependency-k komplexitása és statikus linkelés miatt, én nem készítek .deb és .rpb binary-kat, mivel megbízhatatlanok lennének és nem éri meg a kavarodást. A Snapcraft autómatikusan kezel minden dependency-t valamint futtatót és ez az ajánlott módja a Picocrypt futtatásának bármely fő Linux disztribúción. Ezen felül, a Snapcraft jobb biztonságot és kompartmentalizációt nyújt, mint a Flatpak-ek és AppImage-ek, ami fontos egy titkosító eszköznél mint a Picocrypt. Ha nem szeretnél a Canonical-lal foglalkozni, emlékeztetlek, hogy a forrásból buildelés mindig lehetséges.

# Miért a Picocrypt?
Miért használd a Picocryptet a BitLocker, NordLocker, VeraCrypt, AxCrypt, vagy 7-Zip helyett? Itt van pár ok, hogy miért válaszd a Picocryptet:
<ul>
	<li>Nem úgy mint a NordLocker, BitLocker, AxCrypt, és a legtöbb felhőtárhely szolgáltató, Picocrypt és annak dependency-jei teljesen nyílt forráskódúak és auditolhatók. Saját magad is ellenőrizni tudod, hogy nincs-e benne backdoor, vagy hiba.</li>
	<li>A Picocrypt <i>kicsi</i>. Míg a NordLocker 100MB fölött és VeraCrypt 30MB fölött van, addig a Picocrypt csupán csak 3MB, ami körülbelül egy magas felbontású kép mérete. És ez nem minden - A Picocrypt hordozható (nem kell, hogy telepítve legyen) és nem igényel adminisztrátor/root hozzáférést.</li>
	<li>A Picocryptet könnyeb és sokkal produktívabb használni mint a VeraCryptet. Ahhoz, hogy fájlokat titkosíts a VeraCrypt-tel, legalább öt percet el kell töltened egy kötet létrehozásával. A Picocrypt szimpla felhasználói felületével, csak annyit kell tenned, hogy ráhúzod a fájljaidat, beírsz egy jelszót és kiválasztod az Indítást. Az összes komplex folyamat a Picocrypt álltal és saját magán belül van kezelve. Ki mondta, hogy biztonságos titkosítás nem lehet egyszerű?</li>
	<li>A Picocrypt a biztonságra lett tervezve. A 7-Zip egy archíváló program és nem egy titkosító eszköz, ezért a fókusza nem a biztonságra helyeződik. A Picocrypt azonban biztonsággal lett tervezve, mint legfőbb prioritás. A Picocrypt minden egyes része okkal van ott ahol, és bármi ami kihatna Picocrypt biztonságára el van távolítva. A Picocrypt olyan kriptográfiát alkalmaz amiben megbízhatsz.</li>
	<li>A Picocrypt hitelesíti az adatokat azok védelmén kívül, így meggátolja, hogy hekkerek érzékeny információn rosszindulatú módosításokat hajtsanak végre. Ez hasznos amikor titkosított fájlokat küldesz egy nem biztonságos csatornán keresztül és meg akarsz róla győződni, hogy azok érintetlenül érkeznek meg.</li>
	<li>A Picocrypt aktívan védi a titkosított header adatokat az adatsérüléstől Reed-Solomon parity byteok hozzáadásával, így ha a kötet header-jének egy része (ami fontos kriptográfiai komponenseket tartalmaz) sérül (pl. merevlemez bit rot), attól a Picocrypt még mindig nagy valószínűséggel helyre tudja állítani a headert és vissza tudja fejteni az adatokat. Picocrypt ezen kívül encode-olni tudja a teljes kötetet Reed-Solomonnal, hogy megelőzzön bármilyen adatsérülést ami a fontos fájljaidat érintené.</li>
</ul>

Még mindig nem sikerült meggyőzzelek? Itt olvashatsz még több okot, hogy hogyan tűnik ki Picocrypt mások közül.

# Sajátosságok
A Picocrypt egy nagyon egyszerű eszköz és a legtöbb felhasználó pár másodpercen belül ösztönösen rá fog jönni, hogy hogyan kell használni. Alap szinten egyszerűen a fájljaid behúzása, egy jelszó beírása és az Indítás kiválasztása minden amire szükség van a fájljaid titkosításához. Elég egyszerű, ugye?

Bár egyszerű, Picocrypt törekszik arra is, hogy rendkívül hatékony legyen tapasztalt és haladó felhasználók kezében. Így van pár további opció amit használhatsz a szükségleteidnek megfelelően.
<ul>
	<li><strong>Jelszó generátor</strong>: Picocrypt nyújt egy biztonságos jelszó generátort amit kriptográfiailag biztonságos jelszavak létrehozására használhatsz. Tesztreszabhatod a jelszó hosszát, illetve a használt karaktertípusokat.</li>
	<li><strong>Fájl metaadat</strong>: Ezt használni tudod jegyzetek, információ és szöveg tárolására a fájl mellett (ez nem lesz titkosítva). Például, csatolhatsz egy leírást a fájlról még mielőtt elküldöd valakinek. Amikor a személy akinek küldted a fájlt azt beilleszti a Picocryptbe, a leírásod meg fog jelenni neki.</li>
	<li><strong>Kulcsfájlok</strong>: A Picocrypt támogatja kulcsfájlok használatát, mint a hitelesítés egy másik formáját. Nem csak több kulcsfájlt használhatsz, hanem igényelheted, hogy egy sikeres visszafejtés érdekében a kulcsfájlok egy helyes sorrendben legyenek megadva. Egy kifejezetten jó használati módja több kulcsfájlnak, egy közös kötet létrehozása, ahol minden személy rendelkezik egy kulcsfájllal és mindannyiuknak (a kulcsfájlokkal együtt) jelen kell lenniük ahhoz, hogy a közös kötetet vissza lehessen fejteni.</li>
	<li><strong>Paranoiás mód</strong>: Ennek a módnak a használata mind XChaCha20-szal mind Serpent-tel titkosítani forgja az adataidat egy lépcsőzetes formában, és az SHA3 hash funkciót használja az adatok hitelesítéséhez BLAKE2b helyett. Ez a legtitkosabb fájlok védéséhez ajánlott és az elérhető legcélszerűbb gyakorlati biztonságot nyújtja. Ahhoz, hogy egy hekker feltörje a titkosított adataidat, mind az XChaCha20 ciphert és a Serpent ciphert fel kell törnie, feltéve, hogy egy jó jelszót választottál.</li>
	<li><strong>Adatsérülés megelőzése Reed-Solomonnal</strong>: Ez a sajátosság nagyon hasznos, ha fontos adatokat tervezel arhíválni egy felhő szoláltatón vagy külső médiumon hosszú időre. Kiválasztva Picocrypt a Reed-Solomon hibakorrekciós kódot fogja használni, 8 byteot hozzáadva minden 128 bytehoz, hogy megelőzze a fájlsérülést. Ez azt jelenti, hogy akár ~3%-a a fájlodnak sérülhet és a Picocrypt még mindig helyre tudja hozni a hibákat és vissza tudja fejteni a fájljaidat sérülésmentesen. Természetesen, ha a fileod erősen sérül, (pl. leejtetted a merevlemezedet), Picocrypt nem fogja tudni teljesen helyrehozni a fájlt, de mindent megtesz, hogy visszaszerezzen amit csak tud. Érdemes megjegyezni, hogy ez az opció jelentősen le fogja lassítani a titkosítást és visszafejtést.</li>
	<li><strong>Visszafejtett adatok megtartása, akkor is, ha az sérült, vagy módosították</strong>: A Picocrypt autómatikusan ellenőrzi az integritást visszafejtéskor. Ha a fájl módosítva lett, vagy sérült, Picocrypt autómatikusan törölni fogja a visszafejtés végeredményét a felhasználó biztonságának érdekében. Ha meg akarod őrizni a sérült, vagy módosított adatokat visszafejtés után, jelöld be ezt  az opciót. Továbbá, ha ez az opció ki volt választva és a Reed-Solomon használva volt a titkosított fájlon, Picocrypt megpróbál majd visszanyerni a fájlból amennyit csak lehet visszafejtés közben.</li>
	<li><strong>Fájlok darabokra osztása</strong>: Nincs kedved óriási fájlokkal vesződni? Semmi gond! A Picocrypttal fel tudod osztani a kiadott fájlt egyedi méretű darabokra, így nagyobb fájlok jobban kezelhetővé és könnyebben feltölthetővé válnak felhőszolgáltatókhoz. Egyszerűen válassz ki egy mértékegységet (KiB, MiB, or GiB) és írd be a kívánt számot ahhoz az egységhez. Hogy a darabokat visszafejtsd, szimplán húzd be valamelyiket a Picocryptbe, majd a darabok autómatikusan egyesítve lesznek visszafejtéskor.</li>
</ul>

Ezeken a széleskörű opciókon kívül a titkosításhoz és visszafejtéshez, Picoccrypt nyújt egy checksum generátort is érzékeny fájlok sértetlenségének ellenőrzéséhez, ami támogat számos hash funkciót mint az MD5, BLAKE2, és SHA3.

# Biztonság
További információért arról, hogy a Picocrypt hogyan kezel kriptográfiát, lásd a <a href="Internals.md">Működés</a>-t technikai részletekért. Ha a projekt, vagy az én biztonságom miatt aggódsz, hadd biztosítsalak, hogy ez a repository nem lesz eltérítve, vagy backdoorolva. Teljes lemeztitkosításon kívül minden hordozható eszközömön, be van kapcsolva a 2FA (TOTP) kétfrakkos azonosítás az összes fiókon ami a Picocrypthez köthető (GitHub, Google, Reddit, Ubuntu One/Snapcraft, Discord, stb.). További megerősítésként, Picocrypt az én izolált dependency forkjaimat használja és csak akkor fetchelek upstream ha megnéztem a változásokat és úgy hiszem, hogy nincsenek biztonsági problémák. Ez azt jelenti, hogy ha egy dependency-t meghekkelnek, vagy törli a szerző, Picocrypt az én forkomat fogja használni és teljesen érintetlen marad. A Picocryptet magabiztosan hasznáhatod.

# Közösség
Itt van pár hely ahol naprakész maradhatsz a Picocrypttel és részt tudsz venni a közösségben:
<ul>
	<li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
	<li><a href="https://discord.gg/8QM4A2caxH">Discord</a></li>
</ul>
A Picocrypt subredditjéhez való csatlakozást nagyon ajánlani tudom, mert minden update és szavazás oda lesz kiposztolva. Tartsd észben, hogy csak ezekben a közösségi hálózatokban bízz, és vigyázz a hekkerekkel akik lehet megpróbálnak majd megszemélyesíteni. Én sosem fogok a jelszavad iránt érdeklődni és bárki aki igen, az nem én vagyok. Én sosem mondom majd, hogy tölts le egy fájlt egy gyanús linkről és bárki aki igen, az nem én vagyok.

# Ábrándozók
Mi a helyzet a Picocrypttel? Hogy megtudd nézz körül lejjebb.
[![Stargazers over time](https://starchart.cc/HACKERALERT/Picocrypt.svg)](https://starchart.cc/HACKERALERT/Picocrypt)

# Adományok	
Ha hasznosnak találod a Picocryptet, fontold meg, hogy támogatsz <a href="https://paypal.me/evanyiwensu">PayPalon</a>. Ezt a szoftver teljesen díjmentesen nyújtom és szeretném, ha lenne pár támogatóm aki továbbra is motiválná a munkámat a Picocrypten.

# Köszönetek
Egy köszönet a szívem mélyéről azoknak akik jelentős hozzájárulásukat tették Open Collective-en:
<ul>
	<li>jp26 ($50)</li>
	<li>guest-116103ad ($50)</li>
	<li>Tybbs ($10)</li>
	<li>N. Chin ($10)</li>
	<li>Manjot ($10)</li>
	<li>Phil P. ($10)</li>
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

Úgyszintén nagy köszönet ezeknek az embereknek akik segítettek a Picocrypt lefordításában és elérhetőbbé tették azt a világ számára:
<ul>
	<li>@umitseyhan75 Török</li>
	<li>@digitalblossom Német</li>
	<li>@zeeaall Brazil Portugál</li>
	<li>@kurpau Litván</li>
	<li>u/francirc Spanyol</li>
	<li>yn Orosz</li>
	<li>@Etim-Orb Magyar</li>
</ul>

Végül köszönet ezeknek az embereknek amiért segítettek amikor szükségem volt rá:
<ul>
	<li>Fuderal Discordon amiért segített egy Discord szerver létrehozásában</li>
	<li>u/greenreddits az állandó visszajelzésért és támogatásért</li>
	<li>u/Tall_Escape mert segített tesztelni Picocryptet</li>
	<li>u/NSABackdoors a rengeteg tesztelésért</li>
	<li>@samuel-lucas6 visszajelzésért, tanácsokért és támogatásért</li>
</ul>
