package main

/*

Picocrypt v1.14
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

*/

import (
	// Generic
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"math"
	"time"
	"sync"
	_ "embed"
	"strings"
	"strconv"
	"runtime"
	"net/http"
	"runtime/debug"
	"image/color"
	"archive/zip"
	"encoding/hex"
	"path/filepath"

	// Reed-Solomon
	"github.com/HACKERALERT/infectious"

	// Cryptography
	"crypto/rand"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/chacha20poly1305"
	"github.com/HACKERALERT/Picocrypt/src/monocypher"

	// GUI
	"github.com/AllenDang/giu"

	// Helpers
	"github.com/HACKERALERT/clipboard"
	"github.com/HACKERALERT/dialog"
	"github.com/HACKERALERT/browser"
)

//go:embed NotoSans-Regular.ttf
var font []byte

var version = "v1.14"
var exe,_ = os.Executable()
var rootDir = filepath.Dir(exe)

// Global variables
var dpi float32 // Used to scale properly in high-resolution displays
var mode string // Either "encrypt" or "decrypt" (self-explanatory)
var working = false // True if encryption/decryption is in progress
var onlyFiles []string // Only contains files not in a folder
var onlyFolders []string // Only contains names of folders
var allFiles []string // Contains all files including files in folders
var inputFile string // Path of the input file to encrypt/decrypt
var outputFile string // Path of the output file to encrypt/decrypt to
var recombine bool // True if decrypting and the original file was splitted during encryption

// UI-related global variables
var tab = 0 // The index of the currently selected tab
var inputLabel = "Drag and drop file(s) and folder(s) into this window."
var outputEntry string // A modifiable text entry string variable containing path of output
var outputWidth float32 = 376 // How wide 'outputEntry' is in pixels
var orLabel = "or" // The text "or" separating 'outputEntry' and "Save as"
var passwordState = giu.InputTextFlagsPassword // Changes to show/hide password
var showPassword = false
var keyfile = false // True if user chooses/chose to use a keyfile
var keyfilePrompt = "Keyfile (optional):" // Changes if decrypting and keyfile was enabled
var progress float32 = 0 // 0 is 0%, 1 is 100%
var progressInfo = "" // Text inside the progress bar on the encrypting/decrypting window
var status = "Ready." // Status text in encrypting/decrypting window
var _status = "Ready." // Status text in main window
var _status_color = color.RGBA{0xff,0xff,0xff,255} // Changes according to status (success, fail, etc.)
var splitUnits = []string{
	"KB",
	"MB",
	"GB",
} // Three choosable units for splitting output file when encrypting, in powers of 2
var splitSelected int32 // Index of which splitting unit was chosen from above
var shredModes = []string{
	"Normal (4 passes)",
	"Paranoid (Still in development, don't use)",
}
var shredMode int32
var shredProgress float32 // Progress of shredding files
var shredDone float32
var shredTotal float32 // Total files to shred (recursive)
var shredOverlay string // Text in shredding progress bar
var shredding = "Ready."

// User input variables
var password string
var cPassword string // Confirm password text entry string variable
var keyfilePath string
var keyfileLabel = "Use a keyfile"
var metadata string
var keep bool
var erase bool
var reedsolo bool
var split bool
var splitSize string
var fast bool
var _fast bool // Helper variable with the original checkbox state
var kept = false // If a file was corrupted/modified, but the output was kept

// Reed-Solomon encoders
var rs5,_ = infectious.NewFEC(5,15) // 5 data shards, 15 total -> 10 parity shards
var rs10,_ = infectious.NewFEC(10,30)
var rs16,_ = infectious.NewFEC(16,48)
var rs24,_ = infectious.NewFEC(24,72)
var rs32,_ = infectious.NewFEC(32,96)
var rs64,_ = infectious.NewFEC(64,192)

// File checksum generator variables
var cs_md5 string // A string containing a hex-encoded MD5 hash
var cs_sha1 string
var cs_sha256 string
var cs_sha3_256 string
var cs_blake2b string
var cs_blake2s string
var cs_validate string
var md5_color = color.RGBA{0x10,0x10,0x10,255} // Border color that changes upon a match
var sha1_color = color.RGBA{0x10,0x10,0x10,255}
var sha256_color = color.RGBA{0x10,0x10,0x10,255}
var sha3_256_color = color.RGBA{0x10,0x10,0x10,255}
var blake2b_color = color.RGBA{0x10,0x10,0x10,255}
var blake2s_color = color.RGBA{0x10,0x10,0x10,255}
var cs_progress float32 = 0
var md5_selected = false // Whether the checkbox was checked or not
var sha1_selected = false
var sha256_selected = false
var sha3_256_selected = false
var blake2b_selected = false
var blake2s_selected = false

