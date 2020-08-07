package cms

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func (w *WingCMS) kategorije(tip content.TypePrikaz) func() {
	return func() {
		widgets := map[string]interface{}{
			"struktura": &layout.List{
				Axis: layout.Vertical,
			},
		}
		prikazLista := []func(gtx C) D{}
		for _, kat := range tip.Kategorije {
			widgets[kat.Slug+"_dugme"] = new(widget.Clickable)
			widgets[kat.Slug+"_lista"] = &layout.List{
				Axis: layout.Vertical,
			}
		}
		prikaz.w = widgets
		for _, row := range tip.Kategorije {
			var r sadrzaj.Kategorija
			r = row
			if r.Parent == "" {
				podKategorije := []sadrzaj.Kategorija{}

				for _, pk := range tip.Kategorije {
					if pk.Parent == row.Slug {
						podKategorije = append(podKategorije, pk)
					}
				}

				prikazLista = append(prikazLista, w.rowListKategorija(r, podKategorije, prikaz.w[r.Slug+"_lista"].(*layout.List), sadrzajDugme))
			}
		}
		prikaz.e = prikazLista
	}
}

func (w *WingCMS) rowListKategorija(kategorija sadrzaj.Kategorija, podKategorije []sadrzaj.Kategorija, list *layout.List, btn *widget.Clickable) func(gtx C) D {
	return func(gtx C) D {

		return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
			func(gtx C) D {
				return lyt.Format(gtx, "hflexb(middle,f(0.5,_),f(0.5,_),r(_))",
					material.Body1(w.UI.Tema.T, kategorija.Title).Layout,
					material.Body1(w.UI.Tema.T, kategorija.Slug).Layout,
					material.Body1(w.UI.Tema.T, kategorija.Slug).Layout,
				)
				//w.stranaDugme(btn, w.sadrzaj(naziv, struktura), "Uredi", "sadrzaj"),
			},
			func(gtx C) D {
				return list.Layout(gtx, len(podKategorije), func(gtx C, i int) D {
					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
						func(gtx C) D {
							return lyt.Format(gtx, "hflexb(middle,r(_),f(0.5,_),f(0.5,_))",
								material.Body1(w.UI.Tema.T, podKategorije[i].Parent).Layout,
								material.Body1(w.UI.Tema.T, podKategorije[i].Title).Layout,
								material.Body1(w.UI.Tema.T, podKategorije[i].Slug).Layout,
							)
						},
						helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]),
					)
				})
			},
		)
	}
}
