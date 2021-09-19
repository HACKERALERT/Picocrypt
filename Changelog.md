# Features Under Consideration
<ul>
	<li>Get Picocrypt audited (<i>please donate</i> to help fund the audit)</li>
	<li>Hidden volume like VeraCrypt (thanks u/greenreddits)</li>
	<li>CLI (thanks u/your_username)</li>
	<li>File name encryption</li>
</ul>

# v1.XX (a future release)
<ul>
	<li>Associate .pcv file extension with Picocrypt</li>
	<li>Add start button to shredder (improve user flow)</li>
</ul>

# v1.18 (ETA: 1 month)
<ul>
	<li>Add output size to input label</li>
	<li>Customizable Argon2 parameters</li>
	<li>✓ Make UI more consistent (minor DPI issues)</li>
	<li>Improve Reed-Solomon performance</li>
	<li>Remove EXIF data from keyfiles</li>
	<li>✓ Encryption/decryption file naming and extension bugs</li>
</ul>

# v1.17 (ETA: 5 days)
<strong>Note: v1.17 will be incompatible with 1.16 (and therefore 1.15 and 1.14).</strong>
<ul>
	<li>✓ (abandoned due to UI issues with ASCII codes >128) Extended ASCII set in password generator</li>
	<li>✓ Tooltips for all advanced options</li>
	<li>✓ Localization support (use system default where possible)</li>
	<li>✓ Auto detect system locale, fallback to English</li>
	<li>✓ Fix ETA negative number bug</li>
	<li>✓ Add clear button to password field</li>
	<li>✓ Multiple keyfiles support and DND</li>
	<li>✓ Option to require specific keyfile order</li>
	<li>✓ Keyfile generator</li>
	<li>✓ Bug: Red error label shown in main window during successful decryption after selecting incorrect keyfiles</li>
	<li>✓ Prevent duplicate keyfile</li>
	<li>✓ Add a select keyfile button</li>
	<li>✓ Make sure only one of "Fast mode" and "Paranoid mode" can be enabled</li>
	<li>✓ (abandoned because this could cause issues with different language characters) Filter invalid charaters and emojis out of "Save as" text entry</li>
	<li>✓ Fix bug where metadata says "read-only", but the textbox is modifiable</li>
	<li>✓ Add option to delete encrypted files after decryption</li>
</ul>

# v1.16 (Released 08/11/2021)
<ul>
	<li>Fixed bug when entering a wrong password when decrypting a splitted file</li>
	<li>Fixed bug where an existing file is delete when a wrong password is used</li>
	<li>The password generator is now customizable</li>
	<li>Make keyfile support more reliable (keyfile now out of Beta)</li>
	<li>Fix keyfile user flow issue</li>
	<li>Bug fixes</li>
	<li>UI fixes improvements</li>
</ul>

# v1.15 (Released 08/09/2021)
<ul>
	<li>Add cancel button to file shredder and custom number of passes</li>
	<li>Password generator</li>
	<li>Make password strength circle start at top</li>
	<li>Fix shredder UI bugs</li>
</ul>

