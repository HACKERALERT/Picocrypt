# Instructions For Running Manually
If you want to run from source or want to use Picocrypt on any Linux distro, you're in the right place. To run Picocrypt from source, first install Python3 (any version >3.6 will do). Next, install these dependencies via <code>pip</code>: <code>argon2-cffi</code>, <code>pycryptodome</code>, and <code>reedsolo</code>.

# Work in progress, be patient

# A note about reedsolo
Picocrypt can use the raw <code>reedsolo</code> Python module by itself, but it is very slow because it's written in pure Python. It is recommended to compile a Python extension (.pyd/.so) for <code>reedsolo</code>, and name it <code>creedsolo</code> (ie. <code>creedsolo.pyd</code> or <code>creedsolo.so</code>). Make sure to include the extension in the same directory as <code>Picocrypt.py</code>. See <a href="https://github.com/tomerfiliba/reedsolomon#installation" target="_blank">here</a> for instructions on building the extension. The Windows executable already bundles <code>creedsolo.pyd</code>, but for Linux or MacOS, you'll have to build the Python extension yourself if you want better speeds. <strong>Building the extension is not necessary if you don't intend on using the Reed-Solomon feature a lot, or if you are okay with slow Reed-Solomon speeds. Picocrypt will work just fine without the extension and just the way it is.</strong>

<ul>
  <li>Argon2-cffi: https://github.com/hynek/argon2-cffi</li>
  <li>Pycryptodome: https://github.com/Legrandin/pycryptodome</li>
  <li>ReedSolo: https://github.com/tomerfiliba/reedsolomon</li>
</ul>
