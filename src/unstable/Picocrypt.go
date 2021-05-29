package main

/*

Picocrypt v1.13
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

*/

import (
	"io"
	"os"
	"fmt"
	"math"
	"time"
	"sync"
	"os/exec"
	"strings"
	"strconv"
	"runtime"
	"net/http"
	"io/ioutil"
	"image/color"
	"crypto/md5"
	"archive/zip"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"path/filepath"
	"crypto/sha256"
	"runtime/debug"
	"github.com/pkg/browser"
	"github.com/zeebo/blake3"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/argon2"
	g "github.com/AllenDang/giu"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"github.com/atotto/clipboard"
	di "github.com/OpenDiablo2/dialog"
	"github.com/klauspost/reedsolomon"
	ig "github.com/AllenDang/imgui-go"
	"golang.org/x/crypto/chacha20poly1305"
	"github.com/HACKERALERT/Picocypher/monocypher"
)

var version = "v1.13"
var exe,_ = os.Executable()
var rootDir = filepath.Dir(exe)

// Global variables
var dpi float32
var mode string
var working = false
var onlyFiles []string
var onlyFolders []string
var allFiles []string
var inputFile string
var outputFile string

// UI-related global variables
var tab = 0
var inputLabel = "Drag and drop file(s) and folder(s) into this window."
var outputEntry string
var outputWidth float32 = 376
var orLabel = "or"
var passwordState = g.InputTextFlags_Password
var showPassword = false
var keyfile = false
var keyfilePrompt = "Keyfile (optional):"
var progress float32 = 0
var progressInfo = ""
var status = "Ready."
var _status = "Ready."
var _status_color = color.RGBA{0xff,0xff,0xff,255}
var splitUnits = []string{
	"KB",
	"MB",
	"GB",
}
var splitSelected int32
var items = []string{
	"Normal (4 passes)",
	"Paranoid (Still in development, don't use)",
}
var itemSelected int32
var shredProgress float32
var shredDone float32
var shredTotal float32
var shredOverlay string
var shredding = "Ready."

var recombine bool


// User input variables
var password string
var cPassword string
var keyfilePath string
var keyfileLabel = "Use a keyfile"
var metadata string
var keep bool
var erase bool
var reedsolo bool
var split bool
var splitSize string
var fast bool
var _fast bool

var kept = false

// Reed-Solomon encoders
var rs5_128,_ = reedsolomon.New(5,128)
var rs10_128,_ = reedsolomon.New(10,128)
var rs16_128,_ = reedsolomon.New(16,128)
var rs24_128,_ = reedsolomon.New(24,128)
var rs32_128,_ = reedsolomon.New(32,128)
var rs64_128,_ = reedsolomon.New(64,128)

// File checksum generator variables
var cs_md5 string
var cs_sha1 string
var cs_sha256 string
var cs_sha3_256 string
var cs_blake2b string
var cs_blake2s string
var cs_blake3 string
var cs_validate string
var md5_color = color.RGBA{0x10,0x10,0x10,255}
var sha1_color = color.RGBA{0x10,0x10,0x10,255}
var sha256_color = color.RGBA{0x10,0x10,0x10,255}
var sha3_256_color = color.RGBA{0x10,0x10,0x10,255}
var blake2b_color = color.RGBA{0x10,0x10,0x10,255}
var blake2s_color = color.RGBA{0x10,0x10,0x10,255}
var blake3_color = color.RGBA{0x10,0x10,0x10,255}
var cs_progress float32 = 0
var md5_selected = false
var sha1_selected = false
var sha256_selected = false
var sha3_256_selected = false
var blake2b_selected = false
var blake2s_selected = false
var blake3_selected = false

