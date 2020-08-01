package materijal

import (
	"github.com/w-ingsolutions/cms/pkg/sadrzaj"
)

func NewRadovi() []sadrzaj.Sadrzaj {
	return []sadrzaj.Sadrzaj{
		sadrzaj.Sadrzaj{
			ID:         "1",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Skidanje boje sa vrata i prozora hemijskim putem"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Postojece slojeve boje skinuti nanošenjem hemijskog rastvaraca i skidanjem slojeva boje špahtlama i odgovarajucim alatkama \nPostupak ponavljati dok se ne skinu svi slojevi boje i ne dode do zdravog i cistog drveta \nNarocitu pažnju obratiti da ne dode do oštecenja ivica i profilacije"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 ocišcene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.50"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_1"},
			},
		},
		sadrzaj.Sadrzaj{
			ID:         "2",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Skidanje boje sa vrata i prozora upotrebom specijalnog fena za skidanje boje"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Postojece slojeve boje skinuti grejanjem fenom i upotrebom špahtli i odgovarajucih alatki Postupak ponavljati dok se ne skinu svi slojevi boje i ne dode do zdravog i cistog drveta \nPo izvršenom skidanju boje drvo prebrusiti finom šmirglom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 ocišcene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "4.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_2"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "3",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Skidanje boje sa vrata i prozora paljenjem"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Postojece slojeve boje skinuti paljenjem brenerom sa plinom i upotrebom špahtli i odgovarajucih alatki Postupak ponavljati dok se ne skinu svi slojevi boje i ne dode do zdravog i cistog drveta \nPo izvršenom skidanju boje drvo prebrusiti finom šmirglom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 ocišcene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "5.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_3"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "4",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih prozora i vrata u jednom tonu"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta \nPre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov \nPrevuci uljanim kitom prvi put, brusiti i ponovno kitovati \nBojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "5.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_4"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "5",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih prozora i vrata u dva tona"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta \nPre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov \nPrevuci uljanim kitom prvi put, brusiti i ponovno kitovati \nBojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "5.75"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_5"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "6",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih prozora i vrata, kvalitetnije bojenje u jednom tonu"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov Prevuci uljanim kitom prvi put, brusiti i ponovno kitovati Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "6.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_6"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "7",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih prozora i vrata, kvalitetnije bojenje u dva tona"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta \nPre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov Prevuci uljanim kitom prvi put, brusiti i ponovno kitovati Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "6.75"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_7"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "8",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih prozora i vrata, prvorazredno bojenje  u jednom tonu"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov Prevuci uljanim kitom prvi put, brusiti i ponovno prevuci uljani kit, drugi put Brusiti i ponovo kitovati Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "8.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_8"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "9",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih prozora i vrata, prvorazredno bojenje  u dva tona"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov Prevuci uljanim kitom prvi put, brusiti i ponovno prevuci uljani kit, drugi put Brusiti i ponovo kitovati Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "9.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_9"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "10",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih prozora i vrata lazurnim bojama, sa lakiranjem"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Bojiti sadolinom ili nekim slicnim sredstvom po izboru projektanta Pre bojenja sve površine preci finom šmirglom, da ostane glatka površina"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.50"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_10"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "11",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih prozora i vrata lazurnim bojama, bez lakiranja"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Bojiti sadolinom ili nekim slicnim sredstvom po izboru projektanta Pre bojenja sve površine preci finom šmirglom, da ostane glatka površina"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_11"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "12",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata preko postojece boje u jednom tonu"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "4.25"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_12"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "13",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata preko postojece boje u dva tona"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "4.75"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_13"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "14",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata preko postojece boje, kvalitetnije bojenje u jednom tonu"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "4.75"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_14"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "15",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata preko postojece boje, kvalitetnije bojenje u dva tona"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "5.25"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_15"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "16",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata preko postojece boje, prvorazredna izrada u jednom tonu"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "5.50"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_16"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "17",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata preko postojece boje, prvorazredna izrada u dva tona"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "6.50"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_17"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "18",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja u jednom tonu"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "6.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_18"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "19",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja u dva tona tonu"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "6.75"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_19"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "20",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja, kvalitetnije bojenje u jednom tonu"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "7.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_20"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "21",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja, kvalitetnije bojenje u dva tona"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "7.75"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_21"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "22",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja, prvorazredno bojenje u jednom tonu"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "8.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_22"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "23",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja, prvorazredno bojenje u dva tona"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "9.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_23"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "24",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata lazurnim bojama, sa lakiranjem"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Bojiti sadolinom ili nekim slicnim sredstvom po izboru projektanta Pre bojenja sve površine preci finom šmirglom, da ostane glatka površina"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.25"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_24"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "25",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih prozora i vrata lazurnim bojama, sa lakiranja"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Bojiti sadolinom ili nekim slicnim sredstvom po izboru projektanta Pre bojenja sve površine preci finom šmirglom, da ostane glatka površina"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "2.75"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_25"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "26",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Lakiranje površina vrata i prozora, preko lazurne boje"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Sve površine lakirati mat ili sjajnim lakom, po izboru projektanta Lakirati prvi put, posle 24 h preci najfinijom šmirglom, opajati i lakirati drugi put"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 lakirane površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "2.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_26"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "27",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje osnovnom bojom metalnih prozora, sa prethodnim cišcenjem prozora"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Metalne prozore ocistiti od korozije i prašine hemijskim i fizickim sredstvima, brusiti i ocistiti Naneti impregnaciju i obojiti osnovnom bojom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obradene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "1.25"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_27"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "28",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje osnovnom bojom metalnih prozora, sa prethodnim cišcenjem vrata"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Metalne prozore ocistiti od korozije i prašine hemijskim i fizickim sredstvima, brusiti i ocistiti Naneti impregnaciju i obojiti osnovnom bojom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obradene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "1.25"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_28"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "29",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje osnovnom bojom metalne kapije, sa prethodnim cišcenjem kapije"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Metalnu kapiju ocistiti od korozije i prašine hemijskim i fizickim sredstvima, brusiti i ocistiti Naneti impregnaciju i obojiti osnovnom bojom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obradene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "1.75"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_29"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "30",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje osnovnom bojom metalne kapije, sa prethodnim cišcenjem ograde"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Metalnu kapiju ocistiti od korozije i prašine hemijskim i fizickim sredstvima, brusiti i ocistiti Naneti impregnaciju i obojiti osnovnom bojom"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obradene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "1.75"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_30"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "31",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih metalnih prozora, bojom za metal"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.75"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_31"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "32",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih metalnih vrata, bojom za metal"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.50"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_32"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "33",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih metalnih protivpožarnih vrata, bojom za metal"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.25"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_33"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "34",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih garažnih vrata, bojom za metal"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.50"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_34"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "35",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih metalnih rešetki prozora, bojom za metal"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_35"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "36",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih metalnih rešetki vrata, bojom za metal"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.50"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_36"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "37",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje novih metalnihroletni, bojom za metal"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "4.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_37"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "38",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih metalnih prozora, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.50"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_38"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "39",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih metalnih vrata, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "4.50"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_39"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "40",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih protivpožarnih vrata, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "4.25"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_40"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "41",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih garažnih vrata, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "4.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_41"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "42",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih metalnih rešetki prozora, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "3.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_42"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "43",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih metalnih roletni, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "4.50"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_43"},
			},
		},

		sadrzaj.Sadrzaj{
			ID:         "44",
			Kategorija: "molerski_pripremni",
			Struktura: map[string]sadrzaj.PoljeSadrzaja{
				"Naziv":    sadrzaj.PoljeSadrzaja{"Naziv", "Text", "Bojenje starih metalnih rešetki vrata, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     sadrzaj.PoljeSadrzaja{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  sadrzaj.PoljeSadrzaja{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": sadrzaj.PoljeSadrzaja{"Jedinica", "Text", "m2"},
				"Cena":     sadrzaj.PoljeSadrzaja{"Cena", "Num", "4.00"},
				"Slug":     sadrzaj.PoljeSadrzaja{"Slug", "Text", "molerski_pripremni_44"},
			},
		},
	}
}
