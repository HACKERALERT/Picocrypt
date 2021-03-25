#!/usr/bin/env python3

"""

Picocrypt v1.12
Dependencies: argon2-cffi, pycryptodome, reedsolo, tkinterdnd2
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

"""

# Imports
from tkinter import filedialog,messagebox
from threading import Thread
from datetime import datetime
from argon2.low_level import hash_secret_raw,Type
from Crypto.Cipher import ChaCha20_Poly1305
from Crypto.Hash import SHA3_512 as sha3_512
from secrets import compare_digest
from os import urandom,fsync,remove,system
from os.path import getsize,expanduser,isdir
from os.path import dirname,abspath,realpath
from os.path import join as pathJoin,basename
from os.path import split as pathSplit,exists
from tkinterdnd2 import TkinterDnD,DND_FILES
from zipfile import ZipFile
from pathlib import Path
from shutil import rmtree,copyfile,copytree
from time import sleep
import sys
import tkinter
import tkinter.ttk
import tkinter.scrolledtext
import webbrowser
import platform
try:
	import winreg as wr
except:
	pass
from creedsolo import RSCodec,ReedSolomonError

# Tk/Tcl is a little barbaric, so I'm disabling
# high DPI so it doesn't scale bad and look horrible
try:
	from ctypes import windll
	windll.shcore.SetProcessDpiAwareness(0)
except:
	pass

# Global variables and strings
rootDir = dirname(realpath(__file__))

inputFile = ""
outputFile = ""
outputPath = ""
password = ""
ad = ""
kept = False
working = False
gMode = None
headerRsc = False
allFiles = False
draggedFolderPaths = False
files = False
filesLoaded = False
adString = "File metadata (used to store some text along with the file):"
compressingNotice = "Compressing files together..."
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
erasingNotice = "Securely erasing original file(s)..."
overwriteNotice = "Output file already exists. Would you like to overwrite it?"
cancelNotice = "Exiting now will lead to broken output. Are you sure?"
rsNotice = "Prevent corruption using Reed-Solomon"
rscNotice = "Creating Reed-Solomon tables..."
unknownErrorNotice = "Unknown error occured. Please try again."

# Create root Tk
tk = TkinterDnD.Tk()
tk.geometry("480x540")
tk.title("Picocrypt")
if platform.system()=="Darwin":
	tk.configure(background="#edeced")
else:
	tk.configure(background="#ffffff")
tk.resizable(0,0)

# Try setting window icon if included with Picocrypt
try:
	favicon = tkinter.PhotoImage(file="./key.png")
	tk.iconphoto(False,favicon)
except:
	pass

# Some styling
s = tkinter.ttk.Style()
s.configure("TCheckbutton",background="#ffffff")

