package cms

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/icontextbtn"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
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

func Meni(th *theme.DuoUItheme, tipoviSadrzajaPrikaz []sadrzaj.TipSadrzajaPrikaz, slug string) func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(th, 4, th.Colors["DarkGrayI"]).Layout(gtx, layout.N, func(gtx C) D {
			gtx.Constraints.Max.X = 180
			gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
			return lyt.Format(gtx, "vflexe(middle,r(_),r(_),r(_))",
				func(gtx C) D {
					ic := th.Icons["logo"]
					ic.Color = helper.HexARGB("ffb8df42")
					return ic.Layout(gtx, unit.Dp(32))
				},
				func(gtx C) D {
					return meniList.Layout(gtx, len(tipoviSadrzajaPrikaz), func(gtx C, i int) D {
						tipSadrzaja := tipoviSadrzajaPrikaz[i]
						btn := icontextbtn.IconTextBtn(th.T, tipSadrzaja.Link, th.Icons[tipSadrzaja.SlugMnozina], unit.Dp(48), th.Colors["Light"], tipSadrzaja.NazivMnozina)
						btn.TextSize = unit.Dp(16)
						btn.CornerRadius = unit.Dp(0)
						btn.Background = helper.HexARGB(th.Colors["Gray"])
						LinkoviMenijaKlik(tipSadrzaja)
						b := lyt.Format(gtx, "vflexb(middle,r(_))",
							func(gtx C) D {
								return btn.Layout(gtx)
							},
						)
						if slug == tipSadrzaja.SlugMnozina {
							b = lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
								func(gtx C) D {
									return btn.Layout(gtx)
								},
								tipSadrzajaPodMeni(tipSadrzaja),
							)
						}
						return b
					})
				},
				w.stranaDugme(noviTipDugme, w.podesavanjaTipa(sadrzaj.TipSadrzajaPrikaz{}), "Dodaj Novi Tip", "novi_tip"),
			)
		})
	}
}

func LinkoviMenijaKlik(w *WingCMS, l sadrzaj.TipSadrzajaPrikaz) {
	for l.Link.Clicked() {
		w.Strana = WingStrana{l.Naziv, l.SlugMnozina}
		//w.Prikaz = w.Db.DbReadAll(l.SlugMnozina)
	}
	return
}

func tipSadrzajaPodMeni(s sadrzaj.TipSadrzajaPrikaz) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflexb(middle,r(_),r(_),r(_),r(_),r(_))",
			w.tipSadrzajaPodMeniDugme(sveOdTipaSadrzajaDugme, w.sveOdTipa(s), "Sve od "+s.NazivMnozina, s.SlugMnozina),
			w.tipSadrzajaPodMeniDugme(dodajSadrzajDugme, w.sadrzaj(s.SlugMnozina, s.Struktura), "Dodaj novi "+s.Naziv, "novi_"+s.Slug),
			w.tipSadrzajaPodMeniDugme(kategorijeSadrzajaDugme, w.kategorije(s), "Kategorije "+s.NazivMnozina, "kategorije_"+s.SlugMnozina),
			w.tipSadrzajaPodMeniDugme(oznakeSadrzajaDugme, func() {}, "Oznake "+s.NazivMnozina, "oznake_"+s.SlugMnozina),
			w.tipSadrzajaPodMeniDugme(podesavanjeSadrzajaDugme, w.podesavanjaTipa(s), "Podesavanja "+s.NazivMnozina, "podesavanja_"+s.SlugMnozina),
		)
	}
}

func (w *WingCMS) tipSadrzajaPodMeniDugme(b *widget.Clickable, f func(), t, p string) func(gtx C) D {
	return func(gtx C) D {
		btn := icontextbtn.IconTextBtn(w.UI.Tema.T, b, w.UI.Tema.Icons["stolarski"], unit.Dp(48), w.UI.Tema.Colors["Light"], " "+t)
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(12)
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["LightGrayII"])
		for b.Clicked() {
			f()
			w.Strana = WingStrana{t, p}
		}
		return btn.Layout(gtx)
	}
}