# v1.14 (Released 08/07/2021)
<strong>v1.14 will unfortunately be incompatible with v1.13, as I have dropped Poly1305 in favour of an HMAC. Because I have appended the encryptor version to the encrypted data in v1.13, however, v1.14 and on will be able to tell you which version your data was encrypted with so that you can fetch the correct verion of Picocrypt. There are no security issues in any previous versions of Picocrypt as long as you use a strong password.</strong>
<ul>
	<li>Low-severity security fix for the recently discovered <a href="https://eprint.iacr.org/2020/1491.pdf">partitioning oracle attacks</a></li>
	<li>Move from Monocypher to Go's standard supplemental ChaCha20 in favour of being the latter being stateful</li>
	<li>Add SHA3 (normal mode) and BLAKE2b (fast mode) as HMAC to replace Poly1305 and prevent partitioning oracle attacks</li>
	<li>Removed ~100 lines of unnecessary code now that Picocrypt uses Go's ChaCha20 (cleaner and stabler code)</li>
	<li>Added window icons</li>
	<li>Switch to a new Reed-Solomon encoder that automatically corrects errors</li>
	<li>Add a "Paranoid mode", which will use the Serpent cipher in addition to XChaCha20</li>
	<li>Cleaner code with plenty of comments for people taking a look</li>
	<li>Metadata is now Reed-Solomon encoded (everything bit of header data is now RS-encoded for redundancy)</li>
	<li>Reed-Solomon checkbox is now enabled and Reed-Solomon works</li>
	<li>Implemented Dropbox's zxcvbn password strength checker</li>
	<li>Removed paranoid shredding as it is too hard to implement correctly and not cross platform</li>
	<li>Fixed Windows zip extract error notice that doesn't appear in 7-Zip (edit: it was a backslash and forwardslash issue)</li>
	<li>Optional shred temporary files checkbox</li>
	<li>Remove BLAKE3 from the checksum generator tab, as it has no practical use and requires a non-standard library</li>
	<li>Advanced options are shown dynamically depending on whether encrypting or decrypting</li>
	<li>Window closing disabled during encryption/decryption/shredding to prevent leakage of temporary files</li>
	<li>Reduce padding of metadataLength from 10 to 5 (you probably won't type more than 99999 metadata characters)</li>
	<li>Use regex to check if an input file is a valid Picocrypt volume or not (during decryption)</li>
	<li>Improved user flow as well as fix UI bugs</li>
	<li>Code optimizations</li>
	<li>Many bug fixes/stability improvments</li>
</ul>

# v1.13 (Released 5/29/2021)
<strong>Note: v1.13 will be INCOMPATIBLE with all previous releases! <s>This will likely be the last change in the header format, meaning that all future releases _should_ be compatible with v1.13 and on.</s></strong>
<ul>
	<li>Picocrypt has been ported from Python to Go, thus completely rewritten</li>
	<li>Added fast mode, which can achieve ~250MB/s</li>
	<li>Added file shredder and file checksum generator</li>
	<li>Automatically checks for newer versions</li>
	<li>Added file chunking support</li>
</ul>

# v1.12.1 (Bug fix patch, released 04/11/2021)
There was a major bug in v1.12 that caused the "Secure wipe" feature to show "Unknown error" when done. This wasn't anything serious security-wise and v1.12.1 has the fix for this bug. Also, a bug that causes "Secure wipe" to hang has been fixed.

# v1.12 (Released 04/07/2021)
<ul>
	<li>Beautiful UI</li>
	<li>More than x2 as fast as previous versions</li>
	<li>Add cancel button to cancel encryption/decryption</li>
	<li>(Bug) Delete existing file only if password is correct</li>
	<li>Minor aesthetic fixes</li>
	<li>Complete rewrite from scratch, to ensure reliability and security</li>
	<li>Better anti-corruption (re-defined header format)</li>
	<li>Switch to Argon2d instead Argon2id for better security</li>
	<li>Switch from SHA3 to BLAKE3 for corruption check</li>
	<li>Better user flow</li>
</ul>
<strong>Note: v1.12 will be INCOMPATIBLE with all previous releases!</strong>

# v1.11 (released 03/23/2021)
<ul>
	<li>Much more secure wipe via <code>sdelete64</code> for Windows, <code>shred</code> for Linux, and <code>rm -P</code> for MacOS</li>
	<li>Much more beautiful UI for MacOS</li>
	<li>Robust secure wipe support for drag and dropped files/folders</li>
	<li>Only open input files in read mode, since write mode is unnecessary</li>
	<li>Clean up source code, add better comments</li>
	<li><strong>New: </strong>Drag and drop support (multiple files, a folder, a file and a folder, etc.)</li>
</ul>
