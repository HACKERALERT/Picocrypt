package main

/*

Picocrypt v1.14
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

*/

import (
	_ "embed"

	// Generic
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"math"
	"time"
	"sync"
	"hash"
	"image"
	"bytes"
	"regexp"
	"strings"
	"strconv"
	"runtime"
	"net/http"
	"runtime/debug"
	"image/png"
	"image/color"
	"archive/zip"
	"encoding/hex"
	"path/filepath"

	// Cryptography
	"crypto/rand"
	"crypto/hmac"
	"crypto/subtle"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/chacha20"
	"github.com/HACKERALERT/serpent" // v0.0.0-20210716182301-293b29869c66

	// GUI
	"github.com/AllenDang/giu"

	// Reed-Solomon
	"github.com/HACKERALERT/infectious" // v0.0.0-20210730231340-8af02cb9ed0a

	// Helpers
	"github.com/HACKERALERT/clipboard" // v0.1.5-0.20210716140604-61d96bf4fc94
	"github.com/HACKERALERT/dialog" // v0.0.0-20210716143851-223edea1d840
	"github.com/HACKERALERT/browser" // v0.0.0-20210730230128-85901a8dd82f
	"github.com/HACKERALERT/zxcvbn-go" // v0.0.0-20210730224720-b29e9dba62c2

)

var version = "v1.14"

//go:embed NotoSans-Regular.ttf
var font []byte
//go:embed sdelete64.exe
var sdelete64bytes []byte
//go:embed icon.png
var iconBytes []byte

// Languages
var languages = []string{
	"English",
}
var languageSelected int32

// Global variables
var dpi float32 // Used to scale properly in high-resolution displays
var mode string = "" // "encrypt", "decrypt", or ""
var working = false // True if encryption/decryption is in progress
var onlyFiles []string // Only contains files not in a folder
var onlyFolders []string // Only contains names of folders
var allFiles []string // Contains all files including files in folders
var inputFile string
var outputFile string
var recombine bool // True if decrypting and the original file was splitted during encryption
var sdelete64path string // The temporary file path where sdelete64.exe will be stored

// UI-related global variables
var tab = 0 // The index of the currently selected tab
var inputLabel = "Drag and drop file(s) and folder(s) into this window."
var outputEntry string // A modifiable text entry string variable containing path of output
var outputWidth float32 = 370
var orLabel = "or"
var passwordStrength int
var keyfile = false // True if user chooses/chose to use a keyfile
var keyfilePrompt = "Keyfile (optional):" // Changes if decrypting and keyfile was enabled
var showConfirmation = false
var showProgress = false
var progress float32 = 0 // 0 is 0%, 1 is 100%
var progressInfo = "" // Text inside the progress bar on the encrypting/decrypting window
var status = "Ready." // Status text in encrypting/decrypting window
var _status = "Ready." // Status text in main window
var _status_color = color.RGBA{0xff,0xff,0xff,255} // Changes according to status (success, fail, etc.)
var splitUnits = []string{
	"KiB",
	"MiB",
	"GiB",
} // Three choosable units for splitting output file when encrypting, in powers of 2
var splitSelected int32 // Index of which splitting unit was chosen from above
var shredProgress float32 // Progress of shredding files
var shredDone float32
var shredTotal float32 // Total files to shred (recursive)
var shredOverlay string // Text in shredding progress bar
var shredding = "Ready."

// User input variables
var password string
var cPassword string // Confirm password text entry string variable
var keyfilePath string
var keyfileLabel = "Use a keyfile (experimental)"
var metadata string
var shredTemp bool
var paranoid bool
var keep bool
var reedsolo bool
var split bool
var splitSize string
var fast bool
var kept = false // If a file was corrupted/modified, but the output was kept