// Create the UI
func startUI(){
	giu.SingleWindow().Layout(
		giu.Style().SetColor(giu.StyleColorBorder,color.RGBA{0x10,0x10,0x10,255}).To(
			// The tab bar, which contains different tabs for different functions
			giu.TabBar().TabItems(
				// File encryption/decryption tab
				giu.TabItem("Encryption/decryption").Layout(
					// Update 'tab' to indicate that this is the active tab
					giu.Custom(func(){
						if giu.IsItemActive(){
							tab = 0
						}
					}),
					
					// Confirm overwrite with a modal
					giu.PopupModal("Confirmation").Layout(
						giu.Label("Output already exists. Overwrite?"),
						giu.Row(
							giu.Button("No").Size(110,0).OnClick(func(){
								giu.CloseCurrentPopup()
							}),
							giu.Dummy(-118,0),
							giu.Button("Yes").Size(110,0).OnClick(func(){
								giu.CloseCurrentPopup()
								go work()
							}),
							
						),
					),

					// Label listing the input files and a button to clear them
					giu.Row(
						giu.Label(inputLabel),
						giu.Dummy(-55,0),
						giu.Button("Clear").Size(45,0).OnClick(resetUI),
					),

					// Allow user to choose a custom output path and/or name
					giu.Label("Save output as:"),
					giu.Row(
						giu.InputText(&outputEntry).Size(outputWidth/dpi),
						giu.Label(orLabel),
						giu.Button("Save as").OnClick(func(){
							file,_ := dialog.File().Title("Save as").Save()

							// Return if user canceled the file dialog
							if file==""{
								return
							}

							// Remove the extra ".pcv" extension if necessary
							if strings.HasSuffix(file,".pcv"){
								file = file[:len(file)-4]
							}
							outputEntry = file
						}),
					),

					// Prompt for password
					giu.Row(
						giu.Label("Password:"),
						giu.Dummy(-220,0),
						giu.Label(keyfilePrompt),
					),
					giu.Row(
						giu.InputText(&password).Size(200/dpi).Flags(passwordState),
						giu.Checkbox("##showPassword",&showPassword).OnChange(func(){
							if passwordState==giu.InputTextFlagsPassword{
								passwordState = giu.InputTextFlagsNone
							}else{
								passwordState = giu.InputTextFlagsPassword
							}
							giu.Update()
						}),
						giu.Dummy(-220,0),
						giu.Checkbox(keyfileLabel,&keyfile).OnChange(func(){
							if !keyfile{
								keyfileLabel = "Use a keyfile"
								return
							}
							filename,err := dialog.File().Load()
							if err!=nil{
								keyfile = false
								return
							}
							keyfileLabel = filename
							keyfilePath = filename
						}),
					),

					// Prompt to confirm password
					giu.Label("Confirm password:"),
					giu.InputText(&cPassword).Size(200/dpi).Flags(passwordState),

					// Optional metadata
					giu.Label("Metadata (optional):"),
					giu.InputTextMultiline(&metadata).Size(230,126),

					// Advanced options can be enabled with checkboxes
					giu.Checkbox("Keep decrypted output even if it's corrupted or modified",&keep),
					giu.Checkbox("Securely shred the original file(s) and folder(s)",&erase),
					giu.Row(
						giu.Checkbox("(Not ready) Encode with Reed-Solomon to prevent corruption",&reedsolo),
						giu.Button("?").OnClick(func(){
							browser.OpenURL("https://bit.ly/reedsolomonwikipedia")
						}),
					),
					giu.Row(
						giu.Checkbox("Split output into chunks of",&split),
						giu.InputText(&splitSize).Size(30).Flags(giu.InputTextFlagsCharsDecimal),
						giu.Combo("##splitter",splitUnits[splitSelected],splitUnits,&splitSelected).Size(40),
					),
					giu.Checkbox("Fast mode (less secure, not as durable)",&fast),

					// Start button
					giu.Button("Start").Size(-1,20).OnClick(func(){
						if password!=cPassword{
							_status = "Passwords don't match."
							_status_color = color.RGBA{0xff,0x00,0x00,255}
							return
						}
						if mode=="encrypt"{
							outputFile = outputEntry+".pcv"
						}else{
							outputFile = outputEntry
						}
						_,err := os.Stat(outputFile)
						if err==nil{
							giu.OpenPopup("Confirmation")
						}else{
							go work()
						}
					}),

					giu.Style().SetColor(giu.StyleColorText,_status_color).To(
						giu.Label(_status),
					),
				),

				// File shredder tab
				giu.TabItem("Shredder").Layout(
					giu.Custom(func(){
						if giu.IsItemActive(){
							tab = 1
						}
					}),

					giu.Label("Select a mode below and drop file(s) and folder(s) here."),
					giu.Label("Warning: Anything dropped here will be shredded immediately!"),
					//giu.Dummy(10,0),
					giu.Combo("##shredder_mode",shredModes[shredMode],shredModes,&shredMode).Size(463),
					//giu.Dummy(10,0),
					giu.ProgressBar(shredProgress).Overlay(shredOverlay).Size(-1,0),
					giu.Label(shredding).Wrapped(true),
				),

				// File checksum generator tab
				giu.TabItem("Checksum").Layout(
					giu.Custom(func(){
						if giu.IsItemActive(){
							tab = 2
						}
					}),

					giu.Label("Toggle the hashes you would like to generate and drop a file here."),

					// MD5
					giu.Row(
						giu.Checkbox("MD5:",&md5_selected),
						giu.Dummy(-45,0),
						giu.Button("Copy##md5").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_md5)
						}),
					),
					giu.Style().SetColor(giu.StyleColorBorder,md5_color).To(
						giu.InputText(&cs_md5).Size(-1).Flags(giu.InputTextFlagsReadOnly),
					),

					// SHA1
					giu.Row(
						giu.Checkbox("SHA1:",&sha1_selected),
						giu.Dummy(-45,0),
						giu.Button("Copy##sha1").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_sha1)
						}),
					),
					giu.Style().SetColor(giu.StyleColorBorder,sha1_color).To(
						giu.InputText(&cs_sha1).Size(-1).Flags(giu.InputTextFlagsReadOnly),
					),

					// SHA256
					giu.Row(
						giu.Checkbox("SHA256:",&sha256_selected),
						giu.Dummy(-45,0),
						giu.Button("Copy##sha256").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_sha256)
						}),
					),
					giu.Style().SetColor(giu.StyleColorBorder,sha256_color).To(
						giu.InputText(&cs_sha256).Size(-1).Flags(giu.InputTextFlagsReadOnly),
					),

					// SHA3-256
					giu.Row(
						giu.Checkbox("SHA3-256:",&sha3_256_selected),
						giu.Dummy(-45,0),
						giu.Button("Copy##sha3_256").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_sha3_256)
						}),
					),
					giu.Style().SetColor(giu.StyleColorBorder,sha3_256_color).To(
						giu.InputText(&cs_sha3_256).Size(-1).Flags(giu.InputTextFlagsReadOnly),
					),

					// BLAKE2b
					giu.Row(
						giu.Checkbox("BLAKE2b:",&blake2b_selected),
						giu.Dummy(-45,0),
						giu.Button("Copy##blake2b").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_blake2b)
						}),
					),
					giu.Style().SetColor(giu.StyleColorBorder,blake2b_color).To(
						giu.InputText(&cs_blake2b).Size(-1).Flags(giu.InputTextFlagsReadOnly),
					),

					// BLAKE2s
					giu.Row(
						giu.Checkbox("BLAKE2s:",&blake2s_selected),
						giu.Dummy(-45,0),
						giu.Button("Copy##blake2s").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_blake2s)
						}),
					),
					giu.Style().SetColor(giu.StyleColorBorder,blake2s_color).To(
						giu.InputText(&cs_blake2s).Size(-1).Flags(giu.InputTextFlagsReadOnly),
					),
					
					// Input entry for validating a checksum
					giu.Label("Validate a checksum:"),
					giu.InputText(&cs_validate).Size(-1).OnChange(func(){
						md5_color = color.RGBA{0x10,0x10,0x10,255}
						sha1_color = color.RGBA{0x10,0x10,0x10,255}
						sha256_color = color.RGBA{0x10,0x10,0x10,255}
						sha3_256_color = color.RGBA{0x10,0x10,0x10,255}
						blake2b_color = color.RGBA{0x10,0x10,0x10,255}
						blake2s_color = color.RGBA{0x10,0x10,0x10,255}
						if cs_validate==""{
							return
						}
						if cs_validate==cs_md5{
							md5_color = color.RGBA{0x00,0xff,0x00,255}
						}else if cs_validate==cs_sha1{
							sha1_color = color.RGBA{0x00,0xff,0x00,255}
						}else if cs_validate==cs_sha256{
							sha256_color = color.RGBA{0x00,0xff,0x00,255}
						}else if cs_validate==cs_sha3_256{
							sha3_256_color = color.RGBA{0x00,0xff,0x00,255}
						}else if cs_validate==cs_blake2b{
							blake2b_color = color.RGBA{0x00,0xff,0x00,255}
						}else if cs_validate==cs_blake2s{
							blake2s_color = color.RGBA{0x00,0xff,0x00,255}
						}
						giu.Update()
					}),

					// Progress bar
					giu.Label("Progress:"),
					giu.ProgressBar(cs_progress).Size(-1,0),
				),
				giu.TabItem("About").Layout(
					giu.Custom(func(){
						if giu.IsItemActive(){
							tab = 3
						}
					}),
					giu.Label("Picocrypt "+version+", created by Evan Su (https://evansu.cc)"),
				),
			),
		),
	)
	if working{
		giu.SingleWindow().IsOpen(&working).Layout(
			giu.Label("Tips:"),
			giu.Label("    - Choose a strong password with more than 16 characters."),
			giu.Label("    - Use a unique password that isn't used anywhere else."),
			giu.Label("    - Trust no one but yourself and never give out your key."),
			giu.Label("    - For highly sensitive files, encrypt them while offline."),
			giu.Label("    - An antivirus can be beneficial to prevent keyloggers."),
			giu.Label("    - Encrypt your root filesystem for maximal security."),
			giu.Dummy(0,-50),
			
			// Progress bar
			giu.Row(
				giu.ProgressBar(progress).Size(-54,0).Overlay(progressInfo),
				giu.Dummy(-59,0),
				giu.Button("Cancel").Size(50,0).OnClick(func(){
					working = false
				}),
			),
			giu.Label(status),
		)
	}
}

