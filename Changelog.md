# v1.12 (soon to be released)
<ul>
	<li>Add cancel button to cancel encryption/decryption</li>
	<li>(Bug) Delete existing file only if password is correct</li>
	<li>Minor aesthetic fixes</li>
	<li>Complete rewrite from scratch, implementing the audited Monocypher library instead of Pycryptodome, which is not audited.</li>
</ul>
<strong>Note: v1.12 will be INCOMPATIBLE with all previous releases!</strong>

# v1.11 (released 3/23/2021)
<ul>
	<li>Much more secure wipe via <code>sdelete64</code> for Windows, <code>shred</code> for Linux, and <code>rm -P</code> for MacOS</li>
	<li>Much more beautiful UI for MacOS</li>
	<li>Robust secure wipe support for drag and dropped files/folders</li>
	<li>Only open input files in read mode, since write mode is unnecessary</li>
	<li>Clean up source code, add better comments</li>
	<li><strong>New: </strong>Drag and drop support (multiple files, a folder, a file and a folder, etc.)</li>
</ul>
