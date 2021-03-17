#!/usr/bin/env python3

"""
Dependencies: argon2-cffi, pycryptodome, reedsolo
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt
"""

# Test if libraries are installed
try:
	from argon2.low_level import hash_secret_raw
	from Crypto.Cipher import ChaCha20_Poly1305
	try:
		from creedsolo import ReedSolomonError
	except:
		from reedsolo import ReedSolomonError
except:
	# Libraries missing, install them
	from os import system
	try:
		# Debian/Ubuntu based
		system("sudo apt-get install python3-tk")
	except:
		# Fedora
		system("sudo dnf install python3-tkinter")

	system("python3 -m pip install argon2-cffi --no-cache-dir")
	system("python3 -m pip install pycryptodome --no-cache-dir")
	system("python3 -m pip install reedsolo --no-cache-dir")

# Imports
from tkinter import filedialog,messagebox
from threading import Thread
from datetime import datetime
from argon2.low_level import hash_secret_raw,Type
from Crypto.Cipher import ChaCha20_Poly1305
from hashlib import sha3_512
from secrets import compare_digest
from os import urandom,fsync,remove
from os.path import getsize,expanduser
import sys
import tkinter
import tkinter.ttk
import tkinter.scrolledtext
import webbrowser
try:
	from creedsolo import RSCodec,ReedSolomonError
except:
	from reedsolo import RSCodec,ReedSolomonError

# Tk/Tcl is a little barbaric, so I'm disabling
# high DPI so it doesn't scale bad and look horrible
try:
	from ctypes import windll
	windll.shcore.SetProcessDpiAwareness(0)
except:
	pass

# Global variables and notices
inputFile = ""
outputFile = ""
password = ""
ad = ""
kept = False
working = False
adString = "File metadata (used to store some text along with the file):"
passwordNotice = "Error. The provided password is incorrect."
corruptedNotice = "Error. The input file is corrupted."
veryCorruptedNotice = "Error. The input file and header keys are badly corrupted."
modifiedNotice = "Error. The input file has been intentionally modified."
kCorruptedNotice = "The input file is corrupted, but the output has been kept."
kModifiedNotice = "The input file has been intentionally modified, but the output has been kept."
kVeryCorruptedNotice = "The input file is badly corrupted, but the output has been kept."
derivingNotice = "Deriving key (takes a few seconds)..."
keepNotice = "Keep decrypted output even if it's corrupted or modified"
eraseNotice = "Securely erase and delete original file"
overwriteNotice = "Output file already exists. Would you like to overwrite it?"
rsNotice = "Prevent corruption using Reed-Solomon"
rscNotice = "Creating Reed-Solomon tables..."
unknownErrorNotice = "Unknown error occured. Please try again."

# Create root Tk
tk = tkinter.Tk()
tk.geometry("480x440")
tk.title("Picocrypt")
tk.configure(background="#f5f6f7")
tk.resizable(0,0)

# Try setting window icon if included with Picocrypt
try:
	favicon = tkinter.PhotoImage(file="./key.png")
	tk.iconphoto(False,favicon)
except:
	pass

# Some styling
s = tkinter.ttk.Style()
s.configure("TCheckbutton",background="#f5f6f7")

