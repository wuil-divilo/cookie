package internal

import "github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"

type handlerRequest struct {
	UserID                    string `eventmapper:"request.authorizer.username" validate:"required"`
	model.CreateDeviceRequest `eventmapper:"request.body"`
}
