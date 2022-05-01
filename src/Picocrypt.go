package main

/*

Picocrypt v1.27
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

*/

import (
	_ "embed"

	"archive/zip"
	"bytes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"hash"
	"image"
	"image/color"
	"io"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/HACKERALERT/clipboard"
	"github.com/HACKERALERT/crypto/argon2"
	"github.com/HACKERALERT/crypto/blake2b"
	"github.com/HACKERALERT/crypto/chacha20"
	"github.com/HACKERALERT/crypto/hkdf"
	"github.com/HACKERALERT/crypto/sha3"
	"github.com/HACKERALERT/dialog"
	"github.com/HACKERALERT/giu"
	"github.com/HACKERALERT/infectious"
	"github.com/HACKERALERT/serpent"
	"github.com/HACKERALERT/zxcvbn-go"
)

// Generic variables
var version = "v1.27"
var window *giu.MasterWindow
var dpi float32
var mode string
var working bool
var recombine bool

// Popup modals
var modalId int           // A hack to keep modals centered
var showPassgen bool      // Password generator
var showKeyfile bool      // Keyfile manager
var showProgress bool     // Encryption/decryption progress
var showConfirmation bool // Confirm overwriting an existing file

// Input and output files
var onlyFiles []string
var onlyFolders []string
var allFiles []string
var inputLabel = "Drop files and folders into this window."
var inputFile string
var outputFile string

// Password and generator variables
var password string
var cpassword string
var passwordStrength int
var passwordState = giu.InputTextFlagsPassword
var passwordStateLabel = "Show"
var passgenCopy = true
var passgenLength int32 = 32
var passgenUpper = true
var passgenLower = true
var passgenNums = true
var passgenSymbols = true

// Keyfile variables
var keyfile bool
var keyfiles []string
var keyfileOrderMatters bool
var keyfilePrompt = "None selected."

// Comments variables
var comments string
var commentsPrompt = "Comments:"
var commentsDisabled bool

// Advanced options
var paranoid bool
var reedsolo bool
var deleteWhenDone bool
var split bool
var splitSize string
var splitUnits = []string{"KiB", "MiB", "GiB", "TiB", "Total"}
var splitSelected int32 = 1
var compress bool
var keep bool
var kept bool

// Status variables
var startLabel = "Start"
var mainStatus = "Ready."
var mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
var popupStatus string

// Progress variables
var progress float32
var progressInfo string

// Reed-Solomon codecs
var rs1, _ = infectious.NewFEC(1, 3) // 1 data shard, 3 total -> 2 parity shards
var rs5, _ = infectious.NewFEC(5, 15)
var rs16, _ = infectious.NewFEC(16, 48)
var rs24, _ = infectious.NewFEC(24, 72)
var rs32, _ = infectious.NewFEC(32, 96)
var rs64, _ = infectious.NewFEC(64, 192)
var rs128, _ = infectious.NewFEC(128, 136) // Used for full Reed-Solomon on files

// A passthrough and related helpers to get compression progress
var compressDone int64
var compressTotal int64

type compressorProgress struct {
	io.Reader
}

func (p *compressorProgress) Read(data []byte) (int, error) {
	read, err := p.Reader.Read(data)
	compressDone += int64(read)
	progress = float32(compressDone) / float32(compressTotal)
	giu.Update()
	return read, err
}

