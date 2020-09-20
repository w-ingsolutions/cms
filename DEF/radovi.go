package osnovna

import (
	"fmt"
	"github.com/w-ingsolutions/c/pkg/icons"
	"github.com/w-ingsolutions/cms/pkg/phi"
)

func RadoviKategorije() map[string]phi.C {
	return map[string]phi.C{
		"pripremni":     phi.C{1, "Pripremni", "Pripremni", "pripremni", "", icons.Pripremni},
		"istrazivacki":  phi.C{2, "Istraživački", "Istraživački", "istrazivacki", "", icons.Istrazivacki},
		"rusenja":       phi.C{3, "Rušenja", "Rušenja", "rusenja", "", icons.Rusenja},
		"zemljani":      phi.C{4, "Zemljani", "Zemljani", "zemljani", "", icons.Zemljani},
		"zidarski":      phi.C{5, "Zidarski", "Zidarski", "zidarski", "", icons.Zidarski},
		"betonski":      phi.C{6, "Betonski", "Betonski", "betonski", "", icons.Betonski},
		"tesarski":      phi.C{7, "Tesarski", "Tesarski", "tesarski", "", icons.Tesarski},
		"pokrivacki":    phi.C{8, "Pokrivački", "Pokrivački", "pokrivacki", "", icons.Pokrivacki},
		"izolaterski":   phi.C{9, "Izolaterski", "Izolaterski", "izolaterski", "", icons.Izolaterski},
		"stolarija":     phi.C{10, "Stolarija", "Stolarija", "stolarija", "", icons.Stolarija},
		"stolarski":     phi.C{11, "Stolarski", "Stolarski", "stolarski", "", icons.Stolarski},
		"bravarski":     phi.C{12, "Bravarski", "Bravarski", "bravarski", "", icons.Bravarski},
		"limarski":      phi.C{13, "Limarski", "Limarski", "limarski", "", icons.Limarski},
		"staklorezacki": phi.C{14, "Staklorezački", "Staklorezački", "staklorezacki", "", icons.Staklorezacki},
		"keramicarski":  phi.C{15, "Keramičarski", "Keramičarski", "keramicarski", "", icons.Keramicarski},
		"teracerski":    phi.C{16, "Teracerski", "Teracerski", "teracerski", "", icons.Teracerski},
		"kamenorezacki": phi.C{17, "Kamenorezački", "Kamenorezački", "kamenorezacki", "", icons.Kamenorezacki},
		"parketarski":   phi.C{18, "Parketarski", "Parketarski", "parketarski", "", icons.Parketarski},
		"podopolagacki": phi.C{19, "Podopolagački", "Podopolagački", "podopolagacki", "", icons.Podopolagacki},
		"tapetarski":    phi.C{20, "Tapetarski", "Tapetarski", "tapetarski", "", icons.Tapetarski},
		"roletnarski":   phi.C{21, "Roletnarski", "Roletnarski", "roletnarski", "", icons.Roletnarski},
		"suvomontazni":  phi.C{22, "Suvomontažni", "Suvomontažni", "suvomontazni", "", icons.Suvomontazni},
		"gipsarski":     phi.C{23, "Gipsarski", "Gipsarski", "gipsarski", "", icons.Gipsarski},
		"fasaderski":    phi.C{24, "Fasaderski", "Fasaderski", "fasaderski", "", icons.Fasaderski},
		"likorezacki":   phi.C{25, "Likorezački", "Likorezački", "likorezacki", "", icons.Likorezacki},
		"molerski":      phi.C{26, "Molerski", "Molerski", "molerski", "", icons.Molerski},
		"livacki":       phi.C{27, "Livački", "Livački", "livacki", "", icons.Livacki},
		"razni":         phi.C{28, "Razni", "Razni", "razni", "", icons.Kovacki},
		"vodovod":       phi.C{29, "Vodovod", "Vodovod", "vodovod", "", icons.Vodovod},
		"kanalizacija":  phi.C{30, "Kanalizacija", "Kanalizacija", "kanalizacija", "", icons.Kanalizacija},
		"sanitarije":    phi.C{31, "Sanitarije", "Sanitarije", "sanitarije", "", icons.Sanitarije},
	}
}

func MolerskiPodKategorije() map[string]phi.C {
	return map[string]phi.C{
		"pripremni":                        phi.C{1, "Pripremni", "pripremni", "pripremni", "molerski", icons.Molerski},
		"vrata_i_prozori":                  phi.C{2, "Vrata i prozori", "vrata_i_prozori", "vrata_i_prozori", "molerski", icons.Molerski},
		"zidovi_i_plafoni":                 phi.C{3, "Zidovi i plafoni", "zidovi_i_plafoni", "zidovi_i_plafoni", "molerski", icons.Molerski},
		"tapete":                           phi.C{4, "Tapete", "tapete", "tapete", "molerski", icons.Molerski},
		"lamperije":                        phi.C{5, "Lamperije", "lamperije", "lamperije", "molerski", icons.Molerski},
		"strehe":                           phi.C{6, "Strehe", "strehe", "strehe", "molerski", icons.Molerski},
		"stepenista":                       phi.C{7, "Stepeništa", "stepenista", "stepenista", "molerski", icons.Molerski},
		"podovi":                           phi.C{8, "Podovi", "podovi", "podovi", "molerski", icons.Molerski},
		"obrada_fasada":                    phi.C{9, "Obrada fasada", "obrada_fasada", "obrada_fasada", "molerski", icons.Molerski},
		"bojenja_fasada":                   phi.C{10, "Bojenja fasada", "bojenja_fasada", "bojenja_fasada", "molerski", icons.Molerski},
		"zastita_fasada":                   phi.C{11, "Zaštita fasada", "zastita_fasada", "zastita_fasada", "molerski", icons.Molerski},
		"limarija":                         phi.C{12, "Limarija", "limarija", "limarija", "molerski", icons.Molerski},
		"ograde":                           phi.C{13, "Ograde", "ograde", "ograde", "molerski", icons.Molerski},
		"instalacije":                      phi.C{14, "Instalacije", "instalacije", "instalacije", "molerski", icons.Molerski},
		"mobilijar":                        phi.C{15, "Mobilijar", "mobilijar", "mobilijar", "molerski", icons.Molerski},
		"ostali_molersko_farbarski_radovi": phi.C{16, "Ostali molersko - farbarski radovi", "ostali_molersko_farbarski_radovi", "ostali_molersko_farbarski_radovi", "molerski", icons.Molerski},
	}
}

func Radovi() phi.T {
	return phi.T{
		Title:       "Rad",
		TitlePlural: "Radovi",
		Slug:        "rad",
		SlugPlural:  "radovi",
		Struct:phi.S{
			"Title":     phi.F{"Title", "Text", ""},
			"Opis":      phi.F{"Opis", "Text", ""},
			"Obracun":   phi.F{"Obracun", "Text", ""},
			"Jedinica":  phi.F{"Jedinica", "Text", ""},
			"Cena":      phi.F{"Cena", "Num", ""},
			"Slug":      phi.F{"Slug", "Text", ""},
			//"Neophodan": phi.F{"Neophodan", "Array", map[int]map[string]string{}},
		},
		Categories: RadoviKategorije(),
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
