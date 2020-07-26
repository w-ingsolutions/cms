package cms

import (
	"gioui.org/widget"
	"github.com/w-ingsolutions/c/model"
)

func tipoviSadrzaja() map[string]model.TipSadrzaja {
	return map[string]model.TipSadrzaja{
		"radovi": model.TipSadrzaja{
			Naziv:        "Rad",
			NazivMnozina: "Radovi",
			Slug:         "rad",
			SlugMnozina:  "radovi",
			Struktura:    new(model.WingVrstaRadova),
			Link:         new(widget.Clickable),
		},
		"materijal": model.TipSadrzaja{
			Naziv:        "Materijal",
			NazivMnozina: "Materijali",
			Slug:         "materijal",
			SlugMnozina:  "materijali",
			Struktura:    new(model.WingMaterijal),
			Link:         new(widget.Clickable),
		},
		"pravno_lice": model.TipSadrzaja{
			Naziv:        "Pravno lice",
			NazivMnozina: "Pravna lica",
			Slug:         "pravno_lice",
			SlugMnozina:  "pravna_lica",
			Struktura:    new(model.WingPravnoLice),
			Link:         new(widget.Clickable),
		},
		"fizicko_lice": model.TipSadrzaja{
			Naziv:        "Fizicko lice",
			NazivMnozina: "Fizicka lica",
			Slug:         "fizicko_lice",
			SlugMnozina:  "fizicka_lica",
			Struktura:    new(model.WingFizickoLice),
			Link:         new(widget.Clickable),
		},
		"objekat": model.TipSadrzaja{
			Naziv:        "Objekat",
			NazivMnozina: "Objekti",
			Slug:         "objekat",
			SlugMnozina:  "objekti",
			Struktura:    new(model.WingObjekat),
			Link:         new(widget.Clickable),
		},
		"vrsta_objekta": model.TipSadrzaja{
			Naziv:        "Vrsta objekta",
			NazivMnozina: "Vrste objekta",
			Slug:         "vrsta_objekta",
			SlugMnozina:  "vrste_objekta",
			Struktura:    new(model.WingObjekat),
			Link:         new(widget.Clickable),
		},
	}
}

func (w *WingCMS) tipoviSadrzajaPrikaz() {
	var tipovi []model.TipSadrzaja
	for _, t := range w.TipoviSadrzaja {
		tipovi = append(tipovi, t)
		w.TipoviSadrzajaPrikaz = tipovi
	}
}
