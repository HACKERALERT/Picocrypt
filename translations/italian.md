
<p  align="center"><img  align="center"  src="/images/logo.svg"  width="512"  alt="Picocrypt"></p>

Picocrypt è uno strumento di crittografia piccolissimo (ecco perché <i>Pico</i>) e semplicissimo, ma, allo stesso tempo, estremamente sicuro che puoi usare per proteggere i tuoi file. È progettato per essere lo strumento <i>go-to</i> quando si parla di crittografia, con particolare attenzione su sicurezza, semplicità d'uso e affidabilità. Picocrypt usa la cifratura sicura Xchacha20 e la funzione di derivazione della password Argon2id per garantire un elevato standard di sicurezza, anche da agenzie "a 3 lettere" come l'NSA. È progettato per garantire la massima sicurezza senza alcun compromesso ed è realizzato con i moduli standard x/crypto di Go. <strong>La tua sicurezza e la tua privacy sono sotto attacco. Difenditi in modo sicuro proteggendo i tuoi file con Picocrypt.</strong>

<p  align="center"><img  align="center"  src="/images/screenshot.png"  width="384"  alt="Picocrypt"></p>

# Donazioni

Fai una donazione a Picocrypt su <a  href="https://opencollective.com/picocrypt">Open Collective</a> (è possibile usare anche le criptovalute) per raccogliere fondi in modo da poter richiedere una revisione da parte di Cure53. Dato che, nonostante io spenda moltissimo tempo su questo progetto, non ci guadagno assolutamente nulla, non posso permettermi di pagare una revisione di tasca mia. <i>Picocrypt ha bisogno di aiuto dalla sua community.</i>

# Download

<strong>Importante:</strong> Nella rete si trova, cercando Picocrypt, anche un programma, ormai inutile e non più seguito, chiamato "PicoCrypt". PicoCrypt non è correlato in alcun modo a Picocrypt (questo progetto). Assicurati di scaricare Picocrypt da questa pagina per essere certo di ottenere una copia di Picocrypt autentica e priva di backdoor.

## Windows

Scaricare Picocrypt per Windows è semplicissimo. Scarica l'ultima versione, standalone e portable, degli eseguibili per Windows, cliccando <a  href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.24/Picocrypt.exe">qui</a>. Se Windows Defender o il tuo antivirus segnala Picocrypt come un virus, per piacere, per il bene di tutti, contribuisci al progetto segnalandolo come un falso positivo.

## macOS

Scaricare Picocrypt per macOS è altrettanto semplice. Scarica Picocrypt <a  href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.24/Picocrypt.app.zip">qui</a>, estrai l'archivio zip ed esegui il Picocrypt che troverai al suo interno. Se non riesci ad aprire Picocrypt perché non proviene da uno sviluppatore certificato, premi il tasto destro su Picocrypt e seleziona "Apri". Se nonostante tutto ricevi ancora un avviso, ripremi il tasto destro su Picocrypt e seleziona "Apri" ancora una volta. A questo punto dovresti poter eseguire Picocrypt.

## Linux

Uno Snap è disponibile per Linux. Ammesso che tu sia su un sistema basato su Debian, un semplice `apt install snapd` e `snap install picocrypt` sarà sufficiente. Per altre distro, come Fedora, sono disponibili istruzioni dettagliate su https://snapcraft.io/picocrypt. A causa della complessità delle dipendenze e del collegamento statico, non distribuirò binari .deb o .rpm standalone perché risulterebbero inaffidabili e non ne varrebbe la pena. Snapcraft gestisce tutte le dipendenze e le runtime automaticamente e rappresenta il metodo consigliato di eseguire Picocrypt sulle principali distro Linux. Inoltre, Snapcraft offre maggior sicurezza ed una containerizzazione migliore rispetto Flatpak ed AppImage, due cose che risultano importanti per uno strumento di crittografia come Picocrypt. Se preferisci non aver a che fare con Canonical, ricorda che puoi sempre compilare Picocrypt dai sorgenti.

