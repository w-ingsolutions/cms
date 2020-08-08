package cms

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/theme"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/w-ingsolutions/c/pkg/lyt"
	osnovna "github.com/w-ingsolutions/cms/DEF"
	"github.com/w-ingsolutions/cms/pkg/phi"
	"github.com/w-ingsolutions/cms/pkg/φ"
)

func sveOdTipa(ctx context.Context, sh *shell.Shell, th *theme.DuoUItheme, tip phi.ContentType) func() {
	return func() {
		widgets := make(map[string]interface{})
		prikazLista := []func(gtx C) D{}

		if tip.SlugPlural == "materijali" {
			files := osnovna.NewMaterijali()
			for _, sadrzajFiles := range files {
				widgets[fmt.Sprint(sadrzajFiles.ID)] = new(widget.Clickable)
				fmt.Println("sadrzajFiles", sadrzajFiles.Category)

				///////////////
				var network bytes.Buffer
				//Create an encoder and send a value.
				enc := gob.NewEncoder(&network)
				err := enc.Encode(podesavanja)
				checkError(err)
				path := "/" + podesavanja.Dir + "/" + "materijali/φ" + fmt.Sprint(sadrzajFiles.ID)
				err = sh.FilesWrite(ctx, path, &network, shell.FilesWrite.Create(true))
				fmt.Println("sadrzajFilesCategoryStruct", path)
				////////////////
			}
			prikaz.w = widgets
			for _, row := range files {
				var r phi.Φ
				r = row
				prikazLista = append(prikazLista, rowList(ctx, sh, th, r.Struct, widgets[fmt.Sprint(r.ID)].(*widget.Clickable), tip.SlugPlural, r.Struct["Title"].Content.(string), r.Struct["Slug"].Content.(string)))
				//fmt.Println("r.Struct", r.Struct["Title"])
			}
		}
		if tip.SlugPlural == "radovi" {
			files := osnovna.NewRadovi()
			for _, sadrzajFiles := range files {
				widgets[sadrzajFiles.Category+"_"+fmt.Sprint(sadrzajFiles.ID)] = new(widget.Clickable)
				//fmt.Println("sadrzajFiles>>>>>", sadrzajFiles.Category+"_"+fmt.Sprint(sadrzajFiles.ID))
				///////////////
				var network bytes.Buffer
				//Create an encoder and send a value.
				enc := gob.NewEncoder(&network)
				err := enc.Encode(podesavanja)
				checkError(err)
				path := "/" + podesavanja.Dir + "/" + "radovi/26/1/φ" + fmt.Sprint(sadrzajFiles.ID)
				err = sh.FilesWrite(ctx, path, &network, shell.FilesWrite.Create(true))
				fmt.Println("sadrzajFilesCategoryStruct", path)
				////////////////
			}
			prikaz.w = widgets
			for _, row := range files {
				var r phi.Φ
				r = row
				sl := r.Struct["Slug"].Content.(string)

				prikazLista = append(prikazLista, rowList(ctx, sh, th, r.Struct, widgets[r.Category+"_"+fmt.Sprint(r.ID)].(*widget.Clickable), tip.SlugPlural, r.Struct["Title"].Content.(string), sl))
				//fmt.Println("slslsl::::::::::::", sl)
			}
		}
		if tip.SlugPlural != "materijali" && tip.SlugPlural != "radovi" {
			files, err := sh.FilesLs(ctx, podesavanja.Dir+"/"+tip.SlugPlural)
			checkError(err)
			for _, sadrzajFiles := range files {
				widgets[sadrzajFiles.Name] = new(widget.Clickable)
			}
			prikaz.w = widgets
			for _, row := range files {
				file, err := sh.FilesRead(ctx, podesavanja.Dir+"/"+tip.SlugPlural+"/"+row.Name)
				checkError(err)
				var s phi.Φ
				dec := gob.NewDecoder(file)
				err = dec.Decode(&s)
				checkError(err)
				prikazLista = append(prikazLista, rowList(ctx, sh, th, s.Struct, sadrzajDugme, tip.SlugPlural, s.Struct["Title"].Content.(string), s.Struct["Slug"].Content.(string)))
			}
		}
		prikaz.e = prikazLista
	}
}

func rowList(ctx context.Context, sh *shell.Shell, th *theme.DuoUItheme, struktura map[string]phi.F, btn *widget.Clickable, tip, naziv, slug string) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,f(0.5,_),f(0.5,_),r(_))",
			material.Body1(th.T, naziv).Layout,
			material.Body1(th.T, slug).Layout,
			stranaDugme(th, btn, φΦφ(ctx, sh, th, tip, struktura), "Uredi", "sadrzaj"),
		)
	}
}
