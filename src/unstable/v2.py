#!/usr/bin/env python3

"""

Picocrypt v1.12
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

"""

# Import dependencies
from threading import Thread
from datetime import datetime
from argon2.low_level import hash_secret_raw
from argon2.low_level import Type as argonType
from Crypto.Cipher import ChaCha20_Poly1305
from Crypto.Hash import SHA3_512
from hmac import compare_digest
from reedsolo import RSCodec,ReedSolomonError
from os import urandom,fsync,remove,system
from os.path import getsize,expanduser,isdir,exists,dirname,abspath,realpath
from os.path import join as pathJoin,split as pathSplit
from pathlib import Path
from zipfile import ZipFile
from tkinterdnd2 import TkinterDnD,DND_FILES
from ttkthemes import ThemedStyle
import re
import sys
import tkinter
import tkinter.ttk
import tkinter.scrolledtext
import webbrowser
import platform

# Global variables
rootDir = dirname(realpath(__file__))
working = False
mode = False
inputFile = False
outputFile = False
rs128 = False
rs13 = False
allFiles = False
onlyFolders = False
onlyFiles = False

# Strings
strings = [
	"File metadata (used to store some text along with the file):",
	"Compressing files together...",
	"Error. The provided password is incorrect.",
	"Error. The input file is corrupted.",
	"Error. The input file and header keys are badly corrupted.",
	"Error. The input file has been intentionally modified.",
	"The input file is corrupted, but the output has been kept.",
	"The input file has been intentionally modified, but the output has been kept.",
	"The input file is badly corrupted, but the output has been kept.",
	"Deriving key (takes a few seconds)...",
	"Keep decrypted output even if it's corrupted or modified",
	"Securely erase and delete original file",
	"Securely erasing original file(s)...",
	"Output file already exists. Would you like to overwrite it?",
	"Exiting now will lead to broken output. Are you sure?",
	"Prevent corruption using Reed-Solomon",
	"Error. Folder(s) and/or file(s) are empty.",
	"Unknown error occured. Please try again.",
	"Drag and drop file(s) and folder(s) into this window.",
	"File metadata (read-only):",
	"Error. The input file couldn't be decoded as UTF-8."
]

# Create root tk
tk = TkinterDnD.Tk()
tk.geometry("480x512")
tk.title("Picocrypt")
tk.resizable(0,0)
tk.configure(background="#f5f6f7")
ThemedStyle(tk).set_theme("arc")

# Enable high DPI on Windows
def Get_HWND_DPI(window_handle):
	from ctypes import windll,pointer,wintypes
	windll.shcore.SetProcessDpiAwareness(1)
	DPI100pc = 96
	DPI_type = 0
	winH = wintypes.HWND(window_handle)
	monitorhandle = windll.user32.MonitorFromWindow(
		winH,wintypes.DWORD(2)
	)
	X = wintypes.UINT()
	Y = wintypes.UINT()
	try:
		windll.shcore.GetDpiForMonitor(
			monitorhandle,DPI_type,pointer(X),pointer(Y)
		)
		return X.value*2,Y.value*2,(X.value+Y.value)/(2*DPI100pc)
	except Exception:
		return 96,96,1
def TkGeometryScale(s,cvtfunc):
	patt = r"(?P<W>\d+)x(?P<H>\d+)\+(?P<X>\d+)\+(?P<Y>\d+)"
	R = re.compile(patt).search(s)
	G = str(cvtfunc(R.group("W")))+"x"
	G += str(cvtfunc(R.group("H")))+"+"
	G += str(cvtfunc(R.group("X")))+"+"
	G += str(cvtfunc(R.group("Y")))
	return G
def MakeTkDPIAware(TKGUI):
	TKGUI.DPI_X,TKGUI.DPI_Y,TKGUI.DPI_scaling = Get_HWND_DPI(TKGUI.winfo_id())
	TKGUI.TkScale = lambda v:int(float(v)*TKGUI.DPI_scaling)
	TKGUI.TkGeometryScale = lambda s:TkGeometryScale(s,TKGUI.TkScale)
if platform.system()=="Windows":
	MakeTkDPIAware(tk)

