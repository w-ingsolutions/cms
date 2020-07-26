package db

import (
	"encoding/json"
	"fmt"
	"gioui.org/widget"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/w-ingsolutions/c/model"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

type DuoUIdb struct {
	DB     *scribble.Driver
	Folder string `json:"folder"`
	Name   string `json:"name"`
}

type Ddb interface {
	//DbReadAllTypes()map[int]model.WingCalGrupaRadova
	DbRead(folder, name string) model.W
	DbReadAll(folder string) []model.W
	DbWrite(folder, name string, data interface{})
}

func DuoUIdbInit(dataDir string) (d *DuoUIdb) {
	d = new(DuoUIdb)
	db, err := scribble.New(dataDir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	d.DB = db
	return
}

var skip = []*unicode.RangeTable{
	unicode.Mark,
	unicode.Sk,
	unicode.Lm,
}

var safe = []*unicode.RangeTable{
	unicode.Letter,
	unicode.Number,
}

var _ Ddb = &DuoUIdb{}

//func (d *DuoUIdb) DbReadAllTypes() map[int]model.WingCalGrupaRadova{
//	items := make(map[int]model.WingCalGrupaRadova)
//	//types := []string{"assets", "config", "apps"}
//
//	//for t := range types {
//	//	items[t] = d.DbReadAll(t)
//	//}
//	for i := 1; i <= 31; i++ {
//		items[i] = d.DbReadAll(fmt.Sprint(i))
//	}
//	return items
//}
//func (d *DuoUIdb) DbReadTypeAll(f string) map[]model.WingVrstaRadova{
//	return d.DbReadAll(f)
//}

func (d *DuoUIdb) DbReadAll(folder string) (items []model.W) {
	itemsRaw, err := d.DB.ReadAll(folder)
	if err != nil {
		fmt.Println("Error", err)
	}
	//fmt.Println("itemsRaw", itemsRaw)

	for _, bt := range itemsRaw {
		item := model.W{}
		if err := json.Unmarshal([]byte(bt), &item); err != nil {
			fmt.Println("Error", err)
		}
		//fmt.Println("item", item)
		item.Izmena = new(widget.Clickable)
		item.Brisanje = new(widget.Clickable)
		items = append(items, item)
		//fmt.Println("items", items)

	}
	return items
}

func (d *DuoUIdb) DbRead(folder, name string) model.W {
	item := model.W{}
	if err := d.DB.Read(folder, name, &item); err != nil {
		fmt.Println("Error", err)
	}
	return item
}
func (d *DuoUIdb) DbWrite(folder, name string, data interface{}) {
	d.DB.Write(folder, name, data)
}

func slug(text string) string {
	buf := make([]rune, 0, len(text))
	dash := false
	for _, r := range norm.NFKD.String(text) {
		switch {
		case unicode.IsOneOf(safe, r):
			buf = append(buf, unicode.ToLower(r))
			dash = true
		case unicode.IsOneOf(skip, r):
		case dash:
			buf = append(buf, '-')
			dash = false
		}
	}
	if i := len(buf) - 1; i >= 0 && buf[i] == '-' {
		buf = buf[:i]
	}
	return string(buf)
}