// Handle files dropped into Picocrypt by user
func onDrop(names []string){
	_status = ""
	recombine = false
	if tab==0{
		// Clear variables
		onlyFiles = nil
		onlyFolders = nil
		allFiles = nil
		files,folders := 0,0

		// Reset UI
		resetUI()

		// Hide the ".pcv" label
		orLabel = "or"
		outputWidth = 376

		// There's only one dropped item
		if len(names)==1{
			stat,_ := os.Stat(names[0])

			// Check if dropped item is a file or a folder
			if stat.IsDir(){
				folders++
				inputLabel = "1 folder selected."

				// Add the folder
				onlyFolders = append(onlyFolders,names[0])

				// Set 'outputEntry' to 'Encrypted.zip' in the same directory
				outputEntry = filepath.Join(filepath.Dir(names[0]),"Encrypted.zip")
				
				mode = "encrypt"
				// Show the ".pcv" file extension
				orLabel = ".pcv or"
				outputWidth = 341
			}else{
				files++
				name := filepath.Base(names[0])

				nums := []string{"0","1","2","3","4","5","6","7","8","9"}
				endsNum := false
				for _,i := range(nums){
					if strings.HasSuffix(names[0],i){
						endsNum = true
					}
				}
				isSplit := strings.Contains(names[0],".pcv.")&&endsNum

				// Decide if encrypting or decrypting
				if strings.HasSuffix(names[0],".pcv")||isSplit{
					mode = "decrypt"
					inputLabel = name+" (will decrypt)"
				
					if isSplit{
						ind := strings.Index(names[0],".pcv")
						names[0] = names[0][:ind]
						outputEntry = names[0]
						recombine = true
					}else{
						outputEntry = names[0][:len(names[0])-4]
					}
					

					// Open input file in read-only mode
					fin,_ := os.Open(names[0])

					// Read metadata and insert into box
					fin.Read(make([]byte,15))
					tmp := make([]byte,30)
					fin.Read(tmp)
					tmp,_ = rsDecode(rs10,tmp)
					metadataLength,_ := strconv.Atoi(string(tmp))
					//fmt.Println(metadataLength)
					tmp = make([]byte,metadataLength)
					fin.Read(tmp)
					metadata = string(tmp)

					flags := make([]byte,15)
					fin.Read(flags)
					flags,_ = rsDecode(rs5,flags)

					if flags[1]==1{
						keyfilePrompt = "Keyfile (required):"
						keyfileLabel = "Click here to select keyfile."
					}

					fin.Close()
				}else{
					mode = "encrypt"
					inputLabel = name+" (will encrypt)"
					outputEntry = names[0]

					// Show the ".pcv" file extension
					orLabel = ".pcv or"
					outputWidth = 341
				}

				// Add the file
				onlyFiles = append(onlyFiles,names[0])

				inputFile = names[0]
			}
		}else{
			mode = "encrypt"
			// Show the ".pcv" file extension
			orLabel = ".pcv or"
			outputWidth = 341

			// There are multiple dropped items, check each one
			for _,name := range names{
				stat,_ := os.Stat(name)

				// Check if item is a file or a directory
				if stat.IsDir(){
					folders++
					onlyFolders = append(onlyFolders,name)
				}else{
					files++
					onlyFiles = append(onlyFiles,name)
					allFiles = append(allFiles,name)
				}
			}

			if folders==0{
				// If folders==0, then there are multiple files
				inputLabel = fmt.Sprintf("%d files selected.",files)
			}else if files==0{
				// If files==0, then there are multiple folders
				inputLabel = fmt.Sprintf("%d folders selected.",folders)
			}else{
				// There are multiple files and folders
				if files==1&&folders>1{
					inputLabel = fmt.Sprintf("1 file and %d folders selected.",folders)
				}else if folders==1&&files>1{
					inputLabel = fmt.Sprintf("%d files and 1 folder selected.",files)
				}else if folders==1&&files==1{
					inputLabel = "1 file and 1 folder selected."
				}else{
					inputLabel = fmt.Sprintf("%d files and %d folders selected.",files,folders)
				}
			}

			// Set 'outputEntry' to 'Encrypted.zip' in the same directory
			outputEntry = filepath.Join(filepath.Dir(names[0]),"Encrypted.zip")
		}

		// If there are folders that were dropped, recusively add all files into 'allFiles'
		if folders>0{
			for _,name := range(onlyFolders){
				filepath.Walk(name,func(path string,_ os.FileInfo,_ error) error{
					stat,_ := os.Stat(path)
					if !stat.IsDir(){
						allFiles = append(allFiles,path)
					}
					return nil
				})
			}
		}
	}else if tab==1{
		go shred(names,true)
	}else if tab==2{
		go generateChecksums(names[0])
	}

	// Update the UI
	giu.Update()
}

