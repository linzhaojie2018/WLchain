package hashid

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/curve25519"
	"log"
	"math/big"
)

func toPlotterID(pubKey [32]byte) *big.Int {
	S256:=sha256.New()
	S256.Write(pubKey[:])
	bp:=S256.Sum(nil)

	result:=big.NewInt(0)
	func(b []byte) {
		w:=0
		for i:=0;i<8;i++ {
			result.Add(result,big.NewInt(int64(bp[i])<<w))
			w+=8
		}
	}(bp)
	return result
}


//accountid  =  plotterid
func GeneratePIDForSP(secretPhrase string) *big.Int {
	pubkey,_ :=GeneratePubPri(secretPhrase)
	return toPlotterID(pubkey)
}

//公钥hex
func GeneratePubKey(secretPhrase string) string {
	pubkey,_ :=GeneratePubPri(secretPhrase)
	return hex.EncodeToString(pubkey[:])
}

//生成公私钥  secretPhrase
func GeneratePubPri(secretPhrase string) (pub [32]byte,pri [32]byte) {
	S256:=sha256.New()
	if _,e:=S256.Write([]byte(secretPhrase)); e!=nil{
		log.Printf("generatePlotterID is err=%s\n",e)
		return pub,pri
	}
	hb:=S256.Sum(nil)
	var pubKey [32]byte
	var priKey [32]byte
	copy(priKey[:],hb)
	curve25519.ScalarBaseMult(&pubKey,&priKey)
	return pubKey,priKey
}

//通过公钥生产PID
func GeneratePIDForPubKey(pub string) *big.Int{
	if b,e:=hex.DecodeString(pub);e==nil{
		var pubKey [32]byte
		//var priKey [32]byte
		copy(pubKey[:],b)
		//curve25519.ScalarBaseMult(&pubKey,&priKey)
		in:=toPlotterID(pubKey)
		return in
	}else {
		return nil
	}
}