package service

import (
	"proj_1/configs"
	"proj_1/internal/helper"
	"proj_1/internal/repository"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config configs.AppConfig
}