# Add some styling
style = tkinter.ttk.Style()

# Try setting window icon if it exists
try:
	favicon = tkinter.PhotoImage(file="./key.png")
	tk.iconphoto(False,favicon)
except:
	pass

# Dummy button used for removing ugly highlights
dummy = tkinter.ttk.Button(tk)
dummy.place(x=480,y=0)

# Label that shows the input files
inputString = tkinter.StringVar(tk)
inputString.set(strings[18])
inputLabel = tkinter.ttk.Label(
	tk,
	textvariable=inputString
)
inputLabel.place(x=20,y=18)

# Clear input files
clearInput = tkinter.ttk.Button(
	tk,
	text="Clear",
	command = lambda:resetUI()
)
clearInput.place(x=400,y=12,width=60,height=28)
clearInput["state"] = "disabled"

# Separator for aesthetics
separator = tkinter.Frame(
	tk,
	bg="#dfe3ed",
	height=1
)
separator.place(x=20,y=39,width=438)

# Label to ask user what to save output as
outputString = tkinter.StringVar(tk)
outputString.set("Save output as:")
outputLabel = tkinter.ttk.Label(
	tk,
	textvariable=outputString
)
outputLabel.place(x=20,y=48)

outputLabel["state"] = "disabled"

# A ".pcv" extension shown next to output box
pcvString = tkinter.StringVar(tk)
pcvString.set(".pcv")
pcvLabel = tkinter.ttk.Label(
	tk,
	textvariable=pcvString
)
pcvLabel.place(x=434,y=68)

# A frame to allow output box to fill width
outputFrame = tkinter.Frame(
	tk,
	width=440,
	height=24
)
outputFrame.place(x=20,y=66)
outputFrame.columnconfigure(0,weight=10)
outputFrame.grid_propagate(False)

# Output box to allow user to change output name and path
outputInput = tkinter.ttk.Entry(outputFrame)
outputInput.grid(sticky="nesw")
outputInput["state"] = "disabled"

# Prompt user to enter password
passwordString = tkinter.StringVar(tk)
passwordString.set("Password:")
passwordLabel = tkinter.ttk.Label(
	tk,
	textvariable=passwordString
)
passwordLabel.place(x=20,y=100)
passwordLabel["state"] = "disabled"

# Allow password input to fill width
passwordFrame = tkinter.Frame(
	tk,
	width=440,
	height=24
)
passwordFrame.place(x=20,y=118)
passwordFrame.columnconfigure(0,weight=10)
passwordFrame.grid_propagate(False)

# Password input box
passwordInput = tkinter.ttk.Entry(
	passwordFrame,
	show="\u2022"
)
passwordInput.grid(sticky="nesw")
passwordInput["state"] = "disabled"

# Prompt user to confirm password
cPasswordString = tkinter.StringVar(tk)
cPasswordString.set("Confirm password:")
cPasswordLabel = tkinter.ttk.Label(
	tk,
	textvariable=cPasswordString
)
cPasswordLabel.place(x=20,y=150)
cPasswordLabel["state"] = "disabled"

# Allow confirm password input to fill width
cPasswordFrame = tkinter.Frame(
	tk,
	width=440,
	height=24
)
cPasswordFrame.place(x=20,y=168)
cPasswordFrame.columnconfigure(0,weight=10)
cPasswordFrame.grid_propagate(False)

# Confirm password input box
cPasswordInput = tkinter.ttk.Entry(
	cPasswordFrame,
	show="\u2022"
)
cPasswordInput.grid(sticky="nesw")
cPasswordInput["state"] = "disabled"

# Prompt user for optional metadata
metadataString = tkinter.StringVar(tk)
metadataString.set(strings[0])
metadataLabel = tkinter.ttk.Label(
	tk,
	textvariable=metadataString
)
metadataLabel.place(x=20,y=202)
metadataLabel["state"] = "disabled"

# Frame so metadata box can fill width
metadataFrame = tkinter.Frame(
	tk,
	width=439,
	height=99
)
metadataFrame.place(x=20,y=220)
metadataFrame.columnconfigure(0,weight=10)
metadataFrame.rowconfigure(0,weight=10)
metadataFrame.grid_propagate(False)
metadataFrame.config(bg="#e5eaf0")

