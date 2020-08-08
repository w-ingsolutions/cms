package osnovna

import (
	"github.com/w-ingsolutions/cms/pkg/φ"
)

func NewRadovi() map[int]phi.Φ {
	return map[int]phi.Φ{
		1: phi.Φ{
			ID:       1,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Skidanje boje sa vrata i prozora hemijskim putem"},
				"Opis":     phi.F{"Opis", "Text", "Postojece slojeve boje skinuti nanošenjem hemijskog rastvaraca i skidanjem slojeva boje špahtlama i odgovarajucim alatkama \nPostupak ponavljati dok se ne skinu svi slojevi boje i ne dode do zdravog i cistog drveta \nNarocitu pažnju obratiti da ne dode do oštecenja ivica i profilacije"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 ocišcene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.50"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_1"},
			},
		},
		2: phi.Φ{
			ID:       2,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Skidanje boje sa vrata i prozora upotrebom specijalnog fena za skidanje boje"},
				"Opis":     phi.F{"Opis", "Text", "Postojece slojeve boje skinuti grejanjem fenom i upotrebom špahtli i odgovarajucih alatki Postupak ponavljati dok se ne skinu svi slojevi boje i ne dode do zdravog i cistog drveta \nPo izvršenom skidanju boje drvo prebrusiti finom šmirglom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 ocišcene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "4.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_2"},
			},
		},
		3: phi.Φ{
			ID:       3,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Skidanje boje sa vrata i prozora paljenjem"},
				"Opis":     phi.F{"Opis", "Text", "Postojece slojeve boje skinuti paljenjem brenerom sa plinom i upotrebom špahtli i odgovarajucih alatki Postupak ponavljati dok se ne skinu svi slojevi boje i ne dode do zdravog i cistog drveta \nPo izvršenom skidanju boje drvo prebrusiti finom šmirglom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 ocišcene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "5.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_3"},
			},
		},
		4: phi.Φ{
			ID:       4,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih prozora i vrata u jednom tonu"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta \nPre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov \nPrevuci uljanim kitom prvi put, brusiti i ponovno kitovati \nBojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "5.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_4"},
			},
		},
		5: phi.Φ{
			ID:       5,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih prozora i vrata u dva tona"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta \nPre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov \nPrevuci uljanim kitom prvi put, brusiti i ponovno kitovati \nBojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "5.75"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_5"},
			},
		},
		6: phi.Φ{
			ID:       6,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih prozora i vrata, kvalitetnije bojenje u jednom tonu"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov Prevuci uljanim kitom prvi put, brusiti i ponovno kitovati Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "6.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_6"},
			},
		},
		7: phi.Φ{
			ID:       7,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih prozora i vrata, kvalitetnije bojenje u dva tona"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta \nPre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov Prevuci uljanim kitom prvi put, brusiti i ponovno kitovati Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "6.75"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_7"},
			},
		},
		8: phi.Φ{
			ID:       8,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih prozora i vrata, prvorazredno bojenje  u jednom tonu"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov Prevuci uljanim kitom prvi put, brusiti i ponovno prevuci uljani kit, drugi put Brusiti i ponovo kitovati Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "8.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_8"},
			},
		},
		9: phi.Φ{
			ID:       9,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih prozora i vrata, prvorazredno bojenje  u dva tona"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo ocistiti od maltera i prašine, natopiti razredenim firnisom i minizirati gvozdeni okov Prevuci uljanim kitom prvi put, brusiti i ponovno prevuci uljani kit, drugi put Brusiti i ponovo kitovati Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "9.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_9"},
			},
		},
		10: phi.Φ{
			ID:       10,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih prozora i vrata lazurnim bojama, sa lakiranjem"},
				"Opis":     phi.F{"Opis", "Text", "Bojiti sadolinom ili nekim slicnim sredstvom po izboru projektanta Pre bojenja sve površine preci finom šmirglom, da ostane glatka površina"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.50"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_10"},
			},
		},
		11: phi.Φ{
			ID:       11,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih prozora i vrata lazurnim bojama, bez lakiranja"},
				"Opis":     phi.F{"Opis", "Text", "Bojiti sadolinom ili nekim slicnim sredstvom po izboru projektanta Pre bojenja sve površine preci finom šmirglom, da ostane glatka površina"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_11"},
			},
		},
		12: phi.Φ{
			ID:       12,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata preko postojece boje u jednom tonu"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "4.25"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_12"},
			},
		},
		13: phi.Φ{
			ID:       13,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata preko postojece boje u dva tona"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "4.75"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_13"},
			},
		},
		14: phi.Φ{
			ID:       14,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata preko postojece boje, kvalitetnije bojenje u jednom tonu"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "4.75"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_14"},
			},
		},
		15: phi.Φ{
			ID:       15,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata preko postojece boje, kvalitetnije bojenje u dva tona"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "5.25"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_15"},
			},
		},
		16: phi.Φ{
			ID:       16,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata preko postojece boje, prvorazredna izrada u jednom tonu"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "5.50"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_16"},
			},
		},
		17: phi.Φ{
			ID:       17,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata preko postojece boje, prvorazredna izrada u dva tona"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja sve površine brusiti, ocistiti i kitovati oštecenja i pukotine Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "6.50"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_17"},
			},
		},
		18: phi.Φ{
			ID:       18,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja u jednom tonu"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "6.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_18"},
			},
		},
		19: phi.Φ{
			ID:       19,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja u dva tona tonu"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom, po sušenju brusiti i nadkitovati uljanim kitom Fino brusiti i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "6.75"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_19"},
			},
		},
		20: phi.Φ{
			ID:       20,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja, kvalitetnije bojenje u jednom tonu"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "7.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_20"},
			},
		},
		21: phi.Φ{
			ID:       21,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja, kvalitetnije bojenje u dva tona"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom prvi put, po sušenju brusiti i nadkitovati uljanim kitom Bojiti uljanom bojom, drugi put Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "7.75"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_21"},
			},
		},
		22: phi.Φ{
			ID:       22,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja, prvorazredno bojenje u jednom tonu"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "8.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_22"},
			},
		},
		23: phi.Φ{
			ID:       23,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata, sa kojih je skinuta stara boja, prvorazredno bojenje u dva tona"},
				"Opis":     phi.F{"Opis", "Text", "Vrsta boje, proizvodac i ton po izboru projektanta Pre bojenja drvo brusiti i ocistiti, a zatim naneti podlogu od firnisa sa dodatkom uljane boje Bojiti uljanom bojom prvi put Brusiti i bojiti emajl lakom Fino brusiti, ispraviti emajl kitom i bojiti emajl lakom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "9.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_23"},
			},
		},
		24: phi.Φ{
			ID:       24,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata lazurnim bojama, sa lakiranjem"},
				"Opis":     phi.F{"Opis", "Text", "Bojiti sadolinom ili nekim slicnim sredstvom po izboru projektanta Pre bojenja sve površine preci finom šmirglom, da ostane glatka površina"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.25"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_24"},
			},
		},
		25: phi.Φ{
			ID:       25,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih prozora i vrata lazurnim bojama, sa lakiranja"},
				"Opis":     phi.F{"Opis", "Text", "Bojiti sadolinom ili nekim slicnim sredstvom po izboru projektanta Pre bojenja sve površine preci finom šmirglom, da ostane glatka površina"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "2.75"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_25"},
			},
		},
		26: phi.Φ{
			ID:       26,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Lakiranje površina vrata i prozora, preko lazurne boje"},
				"Opis":     phi.F{"Opis", "Text", "Sve površine lakirati mat ili sjajnim lakom, po izboru projektanta Lakirati prvi put, posle 24 h preci najfinijom šmirglom, opajati i lakirati drugi put"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 lakirane površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "2.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_26"},
			},
		},
		27: phi.Φ{
			ID:       27,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje osnovnom bojom metalnih prozora, sa prethodnim cišcenjem prozora"},
				"Opis":     phi.F{"Opis", "Text", "Metalne prozore ocistiti od korozije i prašine hemijskim i fizickim sredstvima, brusiti i ocistiti Naneti impregnaciju i obojiti osnovnom bojom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obradene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "1.25"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_27"},
			},
		},

		28: phi.Φ{
			ID:       28,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje osnovnom bojom metalnih prozora, sa prethodnim cišcenjem vrata"},
				"Opis":     phi.F{"Opis", "Text", "Metalne prozore ocistiti od korozije i prašine hemijskim i fizickim sredstvima, brusiti i ocistiti Naneti impregnaciju i obojiti osnovnom bojom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obradene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "1.25"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_28"},
			},
		},
		29: phi.Φ{
			ID:       29,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje osnovnom bojom metalne kapije, sa prethodnim cišcenjem kapije"},
				"Opis":     phi.F{"Opis", "Text", "Metalnu kapiju ocistiti od korozije i prašine hemijskim i fizickim sredstvima, brusiti i ocistiti Naneti impregnaciju i obojiti osnovnom bojom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obradene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "1.75"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_29"},
			},
		},
		30: phi.Φ{
			ID:       30,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje osnovnom bojom metalne kapije, sa prethodnim cišcenjem ograde"},
				"Opis":     phi.F{"Opis", "Text", "Metalnu kapiju ocistiti od korozije i prašine hemijskim i fizickim sredstvima, brusiti i ocistiti Naneti impregnaciju i obojiti osnovnom bojom"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obradene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "1.75"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_30"},
			},
		},
		31: phi.Φ{
			ID:       31,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih metalnih prozora, bojom za metal"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.75"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_31"},
			},
		},
		32: phi.Φ{
			ID:       32,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih metalnih vrata, bojom za metal"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.50"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_32"},
			},
		},
		33: phi.Φ{
			ID:       33,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih metalnih protivpožarnih vrata, bojom za metal"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.25"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_33"},
			},
		},
		34: phi.Φ{
			ID:       34,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih garažnih vrata, bojom za metal"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.50"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_34"},
			},
		},
		35: phi.Φ{
			ID:       35,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih metalnih rešetki prozora, bojom za metal"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_35"},
			},
		},
		36: phi.Φ{
			ID:       36,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih metalnih rešetki vrata, bojom za metal"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.50"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_36"},
			},
		},
		37: phi.Φ{
			ID:       37,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje novih metalnihroletni, bojom za metal"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja sa metala skinuti koroziju hemijskim i fizickim sredstvima, a zatim sve površine brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "4.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_37"},
			},
		},
		38: phi.Φ{
			ID:       38,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih metalnih prozora, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.50"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_38"},
			},
		},
		39: phi.Φ{
			ID:       39,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih metalnih vrata, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "4.50"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_39"},
			},
		},
		40: phi.Φ{
			ID:       40,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih protivpožarnih vrata, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "4.25"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_40"},
			},
		},
		41: phi.Φ{
			ID:       41,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih garažnih vrata, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na prozore naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "4.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_41"},
			},
		},
		42: phi.Φ{
			ID:       42,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih metalnih rešetki prozora, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "3.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_42"},
			},
		},
		43: phi.Φ{
			ID:       43,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih metalnih roletni, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "4.50"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_43"},
			},
		},
		44: phi.Φ{
			ID:       44,
			Category: "molerski_pripremni",
			Struct: map[string]phi.F{
				"Title":    phi.F{"Title", "Text", "Bojenje starih metalnih rešetki vrata, bojom za metal, sa prethodnim skidanjem stare boje"},
				"Opis":     phi.F{"Opis", "Text", "Pre bojenja skinuti staru boju i koroziju hemijskim i fizickim sredstvima, brusiti i ocistiti Na rešetke naneti impregnaciju i osnovnu boju, a zatim predkitovati i brusiti Obojiti drugi put bojom za metal, u tonu po izboru projektanta"},
				"Obracun":  phi.F{"Obracun", "Text", "Obracun po m2 obojene površine"},
				"Jedinica": phi.F{"Jedinica", "Text", "m2"},
				"Cena":     phi.F{"Cena", "Num", "4.00"},
				"Slug":     phi.F{"Slug", "Text", "molerski_pripremni_44"},
			},
		},
	}
}
