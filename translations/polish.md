<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p> 

Picocrypt to bardzo małe (stąd <i>Pico</i>), proste, lecz bardzo bezpieczne narzędzie szyfrujące pozwalające chronić Twoje pliki. Został zaprojektowany jako <i>sprawdzone</i> narzędzie do szyfrowania, z naciskiem na bezpieczeństwo, prostotę i niezawodność. Picocrypt wykorzystuje bezpieczny szyfr XChaCha20 i funkcję wyprowadzania klucza Argon2id, aby zapewnić wysoki poziom bezpieczeństwa, nawet w przypadku trzyliterowych agencji, takich jak NSA. Został zaprojektowany z myślą o maksymalnym bezpieczeństwie, bez żadnych kompromisów pod względem bezpieczeństwa i jest zbudowany ze standardowych modułów x/crypto firmy Go. <strong>Twoja prywatność i bezpieczeństwo są zagrożone. Odzyskaj je bez obaw, chroniąc swoje pliki za pomocą Picocrypt.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="318" alt="Picocrypt"></p>

# Finansowanie
Przekaż darowiznę na rzecz Picocrypt na <a href="https://opencollective.com/picocrypt">Open Collective</a> (kryptowaluty są akceptowane), aby zebrać pieniądze na audyt bezpieczeństwa od Cure53. Ponieważ jest to projekt, nad którym spędzam wiele godzin i na którym nie zarabiam, sam nie mogę zapłacić za audyt. <i>Picocrypt potrzebuje wsparcia ze strony swojej społeczności.</i>

# Pobieranie
**Ważne**: Pod nazwą „Picocrypt” istnieje wiele podmiotów w Sieci. Na przykład istnieje stare narzędzie do szyfrowania o nazwie PicoCrypt, które używa złamanego szyfru. Istnieje również projekt badawczy finansowany przez ERBN o nazwie PICOCRYPT. Istnieją nawet domeny związane z Picocrypt, których nigdy nie rejestrowałem. Proszę nie mylić żadnego z tych niepowiązanych projektów z Picocrypt (ten projekt). Upewnij się, że pobierasz Picocrypt tylko z tego repozytorium, aby upewnić się, że otrzymasz autentyczny i wolny od backdoorów Picocrypt. Udostępniając Picocrypt z innymi, pamiętaj, aby połączyć się z tym repozytorium, aby uniknąć nieporozumień.

## Windows
Picocrypt dla Windows jest tak prosty, jak to tylko możliwe. Aby pobrać najnowszy, samodzielny i przenośny plik wykonywalny dla systemu Windows, kliknij <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.exe">tutaj</a>. Jeśli program Windows Defender lub Twój program antywirusowy oznaczy Picocrypt jako wirusa, zrób nam wszystkim przysługę i oznacz go jako fałszywy alarm dla dobra wszystkich.

Jeśli powyższy plik wykonywalny nie działa, prawdopodobnie oznacza to, że twój system nie obsługuje OpenGL. W tej sytuacji udostępniłem alternatywną implementację, która będzie działać w dowolnym systemie Windows, którą możesz pobrać <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt-NoGL.exe">tutaj</a>.

## macOS
Picocrypt dla macOS jest również bardzo prosty. Pobierz Picocrypt <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.app.zip">tutaj</a>, rozpakuj plik zip i uruchom Picocrypt, który jest w środku. Jeśli nie możesz otworzyć Picocrypt, ponieważ nie pochodzi od zweryfikowanego twórcy, kliknij prawym przyciskiem myszy Picocrypt i naciśnij „Otwórz”. Jeśli nadal otrzymujesz ostrzeżenie, kliknij prawym przyciskiem myszy Picocrypt i ponownie naciśnij „Otwórz”, a powinieneś być w stanie uruchomić Picocrypt.

## Linux
Istnieje wiele sposobów korzystania z Picocrypt w systemie Linux. Zalecanym sposobem jest zainstalowanie Picocrypt z .deb <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.deb">tutaj</a> (Debian 11+ i Ubuntu 20+). Jeśli plik .deb nie odpowiada Twoim potrzebom lub nie używasz dystrybucji opartej na Debianie, możesz użyć AppImage z <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.AppImage">tutaj</a>. Jeśli żadna z powyższych opcji nie działa, możesz zainstalować Picocrypt ze Snapcrafta, który powinien działać we wszystkich dystrybucjach. Instrukcje dotyczące Snapcrafta znajdziesz <a href="https://snapcraft.io/picocrypt">tutaj</a>.

