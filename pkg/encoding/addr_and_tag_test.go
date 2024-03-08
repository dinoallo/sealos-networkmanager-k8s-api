package encoding

import (
	"log"
	"math/rand"
	"testing"

	taglib "github.com/dinoallo/sealos-networkmanager-k8s-api/pkg/tag"
	"github.com/go-faker/faker/v4"
	"github.com/matryer/is"
)

type Addressing struct {
	IPV4 string `faker:"ipv4"`
	IPV6 string `faker:"ipv6"`
}

func TestIPV4(t *testing.T) {
	addressing := Addressing{}
	at := AT{}
	is := is.New(t)
	var code *ATCode
	t.Run("generate a random AT with ipv4", func(t *testing.T) {
		err := faker.FakeData(&addressing)
		is.NoErr(err)
		at.IPAddr = addressing.IPV4
		at.Tag = taglib.TagID(rand.Uint32())
		log.Printf("this original ipv4 address is %v, and the tag is %v", at.IPAddr, at.Tag)
	})
	t.Run("encode this AT", func(t *testing.T) {
		var err error
		code, err = at.Encode()
		is.NoErr(err)
		log.Printf("the code is %v, and its len is %v", *code, len(*code))
	})
	t.Run("decode this code", func(t *testing.T) {
		_at, err := code.Decode()
		is.NoErr(err)
		log.Printf("the decoded ipv4 addrress is: %v, and the tag is %v", _at.IPAddr, _at.Tag)
	})
}

func TestIPV6(t *testing.T) {
	addressing := Addressing{}
	at := AT{}
	is := is.New(t)
	var code *ATCode
	t.Run("generate a random AT with ipv6", func(t *testing.T) {
		err := faker.FakeData(&addressing)
		is.NoErr(err)
		at.IPAddr = addressing.IPV6
		at.Tag = taglib.TagID(rand.Uint32())
		log.Printf("this original ipv6 address is %v, and the tag is %v", at.IPAddr, at.Tag)
	})
	t.Run("encode this AT", func(t *testing.T) {
		var err error
		code, err = at.Encode()
		is.NoErr(err)
		log.Printf("the code is %v, and its len is %v", *code, len(*code))
	})
	t.Run("decode this code", func(t *testing.T) {
		_at, err := code.Decode()
		is.NoErr(err)
		log.Printf("the decoded ipv6 addrress is: %v, and the tag is %v", _at.IPAddr, _at.Tag)
	})
}
