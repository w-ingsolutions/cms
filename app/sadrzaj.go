package cms

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
	"github.com/w-ingsolutions/cms/pkg/utl"
)

func (w *WingCMS) sadrzaj(slug string, struktura map[string]sadrzaj.PoljeSadrzaja) func() {
	return func() {
		widgets := map[string]interface{}{
			"dodajdugme": new(widget.Clickable),
		}
		w.makeWidgets(struktura, widgets)

		//struktura := make(map[string]sadrzaj.PoljeSadrzaja)
		mid := func(gtx C) D {
			return lyt.Format(gtx, "hflexb(middle,r(_))",
				material.Body1(w.UI.Tema.T, "asasas").Layout,
			)
		}
		dugme := w.dugme(struktura, slug)
		prikazLista := w.makeLayout(struktura)
		prikazLista = append(prikazLista, mid)
		prikazLista = append(prikazLista, dugme)
		w.Prikaz.e = prikazLista
	}
}

func (w *WingCMS) isFolder(path, folder string) (i int, b bool) {
	p, err := w.sh.FilesLs(w.ctx, path)
	checkError(err)
	for _, l := range p {
		if l.Name == folder {
			b = true
		}
		i++
	}
	return
}

func (w *WingCMS) makeWidgets(struktura map[string]sadrzaj.PoljeSadrzaja, widgets map[string]interface{}) {
	for _, polje := range struktura {
		var p interface{}
		switch polje.Tip {
		case "Text":
			p = &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     false,
			}
		case "Num":
			p = &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     false,
			}
		}
		widgets[polje.Naziv] = p
	}
	w.Prikaz.w = widgets
}

func (w *WingCMS) makeLayout(struktura map[string]sadrzaj.PoljeSadrzaja) []func(gtx C) D {
	prikazLista := []func(gtx C) D{}
	for _, po := range struktura {
		switch po.Tip {
		case "Text":
			field := w.Prikaz.w[po.Naziv].(*widget.Editor)
			if po.Sadrzaj != "" {
				field.SetText(po.Sadrzaj.(string))
			}
			prikazLista = append(prikazLista, utl.Editor(w.UI.Tema, field, false, po.Naziv, func(e widget.EditorEvent) {}))
		case "Num":
			field := w.Prikaz.w[po.Naziv].(*widget.Editor)
			if po.Sadrzaj != "" {
				field.SetText(po.Sadrzaj.(string))
			}
			prikazLista = append(prikazLista, utl.Editor(w.UI.Tema, field, false, po.Naziv, func(e widget.EditorEvent) {}))
		}
	}
	return prikazLista
}

func (w *WingCMS) dugme(struktura map[string]sadrzaj.PoljeSadrzaja, slug string) func(gtx C) D {
	return func(gtx C) D {
		d := w.Prikaz.w["dodajdugme"].(*widget.Clickable)
		btn := material.Button(w.UI.Tema.T, d, "Dodaj")
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(12)
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["Secondary"])
		for d.Clicked() {
			for _, p := range struktura {
				polje := sadrzaj.PoljeSadrzaja{
					Naziv: p.Naziv,
					Tip:   p.Tip,
				}
				switch p.Tip {
				case "Text":
					polje.Sadrzaj = w.Prikaz.w[p.Naziv].(*widget.Editor).Text()
				case "Num":
					polje.Sadrzaj = w.Prikaz.w[p.Naziv].(*widget.Editor).Text()
				}
				struktura[p.Naziv] = polje
			}
			brojSadrzaja, folder := w.isFolder(w.Podesavanja.Dir+"/"+slug, slug)
			sadrzajUnos := sadrzaj.Sadrzaj{
				ID:        brojSadrzaja + 1,
				Struktura: struktura,
			}
			var write bytes.Buffer
			enc := gob.NewEncoder(&write)
			err := enc.Encode(sadrzajUnos)
			checkError(err)
			if !folder {
				err := w.sh.FilesMkdir(w.ctx, w.Podesavanja.Dir+"/"+slug)
				checkError(err)
			}
			err = w.sh.FilesWrite(w.ctx, w.Podesavanja.Dir+"/"+slug+"/"+fmt.Sprint(brojSadrzaja+1), &write, shell.FilesWrite.Create(true))
			checkError(err)
		}
		return btn.Layout(gtx)
	}
}
