<p>English | <a href="/translations/french.md">Français</a> | <a href="/translations/spanish.md">Español</a> | <a href="/translations/german.md">Deutsch</a> | <a href="/translations/portuguese.md">Português</a> | <a href="/translations/turkish.md">Türkçe</a> | <a href="/translations/chinese.md">中文</a> | <a href="/translations/russian.md">русский</a> | <a href="/translations/hungarian.md">Magyar</a> | <a href="/translations/italian.md">Italiano</a> | <a href="/translations/persian.md">پارسی</a> | <a href="/translations/polish.md">Polski</a></p>
<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p> 

Picocrypt dosyalarınızı korumak için kullanabileceğiniz çok küçük (bu yüzden <i>Pico</i>, çok basit ve buna rağmen çok güvenli bir şifreleme aracıdır. Güvenlik, basitlik ve güvenilirliğin odakta olduğu, şifrelemede <i>ilk-tercih</i> bir araç olması için tasarlanmıştır. Picocrypt, NSA gibi üç harfli teşkilatlara karşı bile üst düzey güvenlik sağlamak için güvenli XChaCha20 şifreleyicisini ve Argon2id anahtar türetme fonksiyonunu kullanır. Olabilecek en üst düzey güvenlik için tasarlanmıştır, güvenlik açısından mutlak suretle taviz vermez ve Go'nun standart x/crypto modülleri kullanılarak hazırlanmıştır. <strong>Gizliliğiniz ve güvenliğiniz saldırı altında. Dosyalarınızı Picocrypt ile koruyor olmanın verdiği güvenle buna bir son verin.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="318" alt="Picocrypt"></p>

# Finansman Sağlama
**Cure53'ten bir güvenlik denetimi için para toplamak üzere, lütfen <a href="https://opencollective.com/picocrypt">Open Collective</a> üzerinden (kripto kabul edilir) Picocrypt'e bağış yapın. Bu projeyi hayata geçirebilmek için saatlerce çalıştım ve bundan herhangi bir parasal gelir elde etmiyorum, bu yüzden denetleme ücretini kendi başıma ödeyemem. <i>Picocrypt'in, topluluğunun desteğine ihtiyacı var.**</i>

# Downloads
**Önemli**: Etrafta "Picocrypt" adıyla teşekkül eden birden fazla varlık var. Mesela PicoCrypt, kırılmış bir şifreleyici kullanan eski bir araç. Ayrıca, ERC (Avrupa Araştırma Konseyi) tarafından finanse edilen PICOCRYPT isimli bir araşma projesi var. Hatta hiçbir zaman kaydını yaptırmadığım ve Picocrypt ile ilişkili alan adları var Lütfen bu alakasız projelerin hiçbirini Picocrypt (bu proje) ile karıştırmayın. Gerçek ve arka kapı içermeyen Picocrypt'i edindiğinizden emin olmak için özellikle ve yalnızca bu depodan indirmeye özen gösterin. Picocrypt'i başkalarıyla paylaşmak istediğinizde oluşabilecek herhangi bir karışıklığı önlemek için bu deponun bağlantısını sağladığınızdan emin olun.

## Windows
Windows için Picocrypt olabildiğince basittir. Windows için en güncel ve taşınabilir uygulamayı indirmek için <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.31/Picocrypt.exe">buraya</a> tıklayın. Eğer Microsoft Defender ya da kullandığınız anti-virüs yazılımı Picocrypt'i virüs olarak işaretlerse, herkesin iyiliği için lütfen üzerinize düşeni yapın ve bunun bir yanlış-pozitif olduğunu ilgili yazılımın üreticisine bildirin.

## macOS
Picocrypt for macOS is very simple as well. Download Picocrypt <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.31/Picocrypt.app.zip">here</a>, extract the zip file, and run Picocrypt which is inside. If you can't open Picocrypt because it's not from a verified developer, control-click on Picocrypt and hit Open to bypass the warning. Keep in mind that Picocrypt runs through Rosetta 2 and requires OpenGL, and may not work in the future should Apple remove either.

## Linux
To use Picocrypt on Linux, you can download the AppImage <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.31/Picocrypt.AppImage">here</a>. While this AppImage should work on most systems, Linux is a mess when it comes to cross-distro and cross-release compatibility, so if the AppImage doesn't work, you can run Picocrypt through Wine or from source using the instructions in the `src/` directory.

## Paranoid Pack
The Paranoid Pack is a compressed archive that contains executables for Windows, macOS, and Linux, including the source code and dependencies. As long as you have it stored in a place you can access, you'll be able to open it and use Picocrypt on any desktop operating system in case this repository mysteriously vanishes or the entire Internet burns down. Think of it as a seed vault for Picocrypt; as long as one person has the Paranoid Pack within reach, they can share it with the rest of the world and keep Picocrypt functional in case of catastrophic events. The best way to ensure Picocrypt is accessible many decades from now is to keep a Paranoid Pack in a safe place. Get your copy <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.31/Paranoid.zip">here</a>.

# Why Picocrypt?
Why should you use Picocrypt instead of VeraCrypt, 7-Zip, BitLocker, or Cryptomator? Here are a few reasons why you should choose Picocrypt:
<ul>
	<li>Unlike BitLocker and most cloud services, Picocrypt and its dependencies are completely open-source and auditable. You can verify for yourself that there aren't any backdoors or flaws.</li>
	<li>Picocrypt is <i>tiny</i>. While Cryptomator is over 50 MiB and VeraCrypt is over 20 MiB, Picocrypt sits at just 3 MiB, about the size of a medium-resolution photo. And that's not all - Picocrypt is portable (doesn't need to be installed) and doesn't require administrator/root privileges.</li>
	<li>Picocrypt is easier and more productive to use than VeraCrypt. To encrypt files with VeraCrypt, you'd have to spend a minute or two just setting up a volume. With Picocrypt's simple UI, all you have to do is drag and drop your files, enter a password, and hit Encrypt. All the complex procedures are handled by Picocrypt internally. Who said secure encryption can't be simple?</li>
	<li>Picocrypt is designed for security. 7-Zip is an archive utility and not an encryption tool, so its focus is not on security. Picocrypt, however, is built with security as the number one priority. Every part of Picocrypt exists for a reason and anything that could impact the security of Picocrypt is removed. Picocrypt is built with cryptography you can trust.</li>
	<li>Picocrypt authenticates data in addition to protecting it, preventing hackers from maliciously modifying sensitive data. This is useful when you are sending encrypted files over an insecure channel and want to be sure that it arrives untouched.</li>
	<li>Picocrypt actively protects header data from corruption by adding extra Reed-Solomon parity bytes, so if part of a volume's header (which contains important cryptographic components) corrupts (e.g., hard drive bit rot), Picocrypt can still recover the header and decrypt your data with a high success rate. Picocrypt can also encode the entire volume with Reed-Solomon to prevent any corruption to your important files.</li>
