package cms

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

type poljeSadrzaja struct {
	Naziv string
	Tip   string
}

func (w *WingCMS) noviTip() func() {
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
			"dodajdugme":    new(widget.Clickable),
			"dodajtipdugme": new(widget.Clickable),
			"vrstesadrzaja": new(widget.Enum),
			"struktura": &layout.List{
				Axis: layout.Vertical,
			},
		}

		w.Prikaz.w = widgets

		naziv := w.Prikaz.w["naziv"].(*widget.Editor)
		nazivmnozina := w.Prikaz.w["nazivmnozina"].(*widget.Editor)
		slug := w.Prikaz.w["slug"].(*widget.Editor)
		slugmnozina := w.Prikaz.w["slugmnozina"].(*widget.Editor)
		vrste := w.Prikaz.w["vrstesadrzaja"].(*widget.Enum)
		nazivTipaSadrzaja := w.Prikaz.w["nazivTipaSadrzaja"].(*widget.Editor)
		strukturaLista := w.Prikaz.w["struktura"].(*layout.List)

		struktura := []poljeSadrzaja{}

		w.Prikaz.e = []func(gtx C) D{
			Editor(w.UI.Tema, naziv, layout.Vertical, "Naziv", func(e widget.EditorEvent) {}),
			Editor(w.UI.Tema, nazivmnozina, layout.Vertical, "Naziv mnozina", func(e widget.EditorEvent) {}),
			Editor(w.UI.Tema, slug, layout.Vertical, "Slug", func(e widget.EditorEvent) {}),
			Editor(w.UI.Tema, slugmnozina, layout.Vertical, "Slug mnozina", func(e widget.EditorEvent) {}),
			func(gtx C) D {
				return strukturaLista.Layout(gtx, len(struktura), func(gtx C, i int) D {
					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
						func(gtx C) D {
							return lyt.Format(gtx, "hflexb(middle,r(_),r(_))",
								material.Body1(w.UI.Tema.T, struktura[i].Naziv).Layout,
								material.Body1(w.UI.Tema.T, struktura[i].Tip).Layout,
							)
						},
						helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]),
					)
				})
			},
			func(gtx C) D {
				d := w.Prikaz.w["dodajtipdugme"].(*widget.Clickable)
				return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_),r(_),r(_),r(_),r(_))",
					Editor(w.UI.Tema, nazivTipaSadrzaja, layout.Horizontal, "Naziv tipa sadrzaja", func(e widget.EditorEvent) {}),
					material.RadioButton(w.UI.Tema.T, vrste, "label", "Label").Layout,
					material.RadioButton(w.UI.Tema.T, vrste, "string", "String").Layout,
					material.RadioButton(w.UI.Tema.T, vrste, "integer", "Integer").Layout,
					material.RadioButton(w.UI.Tema.T, vrste, "float", "Float").Layout,
					material.RadioButton(w.UI.Tema.T, vrste, "bool", "Bool").Layout,
					w.iconLink(d, "counterPlusIcon", func() {
						struktura = append(struktura, poljeSadrzaja{
							Naziv: nazivTipaSadrzaja.Text(),
							Tip:   vrste.Value,
						})
					}),
				)
			},
			func(gtx C) D {
				d := w.Prikaz.w["dodajdugme"].(*widget.Clickable)
				btn := material.Button(w.UI.Tema.T, d, "Dodaj")
				btn.CornerRadius = unit.Dp(0)
				btn.TextSize = unit.Dp(12)
				//btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
				btn.Background = helper.HexARGB(w.UI.Tema.Colors["Secondary"])
				for d.Clicked() {
					noviTip := model.TipSadrzaja{
						Struktura: new(model.WingObjekat),
						Link:      new(widget.Clickable),
					}
					noviTip.Naziv = nazivmnozina.Text()
					noviTip.NazivMnozina = nazivmnozina.Text()
					noviTip.Slug = slug.Text()
					noviTip.SlugMnozina = slugmnozina.Text()

					w.TipoviSadrzaja[noviTip.Slug] = noviTip
					w.tipoviSadrzajaPrikaz()
					fmt.Println("ddd", noviTip)
				}
				return btn.Layout(gtx)
			},
		}
	}
}
