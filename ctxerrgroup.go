package ctxerrgroup

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Group struct {
	*errgroup.Group

	ctx context.Context
}

func WithContext(ctx context.Context) *Group {
	g, ctx := errgroup.WithContext(ctx)

	return &Group{
		Group: g,
		ctx:   ctx,
	}
}

func (g *Group) Context() context.Context {
	return g.ctx
}

func (g *Group) GoWithContext(f func(context.Context) error) {
	g.Group.Go(func() error {
		return f(g.ctx)
	})
}