// The graphical user interface
func draw() {
	giu.SingleWindow().Flags(524351).Layout(
		giu.Custom(func() {
			if showPassgen {
				giu.PopupModal("Generate password:##"+strconv.Itoa(modalId)).Flags(6).Layout(
					giu.Row(
						giu.Label("Length:"),
						giu.SliderInt(&passgenLength, 4, 64).Size(giu.Auto),
					),
					giu.Checkbox("Uppercase", &passgenUpper),
					giu.Checkbox("Lowercase", &passgenLower),
					giu.Checkbox("Numbers", &passgenNums),
					giu.Checkbox("Symbols", &passgenSymbols),
					giu.Checkbox("Copy to clipboard", &passgenCopy),
					giu.Row(
						giu.Button("Cancel").Size(100, 0).OnClick(func() {
							giu.CloseCurrentPopup()
							showPassgen = false
						}),
						giu.Button("Generate").Size(100, 0).OnClick(func() {
							password = genPassword()
							cpassword = password
							passwordStrength = zxcvbn.PasswordStrength(password, nil).Score
							giu.CloseCurrentPopup()
							showPassgen = false
						}),
					),
				).Build()
				giu.OpenPopup("Generate password:##" + strconv.Itoa(modalId))
				giu.Update()
			}

			if showKeyfile {
				giu.PopupModal("Manage keyfiles:##"+strconv.Itoa(modalId)).Flags(70).Layout(
					giu.Label("Drag and drop your keyfiles here."),
					giu.Custom(func() {
						if mode != "decrypt" {
							giu.Checkbox("Require correct order", &keyfileOrderMatters).Build()
							giu.Tooltip("Decryption will require the correct keyfile order.").Build()
						} else if keyfileOrderMatters {
							giu.Label("Correct order is required.").Build()
						}
					}),
					giu.Separator(),
					giu.Custom(func() {
						for _, i := range keyfiles {
							giu.Label(filepath.Base(i)).Build()
						}
					}),
					giu.Row(
						giu.Button("Clear").Size(100, 0).OnClick(func() {
							keyfiles = nil
							if keyfile {
								keyfilePrompt = "Keyfiles required."
							} else {
								keyfilePrompt = "None selected."
							}
							modalId++
						}),
						giu.Tooltip("Remove all keyfiles."),

						giu.Button("Done").Size(100, 0).OnClick(func() {
							giu.CloseCurrentPopup()
							showKeyfile = false
						}),
					),
				).Build()
				giu.OpenPopup("Manage keyfiles:##" + strconv.Itoa(modalId))
				giu.Update()
			}

			if showConfirmation {
				giu.PopupModal("Warning:##"+strconv.Itoa(modalId)).Flags(6).Layout(
					giu.Label("Output already exists. Overwrite?"),
					giu.Row(
						giu.Button("No").Size(100, 0).OnClick(func() {
							giu.CloseCurrentPopup()
							showConfirmation = false
						}),
						giu.Button("Yes").Size(100, 0).OnClick(func() {
							giu.CloseCurrentPopup()
							showConfirmation = false
							modalId++
							showProgress = true
							giu.Update()
							go func() {
								work()
								working = false
								showProgress = false
								giu.Update()
							}()
						}),
					),
				).Build()
				giu.OpenPopup("Warning:##" + strconv.Itoa(modalId))
				giu.Update()
			}

			if showProgress {
				giu.PopupModal(" ##"+strconv.Itoa(modalId)).Flags(6).Layout(
					giu.Row(
						giu.ProgressBar(progress).Size(180, 0).Overlay(progressInfo),
						giu.Button("Cancel").Size(58, 0).OnClick(func() {
							working = false
						}),
					),
					giu.Label(popupStatus),
				).Build()
				giu.OpenPopup(" ##" + strconv.Itoa(modalId))
				giu.Update()
			}
		}),

		giu.Row(
			giu.Label(inputLabel),
			giu.Custom(func() {
				bw, _ := giu.CalcTextSize("Clear")
				p, _ := giu.GetWindowPadding()
				bw += p * 2
				giu.Dummy((bw+p)/-dpi, 0).Build()
				giu.SameLine()
				giu.Style().SetDisabled(len(allFiles) == 0 && len(onlyFiles) == 0).To(
					giu.Button("Clear").Size(bw/dpi, 0).OnClick(resetUI),
					giu.Tooltip("Clear all input files and reset UI state."),
				).Build()
			}),
		),

		giu.Separator(),
		giu.Style().SetDisabled(len(allFiles) == 0 && len(onlyFiles) == 0).To(
			giu.Label("Password:"),
			giu.Row(
				giu.Button(passwordStateLabel).Size(54, 0).OnClick(func() {
					if passwordState == giu.InputTextFlagsPassword {
						passwordState = giu.InputTextFlagsNone
						passwordStateLabel = "Hide"
					} else {
						passwordState = giu.InputTextFlagsPassword
						passwordStateLabel = "Show"
					}
				}),
				giu.Tooltip("Toggle the visibility of password entries."),

				giu.Button("Clear").Size(54, 0).OnClick(func() {
					password = ""
					cpassword = ""
				}),
				giu.Tooltip("Clear the password entries."),

				giu.Button("Copy").Size(54, 0).OnClick(func() {
					clipboard.WriteAll(password)
				}),
				giu.Tooltip("Copy the password into your clipboard."),

				giu.Button("Paste").Size(54, 0).OnClick(func() {
					tmp, _ := clipboard.ReadAll()
					password = tmp
					if mode != "decrypt" {
						cpassword = tmp
					}
					passwordStrength = zxcvbn.PasswordStrength(password, nil).Score
					giu.Update()
				}),
				giu.Tooltip("Paste a password from your clipboard."),

				giu.Style().SetDisabled(mode == "decrypt").To(
					giu.Button("Create").Size(54, 0).OnClick(func() {
						modalId++
						showPassgen = true
					}),
				),
				giu.Tooltip("Generate a cryptographically secure password."),
			),
			giu.Row(
				giu.InputText(&password).Flags(passwordState).Size(302/dpi).OnChange(func() {
					passwordStrength = zxcvbn.PasswordStrength(password, nil).Score
				}),
				giu.Custom(func() {
					c := giu.GetCanvas()
					p := giu.GetCursorScreenPos()

					col := color.RGBA{
						uint8(0xc8 - 31*passwordStrength),
						uint8(0x4c + 31*passwordStrength), 0x4b, 0xff,
					}
					if password == "" || mode == "decrypt" {
						col = color.RGBA{0xff, 0xff, 0xff, 0x00}
					}

					path := p.Add(image.Pt(
						int(math.Round(-20*float64(dpi))),
						int(math.Round(12*float64(dpi))),
					))
					c.PathArcTo(path, 6*dpi, -math.Pi/2, math.Pi*(.4*float32(passwordStrength)-.1), -1)
					c.PathStroke(col, false, 2)
				}),
			),

			giu.Dummy(0, 0),
			giu.Style().SetDisabled(password == "" || mode == "decrypt").To(
				giu.Label("Confirm password:"),
				giu.Row(
					giu.InputText(&cpassword).Flags(passwordState).Size(302/dpi),
					giu.Custom(func() {
						c := giu.GetCanvas()
						p := giu.GetCursorScreenPos()
						col := color.RGBA{0x4c, 0xc8, 0x4b, 0xff}

						if cpassword != password {
							col = color.RGBA{0xc8, 0x4c, 0x4b, 0xff}
						}
						if password == "" || cpassword == "" || mode == "decrypt" {
							col = color.RGBA{0xff, 0xff, 0xff, 0x00}
						}

						path := p.Add(image.Pt(
							int(math.Round(-20*float64(dpi))),
							int(math.Round(12*float64(dpi))),
						))
						c.PathArcTo(path, 6*dpi, 0, 2*math.Pi, -1)
						c.PathStroke(col, false, 2)
					}),
				),
			),

			giu.Dummy(0, 0),
			giu.Style().SetDisabled(mode == "decrypt" && !keyfile).To(
				giu.Row(
					giu.Label("Keyfiles:"),
					giu.Button("Edit").Size(54, 0).OnClick(func() {
						modalId++
						showKeyfile = true
					}),
					giu.Tooltip("Manage your keyfiles."),

					giu.Style().SetDisabled(mode == "decrypt").To(
						giu.Button("Create").Size(54, 0).OnClick(func() {
							f := dialog.File().Title("Choose where to save the keyfile.")
							f.SetStartDir(func() string {
								if len(onlyFiles) > 0 {
									return filepath.Dir(onlyFiles[0])
								}
								return filepath.Dir(onlyFolders[0])
							}())
							f.SetInitFilename("Keyfile")
							file, err := f.Save()
							if file == "" || err != nil {
								return
							}

							fout, _ := os.Create(file)
							data := make([]byte, 1<<20)
							rand.Read(data)
							fout.Write(data)
							fout.Close()
						}),
						giu.Tooltip("Generate a cryptographically secure keyfile."),
					),
					giu.Style().SetDisabled(true).To(
						giu.InputText(&keyfilePrompt).Size(giu.Auto),
					),
				),
			),
		),

		giu.Separator(),
		giu.Style().SetDisabled((mode == "decrypt" && comments == "") ||
			(mode != "decrypt" && ((len(keyfiles) == 0 && password == "") || (password != cpassword)))).To(
			giu.Style().SetDisabled(mode == "decrypt" && comments == "").To(
				giu.Label(commentsPrompt),
				giu.InputText(&comments).Size(giu.Auto).Flags(func() giu.InputTextFlags {
					if commentsDisabled {
						return giu.InputTextFlagsReadOnly
					}
					return giu.InputTextFlagsNone
				}()),
			),
		),
		giu.Style().SetDisabled((len(keyfiles) == 0 && password == "") ||
			(mode == "encrypt" && password != cpassword)).To(
			giu.Label("Advanced:"),
			giu.Custom(func() {
				if mode != "decrypt" {
					giu.Row(
						giu.Checkbox("Paranoid mode", &paranoid),
						giu.Tooltip("Provides the highest level of security attainable."),
						giu.Dummy(-170, 0),
						giu.Style().SetDisabled(!(len(allFiles) > 1 || len(onlyFolders) > 0)).To(
							giu.Checkbox("Compress files", &compress),
							giu.Tooltip("Compress files with Deflate before encrypting."),
						),
					).Build()

					giu.Row(
						giu.Checkbox("Reed-Solomon", &reedsolo),
						giu.Tooltip("Prevent file corruption by erasure coding (slow)."),
						giu.Dummy(-170, 0),
						giu.Checkbox("Delete files", &deleteWhenDone),
						giu.Tooltip("Delete the input files after encryption."),
					).Build()

					giu.Row(
						giu.Checkbox("Split into chunks:", &split),
						giu.Tooltip("Split the output file into smaller chunks."),
						giu.Dummy(-170, 0),
						giu.InputText(&splitSize).Size(86/dpi).Flags(1).OnChange(func() {
							split = splitSize != ""
						}),
						giu.Tooltip("Choose the chunk size."),
						giu.Combo("##splitter", splitUnits[splitSelected], splitUnits, &splitSelected).Size(68),
						giu.Tooltip("Choose the chunk size units."),
					).Build()
				} else {
					giu.Row(
						giu.Checkbox("Force decrypt", &keep),
						giu.Tooltip("Override security measures when decrypting."),
						giu.Dummy(-170, 0),
						giu.Checkbox("Delete volume", &deleteWhenDone),
						giu.Tooltip("Delete the volume after a successful decryption."),
					).Build()
				}
			}),

			giu.Label("Save output as:"),
			giu.Custom(func() {
				w, _ := giu.GetAvailableRegion()
				bw, _ := giu.CalcTextSize("Change")
				p, _ := giu.GetWindowPadding()
				bw += p * 2
				dw := w - bw - p
				giu.Style().SetDisabled(true).To(
					giu.InputText(func() *string {
						tmp := ""
						if outputFile == "" {
							return &tmp
						}
						tmp = filepath.Base(outputFile)
						return &tmp
					}()).Size(dw / dpi / dpi).Flags(16384),
				).Build()

				giu.SameLine()
				giu.Button("Change").Size(bw/dpi, 0).OnClick(func() {
					f := dialog.File().Title("Choose where to save the output. Don't include extensions.")
					f.SetStartDir(func() string {
						if len(onlyFiles) > 0 {
							return filepath.Dir(onlyFiles[0])
						}
						return filepath.Dir(onlyFolders[0])
					}())

					// Prefill the filename
					tmp := strings.TrimSuffix(filepath.Base(outputFile), ".pcv")
					f.SetInitFilename(strings.TrimSuffix(tmp, filepath.Ext(tmp)))
					if mode == "encrypt" && (len(allFiles) > 1 || len(onlyFolders) > 0) {
						f.SetInitFilename("Encrypted")
					}

					file, err := f.Save()
					if file == "" || err != nil {
						return
					}

					// Add the correct extensions
					if mode == "encrypt" {
						if len(allFiles) > 1 || len(onlyFolders) > 0 {
							file += ".zip.pcv"
						} else {
							file += filepath.Ext(inputFile) + ".pcv"
						}
					} else {
						if strings.HasSuffix(inputFile, ".zip.pcv") {
							file += ".zip"
						} else {
							tmp := strings.TrimSuffix(filepath.Base(inputFile), ".pcv")
							file += filepath.Ext(tmp)
						}
					}
					outputFile = file
				}).Build()
				giu.Tooltip("Save the output with a custom name and path.").Build()
			}),

			giu.Dummy(0, 0),
			giu.Separator(),
			giu.Dummy(0, 0),
			giu.Button(startLabel).Size(giu.Auto, 34).OnClick(func() {
				if keyfile && keyfiles == nil {
					mainStatus = "Please select your keyfiles."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					return
				}
				tmp, err := strconv.Atoi(splitSize)
				if split && (splitSize == "" || tmp <= 0 || err != nil) {
					mainStatus = "Invalid split size."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					return
				}
				_, err = os.Stat(outputFile)
				if err == nil {
					modalId++
					showConfirmation = true
					giu.Update()
				} else {
					modalId++
					showProgress = true
					giu.Update()
					go func() {
						work()
						working = false
						showProgress = false
						giu.Update()
					}()
				}
			}),
			giu.Style().SetColor(giu.StyleColorText, mainStatusColor).To(
				giu.Label(mainStatus),
			),
		),

		giu.Custom(func() {
			window.SetSize(int(318*dpi), giu.GetCursorPos().Y+1)
		}),
	)
}

