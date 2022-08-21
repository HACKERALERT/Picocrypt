<p>English | <a href="/translations/french.md">Français</a> | <a href="/translations/spanish.md">Español</a> | <a href="/translations/german.md">Deutsch</a> | <a href="/translations/portuguese.md">Português</a> | <a href="/translations/turkish.md">Türkçe</a> | <a href="/translations/chinese.md">中文</a> | <a href="/translations/russian.md">русский</a> | <a href="/translations/hungarian.md">Magyar</a> | <a href="/translations/italian.md">Italiano</a> | <a href="/translations/persian.md">پارسی</a> | <a href="/translations/polish.md">Polski</a></p>
<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p>

Picocrypt 是一个非常小（因而叫 _Pico_）、简单、而且安全的加密工具，可以用来保护文件。它为成为加密的 _首选_ 工具而设计，着重于安全性、简易性和可靠性。Picocrypt 使用 XChaCha20 安全算法和 Argon2id 密钥推导函数，以提供高水平的、甚至能抵挡 NSA 攻击的安全性。使用 Go 的标准库 x/crypto 而构建，Picocrypt 为最大程度的安全而设计，在安全方面绝不做任何妥协。**你的隐私和安全正受到攻击，快使用 Picocrypt 夺回本属于我们的数据！奥利给！**

<p align="center"><img align="center" src="/images/screenshot.png" width="318" alt="Picocrypt"></p>

# 基金