# Metadata text box
metadataInput = tkinter.scrolledtext.ScrolledText(
	metadataFrame,
	exportselection=0,
	height=5,
	padx=5,
	pady=5
)
metadataInput.config(font=("Consolas",12))
metadataInput.grid(row=0,column=0,sticky="nesw",padx=1,pady=1)
metadataInput.config(borderwidth=0)
metadataInput.config(bg="#fbfcfc")
metadataInput["state"] = "disabled"
metadataInput.bind(
	"<FocusIn>",
	lambda e:metadataBoxUI("in")
)
metadataInput.bind(
	"<FocusOut>",
	lambda e:metadataBoxUI("out")
)
# Tkinter's Text() boxes are ugly, so I beautify it manually
def metadataBoxUI(what):
	if what=="in":
		if metadataInput.cget("bg")=="#ffffff":
			metadataFrame.config(bg="#78a7e5")
	else:
		metadataFrame.config(bg="#d8ddea")

# Check box for keeping corrupted or modified output
keep = tkinter.IntVar(tk)
keepBtn = tkinter.ttk.Checkbutton(
	tk,
	text=strings[10],
	variable=keep,
	onvalue=1,
	offvalue=0,
	command=lambda:dummy.focus()
)
keepBtn.place(x=18,y=329)
keepBtn["state"] = "disabled"

# Check box for securely erasing original files
erase = tkinter.IntVar(tk)
eraseBtn = tkinter.ttk.Checkbutton(
	tk,
	text=strings[11],
	variable=erase,
	onvalue=1,
	offvalue=0,
	command=lambda:dummy.focus()
)
eraseBtn.place(x=18,y=349)
eraseBtn["state"] = "disabled"

# Check box for enabling Reed-Solomon anti-corruption
rs = tkinter.IntVar(tk)
rsBtn = tkinter.ttk.Checkbutton(
	tk,
	text=strings[15],
	variable=rs,
	onvalue=1,
	offvalue=0,
	command=lambda:dummy.focus()
)
rsBtn.place(x=18,y=369)
rsBtn["state"] = "disabled"

# Frame so start and cancel button can fill width
startFrame = tkinter.Frame(
	tk,
	width=440,
	height=29
)
startFrame.place(x=20,y=402)
startFrame.columnconfigure(0,weight=10)
startFrame.grid_propagate(False)
startFrame.config(background="#ffffff")

# Start button
startBtn = tkinter.ttk.Button(
	startFrame,
	text="Start",
	command=lambda:Thread(target=work,daemon=True).start()
)
startBtn.grid(row=0,column=0,stick="nesw")
startBtn["state"] = "disabled"

# Cancel button
cancelBtn = tkinter.ttk.Button(
	startFrame,
	text="Cancel"
)
cancelBtn.grid(stick="nesw")
cancelBtn.grid(row=0,column=1)
cancelBtn["state"] = "disabled"

# Progress bar
progress = tkinter.ttk.Progressbar(
	tk,	
	orient=tkinter.HORIZONTAL,
	length=440,
	mode="determinate"
)
progress.place(x=20,y=439)

# Status label
statusString = tkinter.StringVar(tk)
statusString.set("Ready.")
status = tkinter.ttk.Label(
	tk,
	textvariable=statusString
)
status.place(x=20,y=453)

# Credits :D
hint = "Created by Evan Su. Click for details and source."
creditsString = tkinter.StringVar(tk)
creditsString.set(hint)
credits = tkinter.ttk.Label(
	tk,
	textvariable=creditsString,
	cursor="hand2"
)
credits.place(x=20,y=480)
source = "https://github.com/HACKERALERT/Picocrypt"
credits.bind("<Button-1>",lambda e:webbrowser.open(source))
credits["state"] = "disabled"

# Version
versionString = tkinter.StringVar(tk)
versionString.set("v1.12")
version = tkinter.ttk.Label(
	tk,
	textvariable=versionString
)
version["state"] = "disabled"
version.place(x=430,y=480)

