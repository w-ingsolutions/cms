package cms

import (
	"context"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/theme"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/utl"
	"github.com/w-ingsolutions/cms/pkg/φ"
)

func φΦφ(ctx context.Context, sh *shell.Shell, th *theme.DuoUItheme, slug string, struktura map[string]φ.F) func() {
	return func() {
		widgets := map[string]interface{}{
			"dodajdugme": new(widget.Clickable),
		}

		prikaz.w = makeWidgets(struktura, widgets)
		//struktura := make(map[string]φ.F)
		prikazLista := makeLayout(th, struktura)

		mid := func(gtx C) D {
			return lyt.Format(gtx, "hflexb(middle,r(_))",
				material.Body1(th.T, "asasas").Layout,
			)
		}
		d := dugme(ctx, sh, th, struktura, slug)
		prikazLista = append(prikazLista, mid)
		prikazLista = append(prikazLista, d)

		prikaz.e = prikazLista
	}
}

func isFolder(ctx context.Context, sh *shell.Shell, path, folder string) (i int, b bool) {
	//p, err := sh.FilesLs(ctx, path)
	//checkError(err)
	//for _, l := range p {
	//	if l.Name == folder {
	//		b = true
	//	}
	//	i++
	//}
	return
}

func makeWidgets(struktura map[string]φ.F, widgets map[string]interface{}) map[string]interface{} {
	for _, field := range struktura {
		var p interface{}
		switch field.Type {
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

		widgets[field.Title] = p
	}
	return widgets
}

func makeLayout(th *theme.DuoUItheme, struktura map[string]φ.F) []func(gtx C) D {
	prikazLista := []func(gtx C) D{}
	for _, field := range struktura {
		switch field.Type {
		case "Text":
			fieldItem := prikaz.w[field.Title].(*widget.Editor)
			if field.Content != "" {
				fieldItem.SetText(field.Content.(string))
			}
			prikazLista = append(prikazLista, utl.Editor(th, fieldItem, false, field.Title, func(e widget.EditorEvent) {}))
		case "Num":
			fieldItem := prikaz.w[field.Title].(*widget.Editor)
			if field.Content != "" {
				fieldItem.SetText(field.Content.(string))
			}
			prikazLista = append(prikazLista, utl.Editor(th, fieldItem, false, field.Title, func(e widget.EditorEvent) {}))
		}
	}
	return prikazLista
}

func dugme(ctx context.Context, sh *shell.Shell, th *theme.DuoUItheme, struktura map[string]φ.F, slug string) func(gtx C) D {
	return func(gtx C) D {
		d := prikaz.w["dodajdugme"].(*widget.Clickable)
		btn := material.Button(th.T, d, "Dodaj")
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(12)
		btn.Background = helper.HexARGB(th.Colors["Secondary"])
		for d.Clicked() {
			for _, p := range struktura {
				field := φ.F{
					Title: p.Title,
					Type:  p.Type,
				}
				switch p.Type {
				case "Text":
					field.Content = prikaz.w[p.Title].(*widget.Editor).Text()
				case "Num":
					field.Content = prikaz.w[p.Title].(*widget.Editor).Text()
				}
				struktura[p.Title] = field
			}
			//brojSadrzaja, folder := isFolder(ctx,sh ,podesavanja.Dir+"/"+slug, slug)
			//sadrzajUnos := φ.Φ{
			//	ID:     brojSadrzaja + 1,
			//	Struct: struktura,
			//}
			//var write bytes.Buffer
			//enc := gob.NewEncoder(&write)
			//err := enc.Encode(sadrzajUnos)
			//checkError(err)
			//if !folder {
			//	err := sh.FilesMkdir(ctx, podesavanja.Dir+"/"+slug)
			//	checkError(err)
			//}
			//err = sh.FilesWrite(ctx, podesavanja.Dir+"/"+slug+"/"+fmt.Sprint(brojSadrzaja+1), &write, shell.FilesWrite.Create(true))
			//checkError(err)
		}
		return btn.Layout(gtx)
	}
}