# Event when user selects an input file
def inputSelected():
	global inputFile,working
	dummy.focus()

	# Try to handle when select file is cancelled
	try:
		# Ask for input file
		suffix = ""
		tmp = filedialog.askopenfilename(
			initialdir=expanduser("~")
		)
		if len(tmp)==0:
			# Exception will be caught by except below
			raise Exception("No file selected.")
		inputFile = tmp
		# Decide if encrypting or decrypting (".pcf" is the legacy Picocrypt extension,
		#     ".pcv" is the newer Picocrypt extension. Both are cross-compatible, but
		#     I just think ".pcv" is better because it stands for "Picocrypt Volume")
		if ".pcf" in inputFile.split("/")[-1] or ".pcv" in inputFile.split("/")[-1]:
			suffix = " (will decrypt)"
			fin = open(inputFile,"rb+")
			# Read file metadata
			adlen = b""
			while True:
				letter = fin.read(1)
				if letter!=b"+":
					adlen += letter
				if letter==b"|":
					adlen = adlen[:-1]
					break
			ad = fin.read(int(adlen.decode("utf-8")))
			fin.close()
			# Insert the metadata into its text box
			adArea["state"] = "normal"
			adArea.delete("1.0",tkinter.END)
			adArea.insert("1.0",ad.decode("utf-8"))
			adArea["state"] = "disabled"
			adLabelString.set("File metadata (read only):")
			keepBtn["state"] = "normal"
			eraseBtn["state"] = "disabled"
			rsBtn["state"] = "disabled"
		else:
			# Update the UI
			eraseBtn["state"] = "normal"
			keepBtn["state"] = "disabled"
			rsBtn["state"] = "normal"
			adArea["state"] = "normal"
			adArea.delete("1.0",tkinter.END)
			suffix = " (will encrypt)"
			adLabelString.set(adString)
		# Enable password box, etc.
		inputString.set(inputFile.split("/")[-1]+suffix)
		passwordInput["state"] = "normal"
		passwordInput.delete(0,"end")
		startBtn["state"] = "normal"
		statusString.set("Ready.")
		progress["value"] = 0
	# File decode error
	except UnicodeDecodeError:
		passwordInput["state"] = "normal"
		passwordInput.delete(0,"end")
		statusString.set(corruptedNotice)
	# No file selected, do nothing
	except:
		pass
	# Focus the dummy button to remove ugly borders
	finally:
		dummy.focus()
		working = False

# Button to select input file
selectFileInput = tkinter.ttk.Button(
	tk,
	text="Select file",
	command=inputSelected,
)
selectFileInput.place(x=19,y=20)

# Label that displays selected input file
inputString = tkinter.StringVar(tk)
inputString.set("Please select a file.")
selectedInput = tkinter.ttk.Label(
	tk,
	textvariable=inputString
)
selectedInput.config(background="#f5f6f7")
selectedInput.place(x=104,y=23)

# Label that prompts user to enter a password
passwordString = tkinter.StringVar(tk)
passwordString.set("Password:")
passwordLabel = tkinter.ttk.Label(
	tk,
	textvariable=passwordString
)
passwordLabel.place(x=17,y=56)
passwordLabel.config(background="#f5f6f7")

# A frame to make password input fill width
passwordFrame = tkinter.Frame(
	tk,
	width=440,
	height=22
)
passwordFrame.place(x=20,y=76)
passwordFrame.columnconfigure(0,weight=10)
passwordFrame.grid_propagate(False)
# Password input box
passwordInput = tkinter.ttk.Entry(
	passwordFrame
)
passwordInput.grid(sticky="nesw")
passwordInput["state"] = "disabled"

