package cms

import (
	"fmt"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/pkg/lyt"
	osnovna "github.com/w-ingsolutions/cms/DEF"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func (w *WingCMS) sveOdTipa(tip sadrzaj.TipSadrzajaPrikaz) func() {
	return func() {
		widgets := make(map[string]interface{})
		prikazLista := []func(gtx C) D{}
		if tip.SlugMnozina == "materijali" {
			files := osnovna.NewMaterijali()
			for _, sadrzajFiles := range files {
				widgets[fmt.Sprint(sadrzajFiles.ID)] = new(widget.Clickable)
				fmt.Println("sadrzajFiles.ID", sadrzajFiles.ID)
			}
			w.Prikaz.w = widgets
			for _, row := range files {
				var r sadrzaj.Sadrzaj
				r = row
				prikazLista = append(prikazLista, w.rowList(r.Struktura, sadrzajDugme, tip.SlugMnozina, r.Struktura["Naziv"].Sadrzaj.(string), r.Struktura["Slug"].Sadrzaj.(string)))
				fmt.Println("sadrzajFiles.Name", row.ID)
			}
		}

		if tip.SlugMnozina == "radovi" {
			files := osnovna.NewRadovi()
			for _, sadrzajFiles := range files {
				widgets[sadrzajFiles.Kategorija+"_"+fmt.Sprint(sadrzajFiles.ID)] = new(widget.Clickable)
				fmt.Println("sadrzajFiles>>>>>", sadrzajFiles.Kategorija+"_"+fmt.Sprint(sadrzajFiles.ID))
			}
			w.Prikaz.w = widgets
			for _, row := range files {
				var r sadrzaj.Sadrzaj
				r = row
				sl := r.Struktura["Slug"].Sadrzaj.(string)

				prikazLista = append(prikazLista, w.rowList(r.Struktura, sadrzajDugme, tip.SlugMnozina, r.Struktura["Naziv"].Sadrzaj.(string), sl))
				fmt.Println("slslsl::::::::::::", sl)
			}
		}
		//if tip.SlugMnozina != "materijali" && tip.SlugMnozina != "radovi" {
		//	files, err := w.sh.FilesLs(w.ctx, w.Podesavanja.Dir+"/"+tip.SlugMnozina)
		//	checkError(err)
		//	for _, sadrzajFiles := range files {
		//		widgets[sadrzajFiles.Name] = new(widget.Clickable)
		//	}
		//	w.Prikaz.w = widgets
		//	for _, row := range files {
		//		file, err := w.sh.FilesRead(w.ctx, w.Podesavanja.Dir+"/"+tip.SlugMnozina+"/"+row.Name)
		//		checkError(err)
		//		var s sadrzaj.Sadrzaj
		//		dec := gob.NewDecoder(file)
		//		err = dec.Decode(&s)
		//		checkError(err)
		//		prikazLista = append(prikazLista, w.rowList(s.Struktura, sadrzajDugme, tip.SlugMnozina, s.Struktura["Naziv"].Sadrzaj.(string), s.Struktura["Slug"].Sadrzaj.(string)))
		//	}
		//}
		w.Prikaz.e = prikazLista
	}
}

func (w *WingCMS) rowList(struktura map[string]sadrzaj.PoljeSadrzaja, btn *widget.Clickable, tip, naziv, slug string) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,f(0.5,_),f(0.5,_),r(_))",
			material.Body1(w.UI.Tema.T, naziv).Layout,
			material.Body1(w.UI.Tema.T, slug).Layout,
			w.stranaDugme(btn, w.sadrzaj(tip, struktura), "Uredi", "sadrzaj"),
		)
	}
}
