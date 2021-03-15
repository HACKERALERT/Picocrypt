# Instructions
Picocrypt is written to be cross-platform, so you should be able to run the raw Python file on your OS without any issues. Picocrypt's dependencies will be automatically installed via pip, which should work without any issues. If Picocrypt can't automatically install dependencies, install these dependencies via pip manually: <code>argon2-cffi</code>, <code>pycryptodome</code>, and <code>reedsolo</code>.

# Note about reedsolo
Picocrypt can use the <code>reedsolo</code> pip package itself, but it is very slow because it's pure Python. It is recommended to compile a Python extension (.pyd/.so) for <code>reedsolo</code>, and name it creedsolo (ie. <code>creedsolo.pyd</code> or <code>creedsolo.so</code>). Make sure to include the extension in the same directory as <code>Picocrypt.py</code>. The Windows executable already bundles <code>creedsolo.pyd</code>, but for Linux, you'll have to build the Python extension yourself if you want better speeds. Building the extension is not necessary if you don't intend on using the Reed-Solomon feature, or if you are okay with speeds ~1MB/s.

# External links
Here are the Github links of the dependencies of Picocrypt:

Argon2-cffi: https://github.com/hynek/argon2-cffi

Pycryptodome: https://github.com/Legrandin/pycryptodome

ReedSolo: https://github.com/tomerfiliba/reedsolomon