# Files have been dragged
def filesDragged(draggedFiles):
	global inputFile,rs128,onlyFiles,mode,onlyFolders,allFiles
	resetUI()
	status.config(cursor="")
	status.bind("<Button-1>",lambda e:None)
	
	try:
		# Create lists to track files dragged
		onlyFiles = []
		onlyFolders = []
		allFiles = []
		tmpName = ""
		tmp = [i for i in draggedFiles]
		res = []
		within = False
		
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
					if tmpName:
						res.append(tmpName)
					tmpName = ""
				else:
					tmpName += i
		if tmpName:
			res.append(tmpName)
		
		# Check each item dragged by user
		for i in res:
			# If it's a directory, recursively add all files
			if isdir(i):
				onlyFolders.append(i)
				tmp = Path(i).rglob("*")
				for p in tmp:
					allFiles.append(abspath(p))
			# Just a file, add it to 'onlyFiles'
			else:
				onlyFiles.append(i)
		
		# If there's only one file, set it as 'inputFile'
		if len(onlyFiles)==1 and not len(allFiles):
			inputFile = onlyFiles[0]
			onlyFiles = []
		else:
			inputFile = ""
			
		# Decide if encrypting or decrypting
		if inputFile.endswith(".pcv"):
			mode = "decrypt"
			suffix = " (will decrypt)"
	
			# Read file metadata
			fin = open(inputFile,"rb")
			fin.read(129)
			metadataLength = fin.read(138)
			metadataLength = bytes(rs128.decode(metadataLength)[0])
			metadataLength = metadataLength.replace(b"+",b"")
			metadata = fin.read(int(metadataLength.decode("utf-8")))
			metadata = bytes(rs128.decode(metadata)[0]).decode("utf-8")
			metadataString.set("File metadata (read only):")
			metadataInput["state"] = "normal"
			metadataInput.delete("1.0",tkinter.END)
			metadataInput.insert("1.0",metadata)
			metadataInput["state"] = "disabled"
			fin.close()
			
			# Insert filename into output box
			outputFrame.config(width=440)
			outputInput["state"] = "normal"
			outputInput.delete(0,tkinter.END)
			outputInput.insert(0,inputFile[:-4])
			
			# Update UI
			setDecryptionUI()
		else:
			mode = "encrypt"

			# Update UI
			setEncryptionUI()
			
			# Update output box with appropriate name
			if inputFile:
				outputInput.insert(0,inputFile)
			else:
				if onlyFiles:
					tmp = Path(onlyFiles[0]).parent.absolute()
				else:
					tmp = Path(onlyFolders[0]).parent.absolute()
				tmp = pathJoin(tmp,"Encrypted.zip")
				tmp = tmp.replace("\\","/")
				outputInput.insert(0,tmp)
			suffix = " (will encrypt)"
			
		nFiles = len(onlyFiles)
		nFolders = len(onlyFolders)

		# Show selected file(s) and folder(s)
		if (allFiles or onlyFiles) and not onlyFolders:
			inputString.set(f"{nFiles} files selected (will encrypt).")
		elif onlyFolders and not onlyFiles:
			inputString.set(f"{nFolders} folder{'s' if nFolders!=1 else ''} selected (will encrypt).")
		elif onlyFolders and (allFiles or onlyFiles):
			inputString.set(
				f"{nFiles} file{'s' if nFiles!=1 else ''} and "+
				f"{nFolders} folder{'s' if nFolders!=1 else ''} selected (will encrypt)."
			)
		else:
			inputString.set(inputFile.split("/")[-1]+suffix)
	
	# UTF-8 decode error
	except UnicodeDecodeError:
		statusString.set(strings[20])
		progress["value"] = 100
	
	# Nothing happened
	except:
		pass
	
# Bind drag and drop to window
def onDrop(e):
	global working
	if not working:
		filesDragged(e.data)
	clearInput["state"] = "normal"
	clearInput.config(cursor="hand2")
tk.drop_target_register(DND_FILES)
tk.dnd_bind("<<Drop>>",onDrop)

