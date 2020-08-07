package cms

import (
	"gioui.org/widget"
	osnovni "github.com/w-ingsolutions/cms/DEF"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func tipoviSadrzaja() map[string]content.Type {
	return map[string]content.Type{
		"radovi":    osnovni.Radovi(),
		"materijal": osnovni.Materijal(),
	}
}

func (w *WingCMS) tipoviSadrzajaPrikaz() {
	var tipovi []content.TypePrikaz
	for _, t := range w.Podesavanja.TipoviSadrzaja {
		tt := content.TypePrikaz{
			Title:       t.Title,
			TitlePlural: t.TitlePlural,
			Slug:        t.Slug,
			SlugPlural:  t.SlugPlural,
			Struct:      t.Struct,
			Kategorije:  t.Kategorije,
			Link:        new(widget.Clickable),
		}
		tipovi = append(tipovi, tt)
		w.TipoviSadrzajaPrikaz = tipovi
	}
}
