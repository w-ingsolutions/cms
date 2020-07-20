package cms

import (
	"gioui.org/widget"
	"github.com/w-ingsolutions/c/model"
)

func tipoviSadrzaja() map[int]model.TipSadrzaja {
	return map[int]model.TipSadrzaja{
		0: model.TipSadrzaja{
			Naziv:        "Rad",
			NazivMnozina: "Radovi",
			Slug:         "rad",
			SlugMnozina:  "radovi",
			Link:         new(widget.Clickable),
		},
		1: model.TipSadrzaja{
			Naziv:        "Materijal",
			NazivMnozina: "Materijali",
			Slug:         "materijal",
			SlugMnozina:  "materijali",
			Link:         new(widget.Clickable),
		},
		2: model.TipSadrzaja{
			Naziv:        "Pravno lice",
			NazivMnozina: "Pravna lica",
			Slug:         "pravno_lice",
			SlugMnozina:  "pravna_lice",
			Link:         new(widget.Clickable),
		},
		3: model.TipSadrzaja{
			Naziv:        "Fizicko lice",
			NazivMnozina: "Fizicka lica",
			Slug:         "fizicko_lice",
			SlugMnozina:  "fizicka_lica",
			Link:         new(widget.Clickable),
		},
		4: model.TipSadrzaja{
			Naziv:        "Objekat",
			NazivMnozina: "Objekti",
			Slug:         "objekat",
			SlugMnozina:  "objekti",
			Link:         new(widget.Clickable),
		},
		5: model.TipSadrzaja{
			Naziv:        "Vrsta objekta",
			NazivMnozina: "Vrste objekta",
			Slug:         "vrsta_objekta",
			SlugMnozina:  "vrste_objekta",
			Link:         new(widget.Clickable),
		},
	}
}