## Confezione paranoia

La "Confezione paranoia" è un archivio compresso che contiene gli eseguibili di ogni versione distribuita di Picocrypt per Windows, macOS e Linux. Fino a che lo conservi in un posto a cui hai accesso, potrai aprirlo ed usare qualsiasi versione di Picocrypt nel caso in cui questa pagina svanisca misteriosamente o bruciasse l'intero Internet. Vedila come una cassaforte di seed di Picocrypt. Finché almeno una persona ha una "Confezione paranoia" a portata di mano, questa può condividerla con il resto del mondo facendo così rimanere Picocrypt funzionante anche in caso di eventi catastrofici come, ad esempio, la chiusura improvvisa di GitHub o il mio arresto da parte dell'NSA (non si sa mai, giusto?). Il miglior modo per far sì che Picocrypt sia accessibile per molte decadi è quello di conservare una "Confezione paranoia" in un posto sicuro. Quindi, se ti preoccupa non poter più avere accesso a Picocrypt nel futuro, bhe, ora hai la soluzione.

# Perché Picocrypt?

Perché dovresti usare Picocrypt invece di BitLocker, NordLocker, VeraCrypt, AxCrypt, o 7-Zip? Ecco alcuni motivi per scegliere Picocrypt:

<ul>

<li>A differenza di NordLocker, BitLocker, AxCrypt, e della maggior parte dei cloud, Picocrypt e le sue dipendenze sono completamente open-source e passibili di revisione. Puoi verificare con i tuoi stessi occhi che non ci sono backdoor o falle.</li>

<li>Picocrypt è <i>piccolo</i>. Mentre NordLocker supera i 100MB e VeraCrypt i 30MB, Picocrypt pesa solo 3MB, circa la dimensione di una foto ad alta risoluzione. E non è finita qui - Picocrypt è "portable" (non è necessario installarlo) e non richiede privilegi root/amministratore.</li>

<li>Picocrypt è più semplice e più immediato da usare di VeraCrypt. Per crittografare i file con VeraCrypt, devi spendere almeno 5 minuti per creare un volume. Con l'UI semplice di Picocrypt, tutto quello che devi fare è un "drag & drop" dei tuoi file, inserire una password e premere Start. Tutte le parti complicate sono gestite internamente da Picocrypt. Chi ha detto che una crittografia sicura non possa, allo stesso tempo, essere semplice?</li>

<li>Picocrypt è progettato con particolare attenzione alla sicurezza. 7-Zip è un gestore di archivi, non uno strumento di crittografia, quindi non si concentra sulla sicurezza. Picocrypt, invece, è stato realizzato avendo la sicurezza come priorità numero uno. Ogni parte di Picocrypt esiste per una ragione e tutto quello che potrebbe influenzare negativamente la sicurezza di Picocrypt viene rimosso. Picocrypt è progettato con una crittografia di cui ti puoi fidare.</li>

<li>Picocrypt autentica i dati oltre a proteggerli, impedendo agli hacker di modificare malevolmente i dati sensibili. Questo risulta particolarmente utile quando si inviano file criptati tramite un canale insicuro e si vuole essere certi che arrivino a destinazione intatti.</li>

<li>Picocrypt protegge attivamente i dati dell'intestazione crittografata dalla corruzione aggiungendo ulteriori bit di parità Reed-Solomon, quindi se parte dell'intestazione di un volume (che contiene importanti componenti crittografici) si corrompe (ad esempio, il bit del disco rigido si degrada), Picocrypt può ancora recuperare l'intestazione e decifrare i tuoi dati con un'alta percentuale di successo. Picocrypt può anche codificare l'intero volume tramite Reed-Solomon per prevenire qualsiasi corruzione dei tuoi file importanti.</li>

</ul>