# Event when user drags file(s) and folder(s) into window
def inputSelected(draggedFile):
	global inputFile,working,headerRsc,allFiles,draggedFolderPaths,files
	resetUI()
	dummy.focus()
	status.config(cursor="")
	status.bind("<Button-1>",lambda e:None)

	# Use try to handle errors
	try:
		# Create list of input files
		allFiles = []
		files = []
		draggedFolderPaths = []
		suffix = ""
		tmp = [i for i in draggedFile]
		res = []
		within = False
		tmpName = ""

		"""
		The next for loop parses data return by tkinterdnd2's file drop method.
		When files and folders are dragged, the output (the 'draggedFile' parameter)
		will contain the dropped files/folders and will look something like this:
		
		A single file/folder: "C:\Foo\Bar.txt"
		A single file/folder with a space in path: "{C:\Foo Bar\Lorem.txt}"
		Multiple files/folders: "C:\Foo\Bar1.txt C:\Foo\Ba2.txt"
		Multiple files/folders with spaces in paths: 
			- "C:\Foo\Bar1.txt {C:\Foo Bar\Lorem.txt}"
			- "{C:\Foo Bar\Lorem.txt} C:\Foo\Bar1.txt"
			- "{C:\Foo Bar\Lorem1.txt} {C:\Foo Bar\Lorem2.txt}"
		"""
		for i in tmp:
			if i=="{":
				within = True
			elif i=="}":
				within = False
				res.append(tmpName)
				tmpName = ""
			else:
				if i==" " and not within:
					if tmpName!="":
						res.append(tmpName)
					tmpName = ""
				else:
					tmpName += i
		if tmpName:
			res.append(tmpName)

		allFiles = []
		files = []

		# Check each thing dragged by user
		for i in res:
			# If there is a directory, recursively add all files to 'allFiles'
			if isdir(i):
				# Record the directory for secure wipe (if necessary)
				draggedFolderPaths.append(i)
				tmp = Path(i).rglob("*")
				for p in tmp:
					allFiles.append(abspath(p))
			# Just a file, add it to files
			else:
				files.append(i)

		# If there's only one file, set it as input file
		if len(files)==1 and len(allFiles)==0:
			inputFile = files[0]
			files = []
		else:
			inputFile = ""

		# Decide if encrypting or decrypting
		if inputFile.endswith(".pcv"):
			suffix = " (will decrypt)"
			fin = open(inputFile,"rb")

			# Read file metadata (a little complex)
			tmp = fin.read(139)
			reedsolo = False
			if tmp[0]==43:
				reedsolo = True
				tmp = tmp[1:]
			else:
				tmp = tmp[:-1]
			tmp = bytes(headerRsc.decode(tmp)[0])
			tmp = tmp.replace(b"+",b"")
			tmp = int(tmp.decode("utf-8"))
			if not reedsolo:
				fin.seek(138)
			ad = fin.read(tmp)
			try:
				ad = bytes(headerRsc.decode(ad)[0])
			except ReedSolomonError:
				ad = b"Error decoding file metadata."
			ad = ad.decode("utf-8")
			fin.close()

			# Insert the metadata into its text box
			adArea["state"] = "normal"
			adArea.delete("1.0",tkinter.END)
			adArea.insert("1.0",ad)
			adArea["state"] = "disabled"

			outputFrame.config(width=440)
			outputCheck.delete(0,tkinter.END)
			outputCheck.insert(0,inputFile[:-4])

			# Update UI
			adLabelString.set("File metadata (read only):")
			keepBtn["state"] = "normal"
			eraseBtn["state"] = "disabled"
			rsBtn["state"] = "disabled"
			cpasswordInput["state"] = "normal"
			cpasswordInput.delete(0,"end")
			cpasswordInput["state"] = "disabled"
			cpasswordString.set("Confirm password (N/A):")
		else:
			# Update the UI
			eraseBtn["state"] = "normal"
			keepBtn["state"] = "disabled"
			rsBtn["state"] = "normal"
			adArea["state"] = "normal"
			adArea.delete("1.0",tkinter.END)
			suffix = " (will encrypt)"
			adLabelString.set(adString)
			cpasswordInput["state"] = "normal"
			cpasswordInput.delete(0,"end")
			cpasswordString.set("Confirm password:")
			cpasswordLabel["state"] = "normal"
			adLabel["state"] = "normal"

			outputFrame.config(width=414)
			outputCheck.delete(0,tkinter.END)
			outputCheck.insert(0,inputFile)

		nFiles = len(files)
		nFolders = len(draggedFolderPaths)

		# Show selected file(s) and folder(s)
		if (allFiles or files) and not draggedFolderPaths:
			inputString.set(f"{nFiles} files selected (will encrypt).")
		elif draggedFolderPaths and not files:
			inputString.set(f"{nFolders} folder{'s' if nFolders!=1 else ''} selected (will encrypt).")
		elif draggedFolderPaths and (allFiles or files):
			inputString.set(
				f"{nFiles} file{'s' if nFiles!=1 else ''} and "+
				f"{nFolders} folder{'s' if nFolders!=1 else ''} selected (will encrypt)."
			)
		else:
			inputString.set(inputFile.split("/")[-1]+suffix)

		# Enable password box, etc.
		passwordInput["state"] = "normal"
		passwordInput.delete(0,"end")
		passwordLabel["state"] = "normal"
		startBtn["state"] = "normal"
		statusString.set("Ready.")
		status["state"] = "enabled"
		progress["value"] = 0
		clearInput["state"] = "normal"

	# File decode error
	except UnicodeDecodeError:
		statusString.set(corruptedNotice)
		progress["value"] = 100

	# No file(s) selected, do nothing
	except Exception as e:
		print(e)
		inputString.set("Drag and drop file(s) and folder(s) into this window.")
		resetUI()

	# Focus the dummy button to remove ugly borders
	finally:
		dummy.focus()
		working = False

