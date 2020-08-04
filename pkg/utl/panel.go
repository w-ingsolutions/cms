package utl

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func Panel(th *theme.DuoUItheme, naslov string, stavke, sadrzaj, footer func(gtx layout.Context) layout.Dimensions) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
		return container.DuoUIcontainer(th, 1, th.Colors["LightGrayIII"]).Layout(gtx, layout.Center, func(gtx layout.Context) layout.Dimensions {
			return container.DuoUIcontainer(th, 0, th.Colors["White"]).Layout(gtx, layout.Center, func(gtx layout.Context) layout.Dimensions {
				return lyt.Format(gtx, "vflexb(middle,r(_),r(_),f(1,_),r(_))",
					func(gtx layout.Context) layout.Dimensions {
						return container.DuoUIcontainer(th, 4, th.Colors["Primary"]).Layout(gtx, layout.W, func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							suma := material.H6(th.T, naslov)
							suma.Alignment = text.Start
							return suma.Layout(gtx)
						})
					},
					func(gtx layout.Context) layout.Dimensions {
						return container.DuoUIcontainer(th, 4, th.Colors["Gray"]).Layout(gtx, layout.W, stavke)
					},
					sadrzaj,
					func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min.X = gtx.Constraints.Max.X
						return container.DuoUIcontainer(th, 0, th.Colors["LightGrayII"]).Layout(gtx, layout.SE, footer)
					})
			})
		})
	}
}
