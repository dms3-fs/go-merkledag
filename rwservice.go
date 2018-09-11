package merkledag

import (
	"context"

	cid "github.com/dms3-fs/go-cid"
	dms3ld "github.com/dms3-fs/go-ld-format"
)

// ComboService implements dms3ld.DAGService, using 'Read' for all fetch methods,
// and 'Write' for all methods that add new objects.
type ComboService struct {
	Read  dms3ld.NodeGetter
	Write dms3ld.DAGService
}

var _ dms3ld.DAGService = (*ComboService)(nil)

// Add writes a new node using the Write DAGService.
func (cs *ComboService) Add(ctx context.Context, nd dms3ld.Node) error {
	return cs.Write.Add(ctx, nd)
}

// AddMany adds nodes using the Write DAGService.
func (cs *ComboService) AddMany(ctx context.Context, nds []dms3ld.Node) error {
	return cs.Write.AddMany(ctx, nds)
}

// Get fetches a node using the Read DAGService.
func (cs *ComboService) Get(ctx context.Context, c *cid.Cid) (dms3ld.Node, error) {
	return cs.Read.Get(ctx, c)
}

// GetMany fetches nodes using the Read DAGService.
func (cs *ComboService) GetMany(ctx context.Context, cids []*cid.Cid) <-chan *dms3ld.NodeOption {
	return cs.Read.GetMany(ctx, cids)
}

// Remove deletes a node using the Write DAGService.
func (cs *ComboService) Remove(ctx context.Context, c *cid.Cid) error {
	return cs.Write.Remove(ctx, c)
}

// RemoveMany deletes nodes using the Write DAGService.
func (cs *ComboService) RemoveMany(ctx context.Context, cids []*cid.Cid) error {
	return cs.Write.RemoveMany(ctx, cids)
}
