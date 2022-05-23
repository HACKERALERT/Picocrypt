<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p> 

پیکوکریپت یک ابزار رمزگذاری بسیار کوچک («پیکو» یعنی <i>چیز خیلی کوچک و ریز</i>)، بسیار ساده و با این حال بسیار امن است که شما میتوانید برای محافظت از فایل هایتان از آن استفاده کنید. این ابزار با تمرکز بر امنیت، سادگی و قابل اتکا بودن، به گونه ای طراحی شده تا <i>بهترین</i> ابزار برای رمزگذاری باشد. پیکوکریپت از سایفر (الگوریتم قفل گذاریِ) امنِ XChaCha20 و از تابع استخراج کلیدِ Argon2id استفاده میکند تا بتواند سطح بالایی از امنیت را فراهم کند (حتی در برابر سازمان های امنیتی مانند سازمان امنیت آمریکا یا همان NSA). این ابزار برای امنیت حداکثری طراحی شده و هیچ قصور و کوتاهی از بابت امنیت انجام نمیدهد و با ماژول های استانداردِ x/crypto از زبان Go ساخته شده است. <strong>حریم خصوصی و امنیت شما تحت حمله است. با محافظت کردن فایل هایتان توسط پیکوکریپت، حریم خصوصی و امنیت خود را با اطمینان خاطر پس بگیرید.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="318" alt="Picocrypt"></p>

# سرمایه گذاری
لطفاً در سایت <a href="https://opencollective.com/picocrypt">Open Collective</a> (رمزارز هم پذیرفته میشود) به پیکوکریپت کمک مالی کنید تا پول مورد نیاز برای انجام یک ممیزی امنیتی روی پیکوکریپت از طرف گروه امنیتی Cure53 جمع آوری شود. چون این پروژه ای است که من ساعت ها وقت میگذارم و درآمدی از آن ندارم، نمیتوانم خودم پول برای یک ممیزی امنیتی پرداخت کنم. <i>پیکوکریپت به حمایت از طرف جامعه کاربرانش نیاز دارد.</i>

# دانلود
**نکته مهم**: یک نرم افزار رها شده، قدیمی و بی مصرف به نام PicoCrypt که آخرین بار در سال 2005 (1384 خورشیدی) بروز شده، در اینترنت وجود دارد. PicoCrypt به هیچ وجه با Picocrypt (این پروژه) ارتباطی ندارد. دقت داشته باشید که پیکوکریپت را فقط از این صفحه دانلود کنید تا بتوانید اطمینان داشته باشید که پیکوکریپتِ اصل و بدون درپشتی (حفره امنیتی) را دریافت کرده اید.

## ویندوز
پیکوکریپت برای ویندوز به ساده ترین شکل ممکن است. برای دانلود نسخه بروز و خوداتکا (Standalone) و قابل حمل پیکوکریپت برای ویندوز <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.28/Picocrypt.exe">اینجا</a> کلیک کنید. اگر ویندوز دیفندر یا آنتی ویروس شما پیکوکریپت را به عنوان ویروس شناسایی میکند، لطفاً لطف کنید و پیکوکریپت را به عنوان مثبت کاذب (False Positive) ثبت کنید تا به نفع همه بشود
.
اگر فایل اجرایی کار نمیکند، احتمالاً به این معناست که سیستم شما از OpenGL (یک نوع استاندارد گرافیکی) پشتیبانی نمیکند. برای حل این مشکل، من یک نسخه دیگر فراهم کرده ام که روی هر سیستم ویندوزی کار میکند که میتوانید از <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.28/Picocrypt-NoGL.exe">اینجا</a> آن را دریافت کنید.

## مکOS
پیکوکریپت برای مکOS هم بسیار ساده است. آن را از <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.28/Picocrypt.app.zip">اینجا</a> دانلود کنید و فایل را از حالت فشرده در بیاورید و پیکوکریپت داخل پوشه را اجرا کنید. اگر به علت اینکه پیکوکریپت از طرف یک توسعه دهنده تایید شده نیست نمیتوانید آن را باز کنید، روی فایل پیکوکریپت کلیک راست کنید و گزینه "Open/باز کردن" را بزنید. اگر هنوز هم اخطار دریافت میکنید، روی فایل پیکوکریپت کلیک راست کنید و گزینه "Open/باز کردن" را دوباره بزنید تا بتوانید پیکوکریپت را راه اندازی کنید.

