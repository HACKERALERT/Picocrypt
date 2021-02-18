# Picocrypt
Picocrypt is a very tiny, simple, and secure file encryption tool. It uses the modern XChaCha20-Poly1305 cipher suite as well as Argon2ID, making it about as secure and modern of an encryption tool as you'll get. 

# Download
You can run the raw Python source file, compile it yourself, or download the portable .exe (Windows) that I've precompiled and optimized beyond imagination (recommended, just 4MB in size).

# Instructions
Picocrypt is about as simple as it gets. Just select a file, enter a password, and start. There are some additional options that you can use:

<ul>
	<li>File metadata (changeable for encryption, readonly for decryption): Use this to store notes, information, and text along with the file (it won't be encrypted). For example, you can put a description of the file before sending it to someone. When the person you sent it to selects the file in Picocrypt, your description will be shown to that person.</li>
	<li>Keep decrypted output even if it's corrupted or modified (decryption only): Picocrypt automatically checks for integrity upon decryption. If the file has been modified or corrupted, Picocrypt will delete the output. If you want to keep the corrupt or modified data after decryption, check this option.</li>
	<li>Securely erase and delete original file (encryption only): If checked, Picocrypt will generate pseudo-random data and write it to the original file while encrypting, effectively wiping the original file. The file will be deleted once encryption is complete.</li>
</ul>