## Paranoid Packs
Paranoid Pack to skompresowane archiwum, które zawiera pliki wykonywalne dla każdej wersji Picocrypt, która kiedykolwiek została wydana dla systemów Windows, macOS i Linux. Tak długo, jak przechowujesz je w miejscu, do którego masz dostęp, będziesz mógł go otworzyć i używać dowolnej wersji Picocrypt na wypadek, gdyby to repozytorium w tajemniczy sposób zniknęło lub cały Internet spłonął. Pomyśl o tym jako o skarbcu nasion dla Picocrypt. Dopóki jedna osoba ma Paranoid Packs w zasięgu ręki, może podzielić się nim z resztą świata i zachować funkcjonalność Picocrypt w przypadku katastrofalnych wydarzeń, takich jak nagłe zamknięcie GitHub lub przechwycenie mnie przez NSA (na wszelki wypadek, wiesz?). Najlepszym sposobem, aby Picocrypt był dostępny przez wiele dziesięcioleci, jest trzymanie Paranoid Packs w bezpiecznym miejscu. Więc jeśli martwisz się, że nie będziesz mógł uzyskać dostępu do Picocrypt w przyszłości, cóż, oto twoje rozwiązanie. Po prostu przejdź do zakładki Releases i zdobądź kopię.

# Dlaczego Picocrypt?
Dlaczego powinieneś używać Picocrypt zamiast BitLocker, NierdLocker, VeraCrypt, AxCrypt lub 7-Zip? Oto kilka powodów, dla których warto wybrać Picocrypt:
<ul>
	<li>W przeciwieństwie do NordLocker, BitLocker, AxCrypt i większości dostawców pamięci masowej w chmurze, Picocrypt i jego zależności są całkowicie otwarte i podlegają audytowi. Możesz sam sprawdzić, czy nie ma żadnych tylnych drzwi ani wad.</li>
	<li>Picocrypt jest <i>mały</i>. Podczas gdy NordLocker ma ponad 50 MB, a VeraCrypt ponad 20 MB, Picocrypt ma zaledwie 2 MB, mniej więcej rozmiar zdjęcia o średniej rozdzielczości. A to nie wszystko – Picocrypt jest przenośny (nie musi być zainstalowany) i nie wymaga uprawnień administratora/root'a.</li>
	<li>Picocrypt jest łatwiejszy i bardziej produktywny w użyciu niż VeraCrypt. Aby zaszyfrować pliki za pomocą VeraCrypt, musisz poświęcić co najmniej pięć minut na skonfigurowanie woluminu. Dzięki prostemu interfejsowi użytkownika Picocrypt wszystko, co musisz zrobić, to przeciągnąć i upuścić pliki, wprowadzić hasło i nacisnąć Start. Wszystkie złożone procedury są obsługiwane wewnętrznie przez Picocrypt. Kto powiedział, że bezpieczne szyfrowanie nie może być proste?</li>
	<li>Picocrypt został zaprojektowany z myślą o bezpieczeństwie. 7-Zip to narzędzie do archiwizacji, a nie narzędzie do szyfrowania, więc nie koncentruje się na bezpieczeństwie. Picocrypt jest jednak zbudowany z bezpieczeństwem jako priorytetem numer jeden. Każda część Picocrypt istnieje z jakiegoś powodu i wszystko, co może negatywnie wpłynąć na bezpieczeństwo Picocrypt, jest usuwane. Picocrypt używa kryptografii, której możesz zaufać.</li>
	<li>Picocrypt nie tylko chroni dane, ale także uwierzytelnia je, uniemożliwiając hakerom złośliwą modyfikację poufnych danych. Jest to przydatne, gdy wysyłasz zaszyfrowane pliki przez niezabezpieczony kanał i chcesz mieć pewność, że dotrze on nietknięty.</li>
	<li>Picocrypt aktywnie chroni zaszyfrowane dane nagłówka przed uszkodzeniem, dodając kodowanie korekcyjne Reed'a-Solomon'a, więc jeśli część nagłówka wolumenu (zawierająca ważne składniki kryptograficzne) ulegnie uszkodzeniu (np. gnicie bitów dysku twardego), Picocrypt może nadal odzyskać nagłówek i odszyfrować dane z wysokim wskaźnikiem sukcesu. Picocrypt może również zakodować cały wolumen za pomocą Reed'a-Solomon'a, aby zapobiec uszkodzeniu ważnych plików.</li>
