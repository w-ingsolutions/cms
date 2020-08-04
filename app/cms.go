package cms

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/utl"
)

func (w *WingCMS) GlavniEkran(gtx layout.Context) {
	lyt.Format(gtx, "hflexb(start,r(_),f(1,_))",
		Meni(w.UI.Tema, w.TipoviSadrzajaPrikaz, w.Strana.Slug),
		func(gtx C) D {
			return lyt.Format(gtx, "vflexb(start,r(_),f(1,_),r(_))",
				header(w),
				utl.Strana(w.UI.Tema, w.lista(), w.Strana.Naziv),
				footer(w),
			)
		})
}

func (w *WingCMS) sumaFooter(t string) func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		suma := material.Body2(w.UI.Tema.T, t)
		suma.Alignment = text.End
		return suma.Layout(gtx)
	}
}
