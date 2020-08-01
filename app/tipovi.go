package cms

import (
	"gioui.org/widget"
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func tipoviSadrzaja() map[string]sadrzaj.TipSadrzaja {
	return map[string]sadrzaj.TipSadrzaja{
		"radovi": sadrzaj.TipSadrzaja{
			Naziv:        "Rad",
			NazivMnozina: "Radovi",
			Slug:         "rad",
			SlugMnozina:  "radovi",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":     sadrzaj.PoljeSadrzaja{"Naziv", "Text", ""},
				"Opis":      sadrzaj.PoljeSadrzaja{"Opis", "Text", ""},
				"Obracun":   sadrzaj.PoljeSadrzaja{"Obracun", "Text", ""},
				"Jedinica":  sadrzaj.PoljeSadrzaja{"Jedinica", "Text", ""},
				"Cena":      sadrzaj.PoljeSadrzaja{"Cena", "Num", ""},
				"Slug":      sadrzaj.PoljeSadrzaja{"Slug", "Text", ""},
				"Neophodan": sadrzaj.PoljeSadrzaja{"Neophodan", "Array", ""},
			},
			Kategorije: RadoviKategorije(),
		},
		"materijal": sadrzaj.TipSadrzaja{
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
		},
	}
}

func (w *WingCMS) tipoviSadrzajaPrikaz() {
	var tipovi []sadrzaj.TipSadrzajaPrikaz
	for _, t := range w.Podesavanja.TipoviSadrzaja {
		tt := sadrzaj.TipSadrzajaPrikaz{
			Naziv:        t.Naziv,
			NazivMnozina: t.NazivMnozina,
			Slug:         t.Slug,
			SlugMnozina:  t.SlugMnozina,
			Struktura:    t.Struktura,
			Kategorije:   t.Kategorije,
			Link:         new(widget.Clickable),
		}
		tipovi = append(tipovi, tt)
		w.TipoviSadrzajaPrikaz = tipovi
	}
}

