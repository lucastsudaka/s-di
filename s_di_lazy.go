package s_di

// is lazy because we build the instance when invoke (YOU DONT SAY)
func NewLazyProvider[T any](appDi *Sdi, name string, gProvider GenericProvider[T]) Provider[T] {
	//instance, err := gProvider(appDi)
	//if err != nil {
	//	panic(err)
	//}
	return Provider[T]{
		name:     name,
		provider: gProvider,
		//instance: instance,
	}
}

func (l *Provider[T]) BuildInstance(appDi *Sdi, name string) (T, error) {
	p := appDi.Providers[name].(Provider[T])
	return p.provider(appDi)
}
