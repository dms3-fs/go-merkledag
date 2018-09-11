package merkledag

import (
	"fmt"
	"github.com/dms3-fs/go-block-format"

	cid "github.com/dms3-fs/go-cid"
	u "github.com/dms3-fs/go-fs-util"
	dms3ld "github.com/dms3-fs/go-ld-format"
)

// RawNode represents a node which only contains data.
type RawNode struct {
	blocks.Block
}

// NewRawNode creates a RawNode using the default sha2-256 hash function.
func NewRawNode(data []byte) *RawNode {
	h := u.Hash(data)
	c := cid.NewCidV1(cid.Raw, h)
	blk, _ := blocks.NewBlockWithCid(data, c)

	return &RawNode{blk}
}

// DecodeRawBlock is a block decoder for raw DMS3LD nodes conforming to `node.DecodeBlockFunc`.
func DecodeRawBlock(block blocks.Block) (dms3ld.Node, error) {
	if block.Cid().Type() != cid.Raw {
		return nil, fmt.Errorf("raw nodes cannot be decoded from non-raw blocks: %d", block.Cid().Type())
	}
	// Once you "share" a block, it should be immutable. Therefore, we can just use this block as-is.
	return &RawNode{block}, nil
}

var _ dms3ld.DecodeBlockFunc = DecodeRawBlock

// NewRawNodeWPrefix creates a RawNode using the provided cid builder
func NewRawNodeWPrefix(data []byte, builder cid.Builder) (*RawNode, error) {
	builder = builder.WithCodec(cid.Raw)
	c, err := builder.Sum(data)
	if err != nil {
		return nil, err
	}
	blk, err := blocks.NewBlockWithCid(data, c)
	if err != nil {
		return nil, err
	}
	return &RawNode{blk}, nil
}

// Links returns nil.
func (rn *RawNode) Links() []*dms3ld.Link {
	return nil
}

// ResolveLink returns an error.
func (rn *RawNode) ResolveLink(path []string) (*dms3ld.Link, []string, error) {
	return nil, nil, ErrLinkNotFound
}

// Resolve returns an error.
func (rn *RawNode) Resolve(path []string) (interface{}, []string, error) {
	return nil, nil, ErrLinkNotFound
}

// Tree returns nil.
func (rn *RawNode) Tree(p string, depth int) []string {
	return nil
}

// Copy performs a deep copy of this node and returns it as an dms3ld.Node
func (rn *RawNode) Copy() dms3ld.Node {
	copybuf := make([]byte, len(rn.RawData()))
	copy(copybuf, rn.RawData())
	nblk, err := blocks.NewBlockWithCid(rn.RawData(), rn.Cid())
	if err != nil {
		// programmer error
		panic("failure attempting to clone raw block: " + err.Error())
	}

	return &RawNode{nblk}
}

// Size returns the size of this node
func (rn *RawNode) Size() (uint64, error) {
	return uint64(len(rn.RawData())), nil
}

// Stat returns some Stats about this node.
func (rn *RawNode) Stat() (*dms3ld.NodeStat, error) {
	return &dms3ld.NodeStat{
		CumulativeSize: len(rn.RawData()),
		DataSize:       len(rn.RawData()),
	}, nil
}

var _ dms3ld.Node = (*RawNode)(nil)
