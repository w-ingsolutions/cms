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
	"github.com/gioapp/gel/theme"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/content"
	"github.com/w-ingsolutions/cms/pkg/utl"
)

func sadrzaj(th *theme.DuoUItheme, slug string, struktura map[string]content.PoljeSadrzaja) func() {
	return func() {
		widgets := map[string]interface{}{
			"dodajdugme": new(widget.Clickable),
		}
		w.makeWidgets(struktura, widgets)

		//struktura := make(map[string]content.Field)
		mid := func(gtx C) D {
			return lyt.Format(gtx, "hflexb(middle,r(_))",
				material.Body1(th.T, "asasas").Layout,
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

func (w *WingCMS) makeWidgets(struktura map[string]content.Field, widgets map[string]interface{}) {
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
		widgets[polje.Title] = p
	}
	prikaz.w = widgets
}

func (w *WingCMS) makeLayout(struktura map[string]content.Field) []func(gtx C) D {
	prikazLista := []func(gtx C) D{}
	for _, po := range struktura {
		switch po.Tip {
		case "Text":
			field := w.Prikaz.w[po.Title].(*widget.Editor)
			if po.Sadrzaj != "" {
				field.SetText(po.Sadrzaj.(string))
			}
			prikazLista = append(prikazLista, utl.Editor(th, field, false, po.Title, func(e widget.EditorEvent) {}))
		case "Num":
			field := w.Prikaz.w[po.Title].(*widget.Editor)
			if po.Sadrzaj != "" {
				field.SetText(po.Sadrzaj.(string))
			}
			prikazLista = append(prikazLista, utl.Editor(th, field, false, po.Title, func(e widget.EditorEvent) {}))
		}
	}
	return prikazLista
}

func (w *WingCMS) dugme(struktura map[string]content.Field, slug string) func(gtx C) D {
	return func(gtx C) D {
		d := w.Prikaz.w["dodajdugme"].(*widget.Clickable)
		btn := material.Button(th.T, d, "Dodaj")
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(12)
		btn.Background = helper.HexARGB(th.Colors["Secondary"])
		for d.Clicked() {
			for _, p := range struktura {
				polje := content.Field{
					Title: p.Title,
					Tip:   p.Tip,
				}
				switch p.Tip {
				case "Text":
					polje.Sadrzaj = w.Prikaz.w[p.Title].(*widget.Editor).Text()
				case "Num":
					polje.Sadrzaj = w.Prikaz.w[p.Title].(*widget.Editor).Text()
				}
				struktura[p.Title] = polje
			}
			brojSadrzaja, folder := w.isFolder(w.Podesavanja.Dir+"/"+slug, slug)
			sadrzajUnos := sadrzaj.Sadrzaj{
				ID:     brojSadrzaja + 1,
				Struct: struktura,
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