func onDrop(names []string) {
	if showKeyfile {
		keyfiles = append(keyfiles, names...)

		// Remove duplicate keyfiles
		var tmp []string
		for _, i := range keyfiles {
			duplicate := false
			for _, j := range tmp {
				if i == j {
					duplicate = true
				}
			}
			stat, _ := os.Stat(i)
			if !duplicate && !stat.IsDir() {
				tmp = append(tmp, i)
			}
		}
		keyfiles = tmp

		// Update the keyfile status
		if len(keyfiles) == 1 {
			keyfilePrompt = "Using 1 keyfile."
		} else {
			keyfilePrompt = fmt.Sprintf("Using %d keyfiles.", len(keyfiles))
		}

		// Recenter the keyfile modal
		modalId++
		return
	}

	// Clear variables and UI state
	recombine = false
	onlyFiles = nil
	onlyFolders = nil
	allFiles = nil
	files, folders := 0, 0
	size := 0
	resetUI()

	// One item dropped
	if len(names) == 1 {
		stat, _ := os.Stat(names[0])

		// A folder was dropped
		if stat.IsDir() {
			folders++
			mode = "encrypt"
			inputLabel = "1 folder selected."
			startLabel = "Encrypt"
			onlyFolders = append(onlyFolders, names[0])
			inputFile = filepath.Join(filepath.Dir(names[0]), "Encrypted") + ".zip"
			outputFile = inputFile + ".pcv"
		} else { // A file was dropped
			files++
			name := filepath.Base(names[0])

			// Is the file a part of a split volume?
			nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
			endsNum := false
			for _, i := range nums {
				if strings.HasSuffix(names[0], i) {
					endsNum = true
				}
			}
			isSplit := strings.Contains(names[0], ".pcv.") && endsNum

			// Decide if encrypting or decrypting
			if strings.HasSuffix(names[0], ".pcv") || isSplit {
				mode = "decrypt"
				inputLabel = name
				startLabel = "Decrypt"
				commentsPrompt = "Comments (read-only):"
				commentsDisabled = true

				if isSplit {
					ind := strings.Index(names[0], ".pcv")
					names[0] = names[0][:ind+4]
					inputFile = names[0]
					outputFile = names[0][:ind]
					recombine = true
				} else {
					outputFile = names[0][:len(names[0])-4]
				}

				// Open the input file in read-only mode
				var fin *os.File
				if isSplit {
					fin, _ = os.Open(names[0] + ".0")
				} else {
					fin, _ = os.Open(names[0])
				}

				// Use regex to test if the input is a valid Picocrypt volume
				tmp := make([]byte, 30)
				fin.Read(tmp)
				if string(tmp[:5]) == "v1.13" {
					resetUI()
					mainStatus = "Please use v1.13 to decrypt this file."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					fin.Close()
					return
				}
				if valid, _ := regexp.Match(`^v\d\.\d{2}.{10}0?\d+`, tmp); !valid && !isSplit {
					resetUI()
					mainStatus = "This doesn't seem like a Picocrypt volume."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					fin.Close()
					return
				}

				// Use regex to test if the volume is compatible
				fin.Seek(0, 0)
				tmp = make([]byte, 15)
				fin.Read(tmp)
				tmp, _ = rsDecode(rs5, tmp)
				if valid, _ := regexp.Match(`^v1.1[456]$`, tmp); valid {
					resetUI()
					mainStatus = "Please use v1.16 to decrypt this file."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					fin.Close()
					return
				}
				if valid, _ := regexp.Match(`^(v1.1[789])|(v1.2[01])$`, tmp); valid {
					resetUI()
					mainStatus = "Please use v1.21 to decrypt this file."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					fin.Close()
					return
				}

				// Read comments from file and check for corruption
				var err error
				tmp = make([]byte, 15)
				fin.Read(tmp)
				tmp, err = rsDecode(rs5, tmp)
				if err == nil {
					commentsLength, _ := strconv.Atoi(string(tmp))
					tmp = make([]byte, commentsLength*3)
					fin.Read(tmp)
					comments = ""
					for i := 0; i < commentsLength*3; i += 3 {
						t, err := rsDecode(rs1, tmp[i:i+3])
						if err != nil {
							comments = "Comments are corrupted."
							break
						}
						comments += string(t)
					}
				} else {
					comments = "Comments are corrupted."
				}

				// Read flags from file and check for corruption
				flags := make([]byte, 15)
				fin.Read(flags)
				fin.Close()
				flags, err = rsDecode(rs5, flags)
				if err != nil {
					mainStatus = "The volume header is damaged."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					return
				}

				// Update UI and variables according to flags
				if flags[1] == 1 {
					keyfile = true
					keyfilePrompt = "Keyfiles required."
				} else {
					keyfilePrompt = "Not applicable."
				}
				if flags[2] == 1 {
					keyfileOrderMatters = true
				}
			} else { // One file that is not a Picocrypt volume was dropped
				mode = "encrypt"
				inputLabel = name
				startLabel = "Encrypt"
				inputFile = names[0]
				outputFile = names[0] + ".pcv"
			}

			// Add the file
			onlyFiles = append(onlyFiles, names[0])
			inputFile = names[0]
			size += int(stat.Size())
		}
	} else { // There are multiple dropped items
		mode = "encrypt"

		// Go through each dropped item and add to corresponding slices
		for _, name := range names {
			stat, _ := os.Stat(name)
			if stat.IsDir() {
				folders++
				onlyFolders = append(onlyFolders, name)
			} else {
				files++
				onlyFiles = append(onlyFiles, name)
				allFiles = append(allFiles, name)
				size += int(stat.Size())
			}
		}

		// Update UI with the number of files and folders selected
		if folders == 0 {
			inputLabel = fmt.Sprintf("%d files selected.", files)
		} else if files == 0 {
			inputLabel = fmt.Sprintf("%d folders selected.", folders)
		} else {
			if files == 1 && folders > 1 {
				inputLabel = fmt.Sprintf("1 file and %d folders selected.", folders)
			} else if folders == 1 && files > 1 {
				inputLabel = fmt.Sprintf("%d files and 1 folder selected.", files)
			} else if folders == 1 && files == 1 {
				inputLabel = "1 file and 1 folder selected."
			} else {
				inputLabel = fmt.Sprintf("%d files and %d folders selected.", files, folders)
			}
		}
		startLabel = "Encrypt"

		// Set the input and output paths
		inputFile = filepath.Join(filepath.Dir(names[0]), "Encrypted") + ".zip"
		outputFile = inputFile + ".pcv"
	}

	// Recursively add all files in 'onlyFolders' to 'allFiles'
	for _, name := range onlyFolders {
		filepath.Walk(name, func(path string, _ os.FileInfo, _ error) error {
			stat, _ := os.Stat(path)
			if !stat.IsDir() {
				allFiles = append(allFiles, path)
				size += int(stat.Size())
			}
			return nil
		})
	}

	inputLabel = fmt.Sprintf("%s (%s)", inputLabel, sizeify(int64(size)))
}

