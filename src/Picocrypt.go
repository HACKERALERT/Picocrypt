package main

/*

Picocrypt v1.19
Copyright (c) Evan Su (https://evansu.cc)
Released under a GNU GPL v3 License
https://github.com/HACKERALERT/Picocrypt

~ In cryptography we trust ~

*/

import (
	_ "embed"

	// Generic
	"archive/zip"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	// Cryptography
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/subtle"

	"github.com/HACKERALERT/serpent"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/sha3"

	// UI
	"github.com/AllenDang/giu"

	// Reed-Solomon
	"github.com/HACKERALERT/infectious"

	// Helpers
	"github.com/HACKERALERT/clipboard"
	"github.com/HACKERALERT/dialog"
	"github.com/HACKERALERT/jibber_jabber"
	"github.com/HACKERALERT/zxcvbn-go"
)

//go:embed icon.png
var icon []byte

//go:embed font.ttf
var font []byte

//go:embed sdelete64.exe
var sdelete64bytes []byte

//go:embed strings.json
var localeBytes []byte

// Localization
type locale struct {
	iso  string
	data []string
}

var locales []locale
var selectedLocale = "en"
var allLocales = []string{
	"en",
}
var languages = []string{
	"English",
}
var languageSelected int32

// Generic variables
var version = "v1.18"
var dpi float32
var tab = 0
var mode string
var working bool
var recombine bool
var fill float32 = -0.0000001
var sdelete64path string

// Three variables store the input files
var onlyFiles []string
var onlyFolders []string
var allFiles []string

// Input file variables
var inputLabel = "Drop files and folders into this window."
var inputFile string

// Password variables
var password string
var cPassword string
var passwordStrength int
var passwordState = giu.InputTextFlagsPassword
var passwordStateLabel = "Show"

// Password generator variables
var showGenpass = false
var genpassCopy = true
var genpassLength int32 = 32
var genpassUpper = true
var genpassLower = true
var genpassNums = true
var genpassSymbols = true

// Keyfile variables
var keyfile bool
var keyfiles []string
var keyfileOrderMatters bool
var keyfilePrompt = "None selected."
var showKeyfile bool

// Metadata variables
var metadata string
var metadataPrompt = "Metadata:"
var metadataDisabled bool

// Advanced options
var shredTemp bool
var fast bool
var paranoid bool
var reedsolo bool
var deleteWhenDone bool
var split bool
var splitSize string
var splitUnits = []string{
	"KiB",
	"MiB",
	"GiB",
}
var splitSelected int32 = 1
var compress bool
var encryptFilename bool
var keep bool
var kept bool

// Output file variables
var outputFile string

// Status variables
var mainStatus = "Ready."
var mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
var popupStatus string

// Progress variables
var progress float32
var progressInfo string
var showProgress bool

// Confirm overwrite variables
var showConfirmation bool

// Reed-Solomon encoders
var rs1, _ = infectious.NewFEC(1, 3) // 1 data shard, 3 total -> 2 parity shards
var rs5, _ = infectious.NewFEC(5, 15)
var rs6, _ = infectious.NewFEC(6, 18)
var rs16, _ = infectious.NewFEC(16, 48)
var rs24, _ = infectious.NewFEC(24, 72)
var rs32, _ = infectious.NewFEC(32, 96)
var rs64, _ = infectious.NewFEC(64, 192)
var rs128, _ = infectious.NewFEC(128, 136)

// File checksum generator variables
var csMd5 string
var csSha1 string
var csSha256 string
var csSha3 string
var csBlake2b string
var csBlake2s string
var csValidate string
var md5Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
var sha1Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
var sha256Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
var sha3Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
var blake2bColor = color.RGBA{0x00, 0x00, 0x00, 0x00}
var blake2sColor = color.RGBA{0x00, 0x00, 0x00, 0x00}
var csProgress float32 = 0
var md5Selected = true
var sha1Selected = true
var sha256Selected = true
var sha3Selected = false
var blake2bSelected = false
var blake2sSelected = false

// Shredder variables
var shredding string = "Ready."
var shredPasses int32 = 4
var stopShredding bool
var shredProgress float32
var shredDone float32
var shredTotal float32
var shredOverlay string

