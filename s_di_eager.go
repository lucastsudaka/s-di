package s_di

// is eager because we build the instance when creating (YOU DONT SAY)
func NewEagerProvider[T any](appDi *Sdi, name string, gProvider GenericProvider[T]) Provider[T] {
	instance, err := gProvider(appDi)
	if err != nil {
		panic(err)
	}
	return Provider[T]{
		name:     name,
		provider: gProvider,
		instance: instance,
	}
}
