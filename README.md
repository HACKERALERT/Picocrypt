<p>English | <a href="/translations/french.md">Français</a> | <a href="/translations/spanish.md">Español</a> | <a href="/translations/german.md">Deutsch</a> | <a href="/translations/turkish.md">Türkçe</a> | <a href="/translations/chinese.md">中文</a> | <a href="/translations/russian.md">русский</a></p>
<p align="center"><img align="center" src="/images/Picocrypt.svg" width="512" alt="Picocrypt"></p>

Picocrypt is a very small (hence <i>Pico</i>), very simple, yet very secure encryption tool that you can use to protect your files, generate checksums, shred files, and much more. It's designed to be the <i>go-to</i> tool for encryption, with a focus on security, simplicity, and reliability. Picocrypt uses the secure XChaCha20 cipher and the SHA3 hash function to provide a high level of security, even from three-letter agencies like the NSA. It's designed for maximal security, making absolutely no compromises security-wise, and is built with Go's standard x/crypto modules. <strong>Your privacy and security are under attack. Take it back with confidence by protecting your files with Picocrypt.</strong>

<p align="center"><img align="center" src="/images/Picocrypt.png" width="384" alt="Picocrypt"></p>

# Funding
Please donate to Picocrypt on <a href="https://opencollective.com/picocrypt">Open Collective</a> to raise money for a potential audit from Cure53. Because this is a project that I spend many hours on and make no money from, I cannot pay for an audit myself. <i>Picocrypt needs support from its community.</i>

# Downloads
<strong>Important:</strong> There's an outdated and useless piece of abandonware called PicoCrypt on the Internet, which was last updated in 2005. PicoCrypt is not related in any way to Picocrypt (this project). Make sure you only download Picocrypt from the official website or this repository to ensure that you get the authentic and backdoor-free Picocrypt.

## Windows
Picocrypt for Windows is as simple as it gets. To download the latest, standalone, and portable executable for Windows, click <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.16/Picocrypt.exe">here</a>. Older releases are available under the Releases tab.

## macOS
Picocrypt for macOS is very simple as well. Download Picocrypt <a href="https://github.com/HACKERALERT/Picocrypt/releases/download/1.16/Picocrypt.app.zip">here</a>, extract the zip file, and run Picocrypt which is inside.

## Linux
A Snap is available for Linux. Make sure you have Snapcraft installed (`sudo apt install snapd`) and install Picocrypt: `sudo snap install picocrypt`.

# Why Picocrypt?
Why should you use Picocrypt instead of BitLocker, NordLocker, VeraCrypt, AxCrypt, or 7-Zip? Here are a few reasons why you should choose Picocrypt:
<ul>
	<li>Unlike NordLocker, BitLocker, AxCrypt, and most cloud storage providers, Picocrypt and its dependencies are completely open-source and auditable. You can verify for yourself that there aren't any backdoors or flaws.</li>
	<li>Picocrypt is <i>tiny</i>. While NordLocker is over 100MB and VeraCrypt is over 30MB, Picocrypt sits at just 3MB, about the size of a high-resolution image. And that's not all - Picocrypt is portable (doesn't need to be installed) and doesn't require administrator/root privileges.</li>
	<li>Picocrypt is easier and more productive to use than VeraCrypt. To encrypt files with VeraCrypt, you'd have to spend at least five minutes setting up a volume. With Picocrypt's simple UI, all you have to do is drag and drop your files, enter a password, and hit Start. All the complex settings are handled by Picocrypt internally.</li>
	<li>Picocrypt is designed for security. 7-Zip is an archive utility and not an encryption tool, so its focus is not on security. Picocrypt, however, is built with security as the number one priority. Every part of Picocrypt is there for a reason and anything that could impact the security of Picocrypt is removed. Picocrypt is cryptography you can trust.</li>
	<li>Picocrypt authenticates data in addition to protecting it, preventing hackers from maliciously modifying sensitive data. This is useful when you are sending encrypted files over an insecure channel and want to be sure that it arrives untouched. Picocrypt uses HMAC-SHA3 for authenticity, which is a secure and respected hash in a well-known construction.</li>
	<li>Picocrypt actively prevents protects your files from corruption by adding extra Reed-Solomon bits, so if your encrypted data corrupts (e.g., hard drive failure), Picocrypt can still recover your data with a high success rate.</li>
</ul>

