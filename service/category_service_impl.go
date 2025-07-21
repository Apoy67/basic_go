package service

import (
	"basic-restfull-golang/exception"
	"basic-restfull-golang/halper"
	"basic-restfull-golang/model/domain"
	"basic-restfull-golang/model/web"
	"basic-restfull-golang/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, req web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(req)
	halper.PanicIfError(err)

	tx, err := service.DB.Begin() // jika menggunakan transaction harus implement ini
	halper.PanicIfError(err)
	defer halper.CommitOrRollbeck(tx)

	category := domain.Category{
		Name: req.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return halper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, req web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(req)
	halper.PanicIfError(err)

	tx, err := service.DB.Begin() // jika menggunakan transaction harus implement ini
	halper.PanicIfError(err)
	defer halper.CommitOrRollbeck(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = req.Name

	category = service.CategoryRepository.Update(ctx, tx, category)
	return halper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin() // jika menggunakan transaction harus implement ini
	halper.PanicIfError(err)
	defer halper.CommitOrRollbeck(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin() // jika menggunakan transaction harus implement ini
	halper.PanicIfError(err)
	defer halper.CommitOrRollbeck(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return halper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin() // jika menggunakan transaction harus implement ini
	halper.PanicIfError(err)
	defer halper.CommitOrRollbeck(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return halper.ToCategoryResponses(categories)
}
