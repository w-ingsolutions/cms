package main

import (
	"context"
	"fmt"
	format "github.com/ipfs/go-ipld-format"
	"io/ioutil"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
	ipfslite "github.com/w-ingsolutions/cms/db/pkg"
)

type idb struct {
	c  context.Context
	p  *ipfslite.Peer
	cd string
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.SetLogLevel("*", "warn")
	// Bootstrappers are using 1024 keys. See:
	// https://github.com/ipfs/infra/issues/378
	crypto.MinRsaKeyBits = 1024
	ds, err := ipfslite.BadgerDatastore("test")
	if err != nil {
		panic(err)
	}
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		panic(err)
	}
	listen, _ := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/4005")
	h, dht, err := ipfslite.SetupLibp2p(
		ctx,
		priv,
		nil,
		[]multiaddr.Multiaddr{listen},
		ds,
		ipfslite.Libp2pOptionsExtra...,
	)
	if err != nil {
		panic(err)
	}
	p, err := ipfslite.New(ctx, ds, h, dht, nil)
	if err != nil {
		panic(err)
	}
	p.Bootstrap(ipfslite.DefaultBootstrapPeers())
	ib := &idb{
		c:  ctx,
		p:  p,
		cd: "QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv",
	}
	ib.collection()
}

func (ib *idb) collection() {
	c, _ := cid.Decode(ib.cd)
	node, err := ib.p.Get(ib.c, c)
	if err != nil {
		panic(err)
	}
	navNode := format.NewNavigableIPLDNode(node, ib.p.DAGService)
	for i := 0; i < int(navNode.ChildTotal()); i++ {
		childNode, err := navNode.FetchChild(ib.c, uint(i))
		if err != nil {
			panic(err)
		}
		n := format.ExtractIPLDNode(childNode)
		childCID := n.Cid().String()
		fmt.Println("graphOut", childCID)
	}
}

func (ib *idb) item() {
	c, _ := cid.Decode(ib.cd)
	rsc, err := ib.p.GetFile(ib.c, c)
	if err != nil {
		panic(err)
	}
	defer rsc.Close()
	content, err := ioutil.ReadAll(rsc)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