// Reed-Solomon encoders
var rs1,_ = infectious.NewFEC(1,3) // 1 data shards, 3 total -> 2 parity shards
var rs5,_ = infectious.NewFEC(5,15)
var rs10,_ = infectious.NewFEC(10,30)
var rs16,_ = infectious.NewFEC(16,48)
var rs24,_ = infectious.NewFEC(24,72)
var rs32,_ = infectious.NewFEC(32,96)
var rs64,_ = infectious.NewFEC(64,192)
var rs128,_ = infectious.NewFEC(128,136)

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
			giu.Custom(func(){
				pos := giu.GetCursorPos()
				giu.Row(
					giu.Dummy(-108,0),
					giu.Combo("##language",languages[languageSelected],languages,&languageSelected).Size(100),
				).Build()
				giu.SetCursorPos(pos)

				// The tab bar, which contains different tabs for different functions
				giu.TabBar().TabItems(
					// Main file encryption/decryption tab
					giu.TabItem("Main").Layout(
						// Update 'tab' to indicate that this is the active tab
						giu.Custom(func(){
							if giu.IsItemActive(){
								tab = 0
							}
						}),
					
						// Confirm overwrite with a modal
						giu.Custom(func(){
							if showConfirmation{
								giu.PopupModal("Warning:").Layout(
									giu.Label("Output already exists. Overwrite?"),
									giu.Row(
										giu.Button("No").Size(100,0).OnClick(func(){
											giu.CloseCurrentPopup()
											showConfirmation = false
										}),
										giu.Button("Yes").Size(100,0).OnClick(func(){
											giu.CloseCurrentPopup()
											showConfirmation = false
											giu.Update()
											showProgress = true
											giu.Update()
											go func (){
												work()
												working = false
												showProgress = false
												debug.FreeOSMemory()
												giu.Update()
											}()
										}),
									),
								).Build()
								giu.OpenPopup("Warning:")
								giu.Update()
							}
						}),

						// Show encryption/decryption progress with a modal
						giu.Custom(func(){
							if showProgress{
								giu.PopupModal(" ").Layout(
									// Close modal if not working (encryption/decryption done)
									giu.Custom(func(){
										if !working{
											giu.CloseCurrentPopup()
										}
									}),
									// Progress bar
									giu.Row(
										giu.ProgressBar(progress).Size(280,0).Overlay(progressInfo),
										giu.Button("Cancel").Size(58,0).OnClick(func(){
											working = false
										}),
									),
									giu.Label(status),
								).Build()
								giu.OpenPopup(" ")
								giu.Update()
							}
						}),

						// Label listing the input files and a button to clear them
						giu.Row(
							giu.Label(inputLabel),
							giu.Row(
								giu.Dummy(-58,0),
								giu.Button("Clear").Size(50,0).OnClick(resetUI),
							),
						),

						// Allow user to choose a custom output path and/or name
						giu.Label("Save output as:"),
						giu.Row(
							giu.InputText(&outputEntry).Size(outputWidth/dpi),
							giu.Label(orLabel),
							giu.Button("Choose").OnClick(func(){
								file,_ := dialog.File().Title("Save output as...").Save()

								// Return if user canceled the file dialog
								if file==""{
									return
								}

								if len(allFiles)>1||len(onlyFolders)>0{
									// Remove the extra ".zip.pcv" extension if necessary
									if strings.HasSuffix(file,".zip.pcv"){
										file = file[:len(file)-8]
									}
								}else{
									// Remove the extra ".pcv" extension if necessary
									if strings.HasSuffix(file,".pcv"){
										file = file[:len(file)-4]
									}
								}

								outputEntry = file
							}).Size(64,0),
						),

						// Prompt for password
						giu.Row(
							giu.Label("Password:"),
							giu.Dummy(-200,0),
							giu.Label(keyfilePrompt),
						),
						giu.Row(
							giu.InputText(&password).Size(240/dpi).Flags(giu.InputTextFlagsPassword).OnChange(func(){
								passwordStrength = zxcvbn.PasswordStrength(password,nil).Score
							}),
							giu.Custom(func(){
								canvas := giu.GetCanvas()
								pos := giu.GetCursorScreenPos()

								var col color.RGBA
								switch passwordStrength{
									case 0:
										col = color.RGBA{200,76,75,255}
									case 1:
										col = color.RGBA{169,107,75,255}
									case 2:
										col = color.RGBA{138,138,75,255}
									case 3:
										col = color.RGBA{107,169,75,255}
									case 4:
										col = color.RGBA{76,200,75,255}
								}
								if password==""||mode=="decrypt"{
									col = color.RGBA{0,0,0,0}
								}

								path := pos.Add(image.Pt(
									int(math.Round(float64(6*dpi))),
									int(math.Round(float64(12*dpi))),
								))
								canvas.PathArcTo(path,8*dpi,0,float32(passwordStrength+1)/5*(2*math.Pi),-1)
								canvas.PathStroke(col,false,3)
							}),
							giu.Dummy(-200,0),
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
						giu.Row(
							giu.InputText(&cPassword).Size(240/dpi).Flags(giu.InputTextFlagsPassword),
							giu.Custom(func(){
								canvas := giu.GetCanvas()
								pos := giu.GetCursorScreenPos()
								col := color.RGBA{76,200,75,255}
								if cPassword!=password{
									col = color.RGBA{200,76,75,255}
								}
								if password==""||cPassword==""||mode=="decrypt"{
									col = color.RGBA{0,0,0,0}
								}
								path := pos.Add(image.Pt(
									int(math.Round(float64(6*dpi))),
									int(math.Round(float64(12*dpi))),
								))
								canvas.PathArcTo(path,8*dpi,0,2*math.Pi,-1)
								canvas.PathStroke(col,false,3)
							}),
							giu.Dummy(-0.0000001,0),
						),

						// Optional metadata
						giu.Label("Metadata (optional):"),
						giu.InputText(&metadata).Size(-0.0000001),

						giu.Custom(func(){
							if mode!=""{
								giu.Label("Advanced options:").Build()
							}
						}),

						// Advanced options can be enabled with checkboxes
						giu.Custom(func(){
							if mode=="encrypt"{
								giu.Checkbox("Shred temporary files (can be slow for large files)",&shredTemp).Build()
								giu.Checkbox("Fast mode (slightly less secure, not as durable)",&fast).Build()
								giu.Checkbox("Paranoid mode (extremely secure, but a bit slower)",&paranoid).Build()
								giu.Row(
									giu.Checkbox("Encode with Reed-Solomon to prevent corruption",&reedsolo),
									giu.Button("?").Size(24,25).OnClick(func(){
										browser.OpenURL("https://bit.ly/3A2V7LR")
									}),
								).Build()
								giu.Row(
									giu.Checkbox("Split output into chunks of",&split),
									giu.InputText(&splitSize).Size(50).Flags(giu.InputTextFlagsCharsDecimal),
									giu.Combo("##splitter",splitUnits[splitSelected],splitUnits,&splitSelected).Size(52),
								).Build()
								giu.Dummy(0,1).Build()
							}else if mode=="decrypt"{
								giu.Checkbox("Keep decrypted output even if it's corrupted or modified",&keep).Build()
								giu.Dummy(0,112).Build()
							}else{
								giu.Dummy(0,67).Build()
								giu.Label("                                                 No files selected yet.").Build()
								giu.Dummy(0,68).Build()
							}
						}),

						// Start button
						giu.Button("Start").Size(-0.0000001,35).OnClick(func(){
							if mode=="encrypt"&&password!=cPassword{
								_status = "Passwords don't match."
								_status_color = color.RGBA{0xff,0x00,0x00,255}
								return
							}
							if mode=="encrypt"{
								if len(allFiles)>1||len(onlyFolders)>0{
									outputFile = outputEntry+".zip.pcv"
								}else{
									outputFile = outputEntry+".pcv"
								}
							}else{
								outputFile = outputEntry
							}
							_,err := os.Stat(outputFile)
							if err==nil{
								showConfirmation = true
								giu.Update()
							}else{
								showProgress = true
								giu.Update()
								go func (){
									work()
									working = false
									showProgress = false
									debug.FreeOSMemory()
									giu.Update()
								}()
							}
						}),

						giu.Style().SetColor(giu.StyleColorText,_status_color).To(
							giu.Label(_status),
						),
					),

					// File checksum generator tab
					giu.TabItem("Checksum").Layout(
						giu.Custom(func(){
							if giu.IsItemActive(){
								tab = 1
							}
						}),

						giu.Label("Toggle the hashes you would like to generate and drop a file here."),

						// MD5
						giu.Row(
							giu.Checkbox("MD5:",&md5_selected),
							giu.Dummy(-58,0),
							giu.Button("Copy##md5").Size(50,0).OnClick(func(){
								clipboard.WriteAll(cs_md5)
							}),
						),
						giu.Style().SetColor(giu.StyleColorBorder,md5_color).To(
							giu.InputText(&cs_md5).Size(-0.0000001).Flags(giu.InputTextFlagsReadOnly),
						),

						// SHA1
						giu.Row(
							giu.Checkbox("SHA1:",&sha1_selected),
							giu.Dummy(-58,0),
							giu.Button("Copy##sha1").Size(50,0).OnClick(func(){
								clipboard.WriteAll(cs_sha1)
							}),
						),
						giu.Style().SetColor(giu.StyleColorBorder,sha1_color).To(
							giu.InputText(&cs_sha1).Size(-0.0000001).Flags(giu.InputTextFlagsReadOnly),
						),

						// SHA256
						giu.Row(
							giu.Checkbox("SHA256:",&sha256_selected),
							giu.Dummy(-58,0),
							giu.Button("Copy##sha256").Size(50,0).OnClick(func(){
								clipboard.WriteAll(cs_sha256)
							}),
						),
						giu.Style().SetColor(giu.StyleColorBorder,sha256_color).To(
							giu.InputText(&cs_sha256).Size(-0.0000001).Flags(giu.InputTextFlagsReadOnly),
						),

						// SHA3-256
						giu.Row(
							giu.Checkbox("SHA3-256:",&sha3_256_selected),
							giu.Dummy(-58,0),
							giu.Button("Copy##sha3_256").Size(50,0).OnClick(func(){
								clipboard.WriteAll(cs_sha3_256)
							}),
						),
						giu.Style().SetColor(giu.StyleColorBorder,sha3_256_color).To(
							giu.InputText(&cs_sha3_256).Size(-0.0000001).Flags(giu.InputTextFlagsReadOnly),
						),

						// BLAKE2b
						giu.Row(
							giu.Checkbox("BLAKE2b:",&blake2b_selected),
							giu.Dummy(-58,0),
							giu.Button("Copy##blake2b").Size(50,0).OnClick(func(){
								clipboard.WriteAll(cs_blake2b)
							}),
						),
						giu.Style().SetColor(giu.StyleColorBorder,blake2b_color).To(
							giu.InputText(&cs_blake2b).Size(-0.0000001).Flags(giu.InputTextFlagsReadOnly),
						),

						// BLAKE2s
						giu.Row(
							giu.Checkbox("BLAKE2s:",&blake2s_selected),
							giu.Dummy(-58,0),
							giu.Button("Copy##blake2s").Size(50,0).OnClick(func(){
								clipboard.WriteAll(cs_blake2s)
							}),
						),
						giu.Style().SetColor(giu.StyleColorBorder,blake2s_color).To(
							giu.InputText(&cs_blake2s).Size(-0.0000001).Flags(giu.InputTextFlagsReadOnly),
						),
					
						// Input entry for validating a checksum
						giu.Label("Validate a checksum:"),
						giu.InputText(&cs_validate).Size(-0.0000001).OnChange(func(){
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
						giu.ProgressBar(cs_progress).Size(-0.0000001,0),
					),

					// File shredder tab
					giu.TabItem("Shredder").Layout(
						giu.Custom(func(){
							if giu.IsItemActive(){
								tab = 2
							}
						}),

						giu.Label("Drop file(s) and folder(s) here to shred them."),
						giu.ProgressBar(shredProgress).Overlay(shredOverlay).Size(-0.0000001,0),
						giu.Custom(func(){
							if len(shredding)>50{
								shredding = "....."+shredding[len(shredding)-50:]
							}
							giu.Label(shredding).Wrapped(true).Build()
						}),
					),

					// About tab
					giu.TabItem("About").Layout(
						giu.Custom(func(){
							if giu.IsItemActive(){
								tab = 3
							}
						}),
						giu.Label("Picocrypt "+version+", created by Evan Su (https://evansu.cc)"),
					),
				).Build()
			}),
		),
	)
}