## گنو/لینوکس
برای استفاده از پیکوکریپت در گنو/لینوکس چند راه وجود دارد. راه پیشنهادی، نصب پیکوکریپت از <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.28/Picocrypt.deb">این</a> فایل deb. میباشد (پشتیبانی شده در +11 Debian و +20 Ubuntu). اگر فایل نصبی deb. نیاز های شما را برآورده نمیکند و یا از یک توزیع برپایه دبیان استفاده نمیکنید، با خیال راحت از <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.28/Picocrypt.AppImage">این</a> فایل AppImage استفاده کنید. اگر هیچکدام از این دو گزینه کار نمیکند، میتوانید پیکوکریپت را از Snapcraft نصب کنید که باید روی همه توزیع ها کار کند. توضیحات مربوط به نصب از Snapcraft را <a href="https://snapcraft.io/picocrypt">اینجا</a> میتوانید پیدا کنید.

## Paranoid Packs
The Paranoid Pack is a compressed archive that contains executables for every version of Picocrypt ever released for Windows, macOS, and Linux. As long as you have it stored in a place you can access, you'll be able to open it and use any version of Picocrypt in case this repository mysteriously vanishes or the entire Internet burns down. Think of it as a seed vault for Picocrypt. As long as one person has the Paranoid Pack within reach, they can share it with the rest of the world and keep Picocrypt functional in cases of catastrophic events like GitHub shutting down suddenly or the NSA capturing me (just in case, you know?). The best way to ensure Picocrypt is accessible many decades from now is to keep a Paranoid Pack in a safe place. So if you are worried about being unable to access Picocrypt in the future, well, here's your solution. Just head to the Releases tab and get yourself a copy.

# Why Picocrypt?
Why should you use Picocrypt instead of BitLocker, NordLocker, VeraCrypt, AxCrypt, or 7-Zip? Here are a few reasons why you should choose Picocrypt:
<ul>
	<li>Unlike NordLocker, BitLocker, AxCrypt, and most cloud storage providers, Picocrypt and its dependencies are completely open-source and auditable. You can verify for yourself that there aren't any backdoors or flaws.</li>
	<li>Picocrypt is <i>tiny</i>. While NordLocker is over 50MB and VeraCrypt is over 20MB, Picocrypt sits at just 2MB, about the size of a medium-resolution photo. And that's not all - Picocrypt is portable (doesn't need to be installed) and doesn't require administrator/root privileges.</li>
	<li>Picocrypt is easier and more productive to use than VeraCrypt. To encrypt files with VeraCrypt, you'd have to spend at least five minutes setting up a volume. With Picocrypt's simple UI, all you have to do is drag and drop your files, enter a password, and hit Start. All the complex procedures are handled by Picocrypt internally. Who said secure encryption can't be simple?</li>
	<li>Picocrypt is designed for security. 7-Zip is an archive utility and not an encryption tool, so its focus is not on security. Picocrypt, however, is built with security as the number one priority. Every part of Picocrypt exists for a reason and anything that could impact the security of Picocrypt is removed. Picocrypt is built with cryptography you can trust.</li>
	<li>Picocrypt authenticates data in addition to protecting it, preventing hackers from maliciously modifying sensitive data. This is useful when you are sending encrypted files over an insecure channel and want to be sure that it arrives untouched.</li>
	<li>Picocrypt actively protects encrypted header data from corruption by adding extra Reed-Solomon parity bytes, so if part of a volume's header (which contains important cryptographic components) corrupts (e.g., hard drive bit rot), Picocrypt can still recover the header and decrypt your data with a high success rate. Picocrypt can also encode the entire volume with Reed-Solomon to prevent any corruption to your important files.</li>
</ul>

# مقایسه
پیکوکریپت در مقایسه با دیگر ابزار های معروف رمزگذاری.

