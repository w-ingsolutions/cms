package cms

import (
	"gioui.org/app"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/theme"
	wapp "github.com/w-ingsolutions/c/pkg/app"
	"github.com/w-ingsolutions/c/pkg/icons"
	"path/filepath"
)

func NewWingCMS() *WingCMS {
	w := &WingCMS{
		Strana: WingStrana{"Komandna tabla", "komandna_tabla"},
		//Db:     db.DuoUIdbInit("BAZA"),
	}
	w.Podesavanja = &WingPodesavanja{
		Naziv: "Kalkulator",
		Dir:   wapp.Dir("wing", false),
		Cyr:   false,
	}
	w.Podesavanja.File = filepath.Join(w.Podesavanja.Dir, "conf.json")
	w.UI = WingUI{
		Tema: theme.NewDuoUItheme(),
	}
	w.API = WingAPI{
		OK:     true,
		Adresa: "https://wing.marcetin.com/",
	}

	w.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1280), unit.Dp(1024)),
		app.Title("W-ing "),
	)

	counters := WingCounters{
		Kolicina: &counter.DuoUIcounter{
			Value:        1,
			OperateValue: 1,
			From:         1,
			To:           999,
			CounterInput: &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			CounterIncrease: new(widget.Clickable),
			CounterDecrease: new(widget.Clickable),
			CounterReset:    new(widget.Clickable),
		},
	}
	w.UI.Counters = counters
	w.UI.Tema.Icons = icons.NewWingUIicons()

	w.TipoviSadrzaja = tipoviSadrzaja()

	return w
}

//
//func (w *WingCMS) UcitavanjeTipaSadrzaja() {
//	for _, t := range w.TipoviSadrzaja {
//		w.Prikaz = w.Db.DbReadAll(t.SlugMnozina)
//	}
//}