def work():
	global inputFile,outputFile,working,mode,rs13,rs128,onlyFiles,onlyFolders,allFiles
	disableAllInputs()
	dummy.focus()

	# Set and get some variables
	kept = False
	shouldKeep = keep.get()==1
	shouldErase = erase.get()==1
	reedsolo = rs.get()==1
	working = True
	headerBroken = False
	reedsoloFixed = 0
	reedsoloErrors = 0
	password = passwordInput.get().encode("utf-8")
	metadata = metadataInput.get("1.0",tkinter.END).encode("utf-8")
	
	# Decide if encrypting or decrypting
	if mode=="encrypt":
		outputFile = outputInput.get()+".pcv"
	else:
		outputFile = outputInput.get()
	
	# Make sure passwords match
	if passwordInput.get()!=cPasswordInput.get() and mode=="encrypt":
		setEncryptionUI()
		statusString.set("Passwords don't match.")
		return
		
	# Set progress bar indeterminate
	progress.config(mode="indeterminate")
	progress.start(15)

	# Compress files together if necessary
	if onlyFiles or allFiles:
		statusString.set(strings[1])
		tmp = outputFile[:-4]
		if onlyFiles:
			zfPath = Path(onlyFiles[0]).parent.absolute()
		else:
			zfPath = Path(dirname(allFiles[0])).parent.absolute()
		zfOffset = len(str(zfPath))
		zfName = pathJoin(zfPath,tmp)
		zf = ZipFile(zfName,"w")
		for i in allFiles:
			zf.write(i,i[zfOffset:])
		for i in onlyFiles:
			zf.write(i,pathSplit(i)[1])
	
		zf.close()
		inputFile = zfName
		outputFile = zfName+".pcv"
		outputPath = dirname(outputFile)
	
	# Open files
	try:
		fin = open(inputFile,"rb")
	except:
		setEncryptionUI()
		statusString.set(strings[16])
		return
	
	if mode=="encrypt":
		salt = urandom(16)
		nonce = urandom(24)
		fout = open(outputFile,"wb+")
		if reedsolo:
			fout.write(rs128.encode(b"+"))
		else:
			fout.write(rs128.encode(b"-"))
	
		metadata = rs128.encode(metadata)
		tmp = len(metadata)
		tmp = f"{tmp:+<10}"
		tmp = rs128.encode(tmp.encode("utf-8"))
		
		fout.write(tmp)
		fout.write(metadata)
		fout.write(rs128.encode(salt)) # Argon2 salt
		fout.write(rs128.encode(nonce)) # ChaCha20 nonce
		fout.write(b"0"*192) # Hash of key
		fout.write(b"0"*144) # Poly1305 MAC
		fout.write(b"0"*192) # CRC
	else:
		tmp = fin.read(129)
		if bytes(rs128.decode(tmp)[0])==b"+":
			reedsolo = True
		else:
			reedsolo = False

		metadataLength = fin.read(138)
		metadataLength = bytes(rs128.decode(metadataLength)[0])
		metadataLength = metadataLength.replace(b"+",b"")
		fin.read(int(metadataLength.decode("utf-8")))

		salt = fin.read(144)
		nonce = fin.read(152)
		keycs = fin.read(192)
		maccs = fin.read(144)
		crccs = fin.read(192)
		
		salt = bytes(rs128.decode(salt)[0])
		nonce = bytes(rs128.decode(nonce)[0])
		keycs = bytes(rs128.decode(keycs)[0])
		maccs = bytes(rs128.decode(maccs)[0])
		crccs = bytes(rs128.decode(crccs)[0])
				
	statusString.set(strings[9])
	
	key = hash_secret_raw(
		password,
		salt,
		time_cost=8,
		memory_cost=2**10,
		parallelism=8,
		hash_len=32,
		type=argonType.D
	)
	
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0
	
	check = SHA3_512.new(data=key).digest()
	
	if mode=="decrypt":
		if not compare_digest(check,keycs):
			if not headerBroken:
				statusString.set(strings[2])
				fin.close()
				setDecryptionUI()
				return
		fout = open(outputFile,"wb+")
	
	crc = SHA3_512.new()
	cipher = ChaCha20_Poly1305.new(key=key,nonce=nonce)
	
	done = 0
	total = getsize(inputFile)
	
	startTime = datetime.now()
	previousTime = datetime.now()
	while True:
		if mode=="encrypt":
			piece = fin.read(2**20)
		else:
			piece = fin.read(2**20)

		if not piece:
			break

		if mode=="encrypt":
			data = cipher.encrypt(piece)
			fout.write(data)
		else:
			data = cipher.decrypt(piece)
			fout.write(data)
			
		elapsed = (datetime.now()-previousTime).total_seconds() or 0.0001
		sinceStart = (datetime.now()-startTime).total_seconds() or 0.0001
		previousTime = datetime.now()
		percent = done*100/total
		progress["value"] = percent
		done += 2**20
		speed = (done/sinceStart)/10**6 or 0.0001
		eta = round((total-done)/(speed*10**6))
		
		info = f"{percent:.0f}% at {speed:.2f} MB/s (ETA: {eta}s)"

		'''if reedsolo and mode=="decrypt" and reedsoloFixedCount:
			tmp = "s" if reedsoloFixedCount!=1 else ""
			info += f", fixed {reedsoloFixedCount} corrupted byte{tmp}"
		if reedsolo and mode=="decrypt" and reedsoloErrorCount:
			info += f", {reedsoloErrorCount} MB unrecoverable"'''

		statusString.set(info)
		
		
	
	if mode=="encrypt":
		fout.flush()
		fout.close()
		fout = open(outputFile,"r+b")
		fout.seek(129+138+len(metadata)+144+152)
		fout.write(rs128.encode(check))
		fout.write(rs128.encode(cipher.digest()))
		fout.write(rs128.encode(crc.digest()))
	else:
		if not compare_digest(crccs,crc.digest()):
			statusString.set(strings[3])
			progress["value"] = 100
			fin.close()
			fout.close()
			
			if keep.get()!=1:
				remove(outputFile)
				setDecryptionUI()
				return
			else:
				if not kept:
					kept = "corrupted"
		
		try:
			cipher.verify(maccs)
		except:
			#if not reedsoloErrorCount and not headerBroken:
			if True:
				# File is modified
				statusString.set(modifiedNotice)
				progress["value"] = 100
				fin.close()
				fout.close()
				# If keep file not checked...
				if keep.get()!=1:
					remove(outputFile)
					# Reset UI
					setDecryptionUI()
					return
				else:
					if not kept:
						kept = "modified"
						
	# Flush outputs, close files
	#if not kept:
	if True:
		fout.flush()
		fsync(fout.fileno())
	fout.close()
	fin.close()
	
	print("DONEDONEDONEDONEDONEDONEDONEDONEDONEDONE")
	# Securely wipe files as necessary
	if shouldErase:
		if onlyFolders:
			for i in onlyFolders:
				secureWipe(i)
		if onlyFiles:
			for i in range(len(onlyFiles)):
				statusString.set(
					strings[12]+f" ({i}/{len(onlyFiles)}"
				)
				progress["value"] = i/len(onlyFiles)
				secureWipe(onlyFiles[i])
		secureWipe(inputFile)
	# Secure wipe not enabled
	else:
		if allFiles or onlyFiles:
			# Remove temporary zip file if created
			remove(inputFile)

	# Show appropriate notice if file corrupted or modified
	#if not kept:
	if True:
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
	inputFile = ""
	outputFile = ""
	working = False
	allFiles = []
	onlyFolders = []
	onlyFiles = []
	
	# Wipe keys for safety
	#del fin,fout,cipher,key
	
