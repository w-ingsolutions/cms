package cms

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/theme"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
	"github.com/w-ingsolutions/cms/pkg/utl"
)

func (w *WingCMS) podesavanjaTipa(tip content.TypePrikaz) func() {
	return func() {
		widgets := map[string]interface{}{
			"naziv": &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			"nazivmnozina": &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			"slug": &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			"slugmnozina": &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			"nazivTipaSadrzaja": &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			"dodajdugme":         new(widget.Clickable),
			"dodajtipdugme":      new(widget.Clickable),
			"vrstepoljasadrzaja": new(widget.Enum),
			"struktura": &layout.List{
				Axis: layout.Vertical,
			},
			"vrstepolja": &layout.List{
				Axis: layout.Horizontal,
			},
		}

		w.Prikaz.w = widgets

		naziv := w.Prikaz.w["naziv"].(*widget.Editor)
		nazivmnozina := w.Prikaz.w["nazivmnozina"].(*widget.Editor)
		slug := w.Prikaz.w["slug"].(*widget.Editor)
		slugmnozina := w.Prikaz.w["slugmnozina"].(*widget.Editor)
		vrste := w.Prikaz.w["vrstepoljasadrzaja"].(*widget.Enum)
		nazivTipaSadrzaja := w.Prikaz.w["nazivTipaSadrzaja"].(*widget.Editor)
		strukturaLista := w.Prikaz.w["struktura"].(*layout.List)
		vrstePolja := w.Prikaz.w["vrstepolja"].(*layout.List)

		struktura := make(map[string]content.Field)
		if tip.Struct != nil {
			struktura = tip.Struct
		}

		if tip.Title != "" {
			naziv.SetText(tip.Title)
		}
		if tip.TitlePlural != "" {
			nazivmnozina.SetText(tip.TitlePlural)
		}
		if tip.Slug != "" {
			slug.SetText(tip.Slug)
		}
		if tip.SlugPlural != "" {
			slugmnozina.SetText(tip.SlugPlural)
		}

		w.Prikaz.e = []func(gtx C) D{
			utl.Editor(w.UI.Tema, naziv, false, "Title", func(e widget.EditorEvent) {}),
			utl.Editor(w.UI.Tema, nazivmnozina, false, "Title mnozina", func(e widget.EditorEvent) {}),
			utl.Editor(w.UI.Tema, slug, false, "Slug", func(e widget.EditorEvent) {}),
			utl.Editor(w.UI.Tema, slugmnozina, false, "Slug mnozina", func(e widget.EditorEvent) {}),
			strukturaListaIzgled(w.UI.Tema, struktura, strukturaLista),
			w.dodajtipdugme(struktura, nazivTipaSadrzaja, vrstePolja, vrste),
			w.dodajdugme(struktura),
		}
	}
}

func vrstePoljaSadrzaja(th *material.Theme, l *layout.List, enum *widget.Enum) func(gtx C) D {
	v := []string{
		"Text",
		"Num",
		"Array",
	}
	return func(gtx C) D {
		return l.Layout(gtx, len(v), func(gtx C, i int) D {
			return material.RadioButton(th, enum, v[i], v[i]).Layout(gtx)
		})
	}
}

func strukturaListaIzgled(th *theme.DuoUItheme, struktura map[string]content.Field, list *layout.List) func(gtx C) D {
	return func(gtx C) D {
		var strArray []content.Field
		for _, str := range struktura {
			strArray = append(strArray, content.Field{
				Title: str.Title,
				Tip:   str.Tip,
			})
		}
		return list.Layout(gtx, len(strArray), func(gtx C, i int) D {
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				func(gtx C) D {
					return lyt.Format(gtx, "hflexb(middle,r(_),r(_))",
						material.Body1(th.T, strArray[i].Title).Layout,
						material.Body1(th.T, strArray[i].Tip).Layout,
					)
				},
				helper.DuoUIline(false, 0, 0, 1, th.Colors["Gray"]),
			)
		})
	}
}

func (w *WingCMS) dodajtipdugme(struktura map[string]content.Field, nazivTipaSadrzaja *widget.Editor, vrstePolja *layout.List, vrste *widget.Enum) func(gtx C) D {
	return func(gtx C) D {
		d := w.Prikaz.w["dodajtipdugme"].(*widget.Clickable)
		return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_))",
			utl.Editor(w.UI.Tema, nazivTipaSadrzaja, false, "Title tipa sadrzaja", func(e widget.EditorEvent) {}),
			vrstePoljaSadrzaja(w.UI.Tema.T, vrstePolja, vrste),
			w.iconLink(d, "counterPlusIcon", func() {
				struktura[nazivTipaSadrzaja.Text()] = content.Field{
					Title: nazivTipaSadrzaja.Text(),
					Tip:   vrste.Value,
				}
			}))
	}
}

func (w *WingCMS) dodajdugme(struktura map[string]content.Field) func(gtx C) D {
	return func(gtx C) D {
		d := w.Prikaz.w["dodajdugme"].(*widget.Clickable)
		btn := material.Button(w.UI.Tema.T, d, "Dodaj")
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(12)
		//btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["Secondary"])
		for d.Clicked() {
			tipSadrzaja := content.Type{
				Title:       w.Prikaz.w["naziv"].(*widget.Editor).Text(),
				TitlePlural: w.Prikaz.w["nazivmnozina"].(*widget.Editor).Text(),
				Slug:        w.Prikaz.w["slug"].(*widget.Editor).Text(),
				SlugPlural:  w.Prikaz.w["slugmnozina"].(*widget.Editor).Text(),
				Struct:      struktura,
			}
			w.Podesavanja.TipoviSadrzaja[w.Prikaz.w["slug"].(*widget.Editor).Text()] = tipSadrzaja
			w.tipoviSadrzajaPrikaz()
			var write bytes.Buffer
			enc := gob.NewEncoder(&write)
			err := enc.Encode(w.Podesavanja)
			checkError(err)
			err = w.sh.FilesWrite(w.ctx, "/wing/config", &write, shell.FilesWrite.Create(true))
			fmt.Println("PodesavanjaSnima")
		}
		return btn.Layout(gtx)
	}
}
