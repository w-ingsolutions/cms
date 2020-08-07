package content

import (
	"github.com/ipfs/go-cid"
)

type Content struct {
	ID       int
	CID      cid.Cid
	Category string
	Struct   map[string]Field
}

type Field struct {
	Title   string
	Type    string
	Content interface{}
}

type Category struct {
	ID     int
	Title  string
	Opis   string
	Slug   string
	Parent string
}

type Type struct {
	Title       string
	TitlePlural string
	Slug        string
	SlugPlural  string
	Struct      map[string]Field
	Categories  map[string]Category
}

//
//type NovoPoljeSadrzaja struct {
//	Title string
//	Tip   string
//}

type FieldType struct {
	Title string
	Slug  string
}

//type TipSadrzajaPrikaz struct {
//	Title        string
//	TitlePlural string
//	Slug         string
//	SlugPlural  string
//	Struct    map[string]PoljeSadrzaja
//	Kategorije   map[string]Kategorija
//	Link         *widget.Clickable
//}
