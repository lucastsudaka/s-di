package s_di

import (
	"fmt"
)

type Provider[T any] struct {
	name     string
	provider GenericProvider[T]
	instance T
}
type GenericProvider[T any] func(*Sdi) (T, error)

type Sdi struct {
	version   string
	Providers map[string]any
}

func NewSdi() *Sdi {
	p := &Sdi{}
	p.Providers = map[string]interface{}{}
	return p
}

func AddProvider[T any](appDi *Sdi, gProvider GenericProvider[T]) {
	name := generateProviderName[T]()

	appDi.Providers[name] = NewLazyProvider[T](appDi, name, gProvider)
}

func GetProvider[T any](appDi *Sdi) (T, error) {
	name := generateProviderName[T]()
	i := appDi.Providers[name].(Provider[T])
	i.BuildInstance(appDi, name)
	return i.BuildInstance(appDi, name)
}

func AddEagerProvider[T any](appDi *Sdi, gProvider GenericProvider[T]) {
	name := generateProviderName[T]()
	appDi.Providers[name] = NewEagerProvider[T](appDi, name, gProvider)
}

func GetEagerProvider[T any](appDi *Sdi) (T, error) {
	name := generateProviderName[T]()
	i := appDi.Providers[name].(Provider[T])
	return i.instance, nil
}

func generateProviderName[T any]() string {
	var t T
	// struct
	name := fmt.Sprintf("%T", t)
	if name != "<nil>" {
		return name
	}
	// interface
	return fmt.Sprintf("%T", new(T))
}
