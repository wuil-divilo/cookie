package service

//go:generate mockgen -source ./device.go -destination ../testing/service/mock/device.go -package mock

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

// {{cookiecutter.model_name}}ervice Device Service Interface
type {{cookiecutter.model_name}}ervice interface {
	Create(ctx context.Context, device model.Device) (model.Device, error)
}

type {{cookiecutter.model_name}}ervice struct {
	lgr             *zap.SugaredLogger
	deviceRepo      repository.DeviceRepository
	createValidator *validator.Validate
}

// New{{cookiecutter.model_name}}ervice Return a new Device Service
func New{{cookiecutter.model_name}}ervice(logger *zap.SugaredLogger, deviceRepo repository.DeviceRepository) {{cookiecutter.model_name}}ervice {
	return &{{cookiecutter.model_name}}ervice{
		lgr:             logger,
		deviceRepo:      deviceRepo,
		createValidator: newCreateValidator(),
	}
}

func (ds *{{cookiecutter.model_name}}ervice) Create(ctx context.Context, device model.Device) (model.Device, error) {
	lgr := ds.lgr.With("device", device)
	lgr.Debug("Create")
	err := ds.createValidator.Struct(&device)
	if err != nil {
		lgr.Errorw("error validating", zap.Error(err))
		return model.Device{}, err
	}

	dvc, err := ds.deviceRepo.Upsert(ctx, device)
	if err != nil {
		lgr.Errorw("error saving", zap.Error(err))
		return model.Device{}, err
	}
	return dvc, nil
}

func newCreateValidator() *validator.Validate {
	createValidator := validator.New()
	createValidator.SetTagName(tagValidationCreate)
	return createValidator
}
