package mondas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs_First(t *testing.T) {
	args := Args{"foo", "bar", "baz", "foo"}
	assert.Equal(t, "foo", args.First())
}

func TestArgs_Index(t *testing.T) {
	args := Args{"foo", "bar", "baz", "foo"}
	assert.Equal(t, "", args.Index(-1))
	assert.Equal(t, "foo", args.Index(0))
	assert.Equal(t, "bar", args.Index(1))
	assert.Equal(t, "baz", args.Index(2))
	assert.Equal(t, "foo", args.Index(3))
	assert.Equal(t, "", args.Index(4))
}

func TestArgs_Len(t *testing.T) {
	args := Args{"foo", "bar", "baz", "foo"}
	assert.Equal(t, 4, args.Len())
}