def bindContextMenu():
	tmp = Path(rootDir).parent.absolute()
	target = pathJoin(expanduser("~"),"Picocrypt")
	vbs = pathJoin(target,"add_files.vbs")

	if exists(target):
		rmtree(target)
	copytree(tmp,target)

	keyVal = "Directory\\Shell\\Open in Picocrypt\\command"
	try:
		key = wr.OpenKey(
			wr.HKEY_CLASSES_ROOT, 
			keyVal, 
			0, 
			wr.KEY_ALL_ACCESS
		)
	except WindowsError:
		key = wr.CreateKey(wr.HKEY_CLASSES_ROOT,keyVal)

	regEntry = (r'wscript "{}" "%1"'.format(vbs))
	wr.SetValueEx(key,"",0,wr.REG_SZ,regEntry)
	wr.CloseKey(key)

	keyVal = "*\\Shell\\Open in Picocrypt\\command"
	try:
		key = wr.OpenKey(
			wr.HKEY_CLASSES_ROOT, 
			keyVal, 
			0, 
			wr.KEY_ALL_ACCESS
		)
	except WindowsError:
		key = wr.CreateKey(wr.HKEY_CLASSES_ROOT,keyVal)
	regEntry = (r'wscript "{}" "%1"'.format(vbs))
	wr.SetValueEx(key,"",0,wr.REG_SZ,regEntry)
	wr.CloseKey(key)

	a = open(vbs,"rb")
	b = a.read().decode("utf-8")
	a.close()
	b = b.replace("PICOCRYPT_PATH",target)
	a = open(vbs,"wb")
	a.write(b.encode("utf-8"))
	a.close()
	
	a = open(vbs.replace(".vbs",".bat"),"rb")
	b = a.read().decode("utf-8")
	a.close()
	b = b.replace("PICOCRYPT_PATH",target)
	a = open(vbs.replace(".vbs",".bat"),"wb")
	a.write(b.encode("utf-8"))
	a.close()

# Clears the selected files
def clearInputs():
	dummy.focus()
	resetUI()

# Allow drag and drop
def onDrop(e):
	global working
	if not working:
		inputSelected(e.data)
tk.drop_target_register(DND_FILES)
tk.dnd_bind("<<Drop>>",onDrop)

# Label that displays selected input file
inputString = tkinter.StringVar(tk)
inputString.set("Drag and drop file(s) and folder(s) into this window.")
selectedInput = tkinter.ttk.Label(
	tk,
	textvariable=inputString
)
selectedInput.config(background="#ffffff")
selectedInput.place(x=17,y=16)

# Clear input files
clearInput = tkinter.ttk.Button(
	tk,
	text="Clear",
	command=clearInputs
)
if platform.system()=="Darwin":
	clearInput.place(x=398,y=15,width=64,height=25)
else:
	clearInput.place(x=421,y=15,width=40,height=25)
clearInput["state"] = "disabled"

separator = tkinter.ttk.Separator(
	tk
)
separator.place(x=20,y=38,width=440)

outputString = tkinter.StringVar(tk)
outputString.set("Save output as:")
outputLabel = tkinter.ttk.Label(
	tk,
	textvariable=outputString
)
outputLabel.place(x=17,y=46)
outputLabel.config(background="#ffffff")
outputLabel["state"] = "disabled"

pvcString = tkinter.StringVar(tk)
pvcString.set(".pcv")
pvcLabel = tkinter.ttk.Label(
	tk,
	textvariable=pvcString
)
pvcLabel.place(x=436,y=66)
pvcLabel.config(background="#ffffff")

# A frame to make password input fill width
outputFrame = tkinter.Frame(
	tk,
	width=440,
	height=22
)
outputFrame.place(x=(17 if platform.system()=="Darwin" else 20),y=66)
outputFrame.columnconfigure(0,weight=10)
outputFrame.grid_propagate(False)
outputCheck = tkinter.ttk.Entry(
	outputFrame
)
outputCheck.grid(sticky="nesw")

