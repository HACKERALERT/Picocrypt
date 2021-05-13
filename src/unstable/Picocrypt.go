package main

/*

Picocrypt v1.13
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

*/

import (
	"os"
	"fmt"
	"math"
	"time"
	"strings"
	"strconv"
	//"image/color"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"path/filepath"
	"crypto/sha256"
	"github.com/pkg/browser"
	"github.com/zeebo/blake3"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/argon2"
	g "github.com/AllenDang/giu"
	di "github.com/sqweek/dialog"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"github.com/atotto/clipboard"
	//"github.com/AllenDang/imgui-go"
	"github.com/klauspost/reedsolomon"
	"golang.org/x/crypto/chacha20poly1305"
	"github.com/HACKERALERT/Picocypher/monocypher"
)

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
var passwordToggleString = "Show"
var progress float32 = 0
var progressInfo = ""
var status = "Ready."

// User input variables
var password string
var cPassword string
var metadata string
var keep bool
var erase bool
var reedsolo bool
var split bool
var splitSize string
var fast bool

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

				// Label listing the input files and button to clear input files
				g.Dummy(30,0),
				g.Line(
					g.Label(inputLabel),
					g.Button("Clear").OnClick(resetUI),
				),

				// Allow user to choose a custom output path and name
				g.Dummy(10,0),
				g.Label("Save output as:"),
				g.Line(
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
				g.Dummy(10,0),
				g.Label("Password:"),
				g.Line(
					g.InputText("##password",&password).Size(200/dpi).Flags(passwordState),
					g.Button(passwordToggleString).OnClick(func(){
						if passwordState==g.InputTextFlags_Password{
							passwordState = g.InputTextFlags_None
							passwordToggleString = "Hide"
						}else{
							passwordState = g.InputTextFlags_Password
							passwordToggleString = "Show"
						}
						g.Update()
					}),
				),

				// Prompt to confirm password
				g.Dummy(10,0),
				g.Label("Confirm password:"),
				g.InputText("##cPassword",&cPassword).Size(200/dpi).Flags(passwordState),

				// Optional metadata
				g.Dummy(10,0),
				g.Label("Metadata (optional):"),
				g.InputTextMultiline("##metadata",&metadata).Size(200,80),

				// Advanced options can be enabled with checkboxes
				g.Dummy(10,0),
				g.Checkbox("Keep decrypted output even if it's corrupted or modified",&keep),
				g.Checkbox("Securely erase and delete original file(s)",&erase),
				g.Line(
					g.Checkbox("Encode with Reed-Solomon to prevent corruption",&reedsolo),
					g.Button("?").OnClick(func(){
						browser.OpenURL("https://en.wikipedia.org/wiki/Reed%E2%80%93Solomon_error_correction")
					}),
				),
				g.Line(
					g.Checkbox("Split output into chunks of",&split),
					g.InputText("##splitSize",&splitSize).Size(30).Flags(g.InputTextFlags_CharsDecimal),
					g.Label("MB"),
				),
				g.Checkbox("Fast mode (less secure, not as durable)",&fast),

				// Start and cancel buttons
				g.Dummy(10,0),
				g.Button("Start").Size(-1,20).OnClick(func(){
					go work()
				}),

				/*// Progress bar
				g.ProgressBar(progress).Size(-1,0).Overlay(progressInfo),

				// Status label
				g.Dummy(10,0),
				g.Label(status),*/

				// Credits and version
				g.Line(
					g.Label("Created by Evan Su. See the About tab for more info."),
					g.Dummy(46,0),
					g.Label("v1.13"),
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

			),

			// File checksum generator tab
			g.TabItem("Checksum generator").Layout(
				// Update 'tab' to indicate active tab
				g.Custom(func(){
					if g.IsItemActive(){
						tab = 2
					}
				}),

				g.Dummy(30,0),
				g.Label("Toggle the hashes you would like to generate and drop a file here."),
				
				// MD5
				g.Dummy(10,0),
				g.Line(
					g.Checkbox("MD5:",&md5_selected),
					g.Dummy(360,0),
					g.Button("Copy##md5").OnClick(func(){
						clipboard.WriteAll(cs_md5)
					}),
				),
				g.InputText("##cs_md5",&cs_md5).Size(-1).Flags(g.InputTextFlags_ReadOnly),

				// SHA1
				g.Dummy(10,0),
				g.Line(
					g.Checkbox("SHA1:",&sha1_selected),
					g.Dummy(353,0),
					g.Button("Copy##sha1").OnClick(func(){
						clipboard.WriteAll(cs_sha1)
					}),
				),
				g.InputText("##cs_sha1",&cs_sha1).Size(-1).Flags(g.InputTextFlags_ReadOnly),

				// SHA256
				g.Dummy(10,0),
				g.Line(
					g.Checkbox("SHA256:",&sha256_selected),
					g.Dummy(339,0),
					g.Button("Copy##sha256").OnClick(func(){
						clipboard.WriteAll(cs_sha256)
					}),
				),
				g.InputText("##cs_sha256",&cs_sha256).Size(-1).Flags(g.InputTextFlags_ReadOnly),

				// SHA3-256
				g.Dummy(10,0),
				g.Line(
					g.Checkbox("SHA3-256:",&sha3_256_selected),
					g.Dummy(325,0),
					g.Button("Copy##sha3_256").OnClick(func(){
						clipboard.WriteAll(cs_sha3_256)
					}),
				),
				g.InputText("##cs_sha3_256",&cs_sha3_256).Size(-1).Flags(g.InputTextFlags_ReadOnly),

				// BLAKE2b
				g.Dummy(10,0),
				g.Line(
					g.Checkbox("BLAKE2b:",&blake2b_selected),
					g.Dummy(332,0),
					g.Button("Copy##blake2b").OnClick(func(){
						clipboard.WriteAll(cs_blake2b)
					}),
				),
				g.InputText("##cs_blake2b",&cs_blake2b).Size(-1).Flags(g.InputTextFlags_ReadOnly),

				// BLAKE2s
				g.Dummy(10,0),
				g.Line(
					g.Checkbox("BLAKE2s:",&blake2s_selected),
					g.Dummy(332,0),
					g.Button("Copy##blake2s").OnClick(func(){
						clipboard.WriteAll(cs_blake2s)
					}),
				),
				g.InputText("##cs_blake2s",&cs_blake2s).Size(-1).Flags(g.InputTextFlags_ReadOnly),

				// BLAKE3
				g.Dummy(10,0),
				g.Line(
					g.Checkbox("BLAKE3:",&blake3_selected),
					g.Dummy(339,0),
					g.Button("Copy##blake3").OnClick(func(){
						clipboard.WriteAll(cs_blake3)
					}),
				),
				g.InputText("##cs_blake3",&cs_blake3).Size(-1).Flags(g.InputTextFlags_ReadOnly),

				// Progress bar
				g.Dummy(10,0),
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
				g.Dummy(30,0),
				g.Label("Picocrypt v1.13, created by Evan Su (https://evansu.cc)"),
			),
		),
	)
	if working{
		g.SingleWindow("Working..").IsOpen(&working).Layout(
			g.Dummy(30,0),
			g.Label("Tips:"),
			g.Label("    - Choose a strong password with more than 16 characters."),
			g.Label("    - Use a unique password that isn't used anywhere else."),

			// Progress bar
			g.ProgressBar(progress).Size(-1,0).Overlay(progressInfo),

			// Status label
			g.Dummy(10,0),
			g.Label(status),

			g.Button("Cancel").Size(95,20).OnClick(func(){
				working = false
			}),
		)
	}
}