func draw() {
	giu.SingleWindow().Layout(
		giu.Custom(func() {
			pos := giu.GetCursorPos()
			w, _ := giu.CalcTextSize(languages[languageSelected])
			giu.Row(
				giu.Dummy(-w/dpi-34, 0),
				giu.Combo("##language", languages[languageSelected], languages, &languageSelected).OnChange(func() {
					selectedLocale = allLocales[languageSelected]
					shredding = s(shredding)
				}).Size(w/dpi+26),
			).Build()
			giu.SetCursorPos(pos)

			giu.TabBar().TabItems(
				giu.TabItem(s("Encryption")).Layout(
					giu.Custom(func() {
						if giu.IsItemActive() {
							tab = 0
						}
					}),

					giu.Custom(func() {
						if showGenpass {
							giu.PopupModal(s("Generate password:")).
								Flags(giu.WindowFlagsNoMove|giu.WindowFlagsNoResize).Layout(
								giu.Row(
									giu.Label(s("Length:")),
									giu.SliderInt(&genpassLength, 4, 64).Size(fill),
								),
								giu.Checkbox(s("Uppercase"), &genpassUpper),
								giu.Checkbox(s("Lowercase"), &genpassLower),
								giu.Checkbox(s("Numbers"), &genpassNums),
								giu.Checkbox(s("Symbols"), &genpassSymbols),
								giu.Checkbox(s("Copy to clipboard"), &genpassCopy),
								giu.Row(
									giu.Button(s("Cancel")).Size(100, 0).OnClick(func() {
										giu.CloseCurrentPopup()
										showGenpass = false
									}),
									giu.Button(s("Generate")).Size(100, 0).OnClick(func() {
										tmp := genPassword()
										password = tmp
										cPassword = tmp
										passwordStrength = zxcvbn.PasswordStrength(password, nil).Score
										giu.CloseCurrentPopup()
										showGenpass = false
										giu.Update()
									}),
								),
							).Build()
							giu.OpenPopup(s("Generate password:"))
							giu.Update()
						}
					}),

					giu.Custom(func() {
						if showKeyfile {
							giu.PopupModal(s("Manage keyfiles:")).
								Flags(giu.WindowFlagsNoMove|giu.WindowFlagsNoResize).Layout(
								giu.Row(
									giu.Label(s("Drop and drop your keyfiles.")),
								),
								giu.Custom(func() {
									if mode != "decrypt" {
										giu.Checkbox(s("Require correct keyfile order"), &keyfileOrderMatters).Build()
										giu.Tooltip(s("If checked, you will need to drop keyfiles in the correct order.")).Build()
									} else if keyfileOrderMatters {
										giu.Label(s("The correct order of keyfiles is required.")).Build()
									}
								}),

								giu.Custom(func() {
									for _, i := range keyfiles {
										giu.Row(
											giu.Label(filepath.Base(i)),
											giu.Button("Remove").OnClick(func() {
												var tmp []string
												for _, j := range keyfiles {
													if j != i {
														tmp = append(tmp, j)
													}
												}
												keyfiles = tmp
											}),
										).Build()

									}
								}),
								giu.Dummy(0, 200),
								giu.Row(
									giu.Button(s("Clear")).Size(150, 0).OnClick(func() {
										keyfiles = nil
									}),
									giu.Tooltip(s("Remove all keyfiles.")),
									giu.Button(s("Done")).Size(150, 0).OnClick(func() {
										giu.CloseCurrentPopup()
										showKeyfile = false
									}),
								),
							).Build()
							giu.OpenPopup(s("Manage keyfiles:"))
							giu.Update()
						}
					}),

					giu.Custom(func() {
						if showConfirmation {
							giu.PopupModal(s("Warning:")).
								Flags(giu.WindowFlagsNoMove|giu.WindowFlagsNoResize).Layout(
								giu.Label(s("Output already exists. Overwrite?")),
								giu.Row(
									giu.Button(s("No")).Size(100, 0).OnClick(func() {
										giu.CloseCurrentPopup()
										showConfirmation = false
									}),
									giu.Button(s("Yes")).Size(100, 0).OnClick(func() {
										giu.CloseCurrentPopup()
										showConfirmation = false
										showProgress = true
										giu.Update()
										go func() {
											work()
											working = false
											showProgress = false
											debug.FreeOSMemory()
											giu.Update()
										}()
									}),
								),
							).Build()
							giu.OpenPopup(s("Warning:"))
							giu.Update()
						}
					}),

					giu.Custom(func() {
						if showProgress {
							giu.PopupModal(" ").
								Flags(giu.WindowFlagsNoMove|giu.WindowFlagsNoResize|giu.WindowFlagsNoTitleBar).Layout(
								giu.Custom(func() {
									if !working {
										giu.CloseCurrentPopup()
									}
								}),
								giu.Row(
									giu.ProgressBar(progress).Size(280, 0).Overlay(progressInfo),
									giu.Button(s("Cancel")).Size(58, 0).OnClick(func() {
										working = false
									}),
								),
								giu.Label(popupStatus),
							).Build()
							giu.OpenPopup(" ")
							giu.Update()
						}
					}),

					giu.Row(
						giu.Label(inputLabel),
						giu.Custom(func() {
							w, _ := giu.GetAvailableRegion()
							bw, _ := giu.CalcTextSize(s("Clear"))
							p, _ := giu.GetWindowPadding()
							bw += p * 2
							dw := w - bw - p
							giu.Dummy(float32(math.Max(float64(dw/dpi), float64(-bw/dpi-p))), 0).Build()
							giu.SameLine()
							giu.Style().SetDisabled(len(allFiles) == 0 && len(onlyFiles) == 0).To(
								giu.Button(s("Clear")).Size(bw/dpi, 0).OnClick(resetUI),
								giu.Tooltip(s("Clear all input files and reset UI state.")),
							).Build()
						}),
					),

					giu.Separator(),

					giu.Style().SetDisabled(len(allFiles) == 0 && len(onlyFiles) == 0).To(
						giu.Row(
							giu.Label(s("Password:")),
							giu.Dummy(-124, 0),
							giu.Style().SetDisabled(mode == "decrypt" && !keyfile).To(
								giu.Label(s("Keyfiles:")),
							),
						),
						giu.Row(
							giu.Button(s(passwordStateLabel)).Size(54, 0).OnClick(func() {
								if passwordState == giu.InputTextFlagsPassword {
									passwordState = giu.InputTextFlagsNone
									passwordStateLabel = "Hide"
								} else {
									passwordState = giu.InputTextFlagsPassword
									passwordStateLabel = "Show"
								}
							}),

							giu.Button(s("Clear")).Size(54, 0).OnClick(func() {
								password = ""
								cPassword = ""
							}),

							giu.Button(s("Copy")).Size(54, 0).OnClick(func() {
								clipboard.WriteAll(password)
							}),

							giu.Button(s("Paste")).Size(54, 0).OnClick(func() {
								tmp, _ := clipboard.ReadAll()
								password = tmp
								if mode != "decrypt" {
									cPassword = tmp
								}
								passwordStrength = zxcvbn.PasswordStrength(password, nil).Score
								giu.Update()
							}),

							giu.Style().SetDisabled(mode == "decrypt").To(
								giu.Button(s("Create")).Size(54, 0).OnClick(func() {
									showGenpass = true
								}),
							),

							giu.Style().SetDisabled(mode == "decrypt" && !keyfile).To(
								giu.Row(
									giu.Button(s("Edit")).Size(54, 0).OnClick(func() {
										showKeyfile = true
									}),
									giu.Style().SetDisabled(mode == "decrypt").To(
										giu.Button(s("Create")).Size(54, 0).OnClick(func() {
											file, _ := dialog.File().Title(s("Save keyfile as:")).Save()
											if file == "" {
												return
											}
											fout, _ := os.Create(file)
											data := make([]byte, 1048576)
											rand.Read(data)
											fout.Write(data)
											fout.Close()
										}),
									),
								),
							),
						),
						giu.Row(
							giu.InputText(&password).Flags(passwordState).Size(302/dpi).OnChange(func() {
								passwordStrength = zxcvbn.PasswordStrength(password, nil).Score
							}),
							giu.Custom(func() {
								c := giu.GetCanvas()
								p := giu.GetCursorScreenPos()

								var col color.RGBA
								switch passwordStrength {
								case 0:
									col = color.RGBA{0xc8, 0x4c, 0x4b, 0xff}
								case 1:
									col = color.RGBA{0xa9, 0x6b, 0x4b, 0xff}
								case 2:
									col = color.RGBA{0x8a, 0x8a, 0x4b, 0xff}
								case 3:
									col = color.RGBA{0x6b, 0xa9, 0x4b, 0xff}
								case 4:
									col = color.RGBA{0x4c, 0xc8, 0x4b, 0xff}
								}
								if password == "" || mode == "decrypt" {
									col = color.RGBA{0xff, 0xff, 0xff, 0x00}
								}

								path := p.Add(image.Pt(
									int(math.Round(float64(-20*dpi))),
									int(math.Round(float64(12*dpi))),
								))
								c.PathArcTo(path, 6*dpi, -math.Pi/2, float32(passwordStrength+1)/5*2*math.Pi-math.Pi/2, -1)
								c.PathStroke(col, false, 2)
							}),
							giu.Style().SetDisabled(true).To(
								giu.InputText(&keyfilePrompt).Size(fill),
							),
						),
					),

					giu.Style().SetDisabled(password == "").To(
						giu.Row(
							giu.Style().SetDisabled(mode == "decrypt").To(
								giu.Label(s("Confirm password:")),
							),
							giu.Dummy(-124, 0),
							giu.Style().SetDisabled(true).To(
								giu.Label(s("Custom Argon2:")),
							),
						),
					),
					giu.Style().SetDisabled(password == "").To(
						giu.Row(
							giu.Style().SetDisabled(mode == "decrypt").To(
								giu.Row(
									giu.InputText(&cPassword).Flags(passwordState).Size(302/dpi),
									giu.Custom(func() {
										c := giu.GetCanvas()
										p := giu.GetCursorScreenPos()
										col := color.RGBA{0x4c, 0xc8, 0x4b, 0xff}

										if cPassword != password {
											col = color.RGBA{0xc8, 0x4c, 0x4b, 0xff}
										}
										if password == "" || cPassword == "" || mode == "decrypt" {
											col = color.RGBA{0xff, 0xff, 0xff, 0x00}
										}

										path := p.Add(image.Pt(
											int(math.Round(float64(-20*dpi))),
											int(math.Round(float64(12*dpi))),
										))
										c.PathArcTo(path, 6*dpi, 0, 2*math.Pi, -1)
										c.PathStroke(col, false, 2)
									}),
								),
							),
							giu.Style().SetDisabled(true).To(
								giu.Button(s("W.I.P")).Size(fill, 0),
							),
						),
					),

					giu.Dummy(0, 3),
					giu.Separator(),
					giu.Dummy(0, 0),

					giu.Style().SetDisabled(password == "" || (password != cPassword && mode == "encrypt")).To(
						giu.Label(s(metadataPrompt)),
						giu.Style().SetDisabled(metadataDisabled).To(
							giu.InputText(&metadata).Size(fill),
						),

						giu.Label(s("Advanced:")),
						giu.Custom(func() {
							if mode != "decrypt" {
								giu.Row(
									giu.Checkbox(s("Shred temporary files"), &shredTemp),
									giu.Dummy(-221, 0),
									giu.Checkbox(s("Encode with Reed-Solomon"), &reedsolo),
								).Build()
								giu.Row(
									giu.Checkbox(s("Use fast mode"), &fast),
									giu.Dummy(-221, 0),
									giu.Checkbox(s("Delete files when complete"), &deleteWhenDone),
								).Build()
								giu.Row(
									giu.Checkbox(s("Use paranoid mode"), &paranoid),
									giu.Dummy(-221, 0),
									giu.Style().SetDisabled(!(len(allFiles) > 1 || len(onlyFolders) > 0)).To(
										giu.Checkbox(s("Compress files"), &compress),
									),
								).Build()
								giu.Row(
									giu.Style().SetDisabled(true).To(
										giu.Checkbox(s("Encrypt filename (W.I.P)"), &encryptFilename),
									),
									giu.Dummy(-221, 0),
									giu.Checkbox(s("Split every"), &split),
									giu.InputText(&splitSize).Size(55/dpi).Flags(giu.InputTextFlagsCharsHexadecimal).OnChange(func() {
										split = splitSize != ""
									}),
									giu.Combo("##splitter", splitUnits[splitSelected], splitUnits, &splitSelected).Size(52),
								).Build()
							} else {
								giu.Checkbox(s("Keep decrypted output even if it's corrupted or modified"), &keep).Build()
								giu.Checkbox(s("Delete the encrypted files after a successful decryption"), &deleteWhenDone).Build()
								giu.Dummy(0, 52).Build()
							}
						}),

						giu.Label(s("Save output as:")),
						giu.Custom(func() {
							w, _ := giu.GetAvailableRegion()
							bw, _ := giu.CalcTextSize(s("Change"))
							p, _ := giu.GetWindowPadding()
							bw += p * 2
							dw := w - bw - p
							giu.Style().SetDisabled(true).To(
								giu.InputText(&outputFile).Size(dw / dpi / dpi).Flags(giu.InputTextFlagsReadOnly),
							).Build()
							giu.SameLine()
							giu.Button(s("Change")).Size(bw/dpi, 0).OnClick(func() {
								file, _ := dialog.File().Title(s("Save as:")).Save()
								if file == "" {
									return
								}

								if mode == "encrypt" {
									if len(allFiles) > 1 || len(onlyFolders) > 0 {
										file = strings.TrimSuffix(file, ".zip.pcv")
										file = strings.TrimSuffix(file, ".pcv")
										if !strings.HasSuffix(file, ".zip.pcv") {
											file += ".zip.pcv"
										}
									} else {
										file = strings.TrimSuffix(file, ".pcv")
										ind := strings.Index(inputFile, ".")
										file += inputFile[ind:]
										if !strings.HasSuffix(file, ".pcv") {
											file += ".pcv"
										}
									}
								} else {
									ind := strings.Index(file, ".")
									if ind != -1 {
										file = file[:ind]
									}
									if strings.HasSuffix(inputFile, ".zip.pcv") {
										file += ".zip"
									} else {
										tmp := strings.TrimSuffix(filepath.Base(inputFile), ".pcv")
										tmp = tmp[strings.Index(tmp, "."):]
										file += tmp
									}
								}

								outputFile = file
							}).Build()
							giu.Tooltip(s("Save the output with a custom path and name.")).Build()
						}),

						giu.Dummy(0, 2),
						giu.Separator(),
						giu.Dummy(0, 3),

						giu.Button(s("Start")).Size(fill, 34).OnClick(func() {
							if keyfile && keyfiles == nil {
								mainStatus = "Please select your keyfiles."
								mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
								return
							}
							_, err := os.Stat(outputFile)
							if err == nil {
								showConfirmation = true
								giu.Update()
							} else {
								showProgress = true
								giu.Update()
								go func() {
									work()
									working = false
									showProgress = false
									debug.FreeOSMemory()
									giu.Update()
								}()
							}
						}),
						giu.Style().SetColor(giu.StyleColorText, mainStatusColor).To(
							giu.Label(s(mainStatus)),
						),
					),
				),
				giu.TabItem(s("Checksum")).Layout(
					giu.Custom(func() {
						if giu.IsItemActive() {
							tab = 1
						}
					}),
					giu.Label(s("Toggle the hashes you would like to generate and drop a file here.")),

					// MD5
					giu.Custom(func() {
						giu.Checkbox("MD5:", &md5Selected).OnChange(func() {
							csMd5 = ""
						}).Build()
						giu.SameLine()
						w, _ := giu.GetAvailableRegion()
						bw, _ := giu.CalcTextSize(s("Copy"))
						padding, _ := giu.GetWindowPadding()
						bw += 2 * padding
						size := w - bw - padding
						giu.Dummy(size/dpi, 0).Build()
						giu.SameLine()
						giu.Button(s("Copy")+"##md5").Size(bw/dpi, 0).OnClick(func() {
							clipboard.WriteAll(csMd5)
						}).Build()
					}),
					giu.Style().SetColor(giu.StyleColorBorder, md5Color).To(
						giu.Style().SetDisabled(true).To(
							giu.InputText(&csMd5).Size(fill).Flags(giu.InputTextFlagsReadOnly),
						),
					),

					// SHA1
					giu.Custom(func() {
						giu.Checkbox("SHA1:", &sha1Selected).OnChange(func() {
							csSha1 = ""
						}).Build()
						giu.SameLine()
						w, _ := giu.GetAvailableRegion()
						bw, _ := giu.CalcTextSize(s("Copy"))
						padding, _ := giu.GetWindowPadding()
						bw += 2 * padding
						size := w - bw - padding
						giu.Dummy(size/dpi, 0).Build()
						giu.SameLine()
						giu.Button(s("Copy")+"##sha1").Size(bw/dpi, 0).OnClick(func() {
							clipboard.WriteAll(csSha1)
						}).Build()
					}),
					giu.Style().SetColor(giu.StyleColorBorder, sha1Color).To(
						giu.Style().SetDisabled(true).To(
							giu.InputText(&csSha1).Size(fill).Flags(giu.InputTextFlagsReadOnly),
						),
					),

					// SHA256
					giu.Custom(func() {
						giu.Checkbox("SHA256:", &sha256Selected).OnChange(func() {
							csSha256 = ""
						}).Build()
						giu.SameLine()
						w, _ := giu.GetAvailableRegion()
						bw, _ := giu.CalcTextSize(s("Copy"))
						padding, _ := giu.GetWindowPadding()
						bw += 2 * padding
						size := w - bw - padding
						giu.Dummy(size/dpi, 0).Build()
						giu.SameLine()
						giu.Button(s("Copy")+"##sha256").Size(bw/dpi, 0).OnClick(func() {
							clipboard.WriteAll(csSha256)
						}).Build()
					}),
					giu.Style().SetColor(giu.StyleColorBorder, sha256Color).To(
						giu.Style().SetDisabled(true).To(
							giu.InputText(&csSha256).Size(fill).Flags(giu.InputTextFlagsReadOnly),
						),
					),

					// SHA3-256
					giu.Custom(func() {
						giu.Checkbox("SHA3:", &sha3Selected).OnChange(func() {
							csSha3 = ""
						}).Build()
						giu.SameLine()
						w, _ := giu.GetAvailableRegion()
						bw, _ := giu.CalcTextSize(s("Copy"))
						padding, _ := giu.GetWindowPadding()
						bw += 2 * padding
						size := w - bw - padding
						giu.Dummy(size/dpi, 0).Build()
						giu.SameLine()
						giu.Button(s("Copy")+"##sha3").Size(bw/dpi, 0).OnClick(func() {
							clipboard.WriteAll(csSha3)
						}).Build()
					}),
					giu.Style().SetColor(giu.StyleColorBorder, sha3Color).To(
						giu.Style().SetDisabled(true).To(
							giu.InputText(&csSha3).Size(fill).Flags(giu.InputTextFlagsReadOnly),
						),
					),

					// BLAKE2b
					giu.Custom(func() {
						giu.Checkbox("BLAKE2b:", &blake2bSelected).OnChange(func() {
							csBlake2b = ""
						}).Build()
						giu.SameLine()
						w, _ := giu.GetAvailableRegion()
						bw, _ := giu.CalcTextSize(s("Copy"))
						padding, _ := giu.GetWindowPadding()
						bw += 2 * padding
						size := w - bw - padding
						giu.Dummy(size/dpi, 0).Build()
						giu.SameLine()
						giu.Button(s("Copy")+"##blake2b").Size(bw/dpi, 0).OnClick(func() {
							clipboard.WriteAll(csBlake2b)
						}).Build()
					}),
					giu.Style().SetColor(giu.StyleColorBorder, blake2bColor).To(
						giu.Style().SetDisabled(true).To(
							giu.InputText(&csBlake2b).Size(fill).Flags(giu.InputTextFlagsReadOnly),
						),
					),

					// BLAKE2s
					giu.Custom(func() {
						giu.Checkbox("BLAKE2s:", &blake2sSelected).OnChange(func() {
							csBlake2s = ""
						}).Build()
						giu.SameLine()
						w, _ := giu.GetAvailableRegion()
						bw, _ := giu.CalcTextSize(s("Copy"))
						padding, _ := giu.GetWindowPadding()
						bw += 2 * padding
						size := w - bw - padding
						giu.Dummy(size/dpi, 0).Build()
						giu.SameLine()
						giu.Button(s("Copy")+"##blake2s").Size(bw/dpi, 0).OnClick(func() {
							clipboard.WriteAll(csBlake2s)
						}).Build()
					}),
					giu.Style().SetColor(giu.StyleColorBorder, blake2sColor).To(
						giu.Style().SetDisabled(true).To(
							giu.InputText(&csBlake2s).Size(fill).Flags(giu.InputTextFlagsReadOnly),
						),
					),

					giu.Dummy(0, 23),
					// Input entry for validating a checksum
					giu.Row(
						giu.Label(s("Validate a checksum:")),
						giu.Custom(func() {
							bw, _ := giu.CalcTextSize(s("Paste"))
							padding, _ := giu.GetWindowPadding()
							bw += 2 * padding
							giu.Button(s("Paste")).Size(bw/dpi, 0).OnClick(func() {
								tmp, _ := clipboard.ReadAll()
								csValidate = tmp
								md5Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
								sha1Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
								sha256Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
								sha3Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
								blake2bColor = color.RGBA{0x00, 0x00, 0x00, 0x00}
								blake2sColor = color.RGBA{0x00, 0x00, 0x00, 0x00}
								if csValidate == "" {
									return
								}
								csValidate = strings.ToLower(csValidate)
								if csValidate == csMd5 {
									md5Color = color.RGBA{0x00, 0xff, 0x00, 0xff}
								} else if csValidate == csSha1 {
									sha1Color = color.RGBA{0x00, 0xff, 0x00, 0xff}
								} else if csValidate == csSha256 {
									sha256Color = color.RGBA{0x00, 0xff, 0x00, 0xff}
								} else if csValidate == csSha3 {
									sha3Color = color.RGBA{0x00, 0xff, 0x00, 0xff}
								} else if csValidate == csBlake2b {
									blake2bColor = color.RGBA{0x00, 0xff, 0x00, 0xff}
								} else if csValidate == csBlake2s {
									blake2sColor = color.RGBA{0x00, 0xff, 0x00, 0xff}
								}
								giu.Update()
							}).Build()
						}),
						giu.Custom(func() {
							bw, _ := giu.CalcTextSize(s("Paste"))
							padding, _ := giu.GetWindowPadding()
							bw += 2 * padding
							giu.Button(s("Clear")).Size(bw/dpi, 0).OnClick(func() {
								csValidate = ""
								md5Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
								sha1Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
								sha256Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
								sha3Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
								blake2bColor = color.RGBA{0x00, 0x00, 0x00, 0x00}
								blake2sColor = color.RGBA{0x00, 0x00, 0x00, 0x00}
							}).Build()
						}),
					),
					giu.Style().SetDisabled(true).To(
						giu.InputText(&csValidate).Size(fill),
					),

					// Progress bar
					giu.Label(s("Progress:")),
					giu.ProgressBar(csProgress).Size(fill, 0),
				),
				giu.TabItem(s("Shredder")).Layout(
					giu.Custom(func() {
						if giu.IsItemActive() {
							tab = 2
						}
					}),

					giu.Label(s("Drop files and folders here to shred them.")),
					giu.Custom(func() {
						if runtime.GOOS == "darwin" {
							giu.Label(s("Number of passes: Not supported on macOS")).Build()
						} else {
							giu.Row(
								giu.Label(s("Number of passes:")),
								giu.SliderInt(&shredPasses, 1, 32).Size(fill),
							).Build()
						}
					}),
					giu.Dummy(0, -50),
					giu.Custom(func() {
						w, _ := giu.GetAvailableRegion()
						bw, _ := giu.CalcTextSize(s("Cancel"))
						padding, _ := giu.GetWindowPadding()
						bw += 2 * padding
						size := w - bw - padding
						giu.Row(
							giu.ProgressBar(shredProgress).Overlay(shredOverlay).Size(size/dpi, 0),
							giu.Button(s("Cancel")).Size(bw/dpi, 0).OnClick(func() {
								stopShredding = true
								shredding = s("Ready.")
								shredProgress = 0
								shredOverlay = ""
							}),
						).Build()
					}),
					giu.Custom(func() {
						if len(shredding) > 60 {
							shredding = "....." + shredding[len(shredding)-50:]
						}
						giu.Label(shredding).Wrapped(true).Build()
					}),
				),
				giu.TabItem(s("About")).Layout(
					giu.Custom(func() {
						if giu.IsItemActive() {
							tab = 3
						}
					}),
					giu.Label(fmt.Sprintf(s("Picocrypt %s, created by Evan Su (https://evansu.cc/)."), version)),
					giu.Label(s("Released under a GNU GPL v3 License.")),
					giu.Label(s("A warm thank you to all the people listed below.")),
					giu.Label(s("Patrons:")),
					giu.Label("    - Frederick Doe"),
					giu.Label(s("Donators:")),
					giu.Label("    - jp26"),
					giu.Label("    - W.Graham"),
					giu.Label("    - N. Chin"),
					giu.Label("    - Manjot"),
					giu.Label("    - Phil P."),
					giu.Label("    - E. Zahard"),
					giu.Label(s("Translators:")),
					giu.Label("umitseyhan75, digitalblossom, zeeaall, francirc, kurpau"),
					giu.Label(s("Other:")),
					giu.Label("Fuderal, u/greenreddits, u/Tall_Escape, u/NSABackdoors"),
				),
			).Build()
		}),
	)
}

