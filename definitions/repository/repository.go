package repository

import (
	"context"

	"github.com/Nemutagk/godb/v2/definitions/models"
)

type DriverConnection[T any] interface {
	GetTableName() string
	GetOrderColumns() map[string]string
	Get(ctx context.Context, filters models.GroupFilter, opts *models.Options) ([]T, error)
	GetOne(ctx context.Context, filters models.GroupFilter) (T, error)
	Create(ctx context.Context, data map[string]any) (T, error)
	Update(ctx context.Context, filters models.GroupFilter, data map[string]any) (T, error)
	Delete(ctx context.Context, filters models.GroupFilter) error
	Count(ctx context.Context, filters models.GroupFilter) (int64, error)
}

type Repository[T any] struct {
	Driver      DriverConnection[T]
	Table       string
	OrderColumn map[string]string
	Relations   []any
	SoftDelete  *string
}

func NewRepository[T any](driver DriverConnection[T], relations []any, softDelete *string) Repository[T] {
	return Repository[T]{
		Driver:      driver,
		Table:       driver.GetTableName(),
		OrderColumn: driver.GetOrderColumns(),
		SoftDelete:  softDelete,
		Relations:   relations,
	}
}

func (r *Repository[T]) Get(ctx context.Context, filters models.GroupFilter, opts *models.Options) ([]T, error) {
	result, err := r.Driver.Get(ctx, filters, opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository[T]) GetOne(ctx context.Context, filters models.GroupFilter) (T, error) {
	var zero T
	result, err := r.Driver.GetOne(ctx, filters)
	if err != nil {
		return zero, err
	}

	return result, nil
}

func (r *Repository[T]) Create(ctx context.Context, data map[string]any) (T, error) {
	var zero T
	result, err := r.Driver.Create(ctx, data)
	if err != nil {
		return zero, err
	}

	return result, nil
}

func (r *Repository[T]) Update(ctx context.Context, filters models.GroupFilter, data map[string]any) (T, error) {
	var zero T
	result, err := r.Driver.Update(ctx, filters, data)
	if err != nil {
		return zero, err
	}

	return result, nil
}

func (r *Repository[T]) Delete(ctx context.Context, filters models.GroupFilter) error {
	err := r.Driver.Delete(ctx, filters)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository[T]) Count(ctx context.Context, filters models.GroupFilter) (int64, error) {
	count, err := r.Driver.Count(ctx, filters)
	if err != nil {
		return 0, err
	}

	return count, nil
}
