package ctx

import (
	"context"
)

type Ctx struct {
	Context context.Context
}

func New() *Ctx {
	return &Ctx{
		Context: context.Background(),
	}
}

func NewChild(prt context.Context) *Ctx {
	return &Ctx{
		Context: prt,
	}
}

func (c *Ctx) Set(key string, value any) *Ctx {
	c.Context = context.WithValue(c.Context, key, value)
	return c
}

func (c *Ctx) SetMap(key, mapKey string, value any) *Ctx {
	existingMap := c.Get(key).(map[string]any)
	if existingMap == nil {
		existingMap = make(map[string]any)
	}
	existingMap[mapKey] = value
	c.Context = context.WithValue(c.Context, key, existingMap)
	return c
}

func (c *Ctx) Get(key string) any {
	return c.Context.Value(key)
}

func (c *Ctx) GetContext() context.Context {
	return c.Context
}
