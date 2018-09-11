package merkledag

import (
	"context"

	dms3ld "github.com/dms3-fs/go-ld-format"
)

// SessionMaker is an object that can generate a new fetching session.
type SessionMaker interface {
	Session(context.Context) dms3ld.NodeGetter
}

// NewSession returns a session backed NodeGetter if the given NodeGetter
// implements SessionMaker.
func NewSession(ctx context.Context, g dms3ld.NodeGetter) dms3ld.NodeGetter {
	if sm, ok := g.(SessionMaker); ok {
		return sm.Session(ctx)
	}
	return g
}