# Start the encryption/decryption process
def start():
	global inputFile,outputFile,password,ad,kept,working
	dummy.focus()
	reedsolo = False
	chunkSize = 2**20

	# Decide if encrypting or decrypting
	if ".pcf" not in inputFile and ".pcv" not in inputFile:
		mode = "encrypt"
		outputFile = inputFile+".pcv"
		reedsolo = rs.get()==1
	else:
		mode = "decrypt"
		# Check if Reed-Solomon was enabled by checking for "+"
		test = open(inputFile,"rb+")
		decider = test.read(1).decode("utf-8")
		test.close()
		if decider=="+":
			reedsolo = True
		# Decrypted output is just input file without the extension
		outputFile = inputFile[:-4]

	# Check if file already exists (getsize() throws error if file not found)
	try:
		getsize(outputFile)
		force = messagebox.askyesno("Warning",overwriteNotice)
		dummy.focus()
		if force!=1:
			return
	except:
		pass

	# Set progress bar indeterminate
	progress.config(mode="indeterminate")
	progress.start(15)

	# Create Reed-Solomon object
	if reedsolo:
		statusString.set(rscNotice)
		# 13 bytes per 128 bytes, ~10% larger output file
		rsc = RSCodec(13)

	# Set and get some variables
	working = True
	headerBroken = False
	reedsoloFixedCount = 0
	reedsoloErrorCount = 0
	dummy.focus()
	password = passwordInput.get().encode("utf-8")
	ad = adArea.get("1.0",tkinter.END).encode("utf-8")
	wipe = erase.get()==1

	# Disable inputs and buttons while encrypting/decrypting
	selectFileInput["state"] = "disabled"
	passwordInput["state"] = "disabled"
	adArea["state"] = "disabled"
	startBtn["state"] = "disabled"
	eraseBtn["state"] = "disabled"
	keepBtn["state"] = "disabled"
	rsBtn["state"] = "disabled"

	# Open files
	fin = open(inputFile,"rb+")
	if reedsolo and mode=="decrypt":
		# Move pointer one forward
		fin.read(1)
	fout = open(outputFile,"wb+")
	if reedsolo and mode=="encrypt":
		# Signal that Reed-Solomon was enabled with "+"
		fout.write(b"+")

	# Generate values for encryption if encrypting
	if mode=="encrypt":
		salt = urandom(16)
		nonce = urandom(24)
		fout.write(str(len(ad)).encode("utf-8")) # Length of metadata
		fout.write(b"|") # Separator
		fout.write(ad) # Metadata (associated data)

		# Write zeros as placeholder, come back to write over it later
		# Note that 13 additional bytes are added if Reed-Solomon is enabled
		fout.write(b"0"*(64+(13 if reedsolo else 0))) # SHA3-512 of encryption key
		fout.write(b"0"*(64+(13 if reedsolo else 0))) # CRC of file
		fout.write(b"0"*(16+(13 if reedsolo else 0))) # Poly1305 tag
		# If Reed-Solomon is enabled, encode the salt and nonce, otherwise write them raw
		fout.write(bytes(rsc.encode(salt)) if reedsolo else salt) # Argon2 salt
		fout.write(bytes(rsc.encode(nonce)) if reedsolo else nonce) # ChaCha20 nonce
	# If decrypting, read values from file
	else:
		# Read past metadata into actual data
		adlen = b""
		while True:
			letter = fin.read(1)
			adlen += letter
			if letter==b"|":
				adlen = adlen[:-1]
				break
		fin.read(int(adlen.decode("utf-8")))
		# Read the salt, nonce, etc.
		# Read 13 extra bytes if Reed-Solomon is enabled
		cs = fin.read(77 if reedsolo else 64)
		crccs = fin.read(77 if reedsolo else 64)
		digest = fin.read(29 if reedsolo else 16)
		salt = fin.read(29 if reedsolo else 16)
		nonce = fin.read(37 if reedsolo else 24)
		# If Reed-Solomon is enabled, decode each value
		if reedsolo:
			try:
				cs = bytes(rsc.decode(cs)[0])
			except:
				headerBroken = True
				cs = cs[:64]
			try:
				crccs = bytes(rsc.decode(crccs)[0])
			except:
				headerBroken = True
				crccs = crccs[:64]
			try:
				digest = bytes(rsc.decode(digest)[0])
			except:
				headerBroken = True
				digest = digest[:16]
			try:
				salt = bytes(rsc.decode(salt)[0])
			except:
				headerBroken = True
				salt = salt[:16]
			try:
				nonce = bytes(rsc.decode(nonce)[0])
			except:
				headerBroken = True
				nonce = nonce[:24]

			if headerBroken:
				if keep.get()!=1:
					statusString.set(veryCorruptedNotice)
					fin.close()
					fout.close()
					remove(outputFile)
					# Reset UI
					selectFileInput["state"] = "normal"
					passwordInput["state"] = "normal"
					adArea["state"] = "normal"
					startBtn["state"] = "normal"
					keepBtn["state"] = "normal"
					working = False
					progress.stop()
					progress.config(mode="determinate")
					progress["value"] = 100
					return
				else:
					kept = "badlyCorrupted"

	# Show notice about key derivation
	statusString.set(derivingNotice)

	# Derive argon2id key
	key = hash_secret_raw(
		password,
		salt,
		time_cost=8, # 8 iterations
		memory_cost=2**20, # 2^20 Kibibytes (1GiB)
		parallelism=8, # 8 parallel threads
		hash_len=32,
		type=Type.ID
	)

	# Key deriving done, set progress bar determinate
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0

	# Compute hash of derived key
	check = sha3_512(key).digest()

	# If decrypting, check if key is correct
	if mode=="decrypt":
		# If key is incorrect...
		if not compare_digest(check,cs):
			if not headerBroken:
				statusString.set(passwordNotice)
				fin.close()
				fout.close()
				remove(outputFile)
				# Reset UI
				selectFileInput["state"] = "normal"
				passwordInput["state"] = "normal"
				adArea["state"] = "normal"
				startBtn["state"] = "normal"
				keepBtn["state"] = "normal"
				working = False
				progress["value"] = 100
				del key
				return

	# Create XChaCha20-Poly1305 object
	cipher = ChaCha20_Poly1305.new(key=key,nonce=nonce)
	# Cyclic redundancy check for file corruption
	crc = sha3_512()

	# Amount of data encrypted/decrypted, total file size, starting time
	done = 0
	total = getsize(inputFile)

	# If secure wipe enabled, create a wiper object
	if wipe:
		wiper = open(inputFile,"r+b")
		wiper.seek(0)

	# Keep track of time because it flies...
	startTime = datetime.now()
	previousTime = datetime.now()

	# Continously read file in chunks of 1MB
	while True:
		if mode=="decrypt" and reedsolo:
			# Read a chunk plus Reed-Solomon recovery bytes
			piece = fin.read(1104905)
		else:
			piece = fin.read(chunkSize)
		if wipe:
			# If securely wipe, write random trash
			# to original file after reading it
			trash = urandom(len(piece))
			wiper.write(trash)
		# If EOF
		if not piece:
			if mode=="encrypt":
				# Get the cipher MAC tag (Poly1305)
				digest = cipher.digest()
				fout.flush()
				fout.close()
				fout = open(outputFile,"r+b")
				# Compute the offset and seek to it
				rsOffset = 1 if reedsolo else 0
				fout.seek(len(str(len(ad)))+1+len(ad)+rsOffset)
				# Write hash of key, CRC, and Poly1305 MAC tag
				# Reed-Solomon-encode if selected by user
				if reedsolo:
					fout.write(bytes(rsc.encode(check)))
					fout.write(bytes(rsc.encode(crc.digest())))
					fout.write(bytes(rsc.encode(digest)))
				else:
					fout.write(check)
					fout.write(crc.digest())
					fout.write(digest)
			else:
				# If decrypting, verify CRC
				crcdg = crc.digest()
				if not compare_digest(crccs,crcdg):
					# File is corrupted
					statusString.set(corruptedNotice)
					progress["value"] = 100
					fin.close()
					fout.close()
					# If keep file not checked...
					if keep.get()!=1:
						remove(outputFile)
						# Reset UI
						selectFileInput["state"] = "normal"
						passwordInput["state"] = "normal"
						adArea["state"] = "normal"
						startBtn["state"] = "normal"
						keepBtn["state"] = "normal"
						working = False
						del fin,fout,cipher,key
						return
					else:
						if not kept:
							kept = "corrupted"
				# Next, verify MAC tag (Poly1305)
				try:
					# Throws ValueError if incorrect Poly1305
					cipher.verify(digest)
				except:
					if not reedsoloErrorCount and not headerBroken:
						# File is modified
						statusString.set(modifiedNotice)
						progress["value"] = 100
						fin.close()
						fout.close()
						# If keep file not checked...
						if keep.get()!=1:
							remove(outputFile)
							# Reset UI
							selectFileInput["state"] = "normal"
							passwordInput["state"] = "normal"
							adArea["state"] = "normal"
							startBtn["state"] = "normal"
							keepBtn["state"] = "normal"
							working = False
							del fin,fout,cipher,key
							return
						else:
							if not kept:
								kept = "modified"					
			break
		
		# Encrypt/decrypt chunk and update CRC
		if mode=="encrypt":
			# Encrypt piece
			data = cipher.encrypt(piece)
			# Update checksum
			crc.update(data)
			if reedsolo:
				# Encode using Reed-Solomon if user chooses
				data = bytes(rsc.encode(data))
		else:
			# Basically encrypting but in reverse
			if reedsolo:
				try:
					data,_,fixed = rsc.decode(piece)
				except ReedSolomonError:
					# File is really corrupted
					if not reedsoloErrorCount:
						if keep.get()!=1:
							statusString.set(veryCorruptedNotice)
							progress["value"] = 100
					# If keep file not checked...
					if keep.get()!=1:
						fin.close()
						fout.close()
						remove(outputFile)
						# Reset UI
						selectFileInput["state"] = "normal"
						passwordInput["state"] = "normal"
						adArea["state"] = "normal"
						startBtn["state"] = "normal"
						keepBtn["state"] = "normal"
						working = False
						progress["value"] = 100
						del fin,fout,cipher,key
						return
					else:
						kept = "badlyCorrupted"
						# Attempt to recover badly corrupted data
						data = b""
						piece = piece[:-13]
						counter = 0
						while True:
							# Basically just strip the Reed-Solomon bytes
							# and return the original non-encoded data
							if counter<1104905:
								data += piece[counter:counter+242]
								counter += 255 # 255 bytes, 242 original
							else:
								break
						fixed = bytearray()
						reedsoloErrorCount += 1
				data = bytes(data)
				reedsoloFixedCount += len(fixed)
				crc.update(data)
				data = cipher.decrypt(data)
			else:
				crc.update(piece)
				data = cipher.decrypt(piece)

		# Calculate speed, ETA, etc.
		first = False
		elapsed = (datetime.now()-previousTime).total_seconds()
		sinceStart = (datetime.now()-startTime).total_seconds()
		previousTime = datetime.now()
		# Prevent divison by zero
		if not elapsed:
			elapsed = 0.1**6
		percent = done*100/total
		progress["value"] = percent
		rPercent = round(percent)
		speed = (done/sinceStart)/10**6
		# Prevent divison by zero
		if not speed:
			first = True
			speed = 0.1**6
		rSpeed = str(round(speed,2))
		# Right-pad zeros to prevent layout shifts
		while len(rSpeed.split(".")[1])!=2:
			rSpeed += "0"
		eta = round((total-done)/(speed*10**6))
		# Seconds to minutes if seconds more than 59
		if eta>=60:
			eta = f"{eta//60}m {eta%60}"
		if isinstance(eta,int) or isinstance(eta,float):
			if eta<0:
				eta = 0
		# If it's the first round and no data/predictions yet...
		if first:
			statusString.set("...% at ... MB/s (ETA: ...s)")
		else:
			# Update status
			info = f"{rPercent}% at {rSpeed} MB/s (ETA: {eta}s)"
			if reedsolo and mode=="decrypt" and reedsoloFixedCount:
				eng = "s" if reedsoloFixedCount!=1 else ""
				info += f", fixed {reedsoloFixedCount} corrupted byte{eng}"
			if reedsolo and mode=="decrypt" and reedsoloErrorCount:
				info += f", {reedsoloErrorCount} MB unrecoverable"
			statusString.set(info)
		
		# Increase done and write to output
		done += 1104905 if (reedsolo and mode=="decrypt") else chunkSize
		fout.write(data)

	# Show appropriate notice if file corrupted or modified
	if not kept:
		if mode=="encrypt":
			output = inputFile.split("/")[-1]+".pcv"
		else:
			output = inputFile.split("/")[-1].replace(".pcf","").replace(".pcv","")
		statusString.set(f"Completed. (Output: {output})")
		# Show Reed-Solomon stats if it fixed corrupted bytes
		if mode=="decrypt" and reedsolo and reedsoloFixedCount:
			statusString.set(f"Completed with {reedsoloFixedCount} bytes fixed."+
				f" (Output: {output})")
	else:
		if kept=="modified":
			statusString.set(kModifiedNotice)
		elif kept=="corrupted":
			statusString.set(kCorruptedNotice)
		else:
			statusString.set(kVeryCorruptedNotice)
	
	# Reset variables and UI states
	selectFileInput["state"] = "normal"
	adArea["state"] = "normal"
	adArea.delete("1.0",tkinter.END)
	adArea["state"] = "disabled"
	startBtn["state"] = "disabled"
	passwordInput["state"] = "normal"
	passwordInput.delete(0,"end")
	passwordInput["state"] = "disabled"
	progress["value"] = 0
	inputString.set("Please select a file.")
	keepBtn["state"] = "normal"
	keep.set(0)
	keepBtn["state"] = "disabled"
	eraseBtn["state"] = "normal"
	erase.set(0)
	eraseBtn["state"] = "disabled"
	rs.set(0)
	rsBtn["state"] = "disabled"
	if not kept:
		fout.flush()
		fsync(fout.fileno())
	fout.close()
	fin.close()
	if wipe:
		# Make sure to flush file
		wiper.flush()
		fsync(wiper.fileno())
		wiper.close()
		remove(inputFile)
	inputFile = ""
	outputFile = ""
	password = ""
	ad = ""
	kept = False
	working = False
	# Wipe keys for safety
	del fin,fout,cipher,key

