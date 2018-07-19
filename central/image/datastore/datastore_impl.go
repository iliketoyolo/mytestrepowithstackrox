package datastore

import (
	"bitbucket.org/stack-rox/apollo/central/image/index"
	"bitbucket.org/stack-rox/apollo/central/image/search"
	"bitbucket.org/stack-rox/apollo/central/image/store"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
)

type datastoreImpl struct {
	storage  store.Store
	indexer  index.Indexer
	searcher search.Searcher
}

func (ds *datastoreImpl) SearchListImages(request *v1.ParsedSearchRequest) ([]*v1.ListImage, error) {
	return ds.searcher.SearchListImages(request)
}

func (ds *datastoreImpl) ListImage(sha string) (*v1.ListImage, bool, error) {
	return ds.storage.ListImage(sha)
}

func (ds *datastoreImpl) ListImages() ([]*v1.ListImage, error) {
	return ds.storage.ListImages()
}

// SearchImages
func (ds *datastoreImpl) SearchImages(request *v1.ParsedSearchRequest) ([]*v1.SearchResult, error) {
	return ds.searcher.SearchImages(request)
}

// SearchRawImages
func (ds *datastoreImpl) SearchRawImages(request *v1.ParsedSearchRequest) ([]*v1.Image, error) {
	return ds.searcher.SearchRawImages(request)
}

// GetImage
func (ds *datastoreImpl) GetImage(id string) (*v1.Image, bool, error) {
	return ds.storage.GetImage(id)
}

// GetImages
func (ds *datastoreImpl) GetImages() ([]*v1.Image, error) {
	return ds.storage.GetImages()
}

// CountImages
func (ds *datastoreImpl) CountImages() (int, error) {
	return ds.storage.CountImages()
}

// AddImage inserts an alert into storage and into the indexer
func (ds *datastoreImpl) AddImage(alert *v1.Image) error {
	if err := ds.storage.AddImage(alert); err != nil {
		return err
	}
	return ds.indexer.AddImage(alert)
}

// UpdateImage updates an alert in storage and in the indexer
func (ds *datastoreImpl) UpdateImage(alert *v1.Image) error {
	if err := ds.storage.UpdateImage(alert); err != nil {
		return err
	}
	return ds.indexer.AddImage(alert)
}
