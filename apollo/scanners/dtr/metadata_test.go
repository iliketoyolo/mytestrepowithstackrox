package dtr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var metadataPayload = `{
  "state": 0,
  "scanner_version": 3,
  "scanner_updated_at": "2017-11-16T21:07:18.934766247Z",
  "db_version": 279,
  "db_updated_at": "2017-11-17T03:14:02.63437292Z",
  "last_db_update_failed": true,
  "replicas": {
   "d8ae913ef3a1": {
    "db_updated_at": "2017-11-16T00:35:27.408476Z",
    "version": "279",
    "replica_id": "d8ae913ef3a1"
   }
  }
 }`

func getExpectedMetadata() (*ScannerMetadata, error) {
	scannerUpdate, err := time.Parse(time.RFC3339Nano, "2017-11-16T21:07:18.934766247Z")
	if err != nil {
		return nil, err
	}
	dbUpdate, err := time.Parse(time.RFC3339Nano, "2017-11-17T03:14:02.63437292Z")
	if err != nil {
		return nil, err
	}
	replicaUpdate, err := time.Parse(time.RFC3339Nano, "2017-11-16T00:35:27.408476Z")
	if err != nil {
		return nil, err
	}

	return &ScannerMetadata{
		State:              0,
		ScannerVersion:     3,
		ScannerUpdatedAt:   scannerUpdate,
		DBVersion:          279,
		DBUpdatedAt:        dbUpdate,
		LastDBUpdateFailed: true,
		Replicas: map[string]Replica{
			"d8ae913ef3a1": {
				DBUpdatedAt: replicaUpdate,
				Version:     "279",
				ReplicaID:   "d8ae913ef3a1",
			},
		},
	}, nil
}

func TestParseMetadata(t *testing.T) {
	meta, err := parseMetadata([]byte(metadataPayload))
	require.Nil(t, err)

	expectedMeta, err := getExpectedMetadata()
	require.Nil(t, err)
	assert.Equal(t, expectedMeta, meta)
}
