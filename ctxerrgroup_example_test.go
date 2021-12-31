package ctxerrgroup_test

import (
	"context"
	"fmt"
	"time"

	"github.com/EmptyShadow/go-ctxerrgroup"
)

func Example() {
	// Внешний контекст.
	externalCtx := context.Background()

	// Группа для параллельного выполнения функций.
	g := ctxerrgroup.WithContext(externalCtx)

	// Если нужно можно получить внутрений контекст группы.
	internalCtx := g.Context()
	fmt.Println("%+w", internalCtx)

	// Функционал errgroup так же доступен.
	g.Go(func() error {
		timeout := time.Millisecond * 500

		<-time.NewTimer(timeout).C

		fmt.Println("Go without context")

		return nil
	})

	// Запуск функции которая получит на вход внутрений контекс исполнения группы.
	g.GoWithContext(func(ctx context.Context) error {
		timeout := time.Millisecond * 300

		select {
		// обязана обрабатывать контекст.
		case <-ctx.Done():
			// если контект завершился, то обязана вернуть ошибку из контекста.
			return ctx.Err()
		case <-time.NewTimer(timeout).C:
			fmt.Println("Go with context")

			return nil
		}
	})

	// ожидаем завершения функций
	if err := g.Wait(); err != nil {
		panic(err)
	}
}
