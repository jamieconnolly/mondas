package mondas

type Context struct {
	App  *App
	Args Args
}

func NewContext(app *App, args Args) *Context {
	return &Context{
		App:  app,
		Args: args,
	}
}
