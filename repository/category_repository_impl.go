package repository

import (
	// "4kbr/restful-golang/helper"
	"4kbr/restful-golang/model/domain"
	"context"
	"database/sql"
)

type CategoryRepositoryImpl struct {
}

// repostitory *CategoryRepositoryImpl repository.CategoryRepository
// f *File repository.CategoryRepository
func (f *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (f *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (f *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	panic("not implemented") // TODO: Implement
}

func (f *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (f *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	panic("not implemented") // TODO: Implement
}

//c *CategoryRepositoryImpl package.domain.Category

//repository CategoryRepositoryImpl

// func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
// 	SQL := "insert into customer(name) values (?)"
// 	result, err := tx.ExecContext(ctx, SQL, category.Name)
// 	helper.PanicIfError(err)

// 	id, err := result.LastInsertId()
// 	helper.PanicIfError(err)

// 	category.Id = int(id)
// 	return category
// }

// func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
// 	panic("not implemented") // TODO: Implement
// }

// func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
// 	panic("not implemented") // TODO: Implement
// }

// func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Category {
// 	panic("not implemented") // TODO: Implement
// }

// func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
// 	panic("not implemented") // TODO: Implement
// }