Non sei ancora convinto? Qui sotto ci sono ulteriori motivi per cui Picocrypt si distingue dalla massa.

# Paragoni
Ecco Picocrypt paragonato ai software di crittografia più famosi: 

|                | Picocrypt      | VeraCrypt      | 7-Zip (GUI)    | NordLocker     | BitLocker      | AxCrypt        |
| -------------- | -------------- | -------------- | -------------- | -------------- | -------------- | -------------- |
| Gratis           |✅ Si         |✅ Si         |✅ Si         |🟧 In parte    |🟧 In parte   |🟧 In parte    |
| Open Source    |✅ GPLv3       |✅ Multi        |✅ LGPL        |❌ No           |❌ No          |❌ No           |
| Multi-piattaforma |✅ Si         |✅ Si          |❌ No          |❌ No           |❌ No          |❌ No           |
| Dimensione          |✅ 2MB         |❌ 20MB         |✅ 2MB         |❌ 60MB         |✅ Inclusa    |🟧 8MB          |
| Portatile       |✅ Si         |✅ Si          |❌ No          |❌ No           |✅ Si         |✅ Si          |
| Permessi   |✅ Nessuno        |❌ Admin        |❌ Admin       |❌ Admin        |❌ Admin       |❌ Admin        |
| Semplicità d'uso    |✅ Semplice        |❌ Difficile         |✅ Semplice        |🟧 Media       |🟧 Media      |✅ Semplice         |
| Derivazione della chiave |✅ Argon2      |🆗 PBKDF2       |❌ SHA256      |✅ Argon2       |❓ Sconosciuta     |🆗 PBKDF2       |
| Controllo integrità dei dati |✅ Sempre      |❌ No           |❌ No          |✅ Sempre       |❓ Sconosciuta     |✅ Sempre       |
| Reed-Solomon   |✅ Si         |❌ No           |❌ No          |❌ No           |❌ No          |❌ No           |
| Compressione    |✅ Si         |❌ No           |✅ Si         |❌ No           |✅ Si         |✅ Si          |
| Telemetria      |✅ Nessuna       |✅ Nessuna         |✅ Nessuna        |❌ Analytics    |❓ Sconosciuta     |❌ Account     |
| Revisionato        |🟧 Pianificata     |✅ Si          |❌ No          |❓ Sconosciuta      |❓ Sconosciuta     |❌ No           |

# Funzionalità

Picocrypt è uno strumento molto semplice, ed, in pochi secondi, la maggior parte degli utenti capirà intuitivamente come usarlo. Limitandoci ad un livello molto basico, è sufficiente rilasciare i file, inserire una password e premere Start per crittografare i file. Abbastanza semplice, no?

Pur essendo semplice, Picocrypt si sforza anche di essere efficace nelle mani di utenti esperti. Pertanto, ci sono alcune opzioni aggiuntive che è possibile utilizzare in base alle proprie esigenze.

<ul>

<li><strong>Generatore di password</strong>: Picocrypt fornisce un generatore di password sicure che può essere usato usare per creare password crittograficamente sicure. È possibile personalizzare la lunghezza della password, così come i tipi di caratteri da includere.</li>

<li><strong>File metadata</strong>: Usalo per conservare note, informazioni e testo insieme al file (queste note/informazioni non saranno criptate). Per esempio, puoi mettere una descrizione del file che stai criptando prima di inviarlo a qualcuno. Quando la persona a cui l'hai inviato inserisce il file in Picocrypt, potrà vedere la tua descrizione.</li>

<li><strong>File chiave</strong>: Picocrypt supporta l'uso di file chiave come forma aggiuntiva di autenticazione (o come unica forma di autenticazione). Non solo è possibile usare più file chiave, ma si può anche fare in modo che il corretto ordine dei file chiave sia richiesto, affinché la decrittazione abbia successo. Un modo d'uso particolarmente conveniente dei file chiave multipli è la creazione di un volume condiviso, dove ogni persona possiede un file chiave, e tutti loro (con i loro file chiave) devono essere presenti per decifrare il volume condiviso.

