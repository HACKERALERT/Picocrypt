#!/usr/bin/env python3

# Dependencies: argon2-cffi, pycryptodome
# Copyright (c) Evan Su (https://evansu.cc)
# Released under a GNU GPL v3 license
# Source: https://github.com/HACKERALERT/Picocrypt

try:
    from argon2.low_level import hash_secret_raw
    from Crypto.Cipher import ChaCha20_Poly1305
except:
    from os import system
    system("python3 -m pip install argon2-cffi")
    system("python3 -m pip install pycryptodome")

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

inputFile = ""
outputFile = ""
password = ""
ad = ""
adString = "File metadata (used to store some text along with the file):"
passwordNotice = "Error. The provided password is incorrect."
corruptedNotice = "Error. The input file is corrupted."
modifiedNotice = "Error. The input file has been intentionally modified."
kCorruptedNotice = "The input file is corrupted, but the output has been kept."
kModifiedNotice = "The input file has been intentionally modified, but the output has been kept."
derivingNotice = "Deriving key... (takes a few seconds)"
keepNotice = "Keep decrypted output even if it's corrupted or modified"
kept = False
eraseNotice = "Securely erase and delete original file"
working = False
overwriteNotice = "Output file already exists. Would you like to overwrite it?"
unknownErrorNotice = "Unknown error occured. Please try again or view the logs."
log = ""

tk = tkinter.Tk()
tk.geometry("480x420")
tk.title("Picocrypt")
tk.configure(background="#f5f6f7")
tk.resizable(0,0)

favicon = tkinter.PhotoImage(file="./key.png")
tk.iconphoto(False,favicon)

s = tkinter.ttk.Style()
s.configure("TCheckbutton",background="#f5f6f7")

def inputSelected():
    global inputFile,working,log
    dummy.focus()
    log += "File selection dialog opened\n"
    try:
        suffix = ""
        tmp = filedialog.askopenfilename(
            initialdir=expanduser("~")
        )
        if len(tmp)==0:
            raise Exception("No file selected.")
        inputFile = tmp
        if ".pcf" in inputFile.split("/")[-1]:
            suffix = " (will be decrypted)"
            fin = open(inputFile,"rb+")
            adlen = b""
            while True:
                letter = fin.read(1)
                adlen += letter
                if letter==b"|":
                    adlen = adlen[:-1]
                    break
            ad = fin.read(int(adlen.decode("utf-8")))
            fin.close()
            adArea["state"] = "normal"
            adArea.delete("1.0",tkinter.END)
            adArea.insert("1.0",ad.decode("utf-8"))
            adArea["state"] = "disabled"
            adLabelString.set("File metadata (read only):")
            keepBtn["state"] = "normal"
            eraseBtn["state"] = "disabled"
            log += "File selected, will decrypt\n"
        else:
            eraseBtn["state"] = "normal"
            keepBtn["state"] = "disabled"
            adArea["state"] = "normal"
            adArea.delete("1.0",tkinter.END)
            suffix = " (will be encrypted)"
            adLabelString.set(adString)
            log += "File selected, will encrypt\n"
        inputString.set(inputFile.split("/")[-1]+suffix)
        passwordInput["state"] = "normal"
        startBtn["state"] = "normal"
        statusString.set("Ready.")
        progress["value"] = 0
    except UnicodeDecodeError as e:
        log += str(e)+"\n"
        statusString.set(corruptedNotice)
    except Exception as e:
        log += str(e)+"\n"
    finally:
        dummy.focus()
        working = False

selectFileInput = tkinter.ttk.Button(
    tk,
    text="Select file",
    command=inputSelected,
)
selectFileInput.place(x=19,y=20)

inputString = tkinter.StringVar(tk)
inputString.set("Please select a file.")
selectedInput = tkinter.ttk.Label(
    tk,
    textvariable=inputString
)
selectedInput.config(background="#f5f6f7")
selectedInput.place(x=104,y=23)

passwordString = tkinter.StringVar(tk)
passwordString.set("Password:")

