package merkledag

import (
	"context"

	cid "github.com/dms3-fs/go-cid"
	dms3ld "github.com/dms3-fs/go-ld-format"
)

// ErrorService implements dms3ld.DAGService, returning 'Err' for every call.
type ErrorService struct {
	Err error
}

var _ dms3ld.DAGService = (*ErrorService)(nil)

// Add returns the cs.Err.
func (cs *ErrorService) Add(ctx context.Context, nd dms3ld.Node) error {
	return cs.Err
}

// AddMany returns the cs.Err.
func (cs *ErrorService) AddMany(ctx context.Context, nds []dms3ld.Node) error {
	return cs.Err
}

// Get returns the cs.Err.
func (cs *ErrorService) Get(ctx context.Context, c *cid.Cid) (dms3ld.Node, error) {
	return nil, cs.Err
}

// GetMany many returns the cs.Err.
func (cs *ErrorService) GetMany(ctx context.Context, cids []*cid.Cid) <-chan *dms3ld.NodeOption {
	ch := make(chan *dms3ld.NodeOption)
	close(ch)
	return ch
}

// Remove returns the cs.Err.
func (cs *ErrorService) Remove(ctx context.Context, c *cid.Cid) error {
	return cs.Err
}

// RemoveMany returns the cs.Err.
func (cs *ErrorService) RemoveMany(ctx context.Context, cids []*cid.Cid) error {
	return cs.Err
}
