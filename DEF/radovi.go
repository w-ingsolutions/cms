package osnovna

import (
	"github.com/w-ingsolutions/cms/pkg/φ"
)

func RadoviKategorije() map[string]phi.C {
	return map[string]phi.C{
		"pripremni":                phi.C{1, "Pripremni", "Pripremni", "pripremni", ""},
		"istrazivacki":             phi.C{2, "Istraživački", "Istraživački", "istrazivacki", ""},
		"rusenja":                  phi.C{3, "Rušenja", "Rušenja", "rusenja", ""},
		"zemljani":                 phi.C{4, "Zemljani", "Zemljani", "zemljani", ""},
		"zidarski":                 phi.C{5, "Zidarski", "Zidarski", "zidarski", ""},
		"betonski":                 phi.C{6, "Betonski", "Betonski", "betonski", ""},
		"tesarski":                 phi.C{7, "Tesarski", "Tesarski", "tesarski", ""},
		"pokrivacki":               phi.C{8, "Pokrivački", "Pokrivački", "pokrivacki", ""},
		"izolaterski":              phi.C{9, "Izolaterski", "Izolaterski", "izolaterski", ""},
		"stolarija":                phi.C{10, "Stolarija", "Stolarija", "stolarija", ""},
		"stolarski":                phi.C{11, "Stolarski", "Stolarski", "stolarski", ""},
		"bravarski":                phi.C{12, "Bravarski", "Bravarski", "bravarski", ""},
		"limarski":                 phi.C{13, "Limarski", "Limarski", "limarski", ""},
		"staklorezacki":            phi.C{14, "Staklorezački", "Staklorezački", "staklorezacki", ""},
		"keramicarski":             phi.C{15, "Keramičarski", "Keramičarski", "keramicarski", ""},
		"teracerski":               phi.C{16, "Teracerski", "Teracerski", "teracerski", ""},
		"kamenorezacki":            phi.C{17, "Kamenorezački", "Kamenorezački", "kamenorezacki", ""},
		"parketarski":              phi.C{18, "Parketarski", "Parketarski", "parketarski", ""},
		"podopolagacki":            phi.C{19, "Podopolagački", "Podopolagački", "podopolagacki", ""},
		"tapetarski":               phi.C{20, "Tapetarski", "Tapetarski", "tapetarski", ""},
		"roletnarski":              phi.C{21, "Roletnarski", "Roletnarski", "roletnarski", ""},
		"suvomontazni":             phi.C{22, "Suvomontažni", "Suvomontažni", "suvomontazni", ""},
		"gipsarski":                phi.C{23, "Gipsarski", "Gipsarski", "gipsarski", ""},
		"fasaderski":               phi.C{24, "Fasaderski", "Fasaderski", "fasaderski", ""},
		"likorezacki":              phi.C{25, "Likorezački", "Likorezački", "likorezacki", ""},
		"molerski":                 phi.C{26, "Molerski", "Molerski", "molerski", ""},
		"molerski_pripremni":       phi.C{1, "Pripremni", "molerski_pripremni", "molerski_pripremni", "molerski"},
		"molerski_vrata_i_prozori": phi.C{2, "Vrata i prozori", "molerski_vrata_i_prozori", "molerski_vrata_i_prozori", "molerski"},
		"molerski_zastita_fasada":  phi.C{3, "Zaštita fasada", "molerski_zastita_fasada", "molerski_zastita_fasada", "molerski"},
		"molerski_limarija":        phi.C{4, "Limarija", "molerski_limarija", "molerski_limarija", "molerski"},
		"molerski_ograde":          phi.C{5, "Ograde", "molerski_ograde", "molerski_ograde", "molerski"},
		"molerski_instalacije":     phi.C{6, "Instalacije", "molerski_instalacije", "molerski_instalacije", "molerski"},
		"molerski_mobilijar":       phi.C{7, "Mobilijar", "molerski_mobilijar", "molerski_mobilijar", "molerski"},
		"molerski_ostali_molersko_farbarski_radovi": phi.C{8, "Ostali molersko - farbarski radovi", "molerski_ostali_molersko_farbarski_radovi", "molerski_ostali_molersko_farbarski_radovi", "molerski"},
		"molerski_zidovi_i_plafoni":                 phi.C{9, "Zidovi i plafoni", "molerski_zidovi_i_plafoni", "molerski_zidovi_i_plafoni", "molerski"},
		"molerski_tapete":                           phi.C{10, "Tapete", "molerski_tapete", "molerski_tapete", "molerski"},
		"molerski_lamperije":                        phi.C{11, "Lamperije", "molerski_lamperije", "molerski_lamperije", "molerski"},
		"molerski_strehe":                           phi.C{12, "Strehe", "molerski_strehe", "molerski_strehe", "molerski"},
		"molerski_stepenista":                       phi.C{13, "Stepeništa", "molerski_stepenista", "molerski_stepenista", "molerski"},
		"molerski_podovi":                           phi.C{14, "Podovi", "molerski_podovi", "molerski_podovi", "molerski"},
		"molerski_obrada_fasada":                    phi.C{15, "Obrada fasada", "molerski_obrada_fasada", "molerski_obrada_fasada", "molerski"},
		"molerski_bojenja_fasada":                   phi.C{16, "Bojenja fasada", "molerski_bojenja_fasada", "molerski_bojenja_fasada", "molerski"},
		"livacki":                                   phi.C{27, "Livački", "Livački", "livacki", ""},
		"razni":                                     phi.C{28, "Razni", "Razni", "razni", ""},
		"vodovod":                                   phi.C{29, "Vodovod", "Vodovod", "vodovod", ""},
		"kanalizacija":                              phi.C{30, "Kanalizacija", "Kanalizacija", "kanalizacija", ""},
		"sanitarije":                                phi.C{31, "Sanitarije", "Sanitarije", "sanitarije", ""},
	}
}

func Radovi() phi.T {
	return phi.T{
		Title:       "Rad",
		TitlePlural: "Radovi",
		Slug:        "rad",
		SlugPlural:  "radovi",
		Struct: map[string]phi.F{
			"Title":    phi.F{"Title", "Text", ""},
			"Opis":     phi.F{"Opis", "Text", ""},
			"Obracun":  phi.F{"Obracun", "Text", ""},
			"Jedinica": phi.F{"Jedinica", "Text", ""},
			"Cena":     phi.F{"Cena", "Num", ""},
			"Slug":     phi.F{"Slug", "Text", ""},
			//"Neophodan": φ.F{"Neophodan", "Array", ""},
		},
		Categories: RadoviKategorije(),
	}
}
