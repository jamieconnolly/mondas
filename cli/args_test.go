package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs_First(t *testing.T) {
	args := Args{"foo", "bar", "baz"}
	assert.Equal(t, "foo", args.First())
}

func TestArgs_Index(t *testing.T) {
	args := Args{"foo", "bar", "baz"}
	assert.Equal(t, "foo", args.Index(0))
	assert.Equal(t, "bar", args.Index(1))
	assert.Equal(t, "baz", args.Index(2))
	assert.Equal(t, "", args.Index(-1))
	assert.Equal(t, "", args.Index(3))
}

func TestArgs_Len(t *testing.T) {
	args := Args{"foo", "bar", "baz"}
	assert.Equal(t, 3, args.Len())
}
