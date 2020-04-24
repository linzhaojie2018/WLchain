package hashid

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
)

func TestGeneratePlotterID(t *testing.T) {
	u:=GeneratePIDForSP("deeply string pair throw then soldier heavy stair board stare mirror believe")
	fmt.Println(u)
	if i:=u.Cmp(big.NewInt(6287302858570590459));i!=0{
		t.Error("error")
	}
	p:=GeneratePubKey("suppose bread trick afraid been shoot heat game enough stupid number distant")
	strings.EqualFold(p,"6402eed365c0c7a74f8864e8adbef76335c2ad4918c67b12db71264e3ad6ba31")
}

func TestGenerateForPubKey(t *testing.T) {
	id:=GeneratePIDForPubKey("0000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println(id)
	u:=GeneratePIDForSP("Need not perfume women have no future")
	pub:=GeneratePubKey("Need not perfume women have no future")
	pu,pr:=GeneratePubPri("Need not perfume women have no future")
	fmt.Println(u)
	fmt.Println(pub)

	fmt.Println("pub---> ",pu)
	priByte(pu[:])

	fmt.Println("pr---> ",pr)
	priByte(pr[:])
}

func priByte(b []byte)  {
	for i:=0;i<len(b);i++ {
		if i==0 {
			fmt.Printf("[]byte{")
		}
		if i+1 == len(b) {
			fmt.Printf("%d}\n",b[i])
			return
		}
		fmt.Printf("%d, ",b[i])
	}
}