func onDrop(names []string) {
	if tab == 1 {
		go generateChecksums(names[0])
		return
	}
	if tab == 2 {
		go shred(names, true)
		return
	}

	if showKeyfile {
		keyfiles = append(keyfiles, names...)
		tmp := []string{}
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
		if len(keyfiles) == 1 {
			keyfilePrompt = s("Using 1 keyfile.")
		}
		keyfilePrompt = fmt.Sprintf(s("Using %d keyfiles."), len(keyfiles))
		return
	}

	// Clear variables
	recombine = false
	onlyFiles = nil
	onlyFolders = nil
	allFiles = nil
	files, folders := 0, 0
	resetUI()

	if len(names) == 1 {
		stat, _ := os.Stat(names[0])
		if stat.IsDir() {
			// Update variables
			mode = "encrypt"
			folders++
			inputLabel = s("1 folder selected.")

			// Add the folder
			onlyFolders = append(onlyFolders, names[0])

			// Set the input and output paths
			inputFile = filepath.Join(filepath.Dir(names[0]), s("Encrypted")) + ".zip"
			outputFile = filepath.Join(filepath.Dir(names[0]), s("Encrypted")) + ".zip.pcv"
		} else {
			files++
			name := filepath.Base(names[0])
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
				//var err error
				mode = "decrypt"
				inputLabel = name + s(" (will decrypt)")
				metadataPrompt = s("Metadata (read-only):")
				metadataDisabled = true

				if isSplit {
					inputLabel = name + s(" (will recombine and decrypt)")
					ind := strings.Index(names[0], ".pcv")
					names[0] = names[0][:ind+4]
					inputFile = names[0]
					outputFile = names[0][:ind]
					recombine = true
				} else {
					outputFile = names[0][:len(names[0])-4]
				}

				// Open input file in read-only mode
				var fin *os.File
				if isSplit {
					fin, _ = os.Open(names[0] + ".0")
				} else {
					fin, _ = os.Open(names[0])
				}

				// Use regex to test if input is a valid Picocrypt volume
				tmp := make([]byte, 30)
				fin.Read(tmp)
				if string(tmp[:5]) == "v1.13" {
					resetUI()
					mainStatus = "Please use Picocrypt v1.13 to decrypt this file."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					fin.Close()
					return
				}
				if valid, _ := regexp.Match(`^v\d\.\d{2}.{10}0?\d+`, tmp); !valid && !isSplit {
					resetUI()
					mainStatus = "This doesn't seem to be a Picocrypt volume."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					fin.Close()
					return
				}
				fin.Seek(0, 0)

				// Read metadata and insert into box
				var err error
				tmp = make([]byte, 15)
				fin.Read(tmp)
				tmp, _ = rsDecode(rs5, tmp)
				if string(tmp) == "v1.14" || string(tmp) == "v1.15" || string(tmp) == "v1.16" {
					resetUI()
					mainStatus = "Please use Picocrypt v1.16 to decrypt this file."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					fin.Close()
					return
				}
				tmp = make([]byte, 15)
				fin.Read(tmp)
				tmp, err = rsDecode(rs5, tmp)

				if err == nil {
					metadataLength, _ := strconv.Atoi(string(tmp))
					tmp = make([]byte, metadataLength*3)
					fin.Read(tmp)
					metadata = ""

					for i := 0; i < metadataLength*3; i += 3 {
						t, err := rsDecode(rs1, tmp[i:i+3])
						if err != nil {
							metadata = s("Metadata is corrupted.")
							break
						}
						metadata += string(t)
					}
				} else {
					metadata = s("Metadata is corrupted.")
				}

				flags := make([]byte, 18)
				fin.Read(flags)
				fin.Close()
				flags, err = rsDecode(rs6, flags)
				if err != nil {
					mainStatus = "Input file is corrupt and cannot be decrypted."
					mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
					return
				}

				if flags[2] == 1 {
					keyfile = true
					keyfilePrompt = s("Keyfiles required.")
				} else {
					keyfilePrompt = s("Not applicable.")
				}
				if flags[5] == 1 {
					keyfileOrderMatters = true
				}
			} else {
				mode = "encrypt"
				inputLabel = name + s(" (will encrypt)")
				inputFile = names[0]
				outputFile = names[0] + ".pcv"
			}

			// Add the file
			onlyFiles = append(onlyFiles, names[0])
			inputFile = names[0]
		}
	} else {
		mode = "encrypt"

		// There are multiple dropped items, check each one
		for _, name := range names {
			stat, _ := os.Stat(name)

			// Check if item is a file or a directory
			if stat.IsDir() {
				folders++
				onlyFolders = append(onlyFolders, name)
			} else {
				files++
				onlyFiles = append(onlyFiles, name)
				allFiles = append(allFiles, name)
			}
		}

		if folders == 0 {
			inputLabel = fmt.Sprintf(s("%d files selected."), files)
		} else if files == 0 {
			inputLabel = fmt.Sprintf(s("%d folders selected."), files)
		} else {
			if files == 1 && folders > 1 {
				inputLabel = fmt.Sprintf(s("1 file and %d folders selected."), folders)
			} else if folders == 1 && files > 1 {
				inputLabel = fmt.Sprintf(s("%d files and 1 folder selected."), files)
			} else if folders == 1 && files == 1 {
				inputLabel = s("1 file and 1 folder selected.")
			} else {
				inputLabel = fmt.Sprintf(s("%d files and %d folders selected."), files, folders)
			}
		}

		// Set the input and output paths
		inputFile = filepath.Join(filepath.Dir(names[0]), s("Encrypted")) + ".zip"
		outputFile = filepath.Join(filepath.Dir(names[0]), s("Encrypted")) + ".zip.pcv"
	}
	// Recursively add all files to 'allFiles'
	if folders > 0 {
		for _, name := range onlyFolders {
			filepath.Walk(name, func(path string, _ os.FileInfo, _ error) error {
				stat, _ := os.Stat(path)
				if !stat.IsDir() {
					allFiles = append(allFiles, path)
				}
				return nil
			})
		}
	}
}