// Start encryption/decryption
func work(){
	debug.FreeOSMemory()
	// Set some variables
	working = true
	//headerBroken := false
	//reedsoloFixed := 0
	//reedsoloErrors := 0
	var salt []byte
	var nonce []byte
	var keyHash []byte
	var _keyHash []byte
	var khash []byte
	var khash_hash []byte = make([]byte,32)
	var _khash_hash []byte
	var nonces []byte
	_fast = fast
	
	// Set the output file based on mode
	if mode=="encrypt"{
		// Compress files into a zip archive
		if len(allFiles)>1||len(onlyFolders)>0{
			rootDir := filepath.Dir(outputEntry)
			inputFile = outputEntry
			//fmt.Println(inputFile)
			file,_ := os.Create(inputFile)
			
			w := zip.NewWriter(file)
			for _,path := range(allFiles){
				if path==inputFile{
					continue
				}
				stat,_ := os.Stat(path)
				header,_ := zip.FileInfoHeader(stat)
				header.Name = strings.Replace(path,rootDir,"",1)
				header.Method = zip.Deflate
				writer,_ := w.CreateHeader(header)
				file,_ := os.Open(path)
				io.Copy(writer,file)
				file.Close()
			}
			w.Flush()
			w.Close()
			file.Close()
		}
	}

	if recombine{
		status = "Recombining file..."
		total := 0

		for{
			_,err := os.Stat(fmt.Sprintf("%s.%d",inputFile+".pcv",total))
			if err!=nil{
				break
			}
			total++
		}
		fout,_ := os.OpenFile(
			outputEntry+".pcv",
			os.O_RDWR|os.O_CREATE|os.O_TRUNC,
			0755,
		)
		for i:=0;i<total;i++{
			fin,_ := os.Open(fmt.Sprintf("%s.%d",inputFile+".pcv",i))
			for{
				data := make([]byte,1048576)
				read,err := fin.Read(data)
				if err!=nil{
					break
				}
				data = data[:read]
				fout.Write(data)
			}
			fin.Close()
			progressInfo = fmt.Sprintf("%d/%d",i,total)
			progress = float32(i)/float32(total)
			giu.Update()
		}
		fout.Close()
		outputFile = outputEntry
		inputFile = outputEntry+".pcv"
		progressInfo = ""
	}
	
	//fmt.Println(inputFile)
	stat,_ := os.Stat(inputFile)
	total := stat.Size()
	//fmt.Println(total)
	
	// Open input file in read-only mode
	fin,_ := os.Open(inputFile)
	
	var fout *os.File
	
	//fmt.Println(mode)

	// If encrypting, generate values; If decrypting, read values from file
	if mode=="encrypt"{
		status = "Generating values..."
		giu.Update()
		fout,_ = os.OpenFile(
			outputFile,
			os.O_RDWR|os.O_CREATE|os.O_TRUNC,
			0755,
		)

		// Argon2 salt and XChaCha20 nonce
		salt = make([]byte,16)
		nonce = make([]byte,24)
		
		// Write version to file
		fout.Write(rsEncode(rs5,[]byte(version)))

		// Encode the length of the metadata with Reed-Solomon
		metadataLength := []byte(fmt.Sprintf("%010d",len(metadata)))
		//fmt.Println("metadataLength:",metadataLength)
		metadataLength = rsEncode(rs10,metadataLength)
		
		// Write the length of the metadata to file
		fout.Write(metadataLength)
		
		// Write the actual metadata
		fout.Write([]byte(metadata))

		flags := make([]byte,5)
		if fast{
			flags[0] = 1
		}
		if keyfile{
			flags[1] = 1
		}
		//fmt.Println("flags:",flags)
		flags = rsEncode(rs5,flags)
		fout.Write(flags)

		// Fill salt and nonce with Go's CSPRNG
		rand.Read(salt)
		rand.Read(nonce)
		
		//fmt.Println("salt: ",salt)
		//fmt.Println("nonce: ",nonce)

		// Encode salt with Reed-Solomon and write to file
		_salt := rsEncode(rs16,salt)
		fout.Write(_salt)

		// Encode nonce with Reed-Solomon and write to file
		tmp := rsEncode(rs24,nonce)
		fout.Write(tmp)
		
		// Write placeholder for hash of key
		fout.Write(make([]byte,192))
		
		// Write placeholder for hash of hash of keyfile
		fout.Write(make([]byte,96))

		
		pairs := int(math.Ceil(float64(total)/1048576))
		
		offset := 72*pairs+48
		
		// Write placeholder for nonce/Poly1305 pairs
		fout.Write(make([]byte,offset))
	}else{
		var err error
		status = "Reading values..."
		giu.Update()
		version := make([]byte,15)
		fin.Read(version)
		version,err = rsDecode(rs5,version)
		_ = err
		//fmt.Println("version",err,string(version))

		tmp := make([]byte,30)
		fin.Read(tmp)
		tmp,err = rsDecode(rs10,tmp)
		metadataLength,_ := strconv.Atoi(string(tmp))
		//fmt.Println("metadataLength",metadataLength)
		//fmt.Println("metadataLength",err,metadataLength)

		fin.Read(make([]byte,metadataLength))

		flags := make([]byte,15)
		fin.Read(flags)
		flags,err = rsDecode(rs5,flags)
		//fmt.Println("flags",flags)
		//fmt.Println("flags",err,flags)
		fast = flags[0]==1
		keyfile = flags[1]==1

		salt = make([]byte,48)
		fin.Read(salt)
		salt,err = rsDecode(rs16,salt)
		//fmt.Println("salt",err,salt)
		
		nonce = make([]byte,72)
		fin.Read(nonce)
		nonce,err = rsDecode(rs24,nonce)
		//fmt.Println("nonce",err,nonce)
		
		//fmt.Println("salt: ",salt)
		//fmt.Println("nonce: ",nonce)
		
		_keyHash = make([]byte,192)
		fin.Read(_keyHash)
		//fmt.Println("ud",_keyHash)
		_keyHash,err = rsDecode(rs64,_keyHash)
		//fmt.Println("_keyHash",keyHash)
		//fmt.Println("_keyHash",err)
		
		_khash_hash = make([]byte,96)
		fin.Read(_khash_hash)
		_khash_hash,_ = rsDecode(rs32,_khash_hash)
		//fmt.Println("crcHash",crcHash)
		//fmt.Println("_khash_hash",err)
		
		var _tmp float64
		if fast{
			_tmp = math.Ceil(float64(total-int64(metadataLength+468))/float64(1048664))
		}else{
			_tmp = math.Ceil(float64(total-int64(metadataLength+468))/float64(1048696))
		}
		nonces = make([]byte,int(_tmp*72)+48)
		fin.Read(nonces)
		//fmt.Println("Nonces: ",nonces)
	}
	
	giu.Update()
	status = "Deriving key..."
	progress = 0
	
	// Derive encryption/decryption key
	var key []byte
	if fast{
		key = argon2.IDKey(
			[]byte(password),
			salt,
			4,
			131072,
			4,
			32,
		)[:]
	}else{
		key = argon2.IDKey(
			[]byte(password),
			salt,
			8,
			1048576,
			8,
			32,
		)[:]
	}
	
	//fmt.Println("key",key)
	if !working{
		_status = "Operation cancelled by user."
		_status_color = color.RGBA{0xff,0xff,0xff,255}
		fast = _fast
		debug.FreeOSMemory()
		return
	}

	if keyfile{
		kin,_ := os.Open(keyfilePath)
		kstat,_ := os.Stat(keyfilePath)
		//fmt.Println(kstat.Size())
		kbytes := make([]byte,kstat.Size())
		kin.Read(kbytes)
		kin.Close()
		ksha3 := sha3.New256()
		ksha3.Write(kbytes)
		khash = ksha3.Sum(nil)

		khash_sha3 := sha3.New256()
		khash_sha3.Write(khash)
		khash_hash = khash_sha3.Sum(nil)
		//fmt.Println("khash",khash)
		//fmt.Println("khash_hash",khash_hash)
	}
	
	//key = make([]byte,32)
	//fmt.Println("output",outputFile)
	sha3_512 := sha3.New512()
	sha3_512.Write(key)
	keyHash = sha3_512.Sum(nil)
	//fmt.Println("keyHash: ",keyHash)
	
	// Check is password is correct
	if mode=="decrypt"{
		keyCorrect := true
		keyfileCorrect := true
		var tmp bool
		for i,j := range(_keyHash){
			if keyHash[i]!=j{
				keyCorrect = false
				break
			}
		}
		if keyfile{
			for i,j := range(_khash_hash){
				if khash_hash[i]!=j{
					keyfileCorrect = false
					break
				}
			}
			tmp = !keyCorrect||!keyfileCorrect
		}else{
			tmp = !keyCorrect
		}
		if tmp{
			if keep{
				kept = true
			}else{
				fin.Close()
				working = false
				if !keyCorrect{
					_status = "The provided password is incorrect."
				}else{
					_status = "The provided keyfile is incorrect."
				}
				_status_color = color.RGBA{0xff,0x00,0x00,255}
				key = nil
				fast = _fast
				if recombine{
					os.Remove(inputFile)
				}
				debug.FreeOSMemory()
				return
			}
		}

		fout,_ = os.Create(outputFile)
	}

	if keyfile{
		// XOR key and keyfile
		tmp := key
		key = make([]byte,32)
		for i,_ := range(key){
			key[i] = tmp[i]^khash[i]
		}
		//fmt.Println("key",key)
	}

	
	done := 0
	counter := 0
	startTime := time.Now()

	cipher,_ := chacha20poly1305.NewX(key)
	
	if mode=="decrypt"{
		_mac := nonces[len(nonces)-48:]
		_mac,_ = rsDecode(rs16,_mac)
		//fmt.Println("_mac ",_mac)
		nonces = nonces[:len(nonces)-48]
		var tmp []byte
		var chunk []byte
		for i,j := range(nonces){
			chunk = append(chunk,j)
			if (i+1)%72==0{
				chunk,_ = rsDecode(rs24,chunk)
				for _,k := range(chunk){
					tmp = append(tmp,k)
				}
				chunk = nil
			}
		}
		/*for _,j := range(_mac){
			tmp = append(tmp,j)
		}*/

		var authentic bool
		nonces,authentic = monocypher.Unlock(tmp,nonce,key,_mac)
		if !authentic{
			if keep{
				kept = true
			}else{
				fin.Close()
				fout.Close()
				working = false
				_status = "The file is either corrupted or intentionally modified."
				_status_color = color.RGBA{0xff,0x00,0x00,255}
				fast = _fast
				debug.FreeOSMemory()
				return
			}
		}
		//fmt.Println("UNENCRYPTED NONCES: ",nonces)
	}
	for{
		if !working{
			fin.Close()
			fout.Close()
			os.Remove(outputFile)
			_status = "Operation cancelled by user."
			_status_color = color.RGBA{0xff,0xff,0xff,255}
			fast = _fast
			debug.FreeOSMemory()
			return
		}
		//fmt.Println("Encrypt/decrypt loop")
		var _data []byte
		var data []byte
		var _nonce []byte
		if mode=="encrypt"{
			_data = make([]byte,1048576)
		}else{
			if fast{
				_data = make([]byte,1048592)
			}else{
				_data = make([]byte,1048624)
			}
		}

		size,err := fin.Read(_data)
		if err!=nil{
			break
		}
		data = _data[:size]
		

		if mode=="encrypt"{
			_nonce = make([]byte,24)
			rand.Read(_nonce)
			for _,i := range(_nonce){
				nonces = append(nonces,i)
			}
		}else{
			_nonce = nonces[counter*24:counter*24+24]
		}
		
		//fmt.Println("Data nonce: ",_nonce)
		//fmt.Println("Data: ",data)
		if mode=="encrypt"{
			if fast{
				data = cipher.Seal(nil,_nonce,data,nil)
				fout.Write(data)
				//crc.Write(data)
			}else{
				mac,data := monocypher.Lock(data,_nonce,key)
				fout.Write(data)
				fout.Write(rsEncode(rs16,mac))
				//crc.Write(data)
				//crc.Write(mac)
			}

			//fout.Write(data)
		}else{
			//fmt.Println("DECODE LOOP")
			//crc.Write(data)
			if fast{
				data,err = cipher.Open(nil,_nonce,data,nil)
				if err!=nil{
					if keep{
						kept = true
						mac := data[len(data)-16:]
						data = data[:len(data)-16]
						data,_ = monocypher.Unlock(data,_nonce,key,mac)
					}else{
						fin.Close()
						fout.Close()
						broken()
						fast = _fast
						debug.FreeOSMemory()
						return
					}
				}
			}else{
				//crc.Write(data)
				mac,_ := rsDecode(rs16,data[len(data)-48:])
				data = data[:len(data)-48]
				var authentic bool
				data,authentic = monocypher.Unlock(data,_nonce,key,mac)
				if !authentic{
					if keep{
						kept = true
					}else{
						fin.Close()
						fout.Close()
						broken()
						fast = _fast
						debug.FreeOSMemory()
						return
					}
				}
			}
			fout.Write(data)
			//fmt.Println(authentic)
			//fmt.Println("DECRYPTED DATA: ",data)
		}
		
		if mode=="encrypt"{
			done += 1048576
		}else{
			if fast{
				done += 1048592
			}else{
				done += 1048624
			}
		}
		counter++

		progress = float32(done)/float32(total)
		
		elapsed:= float64(int64(time.Now().Sub(startTime)))/float64(1000000000)
		
		speed := (float64(done)/elapsed)/1000000
		eta := math.Abs(float64(total-int64(done))/(speed*1000000))
		
		progressInfo = fmt.Sprintf("%.2f%%",progress*100)
		
		status = fmt.Sprintf("Working at %.2f MB/s (ETA: %.1fs)",speed,eta)
		
		giu.Update()
	}

	if mode=="encrypt"{
		//fmt.Println("'nonces' before RS: ",nonces)
		fout.Seek(int64(180+len(metadata)),0)
		fout.Write(rsEncode(rs64,keyHash))
		fout.Write(rsEncode(rs32,khash_hash))

		_mac,tmp := monocypher.Lock(nonces,nonce,key) 
		//fmt.Println(_mac)
		var chunk []byte

		for i,j := range(tmp){
			chunk = append(chunk,j)
			if (i+1)%24==0{
				fout.Write(rsEncode(rs24,chunk))
				chunk = nil
			}
		}
		fout.Write(rsEncode(rs16,_mac))

	}
	
	fin.Close()
	fout.Close()

	if split{
		status = "Splitting file..."
		stat,_ := os.Stat(outputFile)
		size := stat.Size()
		finished := 0
		var splitted []string
		//fmt.Println(size)
		chunkSize,_ := strconv.Atoi(splitSize)
		//fmt.Println(splitSelected)
		if splitSelected==0{
			chunkSize *= 1024
		}else if splitSelected==1{
			chunkSize *= 1048576
		}else{
			chunkSize *= 1073741824
		}
		chunks := int(math.Ceil(float64(size)/float64(chunkSize)))
		fin,_ := os.Open(outputFile)
		for i:=0;i<chunks;i++{
			//fmt.Println(fmt.Sprintf("%s.%d",outputFile,i))
			fout,_ := os.Create(fmt.Sprintf("%s.%d",outputFile,i))
			done := 0
			for{
				data := make([]byte,1048576)
				read,err := fin.Read(data)
				if err!=nil{
					break
				}
				if !working{
					fin.Close()
					fout.Close()
					_status = "Operation cancelled by user."
					_status_color = color.RGBA{0xff,0xff,0xff,255}
					for _,j := range(splitted){
						os.Remove(j)
					}
					os.Remove(fmt.Sprintf("%s.%d",outputFile,i))
					//fmt.Println("outputFile",outputFile)
					os.Remove(outputFile)
					fast = _fast
					debug.FreeOSMemory()
					return
				}
				data = data[:read]
				fout.Write(data)
				done += read
				if done>=chunkSize{
					break
				}
			}		
			fout.Close()
			finished++
			splitted = append(splitted,fmt.Sprintf("%s.%d",outputFile,i))
			progress = float32(finished)/float32(chunks)
			progressInfo = fmt.Sprintf("%d/%d",finished,chunks)
			giu.Update()
		}
		fin.Close()
		if erase{
			status = "Shredding the unsplitted file..."
			progressInfo = ""
			shred([]string{outputFile}[:],false)
		}else{
			os.Remove(outputFile)
		}
	}

	if recombine{
		os.Remove(inputFile)
	}

	// Delete the temporary zip file
	if len(allFiles)>1{
		if erase{
			progressInfo = ""
			status = "Shredding temporary files..."
			shred([]string{outputEntry}[:],false)
		}else{
			os.Remove(outputEntry)
		}
	}
	//fmt.Println("==============================")
	if erase{
		progressInfo = ""
		shred(onlyFiles,false)
		shred(onlyFolders,false)
	}


	


	resetUI()
	if kept{
		_status = "The input is corrupted and/or modified. Please be careful."
		_status_color = color.RGBA{0xff,0xff,0x00,255}
	}else{
		_status = "Completed."
		_status_color = color.RGBA{0x00,0xff,0x00,255}
	}
	working = false
	kept = false
	key = nil
	status = "Ready."
	debug.FreeOSMemory()
	//fmt.Println("Exit goroutine")
}

