package osnovna

import (
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func Materijal() sadrzaj.TipSadrzaja {
	return sadrzaj.TipSadrzaja{
		Naziv:        "Materijal",
		NazivMnozina: "Materijali",
		Slug:         "materijal",
		SlugMnozina:  "materijali",
		Struktura: map[string]sadrzaj.PoljeSadrzaja{
			"Naziv":             sadrzaj.PoljeSadrzaja{"Naziv", "Text", ""},
			"Opis":              sadrzaj.PoljeSadrzaja{"Opis", "Text", ""},
			"Obracun":           sadrzaj.PoljeSadrzaja{"Obracun", "Text", ""},
			"Proizvodjac":       sadrzaj.PoljeSadrzaja{"Proizvodjac", "Text", ""},
			"OsobineNamena":     sadrzaj.PoljeSadrzaja{"OsobineNamena", "Text", ""},
			"NacinRada":         sadrzaj.PoljeSadrzaja{"NacinRada", "Text", ""},
			"JedinicaPotrosnje": sadrzaj.PoljeSadrzaja{"JedinicaPotrosnje", "Text", ""},
			"Potrosnja":         sadrzaj.PoljeSadrzaja{"Potrosnja", "Num", ""},
			"RokUpotrebe":       sadrzaj.PoljeSadrzaja{"RokUpotrebe", "Text", ""},
			"Jedinica":          sadrzaj.PoljeSadrzaja{"Jedinica", "Text", ""},
			"Pakovanje":         sadrzaj.PoljeSadrzaja{"Pakovanje", "Num", ""},
			"Cena":              sadrzaj.PoljeSadrzaja{"Cena", "Num", ""},
			"Slug":              sadrzaj.PoljeSadrzaja{"Slug", "Text", ""},
		},
		//Link:         new(widget.Clickable),
	}
}