</ul>


# Porównanie
Oto jak Picocrypt wypada w porównaniu z innymi popularnymi narzędziami do szyfrowania.

|                   | Picocrypt     | VeraCrypt    | 7-Zip (GUI)  | BitLocker   | Cryptomator  | NordLocker      | AxCrypt        |
| --------------    | ------------- | ------------ | ------------ | ------------| ------------ | --------------  | ---------------|
| Darmowy           |✅ Tak         |✅ Tak         |✅ Tak        |🟧 Częściowo  |✅ Tak        |🟧 Częściowo     |🟧 Częściowo    |
| Open Source       |✅ GPLv3       |✅ Multi       |✅ LGPL       |❌ Nie        |✅ GPLv3      |❌ Nie           |❌ Nie          |
| Międzyplatformowy |✅ Tak         |✅ Tak         |❌ Nie        |❌ Nie        |✅ Tak        |❌ Nie           |❌ Nie          |
| Rozmiar           |✅ 2MB         |❌ 20MB        |✅ 2MB        |✅ Zawarty    |❌ 50MB       |❌ 60MB          |🟧 8MB          |
| Przenośny         |✅ Tak         |✅ Tak         |❌ Nie        |✅ Tak        |❌ Nie        |❌ Nie           |✅ Tak          |
| Uprawnienia       |✅ Żadne       |❌ Admin       |❌ Admin      |❌ Admin      |❌ Admin      |❌ Admin         |❌ Admin        |
| Trudność          |✅ Łatwe       |❌ Trudny      |✅ Łatwe      |🟧 Średni     |🟧 Średni     |🟧 Średni        |✅ Łatwe        |
| Szyfr             |✅ XChaCha20   |✅ AES-256    	|✅ AES-256    |🟧 AES-128    |✅ AES-256    |✅ AES-256       |🟧 AES-128      |
| Klucz      				|✅ Argon2      |🆗 PBKDF2      |❌ SHA256     |❓ Brak info  |✅ Scrypt     |✅ Argon2        |🆗 PBKDF2       |
| Integralność      |✅ Zawsze      |❌ Nie         |❌ Nie        |❓ Brak info  |✅ Zawsze     |✅ Zawsze        |✅ Zawsze       |
| Reed-Solomon      |✅ Tak         |❌ Nie         |❌ Nie        |❌ Nie        |❌ Nie        |❌ Nie           |❌ Nie          |
| Kompresja         |✅ Tak         |❌ Nie         |✅ Tak        |✅ Tak        |❌ Nie        |❌ Nie           |✅ Tak          |
| Telemetria        |✅ Brak        |✅ Brak        |✅ Brak       |❓ Brak info  |✅ Brak       |❌ Analityka     |❌ Konta        |
| Po audycie        |🟧 Planned     |✅ Tak         |❌ Nie        |❓ Brak info  |✅ Tak        |❓ Brak info     |❌ Nie          |

# Funkcjonalności
Picocrypt to bardzo proste narzędzie, a większość użytkowników intuicyjnie zrozumie, jak z niego korzystać w ciągu kilku sekund. Na podstawowym poziomie wystarczy upuścić pliki, wprowadzić hasło i nacisnąć Start, aby zaszyfrować pliki. Całkiem proste, prawda?

