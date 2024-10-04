package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"math/big"
)

type ecdsaSignature struct {
	R, S *big.Int
}

func main() {
	x, err := base64.RawURLEncoding.DecodeString("RyOjMOIt5qb7NsfT1fzjVS437Se9_fAq9iUPH_-YkbI")

	if err != nil {
		panic(err)
	}
	y, err := base64.RawURLEncoding.DecodeString("ALKGML19DRubnDP10mXZ8B9mbij3O4GkLTSZKDeJgThO")

	if err != nil {
		panic(err)
	}
	//fmt.Println(string(x))
	nx := new(big.Int)
	nx.SetBytes(x)
	fmt.Println(nx)

	ny := new(big.Int)
	ny.SetBytes(y)
	fmt.Println(ny)

	ppk := &ecdsa.PublicKey{
		Curve: btcec.S256(),
		X:     nx,
		Y:     ny}
	fmt.Println(ppk)

	tsig, err := base64.RawURLEncoding.DecodeString("MEUCIQD7YCNgyWV92E9bCGjUhOdGvBmrpuzwntgc1x3lRXt9qQIgXY7tBAfgOv-3bvS2KDgtFkANzwsT0xqK7FAJcWb_XEk")
	var sig ecdsaSignature
	rest, err := asn1.Unmarshal(tsig, &sig)
	if err != nil {
		panic(err)
	}
	if len(rest) > 0 {
		fmt.Println("warring: decode:", rest)
	}
	ok := ecdsa.Verify(ppk, mhs[:], sig.R, sig.S)
	if ok != true {
		panic("not correct sign")
	}
	fmt.Println("ok")
}

type ss struct{}

func sha2() {
	sha := sha256.New()
	str := "hello, world"
	hash := sha.Sum([]byte(str))
	sha.Write()
	hexHash := hex.EncodeToString(hash[:])
	fmt.Printf("SHA256 Hash of %s: %s\n", str, hexHash)
}
