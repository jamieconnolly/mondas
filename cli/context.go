package cli

// Context represents the command line execution context.
type Context struct {
	App  *App
	Args Args
	Env  Env
}

// NewContext creates a new Context object.
func NewContext(app *App, args []string, env []string) *Context {
	return &Context{
		App:  app,
		Args: Args(args),
		Env:  NewEnvFromEnviron(env),
	}
}
