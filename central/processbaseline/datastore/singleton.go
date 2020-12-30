package datastore

import (
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/globalindex"
	"github.com/stackrox/rox/central/processbaseline/index"
	"github.com/stackrox/rox/central/processbaseline/search"
	"github.com/stackrox/rox/central/processbaseline/store/rocksdb"
	"github.com/stackrox/rox/central/processbaselineresults/datastore"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	once sync.Once

	ad DataStore

	log = logging.LoggerForModule()
)

func initialize() {
	storage, err := rocksdb.New(globaldb.GetRocksDB())
	utils.Must(err)

	indexer := index.New(globalindex.GetGlobalTmpIndex())

	searcher, err := search.New(storage, indexer)
	if err != nil {
		panic("unable to load search index for process baseline")
	}

	ad = New(storage, indexer, searcher, datastore.Singleton())
}

// Singleton provides the interface for non-service external interaction.
func Singleton() DataStore {
	once.Do(initialize)
	return ad
}