# Wraps the start() function with error handling
def wrapper():
	global working
	# Try start() and handle errors
	try:
		start()
	except:
		progress["value"] = 100
		selectFileInput["state"] = "normal"
		passwordInput["state"] = "normal"
		adArea["state"] = "normal"
		startBtn["state"] = "normal"
		keepBtn["state"] = "normal"
		rsBtn["state"] = "normal"
		statusString.set(unknownErrorNotice)
		dummy.focus()
		working = False
	finally:
		sys.exit(0)

# Encryption/decrypt is done is a separate thread
# so the UI isn't blocked. This is a wrapper
# to spawn a thread and start it.
def startWorker():
	thread = Thread(target=wrapper,daemon=True)
	thread.start()

# ad stands for "associated data"/metadata
adLabelString = tkinter.StringVar(tk)
adLabelString.set(adString)
adLabel = tkinter.ttk.Label(
	tk,
	textvariable=adLabelString
)
adLabel.place(x=17,y=108)
adLabel.config(background="#f5f6f7")

# Frame so metadata text box can fill width
adFrame = tkinter.Frame(
	tk,
	width=440,
	height=100
)
adFrame.place(x=20,y=128)
adFrame.columnconfigure(0,weight=10)
adFrame.grid_propagate(False)

# Metadata text box
adArea = tkinter.Text(
	adFrame,
	exportselection=0
)
adArea.config(font=("Consolas",12))
adArea.grid(sticky="we")
adArea["state"] = "disabled"