func work() {
	// Show that Picocrypt is encrypting/decrypting
	popupStatus = "Starting..."
	mainStatus = "Working..."
	mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
	working = true
	padded := false
	giu.Update()

	// Cryptography!
	var salt []byte                           // Argon2 salt, 16 bytes
	var hkdfSalt []byte                       // HKDF-SHA3 salt, 32 bytes
	var serpentSalt []byte                    // Serpent salt, 16 bytes
	var nonce []byte                          // 24-byte XChaCha20 nonce
	var keyHash []byte                        // SHA3-512 hash of encryption key
	var _keyHash []byte                       // Same as 'keyHash', but used for comparison
	var keyfileKey []byte                     // The SHA3-256 hashes of keyfiles
	var keyfileHash []byte = make([]byte, 32) // The SHA3-256 of 'keyfileKey'
	var _keyfileHash []byte                   // Same as 'keyfileHash', but used for comparison
	var dataMac []byte                        // 64-byte authentication tag (BLAKE2b or HMAC-SHA3)

	if mode == "encrypt" {
		if compress {
			popupStatus = "Compressing files..."
		} else {
			popupStatus = "Combining files..."
		}

		// Combine/compress all files into a .zip file
		if len(allFiles) > 1 || len(onlyFolders) > 0 {
			var rootDir string
			if len(onlyFolders) > 0 {
				rootDir = filepath.Dir(onlyFolders[0])
			} else {
				rootDir = filepath.Dir(onlyFiles[0])
			}

			file, err := os.Create(inputFile)
			if err != nil {
				mainStatus = "Access denied by operating system."
				mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
				return
			}

			compressTotal = 0
			for _, path := range allFiles {
				stat, _ := os.Stat(path)
				compressTotal += stat.Size()
			}

			w := zip.NewWriter(file)
			for i, path := range allFiles {
				if !working {
					mainStatus = "Operation cancelled by user."
					mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
					w.Close()
					file.Close()
					os.Remove(inputFile)
					compressDone = 0
					return
				}
				progressInfo = fmt.Sprintf("%d/%d", i+1, len(allFiles))
				giu.Update()

				// Don't add the volume to itself
				if path == inputFile {
					continue
				}

				stat, _ := os.Stat(path)
				header, _ := zip.FileInfoHeader(stat)
				header.Name = strings.TrimPrefix(path, rootDir)
				header.Name = filepath.ToSlash(header.Name)
				header.Name = strings.TrimPrefix(header.Name, "/")

				if compress {
					header.Method = zip.Deflate
				} else {
					header.Method = zip.Store
				}
				writer, _ := w.CreateHeader(header)
				file, err := os.Open(path)
				if err != nil {
					mainStatus = "Access denied by operating system."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					os.Remove(inputFile)
					compressDone = 0
					return
				}

				// Use a passthrough to catch compression progress
				prg := &compressorProgress{Reader: file}
				io.Copy(writer, prg)
				file.Close()
			}
			w.Close()
			file.Close()
			compressDone = 0
		}
	}

	// Recombine a split file if necessary
	if recombine {
		popupStatus = "Recombining file..."
		total := 0
		totalBytes := int64(0)
		done := 0

		// Find out the number of splitted chunks
		for {
			stat, err := os.Stat(fmt.Sprintf("%s.%d", inputFile, total))
			if err != nil {
				break
			}
			total++
			totalBytes += stat.Size()
		}

		// Merge all chunks into one file
		fout, _ := os.Create(inputFile)
		for i := 0; i < total; i++ {
			fin, _ := os.Open(fmt.Sprintf("%s.%d", inputFile, i))
			for {
				data := make([]byte, 1<<20)
				read, err := fin.Read(data)
				if err != nil {
					break
				}
				data = data[:read]
				fout.Write(data)
				done += read
				progressInfo = fmt.Sprintf("%d/%d", i+1, total)
				progress = float32(done) / float32(totalBytes)
				giu.Update()
			}
			fin.Close()
		}
		fout.Close()
		progressInfo = ""
	}

	// Subtract the header size from the total size if decrypting
	stat, _ := os.Stat(inputFile)
	total := stat.Size()
	if mode == "decrypt" {
		total -= 789
	}

	// Open input file in read-only mode
	fin, err := os.Open(inputFile)
	if err != nil {
		mainStatus = "Access denied by operating system."
		mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
		if recombine {
			os.Remove(inputFile)
		}
		if len(allFiles) > 1 || len(onlyFolders) > 0 {
			os.Remove(inputFile)
		}
		return
	}
	var fout *os.File

	// If encrypting, generate values and write to file
	if mode == "encrypt" {
		popupStatus = "Generating values..."
		giu.Update()

		// Create the output file
		var err error
		fout, err = os.Create(outputFile)
		if err != nil {
			mainStatus = "Access denied by operating system."
			mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
			fin.Close()
			if len(allFiles) > 1 || len(onlyFolders) > 0 {
				os.Remove(inputFile)
			}
			return
		}

		// Set up cryptographic values
		salt = make([]byte, 16)
		hkdfSalt = make([]byte, 32)
		serpentSalt = make([]byte, 16)
		nonce = make([]byte, 24)

		// Write the program version to file
		fout.Write(rsEncode(rs5, []byte(version)))

		// Encode and write the comment length to file
		commentsLength := []byte(fmt.Sprintf("%05d", len(comments)))
		commentsLength = rsEncode(rs5, commentsLength)
		fout.Write(commentsLength)

		// Encode the comment and write to file
		for _, i := range []byte(comments) {
			fout.Write(rsEncode(rs1, []byte{i}))
		}

		// Configure flags and write to file
		flags := make([]byte, 5)
		if paranoid { // Paranoid mode selected
			flags[0] = 1
		}
		if len(keyfiles) > 0 { // Keyfiles are being used
			flags[1] = 1
		}
		if keyfileOrderMatters { // Order of keyfiles matter
			flags[2] = 1
		}
		if reedsolo { // Full Reed-Solomon encoding is selected
			flags[3] = 1
		}
		if total%(1<<20) >= 1<<20-128 { // Reed-Solomon internals
			flags[4] = 1
		}
		flags = rsEncode(rs5, flags)
		fout.Write(flags)

		// Fill values with Go's CSPRNG
		rand.Read(salt)
		rand.Read(hkdfSalt)
		rand.Read(serpentSalt)
		rand.Read(nonce)

		// Encode values with Reed-Solomon and write to file
		fout.Write(rsEncode(rs16, salt))
		fout.Write(rsEncode(rs32, hkdfSalt))
		fout.Write(rsEncode(rs16, serpentSalt))
		fout.Write(rsEncode(rs24, nonce))

		// Write placeholders for future use
		fout.Write(make([]byte, 192)) // Hash of encryption key
		fout.Write(make([]byte, 96))  // Hash of keyfile key
		fout.Write(make([]byte, 192)) // BLAKE2b/HMAC-SHA3 tag
	} else { // Decrypting, read values from file and decode
		popupStatus = "Reading values..."
		giu.Update()
		errs := make([]error, 10)

		version := make([]byte, 15)
		fin.Read(version)
		_, errs[0] = rsDecode(rs5, version)

		tmp := make([]byte, 15)
		fin.Read(tmp)
		tmp, errs[1] = rsDecode(rs5, tmp)
		commentsLength, _ := strconv.Atoi(string(tmp))
		fin.Read(make([]byte, commentsLength*3))
		total -= int64(commentsLength) * 3

		flags := make([]byte, 15)
		fin.Read(flags)
		flags, errs[2] = rsDecode(rs5, flags)
		paranoid = flags[0] == 1
		reedsolo = flags[3] == 1
		padded = flags[4] == 1

		salt = make([]byte, 48)
		fin.Read(salt)
		salt, errs[3] = rsDecode(rs16, salt)

		hkdfSalt = make([]byte, 96)
		fin.Read(hkdfSalt)
		hkdfSalt, errs[4] = rsDecode(rs32, hkdfSalt)

		serpentSalt = make([]byte, 48)
		fin.Read(serpentSalt)
		serpentSalt, errs[5] = rsDecode(rs16, serpentSalt)

		nonce = make([]byte, 72)
		fin.Read(nonce)
		nonce, errs[6] = rsDecode(rs24, nonce)

		_keyHash = make([]byte, 192)
		fin.Read(_keyHash)
		_keyHash, errs[7] = rsDecode(rs64, _keyHash)

		_keyfileHash = make([]byte, 96)
		fin.Read(_keyfileHash)
		_keyfileHash, errs[8] = rsDecode(rs32, _keyfileHash)

		dataMac = make([]byte, 192)
		fin.Read(dataMac)
		dataMac, errs[9] = rsDecode(rs64, dataMac)

		// If there was an issue during decoding, the header is corrupted
		for _, err := range errs {
			if err != nil {
				if keep { // If the user chooses to force decrypt
					kept = true
				} else {
					mainStatus = "The volume header is damaged."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					fin.Close()
					if recombine {
						os.Remove(inputFile)
					}
					return
				}
			}
		}
	}

	popupStatus = "Deriving key..."
	progress = 0
	progressInfo = ""
	giu.Update()

	// Derive encryption keys and subkeys
	var key []byte
	if paranoid { // Overkilled parameters for paranoid mode
		key = argon2.IDKey(
			[]byte(password),
			salt,
			8,     // 8 passes
			1<<20, // 1 GiB memory
			8,     // 8 threads
			32,    // 32-byte output key
		)
	} else { // High Argon2 parameters by default
		key = argon2.IDKey(
			[]byte(password),
			salt,
			4,
			1<<20,
			4,
			32,
		)
	}

	// If the 'Cancel' button was pressed, cancel and clean up
	if !working {
		mainStatus = "Operation cancelled by user."
		mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
		fin.Close()
		if mode == "encrypt" {
			fout.Close()
		}
		if recombine {
			os.Remove(inputFile)
		}
		if len(allFiles) > 1 || len(onlyFolders) > 0 {
			os.Remove(inputFile)
		}
		os.Remove(outputFile)
		return
	}

	// If keyfiles are being used
	if len(keyfiles) > 0 || keyfile {
		if keyfileOrderMatters { // If order matters, hash progressively
			var keysum = sha3.New256()
			for _, path := range keyfiles {
				kin, _ := os.Open(path)
				kstat, _ := os.Stat(path)
				kbytes := make([]byte, kstat.Size())
				kin.Read(kbytes)
				kin.Close()
				keysum.Write(kbytes)
			}
			keyfileKey = keysum.Sum(nil)
			keyfileSha3 := sha3.New256()
			keyfileSha3.Write(keyfileKey)
			keyfileHash = keyfileSha3.Sum(nil)
		} else { // If order doesn't matter, hash individually and combine
			var keysum []byte
			for _, path := range keyfiles {
				kin, _ := os.Open(path)
				kstat, _ := os.Stat(path)
				kbytes := make([]byte, kstat.Size())
				kin.Read(kbytes)
				kin.Close()
				ksha3 := sha3.New256()
				ksha3.Write(kbytes)
				keyfileKey := ksha3.Sum(nil)
				if keysum == nil {
					keysum = keyfileKey
				} else {
					for i, j := range keyfileKey {
						keysum[i] ^= j
					}
				}
			}
			keyfileKey = keysum
			keyfileSha3 := sha3.New256()
			keyfileSha3.Write(keysum)
			keyfileHash = keyfileSha3.Sum(nil)
		}
	}

	// Hash the encryption key (used to check if a password is correct when decrypting)
	sha3_512 := sha3.New512()
	sha3_512.Write(key)
	keyHash = sha3_512.Sum(nil)

	// Validate the password and/or keyfiles
	if mode == "decrypt" {
		incorrect := false
		keyCorrect := true
		keyfileCorrect := true
		keyCorrect = subtle.ConstantTimeCompare(keyHash, _keyHash) == 1
		if keyfile {
			keyfileCorrect = subtle.ConstantTimeCompare(keyfileHash, _keyfileHash) == 1
			incorrect = !keyCorrect || !keyfileCorrect
		} else {
			incorrect = !keyCorrect
		}

		// If there's an issue with the password and/or keyfiles
		if incorrect {
			if keep {
				kept = true
			} else {
				if !keyCorrect {
					mainStatus = "The provided password is incorrect."
				} else {
					if keyfileOrderMatters {
						mainStatus = "Incorrect keyfiles or order."
					} else {
						mainStatus = "Incorrect keyfiles."
					}
				}
				mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
				fin.Close()
				if recombine {
					os.Remove(inputFile)
				}
				return
			}
		}

		// Create the output file for decryption
		var err error
		fout, err = os.Create(outputFile)
		if err != nil {
			mainStatus = "Access denied by operating system."
			mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
			fin.Close()
			if recombine {
				os.Remove(inputFile)
			}
			return
		}
	}

	if len(keyfiles) > 0 || keyfile {
		// XOR the encryption key with the keyfile to make the master key
		tmp := key
		key = make([]byte, 32)
		for i := range key {
			key[i] = tmp[i] ^ keyfileKey[i]
		}
	}

	done := 0
	counterDone := 0
	counter := 0
	startTime := time.Now()
	chacha, _ := chacha20.NewUnauthenticatedCipher(key, nonce)

	// Use HKDF-SHA3 to generate a subkey
	var mac hash.Hash
	subkey := make([]byte, 32)
	hkdf := hkdf.New(sha3.New256, key, hkdfSalt, nil)
	hkdf.Read(subkey)
	if paranoid {
		mac = hmac.New(sha3.New512, subkey) // HMAC-SHA3
	} else {
		mac, _ = blake2b.New512(subkey) // Keyed BLAKE2b
	}

	// Generate another subkey and cipher (not used unless paranoid mode is checked)
	serpentKey := make([]byte, 32)
	hkdf.Read(serpentKey)
	s, _ := serpent.NewCipher(serpentKey)
	serpent := cipher.NewCTR(s, serpentSalt)

	for {
		// If the user cancels the process, stop and clean up
		if !working {
			mainStatus = "Operation cancelled by user."
			mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
			fin.Close()
			fout.Close()
			if recombine {
				os.Remove(inputFile)
			}
			if len(allFiles) > 1 || len(onlyFolders) > 0 {
				os.Remove(inputFile)
			}
			os.Remove(outputFile)
			return
		}

		// Read in data from the file
		var src []byte
		if mode == "decrypt" && reedsolo {
			src = make([]byte, 1<<20/128*136)
		} else {
			src = make([]byte, 1<<20)
		}
		size, err := fin.Read(src)
		if err != nil {
			break
		}
		src = src[:size]
		dst := make([]byte, len(src))

		// Do the actual encryption
		if mode == "encrypt" {
			if paranoid {
				serpent.XORKeyStream(dst, src)
				copy(src, dst)
			}

			chacha.XORKeyStream(dst, src)
			mac.Write(dst)

			if reedsolo {
				copy(src, dst)
				dst = nil
				// If a full MiB is available
				if len(src) == 1<<20 {
					// Encode every chunk
					for i := 0; i < 1<<20; i += 128 {
						dst = append(dst, rsEncode(rs128, src[i:i+128])...)
					}
				} else {
					// Encode the full chunks
					chunks := math.Floor(float64(len(src)) / 128)
					for i := 0; float64(i) < chunks; i++ {
						dst = append(dst, rsEncode(rs128, src[i*128:(i+1)*128])...)
					}

					// Pad and encode the final partial chunk
					dst = append(dst, rsEncode(rs128, pad(src[int(chunks*128):]))...)
				}
			}
		} else { // Decryption
			if reedsolo {
				copy(dst, src)
				src = nil
				// If a complete 1 MiB block is available
				if len(dst) == 1<<20/128*136 {
					// Decode every chunk
					for i := 0; i < 1<<20/128*136; i += 136 {
						tmp, err := rsDecode(rs128, dst[i:i+136])
						if err != nil {
							if keep {
								kept = true
							} else {
								fin.Close()
								fout.Close()
								broken()
								mainStatus = "The input file is irrecoverably damaged."
								return
							}
						}
						if i == 1113976 && done+1114112 >= int(total) && padded {
							tmp = unpad(tmp)
						}
						src = append(src, tmp...)
					}
				} else {
					// Decode the full chunks
					chunks := len(dst)/136 - 1
					for i := 0; i < chunks; i++ {
						tmp, err := rsDecode(rs128, dst[i*136:(i+1)*136])
						if err != nil {
							if keep {
								kept = true
							} else {
								fin.Close()
								fout.Close()
								broken()
								mainStatus = "The input file is irrecoverably damaged."
								return
							}
						}
						src = append(src, tmp...)
					}

					// Unpad and decode the final partial chunk
					tmp, err := rsDecode(rs128, dst[int(chunks)*136:])
					if err != nil {
						if keep {
							kept = true
						} else {
							fin.Close()
							fout.Close()
							broken()
							mainStatus = "The input file is irrecoverably damaged."
							return
						}
					}
					src = append(src, unpad(tmp)...)
				}
				dst = make([]byte, len(src))
			}

			mac.Write(src)
			chacha.XORKeyStream(dst, src)

			if paranoid {
				copy(src, dst)
				serpent.XORKeyStream(dst, src)
			}
		}
		fout.Write(dst)

		// Update stats
		if mode == "decrypt" && reedsolo {
			done += 1 << 20 / 128 * 136
		} else {
			done += 1 << 20
		}
		counterDone += 1 << 20
		counter++
		progress = float32(done) / float32(total)
		elapsed := float64(time.Since(startTime)) / (1 << 20) / 1000
		speed := float64(done) / elapsed / (1 << 20)
		eta := int(math.Floor(float64(total-int64(done)) / (speed * (1 << 20))))
		progress = float32(math.Min(float64(progress), 1)) // Cap progress to 100%
		progressInfo = fmt.Sprintf("%.2f%%", progress*100)
		popupStatus = fmt.Sprintf("Working at %.2f MiB/s (ETA: %s)", speed, humanize(eta))
		giu.Update()

		// If more than 256 GiB passed, change the nonce to prevent counter overflow
		blocks := counterDone/64 + 1
		if blocks+(1<<20/64) > 1<<32 {
			nonce = make([]byte, 24)
			hkdf.Read(nonce)
			chacha, _ = chacha20.NewUnauthenticatedCipher(key, nonce)
			counterDone = 0
		}
	}

	if mode == "encrypt" {
		// Seek back to header to write important values
		fout.Seek(int64(309+len(comments)*3), 0)
		fout.Write(rsEncode(rs64, keyHash))
		fout.Write(rsEncode(rs32, keyfileHash))
		fout.Write(rsEncode(rs64, mac.Sum(nil)))
	} else {
		// Validate the authenticity of decrypted data
		if subtle.ConstantTimeCompare(mac.Sum(nil), dataMac) == 0 {
			if keep {
				kept = true
			} else {
				fin.Close()
				fout.Close()
				broken()
				return
			}
		}
	}

	fin.Close()
	fout.Close()

	// Split files into chunks
	if split {
		var splitted []string
		popupStatus = "Splitting file..."
		stat, _ := os.Stat(outputFile)
		size := stat.Size()
		finished := 0
		finishedRaw := 0
		chunkSize, _ := strconv.Atoi(splitSize)

		// User can choose KiB, MiB, GiB, TiB, or custom number of chunks
		if splitSelected == 0 {
			chunkSize *= 1 << 10
		} else if splitSelected == 1 {
			chunkSize *= 1 << 20
		} else if splitSelected == 2 {
			chunkSize *= 1 << 30
		} else if splitSelected == 3 {
			chunkSize *= 1 << 40
		} else {
			chunkSize = int(math.Ceil(float64(size) / float64(chunkSize)))
		}

		// Get the number of required chunks
		chunks := int(math.Ceil(float64(size) / float64(chunkSize)))
		progressInfo = fmt.Sprintf("%d/%d", finished+1, chunks)
		giu.Update()
		fin, _ := os.Open(outputFile)

		for i := 0; i < chunks; i++ { // Make the chunks
			fout, _ := os.Create(fmt.Sprintf("%s.%d", outputFile, i))
			done := 0

			// Copy data into the chunk
			for {
				data := make([]byte, 1<<20)
				for done+len(data) > chunkSize {
					data = make([]byte, int(math.Ceil(float64(len(data))/2)))
				}

				read, err := fin.Read(data)
				if err != nil {
					break
				}
				if !working {
					fin.Close()
					fout.Close()
					if len(allFiles) > 1 || len(onlyFolders) > 0 {
						os.Remove(inputFile)
					}
					mainStatus = "Operation cancelled by user."
					mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}

					// If user cancels, remove the unfinished files
					for _, j := range splitted {
						os.Remove(j)
					}
					os.Remove(fmt.Sprintf("%s.%d", outputFile, i))
					os.Remove(outputFile)
					return
				}
				data = data[:read]
				fout.Write(data)
				done += read
				if done >= chunkSize {
					break
				}

				finishedRaw += read
				progress = float32(finishedRaw) / float32(size)
				giu.Update()
			}
			fout.Close()

			// Update stats
			finished++
			if finished == chunks {
				finished--
			}
			splitted = append(splitted, fmt.Sprintf("%s.%d", outputFile, i))
			progressInfo = fmt.Sprintf("%d/%d", finished+1, chunks)
			giu.Update()
		}

		fin.Close()
		os.Remove(outputFile)
	}

	// Remove the temporary file used to combine a splitted volume
	if recombine {
		os.Remove(inputFile)
	}

	// Delete the temporary zip file used to encrypt files
	if len(allFiles) > 1 || len(onlyFolders) > 0 {
		os.Remove(inputFile)
	}

	// Delete the input file(s) if the user chooses
	if deleteWhenDone {
		progressInfo = ""
		popupStatus = "Deleting files..."
		giu.Update()
		if mode == "decrypt" {
			if recombine {
				total := 0
				for {
					_, err := os.Stat(fmt.Sprintf("%s.%d", inputFile, total))
					if err != nil {
						break
					}
					os.Remove(fmt.Sprintf("%s.%d", inputFile, total))
					total++
				}
			} else {
				os.Remove(inputFile)
			}
		} else {
			for _, i := range onlyFiles {
				os.Remove(i)
			}
			for _, i := range onlyFolders {
				os.RemoveAll(i)
			}
		}
	}

	// All done, reset the UI
	resetUI()

	// If the user chose to keep a corrupted/modified file, let them know
	if kept {
		mainStatus = "The input file was modified. Please be careful."
		mainStatusColor = color.RGBA{0xff, 0xff, 0x00, 0xff}
	} else {
		mainStatus = "Completed."
		mainStatusColor = color.RGBA{0x00, 0xff, 0x00, 0xff}
	}

	// Clear some variables
	working = false
	kept = false
	key = nil
	popupStatus = "Ready."
}

