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
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func (w *WingCMS) podesavanjaTipa(tip sadrzaj.TipSadrzajaPrikaz) func() {
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

		struktura := make(map[string]sadrzaj.PoljeSadrzaja)
		if tip.Struktura != nil {
			struktura = tip.Struktura
		}

		if tip.Naziv != "" {
			naziv.SetText(tip.Naziv)
		}
		if tip.NazivMnozina != "" {
			nazivmnozina.SetText(tip.NazivMnozina)
		}
		if tip.Slug != "" {
			slug.SetText(tip.Slug)
		}
		if tip.SlugMnozina != "" {
			slugmnozina.SetText(tip.SlugMnozina)
		}

		w.Prikaz.e = []func(gtx C) D{
			Editor(w.UI.Tema, naziv, layout.Vertical, "Naziv", func(e widget.EditorEvent) {}),
			Editor(w.UI.Tema, nazivmnozina, layout.Vertical, "Naziv mnozina", func(e widget.EditorEvent) {}),
			Editor(w.UI.Tema, slug, layout.Vertical, "Slug", func(e widget.EditorEvent) {}),
			Editor(w.UI.Tema, slugmnozina, layout.Vertical, "Slug mnozina", func(e widget.EditorEvent) {}),
			func(gtx C) D {
				var strArray []sadrzaj.PoljeSadrzaja
				for _, str := range struktura {
					strArray = append(strArray, sadrzaj.PoljeSadrzaja{
						Naziv: str.Naziv,
						Tip:   str.Tip,
					})
				}
				return strukturaLista.Layout(gtx, len(strArray), func(gtx C, i int) D {
					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
						func(gtx C) D {
							return lyt.Format(gtx, "hflexb(middle,r(_),r(_))",
								material.Body1(w.UI.Tema.T, strArray[i].Naziv).Layout,
								material.Body1(w.UI.Tema.T, strArray[i].Tip).Layout,
							)
						},
						helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]),
					)
				})
			},
			func(gtx C) D {
				d := w.Prikaz.w["dodajtipdugme"].(*widget.Clickable)
				return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_))",
					Editor(w.UI.Tema, nazivTipaSadrzaja, layout.Horizontal, "Naziv tipa sadrzaja", func(e widget.EditorEvent) {}),
					vrstePoljaSadrzaja(w.UI.Tema.T, vrstePolja, vrste),
					w.iconLink(d, "counterPlusIcon", func() {
						struktura[nazivTipaSadrzaja.Text()] = sadrzaj.PoljeSadrzaja{
							Naziv: nazivTipaSadrzaja.Text(),
							Tip:   vrste.Value,
						}
					}))
			},
			func(gtx C) D {
				d := w.Prikaz.w["dodajdugme"].(*widget.Clickable)
				btn := material.Button(w.UI.Tema.T, d, "Dodaj")
				btn.CornerRadius = unit.Dp(0)
				btn.TextSize = unit.Dp(12)
				//btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
				btn.Background = helper.HexARGB(w.UI.Tema.Colors["Secondary"])
				for d.Clicked() {
					tipSadrzaja := sadrzaj.TipSadrzaja{
						Naziv:        naziv.Text(),
						NazivMnozina: nazivmnozina.Text(),
						Slug:         slug.Text(),
						SlugMnozina:  slugmnozina.Text(),
						Struktura:    struktura,
					}
					w.Podesavanja.TipoviSadrzaja[slug.Text()] = tipSadrzaja

					w.tipoviSadrzajaPrikaz()

					var write bytes.Buffer
					//var read bytes.Buffer
					enc := gob.NewEncoder(&write) // Will write to network.
					// Encode (send) the value.
					err := enc.Encode(w.Podesavanja)
					checkError(err)
					err = w.sh.FilesWrite(w.ctx, "/wing/config", &write, shell.FilesWrite.Create(true))
					fmt.Println("PodesavanjaSnima")

				}
				return btn.Layout(gtx)
			},
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
