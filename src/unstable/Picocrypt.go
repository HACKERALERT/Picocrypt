package main

/*

Picocrypt v1.13
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

*/

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"image/color"
	g "github.com/AllenDang/giu"
	ig "github.com/AllenDang/imgui-go"
	di "github.com/sqweek/dialog"
	"crypto/rand"
	"github.com/klauspost/reedsolomon"
	_ "golang.org/x/crypto/argon2"
	_ "github.com/HACKERALERT/Monocypher-Go/monocypher"
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
var progressInfo = ""

// User input variables
var password string
var cPassword string
var metadata string
var keep bool
var erase bool
var reedsolo bool

// Reed-Solomon encoders
var rs10_128,_ = reedsolomon.New(10,128)
var rs16_128,_ = reedsolomon.New(16,128)
var rs24_128,_ = reedsolomon.New(24,128)

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
					g.ProgressBar(0).Size(-1,0).Overlay(progressInfo),

					// Status label
					g.Dummy(10,0),
					g.Label("Ready."),

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

	// Set the output file based on mode
	if mode=="encrypt"{
		outputFile = outputEntry+".pcv"
	}else{
		outputFile = outputEntry
	}

	// If encrypting, generate values. If decrypting, read values from file
	if mode=="encrypt"{
		fout,_ := os.OpenFile(
			outputFile,
			os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
			0666,
		)

		// Argon2 salt and XChaCha20 nonce
		salt := make([]byte,16)
		nonce := make([]byte,24)

		// Encode the length of the metadata with Reed-Solomon
		metadataLength := []byte(fmt.Sprintf("%010d",len(metadata)))
		/*shards,_ := rs10_128.Split(metadataLength)
		rs10_128.Encode(shards)
		tmp := make([]byte,138)
		for i,shard := range(shards){
			tmp[i] = shard[0]
		}
		
		fout.Write(tmp)*/
		metadataLength = rsEncode(metadataLength,rs10_128,138)
		fout.Write(metadataLength)

		// Fill salt and nonce with Go's CSPRNG
		rand.Read(salt)
		rand.Read(nonce)

		// Encode salt with Reed-Solomon and write to file
		/*shards,_ = rs16_128.Split(salt)
		rs16_128.Encode(shards)
		tmp = make([]byte,144)
		for i,shard := range(shards){
			tmp[i] = shard[0]
		}
		fout.Write(tmp)*/
		salt = rsEncode(salt,rs16_128,144)
		fout.Write(salt)

		// Encode nonce with Reed-Solomon and write to file
		/*shards,_ = rs24_128.Split(nonce)
		rs24_128.Encode(shards)
		tmp = make([]byte,152)
		for i,shard := range(shards){
			tmp[i] = shard[0]
		}
		fout.Write(tmp)*/
		nonce = rsEncode(nonce,rs24_128,152)
		fout.Write(nonce)

	}else{

	}

}

// Reset the UI to a clean state with no nothing selected
func resetUI(){
	inputLabel = "Drag and drop file(s) and folder(s) into this window."
	outputEntry = ""
	orLabel = "or"
	outputWidth = 376
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

// Create the master window, set callbacks, and start the UI
func main(){
	window := g.NewMasterWindow("Picocrypt",480,470,g.MasterWindowFlagsNotResizable,nil)
	window.SetBgColor(color.RGBA{0xf5,0xf6,0xf7,255})
	window.SetDropCallback(onDrop)
	dpi = g.Context.GetPlatform().GetContentScale()
	window.Run(startUI)
}
