package utl

import (
	"gioui.org/layout"
	"github.com/gioapp/gel/theme"
)

func Strana(th *theme.DuoUItheme, lista func(gtx layout.Context) layout.Dimensions, naziv string) func(gtx layout.Context) layout.Dimensions {
	//return lyt.Format(gtx, "vflexe(middle,r(_),r(_),r(_))",
	//	func(gtx C) D {
	//		ic := w.UI.Tema.Icons["logo"]
	//		ic.Color = helper.HexARGB("ffb8df42")
	//		return ic.Layout(gtx, unit.Dp(32))
	//	},
	//
	return Panel(th, naziv, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{} }, lista, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{} })
}
