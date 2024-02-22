package context

type Context struct {
	store map[string]interface{}
}

func New() Context {
	return Context{
		store: make(map[string]interface{}),
	}
}

func (context Context) Set(name string, value interface{}) {
	context.store[name] = value
}

func (context Context) Get(name string) (interface{}, bool) {
	value, ok := context.store[name]

	return value, ok
}
