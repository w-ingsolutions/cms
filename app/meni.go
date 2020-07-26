package cms

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/icontextbtn"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

var (
	meniList = &layout.List{
		Axis: layout.Vertical,
	}
	sadrzajList = &layout.List{
		Axis: layout.Vertical,
	}
	pomocList = &layout.List{
		Axis: layout.Vertical,
	}
)

func (w *WingCMS) Meni() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 4, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.N, func(gtx C) D {
			gtx.Constraints.Max.X = 180
			gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
			return lyt.Format(gtx, "vflexe(middle,r(_),r(_),r(_))",
				func(gtx C) D {
					ic := w.UI.Tema.Icons["logo"]
					ic.Color = helper.HexARGB("ffb8df42")
					return ic.Layout(gtx, unit.Dp(32))
				},
				func(gtx C) D {
					return meniList.Layout(gtx, len(w.TipoviSadrzajaPrikaz), func(gtx C, i int) D {
						tipSadrzaja := w.TipoviSadrzajaPrikaz[i]
						btn := icontextbtn.IconTextBtn(w.UI.Tema.T, tipSadrzaja.Link, w.UI.Tema.Icons[tipSadrzaja.SlugMnozina], unit.Dp(48), w.UI.Tema.Colors["Light"], tipSadrzaja.NazivMnozina)
						btn.TextSize = unit.Dp(16)
						btn.CornerRadius = unit.Dp(0)
						btn.Background = helper.HexARGB(w.UI.Tema.Colors["Gray"])
						w.LinkoviMenijaKlik(tipSadrzaja)
						b := lyt.Format(gtx, "vflexb(middle,r(_))",
							func(gtx C) D {
								return btn.Layout(gtx)
							},
						)
						if w.Strana.Slug == tipSadrzaja.SlugMnozina {
							b = lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
								func(gtx C) D {
									return btn.Layout(gtx)
								},
								w.tipSadrzajaPodMeni(tipSadrzaja),
							)
						}
						return b
					})
				},
				w.stranaDugme(noviTipDugme, w.noviTip(), "Dodaj Novi Tip", "novi_tip"),
			)
		})
	}
}

func (w *WingCMS) LinkoviMenijaKlik(l model.TipSadrzaja) {
	for l.Link.Clicked() {
		w.Strana = WingStrana{l.Naziv, l.SlugMnozina}
		//w.Prikaz = w.Db.DbReadAll(l.SlugMnozina)
	}
}

func (w *WingCMS) tipSadrzajaPodMeni(s model.TipSadrzaja) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflexb(middle,r(_),r(_),r(_),r(_))",
			w.tipSadrzajaPodMeniDugme(sveOdTipaSadrzajaDugme, "Sve od "+s.NazivMnozina),
			w.tipSadrzajaPodMeniDugme(dodajSadrzajDugme, "Dodaj novi"+s.Naziv),
			w.tipSadrzajaPodMeniDugme(kategorijeSadrzajaDugme, "Kategorije"),
			w.tipSadrzajaPodMeniDugme(oznakeSadrzajaDugme, "Oznake"),
		)

	}
}
func (w *WingCMS) tipSadrzajaPodMeniDugme(l *widget.Clickable, n string) func(gtx C) D {
	return func(gtx C) D {
		btn := icontextbtn.IconTextBtn(w.UI.Tema.T, l, w.UI.Tema.Icons["stolarski"], unit.Dp(48), w.UI.Tema.Colors["Light"], " "+n)
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(12)
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["LightGrayII"])
		return btn.Layout(gtx)
	}
}
