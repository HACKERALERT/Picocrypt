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
	_ "crypto/rand"
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

// User input variables
var password string
var cPassword string
var metadata string
var keep bool
var erase bool
var reedsolo bool

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
						g.Button("Save as"),
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
						g.Button("Start").Size(360,20),
						g.Button("Cancel").Size(95,20),
					),

					// Progress bar
					g.ProgressBar(0).Size(-1,0).Overlay("0%"),

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

	// There's only one dropped item
	if len(names)==1{
		stat,_ := os.Stat(names[0])

		// Check if dropped item is a file or a folder
		if stat.IsDir(){
			folders++
			inputLabel = "1 folder selected."

			// Add the folder
			onlyFolders = append(onlyFolders,names[0])
		}else{
			files++
			name := filepath.Base(names[0])

			// Decide if encrypting or decrypting
			if strings.HasSuffix(names[0],".pcv"){
				mode = "decrypt"
				inputLabel = name+" (will decrypt)"
				outputEntry = names[0][:len(names[0])-4]

				// Hide the ".pcv" file extension
				orLabel = "or"
				outputWidth = 376
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
	}

	// If there are folders that were dropped, recusively add all files into 'allFiles'
	if folders>0{
		for _,name := range(onlyFolders){
			filepath.Walk(name,func(path string,_ os.FileInfo,_ error) error{
				stat,_ := os.Stat(path)
				if !stat.IsDir(){
					fmt.Println(path)
					allFiles = append(allFiles,path)
				}
				return nil
			})
		}
	}

	// Update the UI
	g.Update()

	fmt.Println(onlyFiles)
	fmt.Println(onlyFolders)
	fmt.Println(allFiles)
}

// Reset the UI to a clean state with no nothing selected
func resetUI(){
	inputLabel = "Drag and drop file(s) and folder(s) into this window."
	outputEntry = ""
	orLabel = "or"
	outputWidth = 376
	g.Update()
}

// Create the master window, set callbacks, and start the UI
func main(){
	window := g.NewMasterWindow("Picocrypt",480,470,g.MasterWindowFlagsNotResizable,nil)
	window.SetBgColor(color.RGBA{0xf5,0xf6,0xf7,255})
	window.SetDropCallback(onDrop)
	dpi = g.Context.GetPlatform().GetContentScale()
	window.Run(startUI)
}
