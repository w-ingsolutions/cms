package cms

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func footer(w *WingCMS) func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 4, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.Center, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_))",
				w.iconLink(podesavanjeDugme, "Podesavanja", "podesavanja", "settingsIcon"),
				w.iconLink(pomocDugme, "Pomoc", "pomoc", "settingsIcon"),
				w.footerMenu())
		})
	}
}

func (w *WingCMS) footerMenu() func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,r(_),r(_))",
			helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["DarkGrayI"]),
			w.stranaDugme(materijalDugme, func() {}, "MATERIJAL", "materijal"),
		)
	}
}

func (w *WingCMS) iconLink(b *widget.Clickable, n, s, i string) func(gtx C) D {
	return func(gtx C) D {
		return layout.Center.Layout(gtx, func(gtx C) D {
			btn := material.IconButton(w.UI.Tema.T, b, w.UI.Tema.Icons[i])
			btn.Color = helper.HexARGB(w.UI.Tema.Colors["Danger"])
			btn.Size = unit.Dp(16)
			btn.Inset = layout.Inset{
				Top:    unit.Dp(4),
				Right:  unit.Dp(4),
				Bottom: unit.Dp(4),
				Left:   unit.Dp(4),
			}
			btn.Background = helper.HexARGB(w.UI.Tema.Colors["White"])
			for b.Clicked() {
				w.Strana = WingStrana{n, s}
			}
			return btn.Layout(gtx)
		})
	}
}