Będąc prostym, Picocrypt stara się również być potężnym w rękach doświadczonych użytkowników. W związku z tym istnieje kilka dodatkowych opcji, których możesz użyć, aby dopasować je do swoich potrzeb.
<ul>
	<li><strong>Generator haseł</strong>: Picocrypt zapewnia generator bezpiecznych haseł, którego można używać do tworzenia kryptograficznie bezpiecznych haseł. Możesz dostosować długość hasła, a także typy znaków, które należy uwzględnić.</li>
	<li><strong>Komentarze</strong>: służy do przechowywania wiadomości, informacji i tekstu wraz z plikiem (nie będzie zaszyfrowany). Na przykład możesz umieścić opis szyfrowanego pliku przed wysłaniem go do kogoś. Gdy osoba, do której go wysłałeś, upuszcza plik do Picocrypt, Twój opis zostanie wyświetlony tej osobie.</li>
	<li><strong>Pliki kluczy</strong>: Picocrypt obsługuje użycie plików kluczy jako dodatkowej formy uwierzytelniania (lub jedynej formy uwierzytelniania). Nie tylko możesz używać wielu plików kluczy, ale możesz również wymagać prawidłowej kolejności plików kluczy, aby nastąpiło pomyślne odszyfrowanie. Szczególnie dobrym przypadkiem użycia wielu plików kluczy jest utworzenie współdzielonego woluminu, w którym każda osoba przechowuje plik klucza, a wszystkie z nich (i ich pliki kluczy) muszą być obecne, aby odszyfrować współdzielony wolumin.</li>
	<li><strong>Tryb Paranoid</strong>: użycie tego trybu spowoduje zaszyfrowanie danych za pomocą XChaCha20 i Serpent w sposób kaskadowy oraz użycie HMAC-SHA3 do uwierzytelniania danych zamiast BLAKE2b. Jest to zalecane do ochrony ściśle tajnych plików i zapewnia najwyższy możliwy do osiągnięcia poziom praktycznego bezpieczeństwa. Aby haker mógł złamać zaszyfrowane dane, zarówno szyfr XChaCha20, jak i szyfr Serpent muszą zostać złamane, zakładając, że wybrałeś dobre hasło. Można śmiało powiedzieć, że w tym trybie Twoje pliki są niemożliwe do złamania.</li>
	<li><strong>Reed-Solomon</strong>: Ta funkcja jest bardzo przydatna, jeśli planujesz archiwizować ważne dane u dostawcy chmury lub nośnika zewnętrznego przez długi czas. Jeśli zaznaczone, Picocrypt użyje kodu korekcji błędów Reed-Solomon, aby dodać 8 dodatkowych bajtów na każde 128 bajtów, aby zapobiec uszkodzeniu plików. Oznacza to, że do ~3% twojego pliku może ulec uszkodzeniu, a Picocrypt nadal będzie w stanie naprawić błędy i odszyfrować pliki bez korupcji. Oczywiście, jeśli Twój plik ulegnie bardzo poważnemu uszkodzeniu (np. Upuściłeś dysk twardy), Picocrypt nie będzie w stanie w pełni odzyskać Twoich plików, ale postara się odzyskać to, co może. Pamiętaj, że ta opcja może spowolnić szybkość szyfrowania i deszyfrowania.</li>
	<li><strong>Wymuś odszyfrowanie</strong>: Picocrypt automatycznie sprawdza integralność pliku po odszyfrowaniu. Jeśli plik został zmodyfikowany lub uszkodzony, Picocrypt automatycznie usunie dane wyjściowe dla bezpieczeństwa użytkownika. Jeśli chcesz obejść te zabezpieczenia, zaznacz tę opcję. Ponadto, jeśli ta opcja jest zaznaczona, a funkcja Reed-Solomon została użyta na zaszyfrowanym woluminie, Picocrypt spróbuje odzyskać jak najwięcej pliku podczas odszyfrowywania.</li>
	<li><strong>Podziel pliki na części</strong>: nie masz ochoty zajmować się gigantycznymi plikami? Nie martw się! Dzięki Picocrypt możesz podzielić plik wyjściowy na fragmenty o niestandardowych rozmiarach, dzięki czemu duże pliki będą łatwiejsze w zarządzaniu i łatwiejsze do przesyłania do dostawców chmury. Po prostu wybierz jednostkę (KiB, MiB, GiB lub TiB) i wprowadź żądany rozmiar kawałka dla tej jednostki. Aby odszyfrować fragmenty, po prostu przeciągnij jeden z nich do Picocrypt, a fragmenty zostaną automatycznie ponownie połączone podczas odszyfrowywania.</li>
