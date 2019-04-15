package datastore

import (
	"testing"

	"github.com/blevesearch/bleve"
	"github.com/stackrox/rox/central/globalindex"
	"github.com/stackrox/rox/central/serviceaccount/index"
	serviceAccountSearch "github.com/stackrox/rox/central/serviceaccount/search"
	"github.com/stackrox/rox/central/serviceaccount/store"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/bolthelper"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/fixtures"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stretchr/testify/suite"
)

func TestServiceAccountDataStore(t *testing.T) {
	suite.Run(t, new(ServiceAccountDataStoreTestSuite))
}

type ServiceAccountDataStoreTestSuite struct {
	suite.Suite

	bleveIndex bleve.Index

	indexer   index.Indexer
	searcher  serviceAccountSearch.Searcher
	storage   store.Store
	datastore DataStore
}

func (suite *ServiceAccountDataStoreTestSuite) SetupSuite() {
	var err error
	suite.bleveIndex, err = globalindex.TempInitializeIndices("")
	suite.Require().NoError(err)

	db, err := bolthelper.NewTemp(suite.T().Name() + ".db")
	suite.Require().NoError(err)

	suite.storage = store.New(db)
	suite.searcher = serviceAccountSearch.New(suite.storage, suite.bleveIndex)
	suite.indexer = index.New(suite.bleveIndex)
	suite.datastore, err = New(suite.storage, suite.indexer, suite.searcher)
	suite.Require().NoError(err)
}

func (suite *ServiceAccountDataStoreTestSuite) TearDownSuite() {
	suite.NoError(suite.bleveIndex.Close())
}

func (suite *ServiceAccountDataStoreTestSuite) assertSearchResults(q *v1.Query, s *storage.ServiceAccount) {
	results, err := suite.datastore.SearchServiceAccounts(q)
	suite.Require().NoError(err)
	if s != nil {
		suite.Len(results, 1)
		suite.Equal(s.GetId(), results[0].GetId())
	} else {
		suite.Len(results, 0)
	}
}

func (suite *ServiceAccountDataStoreTestSuite) TestServiceAccountsDataStore() {
	if !features.K8sRBAC.Enabled() {
		return
	}

	sa := fixtures.GetServiceAccount()
	err := suite.datastore.UpsertServiceAccount(sa)
	suite.Require().NoError(err)

	foundSA, found, err := suite.datastore.GetServiceAccount(sa.GetId())
	suite.Require().NoError(err)
	suite.True(found)
	suite.Equal(sa, foundSA)

	_, found, err = suite.datastore.GetServiceAccount("NONEXISTENT")
	suite.Require().NoError(err)
	suite.False(found)

	validQ := search.NewQueryBuilder().AddStrings(search.Cluster, sa.GetClusterName()).ProtoQuery()
	suite.assertSearchResults(validQ, sa)

	invalidQ := search.NewQueryBuilder().AddStrings(search.Cluster, "NONEXISTENT").ProtoQuery()
	suite.assertSearchResults(invalidQ, nil)

	err = suite.datastore.RemoveServiceAccount(sa.GetId())
	suite.Require().NoError(err)

	_, found, err = suite.datastore.GetServiceAccount(sa.GetId())
	suite.Require().NoError(err)
	suite.False(found)

	suite.assertSearchResults(validQ, nil)
}
