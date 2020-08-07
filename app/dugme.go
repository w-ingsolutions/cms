package cms

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/theme"
)

var (
	latcyrDugme       = new(widget.Clickable)
	radoviDugme       = new(widget.Clickable)
	materijalDugme    = new(widget.Clickable)
	sumaDugme         = new(widget.Clickable)
	sumaMaterialDugme = new(widget.Clickable)

	podesavanjeDugme = new(widget.Clickable)
	pomocDugme       = new(widget.Clickable)

	nazadDugme = new(widget.Clickable)

	noviTipDugme = new(widget.Clickable)

	sadrzajDugme             = new(widget.Clickable)
	sveOdTipaSadrzajaDugme   = new(widget.Clickable)
	dodajSadrzajDugme        = new(widget.Clickable)
	kategorijeSadrzajaDugme  = new(widget.Clickable)
	oznakeSadrzajaDugme      = new(widget.Clickable)
	podesavanjeSadrzajaDugme = new(widget.Clickable)
)

func stranaDugme(th *theme.DuoUItheme, b *widget.Clickable, f func(), t, p string) func(gtx C) D {
	return func(gtx C) D {
		btn := material.Button(th.T, b, t)
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(12)
		btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
		btn.Background = helper.HexARGB(th.Colors["Secondary"])
		for b.Clicked() {
			f()
			strana = WingStrana{t, p}
		}
		return btn.Layout(gtx)
	}
}