// This function is run if an issue occurs during decryption
func broken() {
	mainStatus = "The input file is damaged or modified."
	mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}

	// Clean up files since decryption failed
	if recombine {
		os.Remove(inputFile)
	}
	os.Remove(outputFile)
}

// Reset the UI to a clean state with nothing selected or checked
func resetUI() {
	mode = ""
	onlyFiles = nil
	onlyFolders = nil
	allFiles = nil
	inputLabel = "Drop files and folders into this window."
	startLabel = "Start"
	password = ""
	cpassword = ""
	keyfiles = nil
	keyfile = false
	keyfileOrderMatters = false
	keyfilePrompt = "None selected."
	comments = ""
	commentsPrompt = "Comments:"
	commentsDisabled = false
	keep = false
	reedsolo = false
	split = false
	splitSize = ""
	splitSelected = 1
	deleteWhenDone = false
	paranoid = false
	compress = false
	inputFile = ""
	outputFile = ""
	progress = 0
	progressInfo = ""
	mainStatus = "Ready."
	mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
	giu.Update()
}

// Reed-Solomon encoder
func rsEncode(rs *infectious.FEC, data []byte) []byte {
	res := make([]byte, rs.Total())
	rs.Encode(data, func(s infectious.Share) {
		res[s.Number] = s.Data[0]
	})
	return res
}

