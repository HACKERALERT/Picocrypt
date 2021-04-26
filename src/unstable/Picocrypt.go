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
	"image/color"
	"crypto/rand"
	"path/filepath"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/blake2b"
	g "github.com/AllenDang/giu"
	di "github.com/sqweek/dialog"
	ig "github.com/AllenDang/imgui-go"
	"github.com/klauspost/reedsolomon"
	"golang.org/x/crypto/chacha20poly1305"
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
var inputLabel = "Drag and drop file(s) and folder(s) into this window."
var outputEntry string
var outputWidth float32 = 376
var orLabel = "or"
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

// Reed-Solomon encoders
var rs5_128,_ = reedsolomon.New(5,128)
var rs10_128,_ = reedsolomon.New(10,128)
var rs16_128,_ = reedsolomon.New(16,128)
var rs24_128,_ = reedsolomon.New(24,128)
var rs32_128,_ = reedsolomon.New(32,128)
var rs64_128,_ = reedsolomon.New(64,128)

// Create the user interface
func startUI(){
	g.SingleWindow("Picocrypt").Layout(
		// Some styling for aesthetics
		g.Style().SetColor(
			ig.StyleColorBorder,
			color.RGBA{0x06,0x34,0x55,255},
		).To(
			// The tab bar, which contains different tabs for different features
			g.TabBar("TabBar").Layout(
				// File encryption/decryption tab
				g.TabItem("Encryption/decryption").Layout(
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
					g.InputText("##password",&password).Size(200/dpi),

					// Prompt to confirm password
					g.Dummy(10,0),
					g.Label("Confirm password:"),
					g.InputText("##cPassword",&cPassword).Size(200/dpi),

					// Optional metadata
					g.Dummy(10,0),
					g.Label("Metadata (optional):"),
					g.InputTextMultiline("##metadata",&metadata).Size(200,80),

					// Advanced options can be enabled with checkboxes
					g.Dummy(10,0),
					g.Checkbox("Keep decrypted output even if it's corrupted or modified",&keep),
					g.Checkbox("Securely erase and delete original file(s)",&erase),
					g.Checkbox("Encode with Reed-Solomon to prevent corruption",&reedsolo),

					// Start and cancel buttons
					g.Dummy(10,0),
					g.Line(
						g.Button("Start").Size(360,20).OnClick(func(){
							go work()
						}),
						g.Button("Cancel").Size(95,20),
					),

					// Progress bar
					g.ProgressBar(progress).Size(-1,0).Overlay(progressInfo),

					// Status label
					g.Dummy(10,0),
					g.Label(status),

					// Credits and version
					g.Line(
						g.Label("Created by Evan Su."),
						g.Label("v1.13"),
					),
				),

				// File shredder tab
				g.TabItem("Shredder").Layout(

				),

				// File checksum generator tab
				g.TabItem("Checksum generator").Layout(

				),
			),
		),
	)
}

// Handle files dropped into Picocrypt by user
func onDrop(names []string){
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

			// Set the file as 'outputEntry'
			outputEntry = names[0]
			
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

		// Fill salt and nonce with Go's CSPRNG
		rand.Read(salt)
		rand.Read(nonce)
		
		//fmt.Println("Encrypting salt: ",salt)
		//fmt.Println("Encrypting nonce: ",nonce)

		// Encode salt with Reed-Solomon and write to file
		salt = rsEncode(salt,rs16_128,144)
		fout.Write(salt)

		// Encode nonce with Reed-Solomon and write to file
		tmp := rsEncode(nonce,rs24_128,152)
		fout.Write(tmp)
		
		// Write placeholder for hash of key
		fout.Write(make([]byte,192))
		
		// Write placeholder for Blake2 CRC
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
		
		_tmp := math.Ceil(float64(total-int64(metadataLength+919))/float64(1048744))
		nonces = make([]byte,int(_tmp*152)+144)
		fin.Read(nonces)
		
		//fmt.Println("Nonces: ",nonces)
	}
	
	g.Update()
	status = "Deriving key..."
	
	// Derive encryption/decryption key
	key := argon2.IDKey(
		[]byte(password),
		salt,
		4,
		1048576/2,
		4,
		32,
	)[:]
	
	key = make([]byte,32)
	
	sha3_512 := sha3.New512()
	sha3_512.Write(key)
	keyHash = sha3_512.Sum(nil)
	//fmt.Println("keyHash: ",keyHash)
	
	// Check is password is correct
	
	if mode=="decrypt"{
		fout,_ = os.OpenFile(
			outputFile,
			os.O_RDWR|os.O_CREATE|os.O_TRUNC,
			0755,
		)
		defer fout.Close()
	}

	crc,_ := blake2b.New256(nil)
	
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
			data = cipher.Seal(nil,_nonce,data,nil)
			crc.Write(data)
			fout.Write(data)
		}else{
			//fmt.Println("DECODE LOOP")
			crc.Write(data)
			data,_ := cipher.Open(nil,_nonce,data,nil)
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
		data = nil
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
	window := g.NewMasterWindow("Picocrypt",480,470,g.MasterWindowFlagsNotResizable,nil)
	window.SetBgColor(color.RGBA{0xf5,0xf6,0xf7,255})
	window.SetDropCallback(onDrop)
	dpi = g.Context.GetPlatform().GetContentScale()
	window.Run(startUI)
}