请在 [Open Collective](https://opencollective.com/picocrypt) 上向 Picocrypt 捐赠（接受加密货币），Picocrypt 需要为 Cure53 的安全审计筹集资金。我为本项目花了很多时间，而且没有盈利，我无法单独支付审计费用。_Picocrypt 需要来自社区的支持。_

# 下载

**注意**：“Picocrypt”这个名字下有多个独立条目。例如，有一个旧的加密工具叫 PicoCrypt，它使用算法已被破解。还有一个 ERC 资助的研究项目叫 PICOCRYPT。甚至还有 Picocrypt 相关的域名，都不是我注册的。请不要将这些无关的项目与 Picocrypt（本项目）混淆。务必只从这个仓库下载 Picocrypt，以确保你得到可信的、无后门的 Picocrypt。当与他人分享 Picocrypt 时，一定请使用此仓库的链接地址，以防止任何混淆。

## Windows

适用于 Windows 的 Picocrypt 非常简单。我们提供了最新的、单文件、绿色版的 Windows 可执行文件，下载请点击[这里](https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.exe)。如果 Windows Defender 或杀毒软件将 Picocrypt 标记为病毒，请尽自己的一份力量并将其作为误报提交，以造福所有人。

如果上面的可执行文件不能工作，这很可能意味着您的系统不支持 OpenGL。对于这种情况，我提供了一个可以在任何 Windows 系统上运行的替代实现，你可以在[这里](https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt-NoGL.exe)下载。

## macOS

用于 macOS 的 Picocrypt 也非常简单。在[这里](https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.app.zip)下载，解压，并运行里面的 Picocrypt。如果提示 Picocrypt 不是来自一个经过验证的开发商，打不开，请按住 control 点击 Picocrypt 并点“打开”以绕过警告。

需要记住，Picocrypt 在 macOS 上依赖 Rosetta 2 和 OpenGL，而这两者未来可能会被苹果移除，因为苹果公司正在推广他们的专有图像和芯片。这意味着 Picocrypt 可能将来在 macOS 上无法运行。解决办法？不要购买苹果。他们从不在意开发者的处境。

## Linux

有多种方法可以在 Linux 上使用 Picocrypt。推荐通过[这里](https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.deb)的 .deb 包 安装 Picocrypt（Debian 11+ 和 Ubuntu 20+）。如果 .deb 不适合，或者您不使用基于 Debian 的发行版，可以使用[这里](https://github.com/HACKERALERT/Picocrypt/releases/download/1.29/Picocrypt.AppImage)的 AppImage。如果以上两种方法都不奏效，你可以从 Snapcraft 安装 Picocrypt，它应该能在所有发行版上工作。在[这里](https://snapcraft.io/picocrypt)找到 Snapcraft 的说明。

## Paranoid Packs

Paranoid Pack 是一个压缩档案，它包含了 Picocrypt 为 Windows、macOS 和 Linux 发布的每个版本的可执行文件。只要你把它存放在一个你可以访问的地方，当这个存储库神秘地消失或整个互联网被烧毁时，你就可以打开它并使用任何版本的 Picocrypt。它可以看作是 Picocrypt 的种子库。如果发生了 GitHub 突然关闭或 NSA 拿人等灾难性事件（以防万一，懂的都懂），只要有一个人有 Paranoid Pack，他就能和世界上的其他人共享，维持 Picocrypt 的功能。确保 Picocrypt 在几十年后仍可使用的最好方法，就是在安全的地方保存一个 Paranoid Pack。因此，如果担心未来无法访问 Picocrypt，这就是您的解决方案。只需前往 Release 菜单，自己获得一份副本。

# 为什么选择 Picocrypt？

为什么您应该使用 Picocrypt 而不是 BitLocker、NordLocker、VeraCrypt、AxCrypt 或 7-Zip？有以下几个原因：

- 与 NordLocker、BitLocker、AxCrypt 和大多数云存储供应商不同，Picocrypt 及其依赖关系是完全开源和可审计的。你可以自己验证是否有任何后门或缺口。
- Picocrypt 很 _小_。NordLocker 超过 50 兆，VeraCrypt 超过 20 兆，但是 Picocrypt 只有 2 兆，这体积大约相当于一张中等分辨率的照片。而且 Picocrypt 是绿色的（不需要安装），不需要管理员/root 权限。
- Picocrypt 比 VeraCrypt 更简单、高效。用 VeraCrypt 加密文件，你至少要花 5 分钟来设置一个卷。通过 Picocrypt 简洁的 UI ，你所要做的就是拖放你的文件，输入密码，然后点击开始。所有复杂的程序都由 Picocrypt 内部处理。加密难道不能简单吗？
- Picocrypt 专为安全设计。7-Zip 是一个存档工具，而不是一个加密工具，所以它的重点不在安全上。然而，Picocrypt 是以安全为首要考虑的。Picocrypt 的每个部分都有其存在的理由，任何可能影响 Picocrypt 安全的东西都被移除。Picocrypt 是用你可以信任的密码学构建的。
- Picocrypt 除了保护数据外，还对数据进行认证，防止黑客恶意修改敏感数据。当您在不安全的通道上发送加密文件，并希望确保其到达时未被触及时，这一点很有用。
- Picocrypt 通过增加额外的 Reed-Solomon 奇偶校验字节来积极地保护头数据不被破坏。因此如果卷头（包含重要的加密暗号组件）的部分被破坏（比如硬盘坏道），Picocrypt 仍然可以恢复卷头并以高成功率解密您的数据。Picocrypt 还可以用 Reed-Solomon 对整个卷进行编码，防止您重要文件的任何损坏。

# 对比

Picocrypt 与其他主流加密工具对比。

|              | Picocrypt    | VeraCrypt  | 7-Zip (GUI) | BitLocker    | Cryptomator | NordLocker   | AxCrypt      |
| :----------: | :----------- | :--------- | :---------- | :----------- | :---------- | :----------- | :----------- |
|     免费     | ✅ Yes       | ✅ Yes     | ✅ Yes      | 🟧 Partially | ✅ Yes      | 🟧 Partially | 🟧 Partially |
|     开源     | ✅ GPLv3     | ✅ Multi   | ✅ LGPL     | ❌ No        | ✅ GPLv3    | ❌ No        | ❌ No        |
|    跨平台    | ✅ Yes       | ✅ Yes     | ❌ No       | ❌ No        | ✅ Yes      | ❌ No        | ❌ No        |
|     体积     | ✅ 2 MiB     | ❌ 20 MiB  | ✅ 2 MiB    | ✅ Included  | ❌ 50 MiB   | ❌ 60 MiB    | 🟧 8 MiB     |
|    便携性    | ✅ Yes       | ✅ Yes     | ❌ No       | ✅ Yes       | ❌ No       | ❌ No        | ✅ Yes       |
|   权限需求   | ✅ None      | ❌ Admin   | ❌ Admin    | ❌ Admin     | ❌ Admin    | ❌ Admin     | ❌ Admin     |
|     交互     | ✅ Easy      | ❌ Hard    | ✅ Easy     | 🟧 Medium    | 🟧 Medium   | 🟧 Medium    | ✅ Easy      |
|     算法     | ✅ XChaCha20 | ✅ AES-256 | ✅ AES-256  | 🟧 AES-128   | ✅ AES-256  | ✅ AES-256   | 🟧 AES-128   |
|   密钥推导   | ✅ Argon2    | 🆗 PBKDF2  | ❌ SHA-256  | ❓ Unknown   | ✅ Scrypt   | ✅ Argon2    | 🆗 PBKDF2    |
|  数据完整性  | ✅ Always    | ❌ No      | ❌ No       | ❓ Unknown   | ✅ Always   | ✅ Always    | ✅ Always    |
| Reed-Solomon | ✅ Yes       | ❌ No      | ❌ No       | ❌ No        | ❌ No       | ❌ No        | ❌ No        |
|     压缩     | ✅ Yes       | ❌ No      | ✅ Yes      | ✅ Yes       | ❌ No       | ❌ No        | ✅ Yes       |
|     遥测     | ✅ None      | ✅ None    | ✅ None     | ❓ Unknown   | ✅ None     | ❌ Analytics | ❌ Accounts  |
|   安全审计   | 🟧 Planned   | ✅ Yes     | ❌ No       | ❓ Unknown   | ✅ Yes      | ❓ Unknown   | ❌ No        |

# 特点

Picocrypt 是一个非常简单的工具，几秒钟内，大多数用户就能直观地懂得如何使用它。基础使用，只需拖拽你的文件，输入密码，然后点击开始，就可以对你的文件进行加密。很简单，不是嘛？

在简单的同时，Picocrypt 也力求在专业的高级用户手中变得强大。因此，有一些额外的选项，您可以根据自己的需要来使用。

- **密码生成器**：Picocrypt 提供了一个安全的密码生成器，您可以用它来创建加密所需的密码。您可以自定义密码的长度，以及密码包含的字符类型。
- **备注**：使用它可以将注释、信息和文本与文件一起存储（它不会被加密）。例如，你可以在寄件前，写上你要加密的文件的描述。当收件人将加密后的文件放入 Picocrypt 时，您的描述将被显示出来。
- **多密钥文件**：Picocrypt 支持使用多密钥文件作为一种额外的认证形式（或唯一的认证形式）。您不仅可以使用多个密钥文件，还可以要求所需密钥文件的顺序。多密钥文件的一个特别好的用例是创建共享卷，每个人都持有一个密钥文件，所有的人（和他们的密钥文件）都必须在场才能解密共享卷。
- **偏执模式**：这种模式将以层叠方式用 XChaCha20 和 Serpent 加密你的数据，并使用 HMAC-SHA3 替代 BLAKE2b 来验证数据。它被推荐用于保护最高机密的文件，提供实际可达到的最高安全。假设你的密码很棒，那么对于黑客来说，要破解你的加密数据，XChaCha20 算法和 Serpent 算法两者都必须被破解。可以说在这种模式下，你的文件是无法破解的。
- **Reed-Solomon**：如果您打算将重要的数据长期归档在云提供商或外部介质上，这个功能非常有用。如果选中，Picocrypt 将使用 Reed-Solomon 纠错码，为每 128 字节的数据增加额外 8 个字节，以防止文件损坏。这意味着您的文件最多有 ~3% 的损坏时，Picocrypt 仍然能够纠正错误并正确解密您的文件。当然，如果您的文件损坏得非常严重（例如，硬盘摔了），Picocrypt 将不能完全恢复您的文件，但它会尽力恢复它能恢复的部分。请注意，这个选项可能会减慢加密和解密的速度。
- **强制解密**：Picocrypt 在解密时自动检查文件的完整性。如果文件被修改或损坏，为了用户安全，Picocrypt 将自动删除已输出部分。如果您想改变该保障措施，请选中此选项。另外，如果该选项被选中并且在加密卷上使用了 Reed-Solomon 功能，Picocrypt 将试图在解密过程中尽可能多地恢复文件的内容。
- **文件分块**：不喜欢处理巨大的文件？不用担心！使用 Picocrypt，您可以选择将您的输出分割成自定义大小的文件块，这样大文件可以更容易管理，更容易上传到云服务提供商。只需选择一个单位（KiB、MiB、GiB 或 TiB），并输入您想要的分块大小。要解密这些分块，只需将其中一个块拖入 Picocrypt，在解密过程中，这些分块将会自动重新组合起来。

# 安全性

关于 Picocrypt 密码学的更多信息，请参阅 [Internals](Internals.md) 了解技术细节。如果你担心我或这个项目的安全，让我向你保证，这个仓库不会被劫持或留后门。我在所有与 Picocrypt 有关的账户（GitHub、Google、Reddit、Ubuntu One/Snapcraft、Discord 等）都启用了 2FA（TOTP），此外，我还在所有的便携设备上进行了全盘加密。为了进一步加固，Picocrypt 使用的依赖库都是 fork，只有当我认为没有任何安全问题时，我才会从上游获取同步。这意味着，如果依赖库被黑客攻击或被作者删除，Picocrypt 使用的 fork 都完全不受影响。你可以放心地使用 Picocrypt。

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

# 社区

考虑加入官方的 subreddit（[r/Picocrypt](https://www.reddit.com/r/Picocrypt/)）。虽然我本人不会在 subreddit 上活跃，但它仍然是一个提问和互相帮助的好地方，特别是如果我或这个软件库将来发生什么事的时候。记住只能相信 subreddit，注意黑客可能试图在其他平台冒充我。我永远不会索要你的密码，任何这样做的人都不是我。我永远不会让你从一个可疑的链接下载文件，任何这样做的人都不是我。

# Stargazers

How's Picocrypt doing? Take a look below to find out.
![Stargazers Over Time](https://starchart.cc/HACKERALERT/Picocrypt.svg)

# 捐赠

当 Picocrypt 在积极开发时，我接受捐赠，但现在 Picocrypt 已经完成并可以生产环境使用，就没有必要了。作为替代，你可以花点时间和精力与他人分享来自 Picocrypt 的快乐。捐赠是极好的，但对我来说，能够帮助别人比一些闲钱更有价值。看！Picocrypt 正在帮助人们保护他们的文件，这对我来说已经足够了。

# FAQ

**Picocrypt 是否接受新功能？**

不，Picocrypt 功能完整，不会添加任何新功能。与其他试图不断增加新功能的工具不同（这会引入新的 bug 和安全漏洞），Picocrypt 只关注少数核心功能，然后把每项功能都做得非常好。记住 Picocrypt 的理念：小、简单、安全。

**会支持 Android/iOS 吗？**

不会，我没有计划，因为它们与传统的桌面操作系统有很大不同，需要不同的工具链来开发。然而，由于开源，未来可能会出现社区构建的 Android 或 iOS 版本。

**为什么 Picocrypt 不经常更新？**

人们似乎有这样的意识：软件必须不断地更新以保持其相关性和安全性。虽然这对我们今天使用的很多软件来说可能是对的，但 Picocrypt 不是这样。Picocrypt 是 “优秀的软件”，优秀的软件不需要不断的更新来保持相关性和安全性。优秀，驷马难求。

# Thank Yous

A thank you from the bottom of my heart to the people on Open Collective who have made a significant contribution:

<ul>
	<li>Guest ($842)</li>
	<li>YellowNight ($818)</li>
	<li>evelian ($50)</li>
	<li>jp26 ($50)</li>
	<li>guest-116103ad ($50)</li>
	<li>oli ($20)</li>
	<li>Markus ($15)</li>
	<li>Tybbs ($10)</li>
	<li>N. Chin ($10)</li>
	<li>Manjot ($10)</li>
	<li>Phil P. ($10)</li>
	<li>Raymond ($10)</li>
	<li>donor39 (backer)</li>
	<li>Pokabu (backer)</li>
</ul>

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
</ul>