# Instructions
Picocrypt is a very simple tool, and most users will intuitively understand how to use it in a few seconds. There are some advanced options, however, that you may use.
<ul>
	<li><strong>File metadata</strong>: Use this to store notes, information, and text along with the file (it won't be encrypted). For example, you can put a description of the file before sending it to someone. When the person you sent it to drops the file into Picocrypt, your description will be shown to that person.</li>
	<li><strong>Keep decrypted output even if it's corrupted or modified</strong>: Picocrypt automatically checks for integrity upon decryption. If the file has been modified or is corrupted, Picocrypt will automatically delete the output for the user's safety. If you want to keep the corrupted or modified data after decryption, check this option. Also, if this option is checked and the Reed-Solomon feature was used on the encrypted file, Picocrypt will attempt to recover as much of the file as possible during decryption.</li>
	<li><strong>Securely shred the original file(s) and folder(s)</strong>: If checked, Picocrypt will use system internals to shred the input file(s). This method is more secure than just deleting the original file because the original file, in that case, can still be recovered by hackers using special software. Picocrypt securely deletes the original file making sure that it's impossible to retrieve the original file(s) after being shredded.</li>
	<li><strong>Prevent corruption using Reed-Solomon</strong>: This feature is very useful if you are planning to archive important data on a cloud provider or external hard drive for a long time. If checked, Picocrypt will use the Reed-Solomon error correction code to add 6 extra bytes for every 128 bytes to prevent file corruption. This means that up to ~3% of your file can corrupt and Picocrypt will still be able to correct the errors and decrypt your files with no corruption. Of course, if your file corrupts very badly (e.g., you dropped your hard drive), Picocrypt won't be able to fully recover your files, but it will try its best to recover what it can. Note that this option will slow down encryption and decryption a good bit.</li>
</ul>

# Security
For more information on how Picocrypt handles cryptography, see <a href="Internals.md">Internals</a> for the technical details. If you're worried about the safety of me or this project, let me assure you that this repository won't be hijacked or backdoored. I have 2FA (TOTP) enabled on all accounts with a tie to Picocrypt (GitHub, Google, Reddit, Discord, etc.), in addition to full-disk encryption on all of my portable devices. For further hardening, Picocrypt uses my "offline" forks of dependencies and I fetch upstream only when I have taken a look at the changes and believe that there aren't any security issues. You can feel confident about using Picocrypt.

# Roadmap
See <a href="Roadmap.md">Roadmap</a> for a list of new features and improvements that I'm considering.

# Community
Here's some places where you can see the latest news about Picocrypt and get involved:
<ul>
	<li><a href="https://www.reddit.com/r/Picocrypt/">Reddit</a></li>
	<li><a href="https://discord.gg/8QM4A2caxH">Discord</a></li>
</ul>
Only trust these social networks and be aware of scammers and hackers that might try to impersonate me and get you to download malicious software.

# Stargazers
[![Stargazers over time](https://starchart.cc/HACKERALERT/Picocrypt.svg)](https://starchart.cc/HACKERALERT/Picocrypt)

# Contribution
If you find a bug or have a feature request, please create an Issue. I'm open to collaborators, and all Issues and PRs are appreciated. If you find a security issue, please follow the instructions <a href="Security.md">here</a> to report it.

If you are multilingual and you know a language in which Picocrypt hasn't been translated into yet, I would love to have your help translating this page as well as Picocrypt's interface. The more languages, the more welcoming!

# Donations
If you find Picocrypt useful, please consider supporting me on <a href="https://patreon.com/evansu">Patreon</a>. As I'm providing this software completely free of charge, I would love to have some supporters that will motivate me to continue my work on Picocrypt.

# Thank You's
A thank you from the bottom of my heart to my fellow supporters on Patreon:
<ul>
	<li>Frederick Doe</li>
</ul>

And to people on Open Collective who made a significant contribution:
<ul>
	<li>jp26 ($50)</li>
</ul>
You are the people that inspire me to work on Picocrypt and provide it free of charge to everyone!

Also a huge thanks to the following list of five people, who were the first people to donate and support Picocrypt:
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
	<li>@digitalblossom for German</li>
</ul>

Finally, thanks to these people for helping me out when needed:
<ul>
	<li>Fuderal on Discord for helping me setup a Discord server</li>
	<li>u/greenreddits for constant feedback and support</li>
	<li>u/Tall_Escape for helping me test Picocrypt</li>
</ul>
