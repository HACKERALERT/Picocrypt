<p>English | <a href="/translations/french.md">Français</a> | <a href="/translations/spanish.md">Español</a> | <a href="/translations/german.md">Deutsch</a> | <a href="/translations/portuguese.md">Português</a> | <a href="/translations/turkish.md">Türkçe</a> | <a href="/translations/chinese.md">中文</a> | <a href="/translations/russian.md">русский</a> | <a href="/translations/hungarian.md">Magyar</a> | <a href="/translations/italian.md">Italiano</a></p>
<p align="center"><img align="center" src="/images/logo.svg" width="512" alt="Picocrypt"></p> 

Picocrypt is a very small (hence <i>Pico</i>), very simple, yet very secure encryption tool that you can use to protect your files. It's designed to be the <i>go-to</i> tool for encryption, with a focus on security, simplicity, and reliability. Picocrypt uses the secure XChaCha20 cipher and the Argon2id key derivation function to provide a high level of security, even from three-letter agencies like the NSA. It's designed for maximal security, making absolutely no compromises security-wise, and is built with Go's standard x/crypto modules. <strong>Your privacy and security is under attack. Take it back with confidence by protecting your files with Picocrypt.</strong>

<p align="center"><img align="center" src="/images/screenshot.png" width="318" alt="Picocrypt"></p>

# Funding
Please donate to Picocrypt on <a href="https://opencollective.com/picocrypt">Open Collective</a> (crypto is accepted) to raise money for a potential audit from Cure53. Because this is a project that I spend many hours on and make no money from, I cannot pay for an audit myself. <i>Picocrypt needs support from its community.</i>

# Downloads
**Important**: There's an outdated and useless piece of abandonware called PicoCrypt on the Internet, which was last updated in 2005. PicoCrypt is not related in any way to Picocrypt (this project). Make sure you only download Picocrypt from this repository to ensure that you get the authentic and backdoor-free Picocrypt.

## Windows
Picocrypt for Windows is as simple as it gets. To download the latest, standalone, and portable executable for Windows, click <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.26/Picocrypt.exe">here</a>. If Windows Defender or your antivirus flags Picocrypt as a virus, please do your part and submit it as a false positive for the betterment of everyone.

## macOS
Picocrypt for macOS is very simple as well. Download Picocrypt <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.26/Picocrypt.app.zip">here</a>, extract the zip file, and run Picocrypt which is inside. If you can't open Picocrypt because it's not from a verified developer, right click on Picocrypt and hit "Open". If you still get the warning, right click on Picocrypt and hit "Open" again and you should be able to start Picocrypt.

## Linux
A Snap is available for Linux. Assuming you're on a Debian-based system, a simple `apt install snapd` and `snap install picocrypt` will be enough. For other distros such as Fedora, detailed instructions are available at https://snapcraft.io/picocrypt. Due to the complexity of dependencies and static linking, I don't distribute standalone .deb or .rpm binaries because they would be unreliable and not worth the hassle. Snapcraft manages all dependencies and runtimes automatically and is the recommended way to run Picocrypt on any major Linux distribution. Additionally, Snapcraft provides better security and containerization than Flatpaks and AppImages, which is important for an encryption tool like Picocrypt. If you would prefer not to deal with Canonical, remember that building from source is always an option.

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

# Comparison
Here's how Picocrypt compares to other popular encryption tools.

|                | Picocrypt      | VeraCrypt      | 7-Zip (GUI)    | NordLocker     | BitLocker      | AxCrypt        |
| -------------- | -------------- | -------------- | -------------- | -------------- | -------------- | -------------- |
| Free           |✅ Yes         |✅ Yes          |✅ Yes         |🟧 Partially    |🟧 Partially   |🟧 Partially    |
| Open Source    |✅ GPLv3       |✅ Multi        |✅ LGPL        |❌ No           |❌ No          |❌ No           |
| Cross-Platform |✅ Yes         |✅ Yes          |❌ No          |❌ No           |❌ No          |❌ No           |
| Size           |✅ 2MB         |❌ 20MB         |✅ 2MB         |❌ 60MB         |✅ Included    |🟧 8MB          |
| Portable       |✅ Yes         |✅ Yes          |❌ No          |❌ No           |✅ Yes         |✅ Yes          |
| Permissions    |✅ None        |❌ Admin        |❌ Admin       |❌ Admin        |❌ Admin       |❌ Admin        |
| Ease-Of-Use    |✅ Easy        |❌ Hard         |✅ Easy        |🟧 Medium       |🟧 Medium      |✅ Easy         |
| Key Derivation |✅ Argon2      |🆗 PBKDF2       |❌ SHA256      |✅ Argon2       |❓ Unknown     |🆗 PBKDF2       |
| Data Integrity |✅ Always      |❌ No           |❌ No          |✅ Always       |❓ Unknown     |✅ Always       |
| Reed-Solomon   |✅ Yes         |❌ No           |❌ No          |❌ No           |❌ No          |❌ No           |
| Compression    |✅ Yes         |❌ No           |✅ Yes         |❌ No           |✅ Yes         |✅ Yes          |
| Telemetry      |✅ None        |✅ None         |✅ None        |❌ Analytics    |❓ Unknown     |❌ Accounts     |
| Audited        |🟧 Planned     |✅ Yes          |❌ No          |❓ Unknown      |❓ Unknown     |❌ No           |

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

# Thank You's
A thank you from the bottom of my heart to the people on Open Collective who have made a significant contribution:
<ul>
	<li>YellowNight ($818)</li>
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
	<li>Fuderal on Discord for helping me setup a Discord server</li>
	<li>u/greenreddits for constant feedback and support</li>
	<li>u/Tall_Escape for helping me test Picocrypt</li>
	<li>u/NSABackdoors for doing plenty of testing</li>
	<li>@samuel-lucas6 for feedback, suggestions, and support</li>
	<li><a href="https://privacytools.io">PrivacyToolsIO</a> for listing Picocrypt</li>
	<li><a href="https://privacyguides.org">PrivacyGuides</a> for listing Picocrypt</li>
</ul>
