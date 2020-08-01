package cms

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"gioui.org/app"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/theme"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/ipfs/go-log/v2"
	"github.com/w-ingsolutions/c/pkg/icons"
)

func NewWingCMS() *WingCMS {
	log.SetLogLevel("*", "fatal")
	ctx := context.Background()
	w := &WingCMS{
		Strana: WingStrana{"Komandna tabla", "komandna_tabla"},
		sh:     shell.NewShell("localhost:5001"),
		ctx:    ctx,
	}
	p, err := w.sh.FilesLs(ctx, "/")
	checkError(err)
	//fmt.Println("ppp", p)
	var wing bool
	for _, l := range p {
		if l.Name == "wing" {
			wing = true
		}
	}
	if wing {
		fmt.Println("Ima folder")
		p, err := w.sh.FilesRead(ctx, w.Podesavanja.Dir+"/config")
		checkError(err)
		if p != nil {
			checkError(err)
			//read := bytes.NewReader(p)
			dec := gob.NewDecoder(p) // Will read from network.
			err = dec.Decode(&w.Podesavanja)
			fmt.Println("Ima podesavanje")
		} else {
			fmt.Println("Nema podesavanje")
			w.osnovnoPodesavanje()
		}
		checkError(err)
	} else {
		err := w.sh.FilesMkdir(ctx, w.Podesavanja.Dir)
		checkError(err)
		w.osnovnoPodesavanje()
		fmt.Println("Nema podesavanja")
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

func (w *WingCMS) osnovnoPodesavanje() {
	w.Podesavanja = WingPodesavanja{
		Naziv:          "CMS",
		Dir:            "/wing",
		Cyr:            false,
		TipoviSadrzaja: tipoviSadrzaja(),
	}
	var network bytes.Buffer // Stand-in for the network.
	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(w.Podesavanja)
	checkError(err)
	err = w.sh.FilesWrite(w.ctx, w.Podesavanja.Dir+"/config", &network, shell.FilesWrite.Create(true))
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