// Reed-Solomon decoder
func rsDecode(rs *infectious.FEC, data []byte) ([]byte, error) {
	tmp := make([]infectious.Share, rs.Total())
	for i := 0; i < rs.Total(); i++ {
		tmp[i].Number = i
		tmp[i].Data = append(tmp[i].Data, data[i])
	}
	res, err := rs.Decode(nil, tmp)

	// Force decode for the "Force decrypt" option
	if err != nil {
		if rs.Total() == 136 {
			return data[:128], err
		}
		return data[:rs.Total()/3], err
	}
	return res, nil
}

// PKCS#7 pad (for use with Reed-Solomon)
func pad(data []byte) []byte {
	padLen := 128 - len(data)%128
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

// PKCS#7 unpad
func unpad(data []byte) []byte {
	length := len(data)
	padLen := int(data[length-1])
	return data[:length-padLen]
}

// Generate a cryptographically secure password
func genPassword() string {
	chars := ""
	if passgenUpper {
		chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if passgenLower {
		chars += "abcdefghijklmnopqrstuvwxyz"
	}
	if passgenNums {
		chars += "1234567890"
	}
	if passgenSymbols {
		chars += "-=!@#$^&()_+?"
	}
	if chars == "" {
		return chars
	}
	tmp := make([]byte, passgenLength)
	for i := 0; i < int(passgenLength); i++ {
		j, _ := rand.Int(rand.Reader, new(big.Int).SetUint64(uint64(len(chars))))
		tmp[i] = chars[j.Int64()]
	}
	if passgenCopy {
		clipboard.WriteAll(string(tmp))
	}
	return string(tmp)
}

// Convert seconds to HH:MM:SS
func humanize(seconds int) string {
	hours := int(math.Floor(float64(seconds) / 3600))
	seconds %= 3600
	minutes := int(math.Floor(float64(seconds) / 60))
	seconds %= 60
	hours = int(math.Max(float64(hours), 0))
	minutes = int(math.Max(float64(minutes), 0))
	seconds = int(math.Max(float64(seconds), 0))
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// Convert bytes to KiB, MiB, etc.
func sizeify(size int64) string {
	if size >= int64(1<<40) {
		return fmt.Sprintf("%.2fT", float64(size)/(1<<40))
	} else if size >= int64(1<<30) {
		return fmt.Sprintf("%.2fG", float64(size)/(1<<30))
	} else if size >= int64(1<<20) {
		return fmt.Sprintf("%.0fM", float64(size)/(1<<20))
	} else {
		return fmt.Sprintf("%.0fK", float64(size)/(1<<10))
	}
}

func main() {
	// Create the main window
	window = giu.NewMasterWindow("Picocrypt", 318, 479, giu.MasterWindowFlagsNotResizable)

	// Start the dialog module
	dialog.Init()

	// Set callbacks
	window.SetDropCallback(onDrop)
	window.SetCloseCallback(func() bool {
		return !working
	})

	// Set universal DPI
	dpi = giu.Context.GetPlatform().GetContentScale()

	// Start the UI
	window.Run(draw)
}
