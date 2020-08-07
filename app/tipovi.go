package cms

import (
	"gioui.org/widget"
	osnovni "github.com/w-ingsolutions/cms/DEF"
	"github.com/w-ingsolutions/cms/pkg/φ"
)

func CreateContentTypes() map[string]φ.T {
	return map[string]φ.T{
		"radovi":    osnovni.Radovi(),
		"materijal": osnovni.Materijal(),
	}
}

func displayTypes(tipoviSadrzaja map[string]φ.T) {
	var tipovi []φ.ContentType
	for _, t := range tipoviSadrzaja {
		tt := φ.ContentType{
			Title:       t.Title,
			TitlePlural: t.TitlePlural,
			Slug:        t.Slug,
			SlugPlural:  t.SlugPlural,
			Struct:      t.Struct,
			Categories:  t.Categories,
			Link:        new(widget.Clickable),
		}
		tipovi = append(tipovi, tt)
		tipoviSadrzajaPrikaz = tipovi
	}
}
