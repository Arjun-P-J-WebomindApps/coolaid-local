// internal/repository/category/adapter.go
package category

import "context"

type adapter[T any] struct {
	list func(ctx context.Context) ([]T, error)
}

func (a adapter[T]) List(ctx context.Context) ([]Row, error) {
	items, err := a.list(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]Row, len(items))
	for i := range items {
		out[i] = items[i]
	}
	return out, nil
}
