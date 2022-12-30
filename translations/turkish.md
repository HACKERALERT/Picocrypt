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

# Bağış
Picocrypt'i aktif olarak geliştirirken bağışları kabul ettim, ancak artık Picocrypt tamamlandı ve kullanıma hazır, artık buna gerek yok. Bunun yerine, zamanınızı ve çabanızı Picocrypt sevgisini başkalarıyla paylaşmak için ayırın. Bağışlar güzel ama başkalarına edilen yardım benim için birkaç dolardan çok daha değerli. Picocrypt'in, insanların dosyalarını korumalarına yardımcı olduğunu bilmek benim için yeterli.

# Sıkça Sorulan Sorular

**Picocrypt yeni özellikler alacak mı?**

Hayır, Picocrypt tam-özellikli olarak kabul edilmekte olup yeni özellikler almayacaktır. Sürekli olarak yeni özellikler eklemeye çalışan diğer araçların aksine (-ki bu yeni hatalar ve yeni güvenlik açıklarına sebep olur), Picocrypt yalnızca birkaç temel özelliğe odaklanır, ancak her birini olağanüstü derecede iyi yapar. Picocrypt'in ideolojisi küçük, basit ve güvenli olmasıdır.

**Android/iOS desteklenecek mi?**

Hayır, geleneksel masaüstü işletim sistemlerinden çok farklı oldukları ve uygulama geliştirmek için farklı araç zincirleri gerektirdikleri için Android veya iOS'u desteklemeyi planlamıyorum. Bununla birlikte, açık kaynaklı yazılımın doğası gereği, gelecekte Android veya iOS için topluluk tarafından oluşturulmuş bir Picocrypt sürümü olabilir.

**Neden Picocrypt sıklıkla güncelleme almaz?**

İnsanlar, amaca uygun ve güvenli kalabilmesi için bir yazılımın sürekli olarak güncellenmesi gerektiği fikrine sahip gibi görünüyor. Bu, bugün kullandığımız birçok yazılım için doğru olsa dahi, Picocrypt için geçerli değildir. Picocrypt "iyi yazılım"dır ve iyi yazılımlar amaca uygun ve güvenli kalmak için sürekli güncelleme almaya ihtiyaç duymazlar. İyi yazılım her zaman iyi yazılım olacaktır.

**"Dosyaları sil" özelliği dosyaları gerçekten siliyor mu (shredding)?**

Hayır, hiçbir dosyayı gerçekten silmez sadece dosya yöneticiniz nasıl siliyorsa o şekilde siler. SSD gibi modern depolama birimlerinde dosyaları gerçekten silme diye bir şey yoktur, zira bu birimlerdeki yıpranma hızı belirli bir sektörün üzerine yazmayı imkansız kılmaktar. Dolayısıyla, Picocrypt kullanıcılara yanlış bir güvenlik hissi vermemek için herhangi bir gerçekten silme özelliği içermez.

# Teşekkürler
Open Collective üzerinden hatrı sayılır miktarda bağışlarda bulunan bu kişilere tüm içtenliğimle teşekkür ederim.
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
<!-- Son güncelleme tarihi 18 Kasım 2022 -->

Ayrıca, Picocrypt'e ilk bağışta bulunan ve destekleyen beş kişiden oluşan aşağıdaki listeye de çok teşekkür ederim:
<ul>
	<li>W.Graham</li>
	<li>N. Chin</li>
	<li>Manjot</li>
	<li>Phil P.</li>
	<li>E. Zahard</li>
</ul>

Ayrıca, Picocrypt'in tercüme edilmesine ve dünya tarafından daha erişilebilir hale getirilmesine yardımcı olan bu kişilere çok teşekkür ederim:
<ul>
	<li>Türkçe için @umitseyhan75</li>
	<li>Almance için @digitalblossom ve @Pokabu26</li>
	<li>Portekizce için @zeeaall ve @Rayserzor</li>
	<li>İspanyolca için u/francirc ve @victorhck</li>
	<li>Rusça için yn ve Voron</li>
	<li>Macarca için @Etim-Orb</li>
	<li>İtalyanca için @Minibus93</li>
	<li>Fransızca için Michel</li>
	<li>Farsça için @MasterKia</li>
	<li>Lehçe için @ungespurv</li>
	<li>Çince için @qaqland</li>
</ul>

Ve son olarak, bu insanlara/organizasyonlara ihtiyacım olduğunda yardım ettikleri için teşekkür ederim:
<ul>
	<li>Picocrypt için bir AppImage oluşturmamda bana yardım ettiği için [ REDACTED ]</li>
	<li>PGP imzalarını test etmemde yardımcı olduğu için u/Upstairs-Fishing867</li>
	<li>Düzenli geri bildirimleri ve desteğinden ötürü u/greenreddits</li>
	<li>Picocrypt'i test etmeme yardımcı olduğu için</li>
	<li>Yaptığı bol miktarda test için u/NSABackdoors</li>
	<li>Geri bilgirim, tavsiye ve desteğinden ötürü @samuel-lucas6</li>
	<li>Picocrypt'i listeledikleri için <a href="https://privacytools.io">PrivacyTools</a></li>
	<li>Picocrypt'i listeledikleri için<a href="https://privacyguides.org">PrivacyGuides</a></li>
	<li>...Siz?</li>
</ul>