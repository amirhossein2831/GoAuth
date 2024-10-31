package ctx

import (
	"context"
)

type CTX interface {
	Set(key string, value any) *Ctx
	SetMap(key, mapKey string, value any) *Ctx
	Get(key string) any
	GetMap(key string) map[string]any
	GetContext() context.Context
}

type Ctx struct {
	Context context.Context
}

func New() *Ctx {
	return &Ctx{
		Context: context.Background(),
	}
}

func (c *Ctx) Set(key string, value any) *Ctx {
	c.Context = context.WithValue(c.Context, key, value)
	return c
}

func (c *Ctx) SetMap(key, mapKey string, value any) *Ctx {
	existingMap, ok := c.Get(key).(map[string]any)
	if existingMap == nil || !ok {
		existingMap = make(map[string]any)
	}
	if mapKey == "" && value == nil {
		c.Context = context.WithValue(c.Context, key, existingMap)
		return c
	}
	existingMap[mapKey] = value
	c.Context = context.WithValue(c.Context, key, existingMap)
	return c
}

func (c *Ctx) Get(key string) any {
	return c.Context.Value(key)
}

func (c *Ctx) GetMap(key string) map[string]any {
	return c.Context.Value(key).(map[string]any)
}

func (c *Ctx) GetContext() context.Context {
	return c.Context
}