def secureWipe(fin):
	statusString.set(strings[12])
	# Check platform, erase accordingly
	if platform.system()=="Windows":
		if isdir(fin):
			paths = []
			for i in Path(fin).rglob("*"):
				if dirname(i) not in paths:
					paths.append(dirname(i))
			for i in range(len(paths)):
				statusString.set(strings[12]+f" ({i}/{len(paths)})")
				progress["value"] = 100*i/len(paths)
				system(f'cd "{paths[i]}" && "{rootDir}/sdelete64.exe" * -p 4 -s -nobanner')
			system(f'cd "{rootDir}"')
			rmtree(fin)
		else:
			statusString.set(strings[12])
			progress["value"] = 100
			system(f'sdelete64.exe "{fin}" -p 4 -nobanner')
	elif platform.system()=="Darwin":
		system(f'rm -rfP "{fin}"')
	else:
		system(f'shred -uz "{fin}" -n 4')

# Reset UI to state where no files are selected
def resetUI():
	inputString.set(strings[18])
	inputLabel["state"] = "normal"
	clearInput["state"] = "disabled"
	clearInput.config(cursor="")
	outputLabel["state"] = "disabled"
	outputFrame.config(width=440)
	outputInput["state"] = "normal"
	outputInput.delete(0,"end")
	outputInput["state"] = "disabled"
	passwordLabel["state"] = "disabled"
	passwordInput["state"] = "normal"
	passwordInput.delete(0,"end")
	passwordInput["state"] = "disabled"
	cPasswordString.set("Confirm password:")
	cPasswordLabel["state"] = "disabled"
	cPasswordInput["state"] = "normal"
	cPasswordInput.delete(0,"end")
	cPasswordInput["state"] = "disabled"
	metadataFrame.config(bg="#e5eaf0")
	metadataInput.config(bg="#fbfcfc")
	metadataInput.config(fg="#000000")
	metadataString.set(strings[0])
	metadataLabel["state"] = "disabled"
	metadataInput["state"] = "normal"
	metadataInput.delete("1.0",tkinter.END)
	metadataInput["state"] = "disabled"
	keep.set(0)
	keepBtn["state"] = "disabled"
	erase.set(0)
	eraseBtn["state"] = "disabled"
	rs.set(0)
	rsBtn["state"] = "disabled"
	startBtn["state"] = "disabled"
	cancelBtn["state"] = "disabled"
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0
	dummy.focus()