</ul>

# Bezpieczeństwo
Aby uzyskać więcej informacji o tym, jak Picocrypt obsługuje kryptografię, zobacz <a href="Internals.md">Internals</a>, aby uzyskać szczegółowe informacje techniczne. Jeśli martwisz się o bezpieczeństwo moje lub tego projektu, zapewniam Cię, że to repozytorium nie zostanie przejęte ani przejęte przez użyciu backdoor'a. Mam włączone 2FA (TOTP) na wszystkich kontach z powiązaniem z Picocrypt (GitHub, Google, Reddit, Ubuntu One/Snapcraft, Discord itp.), oprócz pełnego szyfrowania dysku na wszystkich moich urządzeniach przenośnych. W celu dalszego umocnienia Picocrypt używa moich izolowanych widełek zależności i ściągam pod prąd tylko wtedy, gdy przyjrzałem się zmianom i wierzę, że nie ma żadnych problemów z bezpieczeństwem. Oznacza to, że jeśli zależność zostanie zhakowana lub usunięta przez autora, Picocrypt użyje mojego forka i pozostanie całkowicie nienaruszony. Możesz czuć się pewnie, korzystając z Picocrypt.

## Podpisy
Dla paranoików-Picocrypt jest podpisany za pomocą PGP. Fingerprint i klucz publiczny są wymienione poniżej.

<pre>B342A744BDEEA57B6A583E33A247E73798946F55</pre>
<pre>-----BEGIN PGP PUBLIC KEY BLOCK-----

mDMEYoGUHxYJKwYBBAHaRw8BAQdAvmQA+pdbDB/ynJxHhNDpz6Sb5tgkNuuNJIvw
HYwZtqi0CVBpY29jcnlwdIiTBBMWCgA7FiEEs0KnRL3upXtqWD4zokfnN5iUb1UF
AmKBlB8CGwMFCwkIBwICIgIGFQoJCAsCBBYCAwECHgcCF4AACgkQokfnN5iUb1UZ
RgEA8jbIsdqCr21DWxcqW/eLlbxRkuA8kflVYvWWUxtVqsUA/jQPSDpvA8rakvaL
PIbXjQvrAMkEVIc0HbCzLxr1k3sH
=YFwz
-----END PGP PUBLIC KEY BLOCK-----</pre>

# Społeczność
Oto kilka miejsc, w których możesz być na bieżąco z Picocrypt i się zaangażować:
<ul>
	<li><a href="https://discord.gg/8QM4A2caxH">Discord</a></li>
	<li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
</ul>

Gorąco polecam dołączyć do serwera Discord, ponieważ odbywa tam się większość komunikacji. Pamiętaj, aby ufać tylko tym sieciom społecznościowym i uważać na hakerów, którzy mogą próbować podszyć się pode mnie. Nigdy nie poproszę Cię o hasło, a każdy, kto to zrobi, to nie ja. Nigdy nie powiem ci, żebyś ściągnął plik z podejrzanego linku, a ten, kto to zrobi, nie jest mną.

