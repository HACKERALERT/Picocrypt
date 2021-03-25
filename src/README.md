# Running from source
If you want to run from source or would like to use Picocrypt on any platform that supports Python, you're in the right place. To run Picocrypt from source, first make sure you have Python3 installed (any version >3.6 will do). Then, download <code>Picocrypt.py</code> from above (you can download a zip file of the files above from the homepage). Now, install these dependencies via <code>pip</code>: <code>argon2-cffi</code>, <code>pycryptodome</code>, <code>reedsolo</code>, and <code>cython</code>. Make sure tkinter is installed on your platform. On Linux, you might have to <code>sudo apt-get install python3-tk</code>.

Next, download <code>tkinterdnd2.zip</code> from above. Extract it and go into it. Inside, you'll see two Python files and a directory called <code>tkdnd</code>.

Now, find the directory where Python is installed. On Windows, for example, this folder would probably be here:
<code>C:\Users\(Your username)\AppData\Local\Programs\Python\PythonXX</code>. Go to there.

Now, go into <code>Lib/</code> and then <code>site-packages</code>. Create a folder named <code>tkinterdnd2</code>. Go into there and copy <code>&#95;&#95;init&#95;&#95;.py</code> and <code>TkinterDnD.py</code> from the zip you extracted earlier to here.

Next, go back to the directory where Python is installed. Go into <code>tcl/</code> and then <code>tclX.X</code>. Go into the same folder you extracted earlier and go into <code>tkdnd/</code>, and then go into your respective platform folder. Copy all the files there into the <code>tclX.X</code> directory from above.

Finally, download <code>reedsolo.zip</code> from above, extract it, and go into the folder. Open up a terminal and <code>python setup.py install</code>. When that's finished, go into <code>Build/</code> and then <code>lib.XXXXXXX</code>. You should now have a Python extension called <code>creedsolo.pyd</code> on Windows, <code>creedsolo.so</code> or <code>creedsolo.dylib</code> on MacOS, and <code>creedsolo.so</code> on Linux. Copy this file to the same directory where <code>Picocrypt.py</code> is.

Now you're all good. Just <code>python3 Picocrypt.py</code> and enjoy!

# External links
<ul>
  <li>Argon2-cffi: https://github.com/hynek/argon2-cffi</li>
  <li>Pycryptodome: https://github.com/Legrandin/pycryptodome</li>
  <li>ReedSolo: https://github.com/tomerfiliba/reedsolomon</li>
  <li>Tkinterdnd2: https://github.com/pmgagne/tkinterdnd2</li>
</ul>
