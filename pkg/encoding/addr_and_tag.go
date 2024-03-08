package encoding

import (
	"fmt"
	"net/netip"

	taglib "github.com/dinoallo/sealos-networkmanager-k8s-api/pkg/tag"
	"github.com/sqids/sqids-go"
)

const (
	sqidAlphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
)

type ATCode string

type AT struct {
	IPAddr string
	Tag    taglib.TagID
}

func (code *ATCode) String() string {
	return string(*code)
}

func (code *ATCode) Decode() (*AT, error) {
	id := string(*code)
	s, err := sqids.New(sqids.Options{Alphabet: sqidAlphabet})
	if err != nil {
		return nil, err
	}
	array := s.Decode(id)
	arrayLen := len(array)
	if len(array) < 2 {
		return nil, fmt.Errorf("unable to decode and this shouldn't happen")
	}
	var at AT
	tag := array[arrayLen-1]
	at.Tag = taglib.TagID(tag)
	_addr := array[0 : arrayLen-1]
	addr := DecodeUint64SliceToIP(_addr)
	if addr == nil {
		return nil, fmt.Errorf("unable to decode and this shouldn't happen")
	}
	at.IPAddr = addr.String()
	return &at, nil
}

func (at *AT) Encode() (*ATCode, error) {
	_ipAddr, err := netip.ParseAddr(at.IPAddr)
	if err != nil {
		return nil, err
	}
	slice := EncodeIPToUint64Slice(_ipAddr)
	if slice == nil {
		return nil, fmt.Errorf("unable to encode and this shouldn't happen")
	}
	slice = append(slice, uint64(at.Tag))
	s, err := sqids.New(sqids.Options{Alphabet: sqidAlphabet})
	if err != nil {
		return nil, err
	}
	id, err := s.Encode(slice)
	if err != nil {
		return nil, err
	}
	code := ATCode(id)
	return &code, nil
}

func EncodeIPToUint64Slice(addr netip.Addr) []uint64 {
	if addr.Is4() {
		return encodeIPV4ToUint64Slice(addr)
	} else if addr.Is6() {
		return encodeIPV6ToUint64Slice(addr)
	}
	return nil
}

func encodeIPV4ToUint64Slice(addr netip.Addr) []uint64 {
	var out4 uint64
	slice := addr.AsSlice()
	for i, x := range slice {
		out4 |= uint64(x) << uint64(i*8)
	}
	return []uint64{out4}
}

func decodeUint64SliceToIPV4(slice []uint64) *netip.Addr {
	out := slice[0]
	array := [4]byte{}
	var mask uint64 = (1 << 8) - 1
	for i := 0; i < 4; i++ {
		array[i] = byte(out & mask)
		out >>= 8
	}
	addr := netip.AddrFrom4(array)
	return &addr
}

func encodeIPV6ToUint64Slice(addr netip.Addr) []uint64 {
	var out6 [2]uint64
	slice := addr.AsSlice()
	for i := 0; i < 2; i++ {
		var out uint64
		for j := 0; j < 8; j++ {
			out |= uint64(slice[i*8+j]) << uint64(j*8)
		}
		out6[i] = out
	}
	return out6[:]
}

func decodeUint64SliceToIPV6(slice []uint64) *netip.Addr {
	array := [16]byte{}
	var mask uint64 = (1 << 8) - 1
	for i := 0; i < 2; i++ {
		out := slice[i]
		for j := 0; j < 8; j++ {
			array[i*8+j] = byte(out & mask)
			out >>= 8
		}
	}
	addr := netip.AddrFrom16(array)
	return &addr
}

func DecodeUint64SliceToIP(slice []uint64) *netip.Addr {
	if len(slice) == 1 {
		return decodeUint64SliceToIPV4(slice)
	} else if len(slice) == 2 {
		return decodeUint64SliceToIPV6(slice)
	} else {
		return nil
	}

}
