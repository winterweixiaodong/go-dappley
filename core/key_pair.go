package core

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/dappley/go-dappley/util"
	"github.com/dappley/go-dappley/crypto/hash"
	"github.com/dappley/go-dappley/crypto/keystore/secp256k1"
)

const version = byte(0x00)
const addressChecksumLen = 4

type KeyPair struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewKeyPair() *KeyPair {
	private, public := newKeyPair()
	return &KeyPair{private, public}
}

func (w KeyPair) GenerateAddress() Address {
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := util.Base58Encode(fullPayload)

	return NewAddress(fmt.Sprintf("%s", address))
}

//func HashPubKey(pubKey []byte) []byte {
//	publicSHA256 := sha256.Sum256(pubKey)
//
//	RIPEMD160Hasher := ripemd160.New()
//	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
//	if err != nil {
//		log.Panic(err)
//	}
//	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
//
//	return publicRIPEMD160
//}

func HashPubKey(pubKey []byte) []byte {

	sha := hash.Sha3256(pubKey)
	content := hash.Ripemd160(sha)

	//publicSHA256 := sha256.Sum256(pubKey)
	//
	//RIPEMD160Hasher := ripemd160.New()
	//_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	//if err != nil {
	//	log.Panic(err)
	//}
	//publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return content
}

//func checksum(payload []byte) []byte {
//	firstSHA := sha256.Sum256(payload)
//	secondSHA := sha256.Sum256(firstSHA[:])
//
//	return secondSHA[:addressChecksumLen]
//}
//
//func newKeyPair() (ecdsa.PrivateKey, []byte) {
//	curve := elliptic.P256()
//	private, err := ecdsa.GenerateKey(curve, rand.Reader)
//	if err != nil {
//		log.Panic(err)
//	}
//	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
//
//	return *private, pubKey
//}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {

	private, err := secp256k1.NewECDSAPrivateKey()
	//curve := elliptic.P256()
	//private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}