# Set UI to encryption state
def setEncryptionUI():
	outputLabel["state"] = "normal"
	outputInput["state"] = "normal"
	outputFrame.config(width=410)
	passwordLabel["state"] = "normal"
	passwordInput["state"] = "normal"
	cPasswordLabel["state"] = "normal"
	cPasswordString.set("Confirm password:")
	cPasswordInput["state"] = "normal"
	metadataFrame.config(bg="#d8ddea")
	metadataInput.config(bg="#ffffff")
	metadataInput.config(fg="#000000")
	metadataLabel["state"] = "normal"
	metadataInput["state"] = "normal"
	eraseBtn["state"] = "normal"
	rsBtn["state"] = "normal"
	startBtn["state"] = "normal"
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0
	
# Set UI to decryption state
def setDecryptionUI():
	outputLabel["state"] = "normal"
	outputInput["state"] = "normal"
	outputFrame.config(width=440)
	passwordLabel["state"] = "normal"
	passwordInput["state"] = "normal"
	cPasswordString.set("Confirm password (N/A):")
	metadataFrame.config(bg="#e5eaf0")
	metadataInput.config(bg="#fbfcfc")
	metadataInput.config(fg="#666666")
	metadataString.set(strings[19])
	metadataInput["state"] = "disabled"
	keepBtn["state"] = "normal"
	startBtn["state"] = "normal"
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0

# Disable all inputs while encrypting/decrypting
def disableAllInputs():
	clearInput["state"] = "disabled"
	outputInput["state"] = "disabled"
	passwordInput["state"] = "disabled"
	cPasswordInput["state"] = "disabled"
	cPasswordString.set("Confirm password:")
	metadataFrame.config(bg="#e5eaf0")
	metadataInput.config(bg="#fbfcfc")
	metadataInput.config(fg="#666666")
	metadataInput["state"] = "disabled"
	startBtn["state"] = "disabled"
	eraseBtn["state"] = "disabled"
	keepBtn["state"] = "disabled"
	rsBtn["state"] = "disabled"

def prepareRsc():
	global rs13,rs128
	rs13 = RSCodec(13)
	rs128 = RSCodec(128)

# Prepare Reed-Solomon codecs
Thread(target=prepareRsc,daemon=True).start()

# Start tkinter
tk.mainloop()
