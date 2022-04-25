package context

import (
	goContext "context"

	"portto-homework/pkg/log"
)

type Context struct {
	goContext.Context
	*log.Logger
}

func Background() Context {
	return Context{
		Context: goContext.Background(),
		Logger:  log.New(),
	}
}