# Check box for keeping corrupted/modified output
keep = tkinter.IntVar()
keepBtn = tkinter.ttk.Checkbutton(
	tk,
	text=keepNotice,
	variable=keep,
	onvalue=1,
	offvalue=0,
	command=lambda:dummy.focus()
)
keepBtn.place(x=18,y=240)
keepBtn["state"] = "disabled"

# Check box for securely erasing original file
erase = tkinter.IntVar()
eraseBtn = tkinter.ttk.Checkbutton(
	tk,
	text=eraseNotice,
	variable=erase,
	onvalue=1,
	offvalue=0,
	command=lambda:dummy.focus()
)
eraseBtn.place(x=18,y=260)
eraseBtn["state"] = "disabled"

# Check box for Reed Solomon
rs = tkinter.IntVar()
rsBtn = tkinter.ttk.Checkbutton(
	tk,
	text=rsNotice,
	variable=rs,
	onvalue=1,
	offvalue=0,
	command=lambda:dummy.focus()
)
rsBtn.place(x=18,y=280)
rsBtn["state"] = "disabled"

# Frame so start button can fill width
startFrame = tkinter.Frame(
	tk,
	width=442,
	height=25
)
startFrame.place(x=19,y=310)
startFrame.columnconfigure(0,weight=10)
startFrame.grid_propagate(False)
# Start button
startBtn = tkinter.ttk.Button(
	startFrame,
	text="Start",
	command=startWorker
)
startBtn.grid(sticky="nesw")
startBtn["state"] = "disabled"

