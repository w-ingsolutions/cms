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
	"log"
	"os"
)

func main() {
	podesavanja := cms.WingPodesavanja{
		Title: "CMS",
		Dir:   "wing",
	}
	if cfg.Initial {
		fmt.Println("running initial sync")
	}
	//in.Init(podesavanja.File)
	w := cms.NewWingCMS(podesavanja)

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

				w.GlavniEkran(w.UI.Context)

				e.Frame(w.UI.Context.Ops)
			}
			w.UI.Window.Invalidate()
		}
	}
}
