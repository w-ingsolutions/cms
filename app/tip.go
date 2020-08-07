package cms

import (
	"context"
	"gioui.org/layout"
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

func podesavanjaTipa(ctx context.Context, sh *shell.Shell, th *theme.DuoUItheme, t map[string]φ.T, tip φ.ContentType) func() {
	return func() {
		widgets := map[string]interface{}{
			"naziv": &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			"nazivmnozina": &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			"slug": &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			"slugmnozina": &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			"nazivTipaSadrzaja": &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			"dodajdugme":         new(widget.Clickable),
			"dodajtipdugme":      new(widget.Clickable),
			"vrstepoljasadrzaja": new(widget.Enum),
			"struktura": &layout.List{
				Axis: layout.Vertical,
			},
			"vrstepolja": &layout.List{
				Axis: layout.Horizontal,
			},
		}

		prikaz.w = widgets

		naziv := prikaz.w["naziv"].(*widget.Editor)
		nazivmnozina := prikaz.w["nazivmnozina"].(*widget.Editor)
		slug := prikaz.w["slug"].(*widget.Editor)
		slugmnozina := prikaz.w["slugmnozina"].(*widget.Editor)
		vrste := prikaz.w["vrstepoljasadrzaja"].(*widget.Enum)
		nazivTipaSadrzaja := prikaz.w["nazivTipaSadrzaja"].(*widget.Editor)
		strukturaLista := prikaz.w["struktura"].(*layout.List)
		vrstePolja := prikaz.w["vrstepolja"].(*layout.List)

		struktura := make(map[string]φ.F)
		if tip.Struct != nil {
			struktura = tip.Struct
		}

		if tip.Title != "" {
			naziv.SetText(tip.Title)
		}
		if tip.TitlePlural != "" {
			nazivmnozina.SetText(tip.TitlePlural)
		}
		if tip.Slug != "" {
			slug.SetText(tip.Slug)
		}
		if tip.SlugPlural != "" {
			slugmnozina.SetText(tip.SlugPlural)
		}

		prikaz.e = []func(gtx C) D{
			utl.Editor(th, naziv, false, "Title", func(e widget.EditorEvent) {}),
			utl.Editor(th, nazivmnozina, false, "Title mnozina", func(e widget.EditorEvent) {}),
			utl.Editor(th, slug, false, "Slug", func(e widget.EditorEvent) {}),
			utl.Editor(th, slugmnozina, false, "Slug mnozina", func(e widget.EditorEvent) {}),
			strukturaListaIzgled(th, struktura, strukturaLista),
			dodajtipdugme(th, struktura, nazivTipaSadrzaja, vrstePolja, vrste),
			dodajdugme(ctx, sh, th, t, struktura),
		}
	}
}

func vrstePoljaSadrzaja(th *material.Theme, l *layout.List, enum *widget.Enum) func(gtx C) D {
	v := []string{
		"Text",
		"Num",
		"Array",
	}
	return func(gtx C) D {
		return l.Layout(gtx, len(v), func(gtx C, i int) D {
			return material.RadioButton(th, enum, v[i], v[i]).Layout(gtx)
		})
	}
}

func strukturaListaIzgled(th *theme.DuoUItheme, struktura map[string]φ.F, list *layout.List) func(gtx C) D {
	return func(gtx C) D {
		var strArray []φ.F
		for _, str := range struktura {
			strArray = append(strArray, φ.F{
				Title: str.Title,
				Type:  str.Type,
			})
		}
		return list.Layout(gtx, len(strArray), func(gtx C, i int) D {
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				func(gtx C) D {
					return lyt.Format(gtx, "hflexb(middle,r(_),r(_))",
						material.Body1(th.T, strArray[i].Title).Layout,
						material.Body1(th.T, strArray[i].Type).Layout,
					)
				},
				helper.DuoUIline(false, 0, 0, 1, th.Colors["Gray"]),
			)
		})
	}
}

func dodajtipdugme(th *theme.DuoUItheme, struktura map[string]φ.F, nazivTipaSadrzaja *widget.Editor, vrstePolja *layout.List, vrste *widget.Enum) func(gtx C) D {
	return func(gtx C) D {
		d := prikaz.w["dodajtipdugme"].(*widget.Clickable)
		return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_))",
			utl.Editor(th, nazivTipaSadrzaja, false, "Title tipa sadrzaja", func(e widget.EditorEvent) {}),
			vrstePoljaSadrzaja(th.T, vrstePolja, vrste),
			iconLink(th, d, "counterPlusIcon", func() {
				struktura[nazivTipaSadrzaja.Text()] = φ.F{
					Title: nazivTipaSadrzaja.Text(),
					Type:  vrste.Value,
				}
			}))
	}
}

func dodajdugme(ctx context.Context, sh *shell.Shell, th *theme.DuoUItheme, t map[string]φ.T, s map[string]φ.F) func(gtx C) D {
	return func(gtx C) D {
		d := prikaz.w["dodajdugme"].(*widget.Clickable)
		btn := material.Button(th.T, d, "Dodaj")
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(12)
		//btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
		btn.Background = helper.HexARGB(th.Colors["Secondary"])
		for d.Clicked() {
			tipSadrzaja := φ.T{
				Title:       prikaz.w["naziv"].(*widget.Editor).Text(),
				TitlePlural: prikaz.w["nazivmnozina"].(*widget.Editor).Text(),
				Slug:        prikaz.w["slug"].(*widget.Editor).Text(),
				SlugPlural:  prikaz.w["slugmnozina"].(*widget.Editor).Text(),
				Struct:      s,
			}
			t[prikaz.w["slug"].(*widget.Editor).Text()] = tipSadrzaja
			displayTypes(t)
			//var write bytes.Buffer
			//enc := gob.NewEncoder(&write)
			//err := enc.Encode(podesavanja)
			//checkError(err)
			//err = sh.FilesWrite(ctx, "/wing/config", &write, shell.FilesWrite.Create(true))
			//fmt.Println("PodesavanjaSnima")
		}
		return btn.Layout(gtx)
	}
}