</ul>

# Comparison
Here's how Picocrypt compares to other popular encryption tools.

|                | Picocrypt      | VeraCrypt      | 7-Zip GUI      | BitLocker      | Cryptomator    |
| -------------- | -------------- | -------------- | -------------- | -------------- | -------------- |
| Free           |✅ Yes         |✅ Yes          |✅ Yes         |✅ Bundled      |✅ Yes         |
| Open Source    |✅ GPLv3       |✅ Multi        |✅ LGPL        |❌ No           |✅ GPLv3       |
| Cross-Platform |✅ Yes         |✅ Yes          |❌ No          |❌ No           |✅ Yes         |
| Size           |✅ 3 MiB       |❌ 20 MiB       |✅ 2 MiB       |✅ N/A          |❌ 50 MiB      |
| Portable       |✅ Yes         |✅ Yes          |❌ No          |✅ Yes          |❌ No          |
| Permissions    |✅ None        |❌ Admin        |❌ Admin       |❌ Admin        |❌ Admin       |
| Ease-Of-Use    |✅ Easy        |❌ Hard         |✅ Easy        |✅ Easy         |🟧 Medium      |
| Cipher         |✅ XChaCha20   |✅ AES-256      |✅ AES-256     |🟧 AES-128      |✅ AES-256     |
| Key Derivation |✅ Argon2      |🟧 PBKDF2       |❌ SHA-256     |❓ Unknown      |✅ Scrypt      |
| Data Integrity |✅ Always      |❌ No           |❌ No          |❓ Unknown      |✅ Always      |
| Reed-Solomon   |✅ Yes         |❌ No           |❌ No          |❌ No           |❌ No          |
| Compression    |✅ Yes         |❌ No           |✅ Yes         |✅ Yes          |❌ No          |
| Telemetry      |✅ None        |✅ None         |✅ None        |❓ Unknown      |✅ None        |
| Audited        |🟧 Planned     |✅ Yes          |❌ No          |❓ Unknown      |✅ Yes         |

