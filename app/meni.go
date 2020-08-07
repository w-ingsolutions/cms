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

func Meni(th *theme.DuoUItheme, tipoviSadrzajaPrikaz []content.TypePrikaz, slug string) func(gtx C) D {
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
						btn := icontextbtn.IconTextBtn(th.T, tipSadrzaja.Link, th.Icons[tipSadrzaja.SlugPlural], unit.Dp(48), th.Colors["Light"], tipSadrzaja.TitlePlural)
						btn.TextSize = unit.Dp(16)
						btn.CornerRadius = unit.Dp(0)
						btn.Background = helper.HexARGB(th.Colors["Gray"])
						LinkoviMenijaKlik(w, tipSadrzaja)
						b := lyt.Format(gtx, "vflexb(middle,r(_))",
							func(gtx C) D {
								return btn.Layout(gtx)
							},
						)
						if slug == tipSadrzaja.SlugPlural {
							b = lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
								func(gtx C) D {
									return btn.Layout(gtx)
								},
								tipSadrzajaPodMeni(th, tipSadrzaja),
							)
						}
						return b
					})
				},
				stranaDugme(th, noviTipDugme, w.podesavanjaTipa(content.TypePrikaz{}), "Dodaj Novi Tip", "novi_tip"),
			)
		})
	}
}

func LinkoviMenijaKlik(w *WingCMS, l content.TypePrikaz) {
	for l.Link.Clicked() {
		strana = WingStrana{l.Title, l.SlugPlural}
		//w.Prikaz = w.Db.DbReadAll(l.SlugPlural)
	}
	return
}

//Strana               WingStrana
//EditPolja            *model.EditabilnaPoljaVrsteRadova
//TipoviSadrzajaPrikaz []content.TypePrikaz
func tipSadrzajaPodMeni(th *theme.DuoUItheme, s content.TypePrikaz) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflexb(middle,r(_),r(_),r(_),r(_),r(_))",
			tipSadrzajaPodMeniDugme(th, sveOdTipaSadrzajaDugme, w.sveOdTipa(s), "Sve od "+s.TitlePlural, s.SlugPlural),
			tipSadrzajaPodMeniDugme(th, dodajSadrzajDugme, w.sadrzaj(s.SlugPlural, s.Struct), "Dodaj novi "+s.Title, "novi_"+s.Slug),
			tipSadrzajaPodMeniDugme(th, kategorijeSadrzajaDugme, w.kategorije(s), "Kategorije "+s.TitlePlural, "kategorije_"+s.SlugPlural),
			tipSadrzajaPodMeniDugme(th, oznakeSadrzajaDugme, func() {}, "Oznake "+s.TitlePlural, "oznake_"+s.SlugPlural),
			tipSadrzajaPodMeniDugme(th, podesavanjeSadrzajaDugme, w.podesavanjaTipa(s), "Podesavanja "+s.TitlePlural, "podesavanja_"+s.SlugPlural),
		)
	}
}

func tipSadrzajaPodMeniDugme(th *theme.DuoUItheme, b *widget.Clickable, f func(), t, p string) func(gtx C) D {
	return func(gtx C) D {
		btn := icontextbtn.IconTextBtn(th.T, b, th.Icons["stolarski"], unit.Dp(48), th.Colors["Light"], " "+t)
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(12)
		btn.Background = helper.HexARGB(th.Colors["LightGrayII"])
		for b.Clicked() {
			f()
			strana = WingStrana{t, p}
		}
		return btn.Layout(gtx)
	}
}