func RadoviKategorije() map[string]sadrzaj.Kategorija {
	return map[string]sadrzaj.Kategorija{
		"pripremni":                sadrzaj.Kategorija{"Pripremni", "Pripremni", "pripremni", ""},
		"istrazivacki":             sadrzaj.Kategorija{"Istraživački", "Istraživački", "istrazivacki", ""},
		"rusenja":                  sadrzaj.Kategorija{"Rušenja", "Rušenja", "rusenja", ""},
		"zemljani":                 sadrzaj.Kategorija{"Zemljani", "Zemljani", "zemljani", ""},
		"zidarski":                 sadrzaj.Kategorija{"Zidarski", "Zidarski", "zidarski", ""},
		"betonski":                 sadrzaj.Kategorija{"Betonski", "Betonski", "betonski", ""},
		"tesarski":                 sadrzaj.Kategorija{"Tesarski", "Tesarski", "tesarski", ""},
		"pokrivacki":               sadrzaj.Kategorija{"Pokrivački", "Pokrivački", "pokrivacki", ""},
		"izolaterski":              sadrzaj.Kategorija{"Izolaterski", "Izolaterski", "izolaterski", ""},
		"stolarija":                sadrzaj.Kategorija{"Stolarija", "Stolarija", "stolarija", ""},
		"stolarski":                sadrzaj.Kategorija{"Stolarski", "Stolarski", "stolarski", ""},
		"bravarski":                sadrzaj.Kategorija{"Bravarski", "Bravarski", "bravarski", ""},
		"limarski":                 sadrzaj.Kategorija{"Limarski", "Limarski", "limarski", ""},
		"staklorezacki":            sadrzaj.Kategorija{"Staklorezački", "Staklorezački", "staklorezacki", ""},
		"keramicarski":             sadrzaj.Kategorija{"Keramičarski", "Keramičarski", "keramicarski", ""},
		"teracerski":               sadrzaj.Kategorija{"Teracerski", "Teracerski", "teracerski", ""},
		"kamenorezacki":            sadrzaj.Kategorija{"Kamenorezački", "Kamenorezački", "kamenorezacki", ""},
		"parketarski":              sadrzaj.Kategorija{"Parketarski", "Parketarski", "parketarski", ""},
		"podopolagacki":            sadrzaj.Kategorija{"Podopolagački", "Podopolagački", "podopolagacki", ""},
		"tapetarski":               sadrzaj.Kategorija{"Tapetarski", "Tapetarski", "tapetarski", ""},
		"roletnarski":              sadrzaj.Kategorija{"Roletnarski", "Roletnarski", "roletnarski", ""},
		"suvomontazni":             sadrzaj.Kategorija{"Suvomontažni", "Suvomontažni", "suvomontazni", ""},
		"gipsarski":                sadrzaj.Kategorija{"Gipsarski", "Gipsarski", "gipsarski", ""},
		"fasaderski":               sadrzaj.Kategorija{"Fasaderski", "Fasaderski", "fasaderski", ""},
		"likorezacki":              sadrzaj.Kategorija{"Likorezački", "Likorezački", "likorezacki", ""},
		"molerski":                 sadrzaj.Kategorija{"Molerski", "Molerski", "molerski", ""},
		"molerski_pripremni":       sadrzaj.Kategorija{"Pripremni", "molerski_pripremni", "molerski_pripremni", "molerski"},
		"molerski_vrata_i_prozori": sadrzaj.Kategorija{"Vrata i prozori", "molerski_vrata_i_prozori", "molerski_vrata_i_prozori", "molerski"},
		"molerski_zastita_fasada":  sadrzaj.Kategorija{"Zaštita fasada", "molerski_zastita_fasada", "molerski_zastita_fasada", "molerski"},
		"molerski_limarija":        sadrzaj.Kategorija{"Limarija", "molerski_limarija", "molerski_limarija", "molerski"},
		"molerski_ograde":          sadrzaj.Kategorija{"Ograde", "molerski_ograde", "molerski_ograde", "molerski"},
		"molerski_instalacije":     sadrzaj.Kategorija{"Instalacije", "molerski_instalacije", "molerski_instalacije", "molerski"},
		"molerski_mobilijar":       sadrzaj.Kategorija{"Mobilijar", "molerski_mobilijar", "molerski_mobilijar", "molerski"},
		"molerski_ostali_molersko_farbarski_radovi": sadrzaj.Kategorija{"Ostali molersko - farbarski radovi", "molerski_ostali_molersko_farbarski_radovi", "molerski_ostali_molersko_farbarski_radovi", "molerski"},
		"molerski_zidovi_i_plafoni":                 sadrzaj.Kategorija{"Zidovi i plafoni", "molerski_zidovi_i_plafoni", "molerski_zidovi_i_plafoni", "molerski"},
		"molerski_tapete":                           sadrzaj.Kategorija{"Tapete", "molerski_tapete", "molerski_tapete", "molerski"},
		"molerski_lamperije":                        sadrzaj.Kategorija{"Lamperije", "molerski_lamperije", "molerski_lamperije", "molerski"},
		"molerski_strehe":                           sadrzaj.Kategorija{"Strehe", "molerski_strehe", "molerski_strehe", "molerski"},
		"molerski_stepenista":                       sadrzaj.Kategorija{"Stepeništa", "molerski_stepenista", "molerski_stepenista", "molerski"},
		"molerski_podovi":                           sadrzaj.Kategorija{"Podovi", "molerski_podovi", "molerski_podovi", "molerski"},
		"molerski_obrada_fasada":                    sadrzaj.Kategorija{"Obrada fasada", "molerski_obrada_fasada", "molerski_obrada_fasada", "molerski"},
		"molerski_bojenja_fasada":                   sadrzaj.Kategorija{"Bojenja fasada", "molerski_bojenja_fasada", "molerski_bojenja_fasada", "molerski"},

		"livacki":      sadrzaj.Kategorija{"Livački", "Livački", "livacki", ""},
		"razni":        sadrzaj.Kategorija{"Razni", "Razni", "razni", ""},
		"vodovod":      sadrzaj.Kategorija{"Vodovod", "Vodovod", "vodovod", ""},
		"kanalizacija": sadrzaj.Kategorija{"Kanalizacija", "Kanalizacija", "kanalizacija", ""},
		"sanitarije":   sadrzaj.Kategorija{"Sanitarije", "Sanitarije", "sanitarije", ""},
	}
}
