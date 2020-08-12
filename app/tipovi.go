package cms

import (
	"gioui.org/widget"
	osnovni "github.com/w-ingsolutions/cms/DEF"
	"github.com/w-ingsolutions/cms/pkg/phi"
)

func CreateContentTypes() map[string]phi.T {
	return map[string]phi.T{
		"radovi":    osnovni.Radovi(),
		"materijal": osnovni.Materijal(),
	}
}

func displayTypes(tipoviSadrzaja map[string]phi.T) {
	var tipovi []phi.ContentType
	for _, t := range tipoviSadrzaja {
		tt := phi.ContentType{
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
