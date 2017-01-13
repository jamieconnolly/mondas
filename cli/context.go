package cli

type Context struct {
	App  *App
	Args Args
	Env  Env
}

func NewContext(app *App, args []string, env []string) *Context {
	return &Context{
		App:  app,
		Args: Args(args),
		Env:  NewEnvFromEnviron(env),
	}
}
