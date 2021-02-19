# Picocrypt
Picocrypt is a very tiny, very simple, yet very secure file encryption tool. It uses the modern XChaCha20-Poly1305 cipher suite as well as Argon2ID, making it about as secure and modern of an encryption tool as you'll ever get your hands on. Picocrypt's focus is <i>security</i>, so it might be slightly slower than others. 

# Download
You can run the raw Python source file, compile it yourself, or download the portable .exe (for Windows) that I've precompiled and optimized beyond imagination (recommended, because it's just 4MB in size). If you're compiling from source or running the raw Python file, you'll need these two <code>pip</code> dependencies: <code>argon2-cffi</code> and <code>pycryptodome</code>.

# Why Picocrypt?
Why should you use Picocrypt, instead of Bitlocker, NordLocker, VeraCrypt, or 7-Zip? Here are some reasons why you should switch to Picocrypt:

<ul>
	<li>Unlike NordLocker and Bitlocker, Picocrypt is FOSS (free open-source software) and can be audited. You can verify for yourself that there aren't any backdoors.</li>
	<li>Picocrypt is portable and <i>tiny</i> (just 4MB!). It's much lighter than NordLocker (>100MB) and VeraCrypt (>30MB). It can also run on any machine (since it's Python) and the pre-made .exe can run on any Windows PC from 7 and up.</li>
	<li>It's infinitely easier to use than VeraCrypt (no need to create volumes) and a 5-year-old could use Picocrypt.</li>
	<li>Picocrypt is built for security, using modern standards and the most secure settings. See **Security** below for more info.</li>
	<li>It supports file integrity checking through Poly1305, which means that you would know if a hacker has maliciously modified your data.</li>
</ul>

# Instructions
Picocrypt is about as simple as it gets. Just select a file, enter a password, and start. There are some additional options that you can use for more control:

<ul>
	<li>File metadata (editable for encryption, readonly for decryption): Use this to store notes, information, and text along with the file (it won't be encrypted). For example, you can put a description of the file before sending it to someone. When the person you sent it to selects the file in Picocrypt, your description will be shown to that person.</li>
	<li>Keep decrypted output even if it's corrupted or modified (decryption only): Picocrypt automatically checks for integrity upon decryption. If the file has been modified or is corrupted, Picocrypt will delete the output. If you want to keep the corrupt or modified data after decryption, check this option.</li>
	<li>Securely erase and delete original file (encryption only): If checked, Picocrypt will generate pseudo-random data and write it to the original file while encrypting, effectively wiping the original file. The file will be deleted once encryption is complete.</li>
</ul>

# Security
Security is Picocrypt's sole focus. I was in need of a secure, reliable, and future-proof encryption tool that didn't require bloatware and containers, but I couldn't find one, so I created Picocrypt. Picocrypt uses XChaCha20-Poly1305, which is a revision of the eSTREAM winner, Salsa20. XChaCha20-Poly1305 has been through significant amount of cryptanalysis and was selected by security engineers at Google to be used in modern TLS suites. It's considered to be the future of encryption, and makes Picocrypt more secure than Bitlocker, NordLocker, and 7-Zip. For key derivation, Picocrypt uses Argon2ID, winner of the PHC (Password Hashing Competition), which was completed in 2015. Both XChaCha20-Poly1305 and Argon2ID are well recognized within the cryptography community and both are mature and future-proof. Let me get this clear: <i>I did not write the crypto for Picocrypt</i>. Instead, I followed cryptography's number one rule: Don't roll your own crypto. Picocrypt uses two Python libraries, <code>argon2-cffi</code> and <code>pycryptodome</code>, both of which are well known and popular within the Python community.
