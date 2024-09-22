package consts

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/version"
)

const (
	HRP  = "tutorial"
	Name = "tutorialvm"
)

var ID ids.ID

func init() {
	b := make([]byte, ids.IDLen)
	copy(b, []byte(Name))
	vmID, err := ids.ToID(b)
	if err != nil {
		panic(err)
	}
	ID = vmID
}

var Version = &version.Semantic{
	Major: 0,
	Minor: 0,
	Patch: 1,
}
