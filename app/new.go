package cms

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"gioui.org/app"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/theme"
	"github.com/marcetin/jdb"
	wapp "github.com/w-ingsolutions/c/pkg/app"
	"github.com/w-ingsolutions/c/pkg/icons"
	"log"
)

func NewWingCMS() *WingCMS {
	// initialize db options
	opts := jdb.DefaultOptions
	// Set PrivateKey. This should come from an ENV or a secret store in the real world
	opts.PrivateKey, _ = hex.DecodeString("44667768254d593b7ea48c3327c18a651f6031554ca4f5e3e641f6ff1ea72e98")
	db, err := jdb.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	w := &WingCMS{
		Strana: WingStrana{"Komandna tabla", "komandna_tabla"},
		Db:     db,
	}
	p, err := db.Get("podesavanja")
	checkError(err)
	if p != nil {
		read := bytes.NewReader(p)
		dec := gob.NewDecoder(read) // Will read from network.

		var podesavanja WingPodesavanja
		err = dec.Decode(&podesavanja)
		if err != nil {
			log.Fatal("decode error:", err)
		}
		fmt.Println("Ima podesavanja")
		w.Podesavanja = podesavanja
	} else {
		var write bytes.Buffer
		podesavanja := WingPodesavanja{
			Naziv:          "CMS",
			Dir:            wapp.Dir("wing", false),
			Cyr:            false,
			TipoviSadrzaja: tipoviSadrzaja(),
		}
		//var read bytes.Buffer
		enc := gob.NewEncoder(&write) // Will write to network.
		// Encode (send) the value.
		err = enc.Encode(podesavanja)
		if err != nil {
			log.Fatal("encode error:", err)
		}

		err = db.Set("podesavanja", write.Bytes())

		fmt.Println("Nema podesavanja")
		w.Podesavanja = podesavanja
	}
	w.UI = WingUI{
		Tema: theme.NewDuoUItheme(),
	}
	w.API = WingAPI{
		OK:     true,
		Adresa: "https://wing.marcetin.com/",
	}
	w.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1280), unit.Dp(1024)),
		app.Title(w.Podesavanja.Naziv),
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

	//w.Podesavanja.TipoviSadrzaja = tipoviSadrzaja()

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

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