func shred(names []string,separate bool){
	shredTotal = 0
	shredDone = 0
	if separate{
		shredOverlay = "Shredding..."
	}
	for _,name := range(names){
		filepath.Walk(name,func(path string,_ os.FileInfo,err error) error{
			if err!=nil{
				return nil
			}
			stat,_ := os.Stat(path)
			if !stat.IsDir(){
				shredTotal++
			}
			return nil
		})
	}
	for _,name := range(names){
		shredding = name
		if runtime.GOOS=="linux"||runtime.GOOS=="darwin"{
			stat,_ := os.Stat(name)
			if stat.IsDir(){
				var coming []string
				filepath.Walk(name,func(path string,_ os.FileInfo,err error) error{
					if err!=nil{
						return nil
					}
					fmt.Println(path)
					stat,_ := os.Stat(path)
					if !stat.IsDir(){
						if len(coming)==128{
							var wg sync.WaitGroup
							for i,j := range(coming){
								wg.Add(1)
								go func(wg *sync.WaitGroup,id int,j string){
									defer wg.Done()
									cmd := exec.Command("")
									if runtime.GOOS=="linux"{
										cmd = exec.Command("shred","-ufvz","-n","3",j)
									}else{
										cmd = exec.Command("rm","-rfP",j)
									}
									output,err := cmd.Output()
									fmt.Println(err)
									fmt.Println(output)
									shredding = j
									shredDone++
									shredUpdate(separate)
									giu.Update()
								}(&wg,i,j)
								
							}
							wg.Wait()
							coming = nil
						}else{
							coming = append(coming,path)
						}
					}
					return nil
				})
				fmt.Println(coming)
				for _,i := range(coming){
					go func(){
						cmd := exec.Command("")
						if runtime.GOOS=="linux"{
							cmd = exec.Command("shred","-ufvz","-n","3",i)
						}else{
							cmd = exec.Command("rm","-rfP",i)
						}
						output,err := cmd.Output()
						fmt.Println(err)
						fmt.Println(output)
						shredding = i
						shredDone++
						shredUpdate(separate)
						giu.Update()
					}()
				}
				os.RemoveAll(name)
			}else{
				cmd := exec.Command("")
				if runtime.GOOS=="linux"{
					cmd = exec.Command("shred","-ufvz","-n","3",name)
				}else{
					cmd = exec.Command("rm","-rfP",name)
				}
				cmd.Run()
				shredding = name+"/*"
				shredDone++
				shredUpdate(separate)
			}
		}else if runtime.GOOS=="windows"{
			stat,_ := os.Stat(name)
			if stat.IsDir(){
				filepath.Walk(name,func(path string,_ os.FileInfo,err error) error{
					if err!=nil{
						return nil
					}
					fmt.Println(path)
					stat,_ := os.Stat(path)
					if stat.IsDir(){
						t := 0
						files,_ := ioutil.ReadDir(path)
						for _,f := range(files){
							if !f.IsDir(){
								t++
							}
						}
						shredDone += float32(t)
						shredUpdate(separate)
						tmp := filepath.Join(rootDir,"sdelete64.exe")
						cmd := exec.Command(tmp,"*","-p","4")
						cmd.Dir = path
						cmd.Run()
						shredding = strings.ReplaceAll(path,"\\","/")+"/*"
					}
					return nil
				})
				os.RemoveAll(name)
			}else{
				o,e := exec.Command(filepath.Join(rootDir,"sdelete64.exe"),name,"-p","4").Output()
				fmt.Println(string(o),e)
				shredDone++
				shredUpdate(separate)
			}
		}
		fmt.Println(name)
		giu.Update()
	}
	/*if shredMode==1&&runtime.GOOS=="windows"{
		cmd := exec.Command("cipher","/w:\""+name+"\"")
		stdout,err := cmd.Output()
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(string(stdout))
	}*/
	shredding = "Ready."
	shredProgress = 0
	shredOverlay = ""
}

