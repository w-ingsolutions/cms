package utl

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func Editor(th *theme.DuoUItheme, editorControler *widget.Editor, vertical bool, label string, handler func(widget.EditorEvent)) func(gtx layout.Context) layout.Dimensions {
	axis := "hflexb(middle, f(0.28,_),f(0.72,_))"
	if vertical {
		axis = "vflexe(middle, r(_),r(_))"
	}
	return func(gtx layout.Context) layout.Dimensions {
		return container.DuoUIcontainer(th, 4, th.Colors["DarkGrayI"]).Layout(gtx, layout.N, func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			return lyt.Format(gtx, axis,
				func(gtx layout.Context) layout.Dimensions {
					return material.H6(th.T, label).Layout(gtx)
				},
				func(gtx layout.Context) layout.Dimensions {
					return container.DuoUIcontainer(th, 8, "ffffffff").Layout(gtx, layout.NW, func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min.X = gtx.Constraints.Max.X
						e := material.Editor(th.T, editorControler, "")
						e.Font.Typeface = th.Fonts["Mono"]
						e.TextSize = unit.Dp(12)
						return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return e.Layout(gtx)
						})
						for _, e := range editorControler.Events() {
							switch e.(type) {
							case widget.ChangeEvent:
								handler(e)
							}
						}
						return layout.Dimensions{}
					})
				})
		})
	}
}
