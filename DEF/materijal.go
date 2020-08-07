package osnovna

import (
	"github.com/w-ingsolutions/cms/pkg/content"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func Materijal() content.Type {
	return content.Type{
		Title:       "Materijal",
		TitlePlural: "Materijali",
		Slug:        "materijal",
		SlugPlural:  "materijali",
		Struct: map[string]content.Field{
			"Title":             content.Field{"Title", "Text", ""},
			"Opis":              content.Field{"Opis", "Text", ""},
			"Obracun":           content.Field{"Obracun", "Text", ""},
			"Proizvodjac":       content.Field{"Proizvodjac", "Text", ""},
			"OsobineNamena":     content.Field{"OsobineNamena", "Text", ""},
			"NacinRada":         content.Field{"NacinRada", "Text", ""},
			"JedinicaPotrosnje": content.Field{"JedinicaPotrosnje", "Text", ""},
			"Potrosnja":         content.Field{"Potrosnja", "Num", ""},
			"RokUpotrebe":       content.Field{"RokUpotrebe", "Text", ""},
			"Jedinica":          content.Field{"Jedinica", "Text", ""},
			"Pakovanje":         content.Field{"Pakovanje", "Num", ""},
			"Cena":              content.Field{"Cena", "Num", ""},
			"Slug":              content.Field{"Slug", "Text", ""},
		},
		//Link:         new(widget.Clickable),
	}
}
