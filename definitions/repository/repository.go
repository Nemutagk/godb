package repository

import (
	"context"

	"github.com/Nemutagk/godb/v2/definitions/models"
)

type DriverConnection[T any] interface {
	GetTableName() string
	GetOrderColumns() map[string]string
	GetConnection() any
	AddRelation(relationName string, loader RelationLoader) error
	Get(ctx context.Context, filters models.GroupFilter, opts *models.Options) ([]T, error)
	GetOne(ctx context.Context, filters models.GroupFilter, opts *models.Options) (T, error)
	Create(ctx context.Context, data map[string]any, opts *models.Options) (T, error)
	CreateMany(ctx context.Context, data []map[string]any, opts *models.Options) ([]T, error)
	Update(ctx context.Context, filters models.GroupFilter, data map[string]any, opts *models.Options) (T, error)
	Delete(ctx context.Context, filters models.GroupFilter) error
	Count(ctx context.Context, filters models.GroupFilter) (int64, error)
	TransactionStart(ctx context.Context) (*models.Transaction, error)
	TransactionCommit(ctx context.Context, trans *models.Transaction) error
	TransactionRollback(ctx context.Context, trans *models.Transaction) error
}

type RelationModel interface {
	GetTableName() string
}

type Model interface {
	ScanFields() []any
}

type RelationLoader interface {
	Load(ctx context.Context, parentModels []any, childs *[]string) error
}
