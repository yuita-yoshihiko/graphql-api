package dataloader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader/v7"
)

type Dataloader[T comparable, U any] interface {
	NewInterface() dataloader.Interface[T, U]
	BatchFunc() dataloader.BatchFunc[T, U]
}

func DataloaderFromCtx[T, U any](ctx context.Context, key T) U {
	l, ok := ctx.Value(key).(U)
	if !ok {
		panic(fmt.Sprintf("no loader for key %v", key))
	}
	return l
}

func MakeErrorResults[T comparable, U any](keys []T, err error) []*dataloader.Result[U] {
	results := make([]*dataloader.Result[U], len(keys))
	for i := range results {
		results[i] = &dataloader.Result[U]{
			Error: err,
		}
	}
	return results
}

func SortResults[T comparable, U any](keys []T, m map[T]U, nullable bool) []*dataloader.Result[U] {
	results := make([]*dataloader.Result[U], len(keys))
	for i, key := range keys {
		obj, ok := m[key]
		if !ok && !nullable {
			results[i] = &dataloader.Result[U]{
				Error: fmt.Errorf("data not found for key \"%v\"", key),
			}
			continue
		}
		results[i] = &dataloader.Result[U]{
			Data: obj,
		}
	}
	return results
}
