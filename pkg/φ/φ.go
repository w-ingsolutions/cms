package φ

import (
	"gioui.org/widget"
	"github.com/ipfs/go-cid"
)

type Φ struct {
	ID       int
	CID      cid.Cid
	Category string
	Struct   map[string]F
}

type F struct {
	Title   string
	Type    string
	Content interface{}
}

type C struct {
	ID     int
	Title  string
	Desc   string
	Slug   string
	Parent string
}

type T struct {
	Title       string
	TitlePlural string
	Slug        string
	SlugPlural  string
	Struct      map[string]F
	Categories  map[string]C
}

type FieldType struct {
	Title string
	Slug  string
}

type ContentType struct {
	Title       string
	TitlePlural string
	Slug        string
	SlugPlural  string
	Struct      map[string]F
	Categories  map[string]C
	Link        *widget.Clickable
}
