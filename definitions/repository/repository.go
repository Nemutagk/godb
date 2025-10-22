package repository

import (
	"context"

	"github.com/Nemutagk/godb/v2/definitions/models"
)

type DriverConnection[T any] interface {
	GetTableName() string
	GetOrderColumns() map[string]string
	// AddRelation(relation Relation)
	// GetRelations() []Relation
	// SetLoadRelations(load bool)
	// GetLoadRelations() bool
	Get(ctx context.Context, filters models.GroupFilter, opts *models.Options) ([]T, error)
	GetOne(ctx context.Context, filters models.GroupFilter) (T, error)
	Create(ctx context.Context, data map[string]any) (T, error)
	Update(ctx context.Context, filters models.GroupFilter, data map[string]any) (T, error)
	Delete(ctx context.Context, filters models.GroupFilter) error
	Count(ctx context.Context, filters models.GroupFilter) (int64, error)
}

type RelationModel interface {
	GetTableName() string
}

type Model interface {
	ScanFields() []any
}

type Repository[T Model] struct {
	Driver        DriverConnection[T]
	Table         string
	OrderColumn   map[string]string
	LoadRelations bool
	SoftDelete    *string
}

type Relation struct {
	Model      RelationModel
	Key        string
	ForeignKey string
	Type       string
}

func NewRepository[T Model](driver DriverConnection[T], softDelete *string) Repository[T] {
	return Repository[T]{
		Driver:      driver,
		Table:       driver.GetTableName(),
		OrderColumn: driver.GetOrderColumns(),
		SoftDelete:  softDelete,
	}

}

func (r *Repository[T]) GetTableName() string {
	return r.Table
}

// func (r *Repository[T]) AddRelation(relation Relation) {
// 	r.Driver.AddRelation(relation)
// }

// func (r *Repository[T]) GetRelations() []Relation {
// 	return r.Driver.GetRelations()
// }

// func (r *Repository[T]) SetLoadRelations(load bool) {
// 	r.Driver.SetLoadRelations(load)
// }

// func (r *Repository[T]) GetLoadRelations() bool {
// 	return r.Driver.GetLoadRelations()
// }

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
