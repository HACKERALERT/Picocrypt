#!/usr/bin/env python3

# Dependencies: argon2-cffi, pycryptodome
# Copyright (c) Evan Su (https://evansu.cc)
# Released under a GNU GPL v3 license
# https://github.com/HACKERALERT/Picocrypt

# Test if libraries are installed
try:
	from argon2.low_level import hash_secret_raw
	from Crypto.Cipher import ChaCha20_Poly1305
except:
	# Libraries missing, install them
	from os import system
	system("sudo apt-get install python3-tk")
	system("python3 -m pip install argon2-cffi")
	system("python3 -m pip install pycryptodome")

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
modifiedNotice = "Error. The input file has been intentionally modified."
kCorruptedNotice = "The input file is corrupted, but the output has been kept."
kModifiedNotice = "The input file has been intentionally modified, but the output has been kept."
derivingNotice = "Deriving key (takes a few seconds)..."
keepNotice = "Keep decrypted output even if it's corrupted or modified"
eraseNotice = "Securely erase and delete original file"
overwriteNotice = "Output file already exists. Would you like to overwrite it?"
unknownErrorNotice = "Unknown error occured. Please try again."

# Create root Tk
tk = tkinter.Tk()
#tk.tk.call('tk', 'scaling', 2.0)
tk.geometry("480x420")
tk.title("Picocrypt")
tk.configure(background="#f5f6f7")
tk.resizable(0,0)

