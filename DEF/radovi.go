package osnovna

import (
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func RadoviKategorije() map[string]sadrzaj.Kategorija {
	return map[string]sadrzaj.Kategorija{
		"pripremni":                sadrzaj.Kategorija{1, "Pripremni", "Pripremni", "pripremni", ""},
		"istrazivacki":             sadrzaj.Kategorija{2, "Istraživački", "Istraživački", "istrazivacki", ""},
		"rusenja":                  sadrzaj.Kategorija{3, "Rušenja", "Rušenja", "rusenja", ""},
		"zemljani":                 sadrzaj.Kategorija{4, "Zemljani", "Zemljani", "zemljani", ""},
		"zidarski":                 sadrzaj.Kategorija{5, "Zidarski", "Zidarski", "zidarski", ""},
		"betonski":                 sadrzaj.Kategorija{6, "Betonski", "Betonski", "betonski", ""},
		"tesarski":                 sadrzaj.Kategorija{7, "Tesarski", "Tesarski", "tesarski", ""},
		"pokrivacki":               sadrzaj.Kategorija{8, "Pokrivački", "Pokrivački", "pokrivacki", ""},
		"izolaterski":              sadrzaj.Kategorija{9, "Izolaterski", "Izolaterski", "izolaterski", ""},
		"stolarija":                sadrzaj.Kategorija{10, "Stolarija", "Stolarija", "stolarija", ""},
		"stolarski":                sadrzaj.Kategorija{11, "Stolarski", "Stolarski", "stolarski", ""},
		"bravarski":                sadrzaj.Kategorija{12, "Bravarski", "Bravarski", "bravarski", ""},
		"limarski":                 sadrzaj.Kategorija{13, "Limarski", "Limarski", "limarski", ""},
		"staklorezacki":            sadrzaj.Kategorija{14, "Staklorezački", "Staklorezački", "staklorezacki", ""},
		"keramicarski":             sadrzaj.Kategorija{15, "Keramičarski", "Keramičarski", "keramicarski", ""},
		"teracerski":               sadrzaj.Kategorija{16, "Teracerski", "Teracerski", "teracerski", ""},
		"kamenorezacki":            sadrzaj.Kategorija{17, "Kamenorezački", "Kamenorezački", "kamenorezacki", ""},
		"parketarski":              sadrzaj.Kategorija{18, "Parketarski", "Parketarski", "parketarski", ""},
		"podopolagacki":            sadrzaj.Kategorija{19, "Podopolagački", "Podopolagački", "podopolagacki", ""},
		"tapetarski":               sadrzaj.Kategorija{20, "Tapetarski", "Tapetarski", "tapetarski", ""},
		"roletnarski":              sadrzaj.Kategorija{21, "Roletnarski", "Roletnarski", "roletnarski", ""},
		"suvomontazni":             sadrzaj.Kategorija{22, "Suvomontažni", "Suvomontažni", "suvomontazni", ""},
		"gipsarski":                sadrzaj.Kategorija{23, "Gipsarski", "Gipsarski", "gipsarski", ""},
		"fasaderski":               sadrzaj.Kategorija{24, "Fasaderski", "Fasaderski", "fasaderski", ""},
		"likorezacki":              sadrzaj.Kategorija{25, "Likorezački", "Likorezački", "likorezacki", ""},
		"molerski":                 sadrzaj.Kategorija{26, "Molerski", "Molerski", "molerski", ""},
		"molerski_pripremni":       sadrzaj.Kategorija{1, "Pripremni", "molerski_pripremni", "molerski_pripremni", "molerski"},
		"molerski_vrata_i_prozori": sadrzaj.Kategorija{2, "Vrata i prozori", "molerski_vrata_i_prozori", "molerski_vrata_i_prozori", "molerski"},
		"molerski_zastita_fasada":  sadrzaj.Kategorija{3, "Zaštita fasada", "molerski_zastita_fasada", "molerski_zastita_fasada", "molerski"},
		"molerski_limarija":        sadrzaj.Kategorija{4, "Limarija", "molerski_limarija", "molerski_limarija", "molerski"},
		"molerski_ograde":          sadrzaj.Kategorija{5, "Ograde", "molerski_ograde", "molerski_ograde", "molerski"},
		"molerski_instalacije":     sadrzaj.Kategorija{6, "Instalacije", "molerski_instalacije", "molerski_instalacije", "molerski"},
		"molerski_mobilijar":       sadrzaj.Kategorija{7, "Mobilijar", "molerski_mobilijar", "molerski_mobilijar", "molerski"},
		"molerski_ostali_molersko_farbarski_radovi": sadrzaj.Kategorija{8, "Ostali molersko - farbarski radovi", "molerski_ostali_molersko_farbarski_radovi", "molerski_ostali_molersko_farbarski_radovi", "molerski"},
		"molerski_zidovi_i_plafoni":                 sadrzaj.Kategorija{9, "Zidovi i plafoni", "molerski_zidovi_i_plafoni", "molerski_zidovi_i_plafoni", "molerski"},
		"molerski_tapete":                           sadrzaj.Kategorija{10, "Tapete", "molerski_tapete", "molerski_tapete", "molerski"},
		"molerski_lamperije":                        sadrzaj.Kategorija{11, "Lamperije", "molerski_lamperije", "molerski_lamperije", "molerski"},
		"molerski_strehe":                           sadrzaj.Kategorija{12, "Strehe", "molerski_strehe", "molerski_strehe", "molerski"},
		"molerski_stepenista":                       sadrzaj.Kategorija{13, "Stepeništa", "molerski_stepenista", "molerski_stepenista", "molerski"},
		"molerski_podovi":                           sadrzaj.Kategorija{14, "Podovi", "molerski_podovi", "molerski_podovi", "molerski"},
		"molerski_obrada_fasada":                    sadrzaj.Kategorija{15, "Obrada fasada", "molerski_obrada_fasada", "molerski_obrada_fasada", "molerski"},
		"molerski_bojenja_fasada":                   sadrzaj.Kategorija{16, "Bojenja fasada", "molerski_bojenja_fasada", "molerski_bojenja_fasada", "molerski"},
		"livacki":                                   sadrzaj.Kategorija{27, "Livački", "Livački", "livacki", ""},
		"razni":                                     sadrzaj.Kategorija{28, "Razni", "Razni", "razni", ""},
		"vodovod":                                   sadrzaj.Kategorija{29, "Vodovod", "Vodovod", "vodovod", ""},
		"kanalizacija":                              sadrzaj.Kategorija{30, "Kanalizacija", "Kanalizacija", "kanalizacija", ""},
		"sanitarije":                                sadrzaj.Kategorija{31, "Sanitarije", "Sanitarije", "sanitarije", ""},
	}
}

func Radovi() sadrzaj.TipSadrzaja {
	return sadrzaj.TipSadrzaja{
		Naziv:        "Rad",
		NazivMnozina: "Radovi",
		Slug:         "rad",
		SlugMnozina:  "radovi",
		Struktura: map[string]sadrzaj.PoljeSadrzaja{
			"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", ""},
			"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", ""},
			"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", ""},
			"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", ""},
			"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", ""},
			"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", ""},
			//"Neophodan": sadrzaj.PoljeSadrzaja{"Neophodan", "Array", ""},
		},
		Kategorije: RadoviKategorije(),
	}
}
