package cms

import (
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (w *WingCMS) lista() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return sadrzajList.Layout(gtx, len(w.Prikaz.e), func(gtx C, i int) D {
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				w.Prikaz.e[i],
				helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]),
			)
		})
	}
}

//func (w *WingCMS) sveOdTipa() {
//	sadrzaj := w.Prikaz[i]
//	return lyt.Format(gtx, "hflexb(middle,r(_),f(0.1,_),f(0.9,_),r(_))",
//		w.akcijaDugme(sadrzaj.Izmena, "Run"),
//		func(gtx C) D {
//			return material.Body1(w.UI.Tema.T, fmt.Sprint(sadrzaj.Id)).Layout(gtx)
//		},
//		func(gtx C) D {
//			return material.Body1(w.UI.Tema.T, sadrzaj.Naziv).Layout(gtx)
//		},
//		w.akcijaDugme(sadrzaj.Brisanje, "Delete"),
//	)
//}
