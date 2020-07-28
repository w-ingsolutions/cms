package cms

import (
	"gioui.org/widget"
)

func tipoviSadrzaja() map[string]TipSadrzaja {
	return map[string]TipSadrzaja{
		"radovi": TipSadrzaja{
			Naziv:        "Rad",
			NazivMnozina: "Radovi",
			Slug:         "rad",
			SlugMnozina:  "radovi",
			Struktura: []poljeSadrzaja{
				poljeSadrzaja{
					Naziv: "Naziv",
					Tip:   "string",
				},
				poljeSadrzaja{
					Naziv: "Opis",
					Tip:   "string",
				},
				poljeSadrzaja{
					Naziv: "Cena",
					Tip:   "float",
				},
			},
			//Link:         new(widget.Clickable),
		},
		"materijal": TipSadrzaja{
			Naziv:        "Materijal",
			NazivMnozina: "Materijali",
			Slug:         "materijal",
			SlugMnozina:  "materijali",
			Struktura: []poljeSadrzaja{
				poljeSadrzaja{
					Naziv: "Naziv",
					Tip:   "string",
				},
				poljeSadrzaja{
					Naziv: "Opis",
					Tip:   "string",
				},
				poljeSadrzaja{
					Naziv: "Cena",
					Tip:   "float",
				},
			},
			//Link:         new(widget.Clickable),
		},
	}
}

func (w *WingCMS) tipoviSadrzajaPrikaz() {
	var tipovi []TipSadrzajaPrikaz
	for _, t := range w.Podesavanja.TipoviSadrzaja {
		tt := TipSadrzajaPrikaz{
			Naziv:        t.Naziv,
			NazivMnozina: t.NazivMnozina,
			Slug:         t.Slug,
			SlugMnozina:  t.SlugMnozina,
			Struktura:    t.Struktura,
			Link:         new(widget.Clickable),
		}
		tipovi = append(tipovi, tt)
		w.TipoviSadrzajaPrikaz = tipovi
	}
}
