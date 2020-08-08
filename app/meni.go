package cms

import (
	"context"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/icontextbtn"
	"github.com/gioapp/gel/theme"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/phi"
	"github.com/w-ingsolutions/cms/pkg/φ"
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

func Meni(ctx context.Context, sh *shell.Shell, th *theme.DuoUItheme, t map[string]phi.T, tipoviSadrzajaPrikaz []phi.ContentType, slug string) func(gtx C) D {
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
						LinkoviMenijaKlik(tipSadrzaja)
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
								tipSadrzajaPodMeni(ctx, sh, th, t, tipSadrzaja),
							)
						}
						return b
					})
				},
				stranaDugme(th, noviTipDugme, podesavanjaTipa(ctx, sh, th, t, phi.ContentType{}), "Dodaj Novi Tip", "novi_tip"),
			)
		})
	}
}

func LinkoviMenijaKlik(l phi.ContentType) {
	for l.Link.Clicked() {
		currentPage = Page{l.Title, l.SlugPlural}
		//w.Prikaz = w.Db.DbReadAll(l.SlugPlural)
	}
	return
}

//Strana               WingStrana
//EditPolja            *model.EditabilnaPoljaVrsteRadova
//TipoviSadrzajaPrikaz []φ.T
func tipSadrzajaPodMeni(ctx context.Context, sh *shell.Shell, th *theme.DuoUItheme, t map[string]phi.T, s phi.ContentType) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflexb(middle,r(_),r(_),r(_),r(_),r(_))",
			tipSadrzajaPodMeniDugme(th, sveOdTipaSadrzajaDugme, sveOdTipa(ctx, sh, th, s), "Sve od "+s.TitlePlural, s.SlugPlural),
			tipSadrzajaPodMeniDugme(th, dodajSadrzajDugme, φΦφ(ctx, sh, th, s.SlugPlural, s.Struct), "Dodaj novi "+s.Title, "novi_"+s.Slug),
			tipSadrzajaPodMeniDugme(th, kategorijeSadrzajaDugme, categories(th, s), "Kategorije "+s.TitlePlural, "kategorije_"+s.SlugPlural),
			tipSadrzajaPodMeniDugme(th, oznakeSadrzajaDugme, func() {}, "Oznake "+s.TitlePlural, "oznake_"+s.SlugPlural),
			tipSadrzajaPodMeniDugme(th, podesavanjeSadrzajaDugme, podesavanjaTipa(ctx, sh, th, t, s), "Podesavanja "+s.TitlePlural, "podesavanja_"+s.SlugPlural),
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
			currentPage = Page{t, p}
		}
		return btn.Layout(gtx)
	}
}
