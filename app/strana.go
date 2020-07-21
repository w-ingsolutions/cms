package cms

func (w *WingCMS) strana() func(gtx C) D {
	return w.Panel(w.Strana.Naziv, func(gtx C) D { return D{} }, w.lista(), func(gtx C) D { return D{} })
}
