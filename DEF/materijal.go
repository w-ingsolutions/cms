package osnovna

import (
	"github.com/w-ingsolutions/cms/pkg/φ"
)

func Materijal() φ.T {
	return φ.T{
		Title:       "Materijal",
		TitlePlural: "Materijali",
		Slug:        "materijal",
		SlugPlural:  "materijali",
		Struct: map[string]φ.F{
			"Id":                φ.F{"Id", "Text", ""},
			"Title":             φ.F{"Title", "Text", ""},
			"Opis":              φ.F{"Opis", "Text", ""},
			"Obracun":           φ.F{"Obracun", "Text", ""},
			"Proizvodjac":       φ.F{"Proizvodjac", "Text", ""},
			"OsobineNamena":     φ.F{"OsobineNamena", "Text", ""},
			"NacinRada":         φ.F{"NacinRada", "Text", ""},
			"JedinicaPotrosnje": φ.F{"JedinicaPotrosnje", "Text", ""},
			"Potrosnja":         φ.F{"Potrosnja", "Num", ""},
			"RokUpotrebe":       φ.F{"RokUpotrebe", "Text", ""},
			"Jedinica":          φ.F{"Jedinica", "Text", ""},
			"Pakovanje":         φ.F{"Pakovanje", "Num", ""},
			"Cena":              φ.F{"Cena", "Num", ""},
			"Slug":              φ.F{"Slug", "Text", ""},
		},
		//Link:         new(widget.Clickable),
	}
}
