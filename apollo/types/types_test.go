package types

import (
	"testing"

	"bitbucket.org/stack-rox/apollo/pkg/api/generated/api/v1"
	"github.com/docker/docker/reference"
	"github.com/stretchr/testify/assert"
)

func TestNewImage(t *testing.T) {
	image := &v1.Image{
		Registry: reference.DefaultHostname,
		Remote:   "library/nginx",
		Tag:      "latest",
		Sha:      "adea4f68096fded167603ba6663ed615a80e090da68eb3c9e2508c15c8368401",
	}
	newImage := GenerateImageFromString("nginx:latest@sha256:adea4f68096fded167603ba6663ed615a80e090da68eb3c9e2508c15c8368401")
	assert.Equal(t, image, newImage)
}
