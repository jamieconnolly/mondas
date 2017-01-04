package mondas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnvFromEnviron(t *testing.T) {
	environ := []string{"foo=bar", "baz=foo"}

	env := NewEnvFromEnviron(environ)
	assert.Equal(t, 2, len(env))
	assert.Equal(t, "bar", env["foo"])
	assert.Equal(t, "foo", env["baz"])
}

func TestEnv_Environ(t *testing.T) {
	env := Env{"foo": "bar", "baz": "foo"}

	environ := env.Environ()
	assert.Equal(t, 2, len(environ))
	assert.Equal(t, "foo=bar", environ[0])
	assert.Equal(t, "baz=foo", environ[1])
}

func TestEnv_Get(t *testing.T) {
	env := Env{"foo": "bar", "baz": "foo"}

	assert.Equal(t, "bar", env.Get("foo"))
	assert.Equal(t, "foo", env.Get("baz"))
}

func TestEnv_Set(t *testing.T) {
	env := Env{"foo": "bar"}

	env.Set("baz", "foo")
	assert.Equal(t, "bar", env["foo"])
	assert.Equal(t, "foo", env["baz"])
}

func TestEnv_Unset(t *testing.T) {
	env := Env{"foo": "bar", "baz": "foo"}

	env.Unset("baz")
	assert.Equal(t, "bar", env["foo"])
	assert.Equal(t, "", env["baz"])
}
