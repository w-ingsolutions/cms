package cms

import (
	"gioui.org/layout"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/utl"
	"github.com/w-ingsolutions/cms/pkg/φ"
)

var (
	prikaz prikazElementi

	currentPage          = Page{}
	podesavanja          WingPodesavanja
	tipoviSadrzajaPrikaz []phi.ContentType
)

func (w *WingCMS) GlavniEkran(gtx layout.Context) {
	lyt.Format(gtx, "hflexb(start,r(_),f(1,_))",
		Meni(w.ctx, w.sh, w.UI.Tema, w.tipoviSadrzaja, tipoviSadrzajaPrikaz, currentPage.Slug),
		func(gtx C) D {
			return lyt.Format(gtx, "vflexb(start,r(_),f(1,_),r(_))",
				header(w),
				utl.Strana(w.UI.Tema, w.lista(), currentPage.Title),
				footer(w.UI.Tema),
			)
		})
}