# Try setting image if included with Picocrypt
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
		# Decide if encrypting or decrypting
		if ".pcf" in inputFile.split("/")[-1]:
			suffix = " (will be decrypted)"
			fin = open(inputFile,"rb+")
			# Read file metadata
			adlen = b""
			while True:
				letter = fin.read(1)
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
		else:
			# Update the UI
			eraseBtn["state"] = "normal"
			keepBtn["state"] = "disabled"
			adArea["state"] = "normal"
			adArea.delete("1.0",tkinter.END)
			suffix = " (will be encrypted)"
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

	# Decide if encrypting or decrypting
	if ".pcf" not in inputFile:
		mode = "encrypt"
		outputFile = inputFile+".pcf"
	else:
		mode = "decrypt"
		outputFile = inputFile[:-4]

	# Check if file already exists
	try:
		getsize(outputFile)
		force = messagebox.askyesno("Warning",overwriteNotice)
		dummy.focus()
		if force!=1:
			return
	except:
		pass

	# Set and get some variables
	working = True
	dummy.focus()
	password = passwordInput.get().encode("utf-8")
	ad = adArea.get("1.0",tkinter.END).encode("utf-8")
	wipe = erase.get()==1

	selectFileInput["state"] = "disabled"
	passwordInput["state"] = "disabled"
	adArea["state"] = "disabled"
	startBtn["state"] = "disabled"
	keepBtn["state"] = "disabled"

	fin = open(inputFile,"rb+")
	fout = open(outputFile,"wb+")

	# Generate values for encryption if encrypting
	if mode=="encrypt":
		salt = urandom(16)
		nonce = urandom(24)
		fout.write(str(len(ad)).encode("utf-8"))
		fout.write(b"|")
		fout.write(ad)
		fout.write(b"0"*64)
		fout.write(b"0"*64)
		fout.write(b"0"*16)
		fout.write(salt)
		fout.write(nonce)
	# If decrypting, read values from file
	else:
		# Read past metadata into actual data
		while True:
			letter = fin.read(1)
			if letter==b"|":
				break
		fin.read(int(adlen.decode("utf-8")))
		cs = fin.read(64)
		crccs = fin.read(64)
		digest = fin.read(16)
		salt = fin.read(16)
		nonce = fin.read(24)

	# Show notice, set progress bar indeterminate
	statusString.set(derivingNotice)
	progress.config(mode="indeterminate")
	progress.start(15)

	# Derive argon2id key
	key = hash_secret_raw(
		password,
		salt,
		time_cost=8, # 8 iterations
		memory_cost=2**20, # 2^20 Kilobytes (1GB)
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
			statusString.set(passwordNotice)
			fin.close()
			fout.close()
			remove(outputFile)
			selectFileInput["state"] = "normal"
			passwordInput["state"] = "normal"
			adArea["state"] = "normal"
			startBtn["state"] = "normal"
			keepBtn["state"] = "normal"
			working = False
			del key
			return

	# Create XChaCha20-Poly1305 object
	cipher = ChaCha20_Poly1305.new(key=key,nonce=nonce)
	# Cyclic redundancy check for file corruption
	crc = sha3_512()

	done = 0
	total = getsize(inputFile)
	chunkSize = 2**20
	startTime = datetime.now()

	# If secure wipe enabled, create a wiper object
	if wipe:
		wiper = open(inputFile,"r+b")
		wiper.seek(0)

	# Continously read file in chunks of 1MB
	while True:
		piece = fin.read(chunkSize)
		if wipe:
			# If securely wipe, write random trash
			# to original file after reading it
			trash = urandom(len(piece))
			wiper.write(trash)
		# If EOF
		if not piece:
			if mode=="encrypt":
				# Get the cipher MAC tag, write to file
				digest = cipher.digest()
				fout.flush()
				fout.close()
				fout = open(outputFile,"r+b")
				fout.seek(len(str(len(ad)))+1+len(ad))
				fout.write(check)
				fout.write(crc.digest())
				fout.write(digest)
			else:
				# If decrypting, verify MAC tag
				crcdg = crc.digest()
				if not compare_digest(crccs,crcdg):
					# File is corrupted
					statusString.set(corruptedNotice)
					progress["value"] = 100
					fin.close()
					fout.close()
					# If keep file checked...
					if keep.get()!=1:
						remove(outputFile)
						selectFileInput["state"] = "normal"
						passwordInput["state"] = "normal"
						adArea["state"] = "normal"
						startBtn["state"] = "normal"
						keepBtn["state"] = "normal"
						working = False
						del fin,fout,cipher,key
						return
					else:
						kept = "corrupted"
				try:
					# Throws ValueError if incorrect
					cipher.verify(digest)
				except:
					# File is modified
					statusString.set(modifiedNotice)
					progress["value"] = 100
					fin.close()
					fout.close()
					# If keep file checked...
					if keep.get()!=1:
						remove(outputFile)
						selectFileInput["state"] = "normal"
						passwordInput["state"] = "normal"
						adArea["state"] = "normal"
						startBtn["state"] = "normal"
						keepBtn["state"] = "normal"
						working = False
						del fin,fout,cipher,key
						return
					else:
						kept = "modified"					
			break
		
		# Encrypt/decrypt chunk and update CRC
		if mode=="encrypt":
			data = cipher.encrypt(piece)
			crc.update(data)
		else:
			crc.update(piece)
			data = cipher.decrypt(piece)

		# Calculate speed, ETA, etc.
		first = False
		elapsed = (datetime.now()-startTime).total_seconds()
		if elapsed==0:
			elapsed = 0.1**6
		percent = done*100/total
		progress["value"] = percent
		rPercent = round(percent)
		speed = (done/elapsed)/10**6
		if speed==0:
			first = True
			speed = 0.1**6
		rSpeed = round(speed)
		eta = round((total-done)/(speed*10**6))
		if first:
			statusString.set("...% at ... MB/s (ETA: ...s)")
		else:
			info = f"{rPercent}% at {rSpeed} MB/s (ETA: {eta}s)"
			statusString.set(info)
		
		done += chunkSize
		fout.write(data)

	# Show appropriate notice if file corrupted or modified
	if not kept:
		if mode=="encrypt":
			output = inputFile.split("/")[-1]+".pcf"
		else:
			output = inputFile.split("/")[-1].replace(".pcf","")
		statusString.set(f"Completed. (Output: {output})")
	else:
		if kept=="modified":
			statusString.set(kModifiedNotice)
		else:
			statusString.set(kCorruptedNotice)
	
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
		selectFileInput["state"] = "normal"
		passwordInput["state"] = "normal"
		adArea["state"] = "normal"
		startBtn["state"] = "normal"
		keepBtn["state"] = "normal"
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

# Frame so start button can fill width
startFrame = tkinter.Frame(
	tk,
	width=442,
	height=25
)
startFrame.place(x=19,y=290)
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
progress.place(x=20,y=328)

# Status label
statusString = tkinter.StringVar(tk)
statusString.set("Ready.")
status = tkinter.ttk.Label(
	tk,
	textvariable=statusString
)
status.place(x=17,y=356)
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
credits.place(x=17,y=386)
source = "https://github.com/HACKERALERT/Picocrypt"
credits.bind("<Button-1>",lambda e:webbrowser.open(source))

# Version
versionString = tkinter.StringVar(tk)
versionString.set("v1.4")
version = tkinter.ttk.Label(
	tk,
	textvariable=versionString
)
version["state"] = "disabled"
version.config(background="#f5f6f7")
version.place(x=436,y=386)

# Dummy button to remove focus from other buttons
# and prevent ugly border highlighting
dummy = tkinter.ttk.Button(
	tk
)
dummy.place(x=480,y=0)

# Close window only if not encryption or decrypting
def onClose():
	if not working:
		tk.destroy()

# Main tkinter loop
if __name__=="__main__":
	tk.protocol("WM_DELETE_WINDOW",onClose)
	tk.mainloop()
	sys.exit(0)
