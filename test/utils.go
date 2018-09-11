package mdutils

import (
	dag "github.com/dms3-fs/go-merkledag"

	bsrv "github.com/dms3-fs/go-blockservice"
	ds "github.com/dms3-fs/go-datastore"
	dssync "github.com/dms3-fs/go-datastore/sync"
	blockstore "github.com/dms3-fs/go-fs-blockstore"
	offline "github.com/dms3-fs/go-fs-exchange-offline"
	dms3ld "github.com/dms3-fs/go-ld-format"
)

// Mock returns a new thread-safe, mock DAGService.
func Mock() dms3ld.DAGService {
	return dag.NewDAGService(Bserv())
}

// Bserv returns a new, thread-safe, mock BlockService.
func Bserv() bsrv.BlockService {
	bstore := blockstore.NewBlockstore(dssync.MutexWrap(ds.NewMapDatastore()))
	return bsrv.New(bstore, offline.Exchange(bstore))
}