# Label that prompts user to enter a password
passwordString = tkinter.StringVar(tk)
passwordString.set("Password:")
passwordLabel = tkinter.ttk.Label(
	tk,
	textvariable=passwordString
)
passwordLabel.place(x=17,y=96)
passwordLabel.config(background="#ffffff")
passwordLabel["state"] = "disabled"

# A frame to make password input fill width
passwordFrame = tkinter.Frame(
	tk,
	width=(445 if platform.system()=="Darwin" else 440),
	height=22
)
passwordFrame.place(x=(17 if platform.system()=="Darwin" else 20),y=116)
passwordFrame.columnconfigure(0,weight=10)
passwordFrame.grid_propagate(False)
# Password input box
passwordInput = tkinter.ttk.Entry(
	passwordFrame,
	show="\u2022"
)
passwordInput.grid(sticky="nesw")
passwordInput["state"] = "disabled"

cpasswordString = tkinter.StringVar(tk)
cpasswordString.set("Confirm password:")
cpasswordLabel = tkinter.ttk.Label(
	tk,
	textvariable=cpasswordString
)
cpasswordLabel.place(x=17,y=146)
cpasswordLabel.config(background="#ffffff")
cpasswordLabel["state"] = "disabled"

# A frame to make confirm password input fill width
cpasswordFrame = tkinter.Frame(
	tk,
	width=(445 if platform.system()=="Darwin" else 440),
	height=22
)
cpasswordFrame.place(x=(17 if platform.system()=="Darwin" else 20),y=166)
cpasswordFrame.columnconfigure(0,weight=10)
cpasswordFrame.grid_propagate(False)
# Confirm password input box
cpasswordInput = tkinter.ttk.Entry(
	cpasswordFrame,
	show="\u2022"
)
cpasswordInput.grid(sticky="nesw")
cpasswordInput["state"] = "disabled"

