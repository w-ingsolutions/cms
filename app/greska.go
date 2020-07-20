package cms

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func (w *WingCMS) GreskaEkran() {
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(w.UI.Context,
		layout.Rigid(material.H3(w.UI.Tema.T, "Greška u povezivanju sa bazom ").Layout))
}
