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
macOS için Picocrypt'te aynı şekilde olabildiğince basittir. Picocrypt'i <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.31/Picocrypt.app.zip">buradan</a> indirin, zip dosyasından çıkarın, ve içindeki Picocrypt'i çalıştırın. Eğer Picocrypt'i doğrulanmamış bir geliştiriciden olduğu için açamıyorsanız, kontrol tuşuna basarak uygulama simgesine tıklayın, ardından kestirme menüsünde Aç’ı seçin, bu şekilde uyarıyı aşabilirsiniz. Picocrypt'in Rosetta 2 üzerinde çalıştığını ve OpenGL gerektirdiğini, ve gelecekte Apple'ın bunları silmesi durumunda çalışmayabileceğini aklınızda bulundurun.

## Linux
Linux üzerinde Picocrypt kullanabilemek için <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.31/Picocrypt.AppImage">buradan</a> AppImage dosyasını indirebilirsiniz. Bu AppImage muhtelen çoğu sistemde çalışacaktır, ancak Linux çapraz-dağıtım ve çapraz-sürüm uyumluluğu konusunda çok karman çorman bir durumda, dolayısıyla eğer AppImage çalışmazsa, Picocrypt'i Wine üzerinden veya `src/` konumundaki talimatları kullanarak doğrudan kaynaktan çalıştırabilirsiniz.

## Paranoid Paketi
Paranoid Paketi, kaynak kodunun ve bağımlılıkların da dahil olduğu Windows, macOS ve Linux için yürütülebilir uygulamaları içeren sıkıştırılmış bir arşivdir. Erişebileceğiniz bir yerde sakladığınız sürece, bu deponun gizemli bir şekilde kaybolması veya tüm internetin yanıp kül olması durumunda bu arşivi açıp Picocrypt'i herhangi bir masaüstülü işletim sisteminde kullanabileceksiniz. Bunu Picocrypt için bir tohum bankası olarak düşünebilirsiniz; en azından bir kişide Paranoid Paketi olduğu sürece Picocrypt'i dünyanın geri kalanıyla paylaşabilir ve felaket senaryolarında Picocrypt'i çalışır vaziyette tutabilir. Picocrypt'in bundan on yıllarca sonra bile erişilebilir olmasını sağlamanın en iyi yolu, bir Paranoid Paketini güvenli bir yerde muhafaza etmektir. Siz de bir nüshasını <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.31/Paranoid.zip">buradan</a> alabilirsiniz.