// Handle files dropped into Picocrypt by user
func onDrop(names []string){
	_status = "Ready."
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
		outputWidth = 370

		// There's only one dropped item
		if len(names)==1{
			stat,_ := os.Stat(names[0])

			// Check if dropped item is a file or a folder
			if stat.IsDir(){
				folders++
				inputLabel = "1 folder selected."

				// Add the folder
				onlyFolders = append(onlyFolders,names[0])

				// Set 'outputEntry' to 'Encrypted'
				outputEntry = filepath.Join(filepath.Dir(names[0]),"Encrypted")
				
				mode = "encrypt"
				// Show the ".zip.pcv" file extension
				orLabel = ".zip.pcv  or"
				outputWidth = 317
			}else{
				files++
				name := filepath.Base(names[0])

				nums := []string{"0","1","2","3","4","5","6","7","8","9"}
				endsNum := false
				for _,i := range nums{
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
						inputLabel = name+" (will recombine and decrypt)"
						ind := strings.Index(names[0],".pcv")
						names[0] = names[0][:ind]
						outputEntry = names[0]
						recombine = true
					}else{
						outputEntry = names[0][:len(names[0])-4]
					}

					// Open input file in read-only mode
					fin,_ := os.Open(names[0])

					// Use regex to test if input is a valid Picocrypt volume
					tmp := make([]byte,30)
					fin.Read(tmp)
					if string(tmp[:5])=="v1.13"{
						resetUI()
						_status = "Please use Picocrypt v1.13 to decrypt this file."
						_status_color = color.RGBA{0xff,0x00,0x00,255}
						fin.Close()
						return
					}
					if valid,_:=regexp.Match(`^v\d\.\d{2}.{10}0?\d+`,tmp);!valid{
						resetUI()
						_status = "This doesn't seem to be a Picocrypt volume."
						_status_color = color.RGBA{0xff,0x00,0x00,255}
						fin.Close()
						return
					}
					fin.Seek(0,0)

					// Read metadata and insert into box
					var err error
					fin.Read(make([]byte,15))
					tmp = make([]byte,15)
					fin.Read(tmp)
					tmp,err = rsDecode(rs5,tmp)

					if err==nil{
						metadataLength,_ := strconv.Atoi(string(tmp))
						tmp = make([]byte,metadataLength*3)
						fin.Read(tmp)
						metadata = ""

						for i:=0;i<metadataLength*3;i+=3{
							//fmt.Println(tmp[i:i+3])
							t,err := rsDecode(rs1,tmp[i:i+3])
							if err!=nil{
								metadata = "Metadata is corrupted."
								break
							}
							metadata += string(t)
						}
					}else{
						metadata = "Metadata is corrupted."
					}

					flags := make([]byte,15)
					fin.Read(flags)
					flags,err = rsDecode(rs5,flags)
					if err!=nil{
						_status = "File is corrupt and cannot be decrypted."
						_status_color = color.RGBA{0xff,0x00,0x00,255}
						return
					}

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
					orLabel = ".pcv  or"
					outputWidth = 338
				}

				// Add the file
				onlyFiles = append(onlyFiles,names[0])
				inputFile = names[0]
			}
		}else{
			mode = "encrypt"
			// Show the ".zip.pcv" file extension
			orLabel = ".zip.pcv  or"
			outputWidth = 317

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

			// Set 'outputEntry' to 'Encrypted'
			outputEntry = filepath.Join(filepath.Dir(names[0]),"Encrypted")
		}

		// If there are folders that were dropped, recusively add all files into 'allFiles'
		if folders>0{
			for _,name := range onlyFolders{
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
		go generateChecksums(names[0])
	}else if tab==2{
		go shred(names,true)
	}

	// Update the UI
	giu.Update()
}

