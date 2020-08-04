package utl

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/theme"
)

func Cell(th theme.DuoUItheme, align text.Alignment, tekst string) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		cell := material.Caption(th.T, tekst)
		cell.TextSize = unit.Dp(12)
		cell.Alignment = align
		return cell.Layout(gtx)
	}
}