passwordLabel = tkinter.ttk.Label(
    tk,
    textvariable=passwordString
)
passwordLabel.place(x=17,y=56)
passwordLabel.config(background="#f5f6f7")

passwordFrame = tkinter.Frame(
    tk,
    width=440,
    height=22
)
passwordFrame.place(x=20,y=76)
passwordFrame.columnconfigure(0,weight=10)
passwordFrame.grid_propagate(False)

passwordInput = tkinter.ttk.Entry(
    passwordFrame
)
passwordInput.grid(sticky="nesw")
passwordInput["state"] = "disabled"

def start():
    global inputFile,outputFile,password,ad,kept,working,log

    log += "Starting the encryption/decryption process\n"

    if ".pcf" not in inputFile:
        mode = "encrypt"
        outputFile = inputFile+".pcf"
    else:
        mode = "decrypt"
        outputFile = inputFile[:-4]
    try:
        getsize(outputFile)
        force = messagebox.askyesno("Warning",overwriteNotice)
        log += "Overwrite set to true\n"
        dummy.focus()
        if force!=1:
            log += "Cancelled because overwrite set to false\n"
            return
    except:
        pass

    working = True
    dummy.focus()
    password = passwordInput.get().encode("utf-8")
    ad = adArea.get("1.0",tkinter.END).encode("utf-8")
    wipe = erase.get()==1
    if wipe and mode=="encrypt":
        log += "Secure wipe enabled\n"
    elif not wipe and mode=="encrypt":
        log += "Secure wipe disabled\n"

    selectFileInput["state"] = "disabled"
    passwordInput["state"] = "disabled"
    adArea["state"] = "disabled"
    startBtn["state"] = "disabled"
    keepBtn["state"] = "disabled"

    fin = open(inputFile,"rb+")
    fout = open(outputFile,"wb+")

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
    else:
        adlen = b""
        while True:
            letter = fin.read(1)
            adlen += letter
            if letter==b"|":
                adlen = adlen[:-1]
                break
        fin.read(int(adlen.decode("utf-8")))
        cs = fin.read(64)
        crccs = fin.read(64)
        digest = fin.read(16)
        salt = fin.read(16)
        nonce = fin.read(24)

    statusString.set(derivingNotice)

    progress.config(mode="indeterminate")
    progress.start(15)

    log += "Generating key through Argon2\n"

    key = hash_secret_raw(
        password,
        salt,
        time_cost=8,
        memory_cost=1048576,
        parallelism=8,
        hash_len=32,
        type=Type.ID
    )

    progress.stop()
    progress.config(mode="determinate")
    progress["value"] = 0

    check = sha3_512(key).digest()

    if mode=="decrypt":
        log += "Checking if key is correct"
        if not compare_digest(check,cs):
            log += "\nIncorrect password\n"
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
        log += " (yes)\n"

    cipher = ChaCha20_Poly1305.new(key=key,nonce=nonce)
    crc = sha3_512()

    done = 0
    total = getsize(inputFile)
    chunkSize = 2**20
    startTime = datetime.now()

    if wipe:
        wiper = open(inputFile,"r+b")
        wiper.seek(0)

    log += "Encryption/decryption starting\n"
    while True:
        piece = fin.read(chunkSize)
        if wipe:
            trash = urandom(len(piece))
            wiper.write(trash)
        if not piece:
            if mode=="encrypt":
                digest = cipher.digest()
                fout.flush()
                fout.close()
                fout = open(outputFile,"r+b")
                fout.seek(len(str(len(ad)))+1+len(ad))
                fout.write(check)
                fout.write(crc.digest())
                fout.write(digest)
            else:
                crcdg = crc.digest()
                if not compare_digest(crccs,crcdg):
                    log += "Data is corrupted\n"
                    statusString.set(corruptedNotice)
                    progress["value"] = 100
                    fin.close()
                    fout.close()
                    if keep.get()!=1:
                        log += "Corrupted output has been kept\n"
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
                    cipher.verify(digest)
                except:
                    log += "Data has been modified\n"
                    statusString.set(modifiedNotice)
                    progress["value"] = 100
                    fin.close()
                    fout.close()
                    if keep.get()!=1:
                        log += "Modified output has been kept\n"
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
        
        if mode=="encrypt":
            data = cipher.encrypt(piece)
            crc.update(data)
        else:
            crc.update(piece)
            data = cipher.decrypt(piece)

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
    del fin,fout,cipher,key
    log += "Process completed\n"