# Neden Picocrypt?
Neden VeraCrypt, 7-Zip, BitLocker, ya da Cryptomator yerine Picocrypt kullanmalısınız? İşte size Picocrypt kullanmanız için birkaç sebep:
<ul>
	<li>BitLocker ve çoğu bulut servisinin aksine, Picocrypt ve bağımlılıkları tamamen açık kaynaklı ve denetlenebilirdir. Herhangi bir kusur ya da arka kapı olmadığını bizatihi kendiniz doğrulayabilirsiniz.</li>
	<li>Picocrypt <i>ufaktır</i>. Cryptomator 50 MiB'nin üzerinde ve VeraCrypt 20 MiB'nin üzerindeyken, Picocrypt sadece 3 MiB, ki bu yaklaşık olarak orta çözünürlüklü bir fotoğraf boyutudur. Ve hepsi bu kadar da değil - Picocrypt taşınabilirdir (yüklenmesine gerek yoktur) ve yönetici/kök izinlerine ihtiyaç duymaz.</li>
	<li>Picocrypt'in kullanımı VeraCrypt'ten daha kolay ve verimlidir. VeraCrypt ile sırf birim oluşturmak için sadece bir ya da iki dakika harcamak zorundasınız. Picocrypt'in basit kullanıcı arayüzü ile tüm yapmanız gereken dosyalarınızı sürükleyip bırakmak, bir parola girmek ve Şifrele'ye basmaktır. Tüm karmaşık prosedürler dahili olarak Picocrypt tarafından gerçekleştirilir. Güvenli şifrelemenin basit olamayacağını kim söyledi?</li>
	<li>Picocrypt güvenlik için dizayn edilmiştir. 7-Zip bir arşivleme aracıdır, şifreleme aracı değil, dolayısıyla odağı güvenlik üzerine değildir. Öte yandan Picocrypt, birinci önceliği güvenlik olarak hazırlanmıştır. Picocrypt'e ait her bir parça bir sebepten ötürü var ve Picocrypt'in güvenliğine etki edebilecek her şey silinmiştir. Picocrypt güvenebileceğiniz bir kriptografi ile hazırlanmıştır.</li>
	<li>Picocrypt dosyaları korumasının yanı sıra onları tasdik te eder, bilgisayar korsanlarının hassas verileri zarara sebep olacak şekilde değiştirmesinin önüne geçer. Bu, şifrelenmiş dosyalarınızı güvenli olmayan bir kanal üzerinden gönderirdiğinizde, el değmeden ulaştığına emin olmanızı sağlar.</li>
	<li>Picocrypt dosyalarınızın bozulmalara karşı aktif olarak korunması için mevcut üstbilgilerine ek olarak Reed-Solomon eşitlik bitleri ekler, bu sayede bir birimin üstbilgisinin bir parçası (ki bu kısım önemli kriptografik bileşenleri içerir) bozulsa bile (örn. sabit disk bit rot'u) Picocrypt yüksek bir başarı oranıyla üstbilgiyi kurtarabilir ve dosyalarınızın şifresini çözebilir. Ayrıca Picocrypt bütün bir birimi Reed-Solomon ile şifreleyerek önemli dosyalarınızın kaybolmasının da önünde geçebilir.</li>
</ul>

# Kıyaslama
İşte Picocrypt'in diğer popüler şifreleme araçları ile kıyası.

|                   | Picocrypt      | VeraCrypt      | 7-Zip GUI      | BitLocker      | Cryptomator    |
| --------------    | -------------- | -------------- | -------------- | -------------- | -------------- |
|Ücretsiz           |✅Evet          |✅Evet          |✅Evet         |✅Gömülü        |✅Evet         |
|Açık Kaynak        |✅GPLv3         |✅Çoklu-Lisans  |✅LGPL         |❌Hayır         |✅GPLv3        |
|Çapraz-Platform    |✅Evet          |✅Evet          |❌Hayır        |❌Hayır         |✅Evet         |
|Boyut              |✅3 MiB         |❌20 MiB        |✅2 MiB        |✅Belirsiz      |❌50 MiB       |
|Taşınabilir        |✅Evet          |✅Evet          |❌Hayır        |✅Evet          |❌Hayır        |
|İzinler            |✅Yok           |❌Yönetici      |❌Yönetici     |❌Yönetici      |❌Yönetici     |
|Kullanım Kolaylığı |✅Kolay         |❌Zor           |✅Kolay        |✅Kolay         |🟧Orta         |
|Şifreleyici        |✅XChaCha20     |✅AES-256       |✅AES-256      |🟧AES-128       |✅AES-256      |
|Anahtar Türetimi   |✅Argon2        |🟧PBKDF2        |❌SHA-256      |❓Bilinmiyor    |✅Scrypt       |
|Veri Bütünlüğü     |✅Her zaman     |❌Hayır         |❌Hayır        |❓Bilinmiyor    |✅Her zaman    |
|Reed-Solomon       |✅Evet          |❌Hayır         |❌Hayır        |❌Hayır         |❌Hayır        |
|Sıkıştırma         |✅Evet          |❌Hayır         |✅Evet         |✅Evet          |❌Hayır        |
|Telemetri          |✅Hiç yok       |✅Hiç yok       |✅Hiç yok      |❓Bilinmiyor    |✅Hiç yok      |
|Denetlendi         |🟧Planlandı     |✅Evet          |❌Hayır        |❓Bilinmiyor    |✅Evet         |

# Özellikler
Picocrypt çok basit bir araçtır ve çoğu kullanıcı onu nasıl kullanacağını birkaç saniye içinde sezgisel olarak anlayacaktır. Temel olarak dosyalarınızı şifrelemek için onları sürükleyip bırakarak ve akabinde bir parola girip Şifrele'ye basmak yeterlidir. Benzer şekilde, şifrelenmiş birimi sürükleip bırakmak, şifrelerken belirlediğiniz parolayı girmek ve Deşifrele'ye basmak dosyalarınızın şifresini çözmek için yeterlidir. Oldukça basit, değil mi?

Picocrypt basit olmakla birlikte, bilgili ve ileri düzey kullanıcıların elinde güçlü olmaya da çalışır. Dolayısıyla, ihtiyaçlarınıza uygun olarak kullanabileceğiniz bazı ek seçenekler sunar.
<ul>
	<li><strong>Parola oluşturucu</strong>: Picocrypt, kriptografik olarak güvenli parolalar oluşturmak için kullanabileceğiniz güvenli bir parola oluşturucu sağlar. Parola uzunluğunu ve eklenecek karakter türlerini istediğiniz gibi özelleştirebilirsiniz.</li>
	<li><strong>Yorumlar</strong>: Bu özelliği şifrelediğiniz dosyalara not eklemek veya başka herhangi bir düz metin formatındaki yazıyı eklemek için kullanabilirsiniz. Eklediğiniz yorumların şifrelenmeyeceğini unutmayın. Örnek vermek gerekirse, birisine göndermeden önce, şifreleyeceğiniz dosyaya bir açıklama ekleyebilirsiniz. Gönderdiğiniz kişi dosyayı Picocrypt'e sürükleyip bıraktığında, yazmış olduğunuz açıklama deşifreleme öncesinde o kişiye gösterilecektir.</li>
	<li><strong>Anahtar dosyası</strong>: Picocrypt, ek bir doğrulama biçimi (veya tek doğrulama biçimi) olarak anahtar dosyalarının kullanımını destekler. Birden fazla anahtar dosyasını aynı anda kullanabilmenizin yanı sıra, dosyların başarılı bir şekilde deşifre edilebilmesi için anahtar dosyalarının hangi sırayla eklendiği bilgisininin belirtmesini de zorunlu kılabilirsiniz. Herkesin kendi anahtar dosyasını kendine sakladığı bölünmüş bir birim oluşturmak birden çok anahtar dosyası kullanımına iyi bir örnek olarak verilebilir. Bu durumda her bir birey şifrelenmiş birimin kendisine ait bir parçasına ve anahtar dosyasına sahip olur. Herkes kendi bölünmüş parçasını ve anahtar dosyasını (doğru sıralamada olmak kaydıyla) sunmadığı müddetçe dosya(lar) deşifre edilemez. "Doğru sıralama iste" kutucuğunu işaretlemek ve kendi anahtar dosyanızı en son eklemek suretiyle, Deşifrele düğmesine tıklayacak kişinin her zaman siz olacağınızı garanti edebilirsiniz.</li>
	<li><strong>Paranoid modu</strong>: Bu modun kullanılması, verilerinizi hem XChaCha20 hem de Serpent ile kademeli bir şekilde şifreler ve verilerin kimliğini doğrulamak için BLAKE2b yerine HMAC-SHA3'ü kullanır. Bu seçenek, çok-gizli dosyaların korunması için önerilir ve gerçekçi olarak olabilecek en yüksek düzeyde güvenlik sağlar. Güçlü bir parola seçtiğinizi varsayarsak, bir bilgisayar korsanının dosyalarınızı deşifre edebilmesi için hem XChaCha20 şifrelemesini hem de Serpent şifrelemesini kırması gerekir. Bu modda dosyalarınızın kırılmasının imkansız olduğunu rahatlıkla söyleyebilirim. </li>
	<li><strong>Reed-Solomon</strong>: Bu özellik eğer dosyalarınızı bir bulut servisinde ya da herhangi bir harci bir birimde uzun süreler boyunca depolamayı planlıyorsanız özellikle kullanışlıdır. İşaretlendiği takdirde, Picocrypt dosya bütünlüğünü korumak için her 128 bayt veri için 8 ekstra bayt eklemek üzere Reed-Solomon hata düzeltme kodunu kullanır. Bu, dosyanızın ~%3 kadarının bozulsa dahi Picocrypt'in yine de hataları düzeltebileceği ve bozulma olmadan dosyalarınızın şifresini çözebileceği anlamına gelir. Elbette, dosyanız çok kötü bir şekilde bozulursa (örneğin, sabit diskinizi düşürdüyseniz), Picocrypt dosyalarınızı tam olarak kurtaramayacaktır, ancak kurtarmak için elinden geleni yapacaktır. Bu seçeneğin şifreleme ve deşifreleme hızlarını yavaşlatabileceğini unutmayın.</li>
	<li><strong>Deşifrelemeye zorla</strong>: Picocrypt, şifre çözme sonrasında dosya bütünlüğünü otomatik olarak kontrol eder. Dosya değiştirilmiş veya bozulmuşsa, Picocrypt kullanıcının güvenliği için deşifre edilen çıktıyı otomatik olarak silecektir. Bu korumaları geçersiz kılmak istiyorsanız, bu seçeneği işaretleyin. Ayrıca, birim şifrelenirken Reed-Solomon özelliği kullanılmışsa ve deşifre ederken bu seçenek işaretlenirse, Picocrypt şifre çözme sırasında dosyanın mümkün olduğunca büyük bir kısmını kurtarmaya çalışacaktır.</li>
	<li><strong>Parçalara böl</strong>: Devasa dosyalarla uğraşmak istemiyor musunuz? Endişelenmeyin! Picocrypt ile çıktı dosyanızı istediğiniz boyutta parçalara ayırmayı seçebilirsiniz, böylece büyük dosyalar daha yönetilebilir bir hale gelir ve bulut sağlayıcılarına yüklemek kolaylaşır. Bir birim (KiB, MiB, GiB veya TiB) seçin ve o birim için istediğiniz parça boyutunu girin. Parçaları deşifrelemek için, bir tanesini Picocrypt'e sürüklemeniz yeterlidir; parçalar deşifreleme sırasında otomatik olarak yeniden birleştirilir.</li>
	<li><strong>Dosyaları sıkıştır</strong>: Varsayılan olarak Picocrypt, birden çok dosyayı şifrelerken dosyaları hızlı bir şekilde birleştirmek için sıkıştırmasız bir zip dosyası kullanır. Ancak bu dosyaları sıkıştırmak isterseniz, bu kutuyu işaretlemeniz yeterlidir; şifreleme sırasında standart Deflate sıkıştırma algoritması uygulanacaktır.</li>
</ul>

# Güvenlik
Picocrypt'in kriptografiyi nasıl ele aldığı hakkında daha fazla bilgi için <a href="Internals.md">Teknik Bilgiler</a>'e göz atabilirsiniz. Benim veya bu projenin güvenliğinden endişe ediyorsanız, sizi temin ederim ki bu depo çalınmayacak veya arka kapı saldırısına uğramayacak. Tüm taşınabilir cihazlarımdaki tam disk şifrelemesine ek olarak, Picocrypt ile bağlantılı tüm hesaplarda (GitHub, Reddit, Google, vb.) 2FA'yı (TOTP) etkinleştirdim. FDahası, Picocrypt ihtiyaç duyduğu bağımlılıkları tarafımca çatallanmış izole depolardan kullanır, yalnızca değişikliklere göz attıktan sonra ve herhangi bir güvenlik sorunu olmadığına emin olduğumda ana depoyu fetch ediyorum. Bu, eğer ilgili bağımlılığın bir saldırıya uğraması veya sahibi tarafından silinmesi durumunda, Picocrypt'in benim çatallarımı kullandığı için hiçbir şekilde etkilenmeyeceği anlamına gelir. Picocrypt'i kullanmakla ilgili kafanız rahat olabilir.

## İmzalar
Paranoid paketi için, Picocrypt PGP ile şifrelenmiştir. Parmak izi ve açık anahtar aşağıda sunulmuştur.

<pre>B342A744BDEEA57B6A583E33A247E73798946F55</pre>
<pre>-----BEGIN PGP PUBLIC KEY BLOCK-----

mDMEYoGUHxYJKwYBBAHaRw8BAQdAvmQA+pdbDB/ynJxHhNDpz6Sb5tgkNuuNJIvw
HYwZtqi0CVBpY29jcnlwdIiTBBMWCgA7FiEEs0KnRL3upXtqWD4zokfnN5iUb1UF
AmKBlB8CGwMFCwkIBwICIgIGFQoJCAsCBBYCAwECHgcCF4AACgkQokfnN5iUb1UZ
RgEA8jbIsdqCr21DWxcqW/eLlbxRkuA8kflVYvWWUxtVqsUA/jQPSDpvA8rakvaL
PIbXjQvrAMkEVIc0HbCzLxr1k3sH
=YFwz
-----END PGP PUBLIC KEY BLOCK-----</pre>

# Topluluk
<a href="https://www.reddit.com/r/Picocrypt/">r/Picocrypt</a>'e katılmayı düşünebilirsiniz. Bu subreddit'te kendim aktif olmayacak olsam dahi, özellikle gelecekte bana veya bu depoya bir şey olması durumunda soru sormak ve başkalarına yardım etmek için harika bir ortamdır. Sadece bu spesifik subreddit'e güvenin ve diğer ortamlarda beni taklit etmeye ve sizi kandırmaya çalışabilecek bilgisayar korsanlarına karşı dikkatli olun. Size asla şifrenizi sormam ve eğer soruyorsam bilin ki o ben değilimdir. Sizden asla şüpheli bir bağlantıdan dosya indirmenizi istemeyeceğim ve eğer istiyorsam bilin ki o ben değilimdir.

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
Neden VeraCrypt, 7-Zip, BitLocker, ya da Cryptomator yerine Picocrypt kullanmalısınız? İşte size Picocrypt kullanmanız için birkaç sebep:
<ul>
	<li>NordLocker, BitLocker, AxCrypt, ve diğer pek çok bulut depolama çözümlerinin aksine, Picocrypt ve bağımlılıkları tamamen açık kaynaklı ve denetime müsaittir. Herhangi bir arka kapı veya kusur olmadığını bizzat kendiniz doğrulayabilirsiniz.</li>
	<li>Picocrypt <i>ufaktır</i>. NordLocker 100MB'ın üzerinde ve VeraCrypt 30MB'tan fazlayken, Picocrypt sadece 3MB, -ki bu da yaklaşık olarak yüksek çözünürlüklü bir görsel kadardır. Ve hepsi bu kadar da değil, Picocrypt taşınabilirdir, yükleme gerektirmez, ve yönetici hakları ya da kök erişim iznine ihtiyaç duymaz.</li>
	<li>Picocrypt, VeraCrypt'a oranla kullanımı daha basit ve daha verimlidir. VeraCrypt ile dosyalarınızı şifreleyebilmek için, bölüm oluştururken hiç olmazsa en azından beş dakika harcamanız gerekir. Picocrypt'in basit kullanıcı arayüzü sayesinde, tüm yapmanız gereken dosyaları sürükleyip bırakmak, bir parola belirlemek ve Başla tuşuna basmak. Picocrypt tüm karmaşık ayarları kendisi halleder.</li>
	<li>Picocrypt güvenlik amacı güdülerek tasarlanmıştır. 7-Zip bir arşiv yazılımı olup asıl odağı şifreleme ya da güvenlik değildir. Lâkin Picocrypt, güvenliği her şeyin üstünde tutacak şekilde oluşturulmuştur. Picocrypt'e ait her bir parça bir sebepten ötürü var ve Picocrypt'in güvenliğine etki edebilecek her şey silinmiştir. Picocrypt güvenebileceğiniz bir kriptografidir.</li>
	<li>Picocrypt dosyaları korumasının yanı sıra onları tasdik te eder, bilgisayar korsanlarının hassas verileri zarara sebep olacak şekilde değiştirmesinin önüne geçer. Bu, güvenli olmayan bir kanal üzerinden şifrelenmiş dosyaları gönderirken ve el değmeden ulaştığından emin olmak istediğinizde kullanışlı bir özellik halini alır. Picocrypt tasdik için güvenli saygınlığı yüksek bir mesaj doğrulama kodu olan Poly1305 kullanır.</li>
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
