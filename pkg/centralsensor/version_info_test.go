package centralsensor

import (
	"context"
	"reflect"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/stackrox/rox/pkg/version/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersionInfoClientToServer(t *testing.T) {
	const testVersion = "2.5.27.0"
	testutils.SetMainVersion(t, testVersion)
	ctx, err := AppendSensorVersionInfoToContext(context.Background())
	require.NoError(t, err)
	derived, err := deriveSensorVersionInfo(metautils.ExtractOutgoing(ctx))
	require.NoError(t, err)
	require.NotNil(t, derived)
	assert.Equal(t, testVersion, derived.MainVersion)
}

func TestVersionInfoOldSensors(t *testing.T) {
	derived, err := DeriveSensorVersionInfo(context.Background())
	assert.NoError(t, err)
	assert.Nil(t, derived)
}

// This unit test helps enforce that nobody accidentally removes a field from
// the SensorVersionInfo object. This does not (and cannot) protect against someone
// intentionally doing so; it's intended mainly as a helpful reminder.
func TestVersionInfoHasAllOldFields(t *testing.T) {
	type fieldNameWithTag struct {
		fieldName string
		jsonTag   string
	}

	allKnownFields := []fieldNameWithTag{
		{"MainVersion", "mainVersion"},
	}

	var seenFields []fieldNameWithTag

	versionInfoType := reflect.TypeOf(SensorVersionInfo{})
	for i := 0; i < versionInfoType.NumField(); i++ {
		field := versionInfoType.Field(i)
		seenFields = append(seenFields, fieldNameWithTag{field.Name, field.Tag.Get("json")})
	}

	assert.ElementsMatch(t, allKnownFields, seenFields, "We never want to remove old fields from SensorVersionInfo "+
		"so please don't do that. Also, if you added a new field, pls add that to allKnownFields above so that this test "+
		"ensures that nobody removes it later")
}
