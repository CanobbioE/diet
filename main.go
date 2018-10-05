package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

const (
	version = "v0.0"
)

var (
	showHelp    bool
	mealToAdd   string
	newFoodItem string
	debug       bool
	w           *astilectron.Window
	AppName     string
	BuiltAt     string
)

func usage(appName, version string) {
	fmt.Printf("Usage: %s [OPTIONS]\n", appName)
	fmt.Println("OPTIONS:\n")
	flag.VisitAll(func(f *flag.Flag) {
		if len(f.Name) > 1 {
			fmt.Printf("\t-%s, -%s\t%s\n", f.Name[0:1], f.Name, f.Usage)
		}
	})
	fmt.Printf("\nVersion: %s\n", version)
}

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "event.name":
		var s string
		if err = json.Unmarshal(m.Payload, &s); err != nil {
			payload = err.Error()
			return
		}
	}
	return
}

func init() {
	// TODO: create db if it doesn't exist
	// read flags
	flag.BoolVar(&showHelp, "h", false, "dispaly help")
	flag.BoolVar(&showHelp, "help", false, "dispaly help")
	flag.BoolVar(&debug, "d", false, "enables debug mode")
	flag.BoolVar(&debug, "debug", false, "enables debug mode")
}

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	astilog.Debugf("Running app built at %s", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
		},
		Debug: debug,
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.PtrStr("Menu"),
		}},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			go func() {
				time.Sleep(5 * time.Second)
				if err := bootstrap.SendMessage(w, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
					astilog.Error(errors.Wrap(err, "sending check.out.menu event failed"))
				}
			}()
			return nil
		},
		Windows: []*bootstrap.Window{{
			Homepage: "index.html",
			Options: &astilectron.WindowOptions{
				BackgroundColor: astilectron.PtrStr("#333"),
				Center:          astilectron.PtrBool(true),
				Height:          astilectron.PtrInt(700),
				Width:           astilectron.PtrInt(700),
			},
			MessageHandler: handleMessages,
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
