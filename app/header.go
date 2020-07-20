package cms

import (
	"gioui.org/layout"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func header(w *WingCMS) func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 4, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.Center, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			return lyt.Format(gtx, "hflexb(middle,f(1,_))",
				w.headerMenu(),
			)
		})
	}
}
func (w *WingCMS) headerMenu() func(gtx C) D {
	r := func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_),r(_),r(_))",
			w.stranaDugme(radoviDugme, func() {}, "RADOVI", "radovi"),
			helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["DarkGrayI"]),
			w.stranaDugme(sumaDugme, func() {}, "SUMA", "sumaRadovi"),
			helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["DarkGrayI"]),
			w.stranaDugme(sumaMaterialDugme, func() {}, "SUMA MATERIJAL", "sumaMaterijal"),
		)
	}
	if w.UI.Device != "mob" {
		r = func(gtx C) D {
			return lyt.Format(gtx, "hflexb(middle,r(_))",
				w.stranaDugme(radoviDugme, func() {}, "RADOVI", "radovi"),
			)
		}
	}
	return r
}
