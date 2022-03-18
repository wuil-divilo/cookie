package service

//go:generate mockgen -source ./{{cookiecutter.model_name}}.go -destination ../testing/service/mock/{{cookiecutter.model_name}}.go -package mock

import (
	"context"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/repository"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const (
	tagValidationCreate = "validateCreate"
)

// {{cookiecutter.model_name.capitalize()}}Service {{cookiecutter.model_name.capitalize()}} Service Interface
type {{cookiecutter.model_name.capitalize()}}Service interface {
	Create(ctx context.Context, {{cookiecutter.model_name}} model.{{cookiecutter.model_name.capitalize()}}) (model.{{cookiecutter.model_name.capitalize()}}, error)
}

type {{cookiecutter.model_name}}Service struct {
	lgr             *zap.SugaredLogger
	{{cookiecutter.model_name}}Repo      repository.{{cookiecutter.model_name.capitalize()}}Repository
	createValidator *validator.Validate
}

// New{{cookiecutter.model_name.capitalize()}}Service Return a new {{cookiecutter.model_name.capitalize()}} Service
func New{{cookiecutter.model_name.capitalize()}}Service(logger *zap.SugaredLogger, {{cookiecutter.model_name}}Repo repository.{{cookiecutter.model_name.capitalize()}}Repository) {{cookiecutter.model_name.capitalize()}}Service {
	return &{{cookiecutter.model_name}}Service{
		lgr:             logger,
		{{cookiecutter.model_name}}Repo:      {{cookiecutter.model_name}}Repo,
		createValidator: newCreateValidator(),
	}
}

func (ds *{{cookiecutter.model_name}}Service) Create(ctx context.Context, {{cookiecutter.model_name}} model.{{cookiecutter.model_name.capitalize()}}) (model.{{cookiecutter.model_name.capitalize()}}, error) {
	lgr := ds.lgr.With("{{cookiecutter.model_name}}", {{cookiecutter.model_name}})
	lgr.Debug("Create")
	err := ds.createValidator.Struct(&{{cookiecutter.model_name}})
	if err != nil {
		lgr.Errorw("error validating", zap.Error(err))
		return model.{{cookiecutter.model_name.capitalize()}}{}, err
	}

	dvc, err := ds.{{cookiecutter.model_name}}Repo.Upsert(ctx, {{cookiecutter.model_name}})
	if err != nil {
		lgr.Errorw("error saving", zap.Error(err))
		return model.{{cookiecutter.model_name.capitalize()}}{}, err
	}
	return dvc, nil
}

func newCreateValidator() *validator.Validate {
	createValidator := validator.New()
	createValidator.SetTagName(tagValidationCreate)
	return createValidator
}
