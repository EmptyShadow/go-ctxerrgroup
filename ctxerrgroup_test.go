package ctxerrgroup_test

import (
	"context"
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/EmptyShadow/go-ctxerrgroup"
)

func Test_GoWithContext(t *testing.T) {
	g := ctxerrgroup.WithContext(context.Background())

	f := func(ctx context.Context) error {
		timeout := time.Millisecond * time.Duration(rand.Int()%100)

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		<-ctx.Done()

		return ctx.Err()
	}

	g.GoWithContext(f)
	g.GoWithContext(f)
	g.GoWithContext(f)
	g.GoWithContext(f)
	g.GoWithContext(f)

	if err := g.Wait(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err '%s' is not '%s'", err, context.DeadlineExceeded)
	}
}
