package internal

import "github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"

func toModel{{cookiecutter.model_name.capitalize()}}(dvcReq *model.Create{{cookiecutter.model_name.capitalize()}}Request) *model.{{cookiecutter.model_name.capitalize()}} {
	return &model.{{cookiecutter.model_name.capitalize()}}{
		{{cookiecutter.model_name.capitalize()}}Id:      dvcReq.{{cookiecutter.model_name.capitalize()}}Id,
		Model:         dvcReq.Model,
		HwVersion:     dvcReq.HardwareVersion,
		OSVersion:     dvcReq.OperatingSystemVersion,
		AppVersion:    dvcReq.AppVersion,
		ApiLevel:      dvcReq.ApiLevel,
		SecurityPatch: dvcReq.SecurityPatch,
		NFCAvailable:  dvcReq.IsNfcAvailable,
		NFCEnabled:    dvcReq.IsNfcEnabled,
	}
}