</li>

<li><strong>Modalità paranoia</strong>: Usando questa modalità i tuoi dati saranno criptati sia con XChaCha20 che con Serpent in modo a cascata e sarà usato HMAC-SHA3 per autenticare i dati invece di BLAKE2b. Questo è il metodo raccomandato per proteggere i file più importanti, top-secret, dato che fornisce il più alto livello di sicurezza pratica raggiungibile. Affinché un hacker possa decifrare i tuoi dati criptati, sia il cifrario XChaCha20 che il cifrario Serpent devono essere violati, assumendo che tu abbia scelto una buona password.</li>

<li><strong>Previeni la corruzione dei file usando Reed-Solomon</strong>: Questa funzionalità risulta molto utile se avete intenzione di archiviare dati importanti su un cloud o su un supporto esterno per un lungo periodo. Se selezionata, Picocrypt utilizzerà il codice di correzione degli errori Reed-Solomon per aggiungere 8 byte extra ogni 128 byte per prevenire la corruzione del file. Questo significa che fino al ~3% del tuo file può corrompersi e, nonostante ciò, Picocrypt sarà ancora in grado di correggere gli errori e decifrare i tuoi file senza corruzione. Naturalmente, se il tuo file si corrompe in modo estremo (per esempio se hai fatto cadere il disco rigido), Picocrypt non sarà in grado di recuperare completamente i tuoi file, ma farà comunque del suo meglio per recuperare ciò che può. Nota che questa opzione rallenterà considerevolmente sia la velocità del processo di crittografia che quello della decrittografia.</li>

<li><strong>Mantieni l'output decriptato anche se è corrotto o modificato</strong>: Picocrypt controlla automaticamente l'integrità al momento della decrittazione. Se il file è stato modificato o è corrotto, Picocrypt cancellerà automaticamente l'output per la sicurezza dell'utente. Se vuoi mantenere i dati corrotti o modificati dopo la decrittazione, allora dovrai selezionare questa opzione. Inoltre, se questa opzione è selezionata e la funzione Reed-Solomon è stata utilizzata sul file crittografato, Picocrypt tenterà di recuperare quanto più possibile del file durante la decrittazione.</li>

<li><strong>Dividi il file in pezzi</strong>: Non avete voglia di avere a che fare con file enormi? Non preoccupatevi! Con Picocrypt, puoi scegliere di dividere il tuo file di output in pezzi di dimensioni personalizzate, così i file di grandi dimensioni possono diventare più gestibili e più facili da caricare sui cloud. Basta scegliere un'unità (KiB, MiB, GiB o TiB) e inserire la dimensione desiderata. Per decriptare i pezzi, basta trascinarne uno in Picocrypt, e i pezzi saranno automaticamente ricombinati durante la decriptazione.</li>

</ul>

Oltre a tutte queste opzioni per la crittografia e la decrittografia, Picocrypt offre anche un generatore di checksum, per convalidare l'integrità dei file, che supporta numerose funzioni di hash come MD5, BLAKE2 e SHA3.

# Sicurezza

Per maggiori informazioni su come Picocrypt gestisce la crittografia, puoi trovare e leggere i dettagli tecnici su <a  href="Internals.md">Internals</a>. Se sei preoccupate per quanto riguardo la mia sicurezza o di quella di questo progetto, lascia che ti rassicuri dicendo che questa pagina non sarà mai hackerata. Ho 2FA (TOTP) abilitato su tutti gli account che hanno un legame con Picocrypt (GitHub, Google, Reddit, Ubuntu One/Snapcraft, Discord, ecc.), oltre alla crittografia full-disk su tutti i miei dispositivi portatili. Per una ulteriore sicurezza, Picocrypt utilizza i miei fork delle dipendenze e recupero l'upstream solo quando ho dato un'occhiata ai cambiamenti e sono sicuro che non ci siano problemi di sicurezza. Questo significa che se una dipendenza viene hackerata o rimossa dall'autore, Picocrypt userà il mio fork di essa e non subirà alcuna alterazione. Puoi sentirti sicuro nell'usare Picocrypt.

