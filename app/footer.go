package cms

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func footer(th *theme.DuoUItheme) func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(th, 4, th.Colors["DarkGrayI"]).Layout(gtx, layout.Center, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_))",
				iconLink(th, podesavanjeDugme, "settingsIcon", func() {
					currentPage = Page{"Podesavanja", "podesavanja"}
				}),
				iconLink(th, pomocDugme, "settingsIcon", func() {
					currentPage = Page{"Pomoc", "pomoc"}
				}),
				footerMenu(th))
		})
	}
}

func footerMenu(th *theme.DuoUItheme) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,r(_),r(_))",
			helper.DuoUIline(true, 0, 2, 2, th.Colors["DarkGrayI"]),
			stranaDugme(th, materijalDugme, func() {}, "MATERIJAL", "materijal"),
		)
	}
}

func iconLink(th *theme.DuoUItheme, b *widget.Clickable, i string, f func()) func(gtx C) D {
	return func(gtx C) D {
		return layout.Center.Layout(gtx, func(gtx C) D {
			btn := material.IconButton(th.T, b, th.Icons[i])
			btn.Color = helper.HexARGB(th.Colors["Danger"])
			btn.Size = unit.Dp(16)
			btn.Inset = layout.Inset{
				Top:    unit.Dp(4),
				Right:  unit.Dp(4),
				Bottom: unit.Dp(4),
				Left:   unit.Dp(4),
			}
			btn.Background = helper.HexARGB(th.Colors["White"])
			for b.Clicked() {
				f()
			}
			return btn.Layout(gtx)
		})
	}
}