# Progress bar
progress = tkinter.ttk.Progressbar(
	tk,
	orient=tkinter.HORIZONTAL,
	length=440,
	mode="determinate"
)
progress.place(x=20,y=348)

# Status label
statusString = tkinter.StringVar(tk)
statusString.set("Ready.")
status = tkinter.ttk.Label(
	tk,
	textvariable=statusString
)
status.place(x=17,y=376)
status.config(background="#f5f6f7")

# Credits :)
hint = "Created by Evan Su. Click for details and source."
creditsString = tkinter.StringVar(tk)
creditsString.set(hint)
credits = tkinter.ttk.Label(
	tk,
	textvariable=creditsString,
	cursor="hand2"
)
credits["state"] = "disabled"
credits.config(background="#f5f6f7")
credits.place(x=17,y=406)
source = "https://github.com/HACKERALERT/Picocrypt"
credits.bind("<Button-1>",lambda e:webbrowser.open(source))

# Version
versionString = tkinter.StringVar(tk)
versionString.set("v1.9")
version = tkinter.ttk.Label(
	tk,
	textvariable=versionString
)
version["state"] = "disabled"
version.config(background="#f5f6f7")
version.place(x=436,y=406)

# Dummy button to remove focus from other buttons
# and prevent ugly border highlighting
dummy = tkinter.ttk.Button(
	tk
)
dummy.place(x=480,y=0)

# Close window only if not encrypting or decrypting
def onClose():
	if not working:
		tk.destroy()

# Main tkinter loop
if __name__=="__main__":
	tk.protocol("WM_DELETE_WINDOW",onClose)
	tk.mainloop()
	sys.exit(0)
