package cms

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/translate"
	"github.com/w-ingsolutions/cms/db"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	selected int
)

type WingCMS struct {
	Strana         string
	EditPolja      *model.EditabilnaPoljaVrsteRadova
	TipoviSadrzaja map[int]model.TipSadrzaja
	Db             *db.DuoUIdb
	UI             WingUI
	API            WingAPI
	Podesavanja    *WingPodesavanja
	Prikaz         []model.W
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
	Naziv string
	Dir   string
	File  string
	Cyr   bool
}

type WingUloge struct {
	Projektanti []*model.WingLice
	Investotori []*model.WingLice
}
type WingCounters struct {
	Kolicina  *counter.DuoUIcounter
	Radovi    *counter.DuoUIcounter
	Materijal *counter.DuoUIcounter
}
