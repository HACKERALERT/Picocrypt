# v1.14 (Soon to be released)

# v1.13 (Released 5/29/2021)
<strong>Note: v1.3 will be INCOMPATIBLE with all previous releases! This will likely be the last change in the header format, meaning that all future releases _should_ be compatible with v1.13 and on.</strong>
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
