package cms

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/theme"
	"github.com/marcetin/jdb"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/translate"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	selected int
)

type WingCMS struct {
	Strana               WingStrana
	EditPolja            *model.EditabilnaPoljaVrsteRadova
	TipoviSadrzajaPrikaz []TipSadrzajaPrikaz
	Db                   *jdb.JavazacDB
	UI                   WingUI
	API                  WingAPI
	Podesavanja          WingPodesavanja
	Prikaz               prikaz
}

type prikaz struct {
	w map[string]interface{}
	e []func(gtx C) D
}

type WingUI struct {
	Device   string
	Window   *app.Window
	Tema     *theme.DuoUItheme
	Context  layout.Context
	Ekran    func(gtx layout.Context) layout.Dimensions
	Ops      op.Ops
	Counters WingCounters
}

type WingAPI struct {
	OK     bool
	Adresa string
}

type WingJezik struct {
	t        translate.Translation
	izabrani string
	dostupni []string
	linkovi  map[string]*widget.Clickable
}

type WingPodesavanja struct {
	Naziv          string
	Dir            string
	File           string
	Cyr            bool
	TipoviSadrzaja map[string]TipSadrzaja
}
type TipSadrzaja struct {
	Naziv        string
	NazivMnozina string
	Slug         string
	SlugMnozina  string
	Struktura    []poljeSadrzaja
}
type TipSadrzajaPrikaz struct {
	Naziv        string
	NazivMnozina string
	Slug         string
	SlugMnozina  string
	Struktura    []poljeSadrzaja
	Link         *widget.Clickable
}

type WingUloge struct {
	Projektanti []*model.WingPravnoLice
	Investotori []*model.WingPravnoLice
}
type WingCounters struct {
	Kolicina  *counter.DuoUIcounter
	Radovi    *counter.DuoUIcounter
	Materijal *counter.DuoUIcounter
}

type WingStrana struct {
	Naziv string
	Slug  string
}
