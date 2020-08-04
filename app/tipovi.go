package cms

import (
	"gioui.org/widget"
	osnovni "github.com/w-ingsolutions/cms/DEF"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func tipoviSadrzaja() map[string]sadrzaj.TipSadrzaja {
	return map[string]sadrzaj.TipSadrzaja{
		"radovi":    osnovni.Radovi(),
		"materijal": osnovni.Materijal(),
	}
}

func (w *WingCMS) tipoviSadrzajaPrikaz() {
	var tipovi []sadrzaj.TipSadrzajaPrikaz
	for _, t := range w.Podesavanja.TipoviSadrzaja {
		tt := sadrzaj.TipSadrzajaPrikaz{
			Naziv:        t.Naziv,
			NazivMnozina: t.NazivMnozina,
			Slug:         t.Slug,
			SlugMnozina:  t.SlugMnozina,
			Struktura:    t.Struktura,
			Kategorije:   t.Kategorije,
			Link:         new(widget.Clickable),
		}
		tipovi = append(tipovi, tt)
		w.TipoviSadrzajaPrikaz = tipovi
	}
}