# Start the encryption/decryption process
def start():
	global inputFile,outputFile,password,ad,kept
	global working,gMode,headerRsc,allFiles,files
	global dragFolderPath
	dummy.focus()
	reedsolo = False
	chunkSize = 2**20

	# Decide if encrypting or decrypting
	if not inputFile.endswith(".pcv"):
		mode = "encrypt"
		gMode = "encrypt"
		outputFile = inputFile+".pcv"
		reedsolo = rs.get()==1
	else:
		mode = "decrypt"
		gMode = "decrypt"
		# Check if Reed-Solomon was enabled by checking for "+"
		test = open(inputFile,"rb")
		decider = test.read(1).decode("utf-8")
		test.close()
		if decider=="+":
			reedsolo = True
		# Decrypted output is just input file without the extension
		outputFile = inputFile[:-4]


	# Disable inputs and buttons while encrypting/decrypting
	disableAllInputs()

	# Make sure passwords match
	if passwordInput.get()!=cpasswordInput.get() and mode=="encrypt":
		resetEncryptionUI()
		statusString.set("Passwords don't match.")
		return

	# Set progress bar indeterminate
	progress.config(mode="indeterminate")
	progress.start(15)
	statusString.set(rscNotice)

	# Create Reed-Solomon object
	if reedsolo:
		# 13 bytes per 128 bytes, ~10% larger output file
		rsc = RSCodec(13)
	
	# Compress files together if user dragged multiple files
	if allFiles or files:
		statusString.set(compressingNotice)
		tmp = datetime.now().strftime("%Y-%m-%d_%H-%M-%S")
		if files:
			zfPath = Path(files[0]).parent.absolute()
		else:
			zfPath = Path(dirname(allFiles[0])).parent.absolute()
		zfOffset = len(str(zfPath))
		zfName = pathJoin(zfPath,tmp+".zip")
		zf = ZipFile(zfName,"w")
		for i in allFiles:
			zf.write(i,i[zfOffset:])
		for i in files:
			zf.write(i,pathSplit(i)[1])

		zf.close()
		inputFile = zfName
		outputFile = zfName+".pcv"
		outputPath = dirname(outputFile)

	# Set and get some variables
	working = True
	headerBroken = False
	reedsoloFixedCount = 0
	reedsoloErrorCount = 0
	dummy.focus()
	password = passwordInput.get().encode("utf-8")
	ad = adArea.get("1.0",tkinter.END).encode("utf-8")
	wipe = erase.get()==1

	# Open files
	try:
		fin = open(inputFile,"rb")
	except:
		resetEncryptionUI()
		statusString.set("Folder is empty.")
		return

	if reedsolo and mode=="decrypt":
		# Move pointer one forward
		fin.read(1)
	fout = open(outputFile,"wb+")
	if reedsolo and mode=="encrypt":
		# Signal that Reed-Solomon was enabled with a "+"
		fout.write(b"+")

	# Generate values for encryption if encrypting
	if mode=="encrypt":
		salt = urandom(16)
		nonce = urandom(24)

		# Reed-Solomon-encode metadata
		ad = bytes(headerRsc.encode(ad))
		# Write the metadata to output
		tmp = str(len(ad)).encode("utf-8")
		# Right-pad with "+"
		while len(tmp)!=10:
			tmp += b"+"
		tmp = bytes(headerRsc.encode(tmp))
		fout.write(tmp) # Length of metadata
		fout.write(ad) # Metadata (associated data)

		# Write zeros as placeholders, come back to write over it later.
		# Note that 128 extra Reed-Solomon bytes are added
		fout.write(b"0"*192) # SHA3-512 of encryption key
		fout.write(b"0"*192) # CRC of file
		fout.write(b"0"*144) # Poly1305 tag
		# Reed-Solomon-encode salt and nonce
		fout.write(bytes(headerRsc.encode(salt))) # Argon2 salt
		fout.write(bytes(headerRsc.encode(nonce))) # ChaCha20 nonce

	# If decrypting, read values from file
	else:
		# Move past metadata into actual data
		tmp = fin.read(138)
		if tmp[0]==43:
			tmp = tmp[1:]+fin.read(1)
		tmp = bytes(headerRsc.decode(tmp)[0])
		tmp = tmp.replace(b"+",b"")
		adlen = int(tmp.decode("utf-8"))
		fin.read(int(adlen))

		# Read the salt, nonce, etc.
		cs = fin.read(192)
		crccs = fin.read(192)
		digest = fin.read(144)
		salt = fin.read(144)
		nonce = fin.read(152)
		# Reed-Solomon-decode each value
		try:
			cs = bytes(headerRsc.decode(cs)[0])
		except:
			headerBroken = True
			cs = cs[:64]
		try:
			crccs = bytes(headerRsc.decode(crccs)[0])
		except:
			headerBroken = True
			crccs = crccs[:64]
		try:
			digest = bytes(headerRsc.decode(digest)[0])
		except:
			headerBroken = True
			digest = digest[:16]
		try:
			salt = bytes(headerRsc.decode(salt)[0])
		except:
			headerBroken = True
			salt = salt[:16]
		try:
			nonce = bytes(headerRsc.decode(nonce)[0])
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
				resetDecryptionUI()
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
		memory_cost=2**10, # 2^20 Kibibytes (1GiB)
		parallelism=8, # 8 parallel threads
		hash_len=32,
		type=Type.ID
	)

	# Key deriving done, set progress bar determinate
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0

	# Compute hash of derived key
	check = sha3_512.new()
	check.update(key)
	check = check.digest()

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
				resetDecryptionUI()
				return

	# Create XChaCha20-Poly1305 object
	cipher = ChaCha20_Poly1305.new(key=key,nonce=nonce)
	# Cyclic redundancy check for file corruption
	crc = sha3_512.new()

	# Amount of data encrypted/decrypted, total file size, starting time
	done = 0
	total = getsize(inputFile)

	# If secure wipe enabled, create a wiper object

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

		# If EOF
		if not piece:
			if mode=="encrypt":
				# Get the cipher MAC tag (Poly1305)
				digest = cipher.digest()
				fout.flush()
				fout.close()
				fout = open(outputFile,"r+b")
				# Compute the offset and seek to it (unshift "+")
				rsOffset = 1 if reedsolo else 0
				fout.seek(138+len(ad)+rsOffset)
				# Write hash of key, CRC, and Poly1305 MAC tag
				fout.write(bytes(headerRsc.encode(check)))
				fout.write(bytes(headerRsc.encode(crc.digest())))
				fout.write(bytes(headerRsc.encode(digest)))
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
						resetDecryptionUI()
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
							resetDecryptionUI()
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
						resetDecryptionUI()
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
		elapsed = (datetime.now()-previousTime).total_seconds() or 0.0001
		sinceStart = (datetime.now()-startTime).total_seconds() or 0.0001
		previousTime = datetime.now()

		percent = done*100/total
		progress["value"] = percent

		speed = (done/sinceStart)/10**6 or 0.0001
		eta = round((total-done)/(speed*10**6))

		# Seconds to minutes if seconds more than 59
		if eta>=60:
			# Set blank ETA if just starting
			if sinceStart<0.5:
				eta = "..."
			else:
				eta = f"{eta//60}m {eta%60}"
		if isinstance(eta,int) or isinstance(eta,float):
			if eta<0:
				eta = 0

		# Update status
		info = f"{percent:.0f}% at {speed:.2f} MB/s (ETA: {eta}s)"

		if reedsolo and mode=="decrypt" and reedsoloFixedCount:
			tmp = "s" if reedsoloFixedCount!=1 else ""
			info += f", fixed {reedsoloFixedCount} corrupted byte{tmp}"
		if reedsolo and mode=="decrypt" and reedsoloErrorCount:
			info += f", {reedsoloErrorCount} MB unrecoverable"

		statusString.set(info)
		
		# Increase done and write to output
		done += 1104905 if (reedsolo and mode=="decrypt") else chunkSize
		fout.write(data)

	# Flush outputs, close files
	if not kept:
		fout.flush()
		fsync(fout.fileno())
	fout.close()
	fin.close()

	# Securely wipe files as necessary
	if wipe:
		if draggedFolderPaths:
			for i in draggedFolderPaths:
				secureWipe(i)
		if files:
			for i in range(len(files)):
				statusString.set(erasingNotice+f" ({i}/{len(files)}")
				progress["value"] = i/len(files)
				secureWipe(files[i])
		secureWipe(inputFile)
	# Secure wipe not enabled
	else:
		if allFiles:
			# Remove temporary zip file if created
			remove(inputFile)

	# Show appropriate notice if file corrupted or modified
	if not kept:
		statusString.set(f"Completed. (Click here to show output)")

		# Show Reed-Solomon stats if it fixed corrupted bytes
		if mode=="decrypt" and reedsolo and reedsoloFixedCount:
			statusString.set(
				f"Completed with {reedsoloFixedCount}"+
				f" bytes fixed. (Output: {output})"
			)
	else:
		if kept=="modified":
			statusString.set(kModifiedNotice)
		elif kept=="corrupted":
			statusString.set(kCorruptedNotice)
		else:
			statusString.set(kVeryCorruptedNotice)
	
	status.config(cursor="hand2")
	
	# A little hack since strings are immutable
	output = "".join([i for i in outputFile])

	# Bind the output file
	if platform.system()=="Windows":
		status.bind("<Button-1>",
			lambda e:showOutput(output.replace("/","\\"))
		)
	else:
		status.bind("<Button-1>",
			lambda e:showOutput(output)
		)
	# Reset variables and UI states
	resetUI()
	status["state"] = "normal"
	inputFile = ""
	outputFile = ""
	password = ""
	ad = ""
	kept = False
	working = False
	allFiles = False
	dragFolderPath = False
	
	# Wipe keys for safety
	del fin,fout,cipher,key

