package service

import (
	"errors"
	"proj_1/configs"
	"proj_1/internal/domain"
	"proj_1/internal/dto"
	"proj_1/internal/helper"
	"proj_1/internal/repository"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config configs.AppConfig
}

func (s CatalogService) CreateCategory(input dto.CreateCategoryRequest) error {
	err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})

	return err
}

func (s CatalogService) EditCategory(id int, input dto.CreateCategoryRequest) (*domain.Category, error) {
	exitCat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("category does not exist")
	}

	if len(input.Name) > 0 {
		exitCat.Name = input.Name
	}

	if input.ParentId > 0 {
		exitCat.ParentId = input.ParentId
	}

	if len(input.ImageUrl) > 0 {
		exitCat.ImageUrl = input.ImageUrl
	}

	if input.DisplayOrder > 0 {
		exitCat.DisplayOrder = input.DisplayOrder
	}

	updatedCat, err := s.Repo.EditCategory(exitCat)

	return updatedCat, err
}

func (s CatalogService) DeleteCategory(id int) error {
	err := s.Repo.DeleteCategory(id)
	if err != nil {
		//log the error
		return errors.New("category doese not exist to delete")
	}

	return nil
}

func (s CatalogService) Getcategories() ([]*domain.Category, error) {
	categories, err := s.Repo.FindCategories()
	if err != nil {
		return nil, errors.New("categories does not exist")
	}

	return categories, err
}

func (s CatalogService) Getcategory(id int) (*domain.Category, error) {
	cat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("category does not exist")
	}

	return cat, nil
}
