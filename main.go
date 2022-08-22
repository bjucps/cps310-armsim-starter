package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"log"
	"os"
)

const htmlAbout = `Golang ARMSim v0.0.1`

// Vars injected via ldflags by bundler
var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

// Application Vars
var (
	help  = flag.Bool("h", false, "shows help message")
	debug = false
	w     *astilectron.Window
)

type LoadFileMessage struct {
	Filename string `json:"filename"`
}

type LoadFileResponse struct {
	Checksum int `json:"checksum"`
}

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "load.file":
		// get the payload into a usable format
		var s string
		if err = json.Unmarshal(m.Payload, &s); err != nil {
			payload = err
			return
		}
		fmt.Println("got body " + s)
		//  convert the payload into one of our user-defined types
		var loadFileMessage LoadFileMessage
		var loadFileResponse LoadFileResponse
		if err = json.Unmarshal([]byte(s), &loadFileMessage); err != nil {
			payload = err
			return
		}
		fmt.Println("got file " + loadFileMessage.Filename)
		// send a message back to the browser
		loadFileResponse.Checksum = 0
		payload = loadFileResponse
	}
	return
}

func main() {
	flag.BoolVar(&debug, "d", false, "enables debug mode")

	// Create logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Parse flags
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Run bootstrap
	l.Printf("Running app built at %s\n", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
		},
		Debug:  debug,
		Logger: l,
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("About"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if err := bootstrap.SendMessage(w, "about", htmlAbout, func(m *bootstrap.MessageIn) {
							// Unmarshal payload
							var s string
							if err := json.Unmarshal(m.Payload, &s); err != nil {
								l.Println(fmt.Errorf("unmarshaling payload failed: %w", err))
								return
							}
							l.Printf("About modal has been displayed and payload is %s!\n", s)
						}); err != nil {
							l.Println(fmt.Errorf("sending about event failed: %w", err))
						}
						return
					},
				},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astikit.StrPtr("#333"),
				Center:          astikit.BoolPtr(true),
				Height:          astikit.IntPtr(700),
				Width:           astikit.IntPtr(700),
			},
		}},
	}); err != nil {
		l.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
	}

}