// Handle files dropped into Picocrypt by user
func onDrop(names []string){
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
			}else{
				files++
				name := filepath.Base(names[0])

				// Decide if encrypting or decrypting
				if strings.HasSuffix(names[0],".pcv"){
					mode = "decrypt"
					inputLabel = name+" (will decrypt)"
					outputEntry = names[0][:len(names[0])-4]

					// Open input file in read-only mode
					fin,_ := os.Open(names[0])
					defer fin.Close()

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

	}else if tab==2{
		go generateChecksums(names[0])
	}

	// Update the UI
	g.Update()
}

// Start encryption/decryption
func work(){
	// Set some variables
	working = true
	//headerBroken := false
	//reedsoloFixed := 0
	//reedsoloErrors := 0
	
	// Declare salt and nonce
	var salt []byte
	var nonce []byte
	var keyHash []byte
	var crcHash []byte
	var nonces []byte
	
	stat,_ := os.Stat(inputFile)
	total := stat.Size()

	// Set the output file based on mode
	if mode=="encrypt"{
		outputFile = outputEntry+".pcv"
	}else{
		outputFile = outputEntry
	}
	
	// Open input file in read-only mode
	fin,_ := os.Open(inputFile)
	defer fin.Close()
	
	var fout *os.File

	// If encrypting, generate values. If decrypting, read values from file
	if mode=="encrypt"{
		status = "Generating values..."
		g.Update()
		fout,_ = os.OpenFile(
			outputFile,
			os.O_RDWR|os.O_CREATE|os.O_TRUNC,
			0755,
		)
		defer fout.Close()

		// Argon2 salt and XChaCha20 nonce
		salt = make([]byte,16)
		nonce = make([]byte,24)
		
		// Write version to file
		fout.Write(rsEncode([]byte("v1.13"),rs5_128,133))

		// Encode the length of the metadata with Reed-Solomon
		metadataLength := []byte(fmt.Sprintf("%010d",len(metadata)))
		metadataLength = rsEncode(metadataLength,rs10_128,138)
		// Write the length of the metadata to file
		fout.Write(metadataLength)
		
		// Write the actual metadata
		fout.Write([]byte(metadata))

		flags := make([]byte,5)
		fmt.Println(flags)
		if fast{
			flags[0] = 1
		}
		flags = rsEncode(flags,rs5_128,133)
		fout.Write(flags)

		// Fill salt and nonce with Go's CSPRNG
		rand.Read(salt)
		rand.Read(nonce)
		
		//fmt.Println("Encrypting salt: ",salt)
		//fmt.Println("Encrypting nonce: ",nonce)

		// Encode salt with Reed-Solomon and write to file
		_salt := rsEncode(salt,rs16_128,144)
		fout.Write(_salt)

		// Encode nonce with Reed-Solomon and write to file
		tmp := rsEncode(nonce,rs24_128,152)
		fout.Write(tmp)
		
		// Write placeholder for hash of key
		fout.Write(make([]byte,192))
		
		// Write placeholder for Blake3 CRC
		fout.Write(make([]byte,160))

		
		pairs := int(math.Ceil(float64(total)/1048576))
		
		offset := 152*pairs+144
		
		// Write placeholder for nonce/Poly1305 pairs
		fout.Write(make([]byte,offset))
	}else{
		g.Update()
		status = "Reading values..."
		version := make([]byte,133)
		fin.Read(version)
		version = rsDecode(version,rs5_128,5)

		tmp := make([]byte,138)
		fin.Read(tmp)
		tmp = rsDecode(tmp,rs10_128,10)
		metadataLength,_ := strconv.Atoi(string(tmp))

		fin.Read(make([]byte,metadataLength))

		flags := make([]byte,5)
		fin.Read(flags)
		flags = rsDecode(flags,rs5_128,5)
		fmt.Println(flags)
		fast = flags[0]==1

		salt = make([]byte,144)
		fin.Read(salt)
		salt = rsDecode(salt,rs16_128,16)
		
		nonce = make([]byte,152)
		fin.Read(nonce)
		nonce = rsDecode(nonce,rs24_128,24)
		
		//fmt.Println("Decrypting salt: ",salt)
		//fmt.Println("Decrypting nonce: ",nonce)
		
		keyHash = make([]byte,192)
		fin.Read(keyHash)
		keyHash = rsDecode(keyHash,rs64_128,64)
		
		crcHash = make([]byte,160)
		fin.Read(crcHash)
		crcHash = rsDecode(crcHash,rs32_128,32)
		
		_tmp := math.Ceil(float64(total-int64(metadataLength+1063))/float64(1048744))
		nonces = make([]byte,int(_tmp*152)+144)
		fin.Read(nonces)
		
		//fmt.Println("Nonces: ",nonces)
	}
	
	g.Update()
	status = "Deriving key..."
	
	fmt.Println("password",[]byte(password))
	fmt.Println("salt",salt)
	// Derive encryption/decryption key
	key := argon2.IDKey(
		[]byte(password),
		salt,
		4,
		1048576/2,
		4,
		32,
	)[:]
	fmt.Println("key",key)
	
	//key = make([]byte,32)
	
	sha3_512 := sha3.New512()
	sha3_512.Write(key)
	keyHash = sha3_512.Sum(nil)
	fmt.Println("keyHash: ",keyHash)
	
	// Check is password is correct
	
	if mode=="decrypt"{
		fout,_ = os.OpenFile(
			outputFile,
			os.O_RDWR|os.O_CREATE|os.O_TRUNC,
			0755,
		)
		defer fout.Close()
	}

	crc := blake3.New()
	
	done := 0
	counter := 0
	startTime := time.Now()

	cipher,_ := chacha20poly1305.NewX(key)
	
	if mode=="decrypt"{
		_mac := nonces[len(nonces)-144:]
		_mac = rsDecode(_mac,rs16_128,16)
		//fmt.Println("_mac len: ",len(_mac))
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
		for _,j := range(_mac){
			tmp = append(tmp,j)
		}
		//fmt.Println("ENCRYPTED NONCES: ",tmp)
		// XXXXXXXXXXXXXXXXFSFSDFFFSFF
		nonces,_ = cipher.Open(nil,nonce,tmp,nil)
		//fmt.Println("UNENCRYPTED NONCES: ",nonces)
	}
	for{
		if !working{
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
			}else{
				mac,data := monocypher.Lock(data,_nonce,key)
				fout.Write(data)
				fout.Write(mac)
			}
			crc.Write(data)
			//fout.Write(data)
		}else{
			//fmt.Println("DECODE LOOP")
			crc.Write(data)
			if fast{
				data,_ = cipher.Open(nil,_nonce,data,nil)
			}else{
				mac := data[len(data)-16:]
				data = data[:len(data)-16]
				data,_ = monocypher.Unlock(data,_nonce,key,mac)
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
		eta := float64(total-int64(done))/(speed*1000000)
		
		progressInfo = fmt.Sprintf("%.2f%%",progress*100)
		
		status = fmt.Sprintf("Working at %.2f MB/s (ETA: %.1fs)",speed,eta)
		
		g.Update()
	}

	if mode=="encrypt"{
		//fmt.Println("'nonces' before RS: ",nonces)
		fout.Seek(int64(567+len(metadata)),0)
		fout.Write(rsEncode(keyHash,rs64_128,192))
		fout.Write(rsEncode(crc.Sum(nil),rs32_128,160))
		//fmt.Println("UNENCRYPTED NONCES: ",nonces)
		tmp := cipher.Seal(nil,nonce,nonces,nil)
		//fmt.Println("ENCRYPTED NONCES: ",tmp)
		_mac := tmp[len(tmp)-16:]
		//fmt.Println("_mac len: ",len(_mac))
		tmp = tmp[:len(tmp)-16]
		var chunk []byte
		//fmt.Println("<Nonces>")
		for i,j := range(tmp){
			chunk = append(chunk,j)
			if (i+1)%24==0{
				fout.Write(rsEncode(chunk,rs24_128,152))
				//fmt.Println(rsEncode(chunk,rs24_128,152))
				chunk = nil
			}
		}
		fout.Write(rsEncode(_mac,rs16_128,144))
		//fmt.Println("</Nonces>")
	}else{
		fmt.Println("crcHash: ",crcHash)
		fmt.Println("crc.Sum: ",crc.Sum(nil))
	}
	fmt.Println("==============================")
	resetUI()
	status = "Completed."
	working = false
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
	metadata = ""
	keep = false
	erase = false
	reedsolo = false
	split = false
	splitSize = ""
	fast = false
	progress = 0
	progressInfo = ""
	g.Update()
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
	window := g.NewMasterWindow("Picocrypt",480,466,g.MasterWindowFlagsNotResizable,nil)
	window.SetDropCallback(onDrop)
	dpi = g.Context.GetPlatform().GetContentScale()
	window.Run(startUI)
}