func shredUpdate(separate bool){
	if separate{
		shredOverlay = fmt.Sprintf("%d/%d",int(shredDone),int(shredTotal))
		shredProgress = shredDone/shredTotal
	}else{
		status = fmt.Sprintf("%d/%d",int(shredDone),int(shredTotal))
		progress = shredDone/shredTotal
	}
	giu.Update()
}

// Generate file checksums
func generateChecksums(file string){
	fin,_ := os.Open(file)

	cs_md5 = ""
	cs_sha1 = ""
	cs_sha256 = ""
	cs_sha3_256 = ""
	cs_blake2b = ""
	cs_blake2s = ""

	if md5_selected{
		cs_md5 = "Calculating..."
	}
	if sha1_selected{
		cs_sha1 = "Calculating..."
	}
	if sha256_selected{
		cs_sha256 = "Calculating..."
	}
	if sha3_256_selected{
		cs_sha3_256 = "Calculating..."
	}
	if blake2b_selected{
		cs_blake2b = "Calculating..."
	}
	if blake2s_selected{
		cs_blake2s = "Calculating..."
	}
	crc_md5 := md5.New()
	crc_sha1 := sha1.New()
	crc_sha256 := sha256.New()
	crc_sha3_256 := sha3.New256()
	crc_blake2b,_ := blake2b.New256(nil)
	crc_blake2s,_ := blake2s.New256(nil)
	stat,_ := os.Stat(file)
	total := stat.Size()
	var done int64 = 0
	for{
		var data []byte
		_data := make([]byte,1048576)
		size,err := fin.Read(_data)
		if err!=nil{
			break
		}
		data = _data[:size]

		if md5_selected{
			crc_md5.Write(data)
		}
		if sha1_selected{
			crc_sha1.Write(data)
		}
		if sha256_selected{
			crc_sha256.Write(data)
		}
		if sha3_256_selected{
			crc_sha3_256.Write(data)
		}
		if blake2b_selected{
			crc_blake2b.Write(data)
		}
		if blake2s_selected{
			crc_blake2s.Write(data)
		}

		done += int64(size)
		cs_progress = float32(done)/float32(total)
		giu.Update()
	}
	cs_progress = 0
	if md5_selected{
		cs_md5 = hex.EncodeToString(crc_md5.Sum(nil))
	}
	if sha1_selected{
		cs_sha1 = hex.EncodeToString(crc_sha1.Sum(nil))
	}
	if sha256_selected{
		cs_sha256 = hex.EncodeToString(crc_sha256.Sum(nil))
	}
	if sha3_256_selected{
		cs_sha3_256 = hex.EncodeToString(crc_sha3_256.Sum(nil))
	}
	if blake2b_selected{
		cs_blake2b = hex.EncodeToString(crc_blake2b.Sum(nil))
	}
	if blake2s_selected{
		cs_blake2s = hex.EncodeToString(crc_blake2s.Sum(nil))
	}
	giu.Update()
}