|                | Picocrypt      | VeraCrypt      | 7-Zip (GUI)    | BitLocker      | Cryptomator    | NordLocker     | AxCrypt        |
| -------------- | -------------- | -------------- | -------------- | -------------- | -------------- | -------------- | -------------- |
| رایگان           |✅ بله         |✅ بله          |✅ بله         |🟧 تا حدودی    |✅ بله         |🟧 تا حدودی    |🟧 تا حدودی   |
| آزاد    |✅ GPLv3       |✅ چند لایسنسه        |✅ LGPL        |❌ خیر           |✅ GPLv3       |❌ خیر           |❌ خیر          |
| چند-سکویی |✅ بله         |✅ بله          |❌ خیر          |❌ خیر           |✅ بله         |❌ خیر           |❌ خیر          |
| حجم           |✅ 2MB         |❌ 20MB         |✅ 2MB         |✅ نامشخص     |❌ 50MB        |❌ 60MB         |🟧 8MB         |
| قابل حمل       |✅ بله         |✅ بله          |❌ خیر          |✅ بله          |❌ خیر          |❌ خیر           |✅ بله         |
| دسترسی ها    |✅ هیچ        |❌ ادمین        |❌ ادمین       |❌ ادمین        |❌ ادمین       |❌ ادمین        |❌ ادمین       |
| آسانی استفاده    |✅ آسان        |❌ سخت         |✅ آسان        |🟧 متوسط       |🟧 متوسط      |🟧 متوسط       |✅ آسان        |
| الگوریتم قفل گذاری (سایفر)         |✅ XChaCha20   |✅ AES-256      |✅ AES-256     |🟧 AES-128      |✅ AES-256     |✅ AES-256      |🟧 AES-128     |
| تابع استخراج کلید |✅ Argon2      |🆗 PBKDF2       |❌ SHA256      |❓ نامشخص      |✅ Scrypt      |✅ Argon2       |🆗 PBKDF2      |
| یکپارچگی داده ها |✅ همیشگی      |❌ خیر           |❌ خیر          |❓ نامشخص      |✅ همیشگی      |✅ همیشگی       |✅ همیشگی      |
| تصحیح خطا (رید-سالامون)   |✅ بله         |❌ خیر           |❌ خیر          |❌ خیر           |❌ خیر          |❌ خیر           |❌ خیر          |
| فشرده سازی    |✅ بله         |❌ خیر           |✅ بله         |✅ بله          |❌ خیر          |❌ خیر           |✅ بله         |
| جمع آوری داده      |✅ هیچ        |✅ هیچ         |✅ هیچ        |❓ نامشخص      |✅ هیچ        |❌ داده های آماری    |❌ حساب کاربری    |
| ممیزی امنیتی        |🟧 تحت اقدام     |✅ بله          |❌ خیر          |❓ نامشخص      |✅ بله         |❓ نامشخص      |❌ خیر          |

# Features
Picocrypt is a very simple tool, and most users will intuitively understand how to use it in a few seconds. On a basic level, simply dropping your files, entering a password, and hitting Start is all that's needed to encrypt your files. Pretty simple, right?

