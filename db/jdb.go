package db

import (
	"bytes"
	"encoding/gob"
	"github.com/marcetin/jdb"
	"github.com/w-ingsolutions/c/model"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

type DuoUIdb struct {
	DB     *jdb.JavazacDB
	Folder string `json:"folder"`
	Name   string `json:"name"`
	//Tipovi map[string]interface{}
}

type Ddb interface {
	//DbReadAllTypes()map[int]model.WingCalGrupaRadova
	DbRead(name string) model.W
	//DbReadAll(folder string) []model.W
	DbWrite(data interface{})
}

func DuoUIdbInit(dataDir string) (d *DuoUIdb) {
	d = new(DuoUIdb)
	d.DB = jdb.New(dataDir)

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

//func (d *DuoUIdb) DbReadAll(folder string) (items []model.W) {
//	itemsRaw, err := d.DB.ReadAll(folder)
//	if err != nil {
//		fmt.Println("Error", err)
//	}
//	//fmt.Println("itemsRaw", itemsRaw)
//
//	for _, bt := range itemsRaw {
//		item := model.W{}
//		if err := json.Unmarshal([]byte(bt), &item); err != nil {
//			fmt.Println("Error", err)
//		}
//		//fmt.Println("item", item)
//		item.Izmena = new(widget.Clickable)
//		item.Brisanje = new(widget.Clickable)
//		items = append(items, item)
//		//fmt.Println("items", items)
//
//	}
//	return items
//}

func (d *DuoUIdb) DbRead(hash string) model.W {
	item := model.W{}
	d.DB.Read(hash, &item)
	return item
}
func (d *DuoUIdb) DbWrite(data interface{}) {
	var bytesBuf bytes.Buffer
	encoder := gob.NewEncoder(&bytesBuf)
	err := encoder.Encode(data)
	if err != nil {

	}
	d.DB.Write(bytesBuf.Bytes())
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