// Reset the UI to a clean state with no nothing selected
func resetUI(){
	inputLabel = "Drag and drop file(s) and folder(s) into this window."
	outputEntry = ""
	orLabel = "or"
	outputWidth = 376
	password = ""
	cPassword = ""
	keyfilePrompt = "Keyfile (optional):"
	keyfileLabel = "Use a keyfile"
	keyfile = false
	metadata = ""
	keep = false
	erase = false
	reedsolo = false
	split = false
	splitSize = ""
	fast = false
	progress = 0
	progressInfo = ""
	_status = ""
	_status_color = color.RGBA{0xff,0xff,0xff,255}
	giu.Update()
}

func broken(){
	working = false
	_status = "The file is either corrupted or intentionally modified."
	_status_color = color.RGBA{0xff,0x00,0x00,255}
	os.Remove(outputFile)
	debug.FreeOSMemory()
}

/*func rsEncode(data []byte,encoder reedsolomon.Encoder,size int) []byte{
	shards,_ := encoder.Split(data)
	encoder.Encode(shards)
	tmp := make([]byte,size)
	for i,shard := range(shards){
		tmp[i] = shard[0]
	}
	return tmp
}

func rsDecode(data []byte,encoder reedsolomon.Encoder,size int) []byte{
	res := make([][]byte,len(data))
	for i,_ := range(data){
		tmp := make([]byte,1)
		tmp[0] = data[i]
		res[i] = tmp
	}
	_ = encoder.Reconstruct(res)
	res = res[:size]
	tmp := make([]byte,size)
	for i,shard := range(res){
		tmp[i] = shard[0]
	}
	return tmp
}*/
func rsEncode(rs *infectious.FEC,data []byte) []byte{
	var res []byte
	rs.Encode(data,func(s infectious.Share){
		res = append(res,s.DeepCopy().Data[0])
	})
	return res
}