func work() {
	popupStatus = s("Starting...")
	mainStatus = "Working..."
	mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
	working = true
	padded := false

	var salt []byte
	var hkdfSalt []byte
	var serpentSalt []byte
	var nonce []byte
	var keyHash []byte
	var _keyHash []byte
	var keyfileKey []byte
	var keyfileHash []byte = make([]byte, 32)
	var _keyfileHash []byte
	var dataMac []byte

	if mode == "encrypt" {
		if compress {
			popupStatus = s("Compressing files...")
		} else {
			popupStatus = s("Combining files...")
		}

		// "Tar" files together (a .zip file with no compression)
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

			w := zip.NewWriter(file)
			for i, path := range allFiles {
				if !working {
					w.Close()
					file.Close()
					os.Remove(inputFile)
					mainStatus = "Operation cancelled by user."
					mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
					return
				}
				progressInfo = fmt.Sprintf("%d/%d", i, len(allFiles))
				progress = float32(i) / float32(len(allFiles))
				giu.Update()
				if path == inputFile {
					continue
				}

				stat, _ := os.Stat(path)
				header, _ := zip.FileInfoHeader(stat)
				header.Name = strings.TrimPrefix(path, rootDir)

				// Windows requires forward slashes in a .zip file
				if runtime.GOOS == "windows" {
					header.Name = strings.ReplaceAll(header.Name, "\\", "/")
					header.Name = strings.TrimPrefix(header.Name, "/")
				}

				if compress {
					header.Method = zip.Deflate
				} else {
					header.Method = zip.Store
				}
				writer, _ := w.CreateHeader(header)
				file, _ := os.Open(path)
				io.Copy(writer, file)
				file.Close()
			}
			w.Flush()
			w.Close()
			file.Close()
		}
	}

	if recombine {
		popupStatus = s("Recombining file...")
		total := 0

		for {
			_, err := os.Stat(fmt.Sprintf("%s.%d", inputFile, total))
			if err != nil {
				break
			}
			total++
		}

		fout, _ := os.Create(inputFile)
		for i := 0; i < total; i++ {
			fin, _ := os.Open(fmt.Sprintf("%s.%d", inputFile, i))
			for {
				data := make([]byte, 1048576)
				read, err := fin.Read(data)
				if err != nil {
					break
				}
				data = data[:read]
				fout.Write(data)
			}
			fin.Close()
			progressInfo = fmt.Sprintf("%d/%d", i, total)
			progress = float32(i) / float32(total)
			giu.Update()
		}
		fout.Close()
		progressInfo = ""
	}

	stat, _ := os.Stat(inputFile)
	total := stat.Size()
	if mode == "decrypt" {
		total -= 789
	}

	// XChaCha20's max message size is 256 GiB
	if total > 256*1073741824 {
		mainStatus = "Total size is larger than 256 GiB, XChaCha20's limit."
		mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
		return
	}

	// Open input file in read-only mode
	fin, err := os.Open(inputFile)
	if err != nil {
		mainStatus = "Access denied by operating system."
		mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
		return
	}

	var fout *os.File

	// If encrypting, generate values; if decrypting, read values from file
	if mode == "encrypt" {
		popupStatus = s("Generating values...")
		giu.Update()

		var err error
		fout, err = os.Create(outputFile)
		if err != nil {
			mainStatus = "Access denied by operating system."
			mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
			return
		}

		// Generate random cryptography values
		salt = make([]byte, 16)
		hkdfSalt = make([]byte, 32)
		serpentSalt = make([]byte, 16)
		nonce = make([]byte, 24)

		// Write version to file
		fout.Write(rsEncode(rs5, []byte(version)))

		// Encode the length of the metadata with Reed-Solomon
		metadataLength := []byte(fmt.Sprintf("%05d", len(metadata)))
		metadataLength = rsEncode(rs5, metadataLength)

		// Write the length of the metadata to file
		fout.Write(metadataLength)

		// Reed-Solomon-encode the metadata and write to file
		for _, i := range []byte(metadata) {
			fout.Write(rsEncode(rs1, []byte{i}))
		}

		flags := make([]byte, 6)
		if fast {
			flags[0] = 1
		}
		if paranoid {
			flags[1] = 1
		}
		if len(keyfiles) > 0 {
			flags[2] = 1
		}
		if reedsolo {
			flags[3] = 1
		}
		if total%1048576 >= 1048448 {
			flags[4] = 1
		}
		if keyfileOrderMatters {
			flags[5] = 1
		}
		flags = rsEncode(rs6, flags)
		fout.Write(flags)

		// Fill salts and nonce with Go's CSPRNG
		rand.Read(salt)
		rand.Read(hkdfSalt)
		rand.Read(serpentSalt)
		rand.Read(nonce)

		// Encode salt with Reed-Solomon and write to file
		_salt := rsEncode(rs16, salt)
		fout.Write(_salt)

		// Encode HKDF salt with Reed-Solomon and write to file
		_hkdfSalt := rsEncode(rs32, hkdfSalt)
		fout.Write(_hkdfSalt)

		// Encode Serpent salt with Reed-Solomon and write to file
		_serpentSalt := rsEncode(rs16, serpentSalt)
		fout.Write(_serpentSalt)

		// Encode nonce with Reed-Solomon and write to file
		_nonce := rsEncode(rs24, nonce)
		fout.Write(_nonce)

		// Write placeholder for hash of key
		fout.Write(make([]byte, 192))

		// Write placeholder for hash of hash of keyfile
		fout.Write(make([]byte, 96))

		// Write placeholder for HMAC-BLAKE2b/HMAC-SHA3 of file
		fout.Write(make([]byte, 192))
	} else {
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

		popupStatus = s("Reading values...")
		giu.Update()

		version := make([]byte, 15)
		fin.Read(version)
		_, err1 = rsDecode(rs5, version)

		tmp := make([]byte, 15)
		fin.Read(tmp)
		tmp, err2 = rsDecode(rs5, tmp)
		metadataLength, _ := strconv.Atoi(string(tmp))

		fin.Read(make([]byte, metadataLength*3))

		flags := make([]byte, 18)
		fin.Read(flags)
		flags, err3 = rsDecode(rs6, flags)
		fast = flags[0] == 1
		paranoid = flags[1] == 1
		reedsolo = flags[3] == 1
		padded = flags[4] == 1

		salt = make([]byte, 48)
		fin.Read(salt)
		salt, err4 = rsDecode(rs16, salt)

		hkdfSalt = make([]byte, 96)
		fin.Read(hkdfSalt)
		hkdfSalt, err5 = rsDecode(rs32, hkdfSalt)

		serpentSalt = make([]byte, 48)
		fin.Read(serpentSalt)
		serpentSalt, err6 = rsDecode(rs16, serpentSalt)

		nonce = make([]byte, 72)
		fin.Read(nonce)
		nonce, err7 = rsDecode(rs24, nonce)

		_keyHash = make([]byte, 192)
		fin.Read(_keyHash)
		_keyHash, err8 = rsDecode(rs64, _keyHash)

		_keyfileHash = make([]byte, 96)
		fin.Read(_keyfileHash)
		_keyfileHash, err9 = rsDecode(rs32, _keyfileHash)

		dataMac = make([]byte, 192)
		fin.Read(dataMac)
		dataMac, err10 = rsDecode(rs64, dataMac)

		// Is there a better way?
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil ||
			err6 != nil || err7 != nil || err8 != nil || err9 != nil || err10 != nil {
			if keep {
				kept = true
			} else {
				mainStatus = "The header is corrupt and the input file cannot be decrypted."
				mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
				fin.Close()
				return
			}
		}
	}

	popupStatus = s("Deriving key...")
	progress = 0
	progressInfo = ""
	giu.Update()

	// Derive encryption/decryption keys and subkeys
	var key []byte
	if fast {
		key = argon2.IDKey(
			[]byte(password),
			salt,
			4,
			131072,
			4,
			32,
		)
	} else if paranoid {
		key = argon2.IDKey(
			[]byte(password),
			salt,
			8,
			1048576,
			8,
			32,
		)
	} else {
		key = argon2.IDKey(
			[]byte(password),
			salt,
			4,
			1048576,
			4,
			32,
		)
	}

	if !working {
		mainStatus = "Operation cancelled by user."
		mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
		if mode == "encrypt" && (len(allFiles) > 1 || len(onlyFolders) > 0) {
			os.Remove(outputFile)
		}
		if recombine {
			os.Remove(inputFile)
		}
		os.Remove(outputFile)
		return
	}

	if len(keyfiles) > 0 || keyfile {
		if keyfileOrderMatters {
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
		} else {
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

	sha3_512 := sha3.New512()
	sha3_512.Write(key)
	keyHash = sha3_512.Sum(nil)

	// Validate password and/or keyfiles
	if mode == "decrypt" {
		keyCorrect := true
		keyfileCorrect := true
		var tmp bool

		keyCorrect = !(subtle.ConstantTimeCompare(keyHash, _keyHash) == 0)
		if keyfile {
			keyfileCorrect = !(subtle.ConstantTimeCompare(keyfileHash, _keyfileHash) == 0)
			tmp = !keyCorrect || !keyfileCorrect
		} else {
			tmp = !keyCorrect
		}

		if tmp || keep {
			if keep {
				kept = true
			} else {
				fin.Close()
				if !keyCorrect {
					mainStatus = "The provided password is incorrect."
				} else {
					if keyfileOrderMatters {
						mainStatus = "Incorrect keyfiles and/or order."
					} else {
						mainStatus = "Incorrect keyfiles."
					}
				}
				mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
				key = nil
				if recombine {
					os.Remove(inputFile)
				}
				return
			}
		}

		var err error
		fout, err = os.Create(outputFile)
		if err != nil {
			mainStatus = "Access denied by operating system."
			mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
			return
		}
	}

	if len(keyfiles) > 0 || keyfile {
		// XOR key and keyfile
		tmp := key
		key = make([]byte, 32)
		for i := range key {
			key[i] = tmp[i] ^ keyfileKey[i]
		}
	}

	done := 0
	counter := 0
	startTime := time.Now()
	chacha20, _ := chacha20.NewUnauthenticatedCipher(key, nonce)

	// Use HKDF-SHA3 to generate a subkey
	var mac hash.Hash
	subkey := make([]byte, 32)
	hkdf := hkdf.New(sha3.New256, key, hkdfSalt, nil)
	hkdf.Read(subkey)
	if fast {
		// Keyed BLAKE2b
		mac, _ = blake2b.New512(subkey)
	} else {
		// HMAC-SHA3
		mac = hmac.New(sha3.New512, subkey)
	}

	// Generate another subkey and cipher (not used unless paranoid mode is checked)
	serpentKey := make([]byte, 32)
	hkdf.Read(serpentKey)
	_serpent, _ := serpent.NewCipher(serpentKey)
	serpentCTR := cipher.NewCTR(_serpent, serpentSalt)

	for {
		if !working {
			mainStatus = "Operation cancelled by user."
			mainStatusColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
			fin.Close()
			fout.Close()
			if mode == "encrypt" && (len(allFiles) > 1 || len(onlyFolders) > 0) {
				os.Remove(outputFile)
			}
			if recombine {
				os.Remove(inputFile)
			}
			os.Remove(outputFile)
			return
		}

		var data []byte
		if mode == "decrypt" && reedsolo {
			data = make([]byte, 1114112)
		} else {
			data = make([]byte, 1048576)
		}

		size, err := fin.Read(data)
		if err != nil {
			break
		}
		data = data[:size]
		_data := make([]byte, len(data))

		// "Actual" encryption is done in the next couple of lines
		if mode == "encrypt" {
			if paranoid {
				serpentCTR.XORKeyStream(_data, data)
				copy(data, _data)
			}

			chacha20.XORKeyStream(_data, data)
			mac.Write(_data)

			if reedsolo {
				copy(data, _data)
				_data = nil
				if len(data) == 1048576 {
					for i := 0; i < 1048576; i += 128 {
						tmp := data[i : i+128]
						tmp = rsEncode(rs128, tmp)
						_data = append(_data, tmp...)
					}
				} else {
					chunks := math.Floor(float64(len(data)) / 128)
					for i := 0; float64(i) < chunks; i++ {
						tmp := data[i*128 : (i+1)*128]
						tmp = rsEncode(rs128, tmp)
						_data = append(_data, tmp...)
					}
					tmp := data[int(chunks*128):]
					_data = append(_data, rsEncode(rs128, pad(tmp))...)
				}
			}
		} else {
			if reedsolo {
				copy(_data, data)
				data = nil
				if len(_data) == 1114112 {
					for i := 0; i < 1114112; i += 136 {
						tmp := _data[i : i+136]
						tmp, err = rsDecode(rs128, tmp)
						if err != nil {
							if keep {
								kept = true
							} else {
								mainStatus = "The input file is too corrupted to decrypt."
								mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
								fin.Close()
								fout.Close()
								broken()
								return
							}
						}
						if i == 1113976 && done+1114112 >= int(total) && padded {
							tmp = unpad(tmp)
						}
						data = append(data, tmp...)
					}
				} else {
					chunks := len(_data)/136 - 1
					for i := 0; i < chunks; i++ {
						tmp := _data[i*136 : (i+1)*136]
						tmp, err = rsDecode(rs128, tmp)
						if err != nil {
							if keep {
								kept = true
							} else {
								mainStatus = "The input file is too corrupted to decrypt."
								mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
								fin.Close()
								fout.Close()
								broken()
								return
							}
						}
						data = append(data, tmp...)
					}
					tmp := _data[int(chunks)*136:]
					tmp, err = rsDecode(rs128, tmp)
					if err != nil {
						if keep {
							kept = true
						} else {
							mainStatus = "The input file is too corrupted to decrypt."
							mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
							fin.Close()
							fout.Close()
							broken()
							return
						}
					}
					tmp = unpad(tmp)
					data = append(data, tmp...)
				}
				_data = make([]byte, len(data))
			}

			mac.Write(data)
			chacha20.XORKeyStream(_data, data)

			if paranoid {
				copy(data, _data)
				serpentCTR.XORKeyStream(_data, data)
			}
		}
		fout.Write(_data)

		// Update stats
		if mode == "decrypt" && reedsolo {
			done += 1114112
		} else {
			done += 1048576
		}
		counter++
		progress = float32(done) / float32(total)
		elapsed := float64(time.Since(startTime)) / math.Pow(10, 9)
		speed := float64(done) / elapsed / math.Pow(10, 6)
		eta := int(math.Floor(float64(total-int64(done)) / (speed * math.Pow(10, 6))))

		if progress > 1 {
			progress = 1
		}

		progressInfo = fmt.Sprintf("%.2f%%", progress*100)
		popupStatus = fmt.Sprintf(s("Working at %.2f MB/s (ETA: %s)"), speed, humanize(eta))
		giu.Update()
	}

	if mode == "encrypt" {
		// Seek back to header and write important data
		fout.Seek(int64(312+len(metadata)*3), 0)
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
		popupStatus = s("Splitting file...")
		stat, _ := os.Stat(outputFile)
		size := stat.Size()
		finished := 0
		chunkSize, _ := strconv.Atoi(splitSize)

		// User can choose KiB, MiB, and GiB
		if splitSelected == 0 {
			chunkSize *= 1024
		} else if splitSelected == 1 {
			chunkSize *= 1048576
		} else {
			chunkSize *= 1073741824
		}
		chunks := int(math.Ceil(float64(size) / float64(chunkSize)))
		fin, _ := os.Open(outputFile)

		for i := 0; i < chunks; i++ {
			fout, _ := os.Create(fmt.Sprintf("%s.%d", outputFile, i))
			done := 0
			for {
				data := make([]byte, 1048576)
				read, err := fin.Read(data)
				if err != nil {
					break
				}
				if !working {
					fin.Close()
					fout.Close()
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
			}
			fout.Close()
			finished++
			splitted = append(splitted, fmt.Sprintf("%s.%d", outputFile, i))
			progress = float32(finished) / float32(chunks)
			progressInfo = fmt.Sprintf("%d/%d", finished, chunks)
			giu.Update()
		}
		fin.Close()
		if shredTemp {
			progressInfo = ""
			popupStatus = s("Shredding temporary files...")
			shred([]string{inputFile + ".pcv"}, false)
		} else {
			os.Remove(inputFile + ".pcv")
		}
	}

	// Remove the temporary file used to combine a splitted Picocrypt volume
	if recombine {
		os.Remove(inputFile)
	}

	// Delete the temporary zip file if user wishes
	if len(allFiles) > 1 || len(onlyFolders) > 0 {
		if shredTemp {
			progressInfo = ""
			popupStatus = s("Shredding temporary files...")
			giu.Update()
			shred([]string{inputFile}, false)
		} else {
			os.Remove(inputFile)
		}
	}

	if deleteWhenDone {
		progressInfo = ""
		popupStatus = s("Deleted files...")
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

	resetUI()

	// If user chose to keep a corrupted/modified file, let them know
	if kept {
		mainStatus = "The input file is corrupted and/or modified. Please be careful."
		mainStatusColor = color.RGBA{0xff, 0xff, 0x00, 0xff}
	} else {
		mainStatus = "Completed."
		mainStatusColor = color.RGBA{0x00, 0xff, 0x00, 0xff}
	}

	// Clear UI state
	working = false
	kept = false
	key = nil
	popupStatus = s("Ready.")
}

// This function is run if an issue occurs during decryption
func broken() {
	mainStatus = "The input file is either corrupted or intentionally modified."
	mainStatusColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
	if recombine {
		os.Remove(inputFile)
	}
	os.Remove(outputFile)
	giu.Update()
}

// Generate file checksums
func generateChecksums(file string) {
	fin, _ := os.Open(file)

	// Clear UI state
	csMd5 = ""
	csSha1 = ""
	csSha256 = ""
	csSha3 = ""
	csBlake2b = ""
	csBlake2s = ""
	md5Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
	sha1Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
	sha256Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
	sha3Color = color.RGBA{0x00, 0x00, 0x00, 0x00}
	blake2bColor = color.RGBA{0x00, 0x00, 0x00, 0x00}
	blake2sColor = color.RGBA{0x00, 0x00, 0x00, 0x00}
	csValidate = ""

	if md5Selected {
		csMd5 = s("Calculating...")
	}
	if sha1Selected {
		csSha1 = s("Calculating...")
	}
	if sha256Selected {
		csSha256 = s("Calculating...")
	}
	if sha3Selected {
		csSha3 = s("Calculating...")
	}
	if blake2bSelected {
		csBlake2b = s("Calculating...")
	}
	if blake2sSelected {
		csBlake2s = s("Calculating...")
	}

	// Create the checksum objects
	crcMd5 := md5.New()
	crcSha1 := sha1.New()
	crcSha256 := sha256.New()
	crcSha3 := sha3.New256()
	crcBlake2b, _ := blake2b.New256(nil)
	crcBlake2s, _ := blake2s.New256(nil)

	stat, _ := os.Stat(file)
	total := stat.Size()
	var done int64 = 0
	for {
		var data []byte
		_data := make([]byte, 1048576)
		size, err := fin.Read(_data)
		if err != nil {
			break
		}
		data = _data[:size]

		if md5Selected {
			crcMd5.Write(data)
		}
		if sha1Selected {
			crcSha1.Write(data)
		}
		if sha256Selected {
			crcSha256.Write(data)
		}
		if sha3Selected {
			crcSha3.Write(data)
		}
		if blake2bSelected {
			crcBlake2b.Write(data)
		}
		if blake2sSelected {
			crcBlake2s.Write(data)
		}

		done += int64(size)
		csProgress = float32(done) / float32(total)
		giu.Update()
	}
	csProgress = 0
	if md5Selected {
		csMd5 = hex.EncodeToString(crcMd5.Sum(nil))
	}
	if sha1Selected {
		csSha1 = hex.EncodeToString(crcSha1.Sum(nil))
	}
	if sha256Selected {
		csSha256 = hex.EncodeToString(crcSha256.Sum(nil))
	}
	if sha3Selected {
		csSha3 = hex.EncodeToString(crcSha3.Sum(nil))
	}
	if blake2bSelected {
		csBlake2b = hex.EncodeToString(crcBlake2b.Sum(nil))
	}
	if blake2sSelected {
		csBlake2s = hex.EncodeToString(crcBlake2s.Sum(nil))
	}

	fin.Close()
	giu.Update()
}

// Recursively shred all file(s) and folder(s) passed in as 'names'
func shred(names []string, separate bool) {
	stopShredding = false
	shredTotal = 0
	shredDone = 0

	// 'separate' is true if this function is being called from the encryption/decryption tab
	if separate {
		shredOverlay = s("Shredding...")
	}

	// Walk through directories to get the total number of files for statistics
	for _, name := range names {
		filepath.Walk(name, func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			stat, _ := os.Stat(path)
			if !stat.IsDir() {
				shredTotal++
			}
			return nil
		})
	}

	for _, name := range names {
		shredding = name

		// Linux and macOS need a command with similar syntax and usage, so they're combined
		if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
			stat, _ := os.Stat(name)
			if stat.IsDir() {
				var coming []string

				// Walk the folder recursively
				filepath.Walk(name, func(path string, _ os.FileInfo, err error) error {
					if err != nil {
						return nil
					}
					if stopShredding {
						return nil
					}
					stat, _ := os.Stat(path)
					if !stat.IsDir() {
						if len(coming) == 128 {
							// Use a WaitGroup to parallelize shredding
							var wg sync.WaitGroup
							for i, j := range coming {
								wg.Add(1)
								go func(wg *sync.WaitGroup, id int, j string) {
									defer wg.Done()
									shredding = j
									var cmd *exec.Cmd
									if runtime.GOOS == "linux" {
										cmd = exec.Command("shred", "-ufvz", "-n", strconv.Itoa(int(shredPasses)), j)
									} else {
										cmd = exec.Command("rm", "-rfP", j)
									}
									cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
									cmd.Run()
									shredDone++
									shredUpdate(separate)
									giu.Update()
								}(&wg, i, j)
							}
							wg.Wait()
							coming = nil
						} else {
							coming = append(coming, path)
						}
					}
					return nil
				})
				for _, i := range coming {
					if stopShredding {
						break
					}
					go func(i string) {
						shredding = i
						var cmd *exec.Cmd
						if runtime.GOOS == "linux" {
							cmd = exec.Command("shred", "-ufvz", "-n", strconv.Itoa(int(shredPasses)), i)
						} else {
							cmd = exec.Command("rm", "-rfP", i)
						}
						cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
						cmd.Run()
						shredDone++
						shredUpdate(separate)
						giu.Update()
					}(i)
				}
				if !stopShredding {
					os.RemoveAll(name)
				}
			} else { // The path is a file, not a directory, so just shred it
				shredding = name
				var cmd *exec.Cmd
				if runtime.GOOS == "linux" {
					cmd = exec.Command("shred", "-ufvz", "-n", strconv.Itoa(int(shredPasses)), name)
				} else {
					cmd = exec.Command("rm", "-rfP", name)
				}
				cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				cmd.Run()
				shredDone++
				shredUpdate(separate)
			}
		} else if runtime.GOOS == "windows" {
			stat, _ := os.Stat(name)
			if stat.IsDir() {
				// Walk the folder recursively
				filepath.Walk(name, func(path string, _ os.FileInfo, err error) error {
					if err != nil {
						return nil
					}
					stat, _ := os.Stat(path)
					if stat.IsDir() {
						if stopShredding {
							return nil
						}

						t := 0
						files, _ := ioutil.ReadDir(path)
						for _, f := range files {
							if !f.IsDir() {
								t++
							}
						}
						shredDone += float32(t)
						shredUpdate(separate)
						shredding = strings.ReplaceAll(path, "\\", "/") + "/*"
						cmd := exec.Command(sdelete64path, "*", "-p", strconv.Itoa(int(shredPasses)))
						cmd.Dir = path
						cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
						cmd.Run()
						giu.Update()
					}
					return nil
				})

				if !stopShredding {
					// sdelete64 doesn't delete the empty folder, so I'll do it manually
					os.RemoveAll(name)
				}
			} else {
				shredding = name
				cmd := exec.Command(sdelete64path, name, "-p", strconv.Itoa(int(shredPasses)))
				cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				cmd.Run()
				shredDone++
				shredUpdate(separate)
			}
		}
		giu.Update()
		if stopShredding {
			return
		}
	}

	// Clear UI state
	shredding = s("Completed.")
	shredProgress = 0
	shredOverlay = ""
}

// Update shredding statistics
func shredUpdate(separate bool) {
	if separate {
		shredOverlay = fmt.Sprintf("%d/%d", int(shredDone), int(shredTotal))
		shredProgress = shredDone / shredTotal
	} else {
		popupStatus = fmt.Sprintf("%d/%d", int(shredDone), int(shredTotal))
		progress = shredDone / shredTotal
	}
	giu.Update()
}

// Reset the UI to a clean state with nothing selected or checked
func resetUI() {
	mode = ""
	onlyFiles = nil
	onlyFolders = nil
	allFiles = nil
	inputLabel = s("Drop files and folders into this window.")
	password = ""
	cPassword = ""
	keyfiles = nil
	keyfile = false
	keyfileOrderMatters = false
	keyfilePrompt = s("None selected.")
	metadata = ""
	metadataPrompt = "Metadata:"
	metadataDisabled = false
	shredTemp = false
	keep = false
	reedsolo = false
	split = false
	splitSize = ""
	splitSelected = 1
	fast = false
	deleteWhenDone = false
	paranoid = false
	compress = false
	encryptFilename = false
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
	var res []byte
	rs.Encode(data, func(s infectious.Share) {
		res = append(res, s.DeepCopy().Data[0])
	})
	return res
}

// Reed-Solomon decoder
func rsDecode(rs *infectious.FEC, data []byte) ([]byte, error) {
	tmp := make([]infectious.Share, rs.Total())
	for i := 0; i < rs.Total(); i++ {
		tmp[i] = infectious.Share{
			Number: i,
			Data:   []byte{data[i]},
		}
	}
	res, err := rs.Decode(nil, tmp)
	if err != nil {
		if rs.Total() == 136 {
			return data[:128], err
		}
		return data[:rs.Total()/3], err
	}
	return res, nil
}

// PKCS7 Pad (for use with Reed-Solomon, not for cryptographic purposes)
func pad(data []byte) []byte {
	padLen := 128 - len(data)%128
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

// PKCS7 Unpad
func unpad(data []byte) []byte {
	length := len(data)
	padLen := int(data[length-1])
	return data[:length-padLen]
}

func genPassword() string {
	chars := ""
	if genpassUpper {
		chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if genpassLower {
		chars += "abcdefghijklmnopqrstuvwxyz"
	}
	if genpassNums {
		chars += "1234567890"
	}
	if genpassSymbols {
		chars += "-=!@#$^&()_+?"
	}
	if chars == "" {
		return chars
	}
	tmp := make([]byte, genpassLength)
	for i := 0; i < int(genpassLength); i++ {
		j, _ := rand.Int(rand.Reader, new(big.Int).SetUint64(uint64(len(chars))))
		tmp[i] = chars[j.Int64()]
	}
	if genpassCopy {
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

func s(term string) string {
	for _, i := range locales {
		if i.iso == selectedLocale {
			for _, j := range locales {
				if j.iso == "en" {
					for k, l := range j.data {
						if l == term {
							return i.data[k]
						}
					}
				}
			}

		}
	}
	return term
}

func main() {
	// Parse locales
	var obj map[string]json.RawMessage
	json.Unmarshal(localeBytes, &obj)
	for i := range obj {
		var tmp []string
		json.Unmarshal(obj[i], &tmp)
		locales = append(locales, locale{
			iso:  i,
			data: tmp,
		})
	}

	// Check system locale
	for _, i := range locales {
		tmp, err := jibber_jabber.DetectIETF()
		if err == nil {
			if strings.HasPrefix(tmp, i.iso) {
				selectedLocale = i.iso
				for j, k := range allLocales {
					if k == i.iso {
						languageSelected = int32(j)
					}
				}
			}
		}
	}

	// Create a temporary file to store sdelete64.exe
	sdelete64, _ := os.CreateTemp("", "sdelete64.*.exe")
	sdelete64path = sdelete64.Name()
	sdelete64.Write(sdelete64bytes)
	sdelete64.Close()
	cmd := exec.Command(sdelete64path, "/accepteula")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()

	// Set a universal font
	giu.SetDefaultFontFromBytes(font, 18)

	// Create the master window
	window := giu.NewMasterWindow("Picocrypt", 442, 532, giu.MasterWindowFlagsNotResizable)
	dialog.Init()

	// Set window icon
	reader := bytes.NewReader(icon)
	decoded, _ := png.Decode(reader)
	window.SetIcon([]image.Image{decoded})

	// Set callbacks
	window.SetDropCallback(onDrop)
	window.SetCloseCallback(func() bool {
		return !working
	})

	// Set universal DPI
	dpi = giu.Context.GetPlatform().GetContentScale()

	// Start a goroutine to check if a newer version is available
	go func() {
		v, err := http.Get("https://raw.githubusercontent.com/HACKERALERT/Picocrypt/main/internals/version.txt")
		if err == nil {
			body, err := io.ReadAll(v.Body)
			v.Body.Close()
			if err == nil {
				if string(body[:5]) != version {
					mainStatus = "A newer version is available."
					mainStatusColor = color.RGBA{0, 255, 0, 255}
				}
			}
		}
	}()

	// Start the UI
	window.Run(draw)

	// Window closed, clean up
	os.Remove(sdelete64path)
}
