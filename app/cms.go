package cms

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (w *WingCMS) GlavniEkran(gtx layout.Context) {
	lyt.Format(gtx, "hflexb(start,r(_),f(1,_))",
		w.Meni(),
		func(gtx C) D {
			return lyt.Format(gtx, "vflexb(start,r(_),f(1,_),r(_))",
				header(w),
				w.strana(),
				footer(w),
			)
		})
}

func (w *WingCMS) cell(align text.Alignment, tekst string) func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		cell := material.Caption(w.UI.Tema.T, tekst)
		cell.TextSize = unit.Dp(12)
		cell.Alignment = align
		return cell.Layout(gtx)
	}
}

func (w *WingCMS) sumaFooter(t string) func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		suma := material.Body2(w.UI.Tema.T, t)
		suma.Alignment = text.End
		return suma.Layout(gtx)
	}
}
