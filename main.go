package main

import (
	"fmt"
	"gioui.org/app"
	_ "gioui.org/app/permission/storage"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/gioapp/gel/helper"
	cms "github.com/w-ingsolutions/cms/app"
	"github.com/w-ingsolutions/cms/cfg"
	in "github.com/w-ingsolutions/cms/cfg/ini"
	"log"
	"os"
)

func main() {

	w := cms.NewWingCMS()

	if cfg.Initial {
		fmt.Println("running initial sync")
	}
	in.Init(w.Podesavanja.File)

	go func() {
		defer os.Exit(0)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *cms.WingCMS) error {
	for {
		select {
		case e := <-w.UI.Window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				w.UI.Context = layout.NewContext(&w.UI.Ops, e)
				helper.Fill(w.UI.Context, helper.HexARGB(w.UI.Tema.Colors["Light"]))

				if !w.API.OK {
					//w.GreskaEkran()
				} else {
					w.GlavniEkran(w.UI.Context)
				}

				e.Frame(w.UI.Context.Ops)
			}
			w.UI.Window.Invalidate()
		}
	}
}