# Features
Picocrypt is a very simple tool, and most users will intuitively understand how to use it in a few seconds. On a basic level, simply dropping your files, entering a password, and hitting Encrypt is all that's needed to encrypt your files. Dropping the output back into Picocrypt, entering the password, and hitting Decrypt is all that's needed to decrypt those files. Pretty simple, right?

While being simple, Picocrypt also strives to be powerful in the hands of knowledgeable and advanced users. Thus, there are some additional options that you may use to suit your needs.
<ul>
	<li><strong>Password generator</strong>: Picocrypt provides a secure password generator that you can use to create cryptographically secure passwords. You can customize the password length, as well as the types of characters to include.</li>
	<li><strong>Comments</strong>: Use this to store notes, information, and text along with the file (it won't be encrypted). For example, you can put a description of the file you're encrypting before sending it to someone. When the person you sent it to drops the file into Picocrypt, your description will be shown to that person.</li>
	<li><strong>Keyfiles</strong>: Picocrypt supports the use of keyfiles as an additional form of authentication (or the only form of authentication). Not only can you use multiple keyfiles, but you can also require the correct order of keyfiles to be present for a successful decryption to occur. A particularly good use case of multiple keyfiles is creating a shared volume, where each person holds a keyfile, and all of them (and their keyfiles) must be present to decrypt the shared volume. By checking the "Require correct order" box and dropping your keyfile in last, you can also ensure that you'll always be the one clicking the Decrypt button.</li>
	<li><strong>Paranoid mode</strong>: Using this mode will encrypt your data with both XChaCha20 and Serpent in a cascade fashion, and use HMAC-SHA3 to authenticate data instead of BLAKE2b. This is recommended for protecting top-secret files and provides the highest level of practical security attainable. For a hacker to crack your encrypted data, both the XChaCha20 cipher and the Serpent cipher must be broken, assuming you've chosen a good password. It's safe to say that in this mode, your files are impossible to crack.</li>
	<li><strong>Reed-Solomon</strong>: This feature is very useful if you are planning to archive important data on a cloud provider or external medium for a long time. If checked, Picocrypt will use the Reed-Solomon error correction code to add 8 extra bytes for every 128 bytes of data to prevent file corruption. This means that up to ~3% of your file can corrupt and Picocrypt will still be able to correct the errors and decrypt your files with no corruption. Of course, if your file corrupts very badly (e.g., you dropped your hard drive), Picocrypt won't be able to fully recover your files, but it will try its best to recover what it can. Note that this option may slow down encryption and decryption speeds.</li>
	<li><strong>Force decrypt</strong>: Picocrypt automatically checks for file integrity upon decryption. If the file has been modified or is corrupted, Picocrypt will automatically delete the output for the user's safety. If you would like to override these safeguards, check this option. Also, if this option is checked and the Reed-Solomon feature was used on the encrypted volume, Picocrypt will attempt to recover as much of the file as possible during decryption.</li>
	<li><strong>Split into chunks</strong>: Don't feel like dealing with gargantuan files? No worries! With Picocrypt, you can choose to split your output file into custom-sized chunks, so large files can become more manageable and easier to upload to cloud providers. Simply choose a unit (KiB, MiB, GiB, or TiB) and enter your desired chunk size for that unit. To decrypt the chunks, simply drag one of them into Picocrypt and the chunks will be automatically recombined during decryption.</li>
	<li><strong>Compress files</strong>: By default, Picocrypt uses a zip file with no compression to quickly merge files together when encrypting multiple files. If you would like to compress these files, however, simply check this box and the standard Deflate compression algorithm will be applied during encryption.</li>
</ul>

# Security
For more information on how Picocrypt handles cryptography, see <a href="Internals.md">Internals</a> for the technical details. If you're worried about the safety of me or this project, let me assure you that this repository won't be hijacked or backdoored. I have 2FA (TOTP) enabled on all accounts with a tie to Picocrypt (GitHub, Reddit, Google, etc.), in addition to full-disk encryption on all of my portable devices. For further hardening, Picocrypt uses my isolated forks of dependencies and I fetch upstream only when I have taken a look at the changes and believe that there aren't any security issues. This means that if a dependency gets hacked or deleted by the author, Picocrypt will be using my fork of it and remain completely unaffected. You can feel confident about using Picocrypt.

## Signatures
For the paranoid, Picocrypt is signed with PGP. The fingerprint and public key are listed below.

<pre>B342A744BDEEA57B6A583E33A247E73798946F55</pre>
<pre>-----BEGIN PGP PUBLIC KEY BLOCK-----

mDMEYoGUHxYJKwYBBAHaRw8BAQdAvmQA+pdbDB/ynJxHhNDpz6Sb5tgkNuuNJIvw
HYwZtqi0CVBpY29jcnlwdIiTBBMWCgA7FiEEs0KnRL3upXtqWD4zokfnN5iUb1UF
AmKBlB8CGwMFCwkIBwICIgIGFQoJCAsCBBYCAwECHgcCF4AACgkQokfnN5iUb1UZ
RgEA8jbIsdqCr21DWxcqW/eLlbxRkuA8kflVYvWWUxtVqsUA/jQPSDpvA8rakvaL
PIbXjQvrAMkEVIc0HbCzLxr1k3sH
=YFwz
-----END PGP PUBLIC KEY BLOCK-----</pre>

# Community
Consider joining <a href="https://www.reddit.com/r/Picocrypt/">r/Picocrypt</a>. While I won't be active in this subreddit myself, it's still a great place to ask questions and help one another out, especially if something happens to me or this repository in the future. Remember to only trust this specific subreddit and be aware of hackers that might try to impersonate me on other platforms. I will never ask you for your password, and anyone who does is not me. I will never tell you to download a file from a suspicious link, and anyone who does is not me.

# Donations
When I was actively developing Picocrypt, I accepted donations, but now that Picocrypt is complete and production-ready, there's no need anymore. Instead, take your time and effort to share the love of Picocrypt with others. Donations are nice, but being able to help others is a lot more valuable to me than a few spare dollars. Knowing that Picocrypt is helping people secure their files is plenty enough for me.

# FAQ

**Is Picocrypt accepting new features?**

No, Picocrypt is considered feature-complete and won't be getting any new features. Unlike other tools which try to constantly add new features (which introduces new bugs and security holes), Picocrypt focuses on just a few core features but does each of them exceptionally well. Remember Picocrypt's ideology: small, simple, and secure.

**Will Android/iOS be supported?**

No, I don't plan on supporting Android or iOS because they are very different from traditional desktop operating systems and require different toolchains to develop apps for. Due to the nature of open-source software, however, a community-built version of Picocrypt for Android or iOS may appear in the future.

**Why is Picocrypt not updated frequently?**

People seem to have the notion that software must be constantly updated to stay relevant and secure. While this may be true for a lot of the software we use today, it is not for Picocrypt. Picocrypt is "good software" and good software doesn't need constant updates to remain relevant and secure. Good software will always be good software.

**Does the "Delete files" feature shred files?**

No, it doesn't shred any files and just deletes them as your file manager would. On modern storage mediums like SSDs, there is no such thing as shredding a file since wear leveling makes it impossible to overwrite a particular sector. Thus, to prevent giving users a false sense of security, Picocrypt doesn't include any shredding features at all.

# Acknowledgements
A thank you from the bottom of my heart to the people on Open Collective who have made a significant contribution:
<ul>
	<li>Guest ($842)</li>
	<li>YellowNight ($818)</li>
	<li>evelian ($50)</li>
	<li>jp26 ($50)</li>
	<li>guest-116103ad ($50)</li>
	<li>Guest ($27)</li>
	<li>oli ($20)</li>
	<li>Markus ($15)</li>
	<li>Bright ($15)</li>
	<li>Tybbs ($10)</li>
	<li>N. Chin ($10)</li>
	<li>Manjot ($10)</li>
	<li>Phil P. ($10)</li>
	<li>Raymond ($10)</li>
	<li>Cohen ($10)</li>
	<li>donor39 (backer)</li>
	<li>Pokabu (backer)</li>
	<li>akp (backer)</li>
</ul>
<!-- Last updated November 18, 2022 -->

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
	<li>@digitalblossom & @Pokabu26 for German</li>
	<li>@zeeaall & @Rayserzor for Portuguese</li>
	<li>u/francirc & @victorhck for Spanish</li>
	<li>yn & Voron for Russian</li>
	<li>@Etim-Orb for Hungarian</li>
	<li>@Minibus93 for Italian</li>
	<li>Michel for French</li>
	<li>@MasterKia for Persian</li>
	<li>@ungespurv for Polish</li>
	<li>@qaqland for Chinese</li>
</ul>

Finally, thanks to these people/organizations for helping me out when needed:
<ul>
	<li>[ REDACTED ] for helping me create an AppImage for Picocrypt</li>
	<li>u/Upstairs-Fishing867 for helping me test PGP signatures</li>
	<li>u/greenreddits for constant feedback and support</li>
	<li>u/Tall_Escape for helping me test Picocrypt</li>
	<li>u/NSABackdoors for doing plenty of testing</li>
	<li>@samuel-lucas6 for feedback, suggestions, and support</li>
	<li><a href="https://privacytools.io">PrivacyTools</a> for listing Picocrypt</li>
	<li><a href="https://privacyguides.org">PrivacyGuides</a> for listing Picocrypt</li>
	<li>...You?</li>
</ul>



















<!--old translation-->
<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p>

Picocrypt; dosyalarınızı şifreleyebilmenizi, sağlama toplamları (checksums) oluşturabilmenizi, dosyalarınızı kalıcı olarak silebilmenizi ve daha fazlasını yapabilmenizi sağlayan oldukça küçük (bu yüzden <i>Pico</i>), oldukça basit ancak oldukça güvenli bir kriptografi aracıdır. Güvenlik, basitlik ve güvenilirliğe odaklanarak şifreleme için <i>ilk-tercih</i> bir araç olması için tasarlanmıştır. Picocrypt, NSA gibi üç harfli teşkilatlara karşı bile üst düzey güvenlik sağlamak için XChaCha20 şifrelemesi ve HMAC-SHA3 mesaj doğrulama kodunu kullanır. Olabilecek en üst düzey güvenlik için tasarlanmıştır, güvenlik açısından mutlak suretle taviz vermez ve denetlenmiş bir kriptografi kütüphanesi ile oluşturulmuştur. <strong>Gizliliğiniz ve güvenliğiniz saldırı altında. Dosyalarınızı Picocrypt ile koruyor olmanın verdiği güvenle buna bir son verin.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="318" alt="Picocrypt"></p>

# Finansman Sağlama
Cure53 tarafından olası bir denetleme yaptırabilmem için gereken parayı toplayabilmem için lütfen Picocrypt'e <a href="https://opencollective.com/picocrypt">Open Collective</a> üzerinden bağışta bulunun. Bu projeyi hayata geçirebilmek için saatlerce çalıştım ve bundan herhangi bir parasal gelir elde etmiyorum, bu yüzden denetleme ücretini kendi başıma ödeyemem. <i>Picocrypt'in, topluluğunun desteğine ihtiyacı var.</i>

# İndirmeler
<strong>Önemli:</strong> İnternette en son 2005'te güncellenmiş PicoCrypt adlı eski, işe yaramaz ve terk edilmiş bir yazılım müsvettesi var. PicoCrypt'in, Picocrypt(bu proje) ile zerre alakası yoktur. Gerçek ve arka kapı içermeyen Picocrypt'i edindiğinizden emin olmak için özellikle ve yalnızca bu depodan indirmeye özen gösterin.

## Windows
Windows için Picocrypt olabildiğince basittir. Windows için en güncel ve taşınabilir uygulamayı indirmek için <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.31/Picocrypt.exe">buraya</a> tıklayın. Bir ihtimal Picocrypt anti-virüs yazılımınızı tetikleyebilir. Picocrypt bir virüs değildir, bu yüzden lütfen gelecekte bu tür sorunların tekrar yaşanmaması için anti-virüs sağlayıcınıza bu durumun bir yanlış pozitif olduğunu belirtin. Daha eski sürümler, Sürümler sekmesinden erişilebilir.

## macOS
<a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.31/Picocrypt.app.zip">buraya</a> tıklayın.

## Linux
Linux için bir Snap paketi mevcut. Sisteminizde Snap yüklü olduğundan emin olun. Daha sonra, çekirdek snap görüntüsünün yüklü olduğunu garanti altına alın: `sudo snap install core`. Son olarak, Picocrypt'i yükleyin: `sudo snap install picocrypt`.

# Neden Picocrypt?
Neden BitLocker, NordLocker, VeraCrypt, AxCrypt, ya da 7-Zip yerine Picocrypt kullanmalısınız? İşte size Picocrypt kullanmanız için birkaç sebep:
<ul>
	<li>NordLocker, BitLocker, AxCrypt, ve diğer pek çok bulut depolama çözümlerinin aksine, Picocrypt ve bağımlılıkları tamamen açık kaynaklı ve denetime müsaittir. Herhangi bir arka kapı veya kusur olmadığını bizzat kendiniz doğrulayabilirsiniz.</li>
	<li>Picocrypt <i>ufaktır</i>. NordLocker 100MB'ın üzerinde ve VeraCrypt 30MB'tan fazlayken, Picocrypt sadece 3MB, -ki bu da yaklaşık olarak yüksek çözünürlüklü bir görsel kadardır. Ve hepsi bu kadar da değil, Picocrypt taşınabilirdir, yükleme gerektirmez, ve yönetici hakları ya da kök erişim iznine ihtiyaç duymaz.</li>
	<li>Picocrypt, VeraCrypt'a oranla kullanımı daha basit ve daha verimlidir. VeraCrypt ile dosyalarınızı şifreleyebilmek için, bölüm oluştururken hiç olmazsa en azından beş dakika harcamanız gerekir. Picocrypt'in basit kullanıcı arayüzü sayesinde, tüm yapmanız gereken dosyaları sürükleyip bırakmak, bir parola belirlemek ve Başla tuşuna basmak. Picocrypt tüm karmaşık ayarları kendisi halleder.</li>
	<li>Picocrypt güvenlik amacı güdülerek tasarlanmıştır. 7-Zip bir arşiv yazılımı olup asıl odağı şifreleme ya da güvenlik değildir. Lâkin Picocrypt, güvenliği her şeyin üstünde tutacak şekilde oluşturulmuştur. Picocrypt'e ait her bir parça bir sebepten ötürü var ve Picocrypt'in güvenliğine etki edebilecek her şey silinmiştir. Picocrypt güvenebileceğiniz bir kriptografidir.</li>
	<li>Picocrypt dosyaları korumasının yanı sıra onları tasdik de eder, bilgisayar korsanlarının ilgili hassas veriyi zarara sebep olacak şekilde etklemesinin önüne geçer. Bu, güvenli olmayan bir kanal üzerinden şifrelenmiş dosyalar gönderirken ve el değmeden ulaştığından emin olmak istediğinizde kullanışlı bir özellik halini alır. Picocrypt tasdik için güvenli saygınlığı yüksek bir mesaj doğrulama kodu olan Poly1305 kullanır.</li>
	<li>Picocrypt dosyalarınızın bozulmalara karşı aktif olarak korunması için fazladan Reed-Solomon bitleri ekler, bu sayede eğer şifrelenmiş veriniz bozulursa (örneğin, sabit disk arızası gibi), Picocrypt verilerinizi yüksek başarı oranıyla kurtarabilir.</li>
</ul>

# Talimatlar
Picocrypt çok basit bir araçtır ve pek çok kullanıcı saniyeler içinde nasıl kullanılacağını anlayabilir. Yine de, kullanabileceğiniz bazı gelişmiş seçenekleri barındırır.
<ul>
	<li><strong>Dosya üstverisi</strong>: Bu, not almak, bilgilendirmek veya herhangi bir şekilde yazı eklemek (şifrelenmemiş bir şekilde) için kullanılabilir. Söz gelimi, dosyayı birisine göndermeden önce açıklama ekleyebilirsiniz. Gönderdiğiniz kişi ilgili dosyayı Picocrypt'a yerleştirdiği sırada, açıklamanız şifreleme öncesinde görüntülenecektir.</li>
	<li><strong>Değişiklik yapılmış ya da bozulmuş olsa dâhi deşifre edilmiş dosyayı sakla</strong>: Picocrypt deşifrasyon sırasında dosyaların bütünlüğünü otomatik olarak kontrol eder. Eğer dosyada değişikik yapılmış ya da bozulmuşsa, Picocrypt çıktı verisini kullanıcının güvenliğinden taviz vermemek için otomatik olarak siler. Eğer üzerinde değişiklik yapılmış ya da bozulmuş çıktıyı yine de saklamak isterseniz bu ayarı seçin. Ayrıca, eğer bu ayar seçilliyse ve şifreli dosyada Reed-Solomon özelliği kullanılmışsa, Picocrypt, -eğer bozulmuşsa- deşifrasyon sırasında elinden geldiğince veriyi kurtarmaya çalışacaktır.</li>
	<li><strong>Girdi dosyası/dosyalarını ve klasörünü/klasörlerini kalıcı olarak sil</strong>: Şâyet seçiliyse, Picocrypt girdi dosyalarını geri döndürülemez şekilde silmek için sistem dahililerini kullanır. Bu usûl ilgili dosyayı sadece silmekten çok daha güvenlidir, zirâ girdi drumundaki özgün dosya(lar) bilgisayar korsanları tarafından çeşitli özel yazılımlar aracılığıyla kurtarılabilmektedir. Picocrypt, özgün dosyaların geri döndürülemeyeceğinden emin olacak şekilde onları siler.</li>
	<li><strong>Reed-Solomon kullanarak bozulmaların önüne geç</strong>: Bu ayar, önemli verileri bir bulut hizmetinde veya harici sabit diskte uzun süre arşivlemeyi planlıyorsanız oldukça kullanışlıdır. Şâyet seçiliyse, Picocrypt Reed-Solomon hata düzeltme kodunu, her 128 byte'a ek olarak 13 fazladan byte ekleyerek dosya bozulmalarının önüne geçer. Bu, dosyaların en fazla %5'e kadar bozulabileceğini ve Picocrypt'in buna rağmen ilgili bozunumu düzeltip hatasız bir şekilde dosyalarınızı deşifre edebileceği anlamına gelir. Yine de, -barizdir ki-,  eğer dosyanız oldukça fena bir şekilde bozulmuşsa (söz gelimi, sabit diskinizi düşürmeniz durumunda), Picocrypt dosyalarınızı tam anlamıyla kurtaramayacaktır, buna rağemen elinden geldiğince fazla veriyi kurtarmayı deneyecektir. Bu ayarın şifreleme ve deşifrasyon hızına ufak dâhi olsa etki edeceğini untumayın..</li>
</ul>

# Güvenlik
Picocrypt'in kriptografi sürecini nasıl yürüttüüğyle ilgili daha fazla bilgi için <a href="Internals.md">Dahililer</a> bilgilendirmesini okuyun. Eğer Picocrypt'in pratikteki güvenliğiyle ilgili endişeleriniz varsa, bu deponun ele geçirilmeyeceğine dair sizi temin ederim. Picocrypt'e bağlı tüm hesaplarda (GitHub, Google, Reddit, Discord, vb.) 2 faktörlü kimlik doğrulamayı (TOTP) etkinleştirdim, böylece Picocrypt'i kullanma konusunda kendinizi güvende hissedebilirsiniz.

# Topluluk
Picocrypt ile ilgili en son haberleri görebileceğiniz ve dahil olabileceğiniz bazı yerler:
<ul>
	<li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
	<li><a href="https://discord.gg/8QM4A2caxH">Discord</a></li>
</ul>
Yalnızca bu topluluklara güvenin ve benim kimliğime bürünüp size kötü amaçlı yazılım indirtmeyi deneyebilecek dolandırıcılara ve bilgisayar korsanlarına karşı dikkatli olun.

# Yıldız Grafiği
[![Stargazers over time](https://starchart.cc/HACKERALERT/Picocrypt.svg)](https://starchart.cc/HACKERALERT/Picocrypt)

# Katkı
Eğer bir hata bulduysanız ya da özellik talebiniz varsa, lütfen bir Sorun girdisi (Issue) oluşturun. Ortaklık edebilecek çalışma arkadaşlarına açığım, ayrıca tüm Sorun girdilerinin ve çekme taleplerinin (Pull Requests) başım üstünde yeri var. Eğer bir güvenlik açığı tespit ettiyseniz, lütfen bildirmek için <a href="Security.md">buradaki</a> talimatları takip edin.

# Teşekkürler
Open Collective üzerinden yaptıkları kayda değer bağışlardan ötürü bu kişilere özellikle çok teşekkür ederim.
<ul>
	<li>jp26 ($50)</li>
</ul>
Ve ayrıca Patreon'daki destekçilerime:
<ul>
	<li>Frederick Doe</li>
</ul>

Aşağıda listelenmiş olan beş kişiye, Picocrypt'e bağışta ve destekte bulunan ilk kişiler olmalarından ötürü çok teşekkür ederim:
<ul>
	<li>W.Graham</li>
	<li>N. Chin</li>
	<li>Manjot</li>
	<li>Phil P.</li>
	<li>E. Zahard</li>
</ul>

Picocrypt'i tercüme ettikleri ve dünya için daha erişilebilir hale getirmeye yardımcı oldukları için bu insanlara çok teşekkür ederim:
<ul>
	<li>Türkçe için @umitseyhan75</li>
	<li>Almanca için @digitalblossom</li>
	<li>Polonya için @ungespurv</li>
</ul> 

Ayrıca bu kişilere de teşekkürü bir borç bilirim:
<ul>
	<li>Bir Discord sunucusu kurmamdaki yarımlarından ötürü Discord'dan Fuderal</li>
	<li>Düzenli geri bildirim ve desteklerinden ötürü u/greenreddits</li>
	<li>Picocrypt'ı test etmemdeki yardımlarından ötürü u/Tall_Escape</li>
	<li>Değerli geri bildirim ve desteklerinden ötürü (GitHub'dan)samuel-lucas6</li>
</ul>
