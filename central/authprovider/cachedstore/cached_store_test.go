package cachedstore

import (
	"fmt"
	"sync"
	"testing"

	// This import is required to register auth providers.
	_ "bitbucket.org/stack-rox/apollo/pkg/authproviders/all"

	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

type mockStore struct {
	authProviders map[string]*v1.AuthProvider
}

func (m *mockStore) GetAuthProvider(id string) (*v1.AuthProvider, bool, error) {
	authProvider, exists := m.authProviders[id]
	return authProvider, exists, nil
}

func (m *mockStore) GetAuthProviders(request *v1.GetAuthProvidersRequest) ([]*v1.AuthProvider, error) {
	providers := make([]*v1.AuthProvider, 0)
	for _, authProvider := range m.authProviders {
		providers = append(providers, authProvider)
	}
	return providers, nil
}

func (m *mockStore) AddAuthProvider(authProvider *v1.AuthProvider) (string, error) {
	id := uuid.NewV4().String()
	authProvider.Id = id
	m.authProviders[id] = authProvider
	return id, nil
}

func (m *mockStore) UpdateAuthProvider(authProvider *v1.AuthProvider) error {
	if authProvider.GetId() == "" {
		panic(fmt.Sprintf("No id for auth provider %#v", authProvider))
	}
	m.authProviders[authProvider.GetId()] = authProvider
	return nil
}

func (m *mockStore) RemoveAuthProvider(id string) error {
	delete(m.authProviders, id)
	return nil
}

func (m *mockStore) RecordAuthSuccess(id string) error {
	m.authProviders[id].Validated = true
	return nil
}

func newMockStoreWithProviders(authProviders map[string]*v1.AuthProvider) *mockStore {
	return &mockStore{
		authProviders: authProviders,
	}
}

func newMockStore() *mockStore {
	return newMockStoreWithProviders(make(map[string]*v1.AuthProvider))
}

var fakeAuthProvider = &v1.AuthProvider{
	Type: "auth0",
	Config: map[string]string{
		"domain":    "blah.com",
		"client_id": "FAKE",
	},
}

func newCachedStoreWithFakeProvider(t *testing.T) (cachedStore CachedStore, providerID string) {
	cachedStore = New(newMockStore())
	fakeAuthProviderCopy := new(v1.AuthProvider)
	*fakeAuthProviderCopy = *fakeAuthProvider
	providerID, err := cachedStore.AddAuthProvider(fakeAuthProviderCopy)
	assert.NoError(t, err)
	assert.NotEmpty(t, providerID)
	return
}

func TestCachedStoreCanAddAndRetrieve(t *testing.T) {
	cachedStore, id := newCachedStoreWithFakeProvider(t)

	authenticators := cachedStore.GetAuthenticators()
	assert.Len(t, authenticators, 1)
	got, ok := authenticators[id]
	assert.True(t, ok)
	assert.False(t, got.Enabled())
	assert.False(t, got.Validated())
}

func TestCachedStoreUpdatesWhenAuthSuccessIsRecorded(t *testing.T) {
	cachedStore, id := newCachedStoreWithFakeProvider(t)

	cachedStore.RecordAuthSuccess(id)
	authenticators := cachedStore.GetAuthenticators()
	assert.Len(t, authenticators, 1)
	got, ok := authenticators[id]
	assert.True(t, ok)
	assert.False(t, got.Enabled())

	// This is the crucial assertion.
	assert.True(t, got.Validated())
}

func TestCacheIsRefreshedToBeginWith(t *testing.T) {
	const fakeID = "FAKEID"
	fakeAuthProviderWithID := new(v1.AuthProvider)
	*fakeAuthProviderWithID = *fakeAuthProvider
	fakeAuthProviderWithID.Id = fakeID
	fakeProviders := map[string]*v1.AuthProvider{
		fakeID: fakeAuthProviderWithID,
	}
	cachedStore := New(newMockStoreWithProviders(fakeProviders))
	authenticators := cachedStore.GetAuthenticators()
	if !assert.Len(t, authenticators, 1) {
		t.Fatal("Didn't find exactly one authenticator")
	}
	got, ok := authenticators[fakeID]
	if !assert.True(t, ok) {
		t.Fatalf("Couldn't find authenticator with %s", fakeID)
	}
	assert.False(t, got.Enabled())
	assert.False(t, got.Validated())
}

func TestCachedStoreHandlesRace(t *testing.T) {
	cachedStore := New(newMockStore())
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(3)
		go func() {
			defer wg.Done()
			fakeAuthProviderCopy := new(v1.AuthProvider)
			*fakeAuthProviderCopy = *fakeAuthProvider
			cachedStore.AddAuthProvider(fakeAuthProviderCopy)
		}()
		go func() {
			defer wg.Done()
			cachedStore.GetAuthenticators()
		}()
		go func() {
			defer wg.Done()
			cachedStore.RefreshCache()
		}()
	}
	wg.Wait()
	// No assertions here, this test is just to make sure we don't hit a concurrent map access or something.
}
