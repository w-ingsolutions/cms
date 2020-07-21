package db

import (
	"github.com/marcetin/jdb"
)

//func cPeer() (context.Context, *Peer, error) {
//
//
//}

func DataBase() {
	d := jdb.DataBase("QmWC7wxt6z69ndm2JPYvkbp6xY6ZJ4ShaKfa98So82wzMy")
	d.Collection()
}