// Start encryption/decryption
func work(){
	status = "Starting..."
	// Set some variables
	working = true
	//headerBroken := false
	//reedsoloFixed := 0
	//reedsoloErrors := 0
	var salt []byte
	var hkdfSalt []byte
	var serpentSalt []byte
	var nonce []byte
	var keyHash []byte
	var _keyHash []byte
	var khash []byte
	var khash_hash []byte = make([]byte,32)
	var _khash_hash []byte
	var fileMac []byte
	
	// Set the output file based on mode
	if mode=="encrypt"{
		status = "Combining files..."

		// "Tar" files into a zip archive with a compression level of 0 (store)
		if len(allFiles)>1||len(onlyFolders)>0{
			var rootDir string
			if len(onlyFolders)>0{
				rootDir = filepath.Dir(onlyFolders[0])
			}else{
				rootDir = filepath.Dir(onlyFiles[0])
			}

			inputFile = outputEntry+".zip"
			outputFile = inputFile+".pcv"
			//fmt.Println(inputFile)
			file,_ := os.Create(inputFile)
			
			w := zip.NewWriter(file)
			for i,path := range allFiles{
				if !working{
					w.Close()
					file.Close()
					os.Remove(inputFile)
					_status = "Operation cancelled by user."
					_status_color = color.RGBA{0xff,0xff,0xff,255}
					return
				}
				progressInfo = fmt.Sprintf("%d/%d",i,len(allFiles))
				progress = float32(i)/float32(len(allFiles))
				giu.Update()
				if path==inputFile{
					continue
				}
				stat,_ := os.Stat(path)
				header,_ := zip.FileInfoHeader(stat)
				header.Name = strings.TrimPrefix(path,rootDir)

				// When Windows contradicts itself :)
				if runtime.GOOS=="windows"{
					header.Name = strings.ReplaceAll(header.Name,"\\","/")
					header.Name = strings.TrimPrefix(header.Name,"/")
				}

				header.Method = zip.Store
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
	
	stat,_ := os.Stat(inputFile)
	total := stat.Size()

	// XChaCha20's max message size is 256 GiB
	if total>256*1073741824{
		_status = "Total size is larger than 256 GiB, XChaCha20's limit."
		_status_color = color.RGBA{0xff,0x00,0x00,255}
		return
	}
	
	// Open input file in read-only mode
	fin,_ := os.Open(inputFile)
	var fout *os.File

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
		hkdfSalt = make([]byte,32)
		serpentSalt = make([]byte,16)
		nonce = make([]byte,24)
		
		// Write version to file
		fout.Write(rsEncode(rs5,[]byte(version)))

		// Encode the length of the metadata with Reed-Solomon
		metadataLength := []byte(fmt.Sprintf("%05d",len(metadata)))
		//fmt.Println("metadataLength:",metadataLength)
		metadataLength = rsEncode(rs5,metadataLength)

		fmt.Println(metadataLength)
		
		// Write the length of the metadata to file
		fout.Write(metadataLength)

		// Reed-Solomon-encode the metadata and write to file
		for _,i := range []byte(metadata){
			fout.Write(rsEncode(rs1,[]byte{i}))
		}

		flags := make([]byte,5)
		if fast{
			flags[0] = 1
		}
		if paranoid{
			flags[1] = 1
		}
		if keyfile{
			flags[2] = 1
		}
		if reedsolo{
			flags[3] = 1
		}
		//fmt.Println("flags:",flags)
		flags = rsEncode(rs5,flags)
		fout.Write(flags)

		// Fill salts and nonce with Go's CSPRNG
		rand.Read(salt)
		rand.Read(hkdfSalt)
		rand.Read(serpentSalt)
		rand.Read(nonce)

		// Encode salt with Reed-Solomon and write to file
		_salt := rsEncode(rs16,salt)
		fout.Write(_salt)

		// Encode HKDF salt with Reed-Solomon and write to file
		_hkdfSalt := rsEncode(rs32,hkdfSalt)
		fout.Write(_hkdfSalt)

		// Encode Serpent salt with Reed-Solomon and write to file
		_serpentSalt := rsEncode(rs16,serpentSalt)
		fout.Write(_serpentSalt)

		// Encode nonce with Reed-Solomon and write to file
		_nonce := rsEncode(rs24,nonce)
		fout.Write(_nonce)
		
		// Write placeholder for hash of key
		fout.Write(make([]byte,192))
		
		// Write placeholder for hash of hash of keyfile
		fout.Write(make([]byte,96))

		// Write placeholder for HMAC-BLAKE2b/HMAC-SHA3 of file
		fout.Write(make([]byte,192))
	}else{
		var err1 error
		var err2 error
		var err3 error
		var err4 error
		var err5 error
		var err6 error
		var err7 error
		var err8 error
		var err9 error
		var err10 error

		status = "Reading values..."
		giu.Update()
		version := make([]byte,15)
		fin.Read(version)
		_,err1 = rsDecode(rs5,version)

		tmp := make([]byte,15)
		fin.Read(tmp)
		tmp,err2 = rsDecode(rs5,tmp)
		metadataLength,_ := strconv.Atoi(string(tmp))

		fin.Read(make([]byte,metadataLength*3))

		flags := make([]byte,15)
		fin.Read(flags)
		flags,err3 = rsDecode(rs5,flags)
		fast = flags[0]==1
		paranoid = flags[1]==1
		keyfile = flags[2]==1

		salt = make([]byte,48)
		fin.Read(salt)
		salt,err4 = rsDecode(rs16,salt)

		hkdfSalt = make([]byte,96)
		fin.Read(hkdfSalt)
		hkdfSalt,err5 = rsDecode(rs32,hkdfSalt)

		serpentSalt = make([]byte,48)
		fin.Read(serpentSalt)
		serpentSalt,err6 = rsDecode(rs16,serpentSalt)
		
		nonce = make([]byte,72)
		fin.Read(nonce)
		nonce,err7 = rsDecode(rs24,nonce)
		
		_keyHash = make([]byte,192)
		fin.Read(_keyHash)
		_keyHash,err8 = rsDecode(rs64,_keyHash)
		
		_khash_hash = make([]byte,96)
		fin.Read(_khash_hash)
		_khash_hash,err9 = rsDecode(rs32,_khash_hash)

		fileMac = make([]byte,192)
		fin.Read(fileMac)
		fileMac,err10 = rsDecode(rs64,fileMac)

		// Is there a better way?
		if err1!=nil||err2!=nil||err3!=nil||err4!=nil||err5!=nil||err6!=nil||err7!=nil||err8!=nil||err9!=nil||err10!=nil{
			if keep{
				kept = true
			}else{
				_status = "The header is corrupt and the file cannot be decrypted."
				_status_color = color.RGBA{0xff,0x00,0x00,255}
				fin.Close()
				return
			}
		}
	}

	
	status = "Deriving key..."
	progress = 0
	progressInfo = ""
	giu.Update()
	
	// Derive encryption/decryption key and subkeys
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
		fin.Close()
		fout.Close()
		if mode=="encrypt"&&(len(allFiles)>1||len(onlyFolders)>0){
			os.Remove(outputEntry+".zip")
		}
		if recombine{
			os.Remove(inputFile)
		}
		os.Remove(outputFile)
		return
	}

	if keyfile{
		kin,_ := os.Open(keyfilePath)
		kstat,_ := os.Stat(keyfilePath)
		kbytes := make([]byte,kstat.Size())
		kin.Read(kbytes)
		kin.Close()
		ksha3 := sha3.New256()
		ksha3.Write(kbytes)
		khash = ksha3.Sum(nil)

		khash_sha3 := sha3.New256()
		khash_sha3.Write(khash)
		khash_hash = khash_sha3.Sum(nil)
	}
	
	sha3_512 := sha3.New512()
	sha3_512.Write(key)
	keyHash = sha3_512.Sum(nil)

	//fmt.Println(keyHash,_keyHash)
	
	// Check is password is correct
	if mode=="decrypt"{
		keyCorrect := true
		keyfileCorrect := true
		var tmp bool
		if subtle.ConstantTimeCompare(keyHash,_keyHash)==0{
			keyCorrect = false
		}
		if keyfile{
			if subtle.ConstantTimeCompare(khash_hash,_khash_hash)==0{
				keyfileCorrect = false
			}
			tmp = !keyCorrect||!keyfileCorrect
		}else{
			tmp = !keyCorrect
		}
		if tmp||keep{
			if keep{
				kept = true
			}else{
				fin.Close()
				if !keyCorrect{
					_status = "The provided password is incorrect."
				}else{
					_status = "The provided keyfile is incorrect."
				}
				_status_color = color.RGBA{0xff,0x00,0x00,255}
				key = nil
				if recombine{
					os.Remove(inputFile)
				}
				os.Remove(outputFile)
				return
			}
		}

		fout,_ = os.Create(outputFile)
	}

	if keyfile{
		// XOR key and keyfile
		tmp := key
		key = make([]byte,32)
		for i,_ := range key{
			key[i] = tmp[i]^khash[i]
		}
		//fmt.Println("key",key)
	}

	
	done := 0
	counter := 0
	startTime := time.Now()
	chacha20,_ := chacha20.NewUnauthenticatedCipher(key,nonce)

	// Use HKDF-SHA3 to generate a subkey
	var mac hash.Hash
	subkey := make([]byte,32)
	hkdf := hkdf.New(sha3.New256,key,hkdfSalt,nil)
	hkdf.Read(subkey)
	if fast{
		// Keyed BLAKE2b
		mac,_ = blake2b.New512(subkey)
	}else{
		// HMAC-SHA3
		mac = hmac.New(sha3.New512,subkey)
	}

	// Generate another subkey and cipher (not used unless paranoid mode is checked
	serpentKey := make([]byte,32)
	hkdf.Read(serpentKey)
	_serpent,_ := serpent.NewCipher(serpentKey)
	serpentCTR := cipher.NewCTR(_serpent,serpentSalt)

	for{
		if !working{
			_status = "Operation cancelled by user."
			_status_color = color.RGBA{0xff,0xff,0xff,255}
			fin.Close()
			fout.Close()
			if mode=="encrypt"&&(len(allFiles)>1||len(onlyFolders)>0){
				os.Remove(outputEntry+".zip")
			}
			if recombine{
				os.Remove(inputFile)
			}
			os.Remove(outputFile)
			return
		}

		//var _data []byte
		var data []byte
		if reedsolo{
			data = make([]byte,1048576)
		}else{
			data = make([]byte,1114112)
		}

		size,err := fin.Read(data)
		if err!=nil{
			break
		}
		data = data[:size]
		_data := make([]byte,len(data))

		// "Actual" encryption is done in the next couple of lines
		fmt.Println(data[:10])
		if mode=="encrypt"{
			if paranoid{
				serpentCTR.XORKeyStream(_data,data)
				fmt.Println(_data[:10])
				copy(data,_data)
				fmt.Println(data[:10])
			}
			chacha20.XORKeyStream(_data,data)
			if reedsolo{
				if len(data)==1048576{
					copy(data,_data)
					_data = nil
					for i:=0;i<1048576;i+=128{
						tmp := data[i:i+128]
						tmp := rsEncode(rs128,tmp)
						for _,j := range tmp{
							_data = append(_data,j)
						}
					}
				}else{

				}
			}
			mac.Write(_data)
		}else{
			mac.Write(data)
			var _data []byte
			if reedsolo{
				if len(data)==1114112{
					copy(_data,data)
					data = nil
					for i:=0;i<1114112;i+=136{
						tmp := data[i:i+136]
						tmp,err := rsDecode(rs128,tmp)
						if err!=nil{
							if keep{
								kept = true
							}else{
								_status = "The input file is too corrupted to decrypt."
								_status_color = color.RGBA{0xff,0x00,0x00,255}
								fin.Close()
								fout.Close()
								broken()
								return
							}
						}
						for _,j := range tmp{
							data = append(data,j)
						}
					}
				}
			}
			chacha20.XORKeyStream(_data,data)
			if paranoid{
				copy(data,_data)
				serpentCTR.XORKeyStream(_data,data)
			}
		}
		fmt.Println("=========")
		fout.Write(_data)
	
		// Update statistics
		if mode=="decrypt"&&reedsolo{
			done += 1114112
		}else{
			done += 1048576
		}
		counter++
		progress = float32(done)/float32(total)
		elapsed:= float64(int64(time.Now().Sub(startTime)))/float64(1000000000)
		speed := (float64(done)/elapsed)/1000000
		eta := math.Abs(float64(total-int64(done))/(speed*1000000))
		
		if progress>1{
			progress = 1
		}

		progressInfo = fmt.Sprintf("%.2f%%",progress*100)
		status = fmt.Sprintf("Working at %.2f MB/s (ETA: %.1fs)",speed,eta)
		giu.Update()
	}

	if mode=="encrypt"{
		fout.Seek(int64(309+len(metadata)*3),0)
		fout.Write(rsEncode(rs64,keyHash))
		fout.Write(rsEncode(rs32,khash_hash))
		fout.Write(rsEncode(rs64,mac.Sum(nil)))
	}else{
		if subtle.ConstantTimeCompare(mac.Sum(nil),fileMac)==0{
			if keep{
				kept = true
			}else{
				fin.Close()
				fout.Close()
				broken()
				return
			}
		}
	}
	fmt.Println(mac.Sum(nil),fileMac)
	
	fin.Close()
	fout.Close()

	// Split files into chunks
	if split{
		status = "Splitting file..."
		stat,_ := os.Stat(outputFile)
		size := stat.Size()
		finished := 0
		var splitted []string
		chunkSize,_ := strconv.Atoi(splitSize)

		// User can choose KiB, MiB, and GiB
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

					// If user cancels, remove the unfinished files
					for _,j := range splitted{
						os.Remove(j)
					}
					os.Remove(fmt.Sprintf("%s.%d",outputFile,i))
					if len(allFiles)>1||len(onlyFolders)>0{
						os.Remove(outputEntry+".zip")
					}
					os.Remove(outputFile)
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
		if shredTemp{
			progressInfo = ""
			status = "Shredding temporary files..."
			fmt.Println(outputFile)
			shred([]string{outputFile}[:],false)
		}else{
			os.Remove(outputFile)
		}
	}

	// Remove the temporary file used to combine a splitted Picocrypt volume
	if recombine{
		os.Remove(inputFile)
	}

	// Delete the temporary zip file if user wishes
	if len(allFiles)>1||len(onlyFolders)>0{
		if shredTemp{
			progressInfo = ""
			status = "Shredding temporary files..."
			giu.Update()
			shred([]string{outputEntry+".zip"}[:],false)
		}else{
			os.Remove(outputEntry+".zip")
		}
	}

	resetUI()

	// If user chose to keep a corrupted/modified file, let them know
	if kept{
		_status = "The input is corrupted and/or modified. Please be careful."
		_status_color = color.RGBA{0xff,0xff,0x00,255}
	}else{
		_status = "Completed."
		_status_color = color.RGBA{0x00,0xff,0x00,255}
	}

	// Clear UI state
	working = false
	kept = false
	key = nil
	status = "Ready."
}

// This function is run if an issue occurs during decryption
func broken(){
	_status = "The file is either corrupted or intentionally modified."
	_status_color = color.RGBA{0xff,0x00,0x00,255}
	if recombine{
		os.Remove(inputFile)
	}
	os.Remove(outputFile)
	giu.Update()
}

// Generate file checksums (pretty straightforward)
func generateChecksums(file string){
	fin,_ := os.Open(file)

	// Clear UI state
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

	// Create the checksum objects
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

// Recursively shred all file(s) and folder(s) passed in as 'names'
func shred(names []string,separate bool){
	shredTotal = 0
	shredDone = 0

	// 'separate' is true if this function is being called from the encryption/decryption tab
	if separate{
		shredOverlay = "Shredding..."
	}

	// Walk through directories to get the total number of files for statistics
	for _,name := range names{
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

	for _,name := range names{
		shredding = name

		// Linux and macOS need a command with similar syntax and usage, so they're combined
		if runtime.GOOS=="linux"||runtime.GOOS=="darwin"{
			stat,_ := os.Stat(name)
			if stat.IsDir(){
				var coming []string

				// Walk the folder recursively
				filepath.Walk(name,func(path string,_ os.FileInfo,err error) error{
					if err!=nil{
						return nil
					}
					stat,_ := os.Stat(path)
					if !stat.IsDir(){
						if len(coming)==128{
							// Use a WaitGroup to parallelize shredding
							var wg sync.WaitGroup
							for i,j := range coming{
								wg.Add(1)
								go func(wg *sync.WaitGroup,id int,j string){
									defer wg.Done()
									cmd := exec.Command("")
									if runtime.GOOS=="linux"{
										cmd = exec.Command("shred","-ufvz","-n","3",j)
									}else{
										cmd = exec.Command("rm","-rfP",j)
									}
									cmd.Run()
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
				for _,i := range coming{
					go func(){
						cmd := exec.Command("")
						if runtime.GOOS=="linux"{
							cmd = exec.Command("shred","-ufvz","-n","3",i)
						}else{
							cmd = exec.Command("rm","-rfP",i)
						}
						cmd.Run()
						shredding = i
						shredDone++
						shredUpdate(separate)
						giu.Update()
					}()
				}
				os.RemoveAll(name)
			}else{ // The path is a file, not a directory, so just shred it
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
				// Walk the folder recursively
				filepath.Walk(name,func(path string,_ os.FileInfo,err error) error{
					if err!=nil{
						return nil
					}
					fmt.Println(path)
					stat,_ := os.Stat(path)
					if stat.IsDir(){
						t := 0
						files,_ := ioutil.ReadDir(path)
						for _,f := range files{
							if !f.IsDir(){
								t++
							}
						}
						shredDone += float32(t)
						shredUpdate(separate)
						cmd := exec.Command(sdelete64path,"*","-p","4")
						cmd.Dir = path
						cmd.Run()
						shredding = strings.ReplaceAll(path,"\\","/")+"/*"
					}
					return nil
				})
				// sdelete64 doesn't delete the empty folder, so I'll do it manually
				os.RemoveAll(name)
			}else{
				exec.Command(sdelete64path,name,"-p","4").Run()
				shredDone++
				shredUpdate(separate)
			}
		}
		giu.Update()
	}
	// Clear UI state
	shredding = "Ready."
	shredProgress = 0
	shredOverlay = ""
}

// Update shredding statistics
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

// Reset the UI to a clean state with nothing selected or checked
func resetUI(){
	mode = ""
	inputLabel = "Drag and drop file(s) and folder(s) into this window."
	outputEntry = ""
	orLabel = "or"
	outputWidth = 370
	password = ""
	cPassword = ""
	keyfilePrompt = "Keyfile (optional):"
	keyfileLabel = "Use a keyfile"
	keyfile = false
	metadata = ""
	shredTemp = false
	keep = false
	reedsolo = false
	split = false
	splitSize = ""
	fast = false
	paranoid = false
	progress = 0
	progressInfo = ""
	_status = "Ready."
	_status_color = color.RGBA{0xff,0xff,0xff,255}
	giu.Update()
}

// Reed-Solomon encoder
func rsEncode(rs *infectious.FEC,data []byte) []byte{
	var res []byte
	rs.Encode(data,func(s infectious.Share){
		res = append(res,s.DeepCopy().Data[0])
	})
	return res
}

// Reed-Solomon decoder
func rsDecode(rs *infectious.FEC,data []byte) ([]byte,error){
	tmp := make([]infectious.Share,rs.Total())
	for i:=0;i<rs.Total();i++{
		tmp[i] = infectious.Share{
			Number:i,
			Data:[]byte{data[i]},
		}
	}
	res,err := rs.Decode(nil,tmp)
	if err!=nil{
		return data[:rs.Total()/3],err
	}
	return res,nil
}

// PKCS7 Pad (for use with Reed-Solomon, not for cryptographic purposes)
func pad(data []byte) []byte{
	padLen := 128-len(data)%128
	padding := bytes.Repeat([]byte{byte(padLen)},padLen)
	return append(data,padding...)
}

// PKCS7 Unpad
func unpad(data []byte) []byte{
	length := len(data)
	padLen := int(data[length-1])
	return data[:length-padLen]
}

func main(){
	// Create a temporary file to store sdelete64.exe
	sdelete64,_ := os.CreateTemp("","sdelete64.*.exe")
	sdelete64path = sdelete64.Name()
	sdelete64.Write(sdelete64bytes)
	sdelete64.Close()
	exec.Command(sdelete64path,"/accepteula").Run()

	// Start a goroutine to check if a newer version is available
	go func(){
		v,err := http.Get("https://raw.githubusercontent.com/HACKERALERT/Picocrypt/main/internals/version.txt")
		if err==nil{
			body,err := io.ReadAll(v.Body)
			v.Body.Close()
			if err==nil{
				if string(body[:5])!=version{
					_status = "A newer version is available."
					_status_color = color.RGBA{0,255,0,255}
				}
			}
		}
	}()

	// Initialize the dialog helper, set default font to NotoSans-Regular
	dialog.Init()
	giu.SetDefaultFontFromBytes(font,18)

	// Create giu window, set window icon
	window := giu.NewMasterWindow("Picocrypt",480,502,giu.MasterWindowFlagsNotResizable)
	r := bytes.NewReader(iconBytes)
	icon,_ := png.Decode(r)
	window.SetIcon([]image.Image{icon})

	// Add callbacks, set the screen DPI, start the UI
	window.SetDropCallback(onDrop)
	window.SetCloseCallback(func() bool{
		// Disable closing window if a Picocrypt is working to prevent temporary files
		if working||shredding!="Ready."{
			return false
		}
		return true
	})
	dpi = giu.Context.GetPlatform().GetContentScale()
	window.Run(startUI)

	// Window closed, clean up
	os.Remove(sdelete64path)
}
