package cms

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
)

func (w WingCMS) akcijaDugme(b *widget.Clickable, icon string) func(gtx C) D {
	return func(gtx C) D {
		btn := material.IconButton(w.UI.Tema.T, b, w.UI.Tema.Icons[icon])
		btn.Inset = layout.Inset{unit.Dp(5), unit.Dp(3), unit.Dp(5), unit.Dp(5)}
		btn.Color = helper.HexARGB(w.UI.Tema.Colors["Danger"])
		btn.Size = unit.Dp(16)
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["White"])
		for b.Clicked() {

		}
		return btn.Layout(gtx)
	}
}