# Wraps the start() function with error handling
def wrapper():
	global working,gMode
	# Try start() and handle errors
	try:
		start()
	except:
		# Reset UI accordingly

		if gMode=="decrypt":
			resetDecryptionUI()
		else:
			resetEncryptionUI()

		statusString.set(unknownErrorNotice)
		dummy.focus()
		working = False
	finally:
		sys.exit(0)

# Encryption/decrypt is done is a separate thread so the UI
# isn't blocked. This is a wrapper to spawn a thread and start it.
def startWorker():
	thread = Thread(target=wrapper,daemon=True)
	thread.start()

def begin(already=False):
	if not already:
		try:
			getsize(outputCheck.get())
			askConfirmOverwrite.pack(anchor=tkinter.W,fill=tkinter.BOTH,expand=True,side=tkinter.LEFT)
		except:
			startWorker()
	else:
		askConfirmOverwrite.pack_forget()
		startWorker()

# Securely wipe file
def secureWipe(fin):
	statusString.set(erasingNotice)
	# Check platform, erase accordingly
	if platform.system()=="Windows":
		if isdir(fin):
			paths = []
			for i in Path(fin).rglob("*"):
				if dirname(i) not in paths:
					paths.append(dirname(i))
			for i in range(len(paths)):
				statusString.set(erasingNotice+f" ({i}/{len(paths)})")
				progress["value"] = 100*i/len(paths)
				system(f'cd "{paths[i]}" && "{rootDir}/sdelete64.exe" * -p 4 -s -nobanner')
			system(f'cd "{rootDir}"')
			rmtree(fin)
		else:
			statusString.set(erasingNotice)
			progress["value"] = 100
			system(f'sdelete64.exe "{fin}" -p 4 -nobanner')
	elif platform.system()=="Darwin":
		system(f'rm -rfP "{fin}"')
	else:
		system(f'shred -uz "{fin}" -n 4')

