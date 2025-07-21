package repository

import (
	"basic-restfull-golang/halper"
	"basic-restfull-golang/model/domain"
	"context"
	"database/sql"
	"errors"
	// "errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

// Save implements CategoryRepository
func (r *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	halper.PanicIfError(err)

	id, err := result.LastInsertId()
	halper.PanicIfError(err)

	category.Id = int(id)
	return category
}

// Update implements CategoryRepository
func (r *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	halper.PanicIfError(err)

	return category
}

// Delete implements CategoryRepository
func (r *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	halper.PanicIfError(err)
}

// FindById implements CategoryRepository
func (r *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	halper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		halper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}

// FindAll implements CategoryRepository
func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	halper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		halper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
