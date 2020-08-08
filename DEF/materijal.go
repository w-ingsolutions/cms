package osnovna

import (
	"github.com/w-ingsolutions/cms/pkg/φ"
)

func Materijal() phi.T {
	return phi.T{
		Title:       "Materijal",
		TitlePlural: "Materijali",
		Slug:        "materijal",
		SlugPlural:  "materijali",
		Struct: map[string]phi.F{
			"Id":                phi.F{"Id", "Text", ""},
			"Title":             phi.F{"Title", "Text", ""},
			"Opis":              phi.F{"Opis", "Text", ""},
			"Obracun":           phi.F{"Obracun", "Text", ""},
			"Proizvodjac":       phi.F{"Proizvodjac", "Text", ""},
			"OsobineNamena":     phi.F{"OsobineNamena", "Text", ""},
			"NacinRada":         phi.F{"NacinRada", "Text", ""},
			"JedinicaPotrosnje": phi.F{"JedinicaPotrosnje", "Text", ""},
			"Potrosnja":         phi.F{"Potrosnja", "Num", ""},
			"RokUpotrebe":       phi.F{"RokUpotrebe", "Text", ""},
			"Jedinica":          phi.F{"Jedinica", "Text", ""},
			"Pakovanje":         phi.F{"Pakovanje", "Num", ""},
			"Cena":              phi.F{"Cena", "Num", ""},
			"Slug":              phi.F{"Slug", "Text", ""},
		},
		//Link:         new(widget.Clickable),
	}
}