# Statystyki
Jak sobie radzi Picocrypt? Spójrz poniżej, aby się dowiedzieć.
![Dane gwiazdki w czasie](https://starchart.cc/HACKERALERT/Picocrypt.svg)

# Dotacje
Jeśli uważasz, że Picocrypt jest przydatny, rozważ przekazanie mi napiwku <a href="https://paypal.me/evanyiwensu">PayPal</a>. Udostępniam to oprogramowanie całkowicie za darmo i chciałbym mieć wsparcie, które zmotywuje mnie do dalszej pracy nad Picocrypt. Obecnie jednak ważniejsze jest finansowanie audytu, więc jeśli chciałbyś przekazać darowiznę na Open Collective na finansowanie audytu, w przeciwieństwie do wspierania mnie, co w tej chwili jest mniejszym priorytetem.

# FAQ

**Czy Picocrypt akceptuje nowe funkcjonalności?**

Nie, Picocrypt jest uważany za kompletny i nie otrzyma żadnych nowych funkcji. W przeciwieństwie do innych narzędzi, które próbują stale dodawać nowe funkcje (co wprowadza nowe błędy i luki w zabezpieczeniach), Picocrypt skupia się tylko na kilku podstawowych funkcjach, ale każdą z nich wykonuje wyjątkowo dobrze. Zapamiętaj ideologię Picocrypt: mały, prosty i bezpieczny.

**Czy system Android/iOS będzie obsługiwany?**

Nie, nie planuję obsługi Androida ani iOS, ponieważ różnią się one bardzo od tradycyjnych systemów operacyjnych na komputery stacjonarne i wymagają różnych łańcuchów narzędzi do tworzenia aplikacji. Jednak ze względu na charakter oprogramowania typu open source możliwe jest, że w przyszłości pojawi się stworzona przez społeczność wersja Picocrypt na Androida lub iOS.

**Dlaczego Picocrypt nie jest często aktualizowany?**

Ludzie uważają, że oprogramowanie musi być stale aktualizowane, aby było aktualne i bezpieczne. Chociaż może to być prawdą w przypadku wielu programów, których używamy dzisiaj, nie dotyczy to Picocrypt. Picocrypt to „dobre oprogramowanie”, a dobre oprogramowanie nie wymaga ciągłych aktualizacji, aby zachować aktualność i bezpieczeństwo. Dobre oprogramowanie zawsze będzie dobrym oprogramowaniem.

# Podziękowania
Z całego serca dziękuję ludziom z Open Collective, którzy wnieśli znaczący wkład:
<ul>
	<li>YellowNight ($818)</li>
	<li>evelian ($50)</li>
	<li>jp26 ($50)</li>
	<li>guest-116103ad ($50)</li>
	<li>Markus ($15)</li>
	<li>Tybbs ($10)</li>
	<li>N. Chin ($10)</li>
	<li>Manjot ($10)</li>
	<li>Phil P. ($10)</li>
	<li>doNier39 (backer)</li>
	<li>Pokabu (backer)</li>
</ul>

Jesteście ludźmi, którzy inspirują mnie do pracy nad Picocrypt i dzięki Wam udostępniany jest wszystkim bezpłatnie!

Ogromne podziękowania należą się również poniższej liście pięciu osób, które jako pierwsze przekazały i wsparły Picocrypt:
<ul>
	<li>W.Graham</li>
	<li>N. Chin</li>
	<li>Manjot</li>
	<li>Phil P.</li>
	<li>E. Zahard</li>
</ul>

Również wielkie podziękowania dla tych osób, które pomogły przetłumaczyć Picocrypt i uczynić go bardziej dostępnym dla świata:
<ul>
	<li>@umitseyhan75 for Turkish</li>
	<li>@digitalblossom & @Pokabu26 for German</li>
	<li>@zeeaall for Brazilian Portuguese</li>
	<li>@kurpau for Lithuanian</li>
	<li>u/francirc for Spanish</li>
	<li>yn for Russian</li>
	<li>@Etim-Orb for Hungarian</li>
	<li>@Minibus93 for Italian</li>
	<li>Michel for French</li>
	<li>@victorhck for Spanish</li>
	<li>@MasterKia for Persian</li>
	<li>@ungespurv for Polish</li>
</ul>

Na koniec dziękuję tym osobom/organizacjom za pomoc:
<ul>
	<li>[ZMIENIONO] za pomoc w tworzeniu AppImage dla Picocrypt</li>
	<li>u/Upstairs-Fishing867 za pomoc w testowaniu sygnatur PGP</li>
	<li>Fuderal na Discord za pomoc w konfiguracji serwera Discord</li>
	<li>u/greenreddits za stałą informację zwrotną i wsparcie</li>
	<li>u/Tall_Escape za pomoc w testowaniu Picocrypt</li>
	<li>u/NSABackdoors za wykonywanie wielu testów</li>
	<li>@samuel-lucas6 za opinię, sugestie i wsparcie</li>
	<li><a href="https://privacytools.io">PrivacyToolsIO</a> za ujęcie Picocrypt na liście</li>
	<li><a href="https://privacyguides.org">PrivacyGuides</a> za ujęcie Picocrypt na liście</li>
</ul>
