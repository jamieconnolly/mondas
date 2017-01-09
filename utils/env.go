package utils

import (
	"fmt"
	"strings"
)

type Env map[string]string

func NewEnvFromEnviron(environ []string) Env {
	env := make(Env)
	for _, variable := range environ {
		parts := strings.SplitN(variable, "=", 2)
		if len(parts) > 1 {
			env[parts[0]] = parts[1]
		} else {
			env[parts[0]] = ""
		}
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
	delete(e, key)
}
