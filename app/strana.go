package cms

func (w *WingCMS) strana() func(gtx C) D {
	komandnaTablaStrana := w.Panel("Komandna Tabla", func(gtx C) D { return D{} }, func(gtx C) D { return D{} }, func(gtx C) D { return D{} })
	s := func(gtx C) D {
		gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
		return D{}
	}

	sadrzajStrana := w.Panel("Materijali", func(gtx C) D { return D{} }, w.lista(), func(gtx C) D { return D{} })

	switch w.Strana {
	case "komandna_tabla":
		s = komandnaTablaStrana
	case "radovi":
		s = sadrzajStrana
	case "materijali":
		s = sadrzajStrana
		//case "radovi":
		//	s = izbornikStrana
		//case "sumaRadovi":
		//	s = sumaRadovaStrana
		//case "sumaMaterijal":
		//	s = sumaMaterijalStrana
		//case "podesavanja":
		//	s = podesavanjaStrana
	}
	//
	//return func(gtx C) D {
	//	return layout.Flex{
	//		Axis: layout.Horizontal,
	//	}.Layout(gtx,
	//		layout.Flexed(1, func(gtx C) D {
	//			return material.Body1(w.UI.Tema.T, "ssss").Layout(gtx)
	//		}))
	//}
	return s
}
