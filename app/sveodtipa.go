package cms

import (
	"fmt"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/pkg/lyt"
	osnovna "github.com/w-ingsolutions/cms/DEF"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func sveOdTipa(th *theme.DuoUItheme, tip content.TypePrikaz) func() {
	return func() {
		widgets := make(map[string]interface{})
		prikazLista := []func(gtx C) D{}
		if tip.SlugPlural == "materijali" {
			files := osnovna.NewMaterijali()
			for _, sadrzajFiles := range files {
				widgets[fmt.Sprint(sadrzajFiles.ID)] = new(widget.Clickable)
				fmt.Println("sadrzajFiles.ID", sadrzajFiles.ID)
			}
			w.Prikaz.w = widgets
			for _, row := range files {
				var r sadrzaj.Sadrzaj
				r = row
				prikazLista = append(prikazLista, rowList(th, r.Struct, sadrzajDugme, tip.SlugPlural, r.Struct["Title"].Sadrzaj.(string), r.Struct["Slug"].Sadrzaj.(string)))
				fmt.Println("sadrzajFiles.Name", row.ID)
			}
		}

		if tip.SlugPlural == "radovi" {
			files := osnovna.NewRadovi()
			for _, sadrzajFiles := range files {
				widgets[sadrzajFiles.Kategorija+"_"+fmt.Sprint(sadrzajFiles.ID)] = new(widget.Clickable)
				fmt.Println("sadrzajFiles>>>>>", sadrzajFiles.Kategorija+"_"+fmt.Sprint(sadrzajFiles.ID))
			}
			w.Prikaz.w = widgets
			for _, row := range files {
				var r sadrzaj.Sadrzaj
				r = row
				sl := r.Struct["Slug"].Sadrzaj.(string)

				prikazLista = append(prikazLista, w.rowList(r.Struct, sadrzajDugme, tip.SlugPlural, r.Struct["Title"].Sadrzaj.(string), sl))
				fmt.Println("slslsl::::::::::::", sl)
			}
		}
		//if tip.SlugPlural != "materijali" && tip.SlugPlural != "radovi" {
		//	files, err := w.sh.FilesLs(w.ctx, w.Podesavanja.Dir+"/"+tip.SlugPlural)
		//	checkError(err)
		//	for _, sadrzajFiles := range files {
		//		widgets[sadrzajFiles.Name] = new(widget.Clickable)
		//	}
		//	w.Prikaz.w = widgets
		//	for _, row := range files {
		//		file, err := w.sh.FilesRead(w.ctx, w.Podesavanja.Dir+"/"+tip.SlugPlural+"/"+row.Name)
		//		checkError(err)
		//		var s sadrzaj.Sadrzaj
		//		dec := gob.NewDecoder(file)
		//		err = dec.Decode(&s)
		//		checkError(err)
		//		prikazLista = append(prikazLista, w.rowList(s.Struct, sadrzajDugme, tip.SlugPlural, s.Struct["Title"].Sadrzaj.(string), s.Struct["Slug"].Sadrzaj.(string)))
		//	}
		//}
		w.Prikaz.e = prikazLista
	}
}

func rowList(th *theme.DuoUItheme, struktura map[string]content.Field, btn *widget.Clickable, tip, naziv, slug string) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,f(0.5,_),f(0.5,_),r(_))",
			material.Body1(th.T, naziv).Layout,
			material.Body1(th.T, slug).Layout,
			stranaDugme(btn, w.sadrzaj(tip, struktura), "Uredi", "sadrzaj"),
		)
	}
}
