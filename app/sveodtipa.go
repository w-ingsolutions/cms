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
				go materijali(ctx, sh, sadrzajFiles)
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
			files261 := osnovna.NewRadovi261()
			for _, sadrzajFiles261 := range files261 {
				widgets[sadrzajFiles261.Category+"_"+fmt.Sprint(sadrzajFiles261.ID)] = new(widget.Clickable)
				//fmt.Println("sadrzajFiles>>>>>", sadrzajFiles.Category+"_"+fmt.Sprint(sadrzajFiles.ID))
				///////////////
				go radovi261(ctx, sh, sadrzajFiles261)
				////////////////
			}
			files263 := osnovna.NewRadovi263()
			for _, sadrzajFiles236 := range files263 {
				widgets[sadrzajFiles236.Category+"_"+fmt.Sprint(sadrzajFiles236.ID)] = new(widget.Clickable)
				//fmt.Println("sadrzajFiles>>>>>", sadrzajFiles.Category+"_"+fmt.Sprint(sadrzajFiles.ID))
				///////////////
				go radovi263(ctx, sh, sadrzajFiles236)
				////////////////
			}
			prikaz.w = widgets
			for _, rowA := range files261 {
				var rA phi.Φ
				rA = rowA
				sl := rA.Struct["Slug"].Content.(string)
				prikazLista = append(prikazLista, rowList(ctx, sh, th, rA.Struct, widgets[rA.Category+"_"+fmt.Sprint(rA.ID)].(*widget.Clickable), tip.SlugPlural, rA.Struct["Title"].Content.(string), sl))
			}
			for _, rowB := range files263 {
				var rB phi.Φ
				rB = rowB
				sl := rB.Struct["Slug"].Content.(string)
				prikazLista = append(prikazLista, rowList(ctx, sh, th, rB.Struct, widgets[rB.Category+"_"+fmt.Sprint(rB.ID)].(*widget.Clickable), tip.SlugPlural, rB.Struct["Title"].Content.(string), sl))
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
				s := readFile(ctx, sh, podesavanja.Dir+"/"+tip.SlugPlural+"/"+row.Name)

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

func materijali(ctx context.Context, sh *shell.Shell, sadrzajFiles phi.Φ) {
	var network bytes.Buffer
	//Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(sadrzajFiles)
	checkError(err)
	path := "/" + podesavanja.Dir + "/" + "materijali/φ" + fmt.Sprint(sadrzajFiles.ID)
	err = sh.FilesWrite(ctx, path, &network, shell.FilesWrite.Create(true))
	fmt.Println("sadrzajFil7900000esCategoryStruct", path)
}

func radovi261(ctx context.Context, sh *shell.Shell, sadrzajFiles261 phi.Φ) {
	var network bytes.Buffer
	//Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(sadrzajFiles261)
	checkError(err)
	path := "/" + podesavanja.Dir + "/" + "radovi/26/1/φ" + fmt.Sprint(sadrzajFiles261.ID)
	err = sh.FilesWrite(ctx, path, &network, shell.FilesWrite.Create(true))
	//fmt.Println("rM1", path)
}

func radovi263(ctx context.Context, sh *shell.Shell, sadrzajFiles236 phi.Φ) {
	var network bytes.Buffer
	//Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(sadrzajFiles236)
	checkError(err)
	path := "/" + podesavanja.Dir + "/" + "radovi/26/3/φ" + fmt.Sprint(sadrzajFiles236.ID)
	err = sh.FilesWrite(ctx, path, &network, shell.FilesWrite.Create(true))
	//fmt.Println("rM3", path)
}

func readFile(ctx context.Context, sh *shell.Shell,path string) (s phi.Φ) {
	file, err := sh.FilesRead(ctx, path)
	checkError(err)
	dec := gob.NewDecoder(file)
	err = dec.Decode(&s)
	checkError(err)
	return
}
