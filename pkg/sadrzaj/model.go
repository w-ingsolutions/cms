package sadrzaj

import (
	"gioui.org/widget"
	"github.com/ipfs/go-cid"
)

type Sadrzaj struct {
	ID         string
	CID        cid.Cid
	Kategorija string
	Struktura  map[string]PoljeSadrzaja
}

type PoljeSadrzaja struct {
	Naziv   string
	Tip     string
	Sadrzaj interface{}
}

type Kategorija struct {
	Naziv  string
	Opis   string
	Slug   string
	Parent string
}

type TipSadrzaja struct {
	Naziv        string
	NazivMnozina string
	Slug         string
	SlugMnozina  string
	Struktura    map[string]PoljeSadrzaja
	Kategorije   map[string]Kategorija
}

//
//type NovoPoljeSadrzaja struct {
//	Naziv string
//	Tip   string
//}

type TipviPoljaSadrzaja struct {
	Naziv string
	Slug  string
}

type TipSadrzajaPrikaz struct {
	Naziv        string
	NazivMnozina string
	Slug         string
	SlugMnozina  string
	Struktura    map[string]PoljeSadrzaja
	Kategorije   map[string]Kategorija
	Link         *widget.Clickable
}
