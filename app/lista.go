package cms

import (
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (w *WingCMS) lista() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return sadrzajList.Layout(gtx, len(prikaz.e), func(gtx C, i int) D {
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				prikaz.e[i],
				helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]),
			)
		})
	}
}