# Disable all inputs while encrypting/decrypting
def disableAllInputs():
	passwordInput["state"] = "disabled"
	cpasswordInput["state"] = "disabled"
	clearInput["state"] = "disabled"
	adArea["state"] = "disabled"
	startBtn["state"] = "disabled"
	eraseBtn["state"] = "disabled"
	keepBtn["state"] = "disabled"
	rsBtn["state"] = "disabled"
	

# Reset UI to encryption state
def resetEncryptionUI():
	global working
	passwordInput["state"] = "normal"
	cpasswordInput["state"] = "normal"
	clearInput["state"] = "disabled"
	adArea["state"] = "normal"
	startBtn["state"] = "normal"
	eraseBtn["state"] = "normal"
	rsBtn["state"] = "normal"
	working = False
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 100

# Reset UI to decryption state
def resetDecryptionUI():
	global working
	passwordInput["state"] = "normal"
	clearInput["state"] = "normal"
	adArea["state"] = "normal"
	startBtn["state"] = "normal"
	keepBtn["state"] = "normal"
	working = False
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 100

# Reset UI to original state (no file selected)
def resetUI():
	adArea["state"] = "normal"
	adArea.delete("1.0",tkinter.END)
	adArea["state"] = "disabled"
	adLabel["state"] = "disabled"
	startBtn["state"] = "disabled"
	passwordInput["state"] = "normal"
	passwordInput.delete(0,"end")
	passwordInput["state"] = "disabled"
	passwordLabel["state"] = "disabled"
	cpasswordInput["state"] = "normal"
	cpasswordInput.delete(0,"end")
	cpasswordInput["state"] = "disabled"
	cpasswordString.set("Confirm password:")
	cpasswordLabel["state"] = "disabled"
	clearInput["state"] = "normal"
	status["state"] = "disabled"
	progress["value"] = 0
	inputString.set("Drag and drop file(s) and folder(s) into this window.")
	keepBtn["state"] = "normal"
	keep.set(0)
	keepBtn["state"] = "disabled"
	eraseBtn["state"] = "normal"
	erase.set(0)
	eraseBtn["state"] = "disabled"
	rs.set(0)
	rsBtn["state"] = "disabled"
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0
	
def showOutput(file):
	if platform.system()=="Windows":
		system(f'explorer /select,"{file}"')
	elif platform.system()=="Darwin":
		system(f'cd "{dirname(file)}"; open -R {pathSplit(file)[1]}')
		system(f'cd "{rootDir}"')
	else:
		system(f'xdg-open "{dirname(file)}"')

# ad stands for "associated data"/metadata
adLabelString = tkinter.StringVar(tk)
adLabelString.set(adString)
adLabel = tkinter.ttk.Label(
	tk,
	textvariable=adLabelString
)
adLabel.place(x=17,y=198)
adLabel.config(background="#ffffff")
adLabel["state"] = "disabled"

# Frame so metadata text box can fill width
adFrame = tkinter.Frame(
	tk,
	width=440,
	height=100
)
adFrame.place(x=20,y=218)
adFrame.columnconfigure(0,weight=10)
adFrame.grid_propagate(False)

# Metadata text box
import tkinter.scrolledtext
adArea = tkinter.scrolledtext.ScrolledText(
	adFrame,
	exportselection=0,
        height = 5,
)
adArea.config(font=("Consolas",12))
adArea.grid(sticky="nesw")
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
keepBtn.place(x=18,y=330)
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
eraseBtn.place(x=18,y=350)
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
rsBtn.place(x=18,y=370)
rsBtn["state"] = "disabled"

