#!/usr/bin/env python3

"""

Picocrypt v2.0
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
from monocypher import lock,unlock,wipe,Blake2b
from hmac import compare_digest
from reedsolo import RSCodec,ReedSolomonError
from os import urandom,fsync,remove,system
from os.path import getsize,expanduser,isdir,exists
from os.path import dirname,abspath,realpath
from os.path import join as pathJoin
from os.path import split as pathSplit
from pathlib import Path
from zipfile import ZipFile
from tkinterdnd2 import TkinterDnD,DND_FILES
from ttkthemes import ThemedStyle
import sys
import tkinter
import tkinter.ttk
import tkinter.scrolledtext
import webbrowser
import platform

# Disable high DPI to prevent layout shifts
try:
	from ctypes import windll
	windll.shcore.SetProcessDpiAwareness(0)
except:
	pass

# Global variables
rootDir = dirname(realpath(__file__))
working = False
mode = False
inputFile = False
outputFile = False
outputPath = False
kept = False
rs128 = False
rs13 = False
allFiles = False
onlyFolders = False
onlyFiles = False
filesLoaded = False

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
	"Creating Reed-Solomon tables...",
	"Unknown error occured. Please try again."
]

# Create root tk with Drag and Drop support
from hdpitkinter import HdpiTk
tk = TkinterDnD.Tk()#HdpiTk()#TkinterDnD.Tk()
tk.geometry("480x520")
tk.title("Picocrypt")
tk.resizable(0,0)
tk.configure(background="#ffffff")
ThemedStyle(tk).set_theme("arc")

# Add some styling
style = tkinter.ttk.Style()
style.configure("TCheckbutton",background="#ffffff")
style.configure("TLabel",background="#ffffff")
style.configure("TSeparator",background="#ffffff")

# Try setting window icon if it exists
try:
	favicon = tkinter.PhotoImage(file="./key.png")
	tk.iconphoto(False,favicon)
except:
	pass

# Label that shows the input files
inputString = tkinter.StringVar(tk)
inputString.set("Drag and drop file(s) and folder(s) into this window.")
inputLabel = tkinter.ttk.Label(
	tk,
	textvariable=inputString
)
inputLabel.place(x=20,y=18)

# Clear input files
clearInput = tkinter.ttk.Button(
	tk,
	text="Clear",
	cursor="hand2"
)
clearInput.place(x=400,y=12,width=60,height=28)

# Separator for aesthetics
separator = tkinter.Frame(#tkinter.ttk.Separator(tk)
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
	width=410,
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
passwordLabel.place(x=20,y=98)
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
cPasswordLabel.place(x=20,y=148)
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
metadataLabel.place(x=20,y=200)
metadataLabel["state"] = "disabled"

# Frame so metadata box can fill width
metadataFrame = tkinter.Frame(
	tk,
	width=444,
	height=99
)
metadataFrame.place(x=20,y=220)
metadataFrame.columnconfigure(0,weight=10)
metadataFrame.rowconfigure(0,weight=10)
metadataFrame.grid_propagate(False)
metadataFrame.config(bg="#dfe4ee")

# Metadata scrollbar
metadataScrollbar = tkinter.ttk.Scrollbar(
	metadataFrame,
	cursor="hand2"
)
metadataScrollbar.grid(row=0,column=1,sticky="nesw")

# Metadata text box
metadataInput = tkinter.Text(
	metadataFrame,
	exportselection=0,
	height=5,
	yscrollcommand=metadataScrollbar.set
)
metadataInput.config(font=("Consolas",12))
metadataInput.grid(row=0,column=0,sticky="nesw",padx=1,pady=1)
metadataInput.config(borderwidth=0)
metadataScrollbar.config(command=metadataInput.yview)
metadataInput["state"] = "disabled"

# Check box for keeping corrupted or modified output
keep = tkinter.IntVar(tk)
keepBtn = tkinter.ttk.Checkbutton(
	tk,
	text=strings[10],
	variable=keep,
	onvalue=1,
	offvalue=0
)
keepBtn.place(x=18,y=332)
keepBtn["state"] = "disabled"


# Check box for securely erasing original files
erase = tkinter.IntVar(tk)
eraseBtn = tkinter.ttk.Checkbutton(
	tk,
	text=strings[12],
	variable=erase,
	onvalue=1,
	offvalue=0
)
eraseBtn.place(x=18,y=352)
eraseBtn["state"] = "disabled"

# Check box for enabling Reed-Solomon anti-corruption
rs = tkinter.IntVar(tk)
rsBtn = tkinter.ttk.Checkbutton(
	tk,
	text=strings[16],
	variable=rs,
	onvalue=1,
	offvalue=0
)
rsBtn.place(x=18,y=372)
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
	text="Start"
)
startBtn.grid(row=0,column=0,stick="nesw")
#startBtn["state"] = "disabled"

# Cancel button
cancelBtn = tkinter.ttk.Button(
	startFrame,
	text="Cancel"
)
cancelBtn.grid(stick="nesw")
cancelBtn.grid(row=0,column=1)
#cancelBtn["state"] = "disabled"

# Progress bar
progress = tkinter.ttk.Progressbar(
	tk,	
	orient=tkinter.HORIZONTAL,
	length=440,
	mode="determinate"
)
progress.place(x=20,y=438)

# Status label
statusString = tkinter.StringVar(tk)
statusString.set("Ready.")
status = tkinter.ttk.Label(
	tk,
	textvariable=statusString
)
status.place(x=20,y=456)
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
credits.place(x=20,y=488)
source = "https://github.com/HACKERALERT/Picocrypt"
credits.bind("<Button-1>",lambda e:webbrowser.open(source))
credits["state"] = "disabled"

# Version
versionString = tkinter.StringVar(tk)
versionString.set("v2.0")
version = tkinter.ttk.Label(
	tk,
	textvariable=versionString
)
version["state"] = "disabled"
version.place(x=435,y=488)


if __name__=="__main__":
	tk.mainloop()