While being simple, Picocrypt also strives to be powerful in the hands of knowledgeable and advanced users. Thus, there are some additional options that you may use to suit your needs.
<ul>
	<li><strong>Password generator</strong>: Picocrypt provides a secure password generator that you can use to create cryptographically secure passwords. You can customize the password length, as well as the types of characters to include.</li>
	<li><strong>Comments</strong>: Use this to store notes, information, and text along with the file (it won't be encrypted). For example, you can put a description of the file you're encrypting before sending it to someone. When the person you sent it to drops the file into Picocrypt, your description will be shown to that person.</li>
	<li><strong>Keyfiles</strong>: Picocrypt supports the use of keyfiles as an additional form of authentication (or the only form of authentication). Not only can you use multiple keyfiles, but you can also require the correct order of keyfiles to be present, for a successful decryption to occur. A particularly good use case of multiple keyfiles is creating a shared volume, where each person holds a keyfile, and all of them (and their keyfiles) must be present in order to decrypt the shared volume.</li>
	<li><strong>Paranoid mode</strong>: Using this mode will encrypt your data with both XChaCha20 and Serpent in a cascade fashion, and use HMAC-SHA3 to authenticate data instead of BLAKE2b. This is recommended for protecting top-secret files and provides the highest level of practical security attainable. In order for a hacker to crack your encrypted data, both the XChaCha20 cipher and the Serpent cipher must be broken, assuming you've chosen a good password. It's safe to say that in this mode, your files are impossible to crack.</li>
	<li><strong>Reed-Solomon</strong>: This feature is very useful if you are planning to archive important data on a cloud provider or external medium for a long time. If checked, Picocrypt will use the Reed-Solomon error correction code to add 8 extra bytes for every 128 bytes to prevent file corruption. This means that up to ~3% of your file can corrupt and Picocrypt will still be able to correct the errors and decrypt your files with no corruption. Of course, if your file corrupts very badly (e.g., you dropped your hard drive), Picocrypt won't be able to fully recover your files, but it will try its best to recover what it can. Note that this option will slow down encryption and decryption considerably.</li>
	<li><strong>Force decrypt</strong>: Picocrypt automatically checks for file integrity upon decryption. If the file has been modified or is corrupted, Picocrypt will automatically delete the output for the user's safety. If you would like to override these safeguards, check this option. Also, if this option is checked and the Reed-Solomon feature was used on the encrypted volume, Picocrypt will attempt to recover as much of the file as possible during decryption.</li>
	<li><strong>Split files into chunks</strong>: Don't feel like dealing with gargantuan files? No worries! With Picocrypt, you can choose to split your output file into custom-sized chunks, so large files can become more manageable and easier to upload to cloud providers. Simply choose a unit (KiB, MiB, GiB, or TiB) and enter your desired chunk size for that unit. To decrypt the chunks, simply drag one of them into Picocrypt and the chunks will be automatically recombined during decryption.</li>
</ul>

# Security
For more information on how Picocrypt handles cryptography, see <a href="Internals.md">Internals</a> for the technical details. If you're worried about the safety of me or this project, let me assure you that this repository won't be hijacked or backdoored. I have 2FA (TOTP) enabled on all accounts with a tie to Picocrypt (GitHub, Google, Reddit, Ubuntu One/Snapcraft, Discord, etc.), in addition to full-disk encryption on all of my portable devices. For further hardening, Picocrypt uses my isolated forks of dependencies and I fetch upstream only when I have taken a look at the changes and believe that there aren't any security issues. This means that if a dependency gets hacked or deleted by the author, Picocrypt will be using my fork of it and remain completely unaffected. You can feel confident about using Picocrypt.

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
Here are some places where you can stay up to date with Picocrypt and get involved:
<ul>
	<li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
	<li><a href="https://discord.gg/8QM4A2caxH">Discord</a></li>
</ul>

I highly recommend you join Picocrypt's subreddit because all updates and polls will be posted there. Remember to only trust these social networks and be aware of hackers that might try to impersonate me. I will never ask you for your password, and anyone who does is not me. I will never tell you to download a file from a suspicious link, and anyone who does is not me.

# Stargazers
How's Picocrypt doing? Take a look below to find out.
![Stargazers Over Time](https://starchart.cc/HACKERALERT/Picocrypt.svg)

# Donations
If you find Picocrypt useful, please consider tipping my <a href="https://paypal.me/evanyiwensu">PayPal</a>. I'm providing this software completely free of charge, and would love to have some supporters that will motivate me to continue my work on Picocrypt.

# Thank Yous
A thank you from the bottom of my heart to the people on Open Collective who have made a significant contribution:
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
	<li>@digitalblossom & @Pokabu26 for German</li>
	<li>@zeeaall for Brazilian Portuguese</li>
	<li>@kurpau for Lithuanian</li>
	<li>u/francirc for Spanish</li>
	<li>yn for Russian</li>
	<li>@Etim-Orb for Hungarian</li>
	<li>@Minibus93 for Italian</li>
	<li>Michel for French</li>
	<li>@victorhck for Spanish</li>
</ul>

Finally, thanks to these people/organizations for helping me out when needed:
<ul>
	<li>[ REDACTED ] for helping me create an AppImage for Picocrypt</li>
	<li>u/Upstairs-Fishing867 for helping me test PGP signatures</li>
	<li>Fuderal on Discord for helping me setup a Discord server</li>
	<li>u/greenreddits for constant feedback and support</li>
	<li>u/Tall_Escape for helping me test Picocrypt</li>
	<li>u/NSABackdoors for doing plenty of testing</li>
	<li>@samuel-lucas6 for feedback, suggestions, and support</li>
	<li><a href="https://privacytools.io">PrivacyToolsIO</a> for listing Picocrypt</li>
	<li><a href="https://privacyguides.org">PrivacyGuides</a> for listing Picocrypt</li>
</ul>
