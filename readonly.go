package merkledag

import (
	"fmt"

	dms3ld "github.com/dms3-fs/go-ld-format"
)

// ErrReadOnly is used when a read-only datastructure is written to.
var ErrReadOnly = fmt.Errorf("cannot write to readonly DAGService")

// NewReadOnlyDagService takes a NodeGetter, and returns a full DAGService
// implementation that returns ErrReadOnly when its 'write' methods are
// invoked.
func NewReadOnlyDagService(ng dms3ld.NodeGetter) dms3ld.DAGService {
	return &ComboService{
		Read:  ng,
		Write: &ErrorService{ErrReadOnly},
	}
}
