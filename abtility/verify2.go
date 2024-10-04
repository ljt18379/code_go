/*
************************************************
Copyright: cmri
Author: lijiangtao@chinamobile.com
Date: 2024-06-28
Description: secret VerifySign
************************************************
*/
package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"errors"
	"fmt"
	//"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"math/big"
)

func main() {
	b, _ := VerifySign()
	fmt.Println(b)
}

type EcdsaSignature struct {
	//R *big.Int `asn1:"explicit,tag:0"`
	//S *big.Int `asn1:"explicit,tag:1"`
	R, S *big.Int
}

func VerifySign() (bool, error) {

	pubKey, err := getPublicKeyTest()

	if err != nil {

		return false, err

	}

	//sign, err := getSign(signMsg)
	sign, err := getSign("MEQCIAzWNcPbKYCzdTH34ukFNQ0KETGkvfNG_cCsELg9__P5AiAhKKHdIp9MFztPsC5y3AuCHSf9qDpSxvpsGwcwuVcJ2A")

	if err != nil {
		return false, err

	}

	s := "hello"
	//x := (*[2]uintptr)(unsafe.Pointer(&s))
	//h := [3]uintptr{x[0], x[1], x[1]}
	//bodyByte := *(*[]byte)(unsafe.Pointer(&h))
	sha := sha256.New()
	bodyHash2 := sha.Sum([]byte(s))

	b := ecdsa.Verify(pubKey, bodyHash2[:], sign.R, sign.S)
	if b != true {

		return false, errors.New("not correct sign")

	}

	return b, nil
}

func getPublicKeyTest() (*ecdsa.PublicKey, error) {

	xBase64, err := base64.RawURLEncoding.DecodeString("AIYai_cpl7TbTPtp-i6DIQb7pqNgtZ_fZNkoA8Rs1MMs")

	if err != nil {

		return nil, err

	}

	nx := new(big.Int)

	nx.SetBytes(xBase64)

	yBase64, err := base64.RawURLEncoding.DecodeString("AJPTCLdgUQveqh8tjcK3yUd-bTH4_HbP9GeXSCYCbBtZ")

	if err != nil {

		return nil, err

	}

	ny := new(big.Int)

	ny.SetBytes(yBase64)

	pubKey := &ecdsa.PublicKey{
		Curve: btcec.S256(),
		X:     nx,
		Y:     ny,
	}

	return pubKey, nil
}

func getSign(signMsg string) (EcdsaSignature, error) {

	tsig, err := base64.RawURLEncoding.DecodeString(signMsg)

	var sig EcdsaSignature

	rest, err := asn1.Unmarshal(tsig, &sig)

	if err != nil {
		return sig, err
	}

	if len(rest) > 0 {
		panic("not correct sign")
	}

	return sig, nil
}