# Frame so start button can fill width
startFrame = tkinter.Frame(
	tk,
	width=442,
	height=24
)
startFrame.place(x=19,y=400)
startFrame.columnconfigure(0,weight=10)
startFrame.grid_propagate(False)
# Start button
startBtn = tkinter.ttk.Button(
	startFrame,
	text="Start",
	command=begin
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
progress.place(x=20,y=428)

# Status label
statusString = tkinter.StringVar(tk)
statusString.set("Ready.")
status = tkinter.ttk.Label(
	tk,
	textvariable=statusString
)
status.place(x=17,y=456)
status.config(background="#ffffff")
status["state"] = "disabled"

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
credits.config(background="#ffffff")
credits.place(x=17,y=486)
source = "https://github.com/HACKERALERT/Picocrypt"
credits.bind("<Button-1>",lambda e:webbrowser.open(source))

# Version
versionString = tkinter.StringVar(tk)
versionString.set("v1.11")
version = tkinter.ttk.Label(
	tk,
	textvariable=versionString
)
version["state"] = "disabled"
version.config(background="#ffffff")
version.place(x=(420 if platform.system()=="Darwin" else 430),y=486)

# Helper to ask/confirm operations
askConfirmOverwrite = tkinter.Frame(tk)
askConfirmOverwrite.config(background="#ffffff")
confirmOverwriteString = tkinter.StringVar(tk)
confirmOverwriteString.set(overwriteNotice)
confirmOverwrite = tkinter.ttk.Label(
	askConfirmOverwrite,
	textvariable=confirmOverwriteString
)
confirmOverwrite.place(x=90,y=170)
confirmOverwrite.config(background="#ffffff")
yes = tkinter.ttk.Button(
	askConfirmOverwrite,
	text="Yes",
	command=lambda:begin(True)
)
yes.place(x=160,y=210)
no = tkinter.ttk.Button(
	askConfirmOverwrite,
	text="No",
	command=lambda:askConfirmOverwrite.pack_forget()
)
no.place(x=240,y=210)

# Dummy button to remove focus from other buttons
# and prevent ugly border highlighting
dummy = tkinter.ttk.Button(
	tk
)
dummy.place(x=480,y=0)

# Function to create Reed-Solomon header codec
def createRsc():
	global headerRsc
	headerRsc = RSCodec(128)
	sys.exit(0)
	
def prepare():
	if platform.system()=="Windows":
		system("sdelete64.exe /accepteula")
		if windll.shell32.IsUserAnAdmin():
			bindContextMenu()

def awaitFiles():
	global filesLoaded
	inputString.set("Loading the file(s) and folder(s) you selected...")
	a = open(files,"rb")
	b = a.read().decode("utf-8")
	a.close()
	while True:
		sleep(5)
		a = open(files,"rb")
		c = a.read().decode("utf-8")
		a.close()
		if b==c:
			b = c
			break

	a.close()
	b = b.replace("\r\n"," ").replace('"',"")
	remove(files)
	remove(files.replace("files.txt","tmp.txt"))
	try:
		inputSelected(b)
	except:
		pass
	filesLoaded = True
	sys.exit(0)

# Close window only if not encrypting or decrypting
def onClose():
	global outputFile,filesLoaded
	if not working and filesLoaded:
		tk.destroy()
	else:
		force = messagebox.askyesno("Confirmation",cancelNotice)
		if force:
			tk.destroy()

# Main application loop
if __name__=="__main__":
	# Create Reed-Solomon header codec
	tmp = Thread(target=createRsc,daemon=True)
	tmp.start()

	# Prepare application
	tmp = Thread(target=prepare,daemon=True)
	tmp.start()

	# Windows context menu
	if platform.system()=="Windows":
		try:
			files = pathJoin(Path(rootDir).parent.absolute(),"files.txt")
			getsize(files)
			tmp = Thread(target=awaitFiles,daemon=True)
			tmp.start()
		except:
			filesLoaded = True

	# Start tkinter
	tk.protocol("WM_DELETE_WINDOW",onClose)
	tk.mainloop()
	sys.exit(0)
