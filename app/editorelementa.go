package cms

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/theme"
)

//func (w *WingCMS) EditorElementaIzgled() func(gtx C) D {
//	return func(gtx C) D {
//		return layout.Flex{
//			Axis: layout.Vertical,
//		}.Layout(gtx,
//			//layout.Rigid(func(gtx C) D {
//			//	container.DuoUIcontainer( w.UI.Tema, 8, w.UI.Tema.Colors["LightGray"]).Layout(gtx, layout.W, func(gtx C) D {
//			//		return material.H5(w.UI.Tema.T, w.PrikazaniElement.Naziv).Layout(gtx)
//			//	})
//			//}),
//			layout.Flexed(1, func(gtx C) D {
//				container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["LightGray"]).Layout(gtx, layout.NW, func(gtx C) D {
//					layout.Flex{Axis: layout.Vertical}.Layout(gtx,
//						layout.Rigid(func(gtx C) D {
//							layout.Flex{Axis: layout.Vertical}.Layout(gtx,
//								layout.Rigid(Editor(gtx, w.UI.Tema, w.EditPolja.Naziv, layout.Vertical, "Naziv", func(e widget.EditorEvent) {})),
//								layout.Rigid(func(gtx C) D {
//									layout.Flex{
//										Axis:    layout.Horizontal,
//										Spacing: layout.SpaceBetween,
//									}.Layout(gtx,
//										layout.Rigid(Editor(gtx, w.UI.Tema, w.EditPolja.Id, layout.Vertical, "Id", func(e widget.EditorEvent) {})),
//										layout.Rigid(Editor(gtx, w.UI.Tema, w.EditPolja.Obracun, layout.Vertical, "Obracun", func(e widget.EditorEvent) {})),
//										layout.Rigid(Editor(gtx, w.UI.Tema, w.EditPolja.Jedinica, layout.Vertical, "Jedinica", func(e widget.EditorEvent) {})),
//										layout.Rigid(Editor(gtx, w.UI.Tema, w.EditPolja.Cena, layout.Vertical, "Cena", func(e widget.EditorEvent) {})))
//								}),
//								layout.Rigid(Editor(gtx, w.UI.Tema, w.EditPolja.Opis, layout.Vertical, "Opis", func(e widget.EditorEvent) {})),
//								layout.Flexed(1, func(gtx C) D {
//									//materijalElementaPanelElement.PanelObject = w.Materijal
//									//materijalElementaPanelElement.PanelObjectsNumber = len(w.Materijal)
//									//materijalElementaPanel := w.UI.Tema.DuoUIpanel()
//									//materijalElementaPanel.ScrollBar = w.Tema.ScrollBar(0)
//									materijalElementaPanel.Layout(gtx, materijalElementaPanelElement, func(i int, in interface{}) {
//										//if in != nil {
//										//addresses := in.([]model.DuoUIaddress)
//										//materijal := w.Materijal[i]
//										layout.Flex{Axis: layout.Vertical}.Layout(gtx,
//											layout.Rigid(func(gtx C) D {
//												layout.Flex{
//													Alignment: layout.Middle,
//												}.Layout(gtx,
//													layout.Rigid(Editor(gtx, w.UI.Tema, w.EditPolja.Materijal[materijal.Id], layout.Horizontal, materijal.Naziv, func(e widget.EditorEvent) {})),
//													//layout.Rigid(func(gtx C) D {
//													//	return material.Body1(materijal.Jedinica).Layout(gtx)
//													//}),
//													//layout.Rigid(func(gtx C) D {
//													//	return material.Caption(materijal.Obracun).Layout(gtx)
//													//}),
//												)
//											}),
//											layout.Rigid(w.UI.Tema.DuoUIline(gtx, 1, 0, 1, w.Tema.Colors["Gray"])),
//										)
//										//}
//									})
//								}),
//							)
//						}),
//
//						//layout.Rigid(Editor(gtx, w.Tema, w.EditabilnaPoljaVrsteRadova[w.PrikazaniElement.Id].Opis, fmt.Sprint(w.PrikazaniElement.PodvrsteRadova[w.PrikazaniElement.Id].Opis), func(e gel.EditorEvent) {})),
//					)
//				})
//			}))
//	}
//}

func Editor(th *theme.DuoUItheme, editorControler *widget.Editor, axis layout.Axis, label string, handler func(widget.EditorEvent)) func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{Axis: axis}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return material.H6(th.T, label).Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return container.DuoUIcontainer(th, 8, "ffffffff").Layout(gtx, layout.NW, func(gtx C) D {
					e := material.Editor(th.T, editorControler, "")
					e.Font.Typeface = th.Fonts["Mono"]
					e.TextSize = unit.Dp(12)
					return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
						return e.Layout(gtx)
					})
					for _, e := range editorControler.Events() {
						switch e.(type) {
						case widget.ChangeEvent:
							handler(e)
						}
					}
					return D{}
				})
			}))
	}
}
