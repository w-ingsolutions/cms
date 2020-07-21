package main

// This example launches an IPFS-Lite peer and fetches a hello-world
// hash from the IPFS network.

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/ipfs/go-cid"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
	"github.com/w-ingsolutions/cms/ipfslite"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	lite, err := ipfslite.New(ctx, ds, h, dht, nil)
	if err != nil {
		panic(err)
	}

	lite.Bootstrap(ipfslite.DefaultBootstrapPeers())

	c, _ := cid.Decode("QmSf6YV2ftaHDgamew2b9CjZ4hu6eDyNfwt7qqFCY2RP3c")
	rsc, err := lite.GetFile(ctx, c)
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