Nota: le versioni attuali di Picocrypt faranno una richiesta di rete a raw.githubusercontent.com per controllare se è disponibile una versione più recente di Picocrypt. In futuro, quando Picocrypt sarà ultimato, rimuoverò il controllo degli aggiornamenti e Picocrypt diventerà completamente isolato dalla rete.

# Community

Ecco alcuni posti in cui puoi rimanere aggiornato su Picocrypt e partecipare attivamente:

<ul>

<li><a  href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>

<li><a  href="https://discord.gg/8QM4A2caxH">Discord</a></li>

</ul>

Consiglio vivamente a tutti di unirsi al subreddit di Picocrypt perché tutti gli aggiornamenti ed i sondaggi saranno pubblicati lì. Ricordatevi di fidarvi solo dei social network che ho linkato e fate attenzione ai truffatori che potrebbero spacciarsi per me. Non vi chiederò mai la vostra password e chiunque lo faccia non sono io. Non vi chiederò mai di scaricare un file da un link sospetto e chiunque lo faccia non sono io.

# Stargazers

Come sta andando Picocrypt? Guarda tu stesso!

[![Stargazers over time](https://starchart.cc/HACKERALERT/Picocrypt.svg)](https://starchart.cc/HACKERALERT/Picocrypt)

# Donazioni

Se trovi Picocrypt utile, considera la possibilità di donare su <a  href="https://paypal.me/evanyiwensu">PayPal</a>. Sto realizzando questo software in modo completamente gratuito e mi piacerebbe avere dei sostenitori che mi motivino a continuare il mio lavoro su Picocrypt.

# Ringraziamenti

Un grazie dal profondo del mio cuore alle persone di Open Collective che hanno dato un contributo significativo:

<ul>

<li>jp26 ($50)</li>

<li>guest-116103ad ($50)</li>

<li>Tybbs ($10)</li>

<li>N. Chin ($10)</li>

<li>Manjot ($10)</li>

<li>Phil P. ($10)</li>

<li>donor39 (backer)</li>

<li>Pokabu (backer)</li>

</ul>

Voi siete le persone che mi ispirano a lavorare su Picocrypt e a renderlo gratuito per tutti!

Inoltre, un enorme grazie alle seguenti cinque persone, che sono state le prime a donare e sostenere Picocrypt:

<ul>

<li>W.Graham</li>

<li>N. Chin</li>

<li>Manjot</li>

<li>Phil P.</li>

<li>E. Zahard</li>

</ul>

Inoltre, un enorme ringraziamento a queste persone, che hanno aiutato a tradurre Picocrypt e a renderlo più accessibile al mondo:

<ul>

<li>@umitseyhan75 per il Turco</li>

<li>@digitalblossom e @Pokabu per il Tedesco</li>

<li>@zeeaall per il Portoghese Brasiliano</li>

<li>@kurpau per il Lituano</li>

<li>u/francirc per lo Spagnolo</li>

<li>yn per il Russo</li>

<li>@Etim-Orb per l'Ungherese</li>

<li>@Minibus93 per l'Italiano</li>

</ul>

Infine, grazie a queste persone per avermi aiutato quando mi serviva:

<ul>

<li>Fuderal su Discord per avermi aiutato a configurare il server discord</li>

<li>u/greenreddits per il costante feedback e il supporto</li>

<li>u/Tall_Escape per avermi aiutato a testare Picocrypt</li>

<li>u/NSABackdoors per aver effettuato molti test</li>

<li>@samuel-lucas6 per i feedback, suggerimenti e il supporto</li>

</ul>
