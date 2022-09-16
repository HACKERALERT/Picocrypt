<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p>

Picocrypt est un outil de chiffrement très petit (d'où <i>Pico</i>), très simple mais très sécurisé que vous pouvez utiliser pour protéger vos fichiers. Il est conçu pour être l'outil <i>de référence</i> pour le chiffrement, en mettant l'accent sur la sécurité, la simplicité et la fiabilité. Picocrypt utilise le chiffrement sécurisé XChaCha20 cipher et la fonction de dérivation de clé Argon2id pour fournir un haut niveau de sécurité, même en face d'agences à trois lettres comme la NSA. Il est conçu pour une sécurité maximale, ne faisant absolument aucun compromis en matière de sécurité, et il est construit avec les modules x/crypto standard de Go. <strong>Votre vie privée et votre sécurité sont attaquées. Reprenez-en le  contrôle en toute confiance en protégeant vos fichiers avec Picocrypt.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="384" alt="Picocrypt"></p>

# Financement
Veuillez faire un don à Picocrypt sur <a href="https://opencollective.com/picocrypt">Open Collective</a> (les cryptomonnaies sont acceptés) pour collecter des fonds pour un audit potentiel de type Cure53. Comme il s'agit d'un projet sur lequel je passe de nombreuses heures, qui ne me rapportent pas d'argent, je ne peux pas payer moi-même cet audit. <i>Picocrypt a besoin du soutien de sa communauté pour se faire.</i>

# Téléghargements
## Windows
Picocrypt pour Windows est aussi simple que possible. Pour télécharger le dernier exécutable autonome et portable pour Windows, cliquez sur <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.exe">ici</a>. Si Windows Defender ou votre antivirus signale Picocrypt comme un virus, faites en sorte de le soumettre comme un faux positif et ce pour le bien de tous.

## macOS
Picocrypt pour macOS est également très simple. Téléchargez Picocrypt <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.app.zip">ici</a>.

## Linux
Un Snap est disponible pour Linux. En supposant que vous êtes sur un système basé sur Debian, un simple `apt install snapd` et `snap install picocrypt` suffiront. Pour les autres distributions telles que Fedora, des instructions détaillées sont disponibles sur https://snapcraft.io/picocrypt. En raison de la complexité des dépendances et des liens statiques, je ne distribue pas de fichiers binaires .deb ou .rpm autonomes car ils ne seraient pas fiables et donc n'en valent pas la peine. Snapcraft gère automatiquement toutes les dépendances, les temps d'exécution et constitue la méthode qui est recommandée pour exécuter Picocrypt sur n'importe quelle distribution Linux majeure. De plus, Snapcraft offre une meilleure sécurité et conteneurisation que Flatpaks et AppImages, ce qui est important pour un outil de chiffrement comme Picocrypt. Si vous préférez ne pas traiter avec Canonical, rappelez-vous que la construction à partir du source est toujours une option.

## Paranoid Packs
Le Paranoid Packs est une archive compressée qui contient des exécutables pour chaque version de Picocrypt historiquement publiée pour Windows, macOS et Linux. Tant que vous l'a sauvegardé dans un endroit auquel vous pouvez accéder, vous pourrez ouvrir l'archive et utiliser n'importe quelle version de Picocrypt au cas où ce référentiel disparaîtrait mystérieusement ou si tout Internet brûlait. Considérez lae Paranoid Pack comme un coffre-fort de semences pour Picocrypt. Tant qu'une personne a le Paranoid Pack à portée de main, elle peut le partager avec le reste du monde et garder Picocrypt fonctionnel en cas d'événements catastrophiques comme GitHub s'arrêtant soudainement ou la NSA me capturant (juste au cas où, vous savez ?) . La meilleure façon de s'assurer que Picocrypt est accessible dans plusieurs décennies est de conserver un pack Paranoid dans un endroit sûr. Donc, si vous craignez de ne plus pouvoir accéder à Picocrypt à l'avenir, eh bien, voici votre solution. Rendez-vous simplement dans l'onglet Versions et procurez-vous en une copie.

# Pourquoi Picocrypt ?
Pourquoi devriez-vous utiliser Picocrypt au lieu de BitLocker, NordLocker, VeraCrypt, AxCrypt ou 7-Zip ? Voici quelques raisons pour lesquelles vous devriez choisir Picocrypt :
<ul>

<li>Contrairement à NordLocker, BitLocker, AxCrypt et à la plupart des fournisseurs de stockage cloud, Picocrypt et ses dépendances sont entièrement open source et auditables. Vous pouvez vérifier par vous-même qu'il n'y a pas de portes dérobées ou de failles.</li>

<li>Picocrypt est <i>minuscule</i>. Alors que NordLocker dépasse 50 Mo et que VeraCrypt dépasse 20 Mo, Picocrypt ne pèse que 3 Mo, soit environ la taille d'une photo haute résolution. Et ce n'est pas tout - Picocrypt est portable (n'a pas besoin d'être installé) et ne nécessite pas de privilèges administrateur/racine.</li>

<li>Picocrypt est plus facile et plus productif à utiliser que VeraCrypt. Pour chiffrer des fichiers avec VeraCrypt, vous devez passer au moins cinq minutes à configurer un volume. Avec l'interface utilisateur simple de Picocrypt, tout ce que vous avez à faire est de faire glisser et déposer vos fichiers, d'entrer un mot de passe et d'appuyer sur Démarrer. Toutes les procédures complexes sont gérées par Picocrypt en interne. Qui a dit que le chiffrement sécurisé ne pouvait pas être simple ?</li>

<li>Picocrypt est conçu pour la sécurité. 7-Zip est un utilitaire d'archivage et non un outil de chiffrement, il n'est donc pas axé sur la sécurité. Picocrypt, lui, est construit avec la sécurité comme priorité numéro un. Chaque partie de Picocrypt existe pour éliminer tout ce qui pourrait avoir un impact sur la sécurité de Picocrypt. Picocrypt est construit avec une cryptographie de confiance.</li>

<li>Picocrypt authentifie les données en plus de les protéger, empêchant les pirates de modifier de manière malveillante les données sensibles. Ceci est utile lorsque vous envoyez des fichiers chiffrés via un canal non sécurisé et que vous voulez être sûr qu'ils arrivent intacts.</li>

<li>Picocrypt protège activement les données d'en-tête chiffrées contre la corruption en ajoutant des octets de parité Reed-Solomon supplémentaires, donc si une partie de l'en-tête d'un volume (qui contient des composants cryptographiques importants) est corrompue (par exemple, le cryptagnon désiré du disque dur), Picocrypt peut toujours récupérer l'en-tête et décryptez vos données avec un taux de réussite élevé. Picocrypt peut également encoder l'intégralité du volume avec Reed-Solomon pour éviter toute corruption de vos fichiers importants.</li>

</ul>

Toujours pas convaincu ? Voir ci-dessous pour encore plus de raisons pour lesquelles Picocrypt se démarque des autres.

# Caractéristiques

Picocrypt est un outil très simple, et la plupart des utilisateurs comprendront intuitivement comment l'utiliser en quelques secondes. À la base, il suffit de déposer vos fichiers, d'entrer un mot de passe et d'appuyer sur Démarrer pour chiffrer vos fichiers. Assez simple, non?

Tout en étant simple, Picocrypt s'efforce également d'être puissant entre les mains d'utilisateurs avertis et avancés. Ainsi, il existe des options supplémentaires que vous pouvez utiliser pour répondre à vos besoins.

<ul>

<li><strong>Générateur de mots de passe</strong> : Picocrypt fournit un générateur de mots de passe sécurisé que vous pouvez utiliser pour créer des mots de passe sécurisés par chiffrement. Vous pouvez personnaliser la longueur du mot de passe, ainsi que les types de caractères à inclure.</li>

<li><strong>Métadonnées du fichier</strong> : utilisez-les pour stocker des notes, des informations et du texte avec le fichier (il ne sera pas chiffré). Par exemple, vous pouvez mettre une description du fichier que vous cryptez avant de l'envoyer à quelqu'un. Lorsque la personne à qui vous l'avez envoyé dépose le fichier dans Picocrypt, votre description lui sera montrée.</li>

<li><strong>Fichiers clés</strong> : Picocrypt prend en charge l'utilisation de fichiers clés comme forme d'authentification supplémentaire. Non seulement vous pouvez utiliser plusieurs fichiers de clés, mais vous pouvez également exiger que le bon ordre des fichiers de clés soit présent, pour qu'un déchiffrement réussisse. Un cas d'utilisation particulièrement intéressant  avec  plusieurs fichiers de clés est la création d'un volume partagé, où chaque personne détient un fichier de clés, et tous (et leurs fichiers de clés) doivent être présents afin de déchiffrer le volume partagé.</li>

<li><strong>Mode paranoïaque</strong> : l'utilisation de ce mode chiffrera vos données avec XChaCha20 et Serpent en cascade, et utilisera HMAC-SHA3 pour authentifier les données au lieu de BLAKE2b. Ceci est recommandé pour protéger les fichiers top-secrets et offre le plus haut niveau de sécurité pratique possible. Pour qu'un pirate informatique déchiffre vos données chiffrées, le chiffrement XChaCha20 et le chiffrement Serpent doivent être cassés, en supposant que vous avez choisi un bon mot de passe.</li>

<li><strong>Prévenir la corruption à l'aide de Reed-Solomon</strong> : cette fonctionnalité est très utile si vous prévoyez d'archiver des données importantes sur un fournisseur de cloud ou un support externe pendant une longue période. Si cette case est cochée, Picocrypt utilisera le code de correction d'erreur Reed-Solomon pour ajouter 8 octets supplémentaires tous les 128 octets afin d'éviter la corruption des fichiers. Cela signifie que jusqu'à ~3% de votre fichier peut être corrompu et Picocrypt pourra toujours corriger les erreurs et décrypter vos fichiers sans corruption. Bien sûr, si votre fichier est très gravement corrompu (par exemple, vous avez fait tomber votre disque dur), Picocrypt ne pourra pas récupérer complètement vos fichiers, mais il fera de son mieux pour récupérer ce qu'il peut. Notez que cette option ralentira considérablement le chiffrement et le déchiffrement.</li>

<li><strong>Conserver la sortie déchiffrée même si elle est corrompue ou modifiée</strong> : Picocrypt vérifie automatiquement l'intégrité lors du déchiffrement. Si le fichier a été modifié ou est corrompu, Picocrypt supprimera automatiquement la sortie pour la sécurité de l'utilisateur. Si vous souhaitez conserver les données corrompues ou modifiées après décryptage, cochez cette option. De plus, si cette option est cochée et que la fonction Reed-Solomon a été utilisée sur le fichier crypté, Picocrypt tentera de récupérer autant que possible le contenu du fichier lors du décryptage.</li>

<li><strong>Divisez les fichiers en morceaux</strong> : vous n'avez pas envie de gérer des fichiers gargantuesques ? Pas de soucis! Avec Picocrypt, vous pouvez choisir de diviser votre fichier de sortie en morceaux de taille personnalisée, afin que les fichiers volumineux deviennent plus faciles à gérer et à télécharger vers les fournisseurs de cloud. Choisissez simplement une unité (Kibibit, Mebibyte, or Gibibyte)) et entrez le nombre souhaité pour cette unité. Pour déchiffrer les morceaux, faites simplement glisser l'un d'entre eux dans Picocrypt, et les morceaux seront automatiquement recombinés pendant le déchiffrement.</li>

</ul>

# Sécurité

Pour plus d'informations sur la façon dont Picocrypt gère la cryptographie, voir <a href="/Internals.md">Internes</a> pour les détails techniques. Si vous vous inquiétez pour ma sécurité ou celle de ce projet, laissez-moi vous assurer que ce référentiel ne sera pas piraté ou détourné. J'ai 2FA (TOTP) activé sur tous les comptes liés à Picocrypt (GitHub, Google, Reddit, Ubuntu One/Snapcraft, Discord, etc.), en plus du cryptage complet du disque sur tous mes appareils portables. Pour un renforcement supplémentaire, Picocrypt utilise mes fourches isolées de dépendances et je ne récupère en amont que lorsque j'ai examiné les modifications et que je pense qu'il n'y a pas de problèmes de sécurité. Cela signifie que si une dépendance est piratée ou supprimée par l'auteur, Picocrypt en utilisera mon fork et ne sera absolument pas affecté. Vous pouvez être sûr d'utiliser Picocrypt.

# Communauté

Voici quelques endroits où vous pouvez rester à jour à propos de  Picocrypt et vous impliquer :

<ul>
  <li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
  <li><a href="https://discord.gg/8QM4A2caxH">Discorde</a></li>
</ul><br>

  
  Je vous recommande fortement de rejoindre le subreddit de Picocrypt car toutes les mises à jour et les sondages y seront publiés. N'oubliez pas de ne faire confiance qu'à ces réseaux sociaux et d'être conscient des pirates qui pourraient essayer de se faire passer pour moi. Je ne vous demanderai jamais votre mot de passe, et quiconque le fait n'est pas moi. Je ne vous dirai jamais de télécharger un fichier à partir d'un lien suspect, et quiconque le fait n'est pas moi.
# Stargazers

Comment va Picocrypt ? Jetez un œil ci-dessous pour le savoir.

[![Stargarzers over time](https://starchart.cc/HACKERALERT/Picocrypt.svg)](https://starchart.cc/HACKERALERT/Picocrypt)

# Dons

Si vous trouvez Picocrypt utile, veuillez envisager de donner un don via mon [PayPal](https://paypal.me/evanyiwensu). Je fournis ce logiciel entièrement gratuitement et j'aimerais avoir des supporters  et qui me motivent  à continuer mon travail sur Picocrypt.

# Merci

Un merci du fond du cœur aux personnes d'Open Collective qui ont apporté une contribution significative :

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

Vous êtes les personnes qui m'inspirent pour travailler sur Picocrypt et le mettre à disposition  gratuitement pour tout le monde !

Aussi, un grand merci à la liste suivante de cinq personnes, qui ont été les premières à faire un don et à soutenir Picocrypt :

<ul>
<li>W.Graham</li>
<li>N. Chin</li>
<li>Manjot</li>
<li>Phil P.</li>
<li>E. Zahard</li>
</ul>

De plus, un grand merci à ces personnes, qui ont aidé à traduire Picocrypt et à le rendre plus accessible au monde :

<ul>
<li>@umitseyhan75 pour le Turc</li>
<li>@digitalblossom pour l'Allemand </li>
<li>@zeeaall pour le Portugais brésilien</li>
<li>@kurpau pour le Lithuanien</li>
<li>u/francirc pour  l'Espagnol</li>
<li>yn pour le Russe</li>
<li>@Etim-Orb pour le Hongrois</li>
<li>@Minibus93 pour l'Italien</li>
  <li>Michel pour francais</li>
</ul>

Enfin, merci à ces personnes de m'avoir aidé dés que le besoin s'en est fait sentir :

<ul>
<li>Fuderal sur Discord pour m'avoir aidé à configurer un serveur Discord</li>
<li>u/greenreddits pour un feedback constant et son support</li>
<li>u/Tall_Escape pour son aide pour tester Picocrypt</li>
<li>u/NSABackdoors pour ses tests nombreux</li>
<li>@samuel-lucas6 pour son feedback, ses suggestions, et son support</li>
</ul>
