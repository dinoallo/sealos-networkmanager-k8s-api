package tag

import "fmt"

type TagID uint32

const (
	// The first 65536 TagIDs are reserved for port tag with number 0~65535
	TIDPortZero TagID = iota
)

const (
	// IdWorld represents any endpoint outside of the cluster
	TIDIdentityWorld TagID = iota + 65536
)

var (
	TagIdentityWorld = Tag{
		TID:    TIDIdentityWorld,
		String: "identity:world",
	}
)

type Tag struct {
	TID    TagID
	String string
}

func GetTagPortN(n uint32) *Tag {
	return &Tag{
		TID:    TIDPortZero + TagID(n),
		String: fmt.Sprintf("port:%v", n),
	}
}