def wrapper():
    global working,log
    try:
        start()
    except Exception as e:
        selectFileInput["state"] = "normal"
        passwordInput["state"] = "normal"
        adArea["state"] = "normal"
        startBtn["state"] = "normal"
        keepBtn["state"] = "normal"
        statusString.set(unknownErrorNotice)
        dummy.focus()
        working = False
        log += str(e)+"\n"
    finally:
        sys.exit(0)
    
def startWorker():
    thread = Thread(target=wrapper,daemon=True)
    thread.start()

adLabelString = tkinter.StringVar(tk)
adLabelString.set(adString)
adLabel = tkinter.ttk.Label(
    tk,
    textvariable=adLabelString
)
adLabel.place(x=17,y=108)
adLabel.config(background="#f5f6f7")

adFrame = tkinter.Frame(
    tk,
    width=440,
    height=100
)
adFrame.place(x=20,y=128)
adFrame.columnconfigure(0,weight=10)
adFrame.grid_propagate(False)

adArea = tkinter.Text(
    adFrame,
    exportselection=0
)
adArea.config(font=("Consolas",12))

adArea.grid(sticky="we")
adArea["state"] = "disabled"

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

startFrame = tkinter.Frame(
    tk,
    width=442,
    height=25
)
startFrame.place(x=19,y=290)
startFrame.columnconfigure(0,weight=10)
startFrame.grid_propagate(False)

startBtn = tkinter.ttk.Button(
    startFrame,
    text="Start",
    command=startWorker
)
startBtn.grid(sticky="nesw")
startBtn["state"] = "disabled"

progress = tkinter.ttk.Progressbar(
    tk,
    orient=tkinter.HORIZONTAL,
    length=440,
    mode="determinate"
)
progress.place(x=20,y=328)

statusString = tkinter.StringVar(tk)
statusString.set("Ready.")
status = tkinter.ttk.Label(
    tk,
    textvariable=statusString
)
status.place(x=17,y=356)
status.config(background="#f5f6f7")

hint = "(v1.4) Created by Evan Su. Click for details and source."
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

versionString = tkinter.StringVar(tk)
versionString.set("Logs")
version = tkinter.ttk.Label(
    tk,
    textvariable=versionString,
    cursor="hand2"
)
version["state"] = "disabled"
version.config(background="#f5f6f7")
version.place(x=430,y=386)
version.bind("<Button-1>",lambda e:showLog())

def showLogWrapper():
    logger = tkinter.Tk()
    logger.geometry("480x420")
    logger.title("Logs")
    logger.resizable(0,0)
    loggerFrame = tkinter.Frame(
        logger,
        width=480,
        height=420
    )
    loggerFrame.place(x=0,y=0)
    loggerFrame.columnconfigure(0,weight=10)
    loggerFrame.grid_propagate(False)
    box = tkinter.scrolledtext.ScrolledText(
        loggerFrame,
    )
    box.config(font=("Consolas",11))
    box.grid(sticky="nesw")
    box.insert("1.0",log)
    box["state"] = "disabled"
    logger.mainloop()
    sys.exit(0)

def showLog():
    t = Thread(target=showLogWrapper,daemon=True)
    t.start()

dummy = tkinter.ttk.Button(
    tk
)
dummy.place(x=480,y=0)

def onClose():
    if not working:
        tk.destroy()

if __name__=="__main__":
    tk.protocol("WM_DELETE_WINDOW",onClose)
    tk.mainloop()
    sys.exit(0)
