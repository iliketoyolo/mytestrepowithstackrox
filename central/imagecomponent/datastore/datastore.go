package datastore

import (
	"context"

	"github.com/stackrox/rox/central/imagecomponent/index"
	"github.com/stackrox/rox/central/imagecomponent/search"
	"github.com/stackrox/rox/central/imagecomponent/store"
	"github.com/stackrox/rox/central/ranking"
	riskDataStore "github.com/stackrox/rox/central/risk/datastore"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	searchPkg "github.com/stackrox/rox/pkg/search"
)

// DataStore is an intermediary to ImageComponent storage.
//go:generate mockgen-wrapper DataStore
type DataStore interface {
	Search(ctx context.Context, q *v1.Query) ([]searchPkg.Result, error)
	SearchImageComponents(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error)
	SearchRawImageComponents(ctx context.Context, q *v1.Query) ([]*storage.ImageComponent, error)

	Exists(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*storage.ImageComponent, bool, error)
	Count(ctx context.Context) (int, error)
	GetBatch(ctx context.Context, id []string) ([]*storage.ImageComponent, error)

	Upsert(ctx context.Context, imagecomponents ...*storage.ImageComponent) error
	Delete(ctx context.Context, ids ...string) error
}

// New returns a new instance of a DataStore.
func New(storage store.Store, indexer index.Indexer, searcher search.Searcher, risks riskDataStore.DataStore, ranker *ranking.Ranker) (DataStore, error) {
	ds := &datastoreImpl{
		storage:              storage,
		indexer:              indexer,
		searcher:             searcher,
		risks:                risks,
		imageComponentRanker: ranker,
	}

	ds.initializeRankers()
	return ds, nil
}
