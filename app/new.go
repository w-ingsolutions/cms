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
	"github.com/w-ingsolutions/cms/db"
	"path/filepath"
)

func NewWingCMS() *WingCMS {
	w := &WingCMS{
		Strana: WingStrana{"Komandna tabla", "komandna_tabla"},
		Db:     db.DuoUIdbInit("BAZA"),
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
	//m := model.WingMaterijal{
	//	Id:                2,
	//	Naziv:             "Masa za španski zid",
	//	Opis:              "Masa za španski zid",
	//	Obracun:           "Obračun po kilogrammu",
	//	Proizvodjac:       "evrojug",
	//	OsobineNamena:     "Španski zid je.",
	//	NacinRada:         "Masa se meša s.",
	//	JedinicaPotrosnje: "m2/kg",
	//	Potrosnja:         2,
	//	RokUpotrebe:       "12 meseci od datuma proizvodnje istaknutog na ambalaži. Cuvati u originalnoj, dobro zatvorenoj i neoštecenoj ambalaži, pri temperaturi od +5°C do +25",
	//	Jedinica:          "kg",
	//	Pakovanje:         25,
	//	Cena:              0.19,
	//	Slug:              "masa_za_spanski_zid",
	//}
	//
	//var bytesBuf bytes.Buffer
	//encoder := gob.NewEncoder(&bytesBuf)
	//err := encoder.Encode(m)
	//if err != nil {
	//
	//}
	//w.Db.DB.Write(bytesBuf.Bytes())
	w.tipoviSadrzajaPrikaz()
	//var materijal model.WingMaterijal
	//w.Db.DB.Read("bafybeihs3p4g232wocqd5ouoakrwm5kocjaexpar6oekkn6qzez5nqj3vu", &materijal)
	//fmt.Println("WingMaterijal", materijal)
	return w
}

//
//func (w *WingCMS) UcitavanjeTipaSadrzaja() {
//	for _, t := range w.TipoviSadrzaja {
//		w.Prikaz = w.Db.DbReadAll(t.SlugMnozina)
//	}
//}
