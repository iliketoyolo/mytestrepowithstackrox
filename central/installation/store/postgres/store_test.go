// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type InstallationInfosStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	pool        *pgxpool.Pool
}

func TestInstallationInfosStore(t *testing.T) {
	suite.Run(t, new(InstallationInfosStoreSuite))
}

func (s *InstallationInfosStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	ctx := sac.WithAllAccess(context.Background())

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.Require().NoError(err)

	Destroy(ctx, pool)

	s.pool = pool
	s.store = New(ctx, pool)
}

func (s *InstallationInfosStoreSuite) TearDownTest() {
	if s.pool != nil {
		s.pool.Close()
	}
	s.envIsolator.RestoreAll()
}

func (s *InstallationInfosStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	installationInfo := &storage.InstallationInfo{}
	s.NoError(testutils.FullInit(installationInfo, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundInstallationInfo, exists, err := store.Get(ctx)
	s.NoError(err)
	s.False(exists)
	s.Nil(foundInstallationInfo)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, installationInfo))
	foundInstallationInfo, exists, err = store.Get(ctx)
	s.NoError(err)
	s.True(exists)
	s.Equal(installationInfo, foundInstallationInfo)

	foundInstallationInfo, exists, err = store.Get(ctx)
	s.NoError(err)
	s.True(exists)
	s.Equal(installationInfo, foundInstallationInfo)

	s.NoError(store.Delete(ctx))
	foundInstallationInfo, exists, err = store.Get(ctx)
	s.NoError(err)
	s.False(exists)
	s.Nil(foundInstallationInfo)

	s.ErrorIs(store.Delete(withNoAccessCtx), sac.ErrResourceAccessDenied)

	installationInfo = &storage.InstallationInfo{}
	s.NoError(testutils.FullInit(installationInfo, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))
	s.NoError(store.Upsert(ctx, installationInfo))

	foundInstallationInfo, exists, err = store.Get(ctx)
	s.NoError(err)
	s.True(exists)
	s.Equal(installationInfo, foundInstallationInfo)

	installationInfo = &storage.InstallationInfo{}
	s.NoError(testutils.FullInit(installationInfo, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))
	s.NoError(store.Upsert(ctx, installationInfo))

	foundInstallationInfo, exists, err = store.Get(ctx)
	s.NoError(err)
	s.True(exists)
	s.Equal(installationInfo, foundInstallationInfo)
}
