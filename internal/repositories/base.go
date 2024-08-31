package repositories

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func (r *BaseRepository[T]) Create(ctx context.Context, model *T) error {
	return r.DB.Create(model).Error
}

func (r *BaseRepository[T]) Update(ctx context.Context, model *T) error {
	return r.DB.Save(model).Error
}

func (r *BaseRepository[T]) Delete(ctx context.Context, model *T) error {
	return r.DB.Delete(model).Error
}

func (r *BaseRepository[T]) GetById(ctx context.Context, model *T, id any) error {
	return r.DB.Where("id = ?", id).Take(model).Error
}

func (r *BaseRepository[T]) GetAll(ctx context.Context, model *T, id any) error {
	return r.DB.Take(model).Error
}
