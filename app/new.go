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
	osnovna "github.com/w-ingsolutions/cms/DEF"
	"github.com/w-ingsolutions/cms/pkg/phi"

	//osnovna "github.com/w-ingsolutions/cms/DEF"
)

func NewWingCMS(settings WingPodesavanja) *WingCMS {
	podesavanja = settings
	log.SetLogLevel("*", "fatal")

	//currentPage = Page{"Komandna tabla", "komandna_tabla"}
	w := &WingCMS{
		sh:  shell.NewShell("localhost:5001"),
		ctx: context.Background(),
	}
	p, err := w.sh.FilesLs(w.ctx, "/")
	checkError(err)
	var wing bool
	for _, l := range p {
		if l.Name == "/"+podesavanja.Dir {
			wing = true
		}
		fmt.Println("lllllllllllllllllllllllll", l)

	}
	if wing {
		fmt.Println("Ima folder")
		p, err := w.sh.FilesRead(w.ctx, "/"+podesavanja.Dir+"/φ")
		checkError(err)
		if p != nil {
			checkError(err)
			//read := bytes.NewReader(p)
			dec := gob.NewDecoder(p)
			err = dec.Decode(&podesavanja)
			fmt.Println("Ima podesavanje")
		} else {
			fmt.Println("Nema podesavanje")
			w.osnovnoPodesavanje(podesavanja)
		}
		checkError(err)
	} else {
		err := w.sh.FilesMkdir(w.ctx, "/"+podesavanja.Dir)
		checkError(err)
		w.osnovnoPodesavanje(podesavanja)
		fmt.Println("Nema podesavanja")
	}

	w.UI = WingUI{
		Tema: theme.NewDuoUItheme(),
	}
	w.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1280), unit.Dp(1024)),
		app.Title("podesavanja.Title"),
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
	//////

	/////////////////

	err = w.sh.FilesMkdir(w.ctx, "/wing/radovi")
	checkError(err)
	err = w.sh.FilesMkdir(w.ctx, "/wing/materijali")
	checkError(err)
	for _, radCat := range osnovna.RadoviKategorije() {

		w.radoviCatUpis(radCat)

	}
	////////////////

	/////////////////
	for _, molerskiCat := range osnovna.MolerskiPodKategorije() {
		w.molerskiCatUpis(molerskiCat)
	}
	////////////////

	//////
	//var bytesBuf bytes.Buffer
	//encoder := gob.NewEncoder(&bytesBuf)
	//err := encoder.Encode(m)
	//if err != nil {
	//
	//}
	//w.Db.DB.Write(bytesBuf.Bytes())
	displayTypes(w.tipoviSadrzaja)

	return w
}

//
//func (w *WingCMS) UcitavanjeTipaSadrzaja() {
//	for _, t := range w.TipoviSadrzaja {
//		w.Prikaz = w.Db.DbReadAll(t.SlugPlural)
//	}
//}

func (w *WingCMS) osnovnoPodesavanje(podesavanja WingPodesavanja) {
	w.tipoviSadrzaja = CreateContentTypes()
	var network bytes.Buffer
	//Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(podesavanja)
	checkError(err)
	err = w.sh.FilesWrite(w.ctx, "/"+podesavanja.Dir+"/φ", &network, shell.FilesWrite.Create(true))
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}

func (w *WingCMS) radoviCatUpis(radCat phi.C ) {
	var network bytes.Buffer
	//Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(radCat)
	checkError(err)
	err = w.sh.FilesMkdir(w.ctx, "/wing/radovi/"+fmt.Sprint(radCat.ID))
	checkError(err)
	path := "/wing/radovi/" + fmt.Sprint(radCat.ID) + "/φ"
	err = w.sh.FilesWrite(w.ctx, path, &network, shell.FilesWrite.Create(true))
	checkError(err)
	fmt.Println("sadrzajFilesCategorySaaaaasasas334truct", path)
}

func (w *WingCMS) molerskiCatUpis(molerskiCat phi.C ) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(molerskiCat)
	checkError(err)
	err = w.sh.FilesMkdir(w.ctx, "/wing/radovi/26/"+fmt.Sprint(molerskiCat.ID))
	checkError(err)
	path := "/wing/radovi/26/" + fmt.Sprint(molerskiCat.ID) + "/φ"
	err = w.sh.FilesWrite(w.ctx, path, &network, shell.FilesWrite.Create(true))
	checkError(err)
}
