package cms

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/φ"
)

func categories(th *theme.DuoUItheme, tip φ.ContentType) func() {
	return func() {
		widgets := map[string]interface{}{
			"struktura": &layout.List{
				Axis: layout.Vertical,
			},
		}
		prikazLista := []func(gtx C) D{}
		for _, kat := range tip.Categories {
			widgets[kat.Slug+"_dugme"] = new(widget.Clickable)
			widgets[kat.Slug+"_lista"] = &layout.List{
				Axis: layout.Vertical,
			}
		}
		prikaz.w = widgets
		for _, row := range tip.Categories {
			var r φ.C
			r = row
			if r.Parent == "" {
				podKategorije := []φ.C{}

				for _, pk := range tip.Categories {
					if pk.Parent == row.Slug {
						podKategorije = append(podKategorije, pk)
					}
				}

				prikazLista = append(prikazLista, rowListKategorija(th, r, podKategorije, prikaz.w[r.Slug+"_lista"].(*layout.List), sadrzajDugme))
			}
		}
		prikaz.e = prikazLista
	}
}

func rowListKategorija(th *theme.DuoUItheme, kategorija φ.C, podKategorije []φ.C, list *layout.List, btn *widget.Clickable) func(gtx C) D {
	return func(gtx C) D {

		return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
			func(gtx C) D {
				return lyt.Format(gtx, "hflexb(middle,f(0.5,_),f(0.5,_),r(_))",
					material.Body1(th.T, kategorija.Title).Layout,
					material.Body1(th.T, kategorija.Slug).Layout,
					material.Body1(th.T, kategorija.Slug).Layout,
				)
				//w.stranaDugme(btn, w.sadrzaj(naziv, struktura), "Uredi", "sadrzaj"),
			},
			func(gtx C) D {
				return list.Layout(gtx, len(podKategorije), func(gtx C, i int) D {
					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
						func(gtx C) D {
							return lyt.Format(gtx, "hflexb(middle,r(_),f(0.5,_),f(0.5,_))",
								material.Body1(th.T, podKategorije[i].Parent).Layout,
								material.Body1(th.T, podKategorije[i].Title).Layout,
								material.Body1(th.T, podKategorije[i].Slug).Layout,
							)
						},
						helper.DuoUIline(false, 0, 0, 1, th.Colors["Gray"]),
					)
				})
			},
		)
	}
}
