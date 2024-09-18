package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"log"
)

func main() {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println("privateKey:", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKey:", hexutil.Encode(publicKeyBytes)[2:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	address = "41" + address[2:]

	fmt.Println("address hex: ", address)

	addb, _ := hex.DecodeString(address)

	hash1 := s256(s256(addb))

	secret := hash1[:4]

	for _, v := range secret {
		addb = append(addb, v)
	}

	fmt.Println("address base58: ", base58.Encode(addb))

}

func s256(s []byte) []byte {

	h := sha256.New()

	h.Write(s)

	bs := h.Sum(nil)

	return bs

}
