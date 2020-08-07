package osnovna

import (
	"github.com/w-ingsolutions/cms/pkg/φ"
)

func RadoviKategorije() map[string]φ.C {
	return map[string]φ.C{
		"pripremni":                φ.C{1, "Pripremni", "Pripremni", "pripremni", ""},
		"istrazivacki":             φ.C{2, "Istraživački", "Istraživački", "istrazivacki", ""},
		"rusenja":                  φ.C{3, "Rušenja", "Rušenja", "rusenja", ""},
		"zemljani":                 φ.C{4, "Zemljani", "Zemljani", "zemljani", ""},
		"zidarski":                 φ.C{5, "Zidarski", "Zidarski", "zidarski", ""},
		"betonski":                 φ.C{6, "Betonski", "Betonski", "betonski", ""},
		"tesarski":                 φ.C{7, "Tesarski", "Tesarski", "tesarski", ""},
		"pokrivacki":               φ.C{8, "Pokrivački", "Pokrivački", "pokrivacki", ""},
		"izolaterski":              φ.C{9, "Izolaterski", "Izolaterski", "izolaterski", ""},
		"stolarija":                φ.C{10, "Stolarija", "Stolarija", "stolarija", ""},
		"stolarski":                φ.C{11, "Stolarski", "Stolarski", "stolarski", ""},
		"bravarski":                φ.C{12, "Bravarski", "Bravarski", "bravarski", ""},
		"limarski":                 φ.C{13, "Limarski", "Limarski", "limarski", ""},
		"staklorezacki":            φ.C{14, "Staklorezački", "Staklorezački", "staklorezacki", ""},
		"keramicarski":             φ.C{15, "Keramičarski", "Keramičarski", "keramicarski", ""},
		"teracerski":               φ.C{16, "Teracerski", "Teracerski", "teracerski", ""},
		"kamenorezacki":            φ.C{17, "Kamenorezački", "Kamenorezački", "kamenorezacki", ""},
		"parketarski":              φ.C{18, "Parketarski", "Parketarski", "parketarski", ""},
		"podopolagacki":            φ.C{19, "Podopolagački", "Podopolagački", "podopolagacki", ""},
		"tapetarski":               φ.C{20, "Tapetarski", "Tapetarski", "tapetarski", ""},
		"roletnarski":              φ.C{21, "Roletnarski", "Roletnarski", "roletnarski", ""},
		"suvomontazni":             φ.C{22, "Suvomontažni", "Suvomontažni", "suvomontazni", ""},
		"gipsarski":                φ.C{23, "Gipsarski", "Gipsarski", "gipsarski", ""},
		"fasaderski":               φ.C{24, "Fasaderski", "Fasaderski", "fasaderski", ""},
		"likorezacki":              φ.C{25, "Likorezački", "Likorezački", "likorezacki", ""},
		"molerski":                 φ.C{26, "Molerski", "Molerski", "molerski", ""},
		"molerski_pripremni":       φ.C{1, "Pripremni", "molerski_pripremni", "molerski_pripremni", "molerski"},
		"molerski_vrata_i_prozori": φ.C{2, "Vrata i prozori", "molerski_vrata_i_prozori", "molerski_vrata_i_prozori", "molerski"},
		"molerski_zastita_fasada":  φ.C{3, "Zaštita fasada", "molerski_zastita_fasada", "molerski_zastita_fasada", "molerski"},
		"molerski_limarija":        φ.C{4, "Limarija", "molerski_limarija", "molerski_limarija", "molerski"},
		"molerski_ograde":          φ.C{5, "Ograde", "molerski_ograde", "molerski_ograde", "molerski"},
		"molerski_instalacije":     φ.C{6, "Instalacije", "molerski_instalacije", "molerski_instalacije", "molerski"},
		"molerski_mobilijar":       φ.C{7, "Mobilijar", "molerski_mobilijar", "molerski_mobilijar", "molerski"},
		"molerski_ostali_molersko_farbarski_radovi": φ.C{8, "Ostali molersko - farbarski radovi", "molerski_ostali_molersko_farbarski_radovi", "molerski_ostali_molersko_farbarski_radovi", "molerski"},
		"molerski_zidovi_i_plafoni":                 φ.C{9, "Zidovi i plafoni", "molerski_zidovi_i_plafoni", "molerski_zidovi_i_plafoni", "molerski"},
		"molerski_tapete":                           φ.C{10, "Tapete", "molerski_tapete", "molerski_tapete", "molerski"},
		"molerski_lamperije":                        φ.C{11, "Lamperije", "molerski_lamperije", "molerski_lamperije", "molerski"},
		"molerski_strehe":                           φ.C{12, "Strehe", "molerski_strehe", "molerski_strehe", "molerski"},
		"molerski_stepenista":                       φ.C{13, "Stepeništa", "molerski_stepenista", "molerski_stepenista", "molerski"},
		"molerski_podovi":                           φ.C{14, "Podovi", "molerski_podovi", "molerski_podovi", "molerski"},
		"molerski_obrada_fasada":                    φ.C{15, "Obrada fasada", "molerski_obrada_fasada", "molerski_obrada_fasada", "molerski"},
		"molerski_bojenja_fasada":                   φ.C{16, "Bojenja fasada", "molerski_bojenja_fasada", "molerski_bojenja_fasada", "molerski"},
		"livacki":                                   φ.C{27, "Livački", "Livački", "livacki", ""},
		"razni":                                     φ.C{28, "Razni", "Razni", "razni", ""},
		"vodovod":                                   φ.C{29, "Vodovod", "Vodovod", "vodovod", ""},
		"kanalizacija":                              φ.C{30, "Kanalizacija", "Kanalizacija", "kanalizacija", ""},
		"sanitarije":                                φ.C{31, "Sanitarije", "Sanitarije", "sanitarije", ""},
	}
}

func Radovi() φ.T {
	return φ.T{
		Title:       "Rad",
		TitlePlural: "Radovi",
		Slug:        "rad",
		SlugPlural:  "radovi",
		Struct: map[string]φ.F{
			"Title":    φ.F{"Title", "Text", ""},
			"Opis":     φ.F{"Opis", "Text", ""},
			"Obracun":  φ.F{"Obracun", "Text", ""},
			"Jedinica": φ.F{"Jedinica", "Text", ""},
			"Cena":     φ.F{"Cena", "Num", ""},
			"Slug":     φ.F{"Slug", "Text", ""},
			//"Neophodan": φ.F{"Neophodan", "Array", ""},
		},
		Categories: RadoviKategorije(),
	}
}
