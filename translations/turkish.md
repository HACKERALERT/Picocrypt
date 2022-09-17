<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p>

Picocrypt; dosyalarınızı şifreleyebilmenizi, sağlama toplamları (checksums) oluşturabilmenizi, dosyalarınızı kalıcı olarak silebilmenizi ve daha fazlasını yapabilmenizi sağlayan oldukça küçük (bu yüzden <i>Pico</i>), oldukça basit ancak oldukça güvenli bir kriptografi aracıdır. Güvenlik, basitlik ve güvenilirliğe odaklanarak şifreleme için <i>ilk-tercih</i> bir araç olması için tasarlanmıştır. Picocrypt, NSA gibi üç harfli teşkilatlara karşı bile üst düzey güvenlik sağlamak için XChaCha20 şifrelemesi ve HMAC-SHA3 mesaj doğrulama kodunu kullanır. Olabilecek en üst düzey güvenlik için tasarlanmıştır, güvenlik açısından mutlak suretle taviz vermez ve denetlenmiş bir kriptografi kütüphanesi ile oluşturulmuştur. <strong>Gizliliğiniz ve güvenliğiniz saldırı altında. Dosyalarınızı Picocrypt ile koruyor olmanın verdiği güvenle buna bir son verin.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="318" alt="Picocrypt"></p>

# Finansman Sağlama
Cure53 tarafından olası bir denetleme yaptırabilmem için gereken parayı toplayabilmem için lütfen Picocrypt'e <a href="https://opencollective.com/picocrypt">Open Collective</a> üzerinden bağışta bulunun. Bu projeyi hayata geçirebilmek için saatlerce çalıştım ve bundan herhangi bir parasal gelir elde etmiyorum, bu yüzden denetleme ücretini kendi başıma ödeyemem. <i>Picocrypt'in, topluluğunun desteğine ihtiyacı var.</i>

# İndirmeler
<strong>Önemli:</strong> İnternette en son 2005'te güncellenmiş PicoCrypt adlı eski, işe yaramaz ve terk edilmiş bir yazılım müsvettesi var. PicoCrypt'in, Picocrypt(bu proje) ile zerre alakası yoktur. Gerçek ve arka kapı içermeyen Picocrypt'i edindiğinizden emin olmak için özellikle bu depodan indirdiğinizden emin olun.

## Windows
Windows için Picocrypt olabildiğince basittir. Windows için en güncel ve taşınabilir uygulamayı indirmek için <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.exe">buraya</a> tıklayın. Bir ihtimal Picocrypt anti-virüs yazılımınızı tetikleyebilir. Picocrypt bir virüs değildir, bu yüzden lütfen gelecekte bu tür sorunların tekrar yaşanmaması için anti-virüs sağlayıcınıza bu durumun bir yanlış pozitif olduğunu belirtin. Daha eski sürümler, Sürümler sekmesinden erişilebilir.

## macOS
<a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.app.zip">buraya</a> tıklayın.

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
