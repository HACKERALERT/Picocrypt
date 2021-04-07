#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""

Picocrypt v1.12
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

"""

# Import dependencies
from threading import Thread
from datetime import datetime,timedelta
from argon2.low_level import hash_secret_raw
from argon2.low_level import Type as argonType
from Crypto.Cipher import ChaCha20_Poly1305
from Crypto.Hash import SHA3_512
from blake3 import blake3
from hmac import compare_digest
from creedsolo import RSCodec,ReedSolomonError
from os import urandom,fsync,remove,system
from os.path import getsize,expanduser,isdir,exists
from os.path import basename,dirname,abspath,realpath
from os.path import join as pathJoin,split as pathSplit
from pathlib import Path
from zipfile import ZipFile
from tkinterdnd2 import TkinterDnD,DND_FILES
from tkinter.filedialog import asksaveasfilename
from ttkthemes import ThemedStyle
from time import sleep
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
startTime = False
previousTime = False
done = False
stopUpdating = False
reedsolo = False
reedsoloFixed = False
reedsoloErrors = False

# A list of strings and notices
strings = [
	"File metadata (optional):",
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
	"The input file's metadata is corrupted.",
	"Prevent corruption using Reed-Solomon",
	"Error. Folder(s) and/or file(s) are empty.",
	"Unknown error occured. Please try again.",
	"Drag and drop file(s) and folder(s) into this window.",
	"File metadata (read-only):",
	"Error. The input file couldn't be decoded as UTF-8.",
	"Error. The input file is corrupted so badly that Picocrypt can't decrypt it."
]

# Create root window
tk = TkinterDnD.Tk()
tk.geometry("480x500")
tk.title("Picocrypt")
tk.resizable(0,0)
tk.configure(background="#f5f6f7")

# Set "arc" theme
ThemedStyle(tk).set_theme("arc")

# Disable high DPI on Windows to prevent ugly scaling
try:
	from ctypes import windll
	windll.shcore.SetProcessDpiAwareness(0)
except:
	pass

# Try setting window icon if it exists
try:
	# PyInstaller bundles files in _MEIPASS
	from sys import _MEIPASS
	favicon = tkinter.PhotoImage(file=pathJoin(_MEIPASS,"key.png"))
	tk.iconphoto(False,favicon)
except:
	try:
		# PyInstaller not used, load from same directory
		favicon = tkinter.PhotoImage(file="key.png")
		tk.iconphoto(False,favicon)
	except:
		pass

# Dummy button used for removing ugly highlights
dummy = tkinter.ttk.Button(tk)
dummy.place(x=480,y=0)

# Label that shows the input file(s)
inputString = tkinter.StringVar(tk)
inputString.set(strings[18])
inputLabel = tkinter.ttk.Label(
	tk,
	textvariable=inputString
)
inputLabel.place(x=20,y=18)

# Clear input file(s)
clearInput = tkinter.ttk.Button(
	tk,
	text="Clear",
	command=lambda:[resetUI(),statusString.set("Ready.")]
)
clearInput.place(x=386,y=13,width=74,height=27)
clearInput["state"] = "disabled"

# Separator for aesthetics
tkinter.Frame(
	tk,
	bg="#dfe3ed",
	height=1
).place(x=20,y=39,width=438)

# Label to ask user what to save output as
outputString = tkinter.StringVar(tk)
outputString.set("Save output as:")
outputLabel = tkinter.ttk.Label(
	tk,
	textvariable=outputString
)
outputLabel.place(x=20,y=51)
outputLabel["state"] = "disabled"

# A ".pcv" extension shown next to output box
pcvString = tkinter.StringVar(tk)
pcvString.set(".pcv")
pcvLabel = tkinter.ttk.Label(
	tk,
	textvariable=pcvString
)
pcvLabel.place(x=314,y=71)

# A frame to allow output box to fill width
outputFrame = tkinter.Frame(
	tk,
	width=320,
	height=24
)
outputFrame.place(x=20,y=69)
outputFrame.columnconfigure(0,weight=10)
outputFrame.grid_propagate(False)

# Output box to allow user to change output name and path
outputInput = tkinter.ttk.Entry(outputFrame)
outputInput.grid(sticky="nesw")
outputInput["state"] = "disabled"

# A "or" label between output box and "Save as" button
orString = tkinter.StringVar(tk)
orString.set("or")
orLabel = tkinter.ttk.Label(
	tk,
	textvariable=orString
)
orLabel.place(x=356,y=71)
orLabel["state"] = "disabled"

# Save as custom file
def saveAs():
	global mode,onlyFiles,onlyFolders
	dummy.focus()

	# Get the root directory to ask user to save as
	if inputFile:
		saveDir = dirname(inputFile)
	elif onlyFiles:
		saveDir = dirname(onlyFiles[0])
	else:
		saveDir = Path(onlyFolders[0]).parent.absolute()
	
	# Ask user to save file
	tmp = asksaveasfilename(
		initialdir=saveDir,
		initialfile=(
			basename(inputFile)[:-4] if mode=="decrypt" else basename(inputFile)+".pcv"
		),
		confirmoverwrite=True
	)
	
	# If a custom output is chosen, update the output box accordingly
	if tmp:
		outputInput.delete(0,tkinter.END)
		outputInput.insert(0,(tmp if mode=="decrypt" else tmp[:-4]))

# Button to allow custom output
saveAsBtn = tkinter.ttk.Button(
	tk,
	text="Save as",
	command=saveAs
)
saveAsBtn.place(x=386,y=67,width=74,height=27)
saveAsBtn["state"] = "disabled"

# Prompt user to enter password
passwordString = tkinter.StringVar(tk)
passwordString.set("Password:")
passwordLabel = tkinter.ttk.Label(
	tk,
	textvariable=passwordString
)
passwordLabel.place(x=20,y=103)
passwordLabel["state"] = "disabled"

# Allow password input to fill width
passwordFrame = tkinter.Frame(
	tk,
	width=210,
	height=24
)
passwordFrame.place(x=20,y=121)
passwordFrame.columnconfigure(0,weight=10)
passwordFrame.grid_propagate(False)

# Password input box
passwordInput = tkinter.ttk.Entry(
	passwordFrame,
	show="•"
)
passwordInput.grid(sticky="nesw")
passwordInput["state"] = "disabled"

# Toggle password visibility
def showPassword(e):
	if passwordInput.cget("show"):
		passwordInput.config(show="")
		cPasswordInput.config(show="")
	else:
		passwordInput.config(show="•")
		cPasswordInput.config(show="•")

# Eye icon to show password
passwordShowString = tkinter.StringVar(tk)
passwordShowString.set(("  " if platform.system()=="Darwin" else "")+"👁")
passwordShow = tkinter.ttk.Label(
	tk,
	textvariable=passwordShowString,
	cursor="hand2",
	font=("TkDefaultFont",14)
)
passwordShow.bind("<Button-1>",showPassword)
passwordShow["state"] = "disabled"
passwordShow.place(
	x=(235 if platform.system()=="Windows" else 230),
	y=(121 if platform.system()=="Windows" else 124)
)

# Prompt user to confirm password
cPasswordString = tkinter.StringVar(tk)
cPasswordString.set("Confirm password:")
cPasswordLabel = tkinter.ttk.Label(
	tk,
	textvariable=cPasswordString
)
cPasswordLabel.place(x=20,y=158)
cPasswordLabel["state"] = "disabled"

# Allow confirm password input to fill width
cPasswordFrame = tkinter.Frame(
	tk,
	width=210,
	height=24
)
cPasswordFrame.place(x=20,y=176)
cPasswordFrame.columnconfigure(0,weight=10)
cPasswordFrame.grid_propagate(False)

# Confirm password input box
cPasswordInput = tkinter.ttk.Entry(
	cPasswordFrame,
	show="•"
)
cPasswordInput.grid(sticky="nesw")
cPasswordInput["state"] = "disabled"

# Show strength of password
def showStrength():
	global mode
	if mode=="decrypt":
		return
	password = passwordInput.get()
	containsLetters = any(i.isalpha() for i in password)
	containsNumbers = any(i.isdigit() for i in password)
	containsSymbols = any(not i.isalnum() for i in password)
	longEnough = len(password)>8
	if containsLetters and containsNumbers and containsSymbols and longEnough:
		passwordStrength.config(width=208)
		passwordStrength.config(bg="#149414")
	elif containsLetters and containsNumbers and containsSymbols:
		passwordStrength.config(width=140)
		passwordStrength.config(bg="#fada52")
	elif containsLetters and containsNumbers or \
		(containsLetters and containsSymbols) or \
		(containsNumbers and containsSymbols):
		passwordStrength.config(width=90)
		passwordStrength.config(bg="#ff781f")
	elif not password:
		passwordStrength.config(width=208)
		passwordStrength.config(bg="#e5eaf0")
	else:
		passwordStrength.config(width=20)
		passwordStrength.config(bg="#e3242b")

# Check if passwords match
def doPasswordsMatch():
	global mode
	if mode=="decrypt":
		if passwordInput.get():
			startBtn["state"] = "normal"
			startBtn.config(cursor="hand2")
		else:
			startBtn["state"] = "disabled"
			startBtn.config(cursor="")
	else:
		matches = passwordInput.get()==cPasswordInput.get()

		# Passwords match
		if passwordInput.get() and matches:
			passwordMatchesString.set("✔️")
			startBtn["state"] = "normal"
			startBtn.config(cursor="hand2")

		# Passwords don't match
		elif passwordInput.get() and not matches:
			passwordMatchesString.set("❌")
			startBtn["state"] = "disabled"
			startBtn.config(cursor="")

		# Password is empty
		elif not passwordInput.get():
			passwordMatchesString.set("")
			startBtn["state"] = "disabled"
			startBtn.config(cursor="")

# Bind key releases to appropriate functions
passwordInput.bind("<KeyRelease>",lambda e:[showStrength(),doPasswordsMatch()])
cPasswordInput.bind("<KeyRelease>",lambda e:doPasswordsMatch())

# Show indicator of password strength
passwordStrength = tkinter.Frame(
	tk,
	height=2,
	width=208
)
passwordStrength.config(bg="#e5eaf0")
passwordStrength.place(x=21,y=146)

# Label that indicates if passwords match
passwordMatchesString = tkinter.StringVar(tk)
passwordMatches = tkinter.ttk.Label(
	tk,
	textvariable=passwordMatchesString
)
passwordMatches.place(x=236,y=180)

# Prompt user for optional metadata
metadataString = tkinter.StringVar(tk)
metadataString.set(strings[0])
metadataLabel = tkinter.ttk.Label(
	tk,
	textvariable=metadataString
)
metadataLabel.place(x=20,y=210)
metadataLabel["state"] = "disabled"

# Frame so metadata box can fill width
metadataFrame = tkinter.Frame(
	tk,
	width=235,
	height=99
)
metadataFrame.place(x=20,y=228)
metadataFrame.columnconfigure(0,weight=10)
metadataFrame.rowconfigure(0,weight=10)
metadataFrame.grid_propagate(False)
metadataFrame.config(bg="#e5eaf0")

# Metadata text box
metadataInput = tkinter.scrolledtext.ScrolledText(
	metadataFrame,
	exportselection=0,
	padx=5,
	pady=5
)
metadataInput.config(font=("Consolas",10))
metadataInput.grid(row=0,column=0,sticky="nesw",padx=1,pady=1)
metadataInput.config(borderwidth=0)
metadataInput.config(bg="#fbfcfc")
metadataInput["state"] = "disabled"

# Tkinter's Text() boxes are ugly, so I (painfully) beautify it manually
metadataInput.bind(
	"<FocusIn>",
	lambda e:metadataBoxUI("in")
)
metadataInput.bind(
	"<FocusOut>",
	lambda e:metadataBoxUI("out")
)
def metadataBoxUI(what):
	if what=="in":
		if metadataInput.cget("bg")=="#ffffff":
			metadataFrame.config(bg="#5294e2")
	else:
		metadataFrame.config(bg="#d8ddea")

# Another separator for aesthetics
tkinter.Frame(
	tk,
	bg="#dfe3ed",
	width=1
).place(x=262,y=106,height=221)

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
keepBtn.place(x=17,y=337)
keepBtn["state"] = "disabled"

# Check box for securely erasing original file(s)
erase = tkinter.IntVar(tk)
eraseBtn = tkinter.ttk.Checkbutton(
	tk,
	text=strings[11],
	variable=erase,
	onvalue=1,
	offvalue=0,
	command=lambda:dummy.focus()
)
eraseBtn.place(x=17,y=357)
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
rsBtn.place(x=17,y=377)
rsBtn["state"] = "disabled"

# "Reed-Solomon" which links to Wikipedia
rsHelpString = tkinter.StringVar(tk)
rsHelpString.set("(?)")
rsHelp = tkinter.ttk.Label(
	tk,
	textvariable=rsHelpString,
	cursor="hand2",
	font=("Helvetica",(7 if platform.system()=="Windows" else 9))
)
rsHelp.place(
	x=(259 if platform.system()=="Windows" else 289),
	y=(382 if platform.system()=="Windows" else 384)
)
rsHelpLink = "https://en.wikipedia.org/wiki/Reed%E2%80%93Solomon_error_correction"
rsHelp.bind("<Button-1>",lambda e:webbrowser.open(rsHelpLink))

# Frame so start and cancel button can fill width
startFrame = tkinter.Frame(
	tk,
	width=440,
	height=(29 if platform.system()=="Windows" else 30)
)
startFrame.place(x=20,y=410)
startFrame.columnconfigure(0,weight=10)
startFrame.grid_propagate(False)
startFrame.config(background="#ffffff")

# Start button
startBtn = tkinter.ttk.Button(
	startFrame,
	text="Start",
	command=lambda:Thread(target=wrapper,daemon=True).start()
)
startBtn.grid(row=0,column=0,stick="nesw")
startBtn["state"] = "disabled"

# Cancel encrypting/decrypting
def cancel():
	global working
	working = False

# Cancel button
cancelBtn = tkinter.ttk.Button(
	startFrame,
	text="Cancel",
	command=cancel
)
cancelBtn.grid(stick="nesw")
cancelBtn.grid(row=0,column=1)
cancelBtn["state"] = "disabled"

# Progress bar
progress = tkinter.ttk.Progressbar(
	tk,	
	orient=tkinter.HORIZONTAL,
	length=(336 if platform.system()=="Windows" else 304),
	mode="determinate"
)
progress.place(x=30,y=420)

# Lift start and cancel button in front of progress bar
startFrame.lift()

# Status label
statusString = tkinter.StringVar(tk)
statusString.set("Ready.")
status = tkinter.ttk.Label(
	tk,
	textvariable=statusString
)
status.place(x=20,y=448)

# Credits
hint = "Created by Evan Su. Click for details and source."
creditsString = tkinter.StringVar(tk)
creditsString.set(hint)
creditsLabel = tkinter.ttk.Label(
	tk,
	textvariable=creditsString,
	cursor="hand2"
)
creditsLabel.place(x=20,y=468)
source = "https://github.com/HACKERALERT/Picocrypt"
creditsLabel.bind("<Button-1>",lambda e:webbrowser.open(source))
creditsLabel["state"] = "disabled"

# Version
versionString = tkinter.StringVar(tk)
versionString.set("v1.12")
version = tkinter.ttk.Label(
	tk,
	textvariable=versionString
)
version["state"] = "disabled"
version.place(x=(430 if platform.system()=="Windows" else 423),y=468)

# Drag files indicator window
prompt = tkinter.Frame(tk)
prompt.config(bg="#f5f6f7")
promptString = tkinter.StringVar(tk)
promptString.set("Drag and drop file(s) and folder(s) here.")
promptLabel = tkinter.ttk.Label(
	prompt,
	textvariable=promptString
)
promptLabel.place(x=135,y=311)
promptIconHor = tkinter.Frame(
	prompt,
	bg="#6f737d",
	height=4
)
promptIconHor.place(x=208,y=261,width=64)
promptIconVer = tkinter.Frame(
	prompt,
	bg="#6f737d",
	width=4
)
promptIconVer.place(x=238,y=231,height=64)

# Ask user to confirm overwrite
confirmOverwrite = tkinter.Frame(tk)
confirmOverwrite.config(bg="#f5f6f7")
confirmOverwriteString = tkinter.StringVar(tk)
confirmOverwriteString.set(strings[13])
confirmOverwriteLabel = tkinter.ttk.Label(
	confirmOverwrite,
	textvariable=confirmOverwriteString
)
confirmOverwriteLabel.place(x=91,y=200)
confirmOverwriteNo = tkinter.ttk.Button(
	confirmOverwrite,
	text="No",
	cursor="hand2",
	command=lambda:confirmOverwrite.pack_forget()
)
confirmOverwriteNo.place(x=147,y=245)

# Start encryption/decryption after user confirms
def overwriteConfirmed():
	confirmOverwrite.pack_forget()
	Thread(target=wrapper,daemon=True,args=(True,)).start()

confirmOverwriteYes = tkinter.ttk.Button(
	confirmOverwrite,
	text="Yes",
	cursor="hand2",
	command=overwriteConfirmed
)
confirmOverwriteYes.place(x=249,y=245)

# Files have been dragged
def filesDragged(draggedFiles):
	global inputFile,rs128,onlyFiles,mode,onlyFolders,allFiles
	resetUI()
	status.config(cursor="")
	status.bind("<Button-1>",lambda e:None)

	# Use try to catch file errors
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
			metadataString.set("File metadata (read only):")
			metadataInput["state"] = "normal"
			metadataInput.delete("1.0",tkinter.END)
			try:
				metadataLength = fin.read(138)
				metadataLength = bytes(rs128.decode(metadataLength)[0])
				metadataLength = metadataLength.replace(b"+",b"")
				metadata = fin.read(int(metadataLength.decode("utf-8")))
				metadata = bytes(rs128.decode(metadata)[0]).decode("utf-8")
				metadataInput.insert("1.0",metadata)
			except:
				# Metadata is corrupted, tell user
				metadataInput.insert("1.0",strings[14])
			metadataInput["state"] = "disabled"
			fin.close()
			
			# Insert filename into output box
			outputFrame.config(width=440)
			outputInput["state"] = "normal"
			outputInput.delete(0,tkinter.END)
			outputInput.insert(0,inputFile[:-4])
			
			# Update UI
			setDecryptionUI()
			startBtn["state"] = "disabled"
			startBtn.config(cursor="")
		else:
			mode = "encrypt"

			# Update UI
			setEncryptionUI()
			startBtn["state"] = "disabled"
			startBtn.config(cursor="")

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
				inputFile = tmp
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
			
		prompt.pack_forget()
		statusString.set("Ready.")
	
	# UTF-8 decode error
	except UnicodeDecodeError:
		statusString.set(strings[20])
	
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
def onDropEnter(e):
	prompt.pack(expand=1,fill=tkinter.BOTH)
	prompt.lift()
def onDropLeave(e):
	prompt.pack_forget()

# Bind functions to window
tk.drop_target_register(DND_FILES)
tk.dnd_bind("<<Drop>>",onDrop)
tk.dnd_bind("<<DropEnter>>",onDropEnter)
tk.dnd_bind("<<DropLeave>>",onDropLeave)

# Main encryption/decryption function
def work():
	global inputFile,outputFile,working,mode,rs13,rs128,reedsolo
	global done,stopUpdating,startTime,previousTime,onlyFiles
	global onlyFolders,allFiles,reedsoloFixed,reedsoloErrors
	disableAllInputs()
	dummy.focus()

	# Set and get some variables
	kept = False
	shouldKeep = keep.get()==1
	shouldErase = erase.get()==1
	reedsolo = rs.get()==1
	working = True
	stopUpdating = False
	headerBroken = False
	reedsoloFixed = 0
	reedsoloErrors = 0
	password = passwordInput.get().encode("utf-8")
	metadata = metadataInput.get("1.0",tkinter.END).encode("utf-8")
	cancelBtn["state"] = "normal"
	cancelBtn.config(cursor="hand2")
	
	# Decide if encrypting or decrypting
	if mode=="encrypt":
		outputFile = outputInput.get()+".pcv"
	else:
		outputFile = outputInput.get()
		
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
	
	# If encrypting, generate values for encryption
	if mode=="encrypt":
		salt = urandom(16) # Argon2 salt
		nonce = urandom(24) # XChaCha20 nonce
		fout = open(outputFile,"wb+")
		
		# Indicate Reed-Solomon with "+"
		if reedsolo:
			fout.write(rs128.encode(b"+"))
		else:
			fout.write(rs128.encode(b"-"))

		# Encode metadata and length of metadata
		metadata = rs128.encode(metadata)
		tmp = len(metadata)
		tmp = f"{tmp:+<10}"
		tmp = rs128.encode(tmp.encode("utf-8"))
		
		# Write to file
		fout.write(tmp)
		fout.write(metadata)
		fout.write(rs128.encode(salt)) # Argon2 salt
		fout.write(rs128.encode(nonce)) # XChaCha20 nonce
		fout.write(b"0"*192) # Hash of key
		fout.write(b"0"*144) # Poly1305 MAC
		fout.write(b"0"*160) # BLAKE3 CRC
	
	# If decrypting, read values from file
	else:
		tmp = fin.read(129)
		try:
			if bytes(rs128.decode(tmp)[0])==b"+":
				reedsolo = True
			else:
				reedsolo = False
		except:
			setDecryptionUI()
			statusString.set(strings[21])
			return

		metadataLength = fin.read(138)
		metadataLength = bytes(rs128.decode(metadataLength)[0])
		metadataLength = metadataLength.replace(b"+",b"")
		fin.read(int(metadataLength.decode("utf-8")))

		# Read values
		salt = fin.read(144)
		nonce = fin.read(152)
		keycs = fin.read(192)
		maccs = fin.read(144)
		crccs = fin.read(160)
		
		# Try to decode each value, increase Reed-Solomon errors fixed if needed
		try:
			salt,_,fixed = rs128.decode(salt)
			salt = bytes(salt)
			reedsoloFixed += len(fixed)
		except:
			headerBroken = True
			salt = salt[:16]
		try:
			nonce,_,fixed = rs128.decode(nonce)
			nonce = bytes(nonce)
			reedsoloFixed += len(fixed)
		except:
			headerBroken = True
			nonce = nonce[:24]
		try:
			keycs,_,fixed = rs128.decode(keycs)
			keycs = bytes(keycs)
			reedsoloFixed += len(fixed)
		except:
			headerBroken = True
			keycs = keycs[:64]
		try:
			maccs,_,fixed = rs128.decode(maccs)
			maccs = bytes(maccs)
			reedsoloFixed += len(fixed)
		except:
			headerBroken = True
			maccs = maccs[:16]
		try:
			crccs,_,fixed = rs128.decode(crccs)
			crccs = bytes(crccs)
			reedsoloFixed += len(fixed)
		except:
			headerBroken = True
			crccs = crccs[:32]
		
		# If the header is broken...
		if headerBroken:
			# Stop if user chose not to keep broken output
			if not shouldKeep:
				statusString.set(strings[4])
				fin.close()
				try:
					remove(outputFile)
				except:
					pass
				setDecryptionUI()
				return
			else:
				kept = "badlyCorrupted"
				
	statusString.set(strings[9])
	
	# Generate Argon2d key from master password
	key = hash_secret_raw(
		password,
		salt,
		time_cost=8,
		memory_cost=2**20,
		parallelism=8,
		hash_len=32,
		type=argonType.D
	)
	
	# Stop the indeterminate progress bar and set determinate
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0
	
	# Hash of the derived Argon2 key
	check = SHA3_512.new(data=key).digest()
	
	# Check if password is correct
	if mode=="decrypt":
		if not compare_digest(check,keycs):
			# If header isn't broken...
			if not headerBroken:
				# Tell user password is incorrect
				statusString.set(strings[2])
				fin.close()
				setDecryptionUI()
				return
		fout = open(outputFile,"wb+")

	crc = blake3() # Blake3 CRC
	cipher = ChaCha20_Poly1305.new(key=key,nonce=nonce) # XChaCha20
	
	# Variables for calculating speeds, etc.
	done = 0
	total = getsize(inputFile)
	startTime = datetime.now()
	previousTime = datetime.now()
	
	# Update progress bar, etc. in another thread
	Thread(target=updateStats,daemon=True,args=(total,)).start()
	
	# Start the encryption/decryption process
	while True:
		# Check if cancel button pressed
		if not working:
			fin.close()
			fout.close()
			remove(outputFile)
			if mode=="encrypt":
				setEncryptionUI()
			else:
				setDecryptionUI()
			statusString.set("Operation canceled by user.")
			dummy.focus()
			return

		# Read from file, read extra if Reed-Solomon was enabled
		if mode=="decrypt" and reedsolo:
			piece = fin.read(1104905)
		else:
			piece = fin.read(2**20)
		
		# End of file
		if not piece:
			break

		# Encrypt, etc.
		if mode=="encrypt":
			data = cipher.encrypt(piece)
			crc.update(data)
			if reedsolo:
				data = bytes(rs13.encode(data))
		
		# Decrypt, etc.
		else:
			if reedsolo:
				try:
					data,_,fixed = rs13.decode(piece)
				except ReedSolomonError:
					# File is really corrupted
					if not reedsoloErrors and not shouldKeep:
						stopUpdating = True
						statusString.set(strings[4])
						fin.close()
						fout.close()
						remove(outputFile)
						setDecryptionUI()
						return

					# Attempt to recover badly corrupted data
					kept = "badlyCorrupted"
					data = b""
					piece = piece[:-13]
					counter = 0
					while True:
						# Basically just strip off the Reed-Solomon bytes
						if counter<1104905:
							data += piece[counter:counter+242]
							counter += 255 # 242 bytes + 13 Reed-Solomon
						else:
							break
					fixed = bytearray()
					reedsoloErrors += 1
				
				reedsoloFixed += len(fixed)
				crc.update(data)
				data = cipher.decrypt(data)

			else:
				crc.update(piece)
				data = cipher.decrypt(piece)
		
		# Write the data, increase the amount done
		fout.write(data)
		done += 1104905 if (mode=="decrypt" and reedsolo) else 2**20
	
	# Stop UI updater from overwriting potential messages
	stopUpdating = True
	
	# Encryption is done, write appropriate values to file
	if mode=="encrypt":
		fout.flush()
		fout.close()
		fout = open(outputFile,"r+b")
		fout.seek(129+138+len(metadata)+144+152)
		fout.write(rs128.encode(check))
		fout.write(rs128.encode(cipher.digest()))
		fout.write(rs128.encode(crc.digest()))
	
	# Decryption is done, check for integrity and authenticity
	else:
		# File is corrupted
		if not compare_digest(crccs,crc.digest()):
			statusString.set(strings[3])
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
			if not reedsoloErrors and not headerBroken:
				# File is modified
				statusString.set(strings[5])
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
	if not kept:
		fout.flush()
		fsync(fout.fileno())
	fout.close()
	fin.close()

	# Securely wipe files as necessary
	if shouldErase:
		if onlyFolders:
			for i in onlyFolders:
				secureWipe(i)
		if onlyFiles:
			for i in range(len(onlyFiles)):
				statusString.set(strings[12]+f" ({i}/{len(onlyFiles)}")
				progress["value"] = i/len(onlyFiles)
				secureWipe(onlyFiles[i])
		secureWipe(inputFile)

	# Secure wipe not enabled
	else:
		# Remove temporary zip file if created
		if allFiles or onlyFiles:
			remove(inputFile)
					
	# Show appropriate notice if file corrupted or modified
	arrow = "" if platform.system()=="Darwin" else "🡪"
	if not kept:
		statusString.set(f"Completed. (Click here to show output {arrow})")
		# Show Reed-Solomon stats if it fixed corrupted bytes
		if mode=="decrypt" and reedsoloFixed:
			tmp = "s" if reedsoloFixed!=1 else ""
			statusString.set(
				f"Completed with {reedsoloFixed} byte{tmp}"+
				f" fixed. (Click here to show output {arrow})"
			)
	else:
		if kept=="modified":
			statusString.set(strings[7])
		elif kept=="corrupted":
			statusString.set(strings[6])
		else:
			statusString.set(strings[8])
	
	status.config(cursor="hand2")
	
	# A little hack to prevent reference nonsense
	output = "".join([i for i in outputFile])

	# Bind the output file to the status label
	if platform.system()=="Windows":
		status.bind("<Button-1>",lambda e:showOutput(output.replace("/","\\")))
	else:
		status.bind("<Button-1>",lambda e:showOutput(output))

	# Reset variables and UI states
	resetUI()
	inputFile = ""
	outputFile = ""
	allFiles = []
	onlyFolders = []
	onlyFiles = []
	working = False

# Wraps the work() function and prevents overwrites
def wrapper(yes=False):
	global working,mode,outputFile
	if mode=="encrypt":
		outputFile = outputInput.get()+".pcv"
	else:
		outputFile = outputInput.get()
	try:
		getsize(outputFile)
		if not yes:
			confirmOverwrite.pack(expand=1,fill=tkinter.BOTH)
			confirmOverwrite.lift()
			return
	except:
		pass
	try:
		work()
	except:
		working = False
		if mode=="encrypt":
			setEncryptionUI()
		else:
			setDecryptionUI()
		statusString.set(strings[17])
	finally:
		dummy.focus()
		sys.exit(0)

# Update visuals with stats
def updateStats(total):
	global startTime,previousTime,done,stopUpdating,reedsolo,reedsoloFixed,reedsoloErrors,working
	while True:
		validStatus = (
			statusString.get().startswith("Working") or statusString.get().startswith("Deriving")
		)
		# Make sure to stop updating if needed
		if not stopUpdating and validStatus and working:
			# Some basic math to calculate speed, ETA, etc.
			elapsed = (datetime.now()-previousTime).total_seconds() or 0.0001
			sinceStart = (datetime.now()-startTime).total_seconds() or 0.0001
			previousTime = datetime.now()
			percent = done*100/total
			progress["value"] = percent
			
			speed = (done/sinceStart)/10**6 or 0.0001
			eta = max(round((total-done)/(speed*10**6)),0)
			eta = str(timedelta(seconds=min(eta,86399))).zfill(8)
			
			info = f"Working... {min(percent,100):.0f}% at {speed:.2f} MB/s (ETA: {eta})"

			# Show the number of fixed and unrecoverable bytes
			if reedsolo and mode=="decrypt" and reedsoloFixed:
				tmp = "s" if reedsoloFixed!=1 else ""
				info += f", fixed {reedsoloFixed} byte{tmp}"

			if reedsolo and mode=="decrypt" and reedsoloErrors:
				info += f", {reedsoloErrors} MB unrecoverable"

			statusString.set(info)
			sleep(0.05)
		else:
			sys.exit(0)
			break

# Securely wipe file(s) via system internals
def secureWipe(fin):
	statusString.set(strings[12])

	# Check platform, erase accordingly
	if platform.system()=="Windows":
		# Recursively delete folders
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

	# MacOS
	elif platform.system()=="Darwin":
		system(f'rm -rfP "{fin}"')
					
	# Linux
	else:
		system(f'shred -uz "{fin}" -n 4')

# Show output file in user's file manager
def showOutput(file):
	if platform.system()=="Windows":
		system(f'explorer /select,"{file}"')
	elif platform.system()=="Darwin":
		system(f'cd "{dirname(file)}"; open -R "{pathSplit(file)[1]}"')
		system(f'cd "{rootDir}"')
	else:
		system(f'xdg-open "{dirname(file)}"')

# Reset UI to state where no files are selected
def resetUI():
	global working
	working = False
	inputString.set(strings[18])
	inputLabel["state"] = "normal"
	clearInput["state"] = "disabled"
	clearInput.config(cursor="")
	outputLabel["state"] = "disabled"
	saveAsBtn.config(cursor="")
	saveAsBtn["state"] = "disabled"
	orLabel.place(x=356,y=71)
	orLabel["state"] = "disabled"
	outputFrame.config(width=320)
	outputInput["state"] = "normal"
	outputInput.delete(0,"end")
	outputInput["state"] = "disabled"
	pcvLabel["state"] = "normal"
	passwordLabel["state"] = "disabled"
	passwordInput["state"] = "normal"
	passwordInput.delete(0,"end")
	passwordInput["state"] = "disabled"
	passwordInput.config(show="•")
	passwordShow["state"] = "disabled"
	cPasswordString.set("Confirm password:")
	cPasswordLabel["state"] = "disabled"
	cPasswordInput["state"] = "normal"
	cPasswordInput.delete(0,"end")
	cPasswordInput["state"] = "disabled"
	cPasswordInput.config(show="•")
	passwordStrength.config(width=208)
	passwordStrength.config(bg="#e5eaf0")
	passwordMatchesString.set("")
	passwordMatches["state"] = "normal"
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
	startFrame.lift()
	startBtn["state"] = "disabled"
	startBtn.config(cursor="")
	cancelBtn["state"] = "disabled"
	cancelBtn.config(cursor="")
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0
	dummy.focus()

# Set UI to encryption state
def setEncryptionUI():
	global working
	working = False
	clearInput["state"] = "normal"
	clearInput.config(cursor="hand2")
	saveAsBtn.config(cursor="hand2")
	saveAsBtn["state"] = "normal"
	orLabel.place(x=353,y=71)
	orLabel["state"] = "normal"
	outputLabel["state"] = "normal"
	outputInput["state"] = "normal"
	outputFrame.config(width=290)
	pcvLabel["state"] = "normal"
	passwordLabel["state"] = "normal"
	passwordInput["state"] = "normal"
	passwordShow["state"] = "normal"
	showStrength()
	cPasswordLabel["state"] = "normal"
	cPasswordString.set("Confirm password:")
	cPasswordInput["state"] = "normal"
	passwordMatches["state"] = "normal"
	metadataFrame.config(bg="#cfd6e6")
	metadataInput.config(bg="#ffffff")
	metadataInput.config(fg="#000000")
	metadataLabel["state"] = "normal"
	metadataInput["state"] = "normal"
	eraseBtn["state"] = "normal"
	rsBtn["state"] = "normal"
	startFrame.lift()
	startBtn["state"] = "normal"
	startBtn.config(cursor="hand2")
	cancelBtn["state"] = "disabled"
	cancelBtn.config(cursor="")
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0
	
# Set UI to decryption state
def setDecryptionUI():
	global working
	working = False
	clearInput["state"] = "normal"
	clearInput.config(cursor="hand2")
	saveAsBtn.config(cursor="hand2")
	saveAsBtn["state"] = "normal"
	orLabel.place(x=356,y=71)
	orLabel["state"] = "normal"
	outputLabel["state"] = "normal"
	outputInput["state"] = "normal"
	outputFrame.config(width=320)
	passwordLabel["state"] = "normal"
	passwordInput["state"] = "normal"
	passwordShow["state"] = "normal"
	cPasswordString.set("Confirm password (N/A):")
	metadataFrame.config(bg="#e5eaf0")
	metadataInput.config(bg="#fbfcfc")
	metadataInput.config(fg="#666666")
	metadataString.set(strings[19])
	metadataInput["state"] = "disabled"
	keepBtn["state"] = "normal"
	startFrame.lift()
	startBtn["state"] = "normal"
	startBtn.config(cursor="hand2")
	cancelBtn["state"] = "disabled"
	cancelBtn.config(cursor="")
	progress.stop()
	progress.config(mode="determinate")
	progress["value"] = 0

# Disable all inputs while encrypting/decrypting
def disableAllInputs():
	clearInput["state"] = "disabled"
	clearInput.config(cursor="")
	saveAsBtn.config(cursor="")
	saveAsBtn["state"] = "disabled"
	orLabel["state"] = "disabled"
	outputInput["state"] = "disabled"
	pcvLabel["state"] = "disabled"
	passwordInput["state"] = "disabled"
	passwordInput.config(show="•")
	passwordShow["state"] = "disabled"
	passwordStrength.config(width=208)
	passwordStrength.config(bg="#e5eaf0")
	cPasswordInput["state"] = "disabled"
	cPasswordInput.config(show="•")
	cPasswordString.set("Confirm password:")
	passwordMatches["state"] = "disabled"
	metadataFrame.config(bg="#e5eaf0")
	metadataInput.config(bg="#fbfcfc")
	metadataInput.config(fg="#666666")
	metadataInput["state"] = "disabled"
	progress.lift()
	startBtn["state"] = "disabled"
	startBtn.config(cursor="")
	eraseBtn["state"] = "disabled"
	keepBtn["state"] = "disabled"
	rsBtn["state"] = "disabled"

# Allow window to close if not encrypting/decrypting
def onClose():
	global working
	if not working:
		tk.destroy()

# Generate Reed-Solomon codecs
def prepare():
	global rs13,rs128
	rs13 = RSCodec(13)
	rs128 = RSCodec(128)
	
	# Accept sdelete64's EULA
	if platform.system()=="Windows":
		system("sdelete64.exe /accepteula")
	sys.exit(0)

# Prepare Reed-Solomon codecs
Thread(target=prepare,daemon=True).start()

# Bind tkinter close event
tk.protocol("WM_DELETE_WINDOW",onClose)

# Start tkinter
tk.mainloop()
sys.exit(0)
