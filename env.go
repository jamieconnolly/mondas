package mondas

import (
	"fmt"
	"strings"
)

type Env map[string]string

func NewEnvFromEnviron(environ []string) Env {
	env := make(Env)
	for _, variable := range environ {
		components := strings.Split(variable, "=")
		env[components[0]] = components[1]
	}
	return env
}

func (e Env) Environ() []string {
	result := make([]string, 0, len(e))
	for key, value := range e {
		result = append(result, fmt.Sprintf("%s=%s", key, value))
	}
	return result
}

func (e Env) Get(key string) string {
	return e[key]
}

func (e Env) Set(key string, value string) {
	e[key] = value
}

func (e Env) Unset(key string) {
	if i, ok := e[key]; ok {
		e[i] = ""
		delete(e, key)
	}
}