// Create the user interface
func startUI(){
	g.SingleWindow("Picocrypt").Layout(
		g.Style().SetColor(ig.StyleColorBorder,color.RGBA{0x10,0x10,0x10,255}).To(
			// The tab bar, which contains different tabs for different features
			g.TabBar("TabBar").Layout(
				// File encryption/decryption tab
				g.TabItem("Encryption/decryption").Layout(
					// Update 'tab' to indicate active tab
					g.Custom(func(){
						if g.IsItemActive(){
							tab = 0
						}
					}),
					
					// Confirm overwrite with a modal
					g.PopupModal("Confirmation").Layout(
						g.Label("Output already exists. Overwrite?"),
						g.Row(
							g.Button("No").Size(110,0).OnClick(func(){
								g.CloseCurrentPopup()
							}),
							g.Dummy(-118,0),
							g.Button("Yes").Size(110,0).OnClick(func(){
								g.CloseCurrentPopup()
								go work()
							}),
							
						),
					),

					// Label listing the input files and a button to clear input files
					//g.Dummy(0,10),
					g.Row(
						g.Label(inputLabel),
						g.Dummy(-55,0),
						g.Button("Clear").Size(46,0).OnClick(resetUI),
					),

					// Allow user to choose a custom output path and name
					//g.Dummy(10,0),
					g.Label("Save output as:"),
					g.Row(
						g.InputText("##output",&outputEntry).Size(outputWidth/dpi),
						g.Label(orLabel),
						g.Button("Save as").OnClick(func(){
							file,_ := di.File().Title("Save as").Save()

							// Return if user canceled the file dialog
							if file==""{
								return
							}

							// Remove the extra ".pcv" extension if needed
							if strings.HasSuffix(file,".pcv"){
								file = file[:len(file)-4]
							}
							outputEntry = file
						}),
					),

					// Prompt for password
					//g.Dummy(10,0),
					g.Row(
						g.Label("Password:"),
						g.Dummy(-220,0),
						g.Label(keyfilePrompt),
					),
					g.Row(
						g.InputText("##password",&password).Size(200/dpi).Flags(passwordState),
						g.Checkbox("##showPassword",&showPassword).OnChange(func(){
							if passwordState==g.InputTextFlags_Password{
								passwordState = g.InputTextFlags_None
							}else{
								passwordState = g.InputTextFlags_Password
							}
							g.Update()
						}),
						g.Dummy(-220,0),
						g.Checkbox(keyfileLabel,&keyfile).OnChange(func(){
							if !keyfile{
								keyfileLabel = "Use a keyfile"
								return
							}
							filename,err := di.File().Load()
							if err!=nil{
								keyfile = false
								return
							}
							keyfileLabel = filename
							keyfilePath = filename
						}),
					),

					// Prompt to confirm password
					//g.Dummy(10,0),
					g.Label("Confirm password:"),
					g.InputText("##cPassword",&cPassword).Size(200/dpi).Flags(passwordState),

					// Optional metadata
					//g.Dummy(10,0),
					g.Label("Metadata (optional):"),
					g.InputTextMultiline("##metadata",&metadata).Size(230,126),

					// Advanced options can be enabled with checkboxes
					//g.Dummy(10,0),
					g.Checkbox("Keep decrypted output even if it's corrupted or modified",&keep),
					g.Checkbox("Securely shred the original file(s) and folder(s)",&erase),
					g.Row(
						g.Checkbox("(Not ready) Encode with Reed-Solomon to prevent corruption",&reedsolo),
						g.Button("?").OnClick(func(){
							browser.OpenURL("https://en.wikipedia.org/wiki/Reed%E2%80%93Solomon_error_correction")
						}),
					),
					g.Row(
						g.Checkbox("Split output into chunks of",&split),
						g.InputText("##splitSize",&splitSize).Size(30).Flags(g.InputTextFlags_CharsDecimal),
						g.Combo("##splitter",splitUnits[splitSelected],splitUnits,&splitSelected).Size(40),
						//g.Label("MB"),
					),
					g.Checkbox("Fast mode (less secure, not as durable)",&fast),

					// Start and cancel buttons
					//g.Dummy(10,0),
					g.Button("Start").Size(-1,20).OnClick(func(){
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
							g.OpenPopup("Confirmation")
						}else{
							go work()
						}
					}),

					//g.Dummy(10,0),
					g.Style().SetColor(ig.StyleColorText,_status_color).To(
						g.Label(_status),
					),
				),

				// File shredder tab
				g.TabItem("Shredder").Layout(
					// Update 'tab' to indicate active tab
					g.Custom(func(){
						if g.IsItemActive(){
							tab = 1
						}
					}),

					g.Label("Select a mode below and drop file(s) and folder(s) here."),
					g.Label("Warning: Anything dropped here will be shredded immediately!"),
					//g.Dummy(10,0),
					g.Combo("##shredder_mode",items[itemSelected],items,&itemSelected).Size(463),
					//g.Dummy(10,0),
					g.ProgressBar(shredProgress).Overlay(shredOverlay).Size(-1,0),
					g.Label(shredding).Wrapped(true),
				),

				// File checksum generator tab
				g.TabItem("Checksum").Layout(
					// Update 'tab' to indicate active tab
					g.Custom(func(){
						if g.IsItemActive(){
							tab = 2
						}
					}),

					g.Label("Toggle the hashes you would like to generate and drop a file here."),
					
					// MD5
					//g.Dummy(10,0),
					g.Row(
						g.Checkbox("MD5:",&md5_selected),
						g.Dummy(-45,0),
						g.Button("Copy##md5").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_md5)
						}),
					),
					g.Style().SetColor(ig.StyleColorBorder,md5_color).To(
						g.InputText("##cs_md5",&cs_md5).Size(-1).Flags(g.InputTextFlags_ReadOnly),
					),

					// SHA1
					//g.Dummy(10,0),
					g.Row(
						g.Checkbox("SHA1:",&sha1_selected),
						g.Dummy(-45,0),
						g.Button("Copy##sha1").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_sha1)
						}),
					),
					g.Style().SetColor(ig.StyleColorBorder,sha1_color).To(
						g.InputText("##cs_sha1",&cs_sha1).Size(-1).Flags(g.InputTextFlags_ReadOnly),
					),

					// SHA256
					//g.Dummy(10,0),
					g.Row(
						g.Checkbox("SHA256:",&sha256_selected),
						g.Dummy(-45,0),
						g.Button("Copy##sha256").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_sha256)
						}),
					),
					g.Style().SetColor(ig.StyleColorBorder,sha256_color).To(
						g.InputText("##cs_sha256",&cs_sha256).Size(-1).Flags(g.InputTextFlags_ReadOnly),
					),

					// SHA3-256
					//g.Dummy(10,0),
					g.Row(
						g.Checkbox("SHA3-256:",&sha3_256_selected),
						g.Dummy(-45,0),
						g.Button("Copy##sha3_256").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_sha3_256)
						}),
					),
					g.Style().SetColor(ig.StyleColorBorder,sha3_256_color).To(
						g.InputText("##cs_sha3_256",&cs_sha3_256).Size(-1).Flags(g.InputTextFlags_ReadOnly),
					),

					// BLAKE2b
					//g.Dummy(10,0),
					g.Row(
						g.Checkbox("BLAKE2b:",&blake2b_selected),
						g.Dummy(-45,0),
						g.Button("Copy##blake2b").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_blake2b)
						}),
					),
					g.Style().SetColor(ig.StyleColorBorder,blake2b_color).To(
						g.InputText("##cs_blake2b",&cs_blake2b).Size(-1).Flags(g.InputTextFlags_ReadOnly),
					),

					// BLAKE2s
					//g.Dummy(10,0),
					g.Row(
						g.Checkbox("BLAKE2s:",&blake2s_selected),
						g.Dummy(-45,0),
						g.Button("Copy##blake2s").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_blake2s)
						}),
					),
					g.Style().SetColor(ig.StyleColorBorder,blake2s_color).To(
						g.InputText("##cs_blake2s",&cs_blake2s).Size(-1).Flags(g.InputTextFlags_ReadOnly),
					),

					// BLAKE3
					//g.Dummy(10,0),
					g.Row(
						g.Checkbox("BLAKE3:",&blake3_selected),
						g.Dummy(-45,0),
						g.Button("Copy##blake3").Size(36,0).OnClick(func(){
							clipboard.WriteAll(cs_blake3)
						}),
					),
					g.Style().SetColor(ig.StyleColorBorder,blake3_color).To(
						g.InputText("##cs_blake3",&cs_blake3).Size(-1).Flags(g.InputTextFlags_ReadOnly),
					),
					
					// Input box for validating checksum
					//g.Dummy(10,0),
					g.Label("Validate a checksum:"),
					g.InputText("##cs_validate",&cs_validate).Size(-1).OnChange(func(){
						md5_color = color.RGBA{0x10,0x10,0x10,255}
						sha1_color = color.RGBA{0x10,0x10,0x10,255}
						sha256_color = color.RGBA{0x10,0x10,0x10,255}
						sha3_256_color = color.RGBA{0x10,0x10,0x10,255}
						blake2b_color = color.RGBA{0x10,0x10,0x10,255}
						blake2s_color = color.RGBA{0x10,0x10,0x10,255}
						blake3_color = color.RGBA{0x10,0x10,0x10,255}
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
						}else if cs_validate==cs_blake3{
							blake3_color = color.RGBA{0x00,0xff,0x00,255}
						}
						g.Update()
					}),

					// Progress bar
					//g.Dummy(10,0),
					g.Label("Progress:"),
					g.ProgressBar(cs_progress).Size(-1,0),
				),
				g.TabItem("About").Layout(
					// Update 'tab' to indicate active tab
					g.Custom(func(){
						if g.IsItemActive(){
							tab = 3
						}
					}),
					//g.Dummy(30,0),
					g.Label("Picocrypt "+version+", created by Evan Su (https://evansu.cc)"),
				),
			),
		),
	)
	if working{
		g.SingleWindow("Working..").IsOpen(&working).Layout(
			//g.Dummy(30,0),
			g.Label("Tips:"),
			g.Label("    - Choose a strong password with more than 16 characters."),
			g.Label("    - Use a unique password that isn't used anywhere else."),
			g.Label("    - Trust no one but yourself and never give out your key."),
			g.Label("    - For highly sensitive files, encrypt them while offline."),
			g.Label("    - An antivirus can be beneficial to prevent keyloggers."),
			g.Label("    - Encrypt your root filesystem for maximal security."),
			g.Dummy(0,-50),
			

			// Progress bar
			g.Row(
				g.ProgressBar(progress).Size(-54,0).Overlay(progressInfo),
				g.Dummy(-59,0),
				g.Button("Cancel").Size(50,0).OnClick(func(){
					working = false
				}),
			),
			//g.Dummy(10,0),
			g.Label(status),
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
					fin.Read(make([]byte,133))
					tmp := make([]byte,138)
					fin.Read(tmp)
					tmp = rsDecode(tmp,rs10_128,10)
					metadataLength,_ := strconv.Atoi(string(tmp))
					//fmt.Println(metadataLength)
					tmp = make([]byte,metadataLength)
					fin.Read(tmp)
					metadata = string(tmp)

					flags := make([]byte,133)
					fin.Read(flags)
					flags = rsDecode(flags,rs5_128,5)

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
	g.Update()
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
	var khash_hash []byte
	var _khash_hash []byte
	var nonces []byte
	_fast = fast
	
	// Set the output file based on mode
	if mode=="encrypt"{
		// Compress files into a zip archive
		if len(allFiles)>1||len(onlyFolders)>0{
			rootDir := filepath.Dir(outputEntry)
			inputFile = outputEntry
			fmt.Println(inputFile)
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
			g.Update()
		}
		fout.Close()
		outputFile = outputEntry
		inputFile = outputEntry+".pcv"
		progressInfo = ""
	}
	
	fmt.Println(inputFile)
	stat,_ := os.Stat(inputFile)
	total := stat.Size()
	fmt.Println(total)
	
	// Open input file in read-only mode
	fin,_ := os.Open(inputFile)
	
	var fout *os.File
	
	fmt.Println(mode)

	// If encrypting, generate values; If decrypting, read values from file
	if mode=="encrypt"{
		status = "Generating values..."
		g.Update()
		fout,_ = os.OpenFile(
			outputFile,
			os.O_RDWR|os.O_CREATE|os.O_TRUNC,
			0755,
		)

		// Argon2 salt and XChaCha20 nonce
		salt = make([]byte,16)
		nonce = make([]byte,24)
		
		// Write version to file
		fout.Write(rsEncode([]byte(version),rs5_128,133))

		// Encode the length of the metadata with Reed-Solomon
		metadataLength := []byte(fmt.Sprintf("%010d",len(metadata)))
		//fmt.Println("metadataLength:",metadataLength)
		metadataLength = rsEncode(metadataLength,rs10_128,138)
		
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
		flags = rsEncode(flags,rs5_128,133)
		fout.Write(flags)

		// Fill salt and nonce with Go's CSPRNG
		rand.Read(salt)
		rand.Read(nonce)
		
		fmt.Println("salt: ",salt)
		fmt.Println("nonce: ",nonce)

		// Encode salt with Reed-Solomon and write to file
		_salt := rsEncode(salt,rs16_128,144)
		fout.Write(_salt)

		// Encode nonce with Reed-Solomon and write to file
		tmp := rsEncode(nonce,rs24_128,152)
		fout.Write(tmp)
		
		// Write placeholder for hash of key
		fout.Write(make([]byte,192))
		
		// Write placeholder for hash of hash of keyfile
		fout.Write(make([]byte,160))

		
		pairs := int(math.Ceil(float64(total)/1048576))
		
		offset := 152*pairs+144
		
		// Write placeholder for nonce/Poly1305 pairs
		fout.Write(make([]byte,offset))
	}else{
		status = "Reading values..."
		g.Update()
		version := make([]byte,133)
		fin.Read(version)
		version = rsDecode(version,rs5_128,5)

		tmp := make([]byte,138)
		fin.Read(tmp)
		tmp = rsDecode(tmp,rs10_128,10)
		metadataLength,_ := strconv.Atoi(string(tmp))
		//fmt.Println("metadataLength",metadataLength)

		fin.Read(make([]byte,metadataLength))

		flags := make([]byte,133)
		fin.Read(flags)
		flags = rsDecode(flags,rs5_128,5)
		//fmt.Println("flags",flags)
		fast = flags[0]==1
		keyfile = flags[1]==1

		salt = make([]byte,144)
		fin.Read(salt)
		salt = rsDecode(salt,rs16_128,16)
		
		nonce = make([]byte,152)
		fin.Read(nonce)
		nonce = rsDecode(nonce,rs24_128,24)
		
		fmt.Println("salt: ",salt)
		fmt.Println("nonce: ",nonce)
		
		_keyHash = make([]byte,192)
		fin.Read(_keyHash)
		_keyHash = rsDecode(_keyHash,rs64_128,64)
		//fmt.Println("keyHash",keyHash)
		
		_khash_hash = make([]byte,160)
		fin.Read(_khash_hash)
		_khash_hash = rsDecode(_khash_hash,rs32_128,32)
		//fmt.Println("crcHash",crcHash)
		
		_tmp := math.Ceil(float64(total-int64(metadataLength+1196))/float64(1048728))
		nonces = make([]byte,int(_tmp*152)+144)
		fin.Read(nonces)
		//fmt.Println("Nonces: ",nonces)
	}
	
	g.Update()
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
	
	fmt.Println("key",key)
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
		fmt.Println(kstat.Size())
		kbytes := make([]byte,kstat.Size())
		kin.Read(kbytes)
		kin.Close()
		ksha3 := sha3.New256()
		ksha3.Write(kbytes)
		khash = ksha3.Sum(nil)

		khash_sha3 := sha3.New256()
		khash_sha3.Write(khash)
		khash_hash = khash_sha3.Sum(nil)
		fmt.Println("khash",khash)
		fmt.Println("khash_hash",khash_hash)
	}
	
	//key = make([]byte,32)
	fmt.Println("output",outputFile)
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

		fout,_ = os.OpenFile(
			outputFile,
			os.O_RDWR|os.O_CREATE|os.O_TRUNC,
			0755,
		)
	}

	if keyfile{
		// XOR key and keyfile
		tmp := key
		key = make([]byte,32)
		for i,_ := range(key){
			key[i] = tmp[i]^khash[i]
		}
		fmt.Println("key",key)
	}

	
	done := 0
	counter := 0
	startTime := time.Now()

	cipher,_ := chacha20poly1305.NewX(key)
	
	if mode=="decrypt"{
		_mac := nonces[len(nonces)-144:]
		_mac = rsDecode(_mac,rs16_128,16)
		//fmt.Println("_mac ",_mac)
		nonces = nonces[:len(nonces)-144]
		var tmp []byte
		var chunk []byte
		for i,j := range(nonces){
			chunk = append(chunk,j)
			if (i+1)%152==0{
				chunk = rsDecode(chunk,rs24_128,24)
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
			_data = make([]byte,1048592)
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
				fout.Write(mac)
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
				mac := data[len(data)-16:]
				data = data[:len(data)-16]
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
		
		done += 1048576
		counter++

		progress = float32(done)/float32(total)
		
		elapsed:= float64(int64(time.Now().Sub(startTime)))/float64(1000000000)
		
		speed := (float64(done)/elapsed)/1000000
		eta := math.Abs(float64(total-int64(done))/(speed*1000000))
		
		progressInfo = fmt.Sprintf("%.2f%%",progress*100)
		
		status = fmt.Sprintf("Working at %.2f MB/s (ETA: %.1fs)",speed,eta)
		
		g.Update()
	}

	if mode=="encrypt"{
		//fmt.Println("'nonces' before RS: ",nonces)
		fout.Seek(int64(700+len(metadata)),0)
		fout.Write(rsEncode(keyHash,rs64_128,192))
		fout.Write(rsEncode(khash_hash,rs32_128,160))

		_mac,tmp := monocypher.Lock(nonces,nonce,key) 
		fmt.Println(_mac)
		//tmp := cipher.Seal(nil,nonce,nonces,nil)
		//fmt.Println("ENCRYPTED NONCES: ",tmp)
		//_mac := tmp[len(tmp)-16:]

		//tmp = tmp[:len(tmp)-16]
		var chunk []byte

		for i,j := range(tmp){
			chunk = append(chunk,j)
			if (i+1)%24==0{
				fout.Write(rsEncode(chunk,rs24_128,152))
				//fmt.Println(rsEncode(chunk,rs24_128,152))
				chunk = nil
			}
		}
		fout.Write(rsEncode(_mac,rs16_128,144))

	}else{
		//fmt.Println("crcHash: ",crcHash)
		//fmt.Println("crc.Sum: ",crc.Sum(nil))
	}
	
	fin.Close()
	fout.Close()

	if split{
		status = "Splitting file..."
		stat,_ := os.Stat(outputFile)
		size := stat.Size()
		finished := 0
		var splitted []string
		fmt.Println(size)
		chunkSize,_ := strconv.Atoi(splitSize)
		fmt.Println(splitSelected)
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
			fmt.Println(fmt.Sprintf("%s.%d",outputFile,i))
			fout,_ := os.OpenFile(
				fmt.Sprintf("%s.%d",outputFile,i),
				os.O_RDWR|os.O_CREATE|os.O_TRUNC,
				0755,
			)
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
					fmt.Println("outputFile",outputFile)
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
			g.Update()
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
	fmt.Println("==============================")
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
	fmt.Println("Exit goroutine")
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
									g.Update()
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
						g.Update()
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
		g.Update()
	}
	/*if itemSelected==1&&runtime.GOOS=="windows"{
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
	g.Update()
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
	cs_blake3 = ""

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
	if blake3_selected{
		cs_blake3 = "Calculating..."
	}

	crc_md5 := md5.New()
	crc_sha1 := sha1.New()
	crc_sha256 := sha256.New()
	crc_sha3_256 := sha3.New256()
	crc_blake2b,_ := blake2b.New256(nil)
	crc_blake2s,_ := blake2s.New256(nil)
	crc_blake3 := blake3.New()
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
		if blake3_selected{
			crc_blake3.Write(data)
		}

		done += int64(size)
		cs_progress = float32(done)/float32(total)
		g.Update()
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
	if blake3_selected{
		cs_blake3 = hex.EncodeToString(crc_blake3.Sum(nil))
	}
	g.Update()
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
	g.Update()
}

func broken(){
	working = false
	_status = "The file is either corrupted or intentionally modified."
	_status_color = color.RGBA{0xff,0x00,0x00,255}
	os.Remove(outputFile)
	debug.FreeOSMemory()
}

func rsEncode(data []byte,encoder reedsolomon.Encoder,size int) []byte{
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
	di.Init()
	if runtime.GOOS=="windows"{
		exec.Command(filepath.Join(rootDir,"sdelete64.exe"),"/accepteula")
	}
	window := g.NewMasterWindow("Picocrypt",480,488,g.MasterWindowFlagsNotResizable,nil)
	window.SetDropCallback(onDrop)
	dpi = g.Context.GetPlatform().GetContentScale()
	window.Run(startUI)
}