func rsDecode(rs *infectious.FEC,data []byte) ([]byte,error){
	tmp := make([]infectious.Share,rs.Total())
	for i:=0;i<rs.Total();i++{
		tmp[i] = infectious.Share{
			Number:i,
			Data:[]byte{data[i]},
		}
	}
	res,err := rs.Decode(nil,tmp)
	return res,err
}


// Create the master window, set callbacks, and start the UI
func main(){
	go func(){
		v,err := http.Get("https://raw.githubusercontent.com/HACKERALERT/Picocrypt/main/internals/version.txt")
		if err==nil{
			fmt.Println(v)
			body,err := io.ReadAll(v.Body)
			v.Body.Close()
			if err==nil{
				if string(body[:5])!=version{
					_status = "A newer version is available."
					_status_color = color.RGBA{0,0xff,0,255}
				}
			}
		}
	}()
	dialog.Init()
	if runtime.GOOS=="windows"{
		exec.Command(filepath.Join(rootDir,"sdelete64.exe"),"/accepteula")
	}

	giu.SetDefaultFontFromBytes(font,18)
	window := giu.NewMasterWindow("Picocrypt",480,500,giu.MasterWindowFlagsNotResizable)
	window.SetDropCallback(onDrop)
	dpi = giu.Context.GetPlatform().GetContentScale()
	window.Run(startUI)
}
